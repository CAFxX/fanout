package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int64; F2 int64 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 int16; F1 int16; F2 int32 }
func genFn0(p0 uint, p1 int32, p2 uint) uint64 {
    var v6 bool = false
    v7 := GenS1{F0: 1, F1: "a", F2: 100}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 396
}

func main() {
    var v1 float32 = (-2.25 / 1.0)
    var v2 int = 9223372036854775807
    v3 := []uint32{1, 4294967295}
    var v4 float32 = v1
    v5 := []int{-1, 9223372036854775807, (v2 + (v2 / ((32767) | 1)))}
    switch (^(v2)) {
    case -9223372036854775808:
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: 484, F1: "", F2: -128}
            fmt.Println(v4)
            _ = fi8
            _ = v9
        }
    }
    v4 = (3.14159 / 1.5)
    v1 -= v1
    switch 882 {
    case 220:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case -1:
        switch (v2 & (v2 * -9223372036854775808)) {
        case 1:
            v1 += (((1e10 + (v1 * (v4 * 1.5))) - ((v4 + (v1 - (3.14159 - float32(3.14159)))) / 3.14159)) - (v1 - ((float32(-2.25) * float32(1e10)) * (v1 - 1e10))))
        case 32768:
            for fi10 := 0; fi10 < 4; fi10++ {
                v2 -= (v2 - fi10)
                v11 := GenS2{F0: 42, F1: -9, F2: 1}
                if false {
                    v1 -= ((1e10 / (1e10 / (v4 / v1))) + float32((1e10 - (0.0 - 1e10))))
                    var v12 float32 = float32(1.5)
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
        case 32767:
            fmt.Println(v1, v2)
        default:
            if (true && false) {
                var v13 int32 = 418
                if true {
                    var v14 bool = true
                    _ = v14
                }
                if false {
                    for fi15 := 0; fi15 < 18; fi15++ {
                        for fi16 := 0; fi16 < 4; fi16++ {
                            var v17 uint8 = 1
                            _ = fi16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 17; fi18++ {
                            var v19 uint16 = 42
                            _ = fi18
                            _ = v19
                        }
                        var v20 int32 = (v13 % ((v13) | 1))
                        _ = fi15
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v22 := GenS3{F0: 1}
                        _ = fi21
                        _ = v22
                    }
                }
                _ = v13
            }
        }
    default:
        v23 := GenS4{F0: 1}
        _ = v23
    }
    var v24 int8 = 42
    switch 951 {
    case -810:
        switch "hello" {
        case "":
            var v25 int64 = 1
            _ = v25
        case "\n\t":
            for fi26 := 0; fi26 < 1; fi26++ {
                fmt.Println(v4, v24)
                var v27 int64 = -9223372036854775808
                for fi28 := 0; fi28 < 6; fi28++ {
                    var v29 bool = true
                    var v30 uint = 2
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi26
                _ = v27
            }
        }
    default:
        var v31 uint32 = 237
        _ = v31
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        var v33 int = (v2 % ((-438) | 1))
        v34 := GenS5{F0: 32767, F1: 9223372036854775807, F2: -101}
        for fi35 := 0; fi35 < 2; fi35++ {
            var v36 float64 = float64(v1)
            var v37 uint8 = 42
            var v38 int16 = -536
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    var v39 uint8 = 173
    v40 := GenS6{F0: "\n\t", F1: 0}
    switch ("foo bar" + ("" + ("hello" + ("hello" + ("hello" + "\n\t"))))) {
    case "foo bar":
        fmt.Println(v24)
    case "hello":
        fmt.Println(v2)
    }
    v41 := GenS7{F0: 0, F1: -1, F2: 1}
    var v42 uint64 = genFn0(762, 1, 894)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
