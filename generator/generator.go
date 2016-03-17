package main

// Product
// Cart
// Login

// func main() {
//   generator.RegisterTheme("Name") // views, configuration, relies

//   app := generator.New()
//   app.UseTheme("shop") // -> product, cart
//   app.Generate()

//   // generate a website render views, layout (looking shared views)
// }

// ////////////////////////////////////////////////////////////////////////////////

//   // generator
//   // generate a layout
//   // generate a controller with auth
//   generator.Generate(&generator.Application{
//     Components: []generator.Component{ // AfterGenerate, RegisterRoute("POST", "/cart", struct{Name string, Code string})
//       {
//         Name: "Product",
//         Options: []string{"Add To Cart"},
//       }, {
//         Name: "Cart",
//       }, {
//         Name: "Auth",
//       },
//     },
//     Entry: "product/index",
//   })
// }

// // Product
// RegisterWidget(Widget{
//   Name: "product/index",
//   Template: "product_index",
//   Setting  *admin.Resource{},
//   Context  func(context Context, setting interface{}) Context {
//   },
// })

// RegisterWidget(Widget{
//   Name: "product/show",
//   Template: "product_show",
//   Setting  *admin.Resource{},
//   Context  func(context Context, setting interface{}) Context {
//   },
// })
