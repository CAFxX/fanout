package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 string; F1 uint16 }

func main() {
    var v1 int = 42
    var v2 float32 = float32(0.1)
    v3 := make(map[int]int)
    v4 := []string{"", ("" + ("foo bar" + "foo bar")), "hello"}
    var v5 int = v1
    var v6 float64 = -2.25
    for fi7 := 0; fi7 < 3; fi7++ {
        switch (v1 & fi7) {
        case 2147483647:
            var v8 float32 = 0.0
            _ = v8
        case 918:
            v9 := GenS1{F0: 18446744073709551615}
            _ = v9
        case -187:
            for fi10 := 0; fi10 < 17; fi10++ {
                if true {
                    _ = len(v3)
                    var v11 uint8 = 42
                    _ = v11
                }
                v5 += 0
                var v12 int8 = -74
                _ = fi10
                _ = v12
            }
        }
        _ = fi7
    }
    v6 = v6
    v4 = append(v4, "a")
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            if (-128 != 1) {
                var v15 int32 = 365
                fmt.Println(v1)
                if true {
                    v6 = ((1.5 * v6) + 1.5)
                    v5 += -628
                } else {
                    v15 += (v15 % ((-2147483648) | 1))
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = v15
            }
            v16 := GenS2{F0: 36}
            switch "\n\t" {
            case "\n\t":
                v17 := GenS2{F0: 1}
                _ = v17
            default:
                v18 := GenS2{F0: 4294967295}
                v6 = 0.0
                switch "" {
                case "":
                    if (9223372036854775807 <= -9223372036854775808) {
                        switch "foo bar" {
                        case "\n\t":
                            var v19 int32 = -2147483648
                            _ = v19
                        case "a":
                            v6 = 0.1
                            switch v5 {
                            case -1:
                                v20 := GenS2{F0: 1}
                                _ = v20
                            case -9223372036854775808:
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                            default:
                                v21 := GenS3{F0: 1, F1: 4294967295}
                                _ = v21
                            }
                        }
                    } else {
                        v6 = float64((-2.25 + v2))
                        if false {
                            fmt.Println(fi14)
                            v5 = (^(fi13))
                            v6 += ((v6 * v6) - v6)
                        } else {
                            fmt.Println(v2, v1)
                            var v22 float64 = v6
                            _ = v22
                        }
                    }
                }
                _ = v18
            }
            v23 := GenS2{F0: 90}
            _ = fi14
            _ = v16
            _ = v23
        }
        switch (-(fi13)) {
        case -368:
            _ = len(v4)
        default:
            var v24 string = ("" + "foo bar")
            _ = v24
        }
        var v25 uint32 = 488
        _ = fi13
        _ = v25
    }
    v1 = 741
    v2 += ((v2 + 1e10) * v2)
    for fi26 := 0; fi26 < 13; fi26++ {
        v6 = (float64((1.5 / float32((((v6 + 1.5) * 1.5) * 3.14159)))) - (1e10 / float64(v2)))
        if false {
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 uint = 1
                _ = fi27
                _ = v28
            }
        } else {
            var v29 string = ""
            _ = len(v3)
            _ = v29
        }
        v5 = v5
        _ = fi26
    }
    v30 := GenS4{F0: 42}
    v6 += v6
    var v31 int = (^(v5))
    v32 := GenS5{F0: 137, F1: -47}
    _ = len(v3)
    if (false || false) {
        if false {
            var v33 uint8 = 247
            _ = v33
        }
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 uint8 = 13
            _ = fi34
            _ = v35
        }
        v36 := GenS6{F0: "a", F1: 32767}
        _ = v36
    }
    var v37 int8 = 42
    switch (v31 + v31) {
    case 127:
        switch "\n\t" {
        case "foo bar":
            v38 := v3[v5]
            _ = v38
        case "\n\t":
            var v39 int64 = 42
            _ = v39
        case "hello":
            fmt.Println(v6, v31)
        default:
            if true {
                v4 = append(v4, ("a" + "\n\t"))
                var v40 float32 = v2
                switch v5 {
                case 4294967295:
                    v4 = append(v4, "a")
                case 9223372036854775807:
                    _ = len(v4)
                case -424:
                    var v41 uint64 = 42
                    _ = v41
                }
                _ = v40
            }
        }
    case 0:
        fmt.Println(v6, v5)
    }
    for fi42 := 0; fi42 < 7; fi42++ {
        if false {
            if false {
                _ = len(v3)
                for fi43 := 0; fi43 < 10; fi43++ {
                    v1 = (-(v5))
                    fmt.Println(fi42)
                    v5 += (-(fi43))
                    _ = fi43
                }
                v5 = v31
            } else {
                var v44 uint8 = 0
                _ = v44
            }
        }
        if false {
            v1 = (v31 ^ (v31 | v1))
            v45 := GenS5{F0: 42, F1: 100}
            _ = v45
        } else {
            v46 := GenS4{F0: -887}
            v31 -= v1
            _ = v46
        }
        var v47 int = (-(fi42))
        _ = fi42
        _ = v47
    }
    _ = len(v3)
    switch "hello" {
    case "":
        var v48 int64 = -9223372036854775808
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
