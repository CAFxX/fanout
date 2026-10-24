package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint; F1 int64; F2 int64 }
type GenS4 struct { F0 bool; F1 string; F2 uint }
func genFn0(p0 int32) bool {
    var v5 uint64 = 445
    if true {
        var v6 uint32 = 18
        p0 = -1
        var v7 float32 = 1e10
        _ = v6
        _ = v7
    } else {
        v8 := []int{0, -179, int(v5), 128}
        v9 := GenS1{F0: 0, F1: ("hello" <= ("\n\t" + "\n\t"))}
        var v10 int16 = -1
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    return ((!(-90 > 127)) || ("foo bar" >= "a"))
}
func genFn1() uint {
    var v11 int = 127
    var v12 int64 = -140
    _ = v11
    _ = v12
    return 696
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 255
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    if true {
        var v13 uint8 = 2
        switch ((("foo bar" + "a") + "foo bar") + "a") {
        case "foo bar":
            v2 += (v2 << ((genFn1()) & 63))
        }
        fmt.Println(v2, v2)
        _ = v13
    } else {
        v2 = 305
        var v14 uint8 = 1
        _ = v14
    }
    v2 = v2
    if true {
        v15 := []int{}
        v16 := GenS2{F0: 4294967295, F1: 2147483647, F2: 0}
        switch 1 {
        case -1:
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 uint = genFn1()
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                v19 := GenS1{F0: 9223372036854775807, F1: ((!false) && (!true))}
                _ = fi17
                _ = v18
                _ = v19
            }
        case 905:
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 float32 = ((((-2.25 + 0.0) * (float32((1e10 / 1.0)) - 0.0)) / ((float32(float64((-2.25 - 1.5))) / 1.5) * 1.5)) * 1.5)
                var v22 int8 = -1
                _ = fi20
                _ = v21
                _ = v22
            }
        case 899:
            switch "foo bar" {
            case "hello":
                for fi23 := 0; fi23 < 16; fi23++ {
                    v24 := GenS3{F0: 662, F1: 0, F2: 745}
                    _ = fi23
                    _ = v24
                }
            case "":
                _ = len(v15)
            default:
                var v25 string = "foo bar"
                _ = v25
            }
        default:
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 int8 = 1
                if (42 <= 4294967295) {
                    v28 := GenS1{F0: -853, F1: false}
                    for fi29 := 0; fi29 < 17; fi29++ {
                        v4[fi29] = 42
                        v27 = (v27 + -1)
                        if ((-1 != 1) || (2 == 2147483647)) {
                            var v30 int64 = -9223372036854775808
                            _ = v30
                        }
                        _ = fi29
                    }
                    _ = v28
                } else {
                    var v31 uint = 255
                    _ = v31
                }
                var v32 float32 = 1e10
                _ = fi26
                _ = v27
                _ = v32
            }
        }
        _ = v15
        _ = v16
    }
    v33 := GenS1{F0: -1, F1: true}
    fmt.Println(v2)
    switch 32767 {
    case 127:
        v34 := []int{-9223372036854775808, 2147483647, 4294967295, 32768}
        _ = v34
    case 128:
        switch -9223372036854775808 {
        case 42:
            fmt.Println(v2)
        case -2147483648:
            v35 := GenS4{F0: (true || true), F1: "a", F2: 643}
            _ = v35
        case 32768:
            if true {
                var v36 uint16 = v2
                v2 = (v2 & v36)
                fmt.Println(v2, v2)
                _ = v36
            }
        default:
            switch -9223372036854775808 {
            case 2147483647:
                if false {
                    v2 = v2
                    v37 := GenS4{F0: (v2 > (v2 * (v2 / (((^(v2))) | 1)))), F1: ("\n\t" + "a"), F2: 1}
                    _ = v37
                }
            default:
                v2 = 1
            }
        }
    default:
        v4[-1] = 128
    }
    v38 := GenS1{F0: -703, F1: (0.1 <= 1.5)}
    v39 := []int{-644, -9223372036854775808, 9223372036854775807, 68}
    fmt.Println(v2)
    for fi40 := 0; fi40 < 19; fi40++ {
        _ = len(v3)
        var v41 uint16 = (v2 + 156)
        fmt.Println(v2, v2, v41)
        _ = fi40
        _ = v41
    }
    var v42 uint16 = 2
    var v43 uint = genFn1()
    var v44 uint8 = 36
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v44
}
