var CommentCountBox = React.createClass({
    getInitialState: function () {
        return {commentCount: 0};
    },
    loadCommentFromServer: function () {
        var currentPage = "sort-in-golang";
        jQuery.ajax({
            url: this.props.commentFetch,
            dataType: "json",
            cache: false,
            type: "post",
            data: currentPage,
            success: function (data) {
                console.log(data.commentCount)
                this.setState({commentCount: data.commentCount});
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
    componentDidMount: function () {
        this.loadCommentFromServer();
        // setInterval(this.loadCommentFromServer, this.props.pollInterval);
    },
    render: function () {
        return (
            <span className='meta_comments'>
                <a className='comment-link' href='#'>{this.state.commentCount} comments</a>
            </span>
        );
    }
});

ReactDOM.render(
    <CommentCountBox commentFetch="/api/comments/fetch" pollInterval={5000}/>,
    document.getElementById("commentCount7")
);