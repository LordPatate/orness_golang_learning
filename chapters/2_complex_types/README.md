# Complex types

## Slices

- A slice has both a length and a capacity.
- The length of a slice is the number of elements it contains.
- The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
- The length and capacity of a slice s can be obtained using the expressions `len(s)` and `cap(s)`.
- The zero value of a slice is nil.
- A nil slice has a length and capacity of 0 and has no underlying array. 
- The `make` function allocates a zeroed array and returns a slice that refers to that array.
- The `append` built-in function appends elements to the end of a slice. 
- If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
- The `range` form of the `for` loop iterates over a slice.

## Maps

- Like Python dict or Java HashMap.
- Initialize maps with a literal expression or `make`.
- Unlike slices, no values can be added from the `nil` map (the zero value).
- Test that a key is present with a two-value assignment.
- Delete from maps with the `delete` builtin.
- The `range` keyword iterates over maps too.

## Structs

- Very similar to C structs.
- When using a pointer to a struct, the fields can be accessed directly with the dot notation without explicit dereference.
- Omitted values from the literal declaration are given their zero values.