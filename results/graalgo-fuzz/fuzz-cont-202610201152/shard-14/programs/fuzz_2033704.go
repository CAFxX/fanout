package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int16; F1 bool }
func genFn0(p0 uint8, p1 bool) int8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{}
        if len(v7) > 0 {
            v7[0] = 42
        }
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: "a"}
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return 33
}
func genFn1() int8 {
    var v10 int32 = 42
    var v11 float32 = (((float32((((float64(1.5) * 3.14159) - 1.5) + 1.5)) - 1e10) / 0.1) + 0.0)
    _ = v10
    _ = v11
    return -29
}

func main() {
    var v1 uint8 = 2
    var v2 int16 = 1
    var v3 uint8 = v1
    v4 := make(map[int]int)
    var v5 uint32 = 1
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int = fi12
        fmt.Println(v1, v13, v3)
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: "\n\t"}
    var v15 int64 = 1
    var v16 string = ("\n\t" + "hello")
    v2 += (v2 - (v2 << ((541) & 63)))
    if (true && (true && (false && true))) {
        if (v2 >= v2) {
            switch 747 {
            case -2147483648:
                _ = len(v4)
            case 9223372036854775807:
                v1 -= 205
            default:
                v17 := []int{}
                _ = v17
            }
            v5 = 42
        }
        v1 = v1
    }
    if (v2 >= 42) {
        v18 := GenS2{F0: (("a" + v16) + (v16 + v16)), F1: v5, F2: (^(v3))}
        _ = v18
    } else {
        fmt.Println(v16, v15, v16)
    }
    switch (("hello" + (v16 + v16)) + v16) {
    case "hello":
        fmt.Println(v16, v3, v1)
    case "foo bar":
        var v19 string = (v16 + ("\n\t" + ""))
        _ = v19
    case "\n\t":
        fmt.Println(v1)
    }
    v20 := GenS3{F0: 1000, F1: (4294967295 <= 436)}
    v21 := []int{127}
    _ = len(v21)
    for fi22 := 0; fi22 < 8; fi22++ {
        fmt.Println(v2)
        v1 -= v3
        _ = fi22
    }
    v5 = (v5 << ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
