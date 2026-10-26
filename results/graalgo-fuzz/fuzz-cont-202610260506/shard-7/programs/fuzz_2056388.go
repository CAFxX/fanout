package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int64; F2 string }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 int64; F1 uint32 }

func main() {
    var v1 bool = (!(true && (true || true)))
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 0
    var v4 uint64 = 0
    v5 := make(map[string]string)
    if (-1 > -2147483648) {
        for fi6 := 0; fi6 < 9; fi6++ {
            v3 += v3
            var v7 uint = 4294967295
            _ = fi6
            _ = v7
        }
        _ = len(v5)
        var v8 uint32 = 2147483648
        _ = v8
    }
    if true {
        v9 := v5["a"]
        v10 := []int{-9223372036854775808, 1, 128}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v3 -= 65535
        _ = fi11
    }
    var v12 bool = (1e10 == 3.14159)
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 uint16 = uint16(v4)
            switch "" {
            case "":
                if (!v1) {
                    switch 0 {
                    case 0:
                        fmt.Println(v1, v3)
                    case 128:
                        var v16 uint64 = v4
                        _ = v16
                    }
                } else {
                    v15 -= 42
                    fmt.Println(fi13)
                }
            case "\n\t":
                v17 := GenS1{F0: 143}
                _ = v17
            case "a":
                var v18 uint = 451
                _ = v18
            default:
                v12 = (2 != 255)
            }
            fmt.Println(fi13, fi13, v4)
            switch fi14 {
            case -9223372036854775808:
                v5[""] = "a"
                if (!true) {
                    if false {
                        v1 = v1
                    }
                    var v19 uint8 = 146
                    if true {
                        v12 = ((true || true) || v12)
                    } else {
                        fmt.Println(v3, v1)
                    }
                    _ = v19
                } else {
                    v20 := []int{fi13, 9223372036854775807, -2147483648}
                    v21 := GenS1{F0: 116}
                    _ = v20
                    _ = v21
                }
            case 2147483647:
                v22 := []int{fi14}
                _ = v22
            }
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 float32 = -2.25
        v25 := GenS2{F0: -1}
        v12 = false
        _ = fi23
        _ = v24
        _ = v25
    }
    switch "hello" {
    case "foo bar":
        if (v1 && (false && v12)) {
            for fi26 := 0; fi26 < 1; fi26++ {
                v4 = v4
                for fi27 := 0; fi27 < 12; fi27++ {
                    v12 = (155 >= 625)
                    var v28 float32 = float32((0.1 - float64(1e10)))
                    v29 := []int{-531, 1, fi26, fi26}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                if (!((!v12) || (0 == 127))) {
                    if (65535 == 2) {
                        v30 := v2[65535]
                        _ = v30
                    }
                    v4 += v4
                } else {
                    fmt.Println(v4, v3)
                    var v31 int8 = -1
                    v32 := GenS3{F0: fi26, F1: 42, F2: "a"}
                    _ = v31
                    _ = v32
                }
                _ = fi26
            }
        }
    case "":
        switch -518 {
        case 9223372036854775807:
            for fi33 := 0; fi33 < 4; fi33++ {
                var v34 float32 = 3.14159
                _ = fi33
                _ = v34
            }
        case 42:
            v4 = v4
        default:
            for fi35 := 0; fi35 < 18; fi35++ {
                v36 := GenS3{F0: (-(fi35)), F1: 42, F2: "foo bar"}
                var v37 int64 = 1
                for fi38 := 0; fi38 < 9; fi38++ {
                    fmt.Println(fi35)
                    switch "a" {
                    case "a":
                        switch (fi35 + -22) {
                        case 127:
                            var v39 int = 0
                            _ = v39
                        case 2147483647:
                            if (true && (v12 && (fi38 == fi38))) {
                                v4 += v4
                                v40 := GenS1{F0: 42}
                                _ = v40
                            } else {
                                fmt.Println(v1)
                                v3 += 32767
                            }
                        }
                    case "\n\t":
                        v41 := GenS3{F0: fi38, F1: -9223372036854775808, F2: (("a" + "") + "foo bar")}
                        _ = v41
                    default:
                        v12 = ((-2.25 * (3.14159 + 0.0)) <= 0.1)
                    }
                    _ = fi38
                }
                _ = fi35
                _ = v36
                _ = v37
            }
        }
    default:
        if false {
            switch (("hello" + ("\n\t" + ("\n\t" + "\n\t"))) + ((((("\n\t" + "") + "a") + "a") + "a") + ("" + ("hello" + (((("\n\t" + "") + "hello") + "\n\t") + ""))))) {
            case "hello":
                fmt.Println(v1, v3)
            case "":
                var v42 int8 = -1
                _ = v42
            default:
                fmt.Println(v12)
            }
        } else {
            var v43 int64 = -587
            _ = v43
        }
    }
    var v44 uint64 = (v4 & v4)
    var v45 int32 = -37
    switch ("foo bar" + ("" + ("\n\t" + (("a" + "foo bar") + "foo bar")))) {
    case "a":
        v4 -= v4
    default:
        for fi46 := 0; fi46 < 10; fi46++ {
            v4 = v44
            v47 := []int{fi46, fi46, (fi46 % ((fi46) | 1)), fi46}
            delete(v2, 0)
            _ = fi46
            _ = v47
        }
    }
    var v48 uint = 4294967295
    v3 -= v3
    for fi49 := 0; fi49 < 16; fi49++ {
        fmt.Println(v45)
        if (v48 == (v48 | v48)) {
            var v50 int32 = -183
            var v51 uint64 = 4294967295
            _ = v50
            _ = v51
        }
        v52 := GenS4{F0: -723, F1: 0}
        _ = fi49
        _ = v52
    }
    v5[""] = "foo bar"
    v44 = uint64(v48)
    switch "" {
    case "a":
        v12 = v1
    case "\n\t":
        switch ("" + "a") {
        case "foo bar":
            v53 := []int{}
            _ = v53
        case "\n\t":
            if (!v12) {
                if (147 < 42) {
                    v54 := []int{int(v45), 2147483647, 9223372036854775807}
                    _ = v54
                }
            }
        case "a":
            v48 = (^(v48))
        }
    default:
        if false {
            v55 := GenS5{F0: 100, F1: "\n\t"}
            var v56 uint32 = 680
            _ = v55
            _ = v56
        } else {
            for fi57 := 0; fi57 < 1; fi57++ {
                fmt.Println(v48)
                if (!v12) {
                    v58 := []int{(fi57 & fi57)}
                    for fi59 := 0; fi59 < 19; fi59++ {
                        if false {
                            var v60 int64 = 84
                            v4 += (v4 >> ((18446744073709551615) & 63))
                            v61 := GenS6{F0: (v60 >> (((v48 ^ v48)) & 63)), F1: 2}
                            _ = v60
                            _ = v61
                        } else {
                            var v62 uint32 = 2147483647
                            _ = v62
                        }
                        _ = fi59
                    }
                    _ = v58
                }
                v44 -= v44
                _ = fi57
            }
            var v63 int16 = -986
            _ = v63
        }
    }
    fmt.Println(v12, v4)
    if false {
        for fi64 := 0; fi64 < 16; fi64++ {
            v65 := GenS4{F0: 0, F1: 1}
            _ = fi64
            _ = v65
        }
        var v66 uint64 = v4
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v48
    _ = v5
}
