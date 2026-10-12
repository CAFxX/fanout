package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int64; F2 int8 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 uint16; F1 int8; F2 int16 }
type GenS8 struct { F0 uint8; F1 int16 }

func main() {
    var v1 int = 0
    var v2 float64 = ((0.0 - 1.5) * float64((float32((1e10 - ((0.1 / 1e10) + 1.5))) + 0.1)))
    var v3 string = ("hello" + "\n\t")
    v4 := []string{}
    var v5 int = v1
    var v6 uint64 = 720
    var v7 int8 = 1
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(v3, v5, fi8)
            var v10 uint = 875
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            v7 += (v7 >> ((4294967295) & 63))
            v12 := GenS1{F0: 52, F1: (255 > 42)}
            v2 -= (((-2.25 + (v2 / (3.14159 * v2))) * (v2 - v2)) * -2.25)
            _ = fi11
            _ = v12
        }
        _ = fi8
    }
    v13 := GenS2{F0: 1}
    v2 = v2
    var v14 bool = false
    v15 := GenS3{F0: 0}
    if v14 {
        v4 = append(v4, "a")
        for fi16 := 0; fi16 < 5; fi16++ {
            v6 = (v6 >> ((874) & 63))
            for fi17 := 0; fi17 < 17; fi17++ {
                v18 := GenS3{F0: 32767}
                fmt.Println(v6, v3, fi17)
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        switch (^(v1)) {
        case 4294967295:
            var v19 uint64 = (v6 * 5)
            _ = v19
        case -617:
            if v14 {
                if true {
                    var v20 int = v1
                    if (v14 && ((18446744073709551615 == 52) || true)) {
                        var v21 int64 = -9223372036854775808
                        _ = v21
                    } else {
                        v22 := GenS2{F0: 10}
                        v23 := GenS2{F0: 2}
                        _ = v22
                        _ = v23
                    }
                    _ = v20
                }
            } else {
                if (1 < -473) {
                    v2 = (1e10 * 1.5)
                    v7 -= 100
                    fmt.Println(v3, v3)
                } else {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        for fi25 := 0; fi25 < 16; fi25++ {
                            v4 = append(v4, ("" + v3))
                            _ = fi25
                        }
                        if len(v4) > 0 {
                            v4[0] = ("hello" + "")
                        }
                        v26 := GenS4{F0: 1, F1: 138, F2: v7}
                        _ = fi24
                        _ = v26
                    }
                }
            }
        case 140:
            if false {
                v27 := GenS1{F0: -391, F1: (526 > 65535)}
                var v28 float64 = v2
                _ = v27
                _ = v28
            }
        default:
            if (1 != -690) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                fmt.Println(v5, v14)
                var v29 float64 = (v2 + (v2 + 0.1))
                _ = v29
            }
        }
    } else {
        var v30 int32 = 578
        v4 = append(v4, v3)
        for fi31 := 0; fi31 < 6; fi31++ {
            v32 := GenS1{F0: v5, F1: v14}
            _ = len(v4)
            fmt.Println(v5, v30, v1)
            _ = fi31
            _ = v32
        }
        _ = v30
    }
    fmt.Println(v1, v6)
    for fi33 := 0; fi33 < 14; fi33++ {
        v2 += v2
        if (0 > 9223372036854775807) {
            v34 := GenS3{F0: 288}
            var v35 uint32 = 4294967295
            v1 = fi33
            _ = v34
            _ = v35
        } else {
            var v36 float64 = v2
            v3 = (v3 + "hello")
            _ = v36
        }
        v37 := GenS5{F0: -686, F1: (v7 << ((42) & 63))}
        _ = fi33
        _ = v37
    }
    v38 := GenS6{F0: 416, F1: 1}
    var v39 uint32 = 65535
    var v40 string = "\n\t"
    v7 -= (v7 * (v7 & (v7 % (((v7 + (v7 + 64))) | 1))))
    v41 := GenS7{F0: 1, F1: (v7 | (v7 | -128)), F2: 42}
    var v42 uint32 = 2147483648
    switch -9223372036854775808 {
    case 2147483647:
        for fi43 := 0; fi43 < 8; fi43++ {
            var v44 uint = 1
            _ = fi43
            _ = v44
        }
    case 4294967295:
        v45 := GenS8{F0: 127, F1: 237}
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
