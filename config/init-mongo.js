db.createUser(
    {
      user: "tinyuser",
      pwd: "tinypass",
      roles: [
        {
          role: "readWrite",
          db: "tinydb"
        }
      ]
    }
);
