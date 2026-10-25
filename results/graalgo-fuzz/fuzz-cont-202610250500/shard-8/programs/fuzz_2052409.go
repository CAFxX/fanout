package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 bool) uint8 {
    p0 = (((("a" + "\n\t") + ("\n\t" + "")) + ("\n\t" + "")) < "\n\t")
    p0 = p0
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float64 = 1e10
        if (457 <= 32767) {
            switch "\n\t" {
            case "":
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 bool = (!false)
                    fmt.Println(fi6, v7)
                    _ = fi8
                    _ = v9
                }
            case "foo bar":
                v10 := []int{fi6, fi6}
                _ = v10
            case "a":
                v11 := []int{}
                _ = v11
            }
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 int = (fi12 ^ 843)
                fmt.Println(fi12, p0, fi12)
                _ = fi12
                _ = v13
            }
            p0 = p0
        }
        _ = fi6
        _ = v7
    }
    v14 := []int{}
    p0 = (42 < -35)
    _ = p0
    _ = v14
    return 128
}
func genFn1() uint64 {
    v15 := GenS1{F0: 48}
    v16 := GenS2{F0: -213, F1: 1000}
    var v17 int = 128
    v17 += v17
    _ = v15
    _ = v16
    _ = v17
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 float64 = 3.14159
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    var v5 int = 4294967295
    if false {
        if (((float32(v2) / -2.25) == 1e10) || true) {
            delete(v3, 128)
        } else {
            if (234 >= 150) {
                v2 -= float64(0.0)
                fmt.Println(v5, v5)
            } else {
                v2 = 1e10
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                fmt.Println(v5, v2)
            }
        }
    }
    v5 = v5
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 uint8 = 42
        switch (v5 >> ((4294967295) & 63)) {
        case -683:
            fmt.Println(v2, fi18, v19)
        case 477:
            v3[fi18] = true
        case 4294967295:
            v20 := v3[2147483647]
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    v21 := GenS3{F0: v5, F1: 0}
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 uint64 = 1
        fmt.Println(v23)
        switch "a" {
        case "\n\t":
            v24 := GenS2{F0: 1, F1: 1}
            _ = v24
        case "a":
            if len(v1) > 0 {
                v1[0] = 247
            }
        default:
            switch "hello" {
            case "foo bar":
                if true {
                    v25 := GenS2{F0: 2147483647, F1: -233}
                    v4[(fi22 % (((v5 % (((fi22 - -9223372036854775808)) | 1))) | 1))] = 357
                    for fi26 := 0; fi26 < 13; fi26++ {
                        var v27 int64 = 42
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                } else {
                    var v28 uint16 = 65535
                    var v29 int8 = 0
                    v2 -= v2
                    _ = v28
                    _ = v29
                }
            case "a":
                if false {
                    v23 -= 1
                }
            case "":
                var v30 uint = 731
                _ = v30
            default:
                v23 = (v23 % ((v23) | 1))
            }
        }
        _ = fi22
        _ = v23
    }
    var v31 uint16 = 860
    fmt.Println(v5, v5)
    for fi32 := 0; fi32 < 10; fi32++ {
        v33 := v4[127]
        for fi34 := 0; fi34 < 7; fi34++ {
            v35 := v3[-2147483648]
            var v36 float32 = float32(v2)
            var v37 uint = 0
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        v5 = (v5 ^ fi32)
        _ = fi32
        _ = v33
    }
    if false {
        if true {
            for fi38 := 0; fi38 < 15; fi38++ {
                var v39 int32 = 42
                _ = fi38
                _ = v39
            }
            var v40 string = "hello"
            _ = v40
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v31 = v31
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
