package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int16, p1 string, p2 int) int8 {
    v5 := GenS1{F0: 2}
    v6 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]bool)
    var v3 uint16 = 65535
    var v4 float32 = v1
    v7 := GenS1{F0: (v3 / ((v3) | 1))}
    v8 := GenS1{F0: v3}
    for fi9 := 0; fi9 < 4; fi9++ {
        v4 += v1
        var v10 int64 = 42
        _ = fi9
        _ = v10
    }
    if (false && (false && true)) {
        v11 := GenS2{F0: 0}
        fmt.Println(v4, v4, v3)
        _ = v11
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 float32 = 3.14159
            _ = fi12
            _ = v13
        }
        var v14 uint = 690
        _ = v14
    }
    v15 := []int{4294967295}
    v16 := GenS1{F0: 42}
    v3 -= v3
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS2{F0: 42}
        _ = fi17
        _ = v18
    }
    if false {
        fmt.Println(v3)
        fmt.Println(v4, v4)
        v19 := GenS1{F0: v3}
        _ = v19
    } else {
        switch ("foo bar" + "foo bar") {
        case "hello":
            v20 := GenS2{F0: -227}
            _ = v20
        case "a":
            v2[32767] = false
        default:
            v3 -= (^(v3))
        }
        var v21 uint8 = 142
        v21 -= v21
        _ = v21
    }
    v4 = v4
    for fi22 := 0; fi22 < 6; fi22++ {
        v15 = append(v15, fi22)
        v23 := GenS2{F0: 1}
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
