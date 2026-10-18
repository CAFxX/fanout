package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 uint; F2 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 int16; F2 int16 }

func main() {
    v1 := []string{}
    v2 := []int{}
    var v3 int16 = 42
    v4 := make(map[string]bool)
    var v5 int = 118
    var v6 float64 = 3.14159
    v6 = v6
    v7 := GenS1{F0: 0, F1: 1, F2: -2147483648}
    fmt.Println(v6, v6, v5)
    v5 += v5
    v8 := GenS2{F0: 4294967295, F1: -420, F2: v5}
    var v9 int8 = 100
    var v10 int8 = (v9 - (v9 >> ((2) & 63)))
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(fi11, v6, v5)
        v10 += (v10 + v10)
        _ = fi11
    }
    if (!((1e10 != v6) && ((2 >= 675) && (228 != 2)))) {
        var v12 uint16 = 667
        if (!true) {
            fmt.Println(v9)
        }
        if (true || ((!((4294967295 >= 543) && ((v9 | v9) <= 42))) && true)) {
            switch "\n\t" {
            case "foo bar":
                switch "\n\t" {
                case "a":
                    fmt.Println(v5, v12)
                }
            case "":
                v13 := v4["a"]
                _ = v13
            case "\n\t":
                var v14 int8 = v10
                _ = v14
            }
            if false {
                if len(v2) > 0 {
                    v2[0] = v5
                }
            } else {
                if (1.5 != float64((0.1 / 0.1))) {
                    var v15 int = (v5 >> ((424) & 63))
                    if true {
                        for fi16 := 0; fi16 < 18; fi16++ {
                            switch (fi16 | (fi16 / (((fi16 * fi16)) | 1))) {
                            case -726:
                                if true {
                                    if false {
                                        for fi17 := 0; fi17 < 11; fi17++ {
                                            if true {
                                                switch v15 {
                                                case -1:
                                                    var v18 uint16 = 1
                                                    _ = v18
                                                case 4294967295:
                                                    var v19 string = ""
                                                    switch -145 {
                                                    case -1:
                                                        var v20 uint64 = 522
                                                        _ = v20
                                                    }
                                                    fmt.Println(v3, fi16)
                                                    _ = v19
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = "hello"
                                                    }
                                                }
                                                v12 -= (v12 >> ((1) & 63))
                                            }
                                            var v21 int8 = -1
                                            fmt.Println(v15, fi16, v15)
                                            switch 1 {
                                            case 108:
                                                fmt.Println(v12, v6)
                                                var v22 float64 = 0.0
                                                _ = v22
                                            default:
                                                if true {
                                                    var v23 uint64 = 42
                                                    v24 := GenS3{F0: "foo bar"}
                                                    v4[""] = false
                                                    _ = v23
                                                    _ = v24
                                                }
                                            }
                                            if ((v12 << ((756) & 63)) == v12) {
                                                if true {
                                                    switch (-(fi16)) {
                                                    case 32768:
                                                        var v25 int32 = -2147483648
                                                        _ = v25
                                                    case -9223372036854775808:
                                                        v26 := v4["a"]
                                                        _ = v26
                                                    case -916:
                                                        v3 += int16(v12)
                                                    }
                                                    switch ("a" + "") {
                                                    case "":
                                                        switch v5 {
                                                        case -982:
                                                            v1 = append(v1, "a")
                                                        }
                                                    case "\n\t":
                                                        if false {
                                                            v15 = (-(v15))
                                                            var v27 int32 = 0
                                                            _ = v27
                                                        }
                                                    case "foo bar":
                                                        if (v3 != (v3 >> ((65535) & 63))) {
                                                            var v28 uint32 = 390
                                                            v5 += fi17
                                                            if true {
                                                                v29 := GenS1{F0: 440, F1: v28, F2: -525}
                                                                var v30 uint64 = 18446744073709551615
                                                                var v31 uint64 = (v30 ^ uint64(v12))
                                                                _ = v29
                                                                _ = v30
                                                                _ = v31
                                                            } else {
                                                                _ = len(v2)
                                                            }
                                                            _ = v28
                                                        } else {
                                                            fmt.Println(v21, fi16)
                                                            v5 = (v5 - v15)
                                                            if (30 != 128) {
                                                                var v32 int32 = 0
                                                                v33 := GenS3{F0: "foo bar"}
                                                                _ = v32
                                                                _ = v33
                                                            }
                                                            _ = len(v4)
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v34 := GenS4{F0: (v3 & v3)}
                                                                _ = v34
                                                            default:
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    delete(v4, "foo bar")
                                                                case "hello":
                                                                    fmt.Println(v10)
                                                                case "foo bar":
                                                                    v15 -= (v5 * v5)
                                                                default:
                                                                    if true {
                                                                        v35 := GenS5{F0: false}
                                                                        fmt.Println(v6, v6)
                                                                        v36 := GenS1{F0: 0, F1: 65535, F2: 0}
                                                                        var v37 int64 = -741
                                                                        if (true && (4294967295 != 1)) {
                                                                            fmt.Println(v9, v5)
                                                                            v38 := GenS6{F0: false, F1: 255, F2: 1}
                                                                            _ = v38
                                                                        }
                                                                        _ = v35
                                                                        _ = v36
                                                                        _ = v37
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        var v39 bool = true
                                                        if true {
                                                            v21 += v21
                                                            v6 += v6
                                                            var v40 float64 = 0.1
                                                            v21 = (v10 % (((v10 % ((v9) | 1))) | 1))
                                                            fmt.Println(v10, v5)
                                                            _ = v40
                                                        }
                                                        _ = v39
                                                    }
                                                } else {
                                                    v9 -= v10
                                                    v41 := GenS6{F0: true, F1: 18446744073709551615, F2: 64}
                                                    var v42 uint = 255
                                                    _ = v41
                                                    _ = v42
                                                }
                                                v6 = (v6 - float64(float32(-2.25)))
                                            }
                                            _ = fi17
                                            _ = v21
                                        }
                                        var v43 int32 = 1
                                        _ = v43
                                    }
                                    for fi44 := 0; fi44 < 12; fi44++ {
                                        v45 := GenS7{F0: 32768}
                                        var v46 float32 = 0.0
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                    }
                                }
                            }
                            _ = fi16
                        }
                    } else {
                        _ = len(v1)
                    }
                    v10 = (v10 + v9)
                    _ = v15
                }
                v47 := GenS6{F0: ((v12 % ((v12) | 1)) > (^(v12))), F1: 18446744073709551615, F2: 0}
                _ = v47
            }
            v48 := GenS8{F0: 128, F1: v3, F2: int16(v5)}
            _ = v48
        } else {
            v49 := GenS7{F0: v5}
            var v50 int32 = 237
            var v51 float64 = 1e10
            _ = v49
            _ = v50
            _ = v51
        }
        _ = v12
    } else {
        v10 = v9
        var v52 uint = 65535
        var v53 float64 = v6
        _ = v52
        _ = v53
    }
    var v54 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
