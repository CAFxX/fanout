package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int8 }
func genFn0(p0 uint, p1 int8, p2 bool) uint64 {
    fmt.Println(p0)
    p0 -= 0
    v5 := GenS1{F0: "foo bar", F1: false, F2: (p1 / ((100) | 1))}
    v6 := GenS1{F0: ("" + "foo bar"), F1: (-2.25 > 0.0), F2: (p1 & (p1 / (((-(p1))) | 1)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 559
}
func genFn1(p0 uint16, p1 uint32) int32 {
    p0 = 255
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{fi7, 2147483647, (fi7 + fi7), 4294967295}
        if (p1 != (p1 * (^(p1)))) {
            p0 -= 255
            var v9 uint16 = (p0 | (p0 + 147))
            _ = len(v8)
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v10 uint64 = 18446744073709551615
    switch ("a" + "hello") {
    case "hello":
        v10 -= (v10 % ((v10) | 1))
    }
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    var v1 int16 = -554
    var v2 int8 = -41
    var v3 float32 = 1e10
    v4 := []uint32{4294967295, 2147483648, 2147483647, 4294967295}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 = (v1 & v1)
    fmt.Println(v1, v1)
    var v11 uint16 = 0
    _ = len(v4)
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            if false {
                if len(v4) > 0 {
                    v4[0] = 1
                }
                v14 := GenS1{F0: "hello", F1: false, F2: -128}
                if len(v4) > 0 {
                    v4[0] = 65535
                }
                _ = v14
            } else {
                v15 := GenS1{F0: "\n\t", F1: false, F2: (v2 % ((100) | 1))}
                var v16 int = fi13
                _ = v15
                _ = v16
            }
            _ = fi13
        }
        _ = fi12
    }
    _ = len(v4)
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 float64 = 0.0
            _ = fi18
            _ = v19
        }
        v1 += v1
        v3 = v3
        _ = fi17
    }
    fmt.Println(v1, v3)
    for fi20 := 0; fi20 < 17; fi20++ {
        fmt.Println(fi20, v3, fi20)
        if (0 > 0) {
            _ = len(v4)
        }
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
