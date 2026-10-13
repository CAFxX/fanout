package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 bool }
type GenS4 struct { F0 int64; F1 string; F2 uint }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS8 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 int = -216
    var v2 float64 = float64(float32(float64(0.0)))
    var v3 int32 = -485
    v4 := make(map[string]int)
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{}
        switch v1 {
        case 32768:
            switch "hello" {
            case "hello":
                v1 += (fi6 % ((fi6) | 1))
            case "foo bar":
                v8 := GenS1{F0: 127, F1: 32767}
                _ = v8
            case "a":
                v3 = int32(fi6)
            default:
                if false {
                    v1 -= (fi6 * (fi6 + 127))
                }
            }
        case 127:
            fmt.Println(v3)
        case -2147483648:
            v3 = v3
        }
        var v9 uint8 = 2
        _ = fi6
        _ = v7
        _ = v9
    }
    switch "\n\t" {
    case "":
        if (false || false) {
            var v10 int8 = -1
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := GenS2{F0: (fi11 * fi11), F1: 32767}
                v13 := GenS1{F0: v10, F1: -472}
                _ = fi11
                _ = v12
                _ = v13
            }
            switch "\n\t" {
            case "foo bar":
                v14 := []int{v1}
                _ = v14
            case "":
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := GenS2{F0: (-(v1)), F1: 255}
                    v17 := GenS2{F0: (^(v1)), F1: 65535}
                    v18 := GenS3{F0: 1, F1: 847, F2: ((3.14159 - (v2 + 0.1)) <= ((v2 / -2.25) * (v2 - 0.0)))}
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            case "hello":
                _ = len(v5)
            }
            _ = v10
        } else {
            if (18446744073709551615 > 779) {
                v3 += 671
                v3 = v3
                var v19 uint8 = 95
                _ = v19
            } else {
                for fi20 := 0; fi20 < 15; fi20++ {
                    v1 = fi20
                    switch "\n\t" {
                    case "\n\t":
                        var v21 uint16 = 805
                        _ = v21
                    }
                    _ = fi20
                }
                var v22 int64 = -504
                v23 := GenS2{F0: (v1 << ((255) & 63)), F1: 2}
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 int32 = v3
                v26 := GenS4{F0: 273, F1: "foo bar", F2: 546}
                _ = fi24
                _ = v25
                _ = v26
            }
            if (-128 != -128) {
                v2 = (3.14159 - 1e10)
            }
        }
    case "a":
        var v27 uint64 = 42
        _ = v27
    case "\n\t":
        if true {
            v1 += int(v3)
            var v28 int8 = 127
            if true {
                v29 := []int{-329, v1}
                _ = v29
            }
            _ = v28
        }
    }
    var v30 int8 = 18
    if true {
        if (2 < v30) {
            var v31 uint64 = 27
            v31 += v31
            _ = v31
        }
        var v32 float32 = 1e10
        if true {
            for fi33 := 0; fi33 < 5; fi33++ {
                v1 = 32768
                _ = fi33
            }
        }
        _ = v32
    } else {
        var v34 uint32 = 680
        _ = v34
    }
    _ = len(v5)
    for fi35 := 0; fi35 < 14; fi35++ {
        for fi36 := 0; fi36 < 10; fi36++ {
            v2 = (3.14159 / 0.1)
            _ = fi36
        }
        _ = fi35
    }
    if true {
        for fi37 := 0; fi37 < 16; fi37++ {
            switch fi37 {
            case 0:
                var v38 bool = false
                _ = v38
            case 2147483647:
                var v39 uint32 = 2
                _ = v39
            case 1:
                if true {
                    var v40 float64 = (float64(0.1) - v2)
                    v41 := GenS4{F0: -399, F1: ("foo bar" + "a"), F2: 757}
                    var v42 int16 = 0
                    _ = v40
                    _ = v41
                    _ = v42
                } else {
                    var v43 uint8 = 6
                    if true {
                        v44 := v4[(((("" + "") + "a") + "foo bar") + ("hello" + "foo bar"))]
                        _ = v44
                    }
                    if false {
                        var v45 uint8 = v43
                        v46 := []int{128, (^(fi37)), (-(fi37))}
                        _ = v45
                        _ = v46
                    }
                    _ = v43
                }
            }
            v47 := GenS5{F0: -345, F1: 932}
            if (false && (false || false)) {
                var v48 string = "hello"
                for fi49 := 0; fi49 < 12; fi49++ {
                    v50 := GenS6{F0: 32768, F1: -948}
                    var v51 int64 = 1
                    var v52 int = (fi37 ^ (^(fi49)))
                    v3 += -1
                    _ = fi49
                    _ = v50
                    _ = v51
                    _ = v52
                }
                _ = v48
            } else {
                v53 := GenS6{F0: 0, F1: -969}
                _ = v53
            }
            _ = fi37
            _ = v47
        }
    }
    var v54 int64 = -9223372036854775808
    v54 += 9223372036854775807
    for fi55 := 0; fi55 < 13; fi55++ {
        v56 := GenS7{F0: 0, F1: 462, F2: 65535}
        _ = fi55
        _ = v56
    }
    v57 := GenS8{F0: 115, F1: 255}
    fmt.Println(v30, v1, v1)
    switch (("hello" + "a") + ("hello" + "\n\t")) {
    case "a":
        if ("foo bar" != (("a" + "\n\t") + "\n\t")) {
            v58 := GenS3{F0: 1, F1: 2, F2: (!(!(1000 >= 0)))}
            _ = v58
        } else {
            fmt.Println(v2, v1)
        }
    }
    v59 := v5["\n\t"]
    var v60 int32 = -1
    var v61 int = (v1 >> ((914) & 63))
    for fi62 := 0; fi62 < 8; fi62++ {
        var v63 uint8 = 128
        v64 := v5[""]
        v54 = (v54 << ((979) & 63))
        _ = fi62
        _ = v63
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v54
    _ = v57
    _ = v59
    _ = v60
    _ = v61
}
