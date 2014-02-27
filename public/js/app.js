App = Ember.Application.create();

App.Router.map(function() {
  // put your routes here
});

App.IndexRoute = Ember.Route.extend({
  model: function() {
    return [
      {title: "Flight Price Checker", description: "Something trivial, everyone does this...", url: ""},
      {title: "Room Finder", description: "Same thing, sigh... no good idea.", url: ""},
      {title: "Coming soon 1", description: "Still thinking, what might be not too bad.", url: ""},
      {title: "Coming soon 2", description: "Just to fill up space", url: ""}];
  }
});
