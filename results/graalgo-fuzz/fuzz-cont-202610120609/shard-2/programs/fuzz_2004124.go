package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }
func genFn0(p0 int16, p1 int, p2 int8) int64 {
    var v6 int8 = p2
    p0 = p0
    var v7 uint8 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return int64(v6)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 42
    v3 := make(map[string]uint32)
    v4 := []uint32{82, 2, 30}
    var v5 bool = false
    switch -88 {
    case 42:
        fmt.Println(v5)
    case -1:
        v8 := GenS1{F0: 1000}
        _ = v8
    default:
        v9 := GenS2{F0: genFn0(-858, -9223372036854775808, 127), F1: 4294967295}
        _ = v9
    }
    var v10 bool = (!v5)
    v5 = v5
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: 1000}
        switch ("hello" + "\n\t") {
        case "":
            var v13 float32 = float32(0.1)
            _ = v13
        case "\n\t":
            v5 = (-128 == -22)
        default:
            var v14 float32 = -2.25
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    v15 := GenS3{F0: (!v10)}
    v16 := GenS2{F0: (v2 >> ((42) & 63)), F1: 1}
    v17 := GenS3{F0: false}
    v2 -= (v2 & v2)
    fmt.Println(v10, v2, v2)
    if (v5 && false) {
        v18 := GenS4{F0: -1}
        var v19 int = 9223372036854775807
        v10 = v5
        _ = v18
        _ = v19
    } else {
        var v20 uint = 2
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 float32 = 3.14159
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 12; fi23++ {
            v4 = append(v4, 42)
            switch ("a" + "a") {
            case "":
                var v24 uint32 = 65535
                _ = v24
            case "hello":
                var v25 int32 = -2147483648
                _ = v25
            }
            _ = fi23
        }
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
