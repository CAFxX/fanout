package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
func genFn0() uint32 {
    v7 := []int{9223372036854775807, 127, 42, -2147483648}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v8 := GenS1{F0: 128, F1: 32767}
    if (-315 < -992) {
        var v9 uint8 = 2
        v10 := GenS1{F0: 85, F1: -789}
        v11 := GenS1{F0: v9, F1: 0}
        _ = v10
        _ = v11
        _ = v9
    }
    _ = v7
    _ = v8
    return 171
}
func genFn1(p0 string) uint {
    p0 = (p0 + (p0 + ((p0 + (p0 + p0)) + p0)))
    p0 = "hello"
    for fi12 := 0; fi12 < 11; fi12++ {
        switch (^(fi12)) {
        case 127:
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 uint32 = 914
                _ = fi13
                _ = v14
            }
        case -2147483648:
            fmt.Println(fi12, p0)
        case -9223372036854775808:
            var v15 float32 = 0.0
            _ = v15
        }
        _ = fi12
    }
    _ = p0
    return 255
}

func main() {
    var v1 int8 = 32
    var v2 string = "foo bar"
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    var v5 uint = 18446744073709551615
    v6 := []uint32{65535}
    if true {
        var v16 uint32 = 42
        for fi17 := 0; fi17 < 3; fi17++ {
            for fi18 := 0; fi18 < 5; fi18++ {
                if (int64(v5) == -779) {
                    v19 := GenS1{F0: 255, F1: (-(fi18))}
                    _ = len(v3)
                    var v20 float64 = float64(-2.25)
                    _ = v19
                    _ = v20
                }
                var v21 uint16 = 39
                var v22 uint = (v5 - (v5 / ((255) | 1)))
                v23 := GenS1{F0: 128, F1: (fi17 & (fi17 % (((fi17 + -869)) | 1)))}
                v16 -= (^(v16))
                v24 := v4[""]
                _ = fi18
                _ = v21
                _ = v22
                _ = v23
                _ = v24
            }
            if (true || (v1 < -128)) {
                v16 = 1
                for fi25 := 0; fi25 < 3; fi25++ {
                    v2 = "a"
                    switch (("foo bar" + "") + (v2 + v2)) {
                    case "a":
                        var v26 int16 = 1000
                        _ = v26
                    case "foo bar":
                        switch -861 {
                        case -111:
                            var v27 bool = false
                            _ = v27
                        case -9223372036854775808:
                            v28 := v4[v2]
                            _ = v28
                        default:
                            var v29 int16 = 624
                            _ = v29
                        }
                    default:
                        var v30 int8 = v1
                        _ = v30
                    }
                    var v31 bool = false
                    v31 = v31
                    _ = fi25
                    _ = v31
                }
            }
            _ = fi17
        }
        if false {
            for fi32 := 0; fi32 < 12; fi32++ {
                v2 = (v2 + v2)
                _ = fi32
            }
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v2, v5, v5)
            case "a":
                var v33 int8 = v1
                _ = v33
            case "\n\t":
                for fi34 := 0; fi34 < 20; fi34++ {
                    var v35 float64 = 1.5
                    var v36 uint64 = 42
                    fmt.Println(v36, v1)
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            v2 = v2
        }
        _ = v16
    } else {
        var v37 uint8 = 188
        var v38 int = 32767
        var v39 uint64 = 42
        _ = v37
        _ = v38
        _ = v39
    }
    if false {
        v40 := GenS2{F0: -680, F1: 0}
        _ = v40
    }
    fmt.Println(v1, v1)
    v41 := v3[760]
    var v42 uint8 = 10
    for fi43 := 0; fi43 < 5; fi43++ {
        fmt.Println(v1, v42, v41)
        var v44 int8 = (v1 | v1)
        v45 := GenS1{F0: 1, F1: (fi43 / ((v41) | 1))}
        _ = fi43
        _ = v44
        _ = v45
    }
    v46 := GenS3{F0: 4294967295, F1: 451, F2: -247}
    var v47 float32 = (float32(1e10) / (3.14159 - 0.1))
    _ = len(v6)
    var v48 int32 = 399
    v5 += genFn1("a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
