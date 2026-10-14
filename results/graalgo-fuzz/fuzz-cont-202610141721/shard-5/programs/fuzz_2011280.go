package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint32; F1 string }
func genFn0(p0 int32, p1 string, p2 uint32) uint64 {
    switch p1 {
    case "foo bar":
        p1 = (p1 + "a")
    case "\n\t":
        switch (("hello" + ((p1 + ("foo bar" + p1)) + ((p1 + "hello") + p1))) + (((("\n\t" + p1) + p1) + ((p1 + "") + p1)) + (p1 + ((p1 + (("foo bar" + p1) + ("hello" + (p1 + (p1 + p1))))) + ("" + ""))))) {
        case "\n\t":
            if (!((p2 >> ((261) & 63)) > p2)) {
                switch 42 {
                case -487:
                    var v7 string = p1
                    _ = v7
                default:
                    p1 = (p1 + (p1 + p1))
                }
                var v8 int64 = 1
                _ = v8
            } else {
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := GenS1{F0: 0, F1: 255, F2: p0}
                    var v11 int16 = 1000
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                switch -1 {
                case -941:
                    fmt.Println(p0, p2)
                case -1:
                    v12 := []int{-1, 42, -2147483648, int(p2)}
                    _ = v12
                default:
                    p1 = "\n\t"
                }
            }
        case "":
            var v13 uint = 2
            _ = v13
        default:
            var v14 float64 = 0.0
            _ = v14
        }
    default:
        var v15 int16 = 32767
        _ = v15
    }
    fmt.Println(p1, p1)
    switch 974 {
    case 0:
        v16 := GenS2{F0: p2, F1: p1}
        _ = v16
    case -268:
        p2 = (p2 & (p2 | (^(p2))))
    default:
        v17 := GenS2{F0: 522, F1: "foo bar"}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}
func genFn1(p0 uint64, p1 string) int64 {
    v18 := GenS1{F0: -135, F1: 196, F2: -2147483648}
    var v19 uint16 = 234
    v20 := GenS1{F0: 1, F1: 127, F2: -343}
    switch 896 {
    case -9223372036854775808:
        var v21 uint16 = (v19 | 2)
        _ = v21
    case 42:
        var v22 uint8 = 67
        _ = v22
    default:
        v23 := []int{}
        _ = v23
    }
    v24 := GenS2{F0: 4294967295, F1: (p1 + "a")}
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v20
    _ = v24
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 226
    var v3 int16 = 42
    v4 := []int{127, -716, -1}
    v5 := []string{}
    var v6 float32 = (3.14159 + 1.5)
    fmt.Println(v2)
    v2 -= v2
    switch ("a" + "\n\t") {
    case "a":
        var v25 int32 = 0
        _ = v25
    case "\n\t":
        var v26 uint8 = 0
        _ = v26
    case "foo bar":
        v4 = append(v4, 626)
    default:
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
    }
    fmt.Println(v2)
    for fi27 := 0; fi27 < 7; fi27++ {
        v6 += v6
        v28 := GenS1{F0: -937, F1: 127, F2: 42}
        _ = fi27
        _ = v28
    }
    v1[860] = 4294967295
    for fi29 := 0; fi29 < 15; fi29++ {
        for fi30 := 0; fi30 < 14; fi30++ {
            switch ("foo bar" + ("hello" + "hello")) {
            case "hello":
                var v31 uint = 0
                _ = v31
            case "\n\t":
                v6 += v6
            }
            _ = fi30
        }
        v32 := GenS2{F0: 2147483647, F1: ("hello" + "hello")}
        _ = fi29
        _ = v32
    }
    var v33 float64 = 1.5
    _ = len(v4)
    v33 += v33
    var v34 uint = 42
    var v35 int16 = v3
    v6 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
