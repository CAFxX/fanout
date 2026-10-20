package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 uint; F2 uint64 }
type GenS8 struct { F0 uint16; F1 uint; F2 int }
type GenS9 struct { F0 bool }

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: (("\n\t" + ("hello" + ("a" + ("" + "\n\t")))) + ("hello" + "\n\t")), F1: 1000}
    fmt.Println(v1)
    v1 = v1
    for fi5 := 0; fi5 < 14; fi5++ {
        if true {
            switch (v1 + 127) {
            case 9223372036854775807:
                v6 := GenS2{F0: 447}
                _ = v6
            default:
                v1 -= (v1 * (v1 >> ((1) & 63)))
            }
            v3[0] = "\n\t"
            v7 := GenS1{F0: ("a" + ("a" + ("" + (("foo bar" + "foo bar") + "a")))), F1: -32768}
            _ = v7
        } else {
            var v8 bool = false
            _ = v8
        }
        _ = fi5
    }
    if false {
        var v9 int = v1
        _ = v9
    } else {
        if (((v1 << ((255) & 63)) > 804) || false) {
            if (!(2 > 0)) {
                if (0 > 255) {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        switch ("\n\t" + ("foo bar" + ("hello" + "a"))) {
                        case "foo bar":
                            for fi11 := 0; fi11 < 2; fi11++ {
                                if true {
                                    switch "foo bar" {
                                    case "foo bar":
                                        switch (v1 - fi10) {
                                        case -9223372036854775808:
                                            v12 := GenS2{F0: 42}
                                            _ = v12
                                        case 861:
                                            v1 += (v1 % (((v1 + v1)) | 1))
                                        case 9223372036854775807:
                                            var v13 bool = false
                                            _ = v13
                                        }
                                    case "hello":
                                        var v14 bool = ("a" != "foo bar")
                                        _ = v14
                                    case "\n\t":
                                        v15 := GenS3{F0: (65535 != 843), F1: 42}
                                        _ = v15
                                    default:
                                        switch (fi10 | (fi10 * 1)) {
                                        case 32768:
                                            v1 -= v1
                                            v1 -= (-(v1))
                                            switch "hello" {
                                            case "":
                                                if false {
                                                    switch v1 {
                                                    case 0:
                                                        v1 = fi11
                                                    case 32768:
                                                        v16 := GenS4{F0: 32767, F1: 745, F2: 18446744073709551615}
                                                        _ = v16
                                                    case 32767:
                                                        v1 += (-(fi11))
                                                    default:
                                                        v17 := []int{1, fi11, fi11}
                                                        _ = v17
                                                    }
                                                    if false {
                                                        v18 := GenS1{F0: "a", F1: 0}
                                                        v1 -= 42
                                                        _ = v18
                                                    } else {
                                                        if (255 > 42) {
                                                            var v19 uint = 4294967295
                                                            _ = v19
                                                        } else {
                                                            v20 := GenS5{F0: v1, F1: 781}
                                                            _ = v20
                                                        }
                                                        if (3.14159 < (((3.14159 - 3.14159) - float64(1e10)) - 1.5)) {
                                                            switch v1 {
                                                            case -328:
                                                                var v21 bool = (-2.25 <= (1e10 - 0.1))
                                                                _ = v21
                                                            }
                                                            switch (^(fi11)) {
                                                            case 2147483647:
                                                                v1 = fi11
                                                            default:
                                                                var v22 uint = 917
                                                                _ = v22
                                                            }
                                                            v1 += -376
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v1, fi10, fi10)
                                                    v1 -= (v1 * (fi10 * -815))
                                                }
                                            case "foo bar":
                                                if false {
                                                    var v23 int64 = -1
                                                    v24 := []int{18, 777}
                                                    v23 = (-(v23))
                                                    _ = v23
                                                    _ = v24
                                                }
                                            case "a":
                                                v25 := GenS1{F0: "\n\t", F1: 356}
                                                _ = v25
                                            }
                                        case -1:
                                            var v26 int32 = 0
                                            _ = v26
                                        }
                                    }
                                    switch (fi11 + (fi11 % ((v1) | 1))) {
                                    case 2147483647:
                                        var v27 int32 = 1
                                        _ = v27
                                    default:
                                        v28 := GenS6{F0: ""}
                                        _ = v28
                                    }
                                }
                                if false {
                                    fmt.Println(v1, v1, fi10)
                                } else {
                                    var v29 uint32 = 2
                                    _ = v29
                                }
                                _ = fi11
                            }
                        case "hello":
                            for fi30 := 0; fi30 < 4; fi30++ {
                                var v31 uint8 = 127
                                if false {
                                    v32 := GenS7{F0: -9223372036854775808, F1: 597, F2: 0}
                                    _ = v32
                                }
                                fmt.Println(v31, v1, fi10)
                                _ = fi30
                                _ = v31
                            }
                        }
                        v1 -= -9223372036854775808
                        fmt.Println(v1, fi10)
                        _ = fi10
                    }
                } else {
                    delete(v2, 97)
                    for fi33 := 0; fi33 < 20; fi33++ {
                        v34 := v2[2147483648]
                        for fi35 := 0; fi35 < 19; fi35++ {
                            v1 += fi35
                            _ = fi35
                        }
                        _ = fi33
                        _ = v34
                    }
                }
                if false {
                    var v36 int32 = -100
                    var v37 uint8 = 1
                    _ = v36
                    _ = v37
                } else {
                    var v38 uint16 = 156
                    _ = v38
                }
                v39 := v2[1]
                _ = v39
            } else {
                if true {
                    switch (^(v1)) {
                    case -1:
                        switch v1 {
                        case 482:
                            v40 := []int{-1}
                            _ = v40
                        case -1:
                            v41 := GenS8{F0: 32767, F1: 18446744073709551615, F2: 9223372036854775807}
                            _ = v41
                        }
                    }
                    v42 := v3[2147483647]
                    _ = v42
                }
                var v43 int = v1
                var v44 uint16 = 622
                _ = v43
                _ = v44
            }
        }
    }
    v45 := GenS1{F0: "hello", F1: -638}
    _ = len(v3)
    fmt.Println(v1)
    for fi46 := 0; fi46 < 6; fi46++ {
        if (!((42 <= 255) || (!true))) {
            v1 = (-(v1))
            var v47 int64 = -888
            if (26 != 255) {
                switch "foo bar" {
                case "foo bar":
                    for fi48 := 0; fi48 < 2; fi48++ {
                        var v49 int = (^(fi46))
                        v1 += -1
                        v47 = (v47 >> ((886) & 63))
                        _ = fi48
                        _ = v49
                    }
                case "a":
                    v50 := GenS8{F0: 65535, F1: 4294967295, F2: v1}
                    _ = v50
                case "\n\t":
                    v3[826] = "a"
                default:
                    var v51 uint = 239
                    _ = v51
                }
                v47 = (v47 * v47)
                if ((true && false) || (!true)) {
                    for fi52 := 0; fi52 < 14; fi52++ {
                        v1 = fi52
                        fmt.Println(v1)
                        _ = fi52
                    }
                }
            } else {
                for fi53 := 0; fi53 < 17; fi53++ {
                    if (!(!true)) {
                        v1 += fi53
                        v54 := GenS9{F0: false}
                        if (2 != 8) {
                            v55 := GenS8{F0: 513, F1: 2, F2: 2147483647}
                            _ = v55
                        }
                        _ = v54
                    }
                    if false {
                        v47 = 67
                    }
                    _ = fi53
                }
                if (2 >= 128) {
                    if false {
                        for fi56 := 0; fi56 < 4; fi56++ {
                            if (true && ((!(127 > -104)) || true)) {
                                var v57 uint16 = 871
                                fmt.Println(fi46, v1, v47)
                                _ = v57
                            }
                            _ = fi56
                        }
                    }
                    switch (v1 * v1) {
                    case 9223372036854775807:
                        var v58 float64 = float64(-2.25)
                        _ = v58
                    case 932:
                        fmt.Println(v1)
                    }
                    v59 := []int{0, v1, v1}
                    _ = v59
                } else {
                    v60 := []int{fi46}
                    if (0 > 127) {
                        var v61 uint = 335
                        _ = v61
                    }
                    _ = v60
                }
                var v62 int = fi46
                _ = v62
            }
            _ = v47
        }
        v63 := GenS6{F0: ("foo bar" + "hello")}
        v1 = -280
        _ = fi46
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
}
