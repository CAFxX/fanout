package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 bool }
type GenS4 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint8, p1 uint) uint8 {
    switch "" {
    case "hello":
        fmt.Println(p1, p1)
    case "a":
        if true {
            switch "a" {
            case "foo bar":
                v4 := []int{-1, 128, -1}
                _ = v4
            case "hello":
                p1 += (p1 + 2)
            }
            var v5 int8 = 0
            v6 := []int{}
            _ = v5
            _ = v6
        }
    case "foo bar":
        p0 += (p0 >> (((^(p1))) & 63))
    default:
        v7 := GenS1{F0: -128, F1: true}
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(fi8, p1)
        if true {
            if (0 <= 4294967295) {
                var v9 float64 = (0.0 * float64(0.0))
                fmt.Println(fi8, p0)
                fmt.Println(fi8, fi8, p0)
                _ = v9
            } else {
                var v10 uint64 = 191
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            p0 = (p0 & 75)
            p0 = 1
            v12 := []int{(fi11 ^ 938), (fi8 << (((p1 % ((p1) | 1))) & 63)), (fi8 + (fi11 + 443))}
            _ = fi11
            _ = v12
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    return 128
}

func main() {
    v1 := []uint32{2, 717, 2147483647}
    var v2 int8 = -1
    var v3 uint = 18446744073709551615
    var v13 uint64 = 42
    v1 = append(v1, 331)
    if (2 > 75) {
        if (!(v2 > (v2 | -1))) {
            var v14 string = "a"
            v1 = append(v1, 2147483647)
            _ = v14
        } else {
            var v15 uint8 = 255
            v16 := GenS1{F0: (v2 + (^(v2))), F1: (788 > 1)}
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := GenS2{F0: 0, F1: 2}
                _ = fi17
                _ = v18
            }
            _ = v15
            _ = v16
        }
    }
    v19 := GenS1{F0: v2, F1: false}
    var v20 int = 597
    for fi21 := 0; fi21 < 15; fi21++ {
        v20 = fi21
        fmt.Println(v20, v13, v2)
        for fi22 := 0; fi22 < 12; fi22++ {
            fmt.Println(v20)
            fmt.Println(fi21, v3, v2)
            var v23 uint = v3
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    if false {
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 float64 = 1.5
            var v26 string = ""
            _ = fi24
            _ = v25
            _ = v26
        }
        switch ("\n\t" + (("hello" + "foo bar") + "hello")) {
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "hello":
            fmt.Println(v3, v3)
        default:
            if (!(true && true)) {
                if true {
                    v13 = (v13 % ((v13) | 1))
                    fmt.Println(v3)
                    v27 := GenS3{F0: (v3 >> (((v3 / ((v3) | 1))) & 63)), F1: v13, F2: true}
                    _ = v27
                } else {
                    v28 := GenS4{F0: 255, F1: -594}
                    _ = v28
                }
            }
        }
        _ = len(v1)
    }
    v29 := GenS2{F0: 0, F1: v3}
    var v30 uint32 = 0
    fmt.Println(v13, v30, v2)
    switch (-(v20)) {
    case 42:
        var v31 int = (v20 % ((415) | 1))
        _ = v31
    }
    v30 = (v30 ^ v30)
    var v32 int32 = -701
    if true {
        var v33 int64 = 1
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v32
}
