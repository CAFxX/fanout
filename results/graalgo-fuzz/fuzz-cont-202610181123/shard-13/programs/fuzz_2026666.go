package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }
type GenS3 struct { F0 int32; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 string; F1 int8 }
func genFn0(p0 int64, p1 bool, p2 uint) int16 {
    for fi5 := 0; fi5 < 14; fi5++ {
        switch ("foo bar" + "foo bar") {
        case "a":
            var v6 uint32 = 0
            _ = v6
        case "foo bar":
            for fi7 := 0; fi7 < 15; fi7++ {
                p2 = 278
                v8 := []int{(fi7 + (fi5 * (fi7 << ((p2) & 63))))}
                _ = fi7
                _ = v8
            }
        case "hello":
            v9 := GenS1{F0: 1, F1: 1}
            _ = v9
        default:
            fmt.Println(fi5, p2, p1)
        }
        _ = fi5
    }
    switch "" {
    case "hello":
        var v10 uint = p2
        _ = v10
    case "":
        v11 := []int{42, 128}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int = 32768
    var v3 uint16 = 42
    var v4 uint8 = 0
    if true {
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(v4, v1)
            if false {
                v2 = 1
                switch (fi12 / ((fi12) | 1)) {
                case 423:
                    switch (fi12 ^ v2) {
                    case -2147483648:
                        v2 = (v2 * fi12)
                    default:
                        var v13 float64 = 1e10
                        _ = v13
                    }
                case 32767:
                    var v14 uint = 65535
                    _ = v14
                case -2147483648:
                    var v15 int16 = 1000
                    _ = v15
                }
            } else {
                v16 := []int{fi12}
                _ = v16
            }
            var v17 int64 = -9223372036854775808
            _ = fi12
            _ = v17
        }
    } else {
        v4 = v4
        var v18 bool = (true || false)
        if v18 {
            switch ((("hello" + "\n\t") + "a") + "") {
            case "":
                if true {
                    v19 := GenS2{F0: -759, F1: -1, F2: -655}
                    v3 -= v3
                    var v20 uint64 = 997
                    _ = v19
                    _ = v20
                } else {
                    if v18 {
                        var v21 int64 = -688
                        if (1e10 != 3.14159) {
                            switch (^(v2)) {
                            case 42:
                                var v22 uint32 = 33
                                _ = v22
                            case 401:
                                for fi23 := 0; fi23 < 1; fi23++ {
                                    v24 := GenS3{F0: -2147483648, F1: v3, F2: 42}
                                    v18 = v18
                                    _ = fi23
                                    _ = v24
                                }
                            default:
                                v25 := GenS1{F0: 2147483647, F1: (v1 & (v1 % ((0) | 1)))}
                                _ = v25
                            }
                            v26 := []int{(v2 % ((643) | 1)), (^(v2)), (v2 * (v2 / ((127) | 1))), 127}
                            _ = v26
                        } else {
                            var v27 uint8 = v4
                            if v18 {
                                v4 -= 42
                                var v28 uint8 = v27
                                _ = v28
                            } else {
                                if (0 >= 42) {
                                    v3 = v3
                                } else {
                                    if (v18 || (2147483648 > 42)) {
                                        if ((v2 | 128) == (v2 ^ (v2 | (v2 << ((517) & 63))))) {
                                            v4 = v4
                                            switch ("\n\t" + "foo bar") {
                                            case "foo bar":
                                                v3 = v3
                                            }
                                        } else {
                                            v4 = v4
                                            v29 := GenS2{F0: 0, F1: v21, F2: -32768}
                                            if v18 {
                                                var v30 string = "a"
                                                if ((^(v3)) <= (v3 >> ((uint(v3)) & 63))) {
                                                    var v31 uint64 = (v1 ^ (v1 / (((v1 ^ v1)) | 1)))
                                                    _ = v31
                                                } else {
                                                    fmt.Println(v1, v27)
                                                    v21 += v21
                                                }
                                                v32 := []int{571, v2, -450, v2}
                                                _ = v30
                                                _ = v32
                                            }
                                            _ = v29
                                        }
                                    }
                                }
                            }
                            v18 = v18
                            _ = v27
                        }
                        var v33 int = v2
                        _ = v21
                        _ = v33
                    } else {
                        var v34 string = ("a" + "")
                        if (v34 == v34) {
                            v35 := []int{}
                            _ = v35
                        } else {
                            for fi36 := 0; fi36 < 1; fi36++ {
                                var v37 uint16 = 217
                                _ = fi36
                                _ = v37
                            }
                            fmt.Println(v3, v3)
                        }
                        _ = v34
                    }
                    v38 := []int{}
                    v39 := GenS4{F0: (v1 * 0), F1: 42}
                    _ = v38
                    _ = v39
                }
            case "foo bar":
                var v40 uint = 405
                _ = v40
            case "a":
                v4 += v4
            }
        } else {
            v3 -= v3
            v3 = (v3 + (v3 + (v3 >> ((65535) & 63))))
        }
        _ = v18
    }
    v41 := GenS3{F0: -1, F1: (^(v3)), F2: 1}
    var v42 string = ""
    for fi43 := 0; fi43 < 17; fi43++ {
        fmt.Println(v3, v2)
        for fi44 := 0; fi44 < 9; fi44++ {
            v45 := GenS5{F0: (v42 + "\n\t"), F1: 42}
            if (42 > 0) {
                v46 := GenS2{F0: -1, F1: 885, F2: 32767}
                var v47 uint32 = 2147483647
                var v48 uint32 = v47
                _ = v46
                _ = v47
                _ = v48
            } else {
                fmt.Println(v2, v3, v2)
                var v49 uint = 493
                _ = v49
            }
            v50 := GenS1{F0: 2147483647, F1: v1}
            _ = fi44
            _ = v45
            _ = v50
        }
        v4 -= v4
        _ = fi43
    }
    switch (v2 % ((127) | 1)) {
    case -871:
        fmt.Println(v4)
    case 676:
        fmt.Println(v1)
    case 0:
        v51 := []int{}
        _ = v51
    }
    v4 = v4
    var v52 int16 = 42
    v53 := GenS2{F0: 42, F1: -236, F2: v52}
    if false {
        var v54 int32 = 1
        var v55 uint = 255
        _ = v54
        _ = v55
    } else {
        v42 = "\n\t"
        v42 = v42
    }
    v56 := []int{v2, int(v1), v2}
    var v57 int64 = 1
    if (true && (false && false)) {
        v2 = (^(v2))
    } else {
        v1 = (v1 * (v1 << ((255) & 63)))
        switch v42 {
        case "foo bar":
            for fi58 := 0; fi58 < 11; fi58++ {
                var v59 uint64 = uint64(v3)
                fmt.Println(fi58, v59)
                var v60 int32 = 1
                _ = fi58
                _ = v59
                _ = v60
            }
        }
    }
    v1 = 1
    v61 := GenS4{F0: 290, F1: v57}
    switch "" {
    case "foo bar":
        if false {
            for fi62 := 0; fi62 < 14; fi62++ {
                if (4294967295 < v1) {
                    if (0.0 == -2.25) {
                        if false {
                            v3 -= (v3 | v3)
                            v42 = (v42 + (v42 + ((v42 + "foo bar") + (v42 + v42))))
                            v52 = (-(v52))
                        }
                    }
                    v3 = v3
                    v3 = v3
                }
                var v63 int8 = 100
                v52 -= v52
                _ = fi62
                _ = v63
            }
            if (!(false || ((true || (42 <= 1)) && true))) {
                switch -1 {
                case -1:
                    v52 += v52
                }
                var v64 string = (v42 + "")
                _ = v64
            }
            v4 -= v4
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v52
    _ = v53
    _ = v56
    _ = v57
    _ = v61
}
