package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int32; F1 uint32; F2 bool }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 int32; F2 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 int16; F2 uint16 }
func genFn0(p0 int8) int16 {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 uint = 255
        v6 -= (v6 - (v6 & 65535))
        v6 += (v6 + (v6 | (v6 ^ 255)))
        _ = fi5
        _ = v6
    }
    switch (("foo bar" + "") + "\n\t") {
    case "\n\t":
        v7 := GenS1{F0: 4294967295, F1: 0}
        _ = v7
    case "hello":
        if true {
            v8 := []int{-51}
            _ = v8
        }
    case "a":
        p0 -= p0
    default:
        fmt.Println(p0)
    }
    _ = p0
    return -794
}

func main() {
    var v1 uint64 = 0
    v2 := []int{32767}
    var v3 uint32 = 2147483647
    v4 := make(map[int]bool)
    fmt.Println(v3, v1)
    v9 := GenS2{F0: 612, F1: 117, F2: false}
    fmt.Println(v3)
    v10 := GenS1{F0: 588, F1: -42}
    if (42 != 65535) {
        v1 -= v1
        if false {
            v11 := GenS1{F0: (^(v1)), F1: 128}
            var v12 int = -255
            var v13 int = v12
            _ = v11
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 uint8 = 128
        switch (fi14 + (fi14 << ((42) & 63))) {
        case 2147483647:
            var v16 float32 = 3.14159
            _ = v16
        case 852:
            v1 += (v1 + (v1 / ((v1) | 1)))
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v1, v3)
    for fi17 := 0; fi17 < 12; fi17++ {
        if true {
            var v18 int8 = 100
            if ((!false) || false) {
                if (1e10 <= 0.0) {
                    if (true || (!(!(!(2 <= (v1 << ((18446744073709551615) & 63))))))) {
                        fmt.Println(v1, fi17, v18)
                    } else {
                        fmt.Println(v3, v18, v18)
                        v19 := GenS3{F0: 0, F1: 277}
                        _ = v19
                    }
                    var v20 uint = 864
                    _ = v20
                } else {
                    var v21 float32 = 0.0
                    v18 = 100
                    _ = v21
                }
                if false {
                    var v22 float32 = 1.5
                    for fi23 := 0; fi23 < 4; fi23++ {
                        if (2 >= 2) {
                            fmt.Println(v22, v22)
                            v22 -= 3.14159
                            if (!(true || false)) {
                                v22 = v22
                                delete(v4, fi23)
                            } else {
                                v1 = (v1 * (v1 >> ((42) & 63)))
                                if (!(("hello" + "\n\t") != ("" + "a"))) {
                                    var v24 int64 = 42
                                    if true {
                                        v24 = (v24 + v24)
                                        var v25 uint64 = (v1 * (v1 | 892))
                                        _ = v25
                                    } else {
                                        fmt.Println(fi23, fi23, v22)
                                        var v26 int32 = -1
                                        _ = v26
                                    }
                                    _ = v24
                                } else {
                                    var v27 float64 = -2.25
                                    v28 := GenS4{F0: (v1 >> ((42) & 63)), F1: 255, F2: (v3 / ((v3) | 1))}
                                    _ = v27
                                    _ = v28
                                }
                                var v29 int32 = 1
                                _ = v29
                            }
                        } else {
                            if true {
                                var v30 int16 = 1000
                                var v31 bool = (!true)
                                _ = v30
                                _ = v31
                            }
                            v18 += v18
                            fmt.Println(v1, v18, fi23)
                        }
                        if true {
                            var v32 string = "foo bar"
                            fmt.Println(v3)
                            _ = len(v4)
                            _ = v32
                        }
                        fmt.Println(v1)
                        _ = fi23
                    }
                    v33 := v4[(fi17 & fi17)]
                    _ = v22
                    _ = v33
                }
            } else {
                v34 := GenS1{F0: v1, F1: fi17}
                var v35 int = fi17
                _ = len(v4)
                _ = v34
                _ = v35
            }
            _ = v18
        } else {
            v1 = v1
        }
        v3 = (v3 << ((0) & 63))
        _ = fi17
    }
    var v36 uint16 = 289
    fmt.Println(v36, v1)
    var v37 uint32 = (^(v3))
    if true {
        if true {
            var v38 uint32 = v37
            _ = v38
        }
    } else {
        for fi39 := 0; fi39 < 15; fi39++ {
            for fi40 := 0; fi40 < 10; fi40++ {
                v41 := GenS4{F0: v1, F1: 0, F2: 0}
                _ = fi40
                _ = v41
            }
            fmt.Println(v3, v37, v36)
            v37 = 0
            _ = fi39
        }
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 int16 = 977
        var v44 float64 = float64(0.1)
        _ = fi42
        _ = v43
        _ = v44
    }
    switch -306 {
    case 1:
        switch 1 {
        case 128:
            var v45 uint = 677
            _ = v45
        case 655:
            if (-387 > 1000) {
                var v46 uint64 = 153
                if len(v2) > 0 {
                    v2[0] = -23
                }
                if false {
                    v46 -= v1
                }
                _ = v46
            }
        case -9223372036854775808:
            switch -861 {
            case 127:
                fmt.Println(v1)
            case -119:
                for fi47 := 0; fi47 < 14; fi47++ {
                    v48 := GenS5{F0: 99}
                    var v49 uint = 71
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            default:
                var v50 uint64 = v1
                _ = v50
            }
        default:
            fmt.Println(v36, v36)
        }
    case -1:
        if true {
            switch ("foo bar" + "") {
            case "a":
                v51 := GenS2{F0: 2147483647, F1: (v3 + 419), F2: true}
                _ = v51
            default:
                v37 -= (v37 << ((42) & 63))
            }
            for fi52 := 0; fi52 < 2; fi52++ {
                fmt.Println(v1)
                v53 := GenS6{F0: 0, F1: 114, F2: -1}
                var v54 uint32 = 0
                _ = fi52
                _ = v53
                _ = v54
            }
            v3 -= v3
        } else {
            var v55 bool = (18446744073709551615 < 42)
            if (-32768 < -1) {
                _ = len(v2)
                v36 = 255
                v2 = append(v2, 127)
            }
            v1 = v1
            _ = v55
        }
    default:
        var v56 uint32 = v3
        _ = v56
    }
    var v57 string = (((("hello" + "hello") + ((("hello" + ("a" + (("hello" + (("a" + "") + ("hello" + ("" + "\n\t")))) + "hello"))) + (("hello" + "") + "")) + "\n\t")) + "a") + "hello")
    for fi58 := 0; fi58 < 5; fi58++ {
        if true {
            if (1 < 127) {
                v59 := GenS7{F0: 9223372036854775807}
                _ = v59
            }
        } else {
            _ = len(v4)
            v60 := GenS8{F0: 65535, F1: 605, F2: 786}
            _ = v60
        }
        v1 -= (v1 + (^(v1)))
        _ = fi58
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v57:", v57)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v57
    _ = v9
}
