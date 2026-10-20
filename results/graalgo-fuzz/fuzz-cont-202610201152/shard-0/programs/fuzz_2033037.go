package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int8; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint32; F2 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 int8; F2 bool }
type GenS7 struct { F0 uint64 }

func main() {
    var v1 uint8 = 43
    var v2 uint = 1
    var v3 uint = 133
    var v4 string = (("foo bar" + ("hello" + ("hello" + "foo bar"))) + ("hello" + ("a" + ("\n\t" + ("foo bar" + "hello")))))
    v5 := []string{v4}
    v6 := []uint32{2147483648}
    v3 = (^(v3))
    var v7 int32 = 1
    for fi8 := 0; fi8 < 14; fi8++ {
        if true {
            v2 = v3
        } else {
            for fi9 := 0; fi9 < 16; fi9++ {
                v1 -= (^(v1))
                if false {
                    switch (fi8 * 68) {
                    case 1:
                        var v10 string = v4
                        _ = v10
                    }
                } else {
                    switch ("hello" + v4) {
                    case "\n\t":
                        if (!(1 < -1)) {
                            var v11 int64 = -9223372036854775808
                            _ = v11
                        }
                        if false {
                            v3 -= (v2 | (v2 | 42))
                            v7 -= v7
                            if (true || (0 < 1)) {
                                var v12 bool = ((42 > 0) || true)
                                v7 = 42
                                if (1 <= v7) {
                                    var v13 int16 = 42
                                    if v12 {
                                        _ = len(v6)
                                        v14 := GenS1{F0: 32767}
                                        if false {
                                            var v15 uint = (v2 * v2)
                                            _ = v15
                                        } else {
                                            if (!(v12 && (2147483647 == 711))) {
                                                if (!v12) {
                                                    v1 = (^(v1))
                                                } else {
                                                    fmt.Println(v13)
                                                    var v16 int64 = 42
                                                    _ = v16
                                                }
                                                if (!((^(v1)) > v1)) {
                                                    switch ((v4 + "hello") + v4) {
                                                    case "a":
                                                        v2 += v3
                                                    case "\n\t":
                                                        v17 := GenS2{F0: v7, F1: -128, F2: 60}
                                                        _ = v17
                                                    case "hello":
                                                        if v12 {
                                                            var v18 uint8 = 152
                                                            _ = v18
                                                        }
                                                    default:
                                                        v19 := GenS3{F0: 1}
                                                        _ = v19
                                                    }
                                                }
                                                if len(v6) > 0 {
                                                    v6[0] = 2147483647
                                                }
                                            } else {
                                                v13 += 1
                                            }
                                        }
                                        _ = v14
                                    } else {
                                        if (v7 == (v7 >> (((v3 / ((v3) | 1))) & 63))) {
                                            var v20 float64 = 1.5
                                            if v12 {
                                                var v21 bool = (v4 < v4)
                                                _ = v21
                                            } else {
                                                fmt.Println(fi9)
                                                var v22 float32 = 1.5
                                                v23 := GenS3{F0: 18446744073709551615}
                                                _ = v22
                                                _ = v23
                                            }
                                            fmt.Println(v13, v7, v20)
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            v24 := GenS1{F0: 42}
                                            var v25 int8 = -128
                                            v7 -= v7
                                            v26 := GenS3{F0: 317}
                                            _ = v20
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        }
                                        var v27 uint = v2
                                        if len(v6) > 0 {
                                            v6[0] = 377
                                        }
                                        _ = v27
                                    }
                                    fmt.Println(v2)
                                    _ = v13
                                }
                                _ = v12
                            } else {
                                v28 := GenS4{F0: 255, F1: 65535, F2: 702}
                                if (18446744073709551615 <= 340) {
                                    if (1.5 >= float32(1.5)) {
                                        v29 := GenS1{F0: 255}
                                        v4 = v4
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v29
                                    }
                                } else {
                                    v7 = v7
                                    var v30 int = (-(fi9))
                                    if len(v5) > 0 {
                                        v5[0] = ""
                                    }
                                    _ = v30
                                }
                                switch "hello" {
                                case "":
                                    v31 := GenS5{F0: v7}
                                    v1 += 42
                                    _ = v31
                                }
                                _ = v28
                            }
                            switch v4 {
                            case "hello":
                                var v32 uint64 = 277
                                _ = v32
                            case "\n\t":
                                v7 -= v7
                            }
                        } else {
                            var v33 int16 = 32767
                            var v34 int16 = 767
                            _ = v33
                            _ = v34
                        }
                        v1 += v1
                    case "hello":
                        var v35 uint64 = 42
                        _ = v35
                    default:
                        fmt.Println(fi8)
                    }
                    var v36 int64 = 42
                    _ = v36
                }
                var v37 int64 = 690
                _ = fi9
                _ = v37
            }
        }
        var v38 float64 = 1.5
        _ = fi8
        _ = v38
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        if false {
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
        } else {
            var v40 string = v4
            var v41 uint8 = v1
            _ = v40
            _ = v41
        }
        switch (fi39 * (fi39 & (fi39 - 869))) {
        case 128:
            v1 -= (v1 << ((v2) & 63))
        case 0:
            v4 = "foo bar"
        default:
            v42 := GenS6{F0: (v7 >= (v7 | (v7 * -421))), F1: int8(v2), F2: (v4 >= v4)}
            _ = v42
        }
        _ = fi39
    }
    for fi43 := 0; fi43 < 17; fi43++ {
        fmt.Println(v2)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v7 -= (v7 + (^(v7)))
        _ = fi43
    }
    var v44 float64 = (3.14159 + (0.1 - (((-2.25 - (1.5 * 0.1)) / 1e10) - (3.14159 * (0.0 / 3.14159)))))
    for fi45 := 0; fi45 < 15; fi45++ {
        for fi46 := 0; fi46 < 8; fi46++ {
            if (!(true || (((v1 * (v1 << ((1) & 63))) < 255) || (1 >= 1)))) {
                if (false || (!((v3 > v2) || false))) {
                    v6 = append(v6, 822)
                    var v47 int64 = 1
                    _ = v47
                } else {
                    switch (fi45 * 42) {
                    case 0:
                        fmt.Println(v1, v4, fi45)
                    default:
                        v2 += (v2 - v2)
                    }
                }
                if (-2.25 < ((float32(v44) - (float32((float64(1e10) - v44)) + float32((v44 + -2.25)))) / (1e10 * 1.5))) {
                    v48 := GenS6{F0: false, F1: 0, F2: (42 >= v1)}
                    v3 += (v3 & 567)
                    if ((fi46 | -806) > (fi45 + (^(fi46)))) {
                        if false {
                            v44 = v44
                        }
                        fmt.Println(fi45, fi46)
                    }
                    _ = v48
                } else {
                    v49 := GenS5{F0: v7}
                    var v50 int32 = (v7 & v7)
                    var v51 string = v4
                    _ = v49
                    _ = v50
                    _ = v51
                }
                v2 += (v2 << ((0) & 63))
            }
            v52 := GenS7{F0: 2}
            v44 -= float64(0.0)
            _ = fi46
            _ = v52
        }
        for fi53 := 0; fi53 < 15; fi53++ {
            v3 -= (v2 & v3)
            var v54 int = fi45
            switch (v54 & 755) {
            case 32768:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 9223372036854775807:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            default:
                if (!true) {
                    var v55 int16 = 78
                    fmt.Println(v4, v3)
                    _ = v55
                }
            }
            _ = fi53
            _ = v54
        }
        fmt.Println(v3, v44, v4)
        _ = fi45
    }
    for fi56 := 0; fi56 < 19; fi56++ {
        var v57 string = v4
        _ = fi56
        _ = v57
    }
    var v58 bool = (true || false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v58
    _ = v6
    _ = v7
}
