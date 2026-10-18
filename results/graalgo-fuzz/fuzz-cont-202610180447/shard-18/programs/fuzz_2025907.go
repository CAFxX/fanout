package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int }
func genFn0(p0 uint8, p1 uint16) uint64 {
    p0 = 255
    p1 = p1
    v4 := []int{}
    var v5 uint64 = 1
    if true {
        fmt.Println(v5, p1)
        v5 = v5
        v6 := GenS1{F0: -176}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return v5
}
func genFn1(p0 int32) uint32 {
    for fi7 := 0; fi7 < 10; fi7++ {
        if ((float64(float32(1.5)) * ((3.14159 * (1e10 / 1.0)) - 3.14159)) <= (0.0 + -2.25)) {
            var v8 int64 = -1
            _ = v8
        } else {
            switch "hello" {
            case "foo bar":
                fmt.Println(p0, fi7)
            case "a":
                if true {
                    var v9 int8 = 72
                    _ = v9
                } else {
                    var v10 int8 = 0
                    var v11 uint = 42
                    _ = v10
                    _ = v11
                }
            case "hello":
                p0 -= 2147483647
            }
            var v12 uint64 = 161
            _ = v12
        }
        _ = fi7
    }
    p0 += (^(p0))
    p0 += (p0 >> ((965) & 63))
    if (0 != 128) {
        p0 = p0
        var v13 int8 = -90
        switch 2147483647 {
        case 2147483647:
            if false {
                if (!(("foo bar" + "hello") != "\n\t")) {
                    var v14 uint8 = 52
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v13 += (v13 % (((v13 % (((v13 << ((42) & 63))) | 1))) | 1))
                        v16 := []int{fi15, fi15, fi15}
                        _ = fi15
                        _ = v16
                    }
                    p0 += p0
                }
                for fi17 := 0; fi17 < 15; fi17++ {
                    if true {
                        v18 := []int{}
                        var v19 int16 = -32768
                        v18 = append(v18, int(v19))
                        _ = v18
                        _ = v19
                    }
                    switch "\n\t" {
                    case "a":
                        for fi20 := 0; fi20 < 17; fi20++ {
                            v21 := GenS1{F0: 506}
                            _ = fi20
                            _ = v21
                        }
                    }
                    _ = fi17
                }
                fmt.Println(p0, v13)
            }
        default:
            fmt.Println(p0, v13)
        }
        _ = v13
    }
    _ = p0
    return 0
}

func main() {
    var v1 int = 127
    var v2 uint = 42
    var v3 int = 2147483647
    var v22 uint = v2
    for fi23 := 0; fi23 < 17; fi23++ {
        v2 += (v2 % (((v22 | 18446744073709551615)) | 1))
        _ = fi23
    }
    v24 := GenS1{F0: 42}
    if (false && false) {
        if true {
            var v25 int64 = -1
            v26 := []int{123, v3, v3}
            _ = v25
            _ = v26
        } else {
            v27 := GenS1{F0: 0}
            _ = v27
        }
        var v28 int16 = -294
        _ = v28
    }
    var v29 uint16 = 884
    if (!false) {
        v3 += 128
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        for fi31 := 0; fi31 < 18; fi31++ {
            v2 = (v2 >> ((v22) & 63))
            _ = fi31
        }
        v32 := GenS1{F0: 42}
        for fi33 := 0; fi33 < 8; fi33++ {
            fmt.Println(fi33, v1)
            _ = fi33
        }
        _ = fi30
        _ = v32
    }
    v34 := GenS2{F0: 32767, F1: (v3 * 32768)}
    v1 -= 128
    v35 := GenS2{F0: (v29 | v29), F1: -165}
    v36 := []int{}
    var v37 int16 = 80
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
