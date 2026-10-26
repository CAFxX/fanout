package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0(p0 string, p1 string, p2 int8) uint16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: fi7, F1: 4294967295}
        p2 += p2
        _ = fi7
        _ = v8
    }
    var v9 int8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{4294967295, 1}
    v2 := []string{"hello", ("\n\t" + ("hello" + "")), "a"}
    v3 := make(map[string]string)
    var v4 uint8 = 127
    v5 := make(map[int]string)
    var v6 int8 = 100
    if (-9223372036854775808 != 42) {
        v10 := GenS2{F0: (3.14159 != float32(0.1)), F1: 110}
        var v11 int = 2147483647
        v4 -= v4
        _ = v10
        _ = v11
    } else {
        var v12 uint = 0
        _ = v12
    }
    if true {
        v6 = 42
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS1{F0: (fi13 ^ (fi13 + fi13)), F1: 1}
            v15 := GenS1{F0: fi13, F1: 671}
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        fmt.Println(v4, v4)
    }
    v16 := v3[("a" + "foo bar")]
    for fi17 := 0; fi17 < 3; fi17++ {
        switch fi17 {
        case 2147483647:
            v18 := GenS2{F0: true, F1: 0}
            _ = v18
        default:
            var v19 int = (fi17 * fi17)
            _ = v19
        }
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 uint64 = 1
            _ = fi20
            _ = v21
        }
        _ = fi17
    }
    var v22 uint = 2
    v23 := GenS2{F0: false, F1: genFn0((v16 + v16), v16, (^(v6)))}
    switch (("" + ("" + "hello")) + (v16 + (v16 + v16))) {
    case "a":
        v4 -= (v4 ^ v4)
    case "\n\t":
        var v24 float64 = -2.25
        _ = v24
    default:
        var v25 int32 = -2147483648
        _ = v25
    }
    v6 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
