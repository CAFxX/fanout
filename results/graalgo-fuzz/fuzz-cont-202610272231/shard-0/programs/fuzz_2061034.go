package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
func genFn0(p0 uint16, p1 int8, p2 int8) int8 {
    fmt.Println(p0, p1)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1() int32 {
    var v4 int8 = -1
    v4 = 0
    var v5 float32 = (1e10 * 1.5)
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    var v3 int8 = -1
    var v6 float32 = 0.1
    var v7 int = 127
    v3 += v3
    if (0 >= 2) {
        switch "hello" {
        case "foo bar":
            v1[0] = ""
        case "\n\t":
            _ = len(v1)
        case "":
            v8 := []int{}
            _ = v8
        }
        var v9 int64 = 478
        for fi10 := 0; fi10 < 6; fi10++ {
            v6 += ((((v6 + 3.14159) / 1.5) + v6) * (v6 * (0.0 + 3.14159)))
            fmt.Println(fi10, fi10)
            var v11 int32 = 0
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        var v12 uint8 = 0
        _ = v12
    }
    var v13 float32 = ((v6 - float32(-2.25)) - (((v6 * ((0.1 - 3.14159) + v6)) * v6) + v6))
    var v14 float64 = 3.14159
    var v15 uint16 = 403
    var v16 uint8 = 0
    switch "foo bar" {
    case "":
        v15 += 255
    }
    _ = len(v1)
    v17 := GenS1{F0: 1000, F1: 1}
    if (!false) {
        v18 := v2[827]
        if false {
            v19 := GenS1{F0: 0, F1: -459}
            v6 = (1.5 + v6)
            if false {
                var v20 uint16 = v15
                switch 4294967295 {
                case 32767:
                    v14 = 1.5
                default:
                    v21 := v2[(v18 * (v18 * (^(v18))))]
                    _ = v21
                }
                _ = v20
            }
            _ = v19
        }
        _ = len(v1)
        _ = v18
    }
    _ = len(v1)
    switch v7 {
    case 4294967295:
        v22 := []int{4294967295, v7, (v7 >> ((833) & 63)), (v7 & 0)}
        _ = v22
    case 276:
        switch ("\n\t" + "\n\t") {
        case "hello":
            var v23 string = ("\n\t" + ("foo bar" + "a"))
            _ = v23
        case "\n\t":
            var v24 bool = (false || (false || false))
            _ = v24
        case "":
            switch "hello" {
            case "\n\t":
                v25 := GenS1{F0: 0, F1: 42}
                _ = v25
            case "foo bar":
                var v26 uint64 = 18446744073709551615
                _ = v26
            }
        default:
            var v27 uint64 = 2
            _ = v27
        }
    case 127:
        v28 := GenS1{F0: 0, F1: -9223372036854775808}
        _ = v28
    default:
        v29 := GenS1{F0: -345, F1: 42}
        _ = v29
    }
    var v30 bool = ((!true) || (!true))
    for fi31 := 0; fi31 < 8; fi31++ {
        var v32 uint32 = 51
        var v33 int64 = 223
        _ = fi31
        _ = v32
        _ = v33
    }
    v34 := []int{(v7 & 411), (^(v7))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v6
    _ = v7
}
