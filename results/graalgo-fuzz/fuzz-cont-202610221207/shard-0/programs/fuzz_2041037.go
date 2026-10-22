package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }

func main() {
    var v1 int32 = 42
    var v2 uint = 1
    var v3 int8 = -128
    var v4 uint64 = 42
    v5 := make(map[int]bool)
    if ((("" + "\n\t") != "hello") || true) {
        switch 128 {
        case 0:
            v1 += v1
        }
        v6 := GenS1{F0: 4294967295}
        _ = v6
    } else {
        switch 127 {
        case 4294967295:
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := GenS2{F0: 231, F1: 128, F2: 2147483647}
                var v9 int16 = 32767
                switch (fi7 << (((v2 << (((v2 - (^(v2)))) & 63))) & 63)) {
                case 42:
                    v10 := GenS3{F0: v4}
                    _ = v10
                case -48:
                    var v11 uint32 = 2147483648
                    _ = v11
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        fmt.Println(v4, v1)
    }
    v4 = v4
    if (!(("hello" + "a") < ("hello" + "foo bar"))) {
        var v12 float64 = 1.5
        _ = v12
    } else {
        var v13 uint32 = 2147483648
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                fmt.Println(fi15)
                if true {
                    v16 := v5[128]
                    if v16 {
                        v4 = (^(v4))
                        v1 += v1
                    } else {
                        fmt.Println(v2)
                        switch "" {
                        case "hello":
                            delete(v5, (fi14 * 4294967295))
                        }
                        if v16 {
                            var v17 int64 = 9223372036854775807
                            _ = v17
                        }
                        switch fi15 {
                        case 32768:
                            if false {
                                if (!(!v16)) {
                                    v13 += 4294967295
                                }
                            } else {
                                _ = len(v5)
                                var v18 bool = v16
                                v19 := []int{fi14, fi15, fi15, int(v4)}
                                _ = v18
                                _ = v19
                            }
                        case 128:
                            if (!v16) {
                                var v20 int8 = v3
                                v21 := GenS3{F0: v4}
                                v22 := GenS4{F0: (fi14 & fi14)}
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        default:
                            if (v16 && (!(uint8(v2) > 111))) {
                                fmt.Println(v3, v4, v4)
                                v23 := GenS5{F0: "hello"}
                                _ = v23
                            }
                            fmt.Println(v4)
                            v13 -= (^(v13))
                            if (v4 > v4) {
                                var v24 int16 = 328
                                _ = v24
                            }
                            if false {
                                fmt.Println(v1, fi15)
                                var v25 uint = v2
                                var v26 string = "hello"
                                switch 9223372036854775807 {
                                case -1:
                                    v27 := GenS3{F0: 4294967295}
                                    _ = v27
                                case -67:
                                    fmt.Println(v2, fi15)
                                case -301:
                                    var v28 int8 = (v3 + v3)
                                    _ = v28
                                default:
                                    switch fi14 {
                                    case 0:
                                        var v29 int16 = -704
                                        _ = v29
                                    case -1:
                                        v30 := []int{fi14, 9223372036854775807}
                                        _ = v30
                                    case -125:
                                        var v31 uint8 = 1
                                        _ = v31
                                    }
                                }
                                _ = v25
                                _ = v26
                            }
                        }
                    }
                    if false {
                        var v32 uint64 = v4
                        var v33 int = -2147483648
                        var v34 int32 = v1
                        if v16 {
                            var v35 float64 = float64((3.14159 + 3.14159))
                            v36 := v5[522]
                            var v37 int64 = 42
                            _ = v35
                            _ = v36
                            _ = v37
                        } else {
                            v34 += (v34 << (((^(v2))) & 63))
                        }
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    v38 := v5[fi15]
                    var v39 string = "hello"
                    v4 -= v4
                    _ = v16
                    _ = v38
                    _ = v39
                }
                _ = fi15
            }
            _ = fi14
        }
        _ = v13
    }
    v4 -= (^(v4))
    switch -275 {
    case 42:
        v40 := []int{535, 0, -9223372036854775808, 32768}
        _ = v40
    default:
        v1 = v1
    }
    delete(v5, 42)
    fmt.Println(v3)
    if ("a" <= (("a" + "") + ("a" + "hello"))) {
        var v41 uint8 = 42
        v4 = (v4 * v4)
        switch "" {
        case "a":
            for fi42 := 0; fi42 < 16; fi42++ {
                switch fi42 {
                case -1:
                    fmt.Println(v3)
                }
                _ = fi42
            }
        case "":
            fmt.Println(v1)
        case "hello":
            for fi43 := 0; fi43 < 8; fi43++ {
                for fi44 := 0; fi44 < 9; fi44++ {
                    v45 := []int{fi43, (fi43 % ((fi44) | 1)), int(v1)}
                    _ = fi44
                    _ = v45
                }
                fmt.Println(fi43, v41)
                _ = fi43
            }
        }
        _ = v41
    } else {
        if (0.1 != 1e10) {
            v46 := GenS3{F0: v4}
            _ = v46
        }
    }
    switch -9223372036854775808 {
    case 4294967295:
        delete(v5, 32768)
    }
    if false {
        switch "" {
        case "hello":
            var v47 int32 = 42
            _ = v47
        case "a":
            v48 := GenS4{F0: 32767}
            _ = v48
        case "foo bar":
            v4 += (v4 & v4)
        default:
            fmt.Println(v4, v3, v2)
        }
    }
    v3 = v3
    var v49 int8 = (v3 & (v3 * v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
}
