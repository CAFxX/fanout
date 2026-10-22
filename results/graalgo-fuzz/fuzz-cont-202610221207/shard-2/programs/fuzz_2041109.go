package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 int64; F1 int; F2 uint64 }
func genFn0(p0 uint8) uint64 {
    if ((2 >= 42) || false) {
        p0 = p0
        if false {
            for fi4 := 0; fi4 < 8; fi4++ {
                v5 := GenS1{F0: 42, F1: 127, F2: -11}
                switch ("hello" + "") {
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        if false {
                            switch (fi4 / ((fi4) | 1)) {
                            case 32768:
                                if (127 == 48) {
                                    var v6 int64 = 0
                                    for fi7 := 0; fi7 < 5; fi7++ {
                                        p0 = p0
                                        var v8 string = ""
                                        switch (fi4 ^ (-(fi4))) {
                                        case -128:
                                            switch (v8 + ("" + "\n\t")) {
                                            case "":
                                                v8 = (v8 + (("foo bar" + ("hello" + "")) + v8))
                                            }
                                        }
                                        _ = fi7
                                        _ = v8
                                    }
                                    _ = v6
                                }
                            case 9223372036854775807:
                                switch fi4 {
                                case 895:
                                    v9 := []int{1, (^(fi4)), fi4, -2147483648}
                                    _ = v9
                                default:
                                    var v10 uint16 = 42
                                    _ = v10
                                }
                            case 587:
                                var v11 bool = true
                                _ = v11
                            }
                        }
                    default:
                        var v12 uint8 = (p0 + 35)
                        _ = v12
                    }
                }
                _ = fi4
                _ = v5
            }
            var v13 uint16 = 1
            v14 := []int{42, 42}
            _ = v13
            _ = v14
        }
    } else {
        v15 := []int{127, 1}
        v16 := GenS1{F0: 4294967295, F1: (p0 + 156), F2: 1}
        _ = v15
        _ = v16
    }
    switch "" {
    case "\n\t":
        var v17 bool = false
        _ = v17
    default:
        p0 += 25
    }
    v18 := GenS1{F0: 886, F1: 0, F2: 9}
    v19 := GenS1{F0: 465, F1: (p0 + (p0 & p0)), F2: 127}
    _ = p0
    _ = v18
    _ = v19
    return uint64(p0)
}
func genFn1() string {
    v20 := []int{42, 2147483647, 9223372036854775807}
    if false {
        var v21 float64 = (0.0 - 3.14159)
        if ((false && true) && (!false)) {
            if false {
                v22 := GenS2{F0: 108, F1: 4294967295}
                _ = v22
            } else {
                for fi23 := 0; fi23 < 2; fi23++ {
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    _ = fi23
                }
            }
        } else {
            v21 += (((v21 + (v21 + (v21 / v21))) - ((0.0 - v21) / v21)) - float64((float32(((1.5 + (1e10 * v21)) / 1e10)) + 1.5)))
            v21 -= v21
            switch -123 {
            case 572:
                v21 = v21
            case 127:
                var v24 int8 = 100
                _ = v24
            case 0:
                fmt.Println(v21)
            }
        }
        _ = v21
    } else {
        var v25 bool = (!(!(!((42 >= 836) || false))))
        var v26 uint16 = 65535
        _ = v25
        _ = v26
    }
    if (true && (((88 != 170) || (!(!(!false)))) && (((-2.25 != -2.25) && false) && (1 >= -32768)))) {
        v20 = append(v20, 42)
    }
    if (("hello" < "\n\t") || false) {
        v27 := GenS3{F0: 78, F1: -1}
        var v28 float32 = 3.14159
        _ = v27
        _ = v28
    }
    var v29 int8 = -1
    _ = v20
    _ = v29
    return "a"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    var v3 int8 = -128
    v3 = v3
    fmt.Println(v3, v2)
    fmt.Println(v2, v3)
    for fi30 := 0; fi30 < 17; fi30++ {
        v3 += (v3 >> ((42) & 63))
        if v2 {
            for fi31 := 0; fi31 < 20; fi31++ {
                switch (fi31 * 4294967295) {
                case 32768:
                    fmt.Println(v2, v3, v3)
                case 332:
                    v32 := GenS4{F0: ("" + "foo bar")}
                    _ = v32
                }
                _ = fi31
            }
            switch "hello" {
            case "a":
                var v33 float32 = (((((-2.25 / 0.1) / 1.0) - 0.1) * (1.5 - 3.14159)) / 1e10)
                _ = v33
            case "foo bar":
                for fi34 := 0; fi34 < 5; fi34++ {
                    v3 += v3
                    _ = fi34
                }
            default:
                switch ("a" + ("foo bar" + "\n\t")) {
                case "\n\t":
                    for fi35 := 0; fi35 < 11; fi35++ {
                        v36 := GenS1{F0: 161, F1: 128, F2: (v3 << ((65535) & 63))}
                        var v37 float32 = (0.1 + 3.14159)
                        v2 = (true && false)
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                case "":
                    v3 = (v3 / ((-128) | 1))
                case "foo bar":
                    v38 := GenS4{F0: (("" + ("a" + "")) + ("hello" + "a"))}
                    _ = v38
                default:
                    v39 := GenS3{F0: 966, F1: -9223372036854775808}
                    _ = v39
                }
            }
            var v40 float32 = (0.0 / 1.0)
            _ = v40
        } else {
            for fi41 := 0; fi41 < 1; fi41++ {
                fmt.Println(fi41)
                _ = fi41
            }
            for fi42 := 0; fi42 < 12; fi42++ {
                if (((0.0 + 0.1) * float32(float64(-2.25))) > 0.0) {
                    if v2 {
                        v2 = false
                    }
                    switch -147 {
                    case 127:
                        delete(v1, 42)
                    case -2147483648:
                        var v43 int = fi30
                        _ = v43
                    case -174:
                        v2 = ("" >= "")
                    default:
                        v2 = true
                    }
                    v44 := GenS5{F0: 366}
                    _ = v44
                }
                fmt.Println(fi42, fi30, v3)
                var v45 uint64 = 18446744073709551615
                v46 := GenS6{F0: "foo bar", F1: 0}
                var v47 float64 = -2.25
                _ = fi42
                _ = v45
                _ = v46
                _ = v47
            }
            fmt.Println(v2, fi30, v2)
        }
        _ = fi30
    }
    switch ("\n\t" + "hello") {
    case "":
        v3 += (v3 * 42)
    case "foo bar":
        v48 := GenS2{F0: int16(v3), F1: -310}
        _ = v48
    case "hello":
        fmt.Println(v3, v3)
    }
    if false {
        v2 = true
        v49 := GenS7{F0: 1, F1: 0, F2: genFn0(213)}
        _ = v49
    }
    for fi50 := 0; fi50 < 10; fi50++ {
        fmt.Println(v2)
        var v51 uint16 = 2
        if ((((!(42 != 2)) && v2) && false) || (v2 || (603 <= 1))) {
            var v52 uint = 42
            var v53 int = fi50
            _ = v52
            _ = v53
        }
        _ = fi50
        _ = v51
    }
    var v54 int = -1
    switch (v54 * v54) {
    case -2147483648:
        v55 := GenS6{F0: "a", F1: 42}
        _ = v55
    default:
        switch ((("" + "foo bar") + "hello") + "a") {
        case "\n\t":
            v56 := v1[49]
            _ = v56
        }
    }
    var v57 uint = 1
    v54 = v54
    var v58 int8 = v3
    for fi59 := 0; fi59 < 15; fi59++ {
        v60 := GenS1{F0: 1, F1: 213, F2: v3}
        var v61 uint64 = 42
        _ = fi59
        _ = v60
        _ = v61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v57
    _ = v58
}
