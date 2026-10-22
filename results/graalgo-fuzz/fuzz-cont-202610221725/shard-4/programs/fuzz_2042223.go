package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int16; F1 uint; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32; F1 int8 }
type GenS10 struct { F0 int32 }

func main() {
    var v1 bool = (!(42 <= 100))
    var v2 bool = false
    var v3 uint16 = 2
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]bool)
    var v6 float64 = float64(3.14159)
    fmt.Println(v2, v3)
    v7 := GenS1{F0: "\n\t", F1: 1}
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(v2, fi8)
        v2 = (("a" + "") >= "foo bar")
        _ = fi8
    }
    v6 -= v6
    v1 = v1
    var v9 int8 = -92
    if (v9 > v9) {
        switch "a" {
        case "a":
            var v10 uint8 = 128
            _ = v10
        default:
            v11 := GenS1{F0: "", F1: 164}
            _ = v11
        }
    }
    switch (("hello" + "") + ("" + "hello")) {
    case "hello":
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 float32 = 0.1
            fmt.Println(v13)
            for fi14 := 0; fi14 < 11; fi14++ {
                v3 -= 1
                switch (("\n\t" + (("foo bar" + ("a" + "a")) + "\n\t")) + ("hello" + "")) {
                case "":
                    var v15 int16 = 1
                    _ = v15
                default:
                    v16 := GenS1{F0: (("hello" + "") + ("" + "a")), F1: 179}
                    _ = v16
                }
                fmt.Println(v6, v6, v2)
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
    case "\n\t":
        if true {
            if true {
                switch "hello" {
                case "hello":
                    if v1 {
                        if (18446744073709551615 < 18446744073709551615) {
                            if (!((("\n\t" + "\n\t") + "\n\t") <= (("\n\t" + (("hello" + "a") + "\n\t")) + ""))) {
                                v17 := GenS2{F0: v9, F1: 0}
                                _ = v17
                            }
                            v2 = (1 <= 128)
                            for fi18 := 0; fi18 < 4; fi18++ {
                                for fi19 := 0; fi19 < 14; fi19++ {
                                    switch (fi19 | fi18) {
                                    case -292:
                                        var v20 int32 = -2147483648
                                        _ = v20
                                    case 32767:
                                        switch "\n\t" {
                                        case "":
                                            v6 = (v6 + v6)
                                        case "foo bar":
                                            var v21 bool = (v1 || (!v2))
                                            _ = v21
                                        }
                                        switch "a" {
                                        case "hello":
                                            if (32767 < 42) {
                                                var v22 uint32 = 687
                                                v23 := GenS3{F0: -1, F1: 1, F2: fi18}
                                                _ = v22
                                                _ = v23
                                            } else {
                                                switch (fi19 | (fi19 & 9223372036854775807)) {
                                                case -341:
                                                    _ = len(v4)
                                                case -1:
                                                    var v24 uint16 = (v3 + v3)
                                                    _ = v24
                                                }
                                                if false {
                                                    _ = len(v5)
                                                    v25 := []int{(-(fi19)), 1}
                                                    switch (fi18 & (fi19 % ((fi19) | 1))) {
                                                    case 1:
                                                        fmt.Println(v6, fi19)
                                                    }
                                                    if len(v25) > 0 {
                                                        v25[0] = (fi19 & (fi19 >> ((4294967295) & 63)))
                                                    }
                                                    var v26 int = (fi18 | (fi18 - (-(fi19))))
                                                    _ = v25
                                                    _ = v26
                                                } else {
                                                    switch (-(fi18)) {
                                                    case -2147483648:
                                                        if true {
                                                            fmt.Println(fi18, fi18)
                                                            v27 := GenS4{F0: v9}
                                                            _ = v27
                                                        }
                                                    case -1:
                                                        fmt.Println(fi18)
                                                    default:
                                                        v1 = (-128 > v9)
                                                    }
                                                }
                                                v9 = -64
                                            }
                                        case "a":
                                            v9 = (v9 * 73)
                                        case "foo bar":
                                            switch (fi18 + (fi19 % ((-356) | 1))) {
                                            case 127:
                                                v6 = float64(0.0)
                                            default:
                                                var v28 int16 = -1
                                                _ = v28
                                            }
                                        default:
                                            var v29 uint16 = (v3 | v3)
                                            _ = v29
                                        }
                                        switch fi18 {
                                        case -615:
                                            v1 = v2
                                        default:
                                            v6 = v6
                                        }
                                    default:
                                        if (v2 && (false || v2)) {
                                            v30 := v5["hello"]
                                            switch "" {
                                            case "":
                                                var v31 uint = 42
                                                _ = v31
                                            case "a":
                                                var v32 bool = (!(970 < fi19))
                                                _ = v32
                                            case "foo bar":
                                                if (128 != 128) {
                                                    fmt.Println(v6)
                                                    if false {
                                                        var v33 uint = 557
                                                        _ = v33
                                                    } else {
                                                        var v34 uint8 = 127
                                                        var v35 bool = (v9 != (-(v9)))
                                                        var v36 int32 = 1
                                                        _ = v34
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                    v37 := GenS5{F0: 1, F1: -840}
                                                    if v1 {
                                                        switch fi19 {
                                                        case -1:
                                                            v5["foo bar"] = false
                                                        case 127:
                                                            fmt.Println(fi18, v6)
                                                        default:
                                                            if (!(0.1 < 0.0)) {
                                                                var v38 int32 = 1
                                                                _ = v38
                                                            } else {
                                                                v6 = -2.25
                                                            }
                                                        }
                                                        _ = len(v5)
                                                    }
                                                    _ = v37
                                                }
                                            default:
                                                v39 := GenS6{F0: 54}
                                                _ = v39
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                var v40 int16 = 32767
                                                _ = v40
                                            default:
                                                v6 = (v6 + v6)
                                            }
                                            _ = v30
                                        }
                                    }
                                    fmt.Println(v3, fi18)
                                    _ = fi19
                                }
                                _ = fi18
                            }
                        } else {
                            for fi41 := 0; fi41 < 4; fi41++ {
                                fmt.Println(v9, v6)
                                for fi42 := 0; fi42 < 13; fi42++ {
                                    v1 = v1
                                    v43 := GenS7{F0: (v9 & 1)}
                                    v44 := GenS3{F0: -114, F1: 18446744073709551615, F2: fi42}
                                    switch ("hello" + ("hello" + ("\n\t" + ""))) {
                                    case "\n\t":
                                        fmt.Println(fi42)
                                    case "":
                                        v1 = v1
                                    case "hello":
                                        var v45 uint8 = 127
                                        _ = v45
                                    default:
                                        var v46 int16 = 1000
                                        v47 := []int{67, fi42, fi42}
                                        _ = v46
                                        _ = v47
                                    }
                                    _ = fi42
                                    _ = v43
                                    _ = v44
                                }
                                _ = fi41
                            }
                            v48 := GenS8{F0: (-(v9))}
                            _ = v48
                        }
                        v49 := GenS9{F0: 42, F1: v9}
                        v50 := GenS10{F0: 2147483647}
                        _ = v49
                        _ = v50
                    }
                case "a":
                    v2 = (!true)
                }
                var v51 int32 = 438
                for fi52 := 0; fi52 < 4; fi52++ {
                    var v53 uint64 = 18446744073709551615
                    switch "a" {
                    case "":
                        v54 := []int{}
                        _ = v54
                    case "hello":
                        v4[2] = 4294967295
                    default:
                        v3 = (^(v3))
                    }
                    _ = fi52
                    _ = v53
                }
                _ = v51
            }
        }
    default:
        fmt.Println(v1)
    }
    v55 := v5[("\n\t" + "a")]
    v6 -= (float64(0.0) - float64((0.1 - 1e10)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v7
    _ = v9
}
