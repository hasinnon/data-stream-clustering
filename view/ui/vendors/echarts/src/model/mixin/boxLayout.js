define(function (require) {

    return {
        getBoxLayoutParams: function () {
            return {
                right: this.get('right'),
                top: this.get('top'),
                left: this.get('left'),
                bottom: this.get('bottom'),
                width: this.get('width'),
                height: this.get('height')
            };
        }
    };
});