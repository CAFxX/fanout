package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    v1 := []uint32{2147483647}
    v2 := []string{"hello", "", (("foo bar" + ("" + (("" + ("" + "")) + ((("a" + "a") + "foo bar") + "a")))) + ("foo bar" + "hello")), ""}
    v3 := []int{127}
    switch 32767 {
    case -9223372036854775808:
        if len(v2) > 0 {
            v2[0] = "a"
        }
    case -183:
        for fi4 := 0; fi4 < 12; fi4++ {
            for fi5 := 0; fi5 < 16; fi5++ {
                if (!false) {
                    switch (^(fi5)) {
                    case 2147483647:
                        if false {
                            v6 := GenS1{F0: 96, F1: 858}
                            var v7 int = -1
                            _ = v6
                            _ = v7
                        }
                    case 629:
                        var v8 uint32 = 326
                        _ = v8
                    }
                } else {
                    switch ("foo bar" + ((("a" + "a") + ("foo bar" + "foo bar")) + "\n\t")) {
                    case "foo bar":
                        var v9 uint32 = 4294967295
                        _ = v9
                    }
                }
                if (4294967295 > 255) {
                    if true {
                        if true {
                            fmt.Println(fi4, fi4)
                            if true {
                                var v10 bool = (!true)
                                _ = v10
                            }
                            fmt.Println(fi5, fi5)
                        }
                        var v11 float32 = 1e10
                        var v12 int32 = 2147483647
                        v13 := GenS1{F0: 1, F1: 0}
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                }
                fmt.Println(fi4, fi4)
                _ = fi5
            }
            var v14 int = -710
            _ = fi4
            _ = v14
        }
    case -1:
        var v15 uint16 = 65535
        _ = v15
    default:
        if (!false) {
            _ = len(v1)
            for fi16 := 0; fi16 < 12; fi16++ {
                v17 := GenS2{F0: -1, F1: 92}
                var v18 uint64 = 944
                _ = fi16
                _ = v17
                _ = v18
            }
            switch (("\n\t" + "hello") + "\n\t") {
            case "a":
                switch ((("a" + "foo bar") + (((("hello" + "\n\t") + "") + "a") + (("\n\t" + "hello") + ("a" + "\n\t")))) + "hello") {
                case "":
                    var v19 int32 = 2147483647
                    _ = v19
                default:
                    var v20 int64 = 902
                    _ = v20
                }
            case "":
                for fi21 := 0; fi21 < 4; fi21++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi21
                }
            }
        } else {
            var v22 int16 = -1
            var v23 float32 = (1e10 + 0.1)
            v22 += v22
            _ = v22
            _ = v23
        }
    }
    v24 := GenS1{F0: 2, F1: 538}
    for fi25 := 0; fi25 < 2; fi25++ {
        fmt.Println(fi25)
        _ = fi25
    }
    var v26 uint = 42
    if false {
        if len(v3) > 0 {
            v3[0] = 0
        }
        var v27 uint = 1
        var v28 float64 = 0.1
        _ = v27
        _ = v28
    }
    v26 = v26
    v26 = v26
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v26 -= v26
    if (0 <= -1) {
        var v29 int = 1
        if len(v3) > 0 {
            v3[0] = int(v26)
        }
        var v30 string = "a"
        _ = v29
        _ = v30
    } else {
        switch 4294967295 {
        case 9223372036854775807:
            switch 32767 {
            case 1:
                v26 = v26
            case 127:
                fmt.Println(v26)
            case 32768:
                fmt.Println(v26)
            }
        case -167:
            for fi31 := 0; fi31 < 17; fi31++ {
                var v32 bool = (!true)
                _ = fi31
                _ = v32
            }
        case 127:
            var v33 string = ("foo bar" + "\n\t")
            _ = v33
        default:
            var v34 bool = true
            _ = v34
        }
        for fi35 := 0; fi35 < 1; fi35++ {
            switch (fi35 | -9223372036854775808) {
            case -9223372036854775808:
                if (2147483647 == 533) {
                    var v36 float64 = 1.5
                    for fi37 := 0; fi37 < 16; fi37++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if true {
                            if true {
                                fmt.Println(v26, fi37, fi37)
                                v38 := GenS3{F0: false, F1: 1}
                                _ = v38
                            }
                        }
                        if (!false) {
                            var v39 int32 = 1
                            _ = v39
                        }
                        fmt.Println(v36)
                        _ = fi37
                    }
                    _ = v36
                } else {
                    v26 += v26
                    var v40 int16 = 0
                    var v41 float64 = 1.5
                    _ = v40
                    _ = v41
                }
            default:
                v26 -= 18446744073709551615
            }
            if true {
                _ = len(v2)
            }
            _ = fi35
        }
        for fi42 := 0; fi42 < 3; fi42++ {
            if true {
                if len(v1) > 0 {
                    v1[0] = 405
                }
            } else {
                for fi43 := 0; fi43 < 1; fi43++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi43
                }
            }
            v26 = 2
            _ = fi42
        }
    }
    var v44 int16 = -1
    if false {
        v44 += v44
        var v45 uint8 = 59
        _ = v45
    } else {
        if len(v2) > 0 {
            v2[0] = ""
        }
        if false {
            if false {
                for fi46 := 0; fi46 < 18; fi46++ {
                    var v47 float32 = -2.25
                    fmt.Println(fi46)
                    _ = fi46
                    _ = v47
                }
                if (false && (25 != -128)) {
                    for fi48 := 0; fi48 < 9; fi48++ {
                        fmt.Println(fi48, fi48)
                        _ = fi48
                    }
                } else {
                    for fi49 := 0; fi49 < 9; fi49++ {
                        if (!(!((fi49 + -9223372036854775808) <= fi49))) {
                            v50 := GenS1{F0: 228, F1: v26}
                            v26 -= 65535
                            _ = v50
                        } else {
                            v44 -= (v44 - (v44 & v44))
                        }
                        v51 := GenS2{F0: fi49, F1: 1}
                        _ = fi49
                        _ = v51
                    }
                }
            }
            v26 = (^(v26))
        }
    }
    fmt.Println(v44, v44)
    if (true || false) {
        v2 = append(v2, ("a" + "foo bar"))
        for fi52 := 0; fi52 < 10; fi52++ {
            v44 -= v44
            _ = fi52
        }
    } else {
        if len(v2) > 0 {
            v2[0] = ("a" + (("foo bar" + "foo bar") + "a"))
        }
        v53 := GenS2{F0: -696, F1: 1}
        _ = v53
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v44
}
