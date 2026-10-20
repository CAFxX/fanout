package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint16) int64 {
    p1 -= uint16(p0)
    var v4 int64 = 0
    v5 := []int{4294967295}
    if (((3.14159 - 1.5) / ((0.1 / 0.1) / 3.14159)) >= 0.1) {
        switch 721 {
        case -870:
            if (!(true && false)) {
                v5 = append(v5, -1)
            }
        default:
            p1 = (p1 << ((4294967295) & 63))
        }
        var v6 string = "foo bar"
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (v4 * v4)
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]int)
    var v3 float32 = (1e10 + -2.25)
    v7 := GenS1{F0: 72}
    var v8 int = -2147483648
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v8, v1)
        v10 := GenS1{F0: 127}
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 47}
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 uint = 4294967295
        _ = fi12
        _ = v13
    }
    var v14 int = 128
    v15 := []int{}
    switch "foo bar" {
    case "hello":
        if true {
            v8 -= 9223372036854775807
        } else {
            if (true || (true && (-32768 != -32768))) {
                v16 := GenS2{F0: 32767}
                _ = v16
            } else {
                v17 := GenS1{F0: -34}
                _ = v17
            }
            v1 -= v1
        }
    default:
        var v18 int64 = 1
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
