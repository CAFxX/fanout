package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int64; F1 bool }
func genFn0() uint8 {
    var v7 string = ("\n\t" + "")
    for fi8 := 0; fi8 < 12; fi8++ {
        v7 = "\n\t"
        for fi9 := 0; fi9 < 5; fi9++ {
            switch fi9 {
            case -929:
                v7 = "\n\t"
            case -2147483648:
                v10 := GenS1{F0: v7}
                _ = v10
            case 42:
                v11 := []int{(fi9 * fi9), (^(fi8)), fi8}
                _ = v11
            }
            var v12 uint8 = 85
            _ = fi9
            _ = v12
        }
        v13 := []int{fi8, (^(fi8))}
        _ = fi8
        _ = v13
    }
    v7 = v7
    var v14 uint16 = 32767
    _ = v14
    _ = v7
    return 255
}
func genFn1(p0 int8, p1 int64) int8 {
    var v15 uint32 = 2
    fmt.Println(p1, p1, p1)
    v16 := GenS2{F0: 456, F1: (^(v15)), F2: 0}
    v17 := GenS2{F0: 65535, F1: v15, F2: 65535}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    _ = v17
    return 10
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.1
    v3 := []int{9223372036854775807}
    var v4 int64 = -647
    var v5 bool = false
    v6 := []int{127, 2147483647}
    v18 := GenS3{F0: 465, F1: "hello"}
    var v19 uint16 = 833
    v4 = -9223372036854775808
    fmt.Println(v2, v19)
    v1["a"] = 65535
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v20 uint8 = genFn0()
    fmt.Println(v19)
    switch 32768 {
    case 127:
        fmt.Println(v5, v5)
    case 42:
        if (!v5) {
            var v21 uint16 = v19
            v22 := GenS3{F0: 65535, F1: ("\n\t" + "hello")}
            _ = v21
            _ = v22
        }
    }
    var v23 uint8 = (v20 | v20)
    var v24 int = 42
    var v25 int16 = -619
    var v26 uint8 = (^(v23))
    if true {
        switch "hello" {
        case "a":
            for fi27 := 0; fi27 < 13; fi27++ {
                v4 -= v4
                _ = fi27
            }
        default:
            var v28 int8 = 100
            _ = v28
        }
        if (!v5) {
            if (((v4 ^ v4) == v4) && true) {
                if (true && (455 != (v19 | v19))) {
                    if v5 {
                        fmt.Println(v5)
                        _ = len(v6)
                        if len(v3) > 0 {
                            v3[0] = (v24 ^ (v24 / ((v24) | 1)))
                        }
                    } else {
                        delete(v1, "foo bar")
                    }
                } else {
                    if false {
                        var v29 int32 = 42
                        v5 = v5
                        _ = v29
                    }
                }
                fmt.Println(v2, v20, v19)
            } else {
                for fi30 := 0; fi30 < 16; fi30++ {
                    var v31 int64 = v4
                    v32 := GenS2{F0: 946, F1: 196, F2: 65535}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        }
    }
    if (false && v5) {
        if v5 {
            v25 = (v25 / (((v25 % ((v25) | 1))) | 1))
        } else {
            v33 := GenS2{F0: 287, F1: 523, F2: 4294967295}
            fmt.Println(v23, v2, v24)
            _ = v33
        }
        v23 = v23
        v24 = v24
    } else {
        switch (v24 | (v24 << ((0) & 63))) {
        case -9223372036854775808:
            switch v24 {
            case 127:
                v34 := GenS4{F0: (-(v4)), F1: ((^(v25)) > v25)}
                _ = v34
            default:
                switch v24 {
                case 9223372036854775807:
                    for fi35 := 0; fi35 < 14; fi35++ {
                        var v36 uint16 = (v19 % ((v19) | 1))
                        var v37 int32 = -616
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                case 42:
                    if (uint64(v23) < 18446744073709551615) {
                        if v5 {
                            fmt.Println(v19, v23, v23)
                        }
                        fmt.Println(v25, v19, v23)
                        var v38 uint8 = (v20 % ((2) | 1))
                        _ = v38
                    } else {
                        v39 := GenS3{F0: 2147483648, F1: ""}
                        _ = v39
                    }
                }
            }
        case 127:
            for fi40 := 0; fi40 < 7; fi40++ {
                fmt.Println(v20, v24)
                var v41 uint = 42
                _ = fi40
                _ = v41
            }
        case -2147483648:
            switch (^(v24)) {
            case 464:
                for fi42 := 0; fi42 < 18; fi42++ {
                    for fi43 := 0; fi43 < 15; fi43++ {
                        var v44 float32 = 0.0
                        var v45 float64 = -2.25
                        var v46 uint8 = 42
                        _ = fi43
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                    _ = fi42
                }
            case 922:
                var v47 int8 = -1
                _ = v47
            }
        }
        fmt.Println(v26)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
