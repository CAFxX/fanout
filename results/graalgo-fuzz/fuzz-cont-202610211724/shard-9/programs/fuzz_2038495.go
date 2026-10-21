package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint32 = 151
    var v2 string = "hello"
    var v3 string = "foo bar"
    v4 := GenS1{F0: 1}
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: 2}
        var v7 int64 = 0
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 113}
    v9 := []int{0, 32767, 128, 31}
    var v10 uint = 4294967295
    var v11 float64 = (-2.25 / 3.14159)
    var v12 bool = false
    var v13 int64 = -1
    for fi14 := 0; fi14 < 10; fi14++ {
        v12 = v12
        v2 = (v3 + (v2 + ""))
        v2 = ((v3 + ((v2 + (v2 + "")) + ((v2 + "foo bar") + "foo bar"))) + v2)
        _ = fi14
    }
    var v15 int = 9223372036854775807
    v16 := GenS1{F0: 356}
    fmt.Println(v1, v1)
    var v17 bool = true
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
