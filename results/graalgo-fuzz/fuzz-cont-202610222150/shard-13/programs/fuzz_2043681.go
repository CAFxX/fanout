package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := []int{128, 128}
    var v2 int64 = -833
    var v3 int32 = 0
    v4 := []uint32{42, 38, 2147483647}
    var v5 int32 = (v3 << ((4294967295) & 63))
    var v6 float64 = 0.1
    var v7 bool = (false || (false || true))
    var v8 bool = (42 != 57)
    var v9 uint32 = 65535
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint = 221
        _ = fi10
        _ = v11
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    if false {
        v12 := GenS1{F0: 0, F1: 42, F2: -614}
        v2 -= (v2 * v2)
        fmt.Println(v5, v5, v6)
        _ = v12
    }
    if false {
        if (false && v8) {
            for fi13 := 0; fi13 < 18; fi13++ {
                v8 = (1 > 1)
                if ((746 == 255) && (-2.25 == float32(1e10))) {
                    v8 = (!false)
                    v14 := GenS2{F0: 2, F1: v2, F2: 255}
                    _ = v14
                } else {
                    if v8 {
                        if (!v8) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            var v15 uint32 = 4294967295
                            var v16 uint = 132
                            switch "a" {
                            case "a":
                                var v17 int = (fi13 / ((fi13) | 1))
                                _ = v17
                            case "foo bar":
                                var v18 uint = 0
                                _ = v18
                            case "hello":
                                v19 := GenS3{F0: 391}
                                _ = v19
                            }
                            _ = v15
                            _ = v16
                        }
                    } else {
                        if (!(0 != 18446744073709551615)) {
                            var v20 uint64 = 0
                            switch fi13 {
                            case -2147483648:
                                v21 := GenS2{F0: 255, F1: (v2 | 9223372036854775807), F2: 2}
                                _ = v21
                            case 0:
                                if len(v1) > 0 {
                                    v1[0] = (fi13 + fi13)
                                }
                            }
                            var v22 int64 = (-(v2))
                            _ = v20
                            _ = v22
                        } else {
                            fmt.Println(v5, v9)
                        }
                        for fi23 := 0; fi23 < 11; fi23++ {
                            v3 += v3
                            var v24 int16 = 42
                            _ = fi23
                            _ = v24
                        }
                        v5 = v5
                    }
                    v9 = (v9 << ((585) & 63))
                    var v25 int8 = 108
                    _ = v25
                }
                fmt.Println(v8)
                _ = fi13
            }
        }
        for fi26 := 0; fi26 < 15; fi26++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (!v7) {
                v6 -= ((((0.0 + v6) / v6) + v6) - 1.5)
                var v27 int8 = 0
                _ = v27
            } else {
                v8 = (127 >= 0)
            }
            _ = fi26
        }
    }
    var v28 int16 = -486
    if len(v1) > 0 {
        v1[0] = -757
    }
    v5 = v3
    if len(v4) > 0 {
        v4[0] = v9
    }
    if (!(v6 < v6)) {
        fmt.Println(v9, v2, v28)
    } else {
        var v29 float64 = v6
        fmt.Println(v3)
        var v30 int32 = v3
        _ = v29
        _ = v30
    }
    switch 4294967295 {
    case 217:
        if len(v4) > 0 {
            v4[0] = (^(v9))
        }
    default:
        v31 := GenS1{F0: v9, F1: 899, F2: 2147483647}
        _ = v31
    }
    v3 += 2147483647
    for fi32 := 0; fi32 < 8; fi32++ {
        v33 := GenS1{F0: (v9 << ((4294967295) & 63)), F1: 0, F2: -1}
        _ = fi32
        _ = v33
    }
    switch ("" + "\n\t") {
    case "hello":
        v34 := GenS1{F0: (^(v9)), F1: 1, F2: 42}
        _ = v34
    case "\n\t":
        for fi35 := 0; fi35 < 3; fi35++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v28 = v28
            fmt.Println(v7, v5)
            _ = fi35
        }
    case "foo bar":
        fmt.Println(v28, v6, v7)
    default:
        if len(v1) > 0 {
            v1[0] = 1
        }
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        v37 := GenS3{F0: 661}
        v3 -= v5
        switch "a" {
        case "a":
            v8 = v7
        case "foo bar":
            if false {
                for fi38 := 0; fi38 < 11; fi38++ {
                    var v39 uint32 = 2
                    _ = fi38
                    _ = v39
                }
                if v8 {
                    var v40 int = (fi36 / ((fi36) | 1))
                    fmt.Println(fi36, v6, v40)
                    _ = v40
                }
            }
        case "hello":
            var v41 uint32 = v9
            _ = v41
        }
        _ = fi36
        _ = v37
    }
    var v42 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
