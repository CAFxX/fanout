package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 bool }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 int }
func genFn0(p0 uint32, p1 uint64, p2 string) int16 {
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 uint8 = 60
        switch fi4 {
        case 2147483647:
            p2 = ("hello" + (p2 + "hello"))
        case 863:
            switch fi4 {
            case 32767:
                var v6 int8 = 1
                _ = v6
            }
        case 949:
            v7 := GenS1{F0: p2, F1: "foo bar", F2: true}
            _ = v7
        default:
            v8 := GenS1{F0: p2, F1: (p2 + p2), F2: (-219 != -1)}
            _ = v8
        }
        var v9 int16 = -872
        _ = fi4
        _ = v5
        _ = v9
    }
    v10 := GenS2{F0: p1, F1: 1}
    v11 := GenS1{F0: "a", F1: p2, F2: (!false)}
    var v12 uint8 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return -32768
}
func genFn1(p0 int, p1 int16, p2 int16) bool {
    v13 := GenS1{F0: "a", F1: ("a" + "a"), F2: (false || (109 > 42))}
    fmt.Println(p1, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return false
}

func main() {
    var v1 int = -9223372036854775808
    v2 := []uint32{}
    v3 := make(map[int]bool)
    v14 := GenS3{F0: ("" + "hello")}
    v1 += (v1 ^ -2147483648)
    var v15 uint8 = 42
    var v16 int8 = -128
    var v17 uint8 = v15
    switch "a" {
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        var v18 float64 = 1e10
        _ = v18
    }
    switch v1 {
    case 42:
        v16 = (v16 + (v16 & -1))
    case 2147483647:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        switch v1 {
        case 966:
            delete(v3, int(v16))
        default:
            _ = len(v3)
        }
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v15 += v17
        var v20 float32 = (0.1 * (3.14159 + 3.14159))
        _ = fi19
        _ = v20
    }
    var v21 float64 = -2.25
    fmt.Println(v1, v1)
    if len(v2) > 0 {
        v2[0] = 2
    }
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    _ = len(v3)
    _ = len(v2)
    if true {
        v21 = v21
        v15 = (v17 / ((v15) | 1))
    }
    fmt.Println(v1, v1, v17)
    var v22 int64 = -9223372036854775808
    if false {
        var v23 uint64 = 18446744073709551615
        _ = v23
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 uint16 = 2
            _ = fi24
            _ = v25
        }
        v26 := GenS4{F0: (-(v16)), F1: 127}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
