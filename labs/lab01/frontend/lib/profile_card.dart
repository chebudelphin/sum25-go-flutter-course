import 'package:flutter/material.dart';

class ProfileCard extends StatelessWidget {
  final String name;
  final String email;
  final int age;
  final String? avatarUrl;

  const ProfileCard({
    Key? key,
    required this.name,
    required this.email,
    required this.age,
    this.avatarUrl,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
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
          ],
        ),
      ),
    );
  }
}
