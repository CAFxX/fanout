package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint8 }

func main() {
    v1 := []int{}
    var v2 int32 = 42
    var v3 int32 = 42
    var v4 string = "hello"
    var v5 int64 = 1
    v5 = (-(v5))
    _ = len(v1)
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint8 = 128
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v2, fi8)
        var v9 float32 = 0.1
        _ = fi8
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        v4 = ("hello" + "\n\t")
    case "":
        var v10 uint64 = 2
        _ = v10
    case "hello":
        v1 = append(v1, 0)
    }
    if true {
        switch "hello" {
        case "":
            for fi11 := 0; fi11 < 12; fi11++ {
                if len(v1) > 0 {
                    v1[0] = (fi11 ^ 4294967295)
                }
                var v12 bool = false
                _ = fi11
                _ = v12
            }
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v5 = v5
        v4 = v4
    } else {
        v5 -= -1
        switch ("foo bar" + v4) {
        case "":
            v13 := GenS1{F0: 452, F1: -356}
            _ = v13
        default:
            var v14 bool = true
            _ = v14
        }
        fmt.Println(v5)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 = v5
    var v15 uint = 2
    if ((!false) && false) {
        var v16 bool = true
        for fi17 := 0; fi17 < 11; fi17++ {
            if len(v1) > 0 {
                v1[0] = fi17
            }
            v15 -= v15
            _ = fi17
        }
        if (v16 || v16) {
            fmt.Println(v3)
        } else {
            var v18 bool = (552 > 42)
            _ = v18
        }
        _ = v16
    } else {
        _ = len(v1)
        for fi19 := 0; fi19 < 18; fi19++ {
            v5 -= v5
            switch (v4 + ("a" + v4)) {
            case "\n\t":
                if (fi19 < (fi19 ^ -119)) {
                    if (1 > 65535) {
                        var v20 int8 = -1
                        _ = v20
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if true {
                        if (32767 != 32767) {
                            if (true && (42 < 0)) {
                                if false {
                                    v4 = ""
                                    for fi21 := 0; fi21 < 2; fi21++ {
                                        if (2 > 42) {
                                            v22 := GenS1{F0: -32768, F1: fi19}
                                            v23 := GenS1{F0: 0, F1: (fi19 * 0)}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v1 = append(v1, fi21)
                                            _ = v22
                                            _ = v23
                                        }
                                        _ = fi21
                                    }
                                    fmt.Println(v3, v15)
                                }
                                v24 := GenS2{F0: 4294967295, F1: v5, F2: 2}
                                if false {
                                    var v25 float32 = (3.14159 - -2.25)
                                    fmt.Println(fi19)
                                    v15 += (v15 & (v15 >> ((v15) & 63)))
                                    _ = v25
                                }
                                _ = v24
                            }
                            var v26 uint = (v15 / (((v15 << ((2) & 63))) | 1))
                            for fi27 := 0; fi27 < 6; fi27++ {
                                var v28 int = (-(fi27))
                                _ = fi27
                                _ = v28
                            }
                            _ = v26
                        } else {
                            v3 = v3
                            for fi29 := 0; fi29 < 5; fi29++ {
                                if false {
                                    v1 = append(v1, 1)
                                }
                                _ = fi29
                            }
                            var v30 int32 = -1
                            _ = v30
                        }
                        v5 += v5
                        v31 := GenS1{F0: 539, F1: fi19}
                        _ = v31
                    }
                    fmt.Println(v5, v5)
                }
            case "foo bar":
                if true {
                    v1 = append(v1, fi19)
                    for fi32 := 0; fi32 < 8; fi32++ {
                        v3 += v2
                        _ = fi32
                    }
                    var v33 int8 = -20
                    _ = v33
                } else {
                    if (56 != 137) {
                        if true {
                            v15 += 255
                            if len(v1) > 0 {
                                v1[0] = fi19
                            }
                        } else {
                            v34 := GenS3{F0: -256, F1: 127}
                            fmt.Println(v3)
                            var v35 uint64 = 326
                            _ = v34
                            _ = v35
                        }
                    }
                    var v36 float64 = (((-2.25 * (0.1 + (1.5 * float64(float32((3.14159 + 1e10)))))) + 0.1) + (0.1 * ((float64(1.5) - (-2.25 * 0.1)) - 0.1)))
                    v4 = ((v4 + "a") + (v4 + ((v4 + v4) + v4)))
                    _ = v36
                }
            case "hello":
                if false {
                    switch "\n\t" {
                    case "a":
                        var v37 float32 = (float32(0.1) - 0.0)
                        _ = v37
                    case "":
                        if false {
                            if (!true) {
                                if len(v1) > 0 {
                                    v1[0] = fi19
                                }
                            } else {
                                switch (fi19 % (((fi19 + fi19)) | 1)) {
                                case 32768:
                                    v3 = 42
                                case 335:
                                    var v38 uint8 = 2
                                    _ = v38
                                case -2147483648:
                                    v39 := GenS2{F0: 850, F1: v5, F2: 536}
                                    _ = v39
                                default:
                                    switch -9223372036854775808 {
                                    case 127:
                                        fmt.Println(v4, v4)
                                    case 904:
                                        if (true || false) {
                                            v2 += v2
                                            if true {
                                                if true {
                                                    _ = len(v1)
                                                    var v40 string = ("\n\t" + v4)
                                                    _ = v40
                                                } else {
                                                    var v41 int64 = int64(v3)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v41 = (^(v41))
                                                    _ = v41
                                                }
                                                for fi42 := 0; fi42 < 20; fi42++ {
                                                    fmt.Println(fi42, fi42, v15)
                                                    _ = fi42
                                                }
                                            }
                                        } else {
                                            v43 := GenS2{F0: 4294967295, F1: v5, F2: 42}
                                            _ = v43
                                        }
                                    case -9223372036854775808:
                                        if ((1 != 32767) && true) {
                                            v3 = v3
                                        } else {
                                            v2 = v2
                                        }
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = (fi19 + (fi19 - fi19))
                                }
                                _ = len(v1)
                            }
                        } else {
                            v44 := GenS2{F0: 2147483647, F1: 1, F2: 100}
                            _ = v44
                        }
                    }
                    for fi45 := 0; fi45 < 4; fi45++ {
                        if (false || (!true)) {
                            if false {
                                v2 = v2
                            }
                            v2 = (v3 + v2)
                        }
                        var v46 string = v4
                        switch (-(fi45)) {
                        case 127:
                            var v47 uint64 = 2
                            _ = v47
                        }
                        _ = fi45
                        _ = v46
                    }
                }
            default:
                for fi48 := 0; fi48 < 8; fi48++ {
                    var v49 uint8 = 200
                    _ = fi48
                    _ = v49
                }
            }
            var v50 int32 = int32(fi19)
            _ = fi19
            _ = v50
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
