package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 bool; F2 string }
type GenS3 struct { F0 int64; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 0
    v3 := []uint32{488, 907, 65535, 911}
    var v4 uint64 = 42
    v5 := GenS1{F0: 5}
    _ = len(v1)
    switch -736 {
    case 32767:
        if false {
            switch 2147483647 {
            case 1:
                var v6 uint32 = 1
                _ = v6
            case 4294967295:
                if (3.14159 == 1e10) {
                    delete(v1, "foo bar")
                    for fi7 := 0; fi7 < 5; fi7++ {
                        var v8 float32 = 0.1
                        _ = fi7
                        _ = v8
                    }
                    if ((-2.25 * 3.14159) <= (1e10 - (float32((float64(((3.14159 + (float32(1e10) * -2.25)) - 0.0)) * ((-2.25 + 0.0) - 1.5))) / float32(0.1)))) {
                        var v9 float32 = 1e10
                        _ = v9
                    } else {
                        for fi10 := 0; fi10 < 3; fi10++ {
                            v11 := GenS1{F0: -9223372036854775808}
                            _ = fi10
                            _ = v11
                        }
                        v12 := GenS1{F0: 0}
                        v2 = 255
                        _ = v12
                    }
                } else {
                    for fi13 := 0; fi13 < 19; fi13++ {
                        if (1 != 923) {
                            v14 := GenS1{F0: 42}
                            _ = v14
                        } else {
                            var v15 int64 = 549
                            var v16 float64 = (0.0 - float64((float32(3.14159) * -2.25)))
                            _ = v15
                            _ = v16
                        }
                        _ = fi13
                    }
                    for fi17 := 0; fi17 < 2; fi17++ {
                        v18 := v1["foo bar"]
                        switch v18 {
                        case 42:
                            v4 = 1
                        case 166:
                            if (1 <= -245) {
                                var v19 float64 = (1e10 * 3.14159)
                                for fi20 := 0; fi20 < 14; fi20++ {
                                    var v21 float64 = v19
                                    switch "hello" {
                                    case "hello":
                                        v22 := GenS1{F0: 42}
                                        _ = v22
                                    case "a":
                                        _ = len(v1)
                                    }
                                    var v23 int32 = 1
                                    _ = fi20
                                    _ = v21
                                    _ = v23
                                }
                                var v24 bool = false
                                _ = v19
                                _ = v24
                            } else {
                                switch (v18 % ((-746) | 1)) {
                                case 1:
                                    v4 = v4
                                default:
                                    for fi25 := 0; fi25 < 8; fi25++ {
                                        switch "a" {
                                        case "\n\t":
                                            switch ("hello" + ("a" + ("" + ("foo bar" + ((("" + "\n\t") + "foo bar") + "foo bar"))))) {
                                            case "":
                                                v26 := GenS2{F0: v4, F1: true, F2: "foo bar"}
                                                _ = v26
                                            case "hello":
                                                var v27 bool = (true || ((0 >= 2147483647) || ((^(v2)) > v2)))
                                                _ = v27
                                            case "a":
                                                fmt.Println(v2, fi17)
                                            default:
                                                var v28 float64 = (1e10 - (-2.25 - 3.14159))
                                                _ = v28
                                            }
                                        default:
                                            _ = len(v1)
                                        }
                                        if true {
                                            switch fi17 {
                                            case 128:
                                                switch "a" {
                                                case "a":
                                                    _ = len(v1)
                                                    v29 := GenS1{F0: 9223372036854775807}
                                                    _ = v29
                                                case "hello":
                                                    v30 := GenS1{F0: 1}
                                                    _ = v30
                                                }
                                            case -628:
                                                v4 = (v4 ^ (^(v4)))
                                            }
                                        } else {
                                            var v31 string = "\n\t"
                                            v1[v31] = fi17
                                            v32 := GenS2{F0: v4, F1: ((1e10 / 1e10) > (float64(0.1) - 1e10)), F2: v31}
                                            _ = v31
                                            _ = v32
                                        }
                                        fmt.Println(fi25, fi25, v2)
                                        _ = fi25
                                    }
                                }
                            }
                        case 32768:
                            v33 := GenS3{F0: 0, F1: (100 >= 52), F2: 4294967295}
                            _ = v33
                        }
                        _ = fi17
                        _ = v18
                    }
                }
            case 479:
                v2 += v2
            }
            v34 := GenS3{F0: 42, F1: true, F2: 42}
            if (4 != 1) {
                fmt.Println(v2, v4)
                v35 := GenS4{F0: 1}
                _ = v35
            } else {
                v4 -= v4
                v2 -= (v2 * (v2 % ((255) | 1)))
                if (-128 != int8(v4)) {
                    switch -2147483648 {
                    case 32767:
                        v2 += (v2 >> ((v2) & 63))
                    case 128:
                        v2 -= v2
                    default:
                        v4 = 2
                    }
                    var v36 int64 = -305
                    _ = v36
                }
            }
            _ = v34
        } else {
            fmt.Println(v2, v2)
        }
    }
    var v37 int32 = 0
    for fi38 := 0; fi38 < 8; fi38++ {
        if (0.1 < (1.5 - -2.25)) {
            v4 = v4
            fmt.Println(v4)
        }
        v2 = v2
        var v39 int16 = -32768
        _ = fi38
        _ = v39
    }
    var v40 uint64 = v4
    var v41 uint16 = 32767
    var v42 int16 = 1
    v43 := GenS5{F0: -1}
    var v44 uint64 = (^(v40))
    for fi45 := 0; fi45 < 18; fi45++ {
        v46 := GenS2{F0: 2, F1: (-1 != 1), F2: "a"}
        v37 += (v37 >> (((v2 | v2)) & 63))
        _ = fi45
        _ = v46
    }
    switch "\n\t" {
    case "hello":
        var v47 uint = v2
        _ = v47
    }
    for fi48 := 0; fi48 < 8; fi48++ {
        v49 := GenS3{F0: 1, F1: false, F2: 141}
        _ = fi48
        _ = v49
    }
    if (v42 > (v42 << (((v2 / ((42) | 1))) & 63))) {
        switch "foo bar" {
        case "a":
            v3 = append(v3, 42)
        case "foo bar":
            var v50 uint32 = 65535
            _ = v50
        }
    } else {
        v51 := GenS5{F0: 0}
        fmt.Println(v44, v40)
        _ = v51
    }
    v52 := GenS1{F0: -57}
    var v53 string = ("\n\t" + "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v52
    _ = v53
}
