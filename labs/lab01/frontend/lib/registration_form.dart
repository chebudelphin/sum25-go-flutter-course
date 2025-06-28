import 'package:flutter/material.dart';

class RegistrationForm extends StatefulWidget {
  const RegistrationForm({Key? key}) : super(key: key);

  @override
  State<RegistrationForm> createState() => _RegistrationFormState();
}

class _RegistrationFormState extends State<RegistrationForm> {
  final _formKey = GlobalKey<FormState>();
  final _name = TextEditingController();
  final _email = TextEditingController();
  final _password = TextEditingController();
  bool _success = false;

  @override
  void dispose() {
    _name.dispose();
    _email.dispose();
    _password.dispose();
    super.dispose();
  }

  void _submit() {
    setState(() => _success = false);
    if (_formKey.currentState!.validate()) {
      setState(() => _success = true);
    }
  }

  String? _nameValidator(String? v) =>
      (v == null || v.isEmpty) ? 'Please enter your name' : null;

  String? _emailValidator(String? v) {
    final re = RegExp(r'^[^@\s]+@[^@\s]+\.[^@\s]+$');
    return (v == null || !re.hasMatch(v))
        ? 'Please enter a valid email'
        : null;
  }

  String? _passValidator(String? v) =>
      (v == null || v.length < 6)
          ? 'Password must be at least 6 characters'
          : null;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Form(
          key: _formKey,
          child: Column(
            children: [
              TextFormField(
                key: const Key('name'),
                controller: _name,
                decoration: const InputDecoration(labelText: 'Name'),
                validator: _nameValidator,
              ),
              TextFormField(
                key: const Key('email'),
                controller: _email,
                decoration: const InputDecoration(labelText: 'Email'),
                validator: _emailValidator,
              ),
              TextFormField(
                key: const Key('password'),
                controller: _password,
                decoration: const InputDecoration(labelText: 'Password'),
                obscureText: true,
                validator: _passValidator,
              ),
              const SizedBox(height: 12),
              ElevatedButton(
                onPressed: _submit,
                child: const Text('Submit'),
              ),
            ],
          ),
        ),
        if (_success)
          const Padding(
            padding: EdgeInsets.only(top: 12),
            child: Text(
              'Registration successful!',
              style: TextStyle(color: Colors.green),
            ),
          ),
      ],
    );
  }
}
