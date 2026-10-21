package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 int8; F1 string; F2 uint16 }
func genFn0(p0 int16, p1 uint, p2 bool) int {
    if false {
        for fi4 := 0; fi4 < 6; fi4++ {
            v5 := []int{fi4, (fi4 >> (((p1 % ((p1) | 1))) & 63))}
            _ = fi4
            _ = v5
        }
        var v6 uint = (p1 % (((^(p1))) | 1))
        _ = v6
    } else {
        fmt.Println(p1)
    }
    var v7 int8 = 32
    p2 = ((29 <= v7) || (((((("" + "\n\t") + (("" + "") + "hello")) + "\n\t") + "hello") + "\n\t") == ""))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32768
}

func main() {
    v1 := []uint32{2147483647, 1}
    var v2 uint16 = 533
    v3 := []uint32{478, 42, 1}
    var v8 uint8 = 128
    if (!true) {
        v9 := GenS1{F0: 4294967295, F1: 0, F2: 600}
        switch "" {
        case "\n\t":
            var v10 uint16 = (v2 | v2)
            _ = v10
        default:
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 uint = 42
                _ = fi11
                _ = v12
            }
        }
        _ = v9
    }
    v13 := GenS2{F0: -111, F1: "foo bar", F2: v2}
    if true {
        v8 = v8
        v8 += v8
        v8 += v8
    }
    var v14 bool = true
    switch "foo bar" {
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "foo bar":
        v2 -= v2
    case "":
        var v15 uint8 = (^(v8))
        _ = v15
    default:
        var v16 float64 = 3.14159
        _ = v16
    }
    fmt.Println(v2)
    fmt.Println(v14)
    v8 = (v8 | (^(v8)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
}
