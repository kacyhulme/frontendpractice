package actions

import (

  "fmt"
  "net/http"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "github.com/gobuffalo/x/responder"
  "frontendpractice/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Post)
// DB Table: Plural (posts)
// Resource: Plural (Posts)
// Path: Plural (/posts)
// View Template Folder: Plural (/templates/posts/)

// PostsResource is the resource for the Post model
type PostsResource struct{
  buffalo.Resource
}

// List gets all Posts. This function is mapped to the path
// GET /posts
func (v PostsResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  posts := &models.Posts{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Posts from the DB
  if err := q.All(posts); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // Add the paginator to the context so it can be used in the template.
    c.Set("pagination", q.Paginator)

    c.Set("posts", posts)
    return c.Render(http.StatusOK, r.HTML("/posts/index.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(posts))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(posts))
  }).Respond(c)
}

// Show gets the data for one Post. This function is mapped to
// the path GET /posts/{post_id}
func (v PostsResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Post
  post := &models.Post{}

  // To find the Post the parameter post_id is used.
  if err := tx.Find(post, c.Param("post_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    c.Set("post", post)

    return c.Render(http.StatusOK, r.HTML("/posts/show.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(post))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(post))
  }).Respond(c)
}

// New renders the form for creating a new Post.
// This function is mapped to the path GET /posts/new
func (v PostsResource) New(c buffalo.Context) error {
  c.Set("post", &models.Post{})

  return c.Render(http.StatusOK, r.HTML("/posts/new.plush.html"))
}
// Create adds a Post to the DB. This function is mapped to the
// path POST /posts
func (v PostsResource) Create(c buffalo.Context) error {
  // Allocate an empty Post
  post := &models.Post{}

  // Bind post to the html form elements
  if err := c.Bind(post); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(post)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the new.html template that the user can
      // correct the input.
      c.Set("post", post)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/posts/new.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "post.created.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/posts/%v", post.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.JSON(post))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.XML(post))
  }).Respond(c)
}

// Edit renders a edit form for a Post. This function is
// mapped to the path GET /posts/{post_id}/edit
func (v PostsResource) Edit(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Post
  post := &models.Post{}

  if err := tx.Find(post, c.Param("post_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  c.Set("post", post)
  return c.Render(http.StatusOK, r.HTML("/posts/edit.plush.html"))
}
// Update changes a Post in the DB. This function is mapped to
// the path PUT /posts/{post_id}
func (v PostsResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Post
  post := &models.Post{}

  if err := tx.Find(post, c.Param("post_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  // Bind Post to the html form elements
  if err := c.Bind(post); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(post)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the edit.html template that the user can
      // correct the input.
      c.Set("post", post)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/posts/edit.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "post.updated.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/posts/%v", post.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(post))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(post))
  }).Respond(c)
}

// Destroy deletes a Post from the DB. This function is mapped
// to the path DELETE /posts/{post_id}
func (v PostsResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Post
  post := &models.Post{}

  // To find the Post the parameter post_id is used.
  if err := tx.Find(post, c.Param("post_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(post); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a flash message
    c.Flash().Add("success", T.Translate(c, "post.destroyed.success"))

    // Redirect to the index page
    return c.Redirect(http.StatusSeeOther, "/posts")
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(post))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(post))
  }).Respond(c)
}
