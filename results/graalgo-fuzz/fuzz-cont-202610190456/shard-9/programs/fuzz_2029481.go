package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 bool }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64, p1 uint8) bool {
    v7 := []int{-2147483648, 32768}
    switch (("a" + "") + ("" + "hello")) {
    case "":
        if ((-2.25 == (((3.14159 - 0.1) * 0.0) * 0.1)) && false) {
            var v8 float32 = 0.1
            var v9 bool = false
            var v10 int64 = int64(p0)
            _ = v10
            _ = v8
            _ = v9
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                switch fi11 {
                case 1:
                    var v12 float64 = 1.5
                    _ = v12
                case 128:
                    var v13 uint64 = p0
                    _ = v13
                case 277:
                    v7 = append(v7, fi11)
                }
                p1 = (p1 + (p1 * (^(p1))))
                p0 = 4294967295
                _ = fi11
            }
        }
    case "\n\t":
        var v14 bool = (-32768 > 42)
        _ = v14
    case "a":
        v15 := GenS1{F0: 32767, F1: -1, F2: (255 <= 0)}
        _ = v15
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v7
    return false
}

func main() {
    v1 := []uint32{2}
    v2 := make(map[string]string)
    v3 := []int{9223372036854775807, -2147483648}
    var v4 string = "hello"
    v5 := make(map[int]string)
    var v6 uint64 = 4294967295
    v3 = append(v3, 32768)
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            fmt.Println(fi17, v6, fi16)
            v6 = v6
            _ = fi17
        }
        switch 4294967295 {
        case 32768:
            v18 := v5[fi16]
            _ = v18
        }
        _ = fi16
    }
    v4 = v4
    v19 := GenS1{F0: 804, F1: 2147483647, F2: false}
    for fi20 := 0; fi20 < 15; fi20++ {
        v6 -= 397
        _ = fi20
    }
    var v21 uint64 = 42
    if true {
        v6 = (v21 & (^(v6)))
        var v22 float64 = -2.25
        var v23 int32 = 0
        _ = v22
        _ = v23
    }
    var v24 int8 = 1
    fmt.Println(v24, v21, v24)
    fmt.Println(v4, v21, v6)
    if (-682 > 1) {
        v4 = "foo bar"
        var v25 uint64 = v6
        for fi26 := 0; fi26 < 13; fi26++ {
            if (-1 > 92) {
                v27 := GenS2{F0: 2147483647}
                var v28 float32 = 0.0
                _ = v27
                _ = v28
            }
            _ = fi26
        }
        _ = v25
    }
    var v29 float64 = (0.1 / 1.5)
    var v30 bool = false
    var v31 int16 = 1
    v31 -= (v31 * v31)
    v30 = v30
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
