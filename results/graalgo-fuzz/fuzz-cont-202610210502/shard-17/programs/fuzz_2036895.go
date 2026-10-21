package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint32 = 304
    var v2 int64 = 759
    v3 := make(map[int]string)
    var v4 float64 = 3.14159
    v2 = v2
    switch -1 {
    case -276:
        var v5 int64 = v2
        _ = v5
    case -1:
        v2 += v2
    default:
        var v6 uint8 = 1
        _ = v6
    }
    var v7 int32 = 969
    switch ("hello" + ("hello" + "hello")) {
    case "":
        v8 := []int{-265, 42, 0, 32768}
        _ = v8
    case "foo bar":
        switch ("\n\t" + (("foo bar" + "hello") + "\n\t")) {
        case "foo bar":
            var v9 uint16 = 255
            _ = v9
        case "a":
            v10 := GenS1{F0: 2, F1: -32768, F2: 298}
            _ = v10
        case "hello":
            switch 317 {
            case -9223372036854775808:
                var v11 uint8 = 128
                _ = v11
            case 9223372036854775807:
                v7 += v7
            default:
                v12 := GenS1{F0: 2, F1: 1, F2: -1}
                _ = v12
            }
        }
    }
    var v13 float32 = 3.14159
    var v14 uint32 = v1
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 int64 = v2
            if false {
                if ((v17 - v17) == v17) {
                    var v18 uint64 = 1
                    if true {
                        v19 := GenS1{F0: 42, F1: 1000, F2: 42}
                        v4 += v4
                        switch ("a" + ((("a" + "foo bar") + (("hello" + "") + "hello")) + ("hello" + "foo bar"))) {
                        case "foo bar":
                            v20 := GenS2{F0: 255, F1: v18, F2: 65535}
                            _ = v20
                        case "hello":
                            v21 := GenS3{F0: v18, F1: (fi15 - (^(fi16)))}
                            _ = v21
                        default:
                            fmt.Println(v18, v1, v18)
                        }
                        _ = v19
                    } else {
                        var v22 int = -622
                        switch "hello" {
                        case "\n\t":
                            v23 := []int{fi15, (v22 - 128)}
                            _ = v23
                        }
                        _ = v22
                    }
                    var v24 int = fi15
                    _ = v18
                    _ = v24
                }
                v2 = -1
            } else {
                var v25 uint16 = 2
                v26 := []int{int(v1), (fi16 | fi16)}
                v14 = (^(v14))
                _ = v25
                _ = v26
            }
            v27 := GenS4{F0: 42}
            _ = fi16
            _ = v17
            _ = v27
        }
        switch ("\n\t" + "a") {
        case "hello":
            if ((v14 / ((v14) | 1)) > (v14 - v14)) {
                if true {
                    var v28 float32 = (v13 - (v13 - ((float32((0.0 - ((-2.25 - (-2.25 * v4)) - v4))) / v13) - (v13 / v13))))
                    if false {
                        v28 = v13
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            var v29 uint = uint(v14)
                            _ = v29
                        case "a":
                            if (1 != 2147483647) {
                                fmt.Println(v28)
                                var v30 bool = true
                                var v31 int64 = 0
                                _ = v30
                                _ = v31
                            } else {
                                for fi32 := 0; fi32 < 17; fi32++ {
                                    fmt.Println(v28)
                                    var v33 string = "\n\t"
                                    var v34 uint64 = 42
                                    _ = fi32
                                    _ = v33
                                    _ = v34
                                }
                                v14 += (v14 * v1)
                                v35 := GenS3{F0: 0, F1: (fi15 % ((-958) | 1))}
                                _ = v35
                            }
                        case "\n\t":
                            var v36 uint32 = (v14 | (v14 / (((v1 + 255)) | 1)))
                            _ = v36
                        }
                        v14 += v14
                        var v37 bool = (fi15 >= (fi15 >> ((557) & 63)))
                        _ = v37
                    }
                    var v38 uint8 = 42
                    _ = v28
                    _ = v38
                } else {
                    v39 := GenS3{F0: 42, F1: (^(fi15))}
                    for fi40 := 0; fi40 < 16; fi40++ {
                        v7 = 1
                        _ = fi40
                    }
                    _ = v39
                }
            } else {
                switch "hello" {
                case "foo bar":
                    v4 -= v4
                case "a":
                    var v41 uint8 = 2
                    _ = v41
                }
            }
        case "a":
            if (42 >= 1) {
                switch "a" {
                case "a":
                    v13 = (v13 - v13)
                case "hello":
                    var v42 float32 = (3.14159 - float32((v4 - v4)))
                    _ = v42
                default:
                    v2 += 42
                }
                var v43 int16 = -32768
                switch "" {
                case "\n\t":
                    for fi44 := 0; fi44 < 15; fi44++ {
                        fmt.Println(v2, v43)
                        var v45 int32 = (v7 % (((v7 * v7)) | 1))
                        _ = fi44
                        _ = v45
                    }
                case "hello":
                    fmt.Println(v13, v1, v2)
                case "a":
                    var v46 uint64 = 1
                    _ = v46
                }
                _ = v43
            } else {
                if false {
                    v2 = -9223372036854775808
                }
            }
        default:
            var v47 uint32 = v1
            _ = v47
        }
        for fi48 := 0; fi48 < 7; fi48++ {
            delete(v3, (fi15 - fi15))
            v14 -= v1
            switch 1 {
            case 32768:
                var v49 uint16 = 65535
                _ = v49
            case -1:
                var v50 float64 = (((0.0 + 0.1) - v4) / (v4 / 1.5))
                _ = v50
            default:
                var v51 float32 = 0.0
                _ = v51
            }
            var v52 float32 = -2.25
            if true {
                var v53 uint64 = 394
                _ = v53
            }
            v7 = (-(v7))
            switch "hello" {
            case "":
                var v54 int64 = int64(v1)
                _ = v54
            }
            v4 = (v4 / (0.1 + (-2.25 / (v4 * 1e10))))
            _ = fi48
            _ = v52
        }
        _ = fi15
    }
    v1 += 0
    var v55 float32 = v13
    v14 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v7
}
