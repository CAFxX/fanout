package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    v6 := GenS1{F0: -98, F1: (4294967295 >= (p0 % ((p0) | 1)))}
    v7 := GenS1{F0: 0, F1: false}
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint8 = 253
        if false {
            v10 := GenS1{F0: 0, F1: true}
            switch "a" {
            case "\n\t":
                switch "a" {
                case "\n\t":
                    var v11 int16 = 533
                    _ = v11
                }
            case "foo bar":
                var v12 bool = true
                _ = v12
            case "a":
                fmt.Println(v9, fi8, v9)
            }
            if false {
                if (true || (!true)) {
                    fmt.Println(fi8)
                    switch 1 {
                    case -730:
                        if (!true) {
                            var v13 string = "\n\t"
                            _ = v13
                        }
                    }
                }
                v14 := GenS1{F0: -1, F1: true}
                v15 := GenS2{F0: 23, F1: ("\n\t" + "")}
                _ = v14
                _ = v15
            }
            _ = v10
        } else {
            v16 := GenS2{F0: 897, F1: ("hello" + (("\n\t" + ("\n\t" + "a")) + ""))}
            _ = v16
        }
        v9 += 128
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 int64) uint64 {
    var v17 int16 = 42
    fmt.Println(v17)
    _ = p0
    _ = v17
    return 56
}

func main() {
    v1 := []int{}
    v2 := make(map[int]bool)
    var v3 uint = 2
    v4 := []uint32{544}
    var v5 int64 = 1
    v5 -= v5
    v3 = v3
    var v18 int = 948
    v19 := GenS3{F0: 2}
    v3 = 4294967295
    if (!true) {
        fmt.Println(v3, v18, v5)
        switch "hello" {
        case "":
            for fi20 := 0; fi20 < 6; fi20++ {
                fmt.Println(v18)
                var v21 uint32 = 1
                _ = fi20
                _ = v21
            }
        }
        if (83 < v5) {
            for fi22 := 0; fi22 < 16; fi22++ {
                _ = len(v1)
                _ = fi22
            }
            var v23 bool = true
            fmt.Println(v3, v23, v3)
            _ = v23
        }
    } else {
        if true {
            fmt.Println(v3, v18, v5)
        }
        var v24 uint16 = 251
        var v25 int = 1
        _ = v24
        _ = v25
    }
    var v26 bool = true
    for fi27 := 0; fi27 < 7; fi27++ {
        v28 := v2[(fi27 * (v18 - fi27))]
        v29 := GenS1{F0: 42, F1: ((3.14159 / float32(float64(((1e10 / 1.5) - 1.5)))) == (float32((1.5 + -2.25)) * 0.0))}
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        v26 = (3.14159 != (float64(-2.25) / 0.1))
        _ = fi30
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := v2[(v18 >> (((v3 << (((v3 - v3)) & 63))) & 63))]
        v33 := GenS2{F0: 0, F1: "hello"}
        v34 := GenS1{F0: -8, F1: (0 >= 1)}
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    v3 += (v3 / ((v3) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
