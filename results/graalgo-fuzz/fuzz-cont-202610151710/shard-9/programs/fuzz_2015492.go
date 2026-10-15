package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int8, p1 uint, p2 uint) int64 {
    p0 = p0
    var v6 uint = p2
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}

func main() {
    v1 := []string{"\n\t", "a", "a", "a"}
    var v2 int8 = -114
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    v5 := make(map[string]bool)
    v4[2147483647] = 1
    switch -1 {
    case 32768:
        if ((42 <= 0) || true) {
            var v7 int = -2147483648
            if false {
                v7 = 909
            } else {
                v7 -= v7
            }
            v7 -= 1
            _ = v7
        }
    case 1:
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: 42, F1: 1}
            var v10 int16 = 1000
            for fi11 := 0; fi11 < 20; fi11++ {
                switch "a" {
                case "hello":
                    v12 := GenS1{F0: 2, F1: v10}
                    _ = v12
                case "\n\t":
                    fmt.Println(fi8, fi11, fi8)
                    v2 = 100
                    v10 = 86
                case "foo bar":
                    v13 := GenS2{F0: true}
                    _ = v13
                }
                _ = fi11
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v14 := v4[128]
    v2 += (v2 * (v2 + (v2 % ((v2) | 1))))
    var v15 int8 = v2
    var v16 uint32 = 846
    switch "foo bar" {
    case "":
        if (true && ((1e10 + (float32(1.5) - 0.1)) > float32(float64((1e10 / (0.1 * (0.1 + 0.1))))))) {
            var v17 uint8 = uint8(v16)
            _ = v17
        } else {
            var v18 int8 = v2
            _ = v18
        }
    default:
        v15 = (-(v2))
    }
    var v19 float32 = 1.5
    var v20 uint16 = 65535
    v21 := GenS1{F0: 331, F1: 42}
    v16 = v16
    for fi22 := 0; fi22 < 11; fi22++ {
        switch (("foo bar" + "") + (("hello" + "a") + "foo bar")) {
        case "\n\t":
            var v23 int32 = -1
            _ = v23
        case "a":
            v2 -= v15
        default:
            if (!(-1 > (v14 >> ((189) & 63)))) {
                fmt.Println(v20)
                for fi24 := 0; fi24 < 5; fi24++ {
                    switch fi24 {
                    case 0:
                        var v25 string = ("" + "hello")
                        var v26 uint = 18446744073709551615
                        var v27 int64 = genFn0((v15 >> (((v26 % (((v26 ^ (v26 << ((v26) & 63)))) | 1))) & 63)), v26, (^(v26)))
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    _ = len(v1)
                    v14 = (fi22 % ((fi22) | 1))
                    _ = fi24
                }
            }
        }
        _ = fi22
    }
    v2 -= (v15 % ((127) | 1))
    var v28 uint = 18446744073709551615
    var v29 int64 = -9223372036854775808
    var v30 uint64 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
