package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int64 }
type GenS2 struct { F0 int8; F1 int16; F2 int32 }
type GenS3 struct { F0 int8 }
func genFn0(p0 string, p1 uint, p2 int16) uint8 {
    var v6 uint8 = 2
    for fi7 := 0; fi7 < 13; fi7++ {
        if (1 != -128) {
            var v8 float32 = 0.0
            _ = v8
        }
        p2 += (p2 & 32767)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 int) uint16 {
    var v9 uint32 = 2147483647
    var v10 uint32 = v9
    v9 = v10
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int64 = 531
    v3 := []string{(("hello" + (("" + "\n\t") + "foo bar")) + "a"), ("foo bar" + "\n\t")}
    v4 := []uint32{}
    v5 := make(map[string]bool)
    var v11 uint8 = 127
    v3 = append(v3, "")
    v4 = append(v4, (^(v1)))
    var v12 string = "hello"
    v4 = append(v4, v1)
    v13 := GenS1{F0: (v1 <= v1), F1: (v12 + v12), F2: v2}
    if (-32768 < 42) {
        v14 := GenS2{F0: -6, F1: 1, F2: 2147483647}
        _ = v14
    }
    v15 := GenS2{F0: 100, F1: 1000, F2: 0}
    if (-32768 >= 123) {
        var v16 bool = true
        _ = v16
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1, v1)
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 int8 = 0
        if (149 <= 255) {
            v19 := GenS3{F0: 127}
            _ = v19
        } else {
            var v20 float32 = float32((1e10 - -2.25))
            _ = v20
        }
        for fi21 := 0; fi21 < 6; fi21++ {
            if true {
                var v22 uint = 1
                v18 = v18
                v11 = v11
                _ = v22
            } else {
                var v23 uint32 = 4294967295
                _ = v23
            }
            v1 = (v1 - (v1 + v1))
            v11 += v11
            _ = fi21
        }
        _ = fi17
        _ = v18
    }
    var v24 uint32 = 573
    v25 := GenS1{F0: ((^(v2)) < 9223372036854775807), F1: v12, F2: (v2 % ((42) | 1))}
    var v26 string = ("a" + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
