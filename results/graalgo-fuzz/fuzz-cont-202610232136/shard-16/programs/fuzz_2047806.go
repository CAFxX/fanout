package main

import (
    "fmt"
)


func main() {
    var v1 int64 = -1
    var v2 float32 = (0.1 / (float32(1.5) * 1.5))
    var v3 uint = 984
    var v4 float32 = v2
    var v5 int32 = -1
    v4 -= 3.14159
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 float32 = v2
        fmt.Println(v4, v5, v1)
        _ = fi6
        _ = v7
    }
    v2 = v4
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{(fi8 << ((2) & 63)), (fi8 - 4294967295)}
        v9 = append(v9, 42)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{fi10}
        v5 += int32(v1)
        _ = fi10
        _ = v11
    }
    v1 -= -873
    var v12 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
