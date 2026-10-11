package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16 }
func genFn0() uint16 {
    if (((true || true) || false) && false) {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int64 = -1
            v9 := []int{}
            for fi10 := 0; fi10 < 18; fi10++ {
                v9 = append(v9, (fi10 + (fi7 ^ 42)))
                _ = len(v9)
                var v11 uint32 = 4294967295
                var v12 float32 = float32((0.0 / 3.14159))
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        var v13 uint16 = 472
        _ = v13
    }
    if (-1 != -128) {
        var v14 bool = false
        v14 = v14
        _ = v14
    }
    if (!(2147483648 >= 980)) {
        v15 := []int{-9223372036854775808, 128}
        if len(v15) > 0 {
            v15[0] = 731
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            fmt.Println(fi16)
            switch fi16 {
            case 1:
                var v17 int64 = -878
                _ = v17
            case 9223372036854775807:
                fmt.Println(fi16)
            default:
                switch "hello" {
                case "\n\t":
                    if true {
                        var v18 int64 = -1
                        _ = len(v15)
                        _ = v18
                    }
                case "hello":
                    var v19 bool = false
                    _ = v19
                }
            }
            var v20 int16 = 0
            _ = fi16
            _ = v20
        }
        _ = v15
    } else {
        if true {
            v21 := []int{}
            _ = v21
        }
        if (42 != 733) {
            v22 := GenS1{F0: 127, F1: 32767, F2: 188}
            var v23 uint8 = 128
            _ = v22
            _ = v23
        }
        var v24 int64 = 9223372036854775807
        _ = v24
    }
    if (42 > 2) {
        v25 := GenS1{F0: 125, F1: 32767, F2: 2}
        switch "a" {
        case "\n\t":
            v26 := []int{9223372036854775807, 32767}
            _ = v26
        case "foo bar":
            v27 := GenS2{F0: 100, F1: -1}
            _ = v27
        default:
            v28 := GenS3{F0: "a", F1: 56}
            _ = v28
        }
        switch "a" {
        case "":
            var v29 uint64 = 1
            _ = v29
        case "hello":
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 int = fi30
                var v32 int16 = 362
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        _ = v25
    }
    switch ("a" + "\n\t") {
    case "a":
        var v33 uint = 51
        _ = v33
    case "hello":
        v34 := []int{-2147483648, -437, 9223372036854775807}
        _ = v34
    case "\n\t":
        var v35 float64 = 0.0
        _ = v35
    default:
        v36 := GenS3{F0: "hello", F1: 127}
        _ = v36
    }
    return 746
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = -2147483648
    var v3 uint32 = 889
    v4 := make(map[string]int)
    v5 := []string{}
    v6 := []int{-2147483648, -379, 128}
    v3 += 4294967295
    for fi37 := 0; fi37 < 12; fi37++ {
        v2 -= (v2 % ((1) | 1))
        for fi38 := 0; fi38 < 17; fi38++ {
            v3 = 0
            if (400 > 42) {
                var v39 float32 = 1e10
                v5 = append(v5, "foo bar")
                _ = v39
            }
            v3 = 0
            _ = fi38
        }
        _ = fi37
    }
    v40 := GenS3{F0: "foo bar", F1: 255}
    var v41 uint = uint(v3)
    var v42 uint8 = 79
    var v43 float64 = 1e10
    fmt.Println(v41, v41)
    v43 += v43
    var v44 float64 = 0.0
    var v45 uint64 = uint64(v42)
    fmt.Println(v43)
    var v46 uint64 = (v1 | (v45 - 42))
    v42 -= (v42 ^ (v42 << (((v41 ^ (v41 & v41))) & 63)))
    switch 32767 {
    case 660:
        if false {
            _ = len(v4)
            v6 = append(v6, 9223372036854775807)
            v4[("" + "a")] = 32767
        }
    case 9223372036854775807:
        fmt.Println(v2)
    }
    for fi47 := 0; fi47 < 4; fi47++ {
        var v48 uint32 = 65535
        for fi49 := 0; fi49 < 10; fi49++ {
            v5 = append(v5, "a")
            var v50 uint32 = 2147483647
            switch fi49 {
            case -577:
                v2 -= 0
            case 4294967295:
                v45 -= v46
            case 127:
                if (false && true) {
                    v44 = v44
                    var v51 int8 = -50
                    v48 -= v3
                    _ = v51
                } else {
                    var v52 uint16 = 65535
                    var v53 string = "\n\t"
                    _ = v52
                    _ = v53
                }
            default:
                var v54 uint32 = v48
                _ = v54
            }
            _ = fi49
            _ = v50
        }
        for fi55 := 0; fi55 < 16; fi55++ {
            if (42 <= (v48 << (((v41 - v41)) & 63))) {
                v1 = (v45 + 567)
                v56 := GenS1{F0: (^(v42)), F1: genFn0(), F2: 83}
                var v57 float64 = (v43 / v43)
                _ = v56
                _ = v57
            } else {
                switch int(v41) {
                case 32767:
                    var v58 float32 = (((float32(0.0) / -2.25) + 1e10) / 3.14159)
                    _ = v58
                default:
                    if (true || ((!((float64((3.14159 + 1e10)) / v43) < (v43 - v43))) || false)) {
                        _ = len(v4)
                    } else {
                        switch (fi55 | (fi47 - 32767)) {
                        case 128:
                            v59 := GenS4{F0: (fi55 * (-(fi55)))}
                            _ = v59
                        case -9223372036854775808:
                            v3 += (v48 & (v3 / (((v48 * v3)) | 1)))
                        case 32767:
                            v60 := GenS5{F0: 32767}
                            _ = v60
                        default:
                            v46 += (v46 ^ (v45 | v45))
                        }
                        var v61 int = 127
                        fmt.Println(v41)
                        _ = v61
                    }
                }
            }
            _ = fi55
        }
        _ = fi47
        _ = v48
    }
    v62 := GenS2{F0: 37, F1: -53}
    if (true && true) {
        var v63 uint16 = 255
        v64 := v4[("\n\t" + "\n\t")]
        _ = v63
        _ = v64
    } else {
        var v65 int = 127
        _ = v65
    }
    switch ((("" + "hello") + "") + "hello") {
    case "foo bar":
        for fi66 := 0; fi66 < 8; fi66++ {
            fmt.Println(v46, v42, v44)
            if (false && false) {
                var v67 uint64 = 1
                fmt.Println(v67)
                _ = v67
            }
            var v68 bool = false
            _ = fi66
            _ = v68
        }
    default:
        v45 = 576
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v62
}
