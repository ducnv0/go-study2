package main

import "fmt"

func main() {
	// ==========================================
	// 1. VARIABLES (Type Inference)
	// ==========================================
	// Python: name = "Alice"
	// Go uses `:=` to completely infer the type upon creation.
	name := "Alice" 
	
	// Python: age: int = 30
	// Go Explicit Declaration (you rarely need this unless setting defaults)
	var age int = 30
	
	// Similar to Python's f-string formatting, Go uses `Printf` 
	fmt.Printf("%s is %d years old.\n\n", name, age)


	// ==========================================
	// 2. SLICES (Python Lists/Arrays)
	// ==========================================
	// Python: nums = [1, 2, 3]
	// In Go, arrays have a fixed size. So we use "Slices" `[]type` which can grow!
	nums := []int{1, 2, 3} 
	
	// Python: nums.append(4)
	// Go's append returns a NEW updated slice, so you must re-assign it!
	nums = append(nums, 4) 
	
	fmt.Println("Slice of numbers:", nums)


	// ==========================================
	// 3. MAPS (Python Dicts)
	// ==========================================
	// Python: user = {"name": "Bob", "role": "admin"}
	// In Go, Maps require strict types for the Keys and the Values.
	// map[KeyType]ValueType
	user := map[string]string{
		"name": "Bob",
		"role": "admin",
	}
	
	// Access and updating works exactly like Python!
	user["role"] = "superadmin"
	
	fmt.Println("User Name:", user["name"])
	fmt.Println("Full Dictionary:", user)
}
