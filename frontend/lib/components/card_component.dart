import 'package:flutter/material.dart';

class CardComponent extends StatelessWidget {
  const CardComponent({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(15.0),
      child: Container(
        height: 150,
        width: 150,
        decoration: BoxDecoration(borderRadius: BorderRadius.circular(20),
        color: const Color.fromARGB(255, 85, 85, 85)),
        child: Padding(
          padding: const EdgeInsets.all(8.0),
          child: Column(crossAxisAlignment: CrossAxisAlignment.start,
            children: [
            Text("Title",style: TextStyle(fontSize: 25),),
            Padding(
              padding: const EdgeInsets.all(15.0),
              child: Text("This is a trial card, with trial description of about 1 line"),
            )
          ],),
        ),
      ),
    );
  }
}