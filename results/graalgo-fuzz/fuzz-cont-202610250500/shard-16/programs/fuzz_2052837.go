package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint64, p1 int8) int32 {
    var v7 int = -1
    switch "foo bar" {
    case "hello":
        v7 += v7
    case "":
        v8 := GenS1{F0: "", F1: 421, F2: -32768}
        _ = v8
    case "a":
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int16 = 1000
            var v11 uint = 574
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    p0 = (p0 % (((p0 * 2)) | 1))
    _ = p0
    _ = p1
    _ = v7
    return 2147483647
}
func genFn1(p0 uint32, p1 uint32) uint8 {
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(p0)
        v13 := GenS2{F0: "a"}
        _ = fi12
        _ = v13
    }
    v14 := []int{-149, -9223372036854775808, 127, 429}
    _ = p0
    _ = p1
    _ = v14
    return 42
}

func main() {
    var v1 string = ("a" + ((("foo bar" + (("\n\t" + ("\n\t" + "foo bar")) + (("\n\t" + "") + "hello"))) + "\n\t") + "\n\t"))
    var v2 int8 = 75
    var v3 int = 127
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]bool)
    var v6 int16 = -914
    v15 := GenS2{F0: ""}
    for fi16 := 0; fi16 < 12; fi16++ {
        _ = len(v5)
        v17 := GenS2{F0: v1}
        var v18 uint16 = 255
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := []int{}
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 uint16 = 255
        if true {
            var v22 uint64 = 423
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 uint16 = v21
                _ = len(v5)
                v25 := v4[793]
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = v22
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            var v27 int8 = v2
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            switch fi20 {
            case 465:
                if false {
                    if (255 > 401) {
                        switch fi26 {
                        case 127:
                            v21 += v21
                        default:
                            delete(v5, 505)
                        }
                        v28 := v5[1]
                        v19 = append(v19, v3)
                        _ = v28
                    } else {
                        if true {
                            var v29 float64 = 0.0
                            _ = v29
                        }
                        fmt.Println(fi20, fi26)
                    }
                }
            case -758:
                if (!false) {
                    v30 := GenS3{F0: v2}
                    v27 = -128
                    switch ((v1 + v1) + v1) {
                    case "hello":
                        v5[550] = true
                    }
                    _ = v30
                } else {
                    v5[2147483648] = (0 <= 932)
                    v31 := GenS2{F0: ("hello" + ("a" + ("\n\t" + (v1 + v1))))}
                    _ = v31
                }
            }
            _ = fi26
            _ = v27
        }
        _ = fi20
        _ = v21
    }
    var v32 int64 = 0
    v32 = v32
    v33 := GenS4{F0: (v6 - -399)}
    _ = len(v5)
    v6 -= v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
