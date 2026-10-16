package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int64; F1 string; F2 bool }
type GenS5 struct { F0 uint8; F1 string; F2 bool }
type GenS6 struct { F0 int32 }

func main() {
    var v1 bool = ((1.5 - (float64(3.14159) / 0.1)) != 1.5)
    var v2 uint32 = 42
    var v3 uint = 1
    v4 := []string{("" + "\n\t"), "hello", ""}
    var v5 string = (("" + ("foo bar" + "foo bar")) + "\n\t")
    v6 := make(map[string]uint32)
    switch v5 {
    case "\n\t":
        switch 0 {
        case 4294967295:
            var v7 uint16 = 62
            _ = v7
        case 32768:
            fmt.Println(v2)
        case -539:
            for fi8 := 0; fi8 < 20; fi8++ {
                v2 -= v2
                _ = fi8
            }
        default:
            var v9 uint = (v3 * v3)
            _ = v9
        }
    case "foo bar":
        if (true && false) {
            var v10 uint64 = 690
            var v11 int8 = 31
            _ = v10
            _ = v11
        }
    default:
        if (v1 && v1) {
            v12 := GenS1{F0: -32768}
            if v1 {
                for fi13 := 0; fi13 < 16; fi13++ {
                    fmt.Println(v1)
                    v3 = v3
                    _ = fi13
                }
                var v14 uint64 = 293
                fmt.Println(v5, v2, v5)
                _ = v14
            } else {
                for fi15 := 0; fi15 < 4; fi15++ {
                    v1 = (1 > 2)
                    v16 := GenS1{F0: 1000}
                    _ = fi15
                    _ = v16
                }
                v17 := GenS2{F0: 18446744073709551615, F1: 4294967295}
                var v18 int16 = 0
                _ = v17
                _ = v18
            }
            _ = v12
        }
    }
    if (v1 && v1) {
        for fi19 := 0; fi19 < 19; fi19++ {
            switch (v5 + v5) {
            case "foo bar":
                if (!(!((!(("foo bar" != "\n\t") || (((("a" + v5) + v5) + v5) == v5))) && v1))) {
                    var v20 uint32 = (^(v2))
                    v5 = ("hello" + (v5 + "hello"))
                    _ = v20
                } else {
                    if v1 {
                        if v1 {
                            v5 = v5
                        } else {
                            var v21 uint32 = 658
                            v22 := GenS3{F0: v3, F1: 42}
                            _ = v21
                            _ = v22
                        }
                        v5 = v5
                    }
                    if (((0.0 - (1.5 + (-2.25 + (3.14159 + (((float64(-2.25) * 0.1) / -2.25) * ((0.0 - ((0.1 * 3.14159) + 0.1)) + 1e10)))))) + -2.25) >= 3.14159) {
                        fmt.Println(v3)
                        if v1 {
                            v5 = "hello"
                            v3 += v3
                            v3 = v3
                        }
                    } else {
                        switch (v5 + "hello") {
                        case "hello":
                            v6["hello"] = (v2 << ((v3) & 63))
                        default:
                            var v23 int = 32768
                            _ = v23
                        }
                        if v1 {
                            var v24 int16 = -202
                            v6["\n\t"] = v2
                            _ = v24
                        } else {
                            if (((v3 / ((2) | 1)) >= v3) || true) {
                                fmt.Println(v5, fi19)
                                v1 = v1
                                v25 := GenS3{F0: (v3 & v3), F1: 110}
                                _ = v25
                            }
                            if len(v4) > 0 {
                                v4[0] = (v5 + (("foo bar" + v5) + (("hello" + v5) + v5)))
                            }
                            v26 := GenS4{F0: -9223372036854775808, F1: v5, F2: false}
                            _ = v26
                        }
                    }
                    if false {
                        v2 -= 849
                        var v27 uint8 = 128
                        _ = v27
                    } else {
                        var v28 uint16 = 65535
                        v3 = 2
                        _ = v28
                    }
                }
            case "\n\t":
                v3 = v3
            case "hello":
                switch fi19 {
                case 422:
                    fmt.Println(v3)
                default:
                    v2 += (^(v2))
                }
            }
            if v1 {
                if true {
                    v2 = (v2 >> ((v3) & 63))
                    var v29 int = (fi19 ^ 351)
                    v5 = (("a" + ("" + v5)) + ("\n\t" + v5))
                    _ = v29
                }
                if true {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        fmt.Println(fi30, v1)
                        fmt.Println(v2)
                        _ = fi30
                    }
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v5, v3)
                    case "hello":
                        v2 = (v2 >> ((v3) & 63))
                    case "a":
                        var v31 float64 = 1e10
                        _ = v31
                    default:
                        if v1 {
                            var v32 uint = (v3 ^ v3)
                            var v33 uint8 = 9
                            _ = len(v6)
                            _ = v32
                            _ = v33
                        }
                    }
                }
                if (!(!(0.1 <= (1e10 + 0.0)))) {
                    switch (fi19 | (fi19 + (fi19 + (fi19 - 172)))) {
                    case -2147483648:
                        v3 = v3
                    default:
                        v3 += (v3 & (v3 * (v3 >> ((v3) & 63))))
                    }
                    fmt.Println(v5)
                    var v34 string = v5
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 18; fi35++ {
                        v5 = "foo bar"
                        v5 = "a"
                        var v36 uint32 = (v2 & (v2 - v2))
                        _ = fi35
                        _ = v36
                    }
                }
            }
            _ = fi19
        }
        var v37 uint16 = 42
        v2 += (v2 / (((v2 << (((v3 << ((v3) & 63))) & 63))) | 1))
        _ = v37
    } else {
        v3 -= v3
        v38 := GenS5{F0: 32, F1: (v5 + (v5 + v5)), F2: (!v1)}
        _ = v38
    }
    v2 += 463
    fmt.Println(v5, v5, v3)
    v3 = (^(v3))
    if ((-1 > -1) && ("foo bar" == v5)) {
        v3 -= (v3 / (((v3 << ((2) & 63))) | 1))
        _ = len(v6)
        v1 = (true || (4294967295 <= 994))
    } else {
        var v39 bool = true
        _ = v39
    }
    v1 = v1
    fmt.Println(v2)
    if (v1 && ((("foo bar" + v5) + v5) <= v5)) {
        var v40 string = "\n\t"
        v1 = v1
        _ = v40
    }
    var v41 int8 = -128
    var v42 uint64 = 1
    for fi43 := 0; fi43 < 20; fi43++ {
        if (!v1) {
            for fi44 := 0; fi44 < 18; fi44++ {
                v3 += 598
                var v45 uint32 = v2
                v41 = v41
                v42 = v42
                var v46 int16 = 32767
                _ = fi44
                _ = v45
                _ = v46
            }
        } else {
            v47 := GenS6{F0: -340}
            fmt.Println(v42)
            _ = v47
        }
        v1 = v1
        _ = fi43
    }
    v3 = (v3 << (((^(v3))) & 63))
    var v48 uint = (v3 * v3)
    v49 := v6[(v5 + v5)]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
