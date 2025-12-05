// Import the necessary discord.js classes
import { Client, Events, GatewayIntentBits } from 'discord.js';

// Create a new client instance
const client = new Client({ intents: [GatewayIntentBits.Guilds] });

// When the client is ready, run this code (only once).
// The distinction between `client: Client<boolean>` and `readyClient: Client<true>` is important for TypeScript developers.
// It makes some properties non-nullable.
client.once(Events.ClientReady, (readyClient) => {
	console.log(`Ready! Logged in as ${readyClient.user.tag}`);
});

// Log in to Discord with your client's token
// Bun automatically loads .env file, so we can access process.env directly
const token = process.env.TOKEN || process.env.DISCORD_TOKEN;
if (!token) {
	throw new Error('Discord token not found in environment variables. Please set TOKEN or DISCORD_TOKEN in your .env file.');
}
client.login(token);