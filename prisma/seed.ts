import { PrismaClient } from "@prisma/client";
const prisma = new PrismaClient();

async function main() {
  const roles = await prisma.role.createMany({
    data: [
      {
        name: "Admin",
        description: "All allowed",
      },
      {
        name: "Moderator",
        description: "Some roles are allowed",
      },
      {
        name: "User",
        description: "No roles",
      },
    ],
  });
}

main();
