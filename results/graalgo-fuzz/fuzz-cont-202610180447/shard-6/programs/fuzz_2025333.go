package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint8 {
    switch -413 {
    case -9223372036854775808:
        var v7 bool = true
        _ = v7
    }
    v8 := []int{}
    var v9 uint32 = 996
    v8 = append(v8, -796)
    _ = v8
    _ = v9
    return 0
}
func genFn1() int8 {
    switch 0 {
    case 9223372036854775807:
        v10 := GenS1{F0: 265}
        _ = v10
    case 4294967295:
        var v11 int32 = -1
        _ = v11
    case 0:
        var v12 uint = 883
        _ = v12
    default:
        if true {
            v13 := []int{589, -2147483648}
            var v14 uint16 = 326
            _ = v13
            _ = v14
        }
    }
    var v15 int32 = 42
    var v16 float32 = 3.14159
    v17 := GenS1{F0: 9223372036854775807}
    _ = v15
    _ = v16
    _ = v17
    return 0
}

func main() {
    v1 := []string{}
    v2 := make(map[string]int)
    var v3 int32 = 0
    var v4 uint32 = 429
    v5 := []int{-371, -844, 2147483647, 4294967295}
    var v6 string = "a"
    switch v6 {
    case "foo bar":
        fmt.Println(v4)
    }
    switch "a" {
    case "\n\t":
        v6 = "a"
    }
    if false {
        v3 += v3
        v4 -= 345
        var v18 int8 = 1
        _ = v18
    } else {
        v19 := GenS1{F0: -371}
        _ = v19
    }
    var v20 uint64 = 18446744073709551615
    _ = len(v2)
    v3 -= 1
    switch 32768 {
    case 608:
        var v21 uint64 = (v20 + (v20 * v20))
        _ = v21
    case 328:
        _ = len(v5)
    case 32767:
        var v22 uint64 = uint64(v4)
        _ = v22
    default:
        fmt.Println(v4, v6)
    }
    var v23 int = 2147483647
    v23 = 0
    var v24 float64 = (1.5 + (0.0 + 1.5))
    for fi25 := 0; fi25 < 5; fi25++ {
        v24 += v24
        _ = fi25
    }
    fmt.Println(v23, v4)
    v20 = v20
    v26 := GenS1{F0: 83}
    _ = len(v5)
    fmt.Println(v23, v6)
    if false {
        for fi27 := 0; fi27 < 12; fi27++ {
            v28 := GenS1{F0: fi27}
            _ = fi27
            _ = v28
        }
        v3 = 554
    }
    v29 := GenS1{F0: v23}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
