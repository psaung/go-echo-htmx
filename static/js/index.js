htmx.defineExtension("reset-on-change", {
  onEvent: function (name, event) {
    console.log(this);
    console.log(name, event);
  },
});
