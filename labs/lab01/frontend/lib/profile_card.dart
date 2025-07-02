import 'package:flutter/material.dart';

class ProfileCard extends StatelessWidget {
  final String name;
  final String email;
  final int age;
  final String? avatarUrl;

  const ProfileCard({
    super.key,
    required this.name,
    required this.email,
    required this.age,
    this.avatarUrl,
  });

  @override
  Widget build(BuildContext context) {
<<<<<<< HEAD
    final avatar = (avatarUrl != null && avatarUrl!.isNotEmpty)
        ? CircleAvatar(backgroundImage: NetworkImage(avatarUrl!))
        : CircleAvatar(child: Text(name.isNotEmpty ? name[0] : '?'));

    return Card(
      elevation: 2,
      child: Padding(
        padding: const EdgeInsets.all(12),
        child: Row(
          children: [
            avatar,
            const SizedBox(width: 12),
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(name,
                    style: const TextStyle(
                        fontSize: 16, fontWeight: FontWeight.bold)),
                Text(email),
                Text('Age: $age'),
              ],
            ),
=======
    return Card(
      margin: const EdgeInsets.all(16.0),
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            // TODO: add a CircleAvatar with radius 50 and backgroundImage NetworkImage(avatarUrl!) if url is not null and text name[0].toUpperCase() if url is null
            
            const SizedBox(height: 16),
            // TODO: add a Text with name and style fontSize: 24, fontWeight: FontWeight.bold
           
            const SizedBox(height: 8),
            // TODO: add a Text with Age: $age and style fontSize: 16
           
            const SizedBox(height: 8),
            // TODO: add a Text with email and style fontSize: 16, color: Colors.grey
            
>>>>>>> e6d76f7 (update lab1 and workflow of submission)
          ],
        ),
      ),
    );
  }
}
