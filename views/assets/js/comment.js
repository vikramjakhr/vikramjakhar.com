var CommentBox = React.createClass({
    getInitialState: function () {
        return {data: [], commentCount:0};
    },
    loadCommentFromServer: function () {
        var currentPage = jQuery("#currentPage").val();
        jQuery.ajax({
            url: this.props.commentFetch,
            dataType: "json",
            cache: false,
            type: "post",
            data: currentPage,
            success: function (data) {
                this.setState({data: data.comments, commentCount:data.commentCount});
                jQuery("#commentCountLink").text(this.state.commentCount + " comments");
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
    componentDidMount: function () {
        this.loadCommentFromServer();
        setInterval(this.loadCommentFromServer, this.props.pollInterval);
    },
    handleCommentSubmit: function (comment) {
        var comments = this.state.data;
        comment.id = Date.now();
        var newComments = comments.concat([comment]);
        this.setState({data: newComments, commentCount:this.state.commentCount+1});
        jQuery(".commentTxtArea").val("");
        jQuery("#commentCountLink").text(this.state.commentCount + " comments");
        var currentPage = jQuery("#currentPage").val();
        jQuery.ajax({
            url: this.props.commentPost,
            dataType: "json",
            cache: false,
            data: JSON.stringify({text: comment.text, id: comment.id, author: "Guest", postName: currentPage}),
            method: "POST",
            success: function (data) {
                this.setState({data: data.comments, commentCount:data.commentCount});
            }.bind(this),
            error: function (xhr, status, err) {
                this.setState({data: comments, commentCount:this.state.commentCount});
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
    render: function () {
        return (
            <div className="commentBox">
                <h4>{this.state.commentCount} comments:</h4>
                <CommentList data={this.state.data}/>
                <CommentForm onCommentSubmit={this.handleCommentSubmit}/>
            </div>
        );
    }
});

var CommentList = React.createClass({
    render: function () {
        var commentNodes = this.props.data.map(function (comment) {
            return (
                <Comment author={comment.author} key={comment.id}>
                    {comment.text}
                </Comment>
            );
        });

        return (
            <div className='commentList'>
                {commentNodes}
            </div>
        );
    }
});

var CommentForm = React.createClass({
    getInitialState: function () {
        return ({text: ''});
    },
    handleTextChange: function (event) {
        this.setState({text: event.target.value});
    },
    handleKeyPress: function (event) {
        this.setState({text: event.target.value});
        var keycode = (event.keyCode ? event.keyCode : event.which);
        if (keycode == '13') {
            var text = this.state.text.trim();
            if (!text) {
                return;
            } else {
                this.props.onCommentSubmit({text: text});
                this.setState({text: ""});
            }
        }
    },
    render: function () {
        return (
            <div id="commentForm">
                <div className="user-thumbnail">
                    <img className="user-thumb" src="/assets/images/user.png"/>
                </div>
                <div className="comment-txtArea-div">
                    <textarea className="commentTxtArea" placeholder="Add your comment here"
                              onKeyDown={this.handleKeyPress}
                              onChange={this.handleTextChange} value={this.state.text}></textarea>
                </div>
            </div>
        );
    }
});

var Comment = React.createClass({
    render: function () {
        return (
            <div>
                <div className="user-thumbnail">
                    <img className="user-thumb" src="/assets/images/user.png"/>
                </div>
                <div className="aComment">
                    <a href="javascript:void(0)">{this.props.author} </a>
                    {this.props.children}
                </div>
            </div>
        );
    }
});

ReactDOM.render(
    <CommentBox commentFetch="/api/comments/fetch" commentPost="/api/comments/post" pollInterval={2000}/>,
    document.getElementById("comments")
);