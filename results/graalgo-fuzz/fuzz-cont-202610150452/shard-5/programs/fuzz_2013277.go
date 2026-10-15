package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int16 }
func genFn0(p0 string, p1 int16) uint {
    p0 = p0
    for fi4 := 0; fi4 < 11; fi4++ {
        if false {
            switch "" {
            case "hello":
                var v5 uint8 = 127
                _ = v5
            default:
                for fi6 := 0; fi6 < 16; fi6++ {
                    p0 = p0
                    var v7 bool = false
                    _ = fi6
                    _ = v7
                }
            }
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(fi4, p0, fi8)
                v9 := []int{}
                p1 = p1
                _ = fi8
                _ = v9
            }
            var v10 float32 = 0.0
            _ = v10
        }
        switch (fi4 ^ 0) {
        case 128:
            switch fi4 {
            case 128:
                fmt.Println(p1, p0, fi4)
            case 2147483647:
                for fi11 := 0; fi11 < 18; fi11++ {
                    if (0 > 127) {
                        switch 325 {
                        case -9223372036854775808:
                            var v12 bool = (749 > -2147483648)
                            _ = v12
                        case -501:
                            if true {
                                fmt.Println(fi11, p1)
                                switch 42 {
                                case -2147483648:
                                    fmt.Println(fi11, p1, fi4)
                                }
                                p0 = p0
                            }
                        case -2147483648:
                            fmt.Println(fi11, fi11)
                        }
                        if (!(487 == 4294967295)) {
                            var v13 uint8 = 0
                            var v14 int = -9223372036854775808
                            fmt.Println(p1, fi4, fi4)
                            v15 := GenS1{F0: 1}
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                    }
                    v16 := []int{fi11, (fi11 / ((fi4) | 1))}
                    _ = fi11
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 uint = 2
                    _ = fi17
                    _ = v18
                }
            }
        }
        _ = fi4
    }
    v19 := []int{32768, 4294967295, 419}
    _ = p0
    _ = p1
    _ = v19
    return 1
}
func genFn1(p0 uint32, p1 int) int16 {
    fmt.Println(p1)
    fmt.Println(p1, p0)
    p0 = p0
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 int32 = 42
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    return 1000
}

func main() {
    var v1 string = ("hello" + "")
    v2 := []uint32{}
    var v3 bool = (((0.0 * 1e10) < ((0.0 - 0.1) / -2.25)) && false)
    var v22 float32 = (float32(0.1) - (1.5 + 1.5))
    v1 = "a"
    v22 += (0.0 + (0.0 - v22))
    v22 += (float32((1e10 - (3.14159 - 0.0))) + float32(3.14159))
    v1 = v1
    var v23 uint8 = 177
    var v24 uint8 = (^(v23))
    switch 329 {
    case -1:
        var v25 float64 = (0.0 / 1e10)
        _ = v25
    case 2147483647:
        switch v1 {
        case "foo bar":
            for fi26 := 0; fi26 < 16; fi26++ {
                var v27 uint8 = v24
                var v28 float64 = 0.1
                _ = fi26
                _ = v27
                _ = v28
            }
        case "":
            switch -234 {
            case 32768:
                var v29 uint64 = 2
                _ = v29
            case 4294967295:
                var v30 int = -1
                _ = v30
            case 127:
                for fi31 := 0; fi31 < 20; fi31++ {
                    if true {
                        v22 = v22
                    }
                    _ = fi31
                }
            default:
                switch "" {
                case "hello":
                    var v32 bool = (!((false || true) && v3))
                    _ = v32
                case "":
                    var v33 bool = (909 < 4294967295)
                    _ = v33
                case "a":
                    var v34 int8 = 42
                    _ = v34
                default:
                    var v35 int32 = -1
                    _ = v35
                }
            }
        case "\n\t":
            fmt.Println(v1, v24, v22)
        }
    default:
        v36 := GenS2{F0: 0, F1: 1000}
        _ = v36
    }
    _ = len(v2)
    if (true || v3) {
        if (2 == 469) {
            if len(v2) > 0 {
                v2[0] = 391
            }
            var v37 int32 = 949
            for fi38 := 0; fi38 < 17; fi38++ {
                if (v3 || true) {
                    fmt.Println(fi38, fi38, v23)
                } else {
                    fmt.Println(v23, v3, v24)
                }
                v39 := GenS3{F0: 875, F1: v23, F2: genFn1(2147483648, (fi38 << ((0) & 63)))}
                _ = fi38
                _ = v39
            }
            _ = v37
        } else {
            var v40 uint = genFn0(v1, 42)
            v22 -= v22
            _ = v40
        }
        if len(v2) > 0 {
            v2[0] = 0
        }
        var v41 uint32 = 1
        _ = v41
    }
    var v42 uint8 = v24
    for fi43 := 0; fi43 < 9; fi43++ {
        v44 := GenS3{F0: 2147483648, F1: (v23 - 135), F2: -353}
        for fi45 := 0; fi45 < 13; fi45++ {
            var v46 int32 = -1
            fmt.Println(v3)
            _ = fi45
            _ = v46
        }
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v42
}
