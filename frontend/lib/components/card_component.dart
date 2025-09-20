import 'package:flutter/material.dart';
import 'package:any_link_preview/any_link_preview.dart';

class CardComponent extends StatelessWidget {
  const CardComponent({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(15.0),
      child: Container(
        // height: 500,
        width: 150,
        decoration: BoxDecoration(borderRadius: BorderRadius.circular(20),
        color: const Color.fromARGB(255, 85, 85, 85)),
        
        child: Padding(
          padding: const EdgeInsets.all(10.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
            Row(
              children: [
                const Expanded(
                  child: Text(
                    "Title",
                    style: TextStyle(fontSize: 27),
                  ),
                ),

                PopupMenuButton<String>(
                  onSelected: (value) {
                    // Add logic here
                    if (value == 'favorite') {
                      print("Favorited!");
                    } else if (value == 'delete') {
                      print("Deleted!");
                    }
                  },
                  itemBuilder: (BuildContext context) => <PopupMenuEntry<String>>[
                    const PopupMenuItem<String>(
                      value: 'favorite',
                      child: ListTile(
                        leading: Icon(Icons.favorite_border),
                        title: Text('Favorite'),
                      ),
                    ),
                    const PopupMenuItem<String>(
                      value: 'delete',
                      child: ListTile(
                        leading: Icon(Icons.delete_outline),
                        title: Text('Delete'),
                      ),
                    ),
                  ],
                ),
              ],
            ),
            
            Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
              SizedBox(
              height: 100,
              width: 150,
              child: AnyLinkPreview(
                link: "https://medium.com/@kmdkhadeer/docker-get-started-9aa7ee662cea",
                displayDirection: UIDirection.uiDirectionHorizontal,
                cache: const Duration(hours: 1),),
            ),

            Flexible(
              child: Padding(
                padding: const EdgeInsets.all(15.0),
                child: Text("This is a trial card, with trial description of about 1 line"),
              ),
            )

            ],)
          ],),
        ),
      ),
    );
  }
}