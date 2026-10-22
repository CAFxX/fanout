package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int64 }
func genFn0(p0 uint64, p1 uint) uint8 {
    var v6 string = "foo bar"
    p0 += (p0 & 4294967295)
    var v7 string = ((v6 + v6) + v6)
    var v8 uint8 = 0
    fmt.Println(v8, p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return v8
}
func genFn1(p0 uint) int64 {
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 42
}

func main() {
    v1 := []uint32{42, 2147483648, 826, 52}
    var v2 uint = 42
    v3 := make(map[string]string)
    var v4 uint8 = 137
    var v5 string = "\n\t"
    fmt.Println(v4)
    v3[(v5 + v5)] = "a"
    delete(v3, "foo bar")
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 float64 = -2.25
        if (2147483647 == 4294967295) {
            v11 := GenS1{F0: 0, F1: 0, F2: v2}
            v2 -= (^(v2))
            _ = v11
        } else {
            for fi12 := 0; fi12 < 2; fi12++ {
                switch v5 {
                case "":
                    var v13 float64 = v10
                    _ = v13
                case "foo bar":
                    var v14 uint = (^(v2))
                    _ = v14
                default:
                    v4 += genFn0(805, v2)
                }
                fmt.Println(v10, v4)
                v5 = v5
                v10 += v10
                _ = fi12
            }
        }
        if true {
            _ = len(v1)
            if (!(0 < (fi9 << ((702) & 63)))) {
                v15 := GenS1{F0: 0, F1: 512, F2: (v2 + v2)}
                v2 += v2
                _ = v15
            }
        }
        _ = fi9
        _ = v10
    }
    v2 -= 65535
    v4 -= genFn0(0, (v2 & 18446744073709551615))
    var v16 uint64 = 42
    var v17 int16 = -32768
    if (!(0 < 32768)) {
        v17 = -32768
    }
    v18 := v3[v5]
    switch "hello" {
    case "a":
        v4 = 228
    case "":
        v19 := GenS1{F0: 42, F1: 27, F2: v2}
        _ = v19
    case "\n\t":
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 float64 = 3.14159
            _ = fi20
            _ = v21
        }
    }
    if false {
        v16 = v16
        for fi22 := 0; fi22 < 1; fi22++ {
            switch v5 {
            case "":
                v23 := GenS2{F0: (v17 >> ((v2) & 63))}
                _ = v23
            case "hello":
                delete(v3, v18)
            default:
                v24 := GenS3{F0: (v2 * v2), F1: genFn1(965)}
                _ = v24
            }
            v16 -= v16
            _ = fi22
        }
    }
    var v25 string = "hello"
    v17 -= -573
    fmt.Println(v4)
    var v26 string = v5
    var v27 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
