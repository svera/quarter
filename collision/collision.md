```csharp
void EntityBase::ResolveCollisions(){
    if(!m_collisions.empty()) {
        std::sort(m_collisions.begin(), m_collisions.end(), SortCollisions);
        Map* gameMap = m_entityManager->GetContext()->m_gameMap;
        unsigned int tileSize = gameMap->GetTileSize();
        for (auto &itr : m_collisions) {
            if (!m_AABB.intersects(itr.m_tileBounds)){
                continue;
            }

            float xDiff = (m_AABB.left + (m_AABB.width / 2)) - (itr.m_tileBounds.left + (itr.m_tileBounds.width / 2));
            float yDiff = (m_AABB.top + (m_AABB.height / 2)) - (itr.m_tileBounds.top + (itr.m_tileBounds.height / 2));

            float resolve = 0
            if(abs(xDiff) > abs(yDiff)) {
                if(xDiff > 0){
                    resolve = (itr.m_tileBounds.left + tileSize) - m_AABB.left;
                } else {
                    resolve = -((m_AABB.left +  m_AABB.width) - itr.m_tileBounds.left);
                }
                Move(resolve, 0)
                m_velocity.x = 0;
                m_collidingOnX = true;
            } else {
                if(yDiff > 0){
                    resolve = (itr.m_tileBounds.top + tileSize) - m_AABB.top;
                } else {
                    resolve = -((m_AABB.top + m_AABB.height) - itr.m_tileBounds.top)
                }
                Move(0, resolve)
                m_velocity.Y = 0;
                if (m_collidingOnY){
                    continue
                }
                m_referenceTile = itr.m_tile;
                m_collidingOnY = true;
            }
        }
        m_collisions.clear();
    }
    if(!m_collidingOnY){
        m_referenceTile = nullptr;
    }
}
```

First, we check if there are any collisions in the container. Sorting of all the elements happensa next. The `std::sort`
function is called and iterators to the beginning and end of the container are passed in, along with the name of the 
function that will do the comparisons between the elements.

The code proceeds to loop over all of the collisions stored in the container. There is
another intersection check here between the bounding box of the entity and the tile.
This is done because resolving a previous collision could have moved an entity in
such a way that it is no longer colliding with the next tile in the container. If there
still is a collision, distances from the center of the entity's bounding box to the center
of the tile's bounding box are calculated. The first purpose these distances serve is
illustrated in the next line, where their absolute values get compared. If the distance
on the x axis is bigger than on the y axis, the resolution takes place on the x axis.
Otherwise, it's resolved on the y axis.

The second purpose of the distance calculation is determining which side of the tile
the entity is on. If the distance is positive, the entity is on the right side of the tile, so
it gets moved in the positive x direction. Otherwise, it gets moved in the negative x
direction. The resolve variable takes in the amount of penetration between the tile and
the entity, which is different based on the axis and the side of the collision.
In the case of both axes, the entity is moved by calling its `Move` method and passing
in the depth of penetration. Killing the entity's velocity on that axis is also important,
in order to simulate the entity hitting a solid. Lastly, the flag for a collision on a
specific axis is set lo `true`.

If a collision is resolved on the y axis, in addition to all the same steps that are taken
in a case of x axis collision resolution, we also check if the flag is set for a y axis
collision. If it hasn't been set yet, we change the `m_referenceTile` data member to
point to the tile type of the current tile the entity is colliding with, which is followed
by that flag getting set to `true` in order to keep the reference unchanged until the next
time collisions are checked. This little snippet of code gives any entity the ability to
behave differently based on which tile it's standing on. For example, the entity can
slide a lot more on ice tiles than on simple grass tiles, as illustrated here: