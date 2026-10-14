package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 19; fi4++ {
        if false {
            switch "" {
            case "\n\t":
                var v5 uint = 633
                _ = v5
            }
        } else {
            v6 := []int{fi4, fi4}
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 bool = (false && (true && (65535 == 508)))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                fmt.Println(fi4, v8, fi7)
                _ = fi7
                _ = v8
            }
            switch "a" {
            case "":
                var v9 int = fi4
                _ = v9
            default:
                var v10 float64 = ((-2.25 - 0.1) / 1.0)
                _ = v10
            }
            _ = v6
        }
        _ = fi4
    }
    if ((!(true && (2 > 0))) || (!(-9223372036854775808 == 42))) {
        var v11 int16 = -698
        switch "a" {
        case "\n\t":
            var v12 int8 = -1
            _ = v12
        case "":
            var v13 uint16 = 951
            _ = v13
        default:
            v11 -= (v11 % ((0) | 1))
        }
        _ = v11
    }
    var v14 uint32 = 2147483648
    _ = v14
    return 255
}
func genFn1(p0 int16, p1 string) int16 {
    var v15 int32 = 2147483647
    fmt.Println(p0, v15, p1)
    if true {
        fmt.Println(p0)
        for fi16 := 0; fi16 < 5; fi16++ {
            if (1 >= 255) {
                v15 = -2147483648
                p1 = (p1 + p1)
                var v17 bool = ("\n\t" == p1)
                _ = v17
            }
            var v18 bool = false
            var v19 float32 = (-2.25 + (1.5 / 0.1))
            _ = fi16
            _ = v18
            _ = v19
        }
    }
    var v20 int = 9223372036854775807
    var v21 int = (v20 >> ((42) & 63))
    _ = p0
    _ = p1
    _ = v15
    _ = v20
    _ = v21
    return 1
}

func main() {
    var v1 uint16 = 2
    v2 := []int{1}
    v3 := []int{1, 531}
    v1 = 769
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := GenS1{F0: uint64(v1), F1: "hello"}
        for fi24 := 0; fi24 < 4; fi24++ {
            _ = len(v3)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 -= genFn0()
            _ = fi24
        }
        v1 += genFn0()
        _ = fi22
        _ = v23
    }
    v3 = append(v3, 2147483647)
    var v25 int8 = 42
    v1 = (v1 << ((uint(v1)) & 63))
    var v26 float64 = float64(1e10)
    if (v26 < float64((0.1 + (0.0 - (0.1 * float32(1.5)))))) {
        if len(v2) > 0 {
            v2[0] = 32767
        }
    }
    v27 := GenS1{F0: 341, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
