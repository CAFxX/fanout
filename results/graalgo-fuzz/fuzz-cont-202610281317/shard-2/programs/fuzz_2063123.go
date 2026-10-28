package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 string; F2 int16 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 uint64 }

func main() {
    var v1 float64 = 0.1
    var v2 bool = false
    var v3 uint64 = 4294967295
    var v4 uint32 = 65535
    var v5 int64 = 0
    var v6 int8 = 81
    fmt.Println(v5)
    var v7 uint8 = 128
    v3 += 649
    v8 := GenS1{F0: 255, F1: 32767, F2: 810}
    var v9 int8 = v6
    if (v1 < v1) {
        if v2 {
            v10 := GenS2{F0: -2147483648, F1: 2147483647}
            if v2 {
                v11 := GenS3{F0: 2147483647, F1: v4}
                if true {
                    if true {
                        v12 := []int{0, 1, -139, -679}
                        _ = v12
                    }
                    v13 := []int{-174, 128, -1}
                    var v14 uint64 = (v3 | 2)
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(v6, v1, v5)
                    var v15 uint8 = (v7 * (v7 + v7))
                    var v16 int8 = (-(v9))
                    _ = v15
                    _ = v16
                }
                _ = v11
            } else {
                v17 := GenS4{F0: -1}
                v1 = v1
                _ = v17
            }
            _ = v10
        } else {
            if false {
                switch -375 {
                case 42:
                    if (18446744073709551615 < (v3 - uint64(v7))) {
                        v18 := GenS5{F0: (v3 & 42), F1: ("\n\t" + ("hello" + "\n\t")), F2: 0}
                        v5 -= (v5 % (((v5 & -743)) | 1))
                        _ = v18
                    }
                case 4294967295:
                    switch "foo bar" {
                    case "foo bar":
                        v9 -= v6
                    }
                }
                var v19 string = "hello"
                for fi20 := 0; fi20 < 14; fi20++ {
                    switch (("hello" + "a") + (("a" + "foo bar") + v19)) {
                    case "\n\t":
                        var v21 uint64 = v3
                        _ = v21
                    case "a":
                        v22 := []int{(fi20 >> ((535) & 63)), fi20, (fi20 - 519)}
                        _ = v22
                    case "hello":
                        for fi23 := 0; fi23 < 17; fi23++ {
                            if (100 > 20) {
                                v24 := []int{(fi23 + fi23), (fi23 % ((-718) | 1)), (fi23 % (((fi20 & fi23)) | 1)), fi20}
                                if false {
                                    fmt.Println(v2, v7)
                                    var v25 uint8 = (^(v7))
                                    v26 := GenS1{F0: 0, F1: -32768, F2: (^(fi20))}
                                    _ = v25
                                    _ = v26
                                } else {
                                    var v27 float64 = float64(((1e10 * 1e10) / ((-2.25 / 1.5) * 3.14159)))
                                    _ = v27
                                }
                                v9 = int8(v5)
                                if true {
                                    if false {
                                        fmt.Println(v7)
                                    } else {
                                        v3 -= (v3 ^ (v3 ^ v3))
                                        if v2 {
                                            var v28 int32 = 42
                                            _ = v28
                                        }
                                    }
                                }
                                var v29 bool = ((fi23 % (((fi20 & (fi23 | (^(fi23))))) | 1)) <= 9223372036854775807)
                                if ((v5 * v5) >= 9223372036854775807) {
                                    if false {
                                        fmt.Println(v7)
                                        switch ((v19 + "\n\t") + v19) {
                                        case "a":
                                            var v30 int64 = -970
                                            _ = v30
                                        }
                                        fmt.Println(v4)
                                    } else {
                                        if false {
                                            var v31 uint8 = 63
                                            v5 = (v5 % ((v5) | 1))
                                            var v32 uint = 255
                                            _ = v31
                                            _ = v32
                                        }
                                        v6 = v6
                                    }
                                }
                                fmt.Println(v6)
                                _ = v24
                                _ = v29
                            }
                            _ = fi23
                        }
                    default:
                        v33 := []int{(fi20 % (((fi20 % (((fi20 + 127)) | 1))) | 1)), fi20}
                        _ = v33
                    }
                    var v34 int = (fi20 / (((fi20 / ((9223372036854775807) | 1))) | 1))
                    _ = fi20
                    _ = v34
                }
                _ = v19
            } else {
                for fi35 := 0; fi35 < 6; fi35++ {
                    var v36 uint8 = v7
                    var v37 uint = 606
                    switch (fi35 >> (((^(v37))) & 63)) {
                    case -671:
                        for fi38 := 0; fi38 < 1; fi38++ {
                            v39 := []int{fi38, fi35, 9223372036854775807, fi38}
                            v6 = v9
                            fmt.Println(v9)
                            _ = fi38
                            _ = v39
                        }
                    }
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                switch int(v5) {
                case 0:
                    if false {
                        v40 := GenS5{F0: (v3 << ((42) & 63)), F1: "\n\t", F2: 1000}
                        v41 := []int{}
                        _ = v40
                        _ = v41
                    } else {
                        switch "a" {
                        case "foo bar":
                            for fi42 := 0; fi42 < 3; fi42++ {
                                v43 := GenS6{F0: v4, F1: (("hello" + "hello") + "\n\t")}
                                v44 := GenS7{F0: (v3 | (v3 * 116))}
                                v3 += (v3 >> ((4294967295) & 63))
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                        case "hello":
                            v2 = (!((-371 != -32768) || v2))
                        case "a":
                            v45 := []int{-1}
                            _ = v45
                        default:
                            fmt.Println(v3, v7)
                        }
                        for fi46 := 0; fi46 < 13; fi46++ {
                            for fi47 := 0; fi47 < 4; fi47++ {
                                v48 := GenS1{F0: 4294967295, F1: 42, F2: 711}
                                _ = fi47
                                _ = v48
                            }
                            var v49 uint = uint(v7)
                            _ = fi46
                            _ = v49
                        }
                        if (float32((float64(0.1) / v1)) >= (0.0 + ((0.0 - float32((v1 / 1.0))) - 1.5))) {
                            var v50 float32 = 1.5
                            v51 := []int{1, 42, 9223372036854775807, -9223372036854775808}
                            _ = v50
                            _ = v51
                        } else {
                            fmt.Println(v9)
                        }
                    }
                default:
                    if v2 {
                        var v52 uint32 = (^(v4))
                        var v53 int8 = -1
                        var v54 uint32 = v52
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                }
            }
        }
        var v55 uint8 = (v7 * 255)
        _ = v55
    } else {
        fmt.Println(v6, v5)
        for fi56 := 0; fi56 < 16; fi56++ {
            var v57 uint8 = (v7 % (((^(v7))) | 1))
            _ = fi56
            _ = v57
        }
        v58 := []int{-733, 32767, -1}
        _ = v58
    }
    var v59 int16 = -309
    v60 := GenS2{F0: int32(v59), F1: -824}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v7
    _ = v8
    _ = v9
}
