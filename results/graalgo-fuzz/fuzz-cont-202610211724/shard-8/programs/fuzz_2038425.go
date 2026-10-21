package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 0
    var v3 uint = 42
    switch int(v2) {
    case -453:
        v4 := v1[""]
        _ = v4
    case 859:
        fmt.Println(v2)
    case 25:
        v5 := []int{9223372036854775807, -852}
        _ = v5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        v3 = (v3 << ((2) & 63))
        _ = fi6
    }
    var v7 int64 = 0
    var v8 bool = (4294967295 < 4294967295)
    if v8 {
        fmt.Println(v2)
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v1[("hello" + "")] = 42
        _ = fi9
    }
    var v10 string = "hello"
    v3 -= v3
    if (!((1.5 * -2.25) != 3.14159)) {
        v11 := []int{42, 9223372036854775807, 0}
        _ = v11
    } else {
        v10 = v10
        v12 := v1["foo bar"]
        fmt.Println(v7, v10)
        _ = v12
    }
    var v13 int16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
