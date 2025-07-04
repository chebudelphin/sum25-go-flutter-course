import 'package:flutter/material.dart';

class CounterApp extends StatefulWidget {
  const CounterApp({Key? key}) : super(key: key);

  @override
  State<CounterApp> createState() => _CounterAppState();
}

class _CounterAppState extends State<CounterApp> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      _counter++;
    });
  }

  void _decrementCounter() {
    setState(() {
      _counter--;
    });
  }

  void _resetCounter() {
    setState(() {
      _counter= 0;
    });   
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Counter App'),
        actions: [
          // TODO: add a refresh button with Icon(Icons.refresh)
          //comment to have just something in the branch
          //why did i do lab in main? I'm also don't know
          IconButton(onPressed: _resetCounter, icon: Icon(Icons.refresh))
        ],
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              '$_counter',
              style: const TextStyle(fontSize: 48),
            ),
            const SizedBox(height: 32),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                // TODO: add a decrement button with Icon(Icons.remove) and onPressed: _decrementCounter
                FloatingActionButton(onPressed: _decrementCounter, child: const Icon(Icons.remove)),
                const SizedBox(width: 32),
                // TODO: add a increment button with Icon(Icons.add) and onPressed: _incrementCounter
                FloatingActionButton(onPressed: _incrementCounter, child: const Icon(Icons.add))                
              ],
            ),
          ],
        ),
      ),
    );
  }
}