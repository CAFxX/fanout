package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int16; F1 int; F2 int16 }
type GenS4 struct { F0 bool; F1 uint8; F2 string }
func genFn0() uint8 {
    var v4 float32 = float32((1.5 - float64(0.0)))
    v4 += 0.1
    _ = v4
    return 1
}
func genFn1(p0 int8, p1 int32) uint8 {
    var v5 int16 = 42
    if (!false) {
        if true {
            var v6 string = "hello"
            _ = v6
        }
        p1 = 0
        fmt.Println(p0)
    } else {
        v7 := GenS1{F0: 0, F1: -176}
        var v8 uint64 = 18446744073709551615
        switch "" {
        case "foo bar":
            v9 := GenS1{F0: 0, F1: -420}
            _ = v9
        case "\n\t":
            switch 127 {
            case 1:
                var v10 uint64 = v8
                _ = v10
            case -9223372036854775808:
                switch 32767 {
                case 32768:
                    v11 := []int{127}
                    _ = v11
                case -650:
                    for fi12 := 0; fi12 < 15; fi12++ {
                        var v13 int16 = v5
                        _ = fi12
                        _ = v13
                    }
                default:
                    v14 := GenS1{F0: 542, F1: 2147483647}
                    _ = v14
                }
            }
        case "":
            if false {
                v15 := GenS2{F0: (p1 | -835), F1: 0}
                if false {
                    switch "hello" {
                    case "hello":
                        v16 := []int{1, 128}
                        _ = v16
                    }
                } else {
                    var v17 uint64 = (v8 << ((7) & 63))
                    _ = v17
                }
                v18 := []int{}
                _ = v15
                _ = v18
            } else {
                v5 += (v5 | (v5 | -1))
                v19 := []int{42, 32768, 0, 32768}
                _ = v19
            }
        default:
            switch -254 {
            case -336:
                for fi20 := 0; fi20 < 5; fi20++ {
                    v8 = (v8 * (v8 + (v8 & v8)))
                    var v21 int8 = (^(p0))
                    _ = fi20
                    _ = v21
                }
            case -1:
                switch "\n\t" {
                case "a":
                    fmt.Println(p0, p0)
                case "":
                    for fi22 := 0; fi22 < 3; fi22++ {
                        v23 := []int{32768}
                        if false {
                            fmt.Println(v8, v8)
                            v24 := GenS2{F0: (p1 ^ (p1 + 2147483647)), F1: 42}
                            _ = v24
                        } else {
                            v5 = v5
                        }
                        _ = fi22
                        _ = v23
                    }
                case "\n\t":
                    v25 := GenS2{F0: p1, F1: 1}
                    _ = v25
                }
            }
        }
        _ = v7
        _ = v8
    }
    v26 := []int{270, 661, 32768}
    var v27 uint64 = 2
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    _ = v5
    return 0
}

func main() {
    var v1 float64 = 1e10
    v2 := []int{}
    var v3 string = ("a" + ("a" + "\n\t"))
    v3 = (v3 + (v3 + v3))
    v1 = (((0.0 / (v1 + float64(1.5))) - ((v1 * v1) * v1)) + v1)
    v28 := GenS1{F0: 1, F1: 920}
    v1 = float64((float32((v1 - (v1 * v1))) + 3.14159))
    if ((!(1 >= 118)) || false) {
        for fi29 := 0; fi29 < 16; fi29++ {
            fmt.Println(v1)
            v30 := GenS3{F0: 1000, F1: (-(fi29)), F2: -568}
            v1 -= float64(((float32(v1) * (3.14159 * (-2.25 * 3.14159))) / 1.0))
            _ = fi29
            _ = v30
        }
    } else {
        var v31 uint16 = 2
        _ = v31
    }
    var v32 uint32 = 2
    for fi33 := 0; fi33 < 5; fi33++ {
        v34 := GenS4{F0: false, F1: 1, F2: v3}
        v32 = v32
        _ = fi33
        _ = v34
    }
    var v35 int64 = -1
    if true {
        switch int(v32) {
        case 9223372036854775807:
            v3 = v3
        case 4294967295:
            v2 = append(v2, 4294967295)
        case 42:
            var v36 int16 = 1
            _ = v36
        }
        var v37 uint16 = 565
        fmt.Println(v1, v3)
        _ = v37
    } else {
        for fi38 := 0; fi38 < 5; fi38++ {
            var v39 uint32 = v32
            if len(v2) > 0 {
                v2[0] = fi38
            }
            _ = fi38
            _ = v39
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v35
}
