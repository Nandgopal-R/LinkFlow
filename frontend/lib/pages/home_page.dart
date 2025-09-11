import 'package:flutter/material.dart';
import 'package:frontend/components/card_component.dart';
import 'package:google_fonts/google_fonts.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {

// Dialog box for adding new blog
  void createNote(){
    showDialog(
      context: context, 
      builder: (context) => AlertDialog(
        title: Text("Add New Blog"),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: [
            TextField(
              decoration: InputDecoration(border: OutlineInputBorder(borderRadius: BorderRadius.circular(20)), 
              hintText: "Title"),
            ),
            SizedBox(height: 10,),
            TextField(
              decoration: InputDecoration(border: OutlineInputBorder(borderRadius: BorderRadius.circular(20)),
                hintText: "Description"),
            ),
            SizedBox(height: 10,),
            TextField(decoration: InputDecoration(border: OutlineInputBorder(borderRadius: BorderRadius.circular(20)),
              hintText: "link"),),
            SizedBox(height: 10,),
          ],
        ),
        actions: [ElevatedButton(onPressed: (){
          Navigator.pop(context);
        },
        child: Text("Add"),)],
      ));
  }


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: false,
      appBar: AppBar(centerTitle: true,title: Text("LinkFlow", style:GoogleFonts.dmSerifText(
        fontSize: 40,
      ),
      ),
      ),
      drawer: Drawer(
        child: Column(
          children: [
            DrawerHeader(child: Text("LinkFLow",style: TextStyle(fontSize: 30),)),
           ListTile(
            leading: Icon(Icons.home),
            title: Text("Home"),
            onTap: (){},
           ),
           ListTile(
            leading: Icon(Icons.star),
            title: Text("Favourites"),
            onTap: () {},
           )
          ],
        ),
      ),
      
      body: ListView(children: [
        CardComponent(),
        CardComponent(),
        CardComponent()
      ],),
      floatingActionButton: FloatingActionButton(onPressed: createNote,
      child: Icon(Icons.add, color: Colors.black, size: 30,),
      backgroundColor: Colors.white,
      shape: CircleBorder(),)
    );
  }
}