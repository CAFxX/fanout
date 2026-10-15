package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int32; F2 uint32 }
func genFn0(p0 int16) int32 {
    var v5 float32 = 0.0
    var v6 int16 = (^(p0))
    var v7 string = "\n\t"
    if false {
        v8 := GenS1{F0: 0, F1: v6, F2: 612}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 11; fi9++ {
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS1{F0: 255, F1: -32768, F2: 42}
                switch v7 {
                case "a":
                    fmt.Println(v6, v7)
                default:
                    v12 := GenS1{F0: 0, F1: (v6 << ((1) & 63)), F2: 2147483648}
                    _ = v12
                }
                v5 -= (1e10 + float32(1.5))
                p0 = p0
                v7 = v7
                var v13 int16 = (p0 % ((0) | 1))
                _ = fi10
                _ = v11
                _ = v13
            }
            _ = fi9
        }
        var v14 uint = 421
        v15 := GenS1{F0: 88, F1: p0, F2: 1}
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int = 128
    var v2 float64 = 3.14159
    var v3 uint16 = 32767
    var v4 int32 = 0
    v3 -= (v3 | (v3 | 102))
    if (false || true) {
        var v16 float32 = 1.5
        v17 := GenS2{F0: 42, F1: 2, F2: v1}
        _ = v16
        _ = v17
    }
    var v18 string = ("" + "hello")
    switch v18 {
    case "foo bar":
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS1{F0: 108, F1: -32768, F2: 2147483647}
            v21 := GenS1{F0: 63, F1: -357, F2: 1}
            _ = fi19
            _ = v20
            _ = v21
        }
    case "hello":
        fmt.Println(v18, v3)
    case "":
        v22 := GenS3{F0: (v3 | (v3 % (((v3 - 842)) | 1)))}
        _ = v22
    }
    var v23 uint32 = 308
    if (128 == 86) {
        if (9223372036854775807 != -1) {
            var v24 int32 = (v4 % ((int32(v23)) | 1))
            _ = v24
        } else {
            if false {
                fmt.Println(v2)
                switch "a" {
                case "foo bar":
                    if (!false) {
                        var v25 bool = true
                        var v26 int16 = int16(v4)
                        _ = v25
                        _ = v26
                    } else {
                        v4 += (v4 * (v4 & 0))
                    }
                }
            } else {
                v27 := []int{v1, -743, v1}
                if (!true) {
                    var v28 uint64 = 4294967295
                    _ = v28
                }
                for fi29 := 0; fi29 < 2; fi29++ {
                    if (!((v23 != (^(v23))) || true)) {
                        var v30 int16 = -32768
                        _ = v30
                    } else {
                        v2 += 1.5
                        fmt.Println(v3, v18)
                    }
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 int16 = 447
                        _ = fi31
                        _ = v32
                    }
                    v27 = append(v27, (v1 & fi29))
                    _ = fi29
                }
                _ = v27
            }
            v3 = v3
            v33 := []int{v1, 994, -998}
            _ = v33
        }
        switch ((v18 + "\n\t") + v18) {
        case "a":
            var v34 bool = true
            _ = v34
        default:
            v1 = (v1 << ((1) & 63))
        }
    } else {
        for fi35 := 0; fi35 < 8; fi35++ {
            v36 := []int{(^(fi35)), 32767, (v1 % ((-893) | 1)), v1}
            _ = fi35
            _ = v36
        }
        v2 = v2
        switch (v1 << ((623) & 63)) {
        case 42:
            if (!((!(float32(v2) <= 1e10)) && (2 != 255))) {
                v37 := GenS2{F0: 128, F1: 65535, F2: 128}
                v3 = (v3 | (v3 | (v3 >> ((42) & 63))))
                fmt.Println(v23, v23, v4)
                _ = v37
            } else {
                if true {
                    for fi38 := 0; fi38 < 2; fi38++ {
                        for fi39 := 0; fi39 < 8; fi39++ {
                            fmt.Println(v3, v18)
                            _ = fi39
                        }
                        _ = fi38
                    }
                    var v40 uint32 = (v23 >> ((4294967295) & 63))
                    if (!(120 >= 42)) {
                        if (1e10 > 1e10) {
                            v23 -= v40
                            var v41 float32 = 1e10
                            for fi42 := 0; fi42 < 17; fi42++ {
                                for fi43 := 0; fi43 < 20; fi43++ {
                                    fmt.Println(v41)
                                    v1 -= int(v4)
                                    v44 := GenS1{F0: 2, F1: -1, F2: 65535}
                                    v45 := []int{}
                                    fmt.Println(v3)
                                    var v46 uint32 = (v40 & v40)
                                    _ = fi43
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                                var v47 int = (-(v1))
                                _ = fi42
                                _ = v47
                            }
                            _ = v41
                        }
                    }
                    _ = v40
                }
                var v48 uint = 624
                v49 := GenS4{F0: -694, F1: (v4 % ((genFn0(682)) | 1)), F2: v23}
                _ = v48
                _ = v49
            }
        }
    }
    var v50 int16 = -32768
    v4 += genFn0(v50)
    var v51 uint32 = (^(v23))
    v23 = 872
    if true {
        v52 := []int{(v1 / (((v1 % ((v1) | 1))) | 1)), -664, v1, 42}
        for fi53 := 0; fi53 < 11; fi53++ {
            if len(v52) > 0 {
                v52[0] = (^(fi53))
            }
            fmt.Println(v50, v18, v18)
            v54 := GenS2{F0: uint8(v3), F1: uint(v3), F2: (v1 & 421)}
            _ = fi53
            _ = v54
        }
        _ = v52
    } else {
        var v55 uint = uint(v3)
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v50
    _ = v51
}
