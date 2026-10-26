package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0() int32 {
    var v6 uint8 = 127
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 16; fi8++ {
        v7 = ("" + v7)
        switch (v7 + v7) {
        case "\n\t":
            v9 := GenS1{F0: 4294967295, F1: true, F2: 359}
            _ = v9
        }
        _ = fi8
    }
    _ = v6
    _ = v7
    return 1
}
func genFn1() bool {
    v10 := []int{-48, 32767, 128, 128}
    var v11 string = "hello"
    v11 = (v11 + v11)
    if true {
        v12 := GenS2{F0: 42}
        _ = v12
    }
    switch v11 {
    case "a":
        v13 := GenS2{F0: 2}
        _ = v13
    case "\n\t":
        v11 = (((v11 + (v11 + "")) + "a") + v11)
    }
    _ = v10
    _ = v11
    return (40 != -1)
}

func main() {
    var v1 float64 = 0.1
    var v2 uint16 = 42
    var v3 int32 = 0
    var v4 float64 = v1
    v5 := make(map[string]uint32)
    for fi14 := 0; fi14 < 12; fi14++ {
        switch "foo bar" {
        case "a":
            v15 := GenS2{F0: 4294967295}
            _ = v15
        case "foo bar":
            for fi16 := 0; fi16 < 18; fi16++ {
                v4 = (v4 - v1)
                if (((-128 < int8(fi16)) || false) && false) {
                    switch "" {
                    case "hello":
                        switch (^(fi16)) {
                        case 148:
                            _ = len(v5)
                        case -9223372036854775808:
                            v17 := GenS2{F0: 42}
                            _ = v17
                        case 1:
                            var v18 bool = false
                            _ = v18
                        default:
                            v19 := GenS2{F0: 42}
                            _ = v19
                        }
                    case "":
                        v20 := GenS1{F0: 65535, F1: (382 < 18446744073709551615), F2: v2}
                        _ = v20
                    case "foo bar":
                        if ("" != ("foo bar" + "\n\t")) {
                            switch (-(fi16)) {
                            case 0:
                                var v21 string = "a"
                                _ = v21
                            case 127:
                                var v22 int8 = 103
                                _ = v22
                            }
                        }
                    default:
                        if (-572 > -9223372036854775808) {
                            var v23 uint16 = 42
                            _ = v23
                        } else {
                            v3 = -1
                            var v24 int = fi14
                            _ = v24
                        }
                        var v25 float64 = (v1 + v1)
                        _ = v25
                    }
                    var v26 uint8 = 214
                    var v27 bool = (!(128 >= 467))
                    _ = v26
                    _ = v27
                } else {
                    v28 := []int{fi14, (fi14 >> ((65535) & 63)), (fi16 >> ((255) & 63)), fi14}
                    var v29 bool = (42 >= -1)
                    v4 = v1
                    v30 := GenS2{F0: 1}
                    _ = v28
                    _ = v29
                    _ = v30
                }
                var v31 int16 = 1
                if (2 > 192) {
                    v1 -= v4
                    if true {
                        fmt.Println(v31)
                    } else {
                        v32 := []int{(fi14 + (fi14 << ((426) & 63))), 4294967295, fi14}
                        _ = v32
                    }
                    v3 += (v3 * -672)
                } else {
                    var v33 int16 = v31
                    _ = v33
                }
                var v34 float64 = 3.14159
                _ = fi16
                _ = v31
                _ = v34
            }
        default:
            switch "foo bar" {
            case "":
                v35 := GenS1{F0: 255, F1: (((v4 > v1) && true) && true), F2: (v2 & (v2 | (v2 - (^(v2)))))}
                _ = v35
            default:
                for fi36 := 0; fi36 < 15; fi36++ {
                    v3 = v3
                    v4 = float64(((1e10 + 0.0) / (1e10 * 1.5)))
                    if true {
                        var v37 int8 = -81
                        _ = v37
                    }
                    fmt.Println(v3)
                    _ = fi36
                }
            }
        }
        _ = fi14
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        v2 -= v2
        _ = fi38
    }
    var v39 uint32 = 0
    switch "hello" {
    case "foo bar":
        var v40 uint16 = (v2 * (v2 + v2))
        _ = v40
    case "":
        v2 = 42
    case "a":
        if (!false) {
            fmt.Println(v3, v2)
            v39 -= uint32(v2)
            var v41 int16 = -32768
            _ = v41
        } else {
            for fi42 := 0; fi42 < 2; fi42++ {
                var v43 int64 = 217
                _ = fi42
                _ = v43
            }
            v2 += (^(v2))
        }
    default:
        switch 597 {
        case -2147483648:
            if false {
                var v44 float32 = 0.0
                _ = v44
            }
        case 42:
            if ((false || (v1 > 0.0)) || (!false)) {
                v45 := GenS2{F0: 598}
                _ = v45
            }
        case -489:
            fmt.Println(v1, v2, v2)
        }
    }
    v46 := v5["\n\t"]
    fmt.Println(v3, v39)
    for fi47 := 0; fi47 < 19; fi47++ {
        fmt.Println(fi47, fi47, fi47)
        switch ("a" + "a") {
        case "hello":
            fmt.Println(v3, v3)
        default:
            for fi48 := 0; fi48 < 10; fi48++ {
                var v49 uint16 = v2
                v50 := GenS3{F0: false, F1: (v2 / ((255) | 1))}
                var v51 int = fi48
                _ = fi48
                _ = v49
                _ = v50
                _ = v51
            }
        }
        switch -1 {
        case 4294967295:
            fmt.Println(v4)
        default:
            v52 := []int{0, fi47}
            _ = v52
        }
        _ = fi47
    }
    v53 := []int{}
    var v54 float32 = -2.25
    for fi55 := 0; fi55 < 14; fi55++ {
        v2 = uint16(v39)
        if (fi55 < fi55) {
            fmt.Println(v46, v46)
        } else {
            fmt.Println(v3, v46)
        }
        _ = fi55
    }
    v56 := v5["a"]
    var v57 int16 = -466
    _ = len(v53)
    v57 += (v57 / (((v57 % ((v57) | 1))) | 1))
    if (v54 != -2.25) {
        if len(v53) >= 2 {
            v53 = v53[1:len(v53)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v46
    _ = v5
    _ = v53
    _ = v54
    _ = v56
    _ = v57
}
