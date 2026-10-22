package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 uint; F2 bool }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 int16; F2 int32 }
type GenS8 struct { F0 int32; F1 int64 }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 int }
func genFn0(p0 uint32, p1 uint8) int32 {
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: 1}
        v7 := GenS1{F0: 42}
        p0 = (p0 * 2147483647)
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{}
    var v9 int16 = 1000
    switch "a" {
    case "\n\t":
        switch "foo bar" {
        case "":
            var v10 uint32 = 466
            _ = v10
        case "\n\t":
            var v11 float32 = 0.1
            _ = v11
        case "hello":
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 int8 = 127
                switch ("hello" + "\n\t") {
                case "a":
                    var v14 uint32 = p0
                    _ = v14
                case "hello":
                    v8 = append(v8, 32768)
                }
                if len(v8) > 0 {
                    v8[0] = (fi12 & fi12)
                }
                _ = fi12
                _ = v13
            }
        }
    case "":
        var v15 bool = (!(42 >= 366))
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -2147483648
}
func genFn1(p0 uint, p1 int64) uint64 {
    v16 := GenS2{F0: 193, F1: (p1 / ((9223372036854775807) | 1))}
    if (-1 != -1) {
        p1 += (p1 + p1)
    }
    var v17 string = "foo bar"
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 18446744073709551615
}

func main() {
    var v1 int64 = 1
    v2 := []string{}
    var v3 int8 = 42
    var v4 int64 = v1
    v3 = (v3 * 100)
    var v18 int64 = (v1 * v4)
    v4 = 0
    v19 := GenS3{F0: 0}
    if len(v2) > 0 {
        v2[0] = ""
    }
    switch "hello" {
    case "hello":
        var v20 uint = 0
        _ = v20
    default:
        switch -1 {
        case 697:
            var v21 bool = ((!true) || false)
            _ = v21
        case 4294967295:
            v18 += (^(v1))
        }
    }
    v4 += v18
    if true {
        var v22 uint64 = 1
        _ = v22
    }
    if true {
        _ = len(v2)
    }
    v3 = 0
    if ("\n\t" == "\n\t") {
        v23 := GenS4{F0: 203, F1: 567, F2: false}
        _ = v23
    } else {
        var v24 float64 = 1.5
        v1 -= v18
        _ = len(v2)
        _ = v24
    }
    var v25 uint16 = 432
    for fi26 := 0; fi26 < 19; fi26++ {
        if ((!true) || (2 == 65535)) {
            var v27 float32 = (0.0 - 0.1)
            v28 := GenS2{F0: -1, F1: (v4 << ((65535) & 63))}
            var v29 int64 = 9223372036854775807
            _ = v27
            _ = v28
            _ = v29
        } else {
            if (-2147483648 > -443) {
                var v30 uint64 = 4294967295
                var v31 bool = true
                _ = v30
                _ = v31
            }
        }
        var v32 int32 = -325
        var v33 uint64 = 120
        _ = fi26
        _ = v32
        _ = v33
    }
    switch -1 {
    case -407:
        switch -1 {
        case 1:
            if true {
                if false {
                    v25 = (v25 ^ 1)
                    if false {
                        v34 := GenS5{F0: 839, F1: -761}
                        v18 -= -894
                        _ = v34
                    }
                    v35 := GenS6{F0: ("" + "foo bar")}
                    _ = v35
                } else {
                    var v36 int32 = genFn0(2, 128)
                    switch "" {
                    case "foo bar":
                        v3 = (^(v3))
                    case "a":
                        v36 -= (v36 ^ (v36 >> ((255) & 63)))
                    default:
                        fmt.Println(v1, v4, v1)
                    }
                    fmt.Println(v1, v36, v3)
                    _ = v36
                }
            }
        }
    case 127:
        v37 := GenS7{F0: -1, F1: 1, F2: 2147483647}
        _ = v37
    }
    switch 0 {
    case -521:
        switch 1 {
        case -468:
            v25 += (v25 + v25)
        default:
            fmt.Println(v18, v18)
        }
    case 32768:
        if true {
            var v38 int32 = -1
            switch "hello" {
            case "hello":
                if (!false) {
                    fmt.Println(v3, v18, v38)
                    v1 += (v4 % (((v18 - (-(v4)))) | 1))
                } else {
                    for fi39 := 0; fi39 < 14; fi39++ {
                        var v40 bool = (!(-1 == v4))
                        v40 = ("hello" == "")
                        fmt.Println(v40, v25)
                        _ = fi39
                        _ = v40
                    }
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                }
            }
            _ = v38
        } else {
            v25 -= 0
            v41 := GenS8{F0: 1, F1: (^(v4))}
            _ = v41
        }
    case -761:
        if (65535 < 124) {
            v42 := GenS9{F0: v3, F1: 42}
            _ = v42
        }
    }
    if (!(v4 <= 1)) {
        for fi43 := 0; fi43 < 11; fi43++ {
            var v44 float32 = (0.1 * 1e10)
            v1 = (v18 << ((0) & 63))
            _ = fi43
            _ = v44
        }
        for fi45 := 0; fi45 < 1; fi45++ {
            switch 815 {
            case 535:
                for fi46 := 0; fi46 < 3; fi46++ {
                    _ = len(v2)
                    var v47 uint8 = 127
                    _ = fi46
                    _ = v47
                }
            case 496:
                v25 -= 65535
            }
            if false {
                for fi48 := 0; fi48 < 16; fi48++ {
                    var v49 uint = 255
                    var v50 int32 = 0
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            } else {
                switch "a" {
                case "a":
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                default:
                    fmt.Println(v4, fi45, v4)
                }
                switch fi45 {
                case 1:
                    var v51 int = (fi45 - fi45)
                    _ = v51
                default:
                    v1 = 953
                }
            }
            _ = fi45
        }
    } else {
        fmt.Println(v18, v18, v18)
        for fi52 := 0; fi52 < 16; fi52++ {
            var v53 int64 = -240
            if false {
                v54 := GenS10{F0: (fi52 * (fi52 & (fi52 & 127)))}
                _ = v54
            }
            _ = fi52
            _ = v53
        }
    }
    var v55 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v55
}
