package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 string; F1 int64; F2 int32 }
func genFn0(p0 int) uint8 {
    if true {
        v4 := []int{p0, 32768, (^(p0)), p0}
        var v5 int32 = 1
        p0 = -130
        _ = v4
        _ = v5
    }
    fmt.Println(p0)
    var v6 int16 = 1000
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS2{F0: 0, F1: -9223372036854775808}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 string) uint32 {
    switch ("\n\t" + p0) {
    case "":
        for fi10 := 0; fi10 < 12; fi10++ {
            p0 = p0
            _ = fi10
        }
    case "\n\t":
        if false {
            var v11 int = -1
            var v12 string = p0
            _ = v11
            _ = v12
        }
    case "foo bar":
        switch 127 {
        case 127:
            p0 = "a"
        default:
            v13 := []int{32768, 1, 410}
            _ = v13
        }
    default:
        var v14 int16 = -662
        _ = v14
    }
    v15 := []int{903}
    v16 := GenS2{F0: 407, F1: -558}
    if false {
        var v17 uint8 = 1
        var v18 int = 4294967295
        var v19 int16 = 1000
        _ = v17
        _ = v18
        _ = v19
    } else {
        v20 := GenS2{F0: 2, F1: 128}
        if (false && false) {
            fmt.Println(p0)
            var v21 int8 = 0
            _ = v21
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            p0 = ("a" + p0)
            p0 = (p0 + (p0 + (p0 + p0)))
            _ = fi22
        }
        _ = v20
    }
    _ = p0
    _ = v15
    _ = v16
    return 1
}

func main() {
    var v1 uint16 = 440
    var v2 string = "\n\t"
    v3 := []uint32{770}
    v23 := GenS2{F0: 65535, F1: 1}
    var v24 float32 = 0.0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        var v25 bool = false
        var v26 string = v2
        _ = v25
        _ = v26
    } else {
        v27 := GenS1{F0: 4294967295}
        var v28 int32 = 42
        _ = v27
        _ = v28
    }
    v29 := GenS2{F0: 0, F1: int(v1)}
    switch 450 {
    case 255:
        v1 = v1
    }
    v1 = 2
    if (true && true) {
        var v30 int8 = 100
        _ = v30
    } else {
        for fi31 := 0; fi31 < 6; fi31++ {
            v32 := GenS3{F0: -816, F1: -1}
            v33 := GenS4{F0: (v2 + "a"), F1: 1, F2: 2147483647}
            v1 = 32767
            _ = fi31
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 20; fi34++ {
            v3 = append(v3, 963)
            switch int(v1) {
            case 9223372036854775807:
                switch fi34 {
                case -888:
                    if (2 >= 916) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v35 int64 = 312
                        var v36 int8 = 1
                        _ = v35
                        _ = v36
                    }
                case 0:
                    _ = len(v3)
                }
            case -9223372036854775808:
                v24 += 0.0
            default:
                if len(v3) > 0 {
                    v3[0] = 1
                }
            }
            _ = len(v3)
            _ = fi34
        }
    }
    _ = len(v3)
    var v37 int = -754
    v1 -= (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v37
}
