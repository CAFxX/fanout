package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
func genFn0(p0 int8) int8 {
    if (1 <= 4294967295) {
        if (true || (853 == -1)) {
            switch "" {
            case "hello":
                switch 32767 {
                case 0:
                    p0 = -2
                default:
                    var v6 int16 = -549
                    _ = v6
                }
            case "\n\t":
                p0 = (p0 | (p0 % (((p0 & 1)) | 1)))
            }
            fmt.Println(p0)
        }
        var v7 float64 = 0.0
        v7 = v7
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 3; fi8++ {
            p0 -= (p0 | p0)
            _ = fi8
        }
        fmt.Println(p0)
        v9 := []int{128, 866, -2147483648, 925}
        _ = v9
    }
    _ = p0
    return (p0 << ((412) & 63))
}

func main() {
    var v1 int64 = 0
    var v2 bool = (4294967295 >= -349)
    var v3 float64 = float64(((0.0 / -2.25) + 1.5))
    v4 := make(map[int]int)
    var v5 int32 = -268
    if v2 {
        v2 = (!true)
        var v10 int32 = (v5 & 0)
        if v2 {
            var v11 uint8 = 127
            if v2 {
                fmt.Println(v1, v3, v5)
                if (4294967295 != 65535) {
                    var v12 float64 = 1e10
                    v12 = v12
                    var v13 int32 = v5
                    _ = v12
                    _ = v13
                }
                if (32767 < uint16(v11)) {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        if v2 {
                            v15 := GenS1{F0: 65535, F1: ((!(!false)) && true), F2: 65535}
                            fmt.Println(fi14, v5, v5)
                            v16 := []int{fi14, fi14, (fi14 * (fi14 >> ((2) & 63))), fi14}
                            _ = v15
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 19; fi17++ {
                            var v18 int64 = (v1 % ((v1) | 1))
                            var v19 int32 = (^(v5))
                            v20 := GenS1{F0: 0, F1: v2, F2: 42}
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        _ = fi14
                    }
                }
            } else {
                switch "hello" {
                case "\n\t":
                    for fi21 := 0; fi21 < 5; fi21++ {
                        var v22 int32 = -1
                        v1 += int64(fi21)
                        v1 -= v1
                        _ = fi21
                        _ = v22
                    }
                case "hello":
                    v3 = v3
                }
                fmt.Println(v2)
                v3 = 3.14159
            }
            fmt.Println(v5, v5, v11)
            _ = v11
        } else {
            if (true && (!v2)) {
                for fi23 := 0; fi23 < 12; fi23++ {
                    var v24 uint32 = 65535
                    _ = fi23
                    _ = v24
                }
            }
        }
        _ = v10
    } else {
        var v25 uint16 = 32767
        var v26 int16 = 0
        _ = v25
        _ = v26
    }
    v27 := []int{32767, 128, 42, 797}
    if false {
        if false {
            v1 = (-(v1))
            var v28 uint = 2
            _ = v28
        }
    }
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        delete(v4, fi29)
        delete(v4, fi29)
        _ = fi29
    }
    fmt.Println(v1)
    v1 += v1
    var v30 string = "a"
    v31 := GenS1{F0: 2147483648, F1: (!(264 < 507)), F2: 65535}
    v27 = append(v27, 0)
    for fi32 := 0; fi32 < 3; fi32++ {
        v27 = append(v27, 356)
        v3 += v3
        _ = fi32
    }
    switch "" {
    case "":
        var v33 int8 = genFn0(0)
        _ = v33
    case "a":
        v27 = append(v27, 2147483647)
    default:
        var v34 uint8 = 42
        _ = v34
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        v36 := GenS1{F0: 372, F1: (1 <= 42), F2: 42}
        var v37 int64 = v1
        _ = len(v4)
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v4
    _ = v5
}
