package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int64; F2 int }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 uint16; F2 string }

func main() {
    var v1 string = "hello"
    var v2 uint64 = 942
    var v3 bool = true
    v4 := []int{32768, 593, 32767}
    v5 := make(map[int]uint32)
    var v6 uint = 470
    if v3 {
        var v7 float32 = -2.25
        _ = v7
    } else {
        switch 42 {
        case 0:
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 uint32 = 724
                if (-32768 > 503) {
                    var v10 uint32 = (v9 ^ uint32(v6))
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 8; fi11++ {
                        v12 := GenS1{F0: v1}
                        v13 := v5[fi8]
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
                var v14 int64 = 199
                _ = fi8
                _ = v14
                _ = v9
            }
        case 1:
            for fi15 := 0; fi15 < 5; fi15++ {
                fmt.Println(v2, v3)
                _ = fi15
            }
        }
        var v16 float32 = 1.5
        if false {
            v17 := GenS2{F0: false, F1: 1, F2: 187}
            v1 = "a"
            _ = v17
        }
        _ = v16
    }
    fmt.Println(v1)
    v18 := GenS3{F0: -60, F1: v6}
    var v19 string = v1
    if ((0 <= 127) || v3) {
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 string = (v1 + (v19 + v1))
            v3 = true
            var v22 float64 = 0.0
            _ = fi20
            _ = v21
            _ = v22
        }
    } else {
        var v23 int32 = -2147483648
        _ = v23
    }
    var v24 float64 = 1e10
    switch (v19 + "a") {
    case "":
        switch (v1 + "foo bar") {
        case "hello":
            v25 := v5[32767]
            _ = v25
        default:
            var v26 uint32 = 42
            _ = v26
        }
    case "a":
        if false {
            v27 := GenS4{F0: 42}
            if (-952 == 52) {
                for fi28 := 0; fi28 < 8; fi28++ {
                    for fi29 := 0; fi29 < 13; fi29++ {
                        v2 = 780
                        _ = fi29
                    }
                    var v30 int32 = 1
                    _ = fi28
                    _ = v30
                }
                for fi31 := 0; fi31 < 11; fi31++ {
                    if v3 {
                        var v32 uint64 = v2
                        for fi33 := 0; fi33 < 15; fi33++ {
                            var v34 int16 = -709
                            switch v19 {
                            case "a":
                                switch (v1 + (v19 + "\n\t")) {
                                case "foo bar":
                                    var v35 float64 = ((v24 + v24) + v24)
                                    _ = v35
                                }
                            case "hello":
                                v3 = v3
                            }
                            var v36 uint8 = uint8(v6)
                            _ = fi33
                            _ = v34
                            _ = v36
                        }
                        v24 = v24
                        _ = v32
                    }
                    if (127 > 212) {
                        fmt.Println(v19, v2)
                        v37 := GenS5{F0: v19, F1: (v6 ^ 42)}
                        if (!(uint32(v2) == 846)) {
                            var v38 string = ("foo bar" + ("\n\t" + v19))
                            var v39 uint64 = (v2 % (((v2 % ((v2) | 1))) | 1))
                            _ = v38
                            _ = v39
                        } else {
                            fmt.Println(v2, v24)
                            if (false && (v3 && true)) {
                                var v40 float64 = (v24 / v24)
                                _ = v40
                            }
                            var v41 uint8 = 128
                            _ = v41
                        }
                        _ = v37
                    } else {
                        var v42 uint16 = 42
                        _ = v42
                    }
                    if v3 {
                        v6 += v6
                        var v43 uint = (v6 + (v6 ^ v6))
                        if true {
                            for fi44 := 0; fi44 < 17; fi44++ {
                                v3 = v3
                                if v3 {
                                    if len(v4) > 0 {
                                        v4[0] = (fi44 >> ((v43) & 63))
                                    }
                                    _ = len(v5)
                                    if v3 {
                                        fmt.Println(fi44)
                                        fmt.Println(v19)
                                    }
                                }
                                _ = fi44
                            }
                            v45 := GenS6{F0: 67, F1: "a"}
                            switch fi31 {
                            case 4294967295:
                                var v46 uint = (v43 % (((v6 - 65535)) | 1))
                                _ = v46
                            case 629:
                                delete(v5, fi31)
                            case -630:
                                fmt.Println(v24)
                            }
                            _ = v45
                        } else {
                            fmt.Println(v19, v43, v2)
                            if true {
                                v47 := GenS5{F0: "a", F1: v43}
                                if (86 != 42) {
                                    v24 += 0.0
                                } else {
                                    var v48 int16 = 0
                                    _ = v48
                                }
                                _ = v47
                            } else {
                                var v49 string = v1
                                v50 := GenS7{F0: -45}
                                _ = v49
                                _ = v50
                            }
                        }
                        _ = v43
                    }
                    _ = fi31
                }
            }
            _ = v27
        } else {
            var v51 float32 = (-2.25 * 1.5)
            _ = v51
        }
    default:
        v52 := GenS1{F0: (v1 + "a")}
        _ = v52
    }
    var v53 uint8 = 128
    switch "foo bar" {
    case "hello":
        _ = len(v4)
    case "\n\t":
        v1 = v19
    case "foo bar":
        var v54 uint32 = 0
        _ = v54
    }
    if true {
        fmt.Println(v2, v3)
        switch 32768 {
        case 32768:
            var v55 int32 = -1
            _ = v55
        }
    } else {
        var v56 int32 = 2147483647
        v57 := GenS8{F0: -1, F1: 397, F2: ("a" + v19)}
        _ = v56
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
