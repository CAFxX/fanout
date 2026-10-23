package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint32; F1 int8; F2 int16 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int64; F1 int; F2 uint16 }
type GenS6 struct { F0 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := []uint32{274, 584}
    var v4 uint = 617
    v5 := []string{"a"}
    var v6 float32 = (0.1 - 3.14159)
    var v7 float64 = ((1e10 + (-2.25 - 1e10)) / 3.14159)
    for fi8 := 0; fi8 < 9; fi8++ {
        if (("\n\t" > "foo bar") && false) {
            v6 = 1.5
            switch "hello" {
            case "a":
                v9 := GenS1{F0: true}
                _ = v9
            case "hello":
                v6 += 1e10
            case "\n\t":
                if (1 != 426) {
                    var v10 int64 = -9223372036854775808
                    switch 42 {
                    case 564:
                        if ((v4 | v4) >= v4) {
                            v4 = (^(v4))
                            v11 := GenS2{F0: 0, F1: v4}
                            _ = v11
                        } else {
                            var v12 uint32 = 4294967295
                            _ = v12
                        }
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v4 = v4
                    _ = v10
                }
            default:
                for fi13 := 0; fi13 < 17; fi13++ {
                    switch ("hello" + "\n\t") {
                    case "a":
                        var v14 int8 = -24
                        _ = v14
                    case "foo bar":
                        var v15 float64 = (v7 - v7)
                        _ = v15
                    case "hello":
                        delete(v2, ("" + ""))
                    default:
                        fmt.Println(fi8, v4, v4)
                    }
                    _ = fi13
                }
            }
            var v16 int = 32767
            _ = v16
        } else {
            if false {
                fmt.Println(fi8, fi8, fi8)
                _ = len(v3)
            } else {
                if true {
                    v17 := GenS3{F0: 0, F1: -128, F2: 0}
                    fmt.Println(v7)
                    v4 = v4
                    _ = v17
                } else {
                    v18 := GenS3{F0: 65535, F1: 67, F2: 1}
                    var v19 int32 = 42
                    _ = v18
                    _ = v19
                }
                delete(v2, "a")
            }
            v6 -= (-2.25 + (v6 / v6))
            v20 := GenS1{F0: true}
            _ = v20
        }
        _ = fi8
    }
    var v21 uint8 = 1
    for fi22 := 0; fi22 < 16; fi22++ {
        switch "" {
        case "foo bar":
            if true {
                v6 -= v6
            } else {
                v23 := v2["hello"]
                fmt.Println(v6, v6)
                v23 = true
                _ = v23
            }
        default:
            var v24 uint8 = 140
            _ = v24
        }
        switch -378 {
        case 127:
            for fi25 := 0; fi25 < 15; fi25++ {
                v4 = v4
                if len(v5) > 0 {
                    v5[0] = ""
                }
                _ = fi25
            }
        case 42:
            fmt.Println(v4)
        case -366:
            var v26 int8 = 100
            _ = v26
        }
        if false {
            var v27 uint = v4
            _ = v27
        }
        _ = fi22
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        v29 := GenS4{F0: v21, F1: "foo bar"}
        v30 := GenS4{F0: 127, F1: ("foo bar" + (("foo bar" + "\n\t") + "foo bar"))}
        v21 += v21
        _ = fi28
        _ = v29
        _ = v30
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        fmt.Println(v21, v6, v4)
    default:
        _ = len(v1)
    }
    if false {
        var v31 bool = true
        fmt.Println(v6, v6)
        _ = v31
    } else {
        _ = len(v5)
        v21 = (v21 - v21)
    }
    var v32 int64 = 332
    var v33 int16 = 814
    fmt.Println(v33)
    if (4294967295 == 18446744073709551615) {
        v4 -= v4
        switch 694 {
        case 945:
            var v34 int32 = 2147483647
            _ = v34
        case 128:
            var v35 int8 = -1
            _ = v35
        default:
            if (!true) {
                v6 = v6
                var v36 float32 = v6
                var v37 uint16 = 2
                _ = v36
                _ = v37
            }
        }
        v38 := GenS5{F0: (v32 / (((^(v32))) | 1)), F1: 9223372036854775807, F2: 1}
        _ = v38
    }
    if (0 <= 0) {
        fmt.Println(v7)
        switch "" {
        case "\n\t":
            v39 := GenS5{F0: 725, F1: -1, F2: 42}
            _ = v39
        }
    } else {
        if false {
            var v40 int8 = 1
            v41 := GenS6{F0: 42}
            _ = v40
            _ = v41
        } else {
            v33 -= (-(v33))
        }
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        fmt.Println(v32, v33)
        fmt.Println(v33)
        v43 := GenS1{F0: (192 > 2)}
        _ = fi42
        _ = v43
    }
    for fi44 := 0; fi44 < 20; fi44++ {
        var v45 uint16 = 1
        switch "" {
        case "foo bar":
            if false {
                fmt.Println(v6, fi44, fi44)
            }
        case "":
            var v46 float64 = 1.5
            _ = v46
        default:
            switch (fi44 << (((^(v4))) & 63)) {
            case -2147483648:
                var v47 float64 = 0.1
                _ = v47
            }
        }
        var v48 string = "\n\t"
        _ = fi44
        _ = v45
        _ = v48
    }
    var v49 int64 = v32
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
