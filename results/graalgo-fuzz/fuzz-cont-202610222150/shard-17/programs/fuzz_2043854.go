package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32, p1 int64) string {
    switch -9223372036854775808 {
    case -361:
        fmt.Println(p1, p0)
    default:
        v5 := []int{399}
        _ = v5
    }
    var v6 int64 = (p1 | p1)
    if (4294967295 != 817) {
        var v7 int = 96
        _ = v7
    } else {
        var v8 float32 = 0.1
        v6 = (-(v6))
        _ = v8
    }
    switch 380 {
    case 9223372036854775807:
        var v9 bool = true
        _ = v9
    default:
        v10 := []int{1, 998, -455}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = ((-2.25 * -2.25) + 1.5)
    var v3 uint16 = 1
    var v4 uint32 = 0
    v11 := GenS1{F0: -600, F1: true, F2: -1}
    if true {
        v12 := v1[("\n\t" + (("foo bar" + ("a" + "\n\t")) + "hello"))]
        for fi13 := 0; fi13 < 17; fi13++ {
            v3 = v3
            fmt.Println(v3, v4)
            v14 := GenS2{F0: 872, F1: 894}
            _ = fi13
            _ = v14
        }
        v12 += v12
        _ = v12
    } else {
        var v15 uint16 = (v3 - 2)
        _ = v15
    }
    switch -323 {
    case 685:
        v4 -= v4
    default:
        v3 = (v3 | (v3 + v3))
    }
    v16 := GenS1{F0: -426, F1: false, F2: 1}
    for fi17 := 0; fi17 < 3; fi17++ {
        v2 = (3.14159 + v2)
        var v18 bool = false
        v19 := []int{(fi17 / (((fi17 & (fi17 & (^(fi17))))) | 1)), (fi17 * (fi17 % ((fi17) | 1)))}
        _ = fi17
        _ = v18
        _ = v19
    }
    if true {
        v2 -= ((v2 - 1.5) + 0.0)
        v2 += 0.1
        switch 42 {
        case 4294967295:
            if (-480 <= -9223372036854775808) {
                var v20 int8 = int8(v3)
                fmt.Println(v20, v3)
                v2 = v2
                _ = v20
            }
        case 1:
            switch -602 {
            case -590:
                v21 := GenS3{F0: (v3 % ((229) | 1))}
                _ = v21
            case -100:
                switch -990 {
                case 455:
                    v22 := GenS1{F0: 9223372036854775807, F1: true, F2: 529}
                    _ = v22
                default:
                    v3 = (^(v3))
                }
            }
        }
    } else {
        var v23 uint32 = (v4 * v4)
        fmt.Println(v4, v4, v4)
        _ = v23
    }
    v4 = uint32(v3)
    var v24 float32 = ((3.14159 * 3.14159) * 1.5)
    var v25 float32 = -2.25
    var v26 uint64 = 18446744073709551615
    var v27 int8 = -1
    v24 = 1.5
    for fi28 := 0; fi28 < 19; fi28++ {
        delete(v1, "\n\t")
        _ = fi28
    }
    v29 := v1["a"]
    var v30 string = "\n\t"
    switch "foo bar" {
    case "\n\t":
        var v31 float32 = v25
        _ = v31
    case "":
        var v32 int16 = -32768
        _ = v32
    case "a":
        v33 := []int{}
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
