package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 int32; F1 uint8 }

func main() {
    v1 := []string{"", "\n\t", ""}
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    switch 13 {
    case -1:
        v4 := GenS1{F0: ""}
        _ = v4
    case -672:
        var v5 uint32 = 4294967295
        _ = v5
    case 4294967295:
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS2{F0: "\n\t", F1: 18446744073709551615, F2: 32767}
            _ = fi6
            _ = v7
        }
    default:
        for fi8 := 0; fi8 < 8; fi8++ {
            switch (^(fi8)) {
            case 128:
                var v9 float32 = ((float32(0.0) * (-2.25 - -2.25)) / 1e10)
                _ = v9
            }
            for fi10 := 0; fi10 < 17; fi10++ {
                switch ("a" + "hello") {
                case "hello":
                    var v11 string = (("foo bar" + ((("hello" + "foo bar") + "\n\t") + (("a" + "a") + "foo bar"))) + "")
                    _ = v11
                }
                _ = fi10
            }
            for fi12 := 0; fi12 < 2; fi12++ {
                if (true || true) {
                    if (18446744073709551615 < 31) {
                        var v13 float64 = 0.0
                        switch ("\n\t" + "\n\t") {
                        case "":
                            var v14 int32 = 2147483647
                            _ = v14
                        }
                        v15 := GenS3{F0: 2}
                        _ = v13
                        _ = v15
                    } else {
                        fmt.Println(fi8, fi8)
                    }
                    var v16 int32 = 42
                    if false {
                        if true {
                            if false {
                                v17 := v3["hello"]
                                _ = v17
                            }
                            if (!(-128 > 42)) {
                                var v18 string = ("\n\t" + "hello")
                                v16 = v16
                                _ = v18
                            } else {
                                var v19 uint64 = 94
                                var v20 float32 = 3.14159
                                fmt.Println(v19, v16)
                                _ = v19
                                _ = v20
                            }
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        }
                    } else {
                        if ((("foo bar" + "\n\t") > "") && (1 <= 255)) {
                            if ((!(65535 <= 690)) || true) {
                                v21 := GenS4{F0: 42}
                                var v22 int8 = -91
                                _ = v21
                                _ = v22
                            }
                            if false {
                                var v23 uint8 = 255
                                _ = v23
                            }
                        } else {
                            v24 := GenS3{F0: 204}
                            v1 = append(v1, (("a" + "a") + (("hello" + "hello") + "foo bar")))
                            v25 := GenS5{F0: (-(v16))}
                            _ = v24
                            _ = v25
                        }
                    }
                    _ = v16
                }
                _ = fi12
            }
            _ = fi8
        }
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 uint32 = 445
            fmt.Println(fi26)
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    if len(v1) > 0 {
        v1[0] = ("a" + "a")
    }
    if ((!(false && (2147483647 >= 0))) || (4294967295 < 726)) {
        switch "hello" {
        case "foo bar":
            var v29 uint16 = 32767
            _ = v29
        default:
            var v30 int64 = -9223372036854775808
            _ = v30
        }
        var v31 float32 = float32(0.0)
        _ = v31
    } else {
        var v32 int16 = 1000
        if true {
            v33 := GenS5{F0: -870}
            v1 = append(v1, ("foo bar" + "a"))
            _ = v33
        } else {
            for fi34 := 0; fi34 < 12; fi34++ {
                for fi35 := 0; fi35 < 13; fi35++ {
                    v32 -= v32
                    switch (fi35 & 32768) {
                    case 1:
                        v32 = 80
                    case 4294967295:
                        var v36 uint64 = 174
                        _ = v36
                    case 32767:
                        fmt.Println(v32, fi34, fi34)
                    }
                    _ = fi35
                }
                _ = fi34
            }
        }
        v37 := GenS2{F0: (("" + "\n\t") + "hello"), F1: 374, F2: 65535}
        _ = v32
        _ = v37
    }
    var v38 uint16 = 32767
    switch -470 {
    case 553:
        if ((false && (!true)) || (730 >= 42)) {
            v38 = 65535
            if (!false) {
                v38 -= (v38 & v38)
            }
        }
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        if false {
            v38 += v38
            switch -901 {
            case 32767:
                if (float32(((0.0 / -2.25) + -2.25)) < (1e10 - (1e10 + ((1.5 * (1e10 + -2.25)) * (3.14159 + -2.25))))) {
                    var v40 int8 = -1
                    fmt.Println(v38, fi39, v40)
                    _ = v40
                } else {
                    var v41 float64 = (1e10 / 3.14159)
                    v41 = 1e10
                    _ = v41
                }
            case 32768:
                v42 := GenS6{F0: 42, F1: 42}
                _ = v42
            }
        } else {
            var v43 string = "a"
            v38 = 42
            switch (fi39 - (fi39 + 148)) {
            case 9223372036854775807:
                if (885 > 2) {
                    fmt.Println(v38, v38, fi39)
                    v44 := v3[v43]
                    _ = v44
                }
            default:
                v45 := GenS4{F0: 0}
                _ = v45
            }
            _ = v43
        }
        var v46 bool = false
        _ = fi39
        _ = v46
    }
    switch 128 {
    case -661:
        fmt.Println(v38)
    case 1:
        fmt.Println(v38)
    }
    for fi47 := 0; fi47 < 1; fi47++ {
        for fi48 := 0; fi48 < 9; fi48++ {
            switch ("hello" + ("" + "")) {
            case "foo bar":
                _ = len(v2)
                v38 = v38
                if true {
                    v38 = v38
                } else {
                    var v49 uint64 = 2
                    if (!(1 == 127)) {
                        fmt.Println(fi47, v49)
                        v49 -= v49
                    }
                    v50 := GenS4{F0: v49}
                    _ = v49
                    _ = v50
                }
            case "":
                var v51 float32 = 3.14159
                _ = v51
            case "hello":
                switch "a" {
                case "":
                    if true {
                        if true {
                            var v52 int = fi47
                            _ = v52
                        } else {
                            var v53 int16 = 0
                            if false {
                                switch "" {
                                case "":
                                    v38 -= v38
                                }
                                switch (((("a" + "") + ("hello" + "hello")) + "\n\t") + "hello") {
                                case "a":
                                    _ = len(v1)
                                }
                            } else {
                                var v54 int16 = int16(fi47)
                                v38 -= v38
                                switch "foo bar" {
                                case "foo bar":
                                    var v55 uint8 = 2
                                    _ = v55
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                var v56 uint = 65535
                                _ = v54
                                _ = v56
                            }
                            _ = v53
                        }
                    } else {
                        var v57 bool = (!(((-9223372036854775808 > 42) && true) && (252 < 480)))
                        var v58 int64 = 9223372036854775807
                        _ = v57
                        _ = v58
                    }
                case "a":
                    var v59 int32 = -997
                    _ = v59
                case "\n\t":
                    var v60 int8 = 42
                    _ = v60
                }
            default:
                var v61 int = (fi47 | (fi47 - fi48))
                _ = v61
            }
            _ = fi48
        }
        delete(v2, (fi47 | (-(fi47))))
        _ = fi47
    }
    if false {
        v62 := GenS7{F0: 1, F1: 106}
        _ = v62
    }
    var v63 uint16 = v38
    v63 -= (v63 ^ (v38 - (v38 << ((1) & 63))))
    if len(v1) > 0 {
        v1[0] = ""
    }
    fmt.Println(v38, v63)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v63
}
