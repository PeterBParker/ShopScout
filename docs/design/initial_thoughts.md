# ShopScout Application Thoughts

## Purpose
ShopScout should make shopping in a new grocery store a quick and (relatively) painless process. ShopScout can make buying new items at a grocery store easy to add to your typical route without a lot of backtracking. ShopScout should be the MapQuest for grocery stores. The audience is anyone who ever shops at a new grocery store either looking for something their typical store doesn't stock, moving to/traveling through a new town, or just wants to spend less time in the store.

## How
ShopScout will start by leveraging the data provided by Kroger APIs to provide a route through the store that lets a person make a single pass through the store and get all the stuff they need. This will mean that, initially, only Kroger-owned stores will be supported. Depending on the store the data may or may not be stale, so user experience may differ. If the Kroger version proves useful, I can explore adding other store APIs as they become available. I should write the source so that the actual API used should be abstracted from the main logic.

## User Story
Users can tell ShopScout what products they want to add to their list (can they save lists?)
Users will tell ShopScout which store from a list/map view (defaults to closest store to user's location)
ShopScout will query the locations of all the items and do its magic to determine a route that results in a single pass
ShopScout will provide a list of the items to get in order of the route and their locations. Two possible views. One is a complete, ordered checklist, and the other is one item at a time that when it is checked off stacks up the bottom space and displays the next, like a stack. They can click that they got the item or if they couldn't find the item. Eventually a mapped route view would be great.
When they are done, if all their items were gotten we'll display a green checkmark. If not, ???

## Bonus Features
Gamify the shopping so that users can gain experience from successful runs and earn funny titles and badges displaying their experience with specific store locations.
