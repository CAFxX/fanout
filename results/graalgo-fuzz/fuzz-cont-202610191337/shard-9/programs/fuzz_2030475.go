package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int; F2 int16 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint8; F1 int32; F2 int32 }
type GenS5 struct { F0 bool; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 int32; F1 int64 }
type GenS9 struct { F0 uint64; F1 uint64 }
func genFn0(p0 string) uint32 {
    v4 := GenS1{F0: 255}
    if (32767 >= 255) {
        v5 := GenS2{F0: 794, F1: 339, F2: 1}
        v6 := GenS3{F0: 127, F1: 9223372036854775807}
        v7 := []int{4294967295, 32767, 128, 127}
        _ = v5
        _ = v6
        _ = v7
    } else {
        p0 = "a"
    }
    v8 := GenS3{F0: 100, F1: -9223372036854775808}
    v9 := []int{4294967295, 722, 9223372036854775807}
    _ = p0
    _ = v4
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := []string{}
    var v2 float32 = 0.1
    v3 := make(map[int]uint32)
    switch (("\n\t" + "\n\t") + "a") {
    case "a":
        v10 := GenS1{F0: 2}
        _ = v10
    case "":
        var v11 string = "\n\t"
        _ = v11
    }
    v1 = append(v1, "foo bar")
    if (!true) {
        v2 -= -2.25
        v2 -= 3.14159
    } else {
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := GenS4{F0: 42, F1: 0, F2: 42}
            _ = fi12
            _ = v13
        }
    }
    switch -1 {
    case 2147483647:
        if true {
            fmt.Println(v2)
            for fi14 := 0; fi14 < 8; fi14++ {
                switch "\n\t" {
                case "foo bar":
                    v3[-9223372036854775808] = 225
                default:
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 uint8 = 60
                        _ = fi15
                        _ = v16
                    }
                }
                v17 := GenS2{F0: 2147483647, F1: (^(fi14)), F2: 929}
                v18 := GenS5{F0: true, F1: -1, F2: 923}
                _ = fi14
                _ = v17
                _ = v18
            }
        }
    case 0:
        if false {
            switch "" {
            case "\n\t":
                v2 += ((v2 + v2) * (v2 - 1e10))
            case "hello":
                var v19 string = ("hello" + "hello")
                _ = v19
            default:
                v20 := GenS5{F0: true, F1: 343, F2: 857}
                _ = v20
            }
        }
    default:
        var v21 int16 = -760
        _ = v21
    }
    if ((1e10 / 0.1) != 3.14159) {
        v2 = (v2 / (v2 / v2))
    }
    v22 := GenS6{F0: 893, F1: 0}
    for fi23 := 0; fi23 < 1; fi23++ {
        if true {
            var v24 int16 = 1000
            _ = v24
        } else {
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 uint8 = 255
                switch (("a" + ("hello" + "foo bar")) + ("foo bar" + "foo bar")) {
                case "a":
                    v27 := GenS7{F0: v26, F1: 65535}
                    _ = v27
                default:
                    v28 := GenS2{F0: 1, F1: (fi23 / ((4294967295) | 1)), F2: 291}
                    _ = v28
                }
                v2 = (v2 + float32(-2.25))
                _ = fi25
                _ = v26
            }
            if false {
                for fi29 := 0; fi29 < 10; fi29++ {
                    var v30 uint8 = 2
                    _ = fi29
                    _ = v30
                }
            } else {
                fmt.Println(v2, fi23)
                switch "a" {
                case "\n\t":
                    fmt.Println(fi23, v2)
                case "a":
                    v2 -= v2
                case "":
                    if true {
                        fmt.Println(fi23)
                        var v31 int16 = -1
                        _ = v31
                    } else {
                        for fi32 := 0; fi32 < 12; fi32++ {
                            v33 := GenS3{F0: 1, F1: 1}
                            if false {
                                v2 -= (float32(3.14159) - (v2 - 0.1))
                                switch ("hello" + "foo bar") {
                                case "":
                                    fmt.Println(fi23, fi23, v2)
                                case "\n\t":
                                    switch "" {
                                    case "hello":
                                        if true {
                                            v34 := GenS4{F0: 255, F1: -798, F2: -1}
                                            fmt.Println(fi32, fi23)
                                            v2 = v2
                                            v35 := GenS5{F0: true, F1: -9223372036854775808, F2: 0}
                                            var v36 uint64 = 17
                                            fmt.Println(v36, v2)
                                            var v37 uint32 = genFn0("a")
                                            v38 := GenS5{F0: true, F1: 42, F2: 721}
                                            v39 := GenS8{F0: 0, F1: 9223372036854775807}
                                            var v40 uint64 = v36
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                        v2 -= 3.14159
                                        v41 := GenS9{F0: 4294967295, F1: 42}
                                        _ = v41
                                    }
                                }
                            }
                            _ = fi32
                            _ = v33
                        }
                        v2 += -2.25
                    }
                default:
                    for fi42 := 0; fi42 < 16; fi42++ {
                        v2 += 0.0
                        var v43 int64 = 1
                        var v44 bool = true
                        var v45 bool = ((v44 || (42 != -2147483648)) && (32767 < 492))
                        if v45 {
                            v2 = (v2 * 1e10)
                            if v45 {
                                var v46 int = (fi42 * fi42)
                                v45 = (v44 || (1 >= 0))
                                var v47 uint8 = 42
                                var v48 uint = 255
                                _ = v46
                                _ = v47
                                _ = v48
                            } else {
                                switch fi23 {
                                case 42:
                                    if (((v2 / -2.25) * (1e10 + 1e10)) == (float32(3.14159) / 3.14159)) {
                                        v2 = ((v2 - 1.5) / -2.25)
                                    } else {
                                        var v49 int = (^(fi42))
                                        _ = v49
                                    }
                                case 2147483647:
                                    fmt.Println(v45, fi23, v45)
                                default:
                                    var v50 uint64 = 493
                                    _ = v50
                                }
                                if (!v44) {
                                    var v51 uint = 0
                                    _ = v51
                                }
                                if v44 {
                                    var v52 int32 = 1
                                    _ = v52
                                }
                                v45 = v44
                            }
                        }
                        v53 := GenS3{F0: -1, F1: (v43 & (v43 / ((v43) | 1)))}
                        _ = fi42
                        _ = v43
                        _ = v44
                        _ = v45
                        _ = v53
                    }
                }
            }
        }
        fmt.Println(fi23, v2)
        _ = fi23
    }
    _ = len(v1)
    v54 := GenS5{F0: false, F1: 206, F2: 32767}
    for fi55 := 0; fi55 < 10; fi55++ {
        fmt.Println(fi55, fi55)
        switch "a" {
        case "foo bar":
            v2 = 0.0
        case "hello":
            var v56 bool = true
            _ = v56
        case "\n\t":
            fmt.Println(fi55)
        }
        _ = fi55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v54
}
