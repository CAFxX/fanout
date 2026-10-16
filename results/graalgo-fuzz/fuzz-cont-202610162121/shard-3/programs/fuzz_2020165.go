package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int16, p1 int32, p2 uint64) uint8 {
    v5 := GenS1{F0: 4294967295, F1: 2147483648}
    fmt.Println(p1, p0)
    fmt.Println(p1, p2)
    for fi6 := 0; fi6 < 18; fi6++ {
        p1 += p1
        _ = fi6
    }
    v7 := GenS1{F0: 9223372036854775807, F1: 750}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 2
}
func genFn1(p0 bool, p1 bool) uint16 {
    fmt.Println(p0, p0)
    if (793 < 255) {
        p1 = p0
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            p1 = true
            var v9 string = ("\n\t" + "foo bar")
            _ = fi8
            _ = v9
        }
        var v10 bool = (!((p1 && ((-2.25 > -2.25) || p1)) && p1))
        _ = v10
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 65535
    var v3 float64 = (0.1 + (0.1 * float64(1e10)))
    var v4 int32 = 2147483647
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 string = "foo bar"
        v2 = v2
        v13 := GenS1{F0: fi11, F1: 2147483648}
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint8 = genFn0(1000, v4, 893)
        var v16 uint64 = 42
        if (32767 < 42) {
            if (201 == 601) {
                if (uint(v2) >= 1) {
                    if false {
                        for fi17 := 0; fi17 < 3; fi17++ {
                            if false {
                                var v18 float32 = 3.14159
                                _ = v18
                            } else {
                                fmt.Println(fi17)
                            }
                            v19 := GenS1{F0: (fi14 * (fi14 * 32767)), F1: 65535}
                            _ = fi17
                            _ = v19
                        }
                        if (false && false) {
                            for fi20 := 0; fi20 < 20; fi20++ {
                                v2 = 32767
                                if (!false) {
                                    var v21 uint32 = 2147483648
                                    var v22 uint32 = 2
                                    v21 += 2147483647
                                    _ = v21
                                    _ = v22
                                }
                                fmt.Println(fi14, v4, v3)
                                _ = fi20
                            }
                            v23 := GenS2{F0: fi14}
                            var v24 uint64 = (v16 - (^(v16)))
                            _ = v23
                            _ = v24
                        }
                    }
                    var v25 string = "hello"
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 8; fi26++ {
                        var v27 uint8 = genFn0(-32768, 876, (v16 ^ v16))
                        _ = fi26
                        _ = v27
                    }
                    switch "" {
                    case "a":
                        switch 4294967295 {
                        case -2147483648:
                            switch -387 {
                            case 32767:
                                if true {
                                    var v28 int64 = 1
                                    for fi29 := 0; fi29 < 8; fi29++ {
                                        var v30 uint = uint(v16)
                                        _ = fi29
                                        _ = v30
                                    }
                                    v31 := []int{fi14, int(v2), (fi14 - (^(fi14))), 0}
                                    _ = v28
                                    _ = v31
                                } else {
                                    var v32 int64 = -1
                                    var v33 uint8 = (^(v15))
                                    _ = v32
                                    _ = v33
                                }
                            }
                        }
                    case "":
                        for fi34 := 0; fi34 < 14; fi34++ {
                            fmt.Println(v3, v15)
                            _ = fi34
                        }
                    default:
                        v16 += 42
                    }
                    var v35 int = fi14
                    _ = v35
                }
                v2 = 65535
                for fi36 := 0; fi36 < 3; fi36++ {
                    var v37 int32 = v4
                    var v38 int16 = -32768
                    var v39 uint32 = 65535
                    v40 := []int{}
                    _ = fi36
                    _ = v37
                    _ = v38
                    _ = v39
                    _ = v40
                }
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    v2 += (v2 / ((v2) | 1))
    v41 := []int{-1, 2147483647}
    v2 = 42
    var v42 uint64 = 665
    v42 -= (v42 << ((42) & 63))
    for fi43 := 0; fi43 < 12; fi43++ {
        fmt.Println(v42)
        var v44 uint = 2
        for fi45 := 0; fi45 < 1; fi45++ {
            v46 := GenS3{F0: 1000, F1: 2147483647, F2: 106}
            fmt.Println(v2, v4)
            v47 := GenS4{F0: 2, F1: 324, F2: 0}
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = fi43
        _ = v44
    }
    v48 := GenS5{F0: 6, F1: (v2 + v2), F2: 4294967295}
    v42 -= v42
    fmt.Println(v3, v2)
    v49 := GenS6{F0: 127}
    v1[1] = "foo bar"
    if ("\n\t" <= "foo bar") {
        var v50 int = 9223372036854775807
        _ = v50
    }
    var v51 int = -55
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v48
    _ = v49
    _ = v51
}
