package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
type GenS2 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 0
    var v3 int16 = -1
    for fi4 := 0; fi4 < 15; fi4++ {
        v2 = v2
        for fi5 := 0; fi5 < 13; fi5++ {
            v2 += (v2 << ((1) & 63))
            _ = fi5
        }
        if false {
            if true {
                v6 := v1[899]
                var v7 uint32 = 0
                _ = v6
                _ = v7
            } else {
                v8 := GenS1{F0: (v2 & (v2 * 65535)), F1: 1, F2: false}
                _ = v8
            }
            fmt.Println(v2)
            v3 -= (v3 / ((v3) | 1))
        }
        _ = fi4
    }
    v9 := GenS1{F0: 65535, F1: 18446744073709551615, F2: false}
    v10 := []int{877, 0, 127}
    var v11 uint32 = 0
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v12 := GenS1{F0: (v2 - 649), F1: 4294967295, F2: ((0 > 32) && (4294967295 < 0))}
    _ = len(v1)
    if (0 != -9223372036854775808) {
        switch (("a" + ("" + "hello")) + "a") {
        case "hello":
            var v13 int16 = v3
            _ = v13
        case "foo bar":
            var v14 float32 = 3.14159
            _ = v14
        case "a":
            v11 += (v11 * 42)
        }
        v15 := GenS1{F0: 255, F1: 42, F2: true}
        v16 := GenS1{F0: 0, F1: 2, F2: true}
        _ = v15
        _ = v16
    } else {
        delete(v1, 1)
        v3 -= v3
    }
    v17 := GenS1{F0: (v2 + (v2 / (((^(v2))) | 1))), F1: 570, F2: false}
    for fi18 := 0; fi18 < 12; fi18++ {
        switch "\n\t" {
        case "a":
            if (42 <= 905) {
                if ((v2 + (^(v2))) > 255) {
                    v19 := GenS1{F0: v2, F1: 2, F2: true}
                    var v20 uint64 = 4294967295
                    if (0 <= 55) {
                        switch (fi18 + (-(fi18))) {
                        case -956:
                            v20 = v20
                        case 32767:
                            v3 = v3
                        default:
                            for fi21 := 0; fi21 < 6; fi21++ {
                                _ = len(v10)
                                _ = fi21
                            }
                        }
                        if (("" + "a") > "") {
                            v20 += v20
                            v2 = v2
                            v11 += (v11 >> ((1) & 63))
                        } else {
                            v22 := GenS2{F0: v2, F1: 17, F2: 42}
                            fmt.Println(v20, v20)
                            _ = v22
                        }
                        v20 = v20
                    }
                    _ = v19
                    _ = v20
                }
                if (!true) {
                    switch "\n\t" {
                    case "":
                        v3 += v3
                    case "hello":
                        for fi23 := 0; fi23 < 14; fi23++ {
                            if (42 <= 128) {
                                switch "hello" {
                                case "hello":
                                    v11 = v11
                                case "a":
                                    v11 = 0
                                }
                                if false {
                                    switch "" {
                                    case "":
                                        switch fi18 {
                                        case 42:
                                            if len(v10) > 0 {
                                                v10[0] = 32767
                                            }
                                        case 334:
                                            if (false && ((^(fi18)) < -146)) {
                                                v24 := v1[(v11 >> ((42) & 63))]
                                                _ = v24
                                            } else {
                                                v25 := GenS3{F0: 32767}
                                                v26 := GenS4{F0: -127}
                                                _ = v25
                                                _ = v26
                                            }
                                        case 32768:
                                            v27 := GenS5{F0: (v3 % (((v3 << ((0) & 63))) | 1)), F1: ""}
                                            _ = v27
                                        }
                                    }
                                    v28 := v1[(v11 - v11)]
                                    if (1 != 0) {
                                        if true {
                                            v1[(v11 & 1)] = "foo bar"
                                        }
                                        var v29 float32 = 3.14159
                                        v11 += 42
                                        _ = v29
                                    } else {
                                        fmt.Println(v11)
                                    }
                                    _ = v28
                                } else {
                                    v11 -= v11
                                    fmt.Println(fi18, v11)
                                    v11 -= v11
                                }
                                var v30 uint = 984
                                _ = v30
                            } else {
                                v11 = (^(v11))
                            }
                            v31 := GenS6{F0: 0}
                            var v32 uint8 = 42
                            _ = fi23
                            _ = v31
                            _ = v32
                        }
                    case "a":
                        switch (("" + "hello") + (((("hello" + "foo bar") + "foo bar") + "hello") + ("hello" + "hello"))) {
                        case "foo bar":
                            v33 := v1[(v11 / (((v11 | v11)) | 1))]
                            _ = v33
                        default:
                            var v34 uint = uint(v2)
                            _ = v34
                        }
                    }
                    v35 := GenS5{F0: v3, F1: "foo bar"}
                    var v36 string = ""
                    _ = v35
                    _ = v36
                }
                var v37 uint64 = 0
                _ = v37
            } else {
                if ((v2 & (v2 / ((255) | 1))) != 42) {
                    if (!(2147483647 < v11)) {
                        v38 := GenS7{F0: -2147483648, F1: fi18}
                        _ = v38
                    } else {
                        switch ((("a" + "a") + "a") + "a") {
                        case "\n\t":
                            switch "" {
                            case "hello":
                                v39 := GenS8{F0: true}
                                _ = v39
                            case "a":
                                for fi40 := 0; fi40 < 15; fi40++ {
                                    v2 = (v2 / ((42) | 1))
                                    _ = fi40
                                }
                            default:
                                for fi41 := 0; fi41 < 15; fi41++ {
                                    v42 := GenS9{F0: v11, F1: v11}
                                    var v43 float64 = (((float64(3.14159) + (-2.25 + float64(0.1))) / (1.5 - (float64(3.14159) - 0.1))) * 3.14159)
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    _ = fi41
                                    _ = v42
                                    _ = v43
                                }
                            }
                        case "":
                            if (!(float64(-2.25) < 3.14159)) {
                                if true {
                                    var v44 uint8 = 128
                                    _ = v44
                                } else {
                                    var v45 uint = 1
                                    _ = v45
                                }
                            }
                        case "a":
                            for fi46 := 0; fi46 < 13; fi46++ {
                                v47 := GenS7{F0: 0, F1: -1}
                                _ = fi46
                                _ = v47
                            }
                        default:
                            v2 -= 42
                        }
                        var v48 int64 = -403
                        _ = v48
                    }
                    fmt.Println(v3, fi18)
                }
            }
        }
        var v49 int64 = -1
        _ = fi18
        _ = v49
    }
    v50 := v1[42]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v50
    _ = v9
}
