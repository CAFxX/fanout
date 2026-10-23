package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 string }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    v3 := make(map[string]uint32)
    var v4 uint32 = 657
    var v5 float32 = 1.5
    var v6 int64 = 42
    v5 = float32(((1.5 * ((3.14159 / -2.25) / 1e10)) / 1e10))
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    var v7 int64 = v6
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int64 = (-(v7))
        var v10 string = "a"
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 255, F1: (v6 >> ((422) & 63))}
    fmt.Println(v5, v4)
    var v12 uint64 = 4294967295
    switch "foo bar" {
    case "foo bar":
        switch "a" {
        case "foo bar":
            var v13 int64 = 42
            _ = v13
        default:
            v14 := GenS1{F0: 68, F1: v6}
            _ = v14
        }
    }
    switch 655 {
    case 2147483647:
        v7 = v6
    case 1:
        if (2 != 1) {
            v2 = append(v2, 128)
        } else {
            for fi15 := 0; fi15 < 9; fi15++ {
                v16 := GenS1{F0: 42, F1: (v6 << ((65535) & 63))}
                delete(v3, ("foo bar" + ("hello" + (("\n\t" + "\n\t") + ("hello" + "a")))))
                v17 := GenS1{F0: 128, F1: 0}
                _ = fi15
                _ = v16
                _ = v17
            }
            fmt.Println(v5, v5, v6)
            v18 := v1["a"]
            _ = v18
        }
    case 32768:
        v19 := GenS1{F0: 192, F1: v6}
        _ = v19
    default:
        if (((((("\n\t" + ("a" + "hello")) + "") + ("" + ("\n\t" + "hello"))) + ("hello" + ("hello" + ""))) > "\n\t") || (0.0 != v5)) {
            var v20 float64 = (0.1 + 0.0)
            if true {
                var v21 int = -122
                var v22 uint16 = 483
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 19; fi23++ {
                    var v24 uint = 0
                    var v25 int8 = 0
                    switch (fi23 | -2147483648) {
                    case 449:
                        fmt.Println(v5)
                    case 4294967295:
                        var v26 int16 = 0
                        _ = v26
                    default:
                        switch fi23 {
                        case 127:
                            if true {
                                v27 := GenS1{F0: 0, F1: int64(v4)}
                                for fi28 := 0; fi28 < 4; fi28++ {
                                    v24 = v24
                                    _ = fi28
                                }
                                _ = v27
                            }
                        case 4294967295:
                            for fi29 := 0; fi29 < 17; fi29++ {
                                var v30 uint16 = 32767
                                switch fi23 {
                                case 32767:
                                    v25 += (v25 | 1)
                                case 32768:
                                    v7 += (v7 * v7)
                                }
                                v24 -= 4294967295
                                _ = fi29
                                _ = v30
                            }
                        case 431:
                            v5 -= -2.25
                        default:
                            var v31 uint16 = 878
                            _ = v31
                        }
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                for fi32 := 0; fi32 < 7; fi32++ {
                    var v33 int64 = v7
                    _ = fi32
                    _ = v33
                }
                switch 127 {
                case 4294967295:
                    v34 := GenS2{F0: -1, F1: 307, F2: "a"}
                    _ = v34
                case -2147483648:
                    for fi35 := 0; fi35 < 20; fi35++ {
                        if (0 < 0) {
                            if (!(2147483648 != 672)) {
                                if ((!(false || (v20 >= (3.14159 / float64((-2.25 - 3.14159)))))) || ((fi35 == (fi35 - 0)) && ((^(v12)) >= (v12 + (v12 << ((150) & 63)))))) {
                                    var v36 uint16 = 225
                                    _ = v36
                                }
                                var v37 int32 = -2147483648
                                _ = v37
                            } else {
                                v38 := GenS1{F0: uint8(v4), F1: v6}
                                _ = v38
                            }
                            v6 -= v6
                        }
                        for fi39 := 0; fi39 < 11; fi39++ {
                            switch ("" + ("a" + "foo bar")) {
                            case "a":
                                var v40 bool = (407 > 0)
                                _ = v40
                            case "foo bar":
                                v41 := v1[""]
                                _ = v41
                            }
                            if ((^(v12)) < (v12 / (((v12 | (^(v12)))) | 1))) {
                                v42 := GenS2{F0: (fi39 - 9223372036854775807), F1: 930, F2: (("" + ("\n\t" + (("hello" + "\n\t") + "\n\t"))) + "a")}
                                if false {
                                    switch "hello" {
                                    case "hello":
                                        v4 = (v4 * 42)
                                    case "a":
                                        var v43 uint64 = 877
                                        _ = v43
                                    case "":
                                        if false {
                                            if len(v2) > 0 {
                                                v2[0] = (fi35 | fi39)
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = 0
                                            }
                                        }
                                    default:
                                        fmt.Println(v12)
                                    }
                                    var v44 int8 = 7
                                    _ = v44
                                }
                                var v45 int8 = -1
                                _ = v42
                                _ = v45
                            } else {
                                var v46 uint64 = 799
                                if len(v2) > 0 {
                                    v2[0] = (fi39 / ((fi35) | 1))
                                }
                                _ = v46
                            }
                            var v47 uint32 = (v4 % (((v4 ^ uint32(v12))) | 1))
                            _ = fi39
                            _ = v47
                        }
                        _ = fi35
                    }
                case 2147483647:
                    var v48 float32 = (((v5 * ((float32(v20) / (v5 + float32(v20))) + v5)) - v5) - (v5 + (v5 * (v5 - v5))))
                    _ = v48
                default:
                    v1["\n\t"] = 42
                }
            }
            var v49 float32 = v5
            _ = v20
            _ = v49
        }
    }
    fmt.Println(v7, v7)
    for fi50 := 0; fi50 < 12; fi50++ {
        fmt.Println(fi50, v12)
        v6 -= int64(fi50)
        switch (fi50 & (^(fi50))) {
        case 1:
            v51 := GenS3{F0: "\n\t"}
            _ = v51
        case 42:
            fmt.Println(fi50, v6)
        default:
            if false {
                var v52 int64 = (v6 << ((42) & 63))
                for fi53 := 0; fi53 < 19; fi53++ {
                    fmt.Println(fi53)
                    var v54 uint32 = (v4 >> ((255) & 63))
                    fmt.Println(v12)
                    _ = fi53
                    _ = v54
                }
                v7 = 1
                _ = v52
            } else {
                v12 = v12
            }
        }
        _ = fi50
    }
    v55 := GenS1{F0: 56, F1: (v6 << ((4294967295) & 63))}
    fmt.Println(v7, v4, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v7
}
