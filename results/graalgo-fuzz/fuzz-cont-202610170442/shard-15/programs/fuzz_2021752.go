package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint16 }

func main() {
    var v1 string = "hello"
    var v2 int8 = -113
    var v3 uint32 = 318
    var v4 int16 = 32767
    var v5 string = ("foo bar" + v1)
    if (1 <= 444) {
        var v6 string = v5
        fmt.Println(v2)
        _ = v6
    } else {
        v4 -= (v4 % (((v4 * v4)) | 1))
    }
    var v7 uint32 = 65535
    var v8 int16 = v4
    v9 := []int{32768, 2147483647}
    if (0.1 != 1.5) {
        var v10 int16 = (v8 + v4)
        _ = v10
    }
    v5 = (v1 + "a")
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS1{F0: v7, F1: 0, F2: true}
        _ = fi11
        _ = v12
    }
    if true {
        for fi13 := 0; fi13 < 15; fi13++ {
            if true {
                v4 = (v4 >> ((1) & 63))
                var v14 float64 = 1.5
                _ = v14
            } else {
                switch "\n\t" {
                case "foo bar":
                    var v15 uint64 = 782
                    _ = v15
                }
                if ((("" + (v5 + ("foo bar" + "foo bar"))) + "foo bar") >= v5) {
                    switch (v1 + v1) {
                    case "\n\t":
                        v16 := GenS1{F0: v7, F1: -1, F2: (-9223372036854775808 < 31)}
                        _ = v16
                    case "a":
                        v3 = v3
                    default:
                        var v17 int32 = -1
                        _ = v17
                    }
                    v18 := GenS1{F0: v3, F1: 2147483647, F2: false}
                    var v19 uint16 = 0
                    _ = v18
                    _ = v19
                }
            }
            fmt.Println(fi13, fi13)
            _ = fi13
        }
        var v20 float64 = (0.0 - float64(0.0))
        if false {
            if ((^(v4)) < v4) {
                fmt.Println(v7, v4)
                for fi21 := 0; fi21 < 1; fi21++ {
                    var v22 uint8 = 255
                    if (true || true) {
                        v5 = v1
                    }
                    _ = fi21
                    _ = v22
                }
                if ((474 > 478) && false) {
                    var v23 bool = true
                    for fi24 := 0; fi24 < 6; fi24++ {
                        var v25 uint32 = v7
                        v2 -= (v2 / ((42) | 1))
                        var v26 uint64 = 4294967295
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    _ = v23
                }
            }
        } else {
            for fi27 := 0; fi27 < 3; fi27++ {
                if false {
                    if (-32768 <= v8) {
                        v28 := GenS1{F0: 1, F1: 714, F2: false}
                        _ = v28
                    } else {
                        var v29 uint = 511
                        _ = len(v9)
                        for fi30 := 0; fi30 < 11; fi30++ {
                            v9 = append(v9, fi30)
                            _ = fi30
                        }
                        _ = v29
                    }
                    for fi31 := 0; fi31 < 5; fi31++ {
                        v9 = append(v9, (fi31 * fi31))
                        var v32 uint16 = 616
                        if true {
                            if len(v9) > 0 {
                                v9[0] = 127
                            }
                            v33 := GenS1{F0: (v7 * v3), F1: -577, F2: true}
                            if false {
                                v34 := GenS2{F0: (v2 - v2)}
                                _ = v34
                            } else {
                                v20 -= v20
                                v35 := GenS3{F0: fi31}
                                switch "\n\t" {
                                case "a":
                                    v2 = v2
                                default:
                                    if false {
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v36 uint16 = 255
                                            _ = v36
                                        }
                                    }
                                }
                                _ = v35
                            }
                            _ = v33
                        }
                        _ = fi31
                        _ = v32
                    }
                } else {
                    var v37 int16 = v4
                    if (0 == uint16(v3)) {
                        var v38 float64 = 1.5
                        for fi39 := 0; fi39 < 6; fi39++ {
                            var v40 float64 = (float64((1.5 / 3.14159)) / -2.25)
                            _ = fi39
                            _ = v40
                        }
                        v20 -= (v38 - (v20 - (v38 + ((1.5 / 0.1) * float64(1e10)))))
                        _ = v38
                    } else {
                        v41 := GenS4{F0: 1, F1: -1, F2: 297}
                        if ((fi27 - fi27) <= 42) {
                            var v42 int8 = v2
                            _ = v42
                        }
                        _ = v41
                    }
                    for fi43 := 0; fi43 < 9; fi43++ {
                        var v44 uint8 = 255
                        _ = fi43
                        _ = v44
                    }
                    _ = v37
                }
                if true {
                    v7 -= v7
                    if len(v9) > 0 {
                        v9[0] = (^(fi27))
                    }
                    if false {
                        var v45 int16 = v8
                        var v46 uint8 = 0
                        _ = v45
                        _ = v46
                    } else {
                        v47 := GenS4{F0: -674, F1: (v8 ^ v8), F2: 868}
                        _ = v47
                    }
                } else {
                    if len(v9) > 0 {
                        v9[0] = fi27
                    }
                }
                _ = fi27
            }
        }
        _ = v20
    } else {
        v7 -= (v3 >> ((255) & 63))
        v48 := GenS5{F0: 24, F1: 291}
        _ = v48
    }
    v2 -= 100
    if (false || (2147483648 < (^(v7)))) {
        if ((42 <= 32767) && false) {
            var v49 uint64 = 4294967295
            for fi50 := 0; fi50 < 16; fi50++ {
                v51 := GenS3{F0: (fi50 & 42)}
                fmt.Println(v4)
                fmt.Println(v2, v4)
                _ = fi50
                _ = v51
            }
            var v52 uint = 520
            _ = v49
            _ = v52
        } else {
            v53 := GenS2{F0: v2}
            var v54 float32 = 1.5
            v1 = v1
            _ = v53
            _ = v54
        }
    } else {
        v55 := GenS6{F0: true, F1: 644}
        _ = v55
    }
    v56 := GenS4{F0: 994, F1: (v4 / (((v8 << ((221) & 63))) | 1)), F2: 341}
    var v57 int = 1
    v8 = -22
    fmt.Println(v57, v2, v3)
    var v58 int = int(v2)
    var v59 int64 = 1
    v60 := GenS4{F0: v59, F1: (v8 % ((32767) | 1)), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v7
    _ = v8
    _ = v9
}
