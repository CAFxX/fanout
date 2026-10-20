package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint; F2 int8 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 int64; F1 int32; F2 bool }
func genFn0() bool {
    v7 := []int{915, 127}
    v8 := GenS1{F0: 42, F1: 1, F2: 32767}
    var v9 uint = 65535
    v10 := GenS2{F0: 4294967295}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1() bool {
    v11 := []int{}
    var v12 int = 375
    v12 = (v12 % (((v12 - 0)) | 1))
    _ = v11
    _ = v12
    return ("" <= (("foo bar" + "") + "a"))
}

func main() {
    v1 := []int{-949, 32767, 42, 127}
    v2 := []int{4294967295, 2147483647, 990, -157}
    v3 := make(map[uint32]int)
    v4 := []string{"foo bar", "hello"}
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]string)
    delete(v3, 714)
    v13 := GenS2{F0: 9223372036854775807}
    var v14 string = "foo bar"
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v14)
    v15 := GenS3{F0: 841, F1: 2, F2: 948}
    for fi16 := 0; fi16 < 10; fi16++ {
        if (238 >= 128) {
            if (1 < -941) {
                switch 729 {
                case -903:
                    delete(v6, (^(fi16)))
                case 127:
                    if (!(((3.14159 + 1e10) == 3.14159) || (1e10 == 1.5))) {
                        fmt.Println(v14)
                        v3[2] = 869
                        if (2 > 2147483647) {
                            fmt.Println(v14, fi16)
                        }
                    } else {
                        switch fi16 {
                        case 1:
                            v14 = v14
                        case 4294967295:
                            delete(v3, 2)
                        case -715:
                            v14 = "hello"
                        }
                        v14 = "\n\t"
                    }
                case 32767:
                    v14 = "foo bar"
                }
                switch "hello" {
                case "a":
                    if true {
                        if ((!false) && true) {
                            var v17 int = -2147483648
                            _ = v17
                        } else {
                            switch v14 {
                            case "a":
                                if false {
                                    for fi18 := 0; fi18 < 9; fi18++ {
                                        var v19 string = v14
                                        v20 := GenS4{F0: 4294967295, F1: 18446744073709551615, F2: 127}
                                        v21 := GenS3{F0: 255, F1: 255, F2: 42}
                                        v14 = v19
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    }
                                    var v22 int64 = -9223372036854775808
                                    v23 := GenS3{F0: 2, F1: 128, F2: 2}
                                    _ = v22
                                    _ = v23
                                } else {
                                    v6[(fi16 * 127)] = v14
                                    if (42 == 65535) {
                                        var v24 uint32 = 65535
                                        _ = v24
                                    } else {
                                        v14 = v14
                                        switch (fi16 % (((fi16 / ((fi16) | 1))) | 1)) {
                                        case -83:
                                            fmt.Println(fi16, v14)
                                        }
                                        v25 := GenS5{F0: 0, F1: 596}
                                        _ = v25
                                    }
                                    var v26 int8 = -4
                                    _ = v26
                                }
                            case "\n\t":
                                var v27 int8 = -103
                                _ = v27
                            default:
                                fmt.Println(fi16, fi16)
                            }
                            var v28 float32 = float32(0.0)
                            _ = v28
                        }
                        if (-1 < -1) {
                            switch -2147483648 {
                            case 0:
                                for fi29 := 0; fi29 < 8; fi29++ {
                                    _ = len(v3)
                                    _ = fi29
                                }
                            case 128:
                                if (fi16 == (fi16 >> ((1) & 63))) {
                                    var v30 string = ""
                                    var v31 uint16 = 255
                                    _ = v30
                                    _ = v31
                                } else {
                                    for fi32 := 0; fi32 < 9; fi32++ {
                                        v14 = ("" + v14)
                                        if true {
                                            switch fi32 {
                                            case 127:
                                                v3[0] = (fi32 & 885)
                                            case 42:
                                                v14 = ("hello" + v14)
                                            default:
                                                if false {
                                                    var v33 uint8 = 1
                                                    _ = len(v2)
                                                    fmt.Println(fi16)
                                                    _ = v33
                                                } else {
                                                    fmt.Println(fi16)
                                                }
                                            }
                                            v14 = v14
                                        }
                                        _ = fi32
                                    }
                                    for fi34 := 0; fi34 < 4; fi34++ {
                                        v35 := GenS5{F0: 0, F1: 4294967295}
                                        if true {
                                            var v36 bool = (1 <= -32768)
                                            _ = v36
                                        } else {
                                            switch "foo bar" {
                                            case "hello":
                                                v14 = "a"
                                                var v37 uint32 = 1
                                                var v38 int16 = 1
                                                var v39 uint64 = 398
                                                var v40 uint64 = v39
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                            }
                                            v14 = v14
                                            v14 = v14
                                        }
                                        var v41 uint64 = 186
                                        _ = fi34
                                        _ = v35
                                        _ = v41
                                    }
                                }
                            default:
                                fmt.Println(v14)
                            }
                        } else {
                            if false {
                                if (385 <= -2147483648) {
                                    v14 = (("foo bar" + v14) + v14)
                                    delete(v3, 65535)
                                } else {
                                    var v42 int16 = -1
                                    _ = v42
                                }
                            } else {
                                v43 := GenS5{F0: 945, F1: 498}
                                var v44 int64 = 42
                                _ = v43
                                _ = v44
                            }
                        }
                        v14 = "\n\t"
                    }
                case "foo bar":
                    v45 := GenS6{F0: 19}
                    _ = v45
                }
                v14 = v14
            }
        }
        var v46 int64 = 9223372036854775807
        _ = fi16
        _ = v46
    }
    switch ("\n\t" + ("foo bar" + v14)) {
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "":
        v14 = v14
    case "hello":
        fmt.Println(v14)
    }
    var v47 bool = true
    var v48 int64 = 1
    v49 := GenS7{F0: 141, F1: 128}
    v1 = append(v1, -9223372036854775808)
    for fi50 := 0; fi50 < 11; fi50++ {
        var v51 int64 = v48
        if (1 >= 0) {
            if true {
                v14 = (v14 + (v14 + v14))
                var v52 float64 = 0.0
                _ = v52
            }
            if (v14 > v14) {
                var v53 uint64 = 778
                _ = v53
            } else {
                switch fi50 {
                case 32767:
                    var v54 uint16 = 625
                    _ = v54
                case 128:
                    switch fi50 {
                    case 351:
                        if (255 != 677) {
                            v14 = v14
                        }
                    case 0:
                        fmt.Println(fi50, v47)
                    case -1:
                        for fi55 := 0; fi55 < 4; fi55++ {
                            v56 := GenS8{F0: v51, F1: -2147483648, F2: ((v47 && (v47 || true)) && (!v47))}
                            fmt.Println(v14)
                            _ = fi55
                            _ = v56
                        }
                    }
                }
                fmt.Println(v48, v47, fi50)
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi50
        _ = v51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
