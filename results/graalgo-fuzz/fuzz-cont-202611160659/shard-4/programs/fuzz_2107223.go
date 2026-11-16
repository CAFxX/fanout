package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int; F1 uint32; F2 bool }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 int; F1 int32; F2 uint32 }
func genFn0(p0 uint, p1 int, p2 int16) uint32 {
    if (32768 < -9223372036854775808) {
        if ((-128 >= 1) || (1 > (p0 + (p0 ^ p0)))) {
            p0 = p0
            var v5 int16 = (^(p2))
            for fi6 := 0; fi6 < 3; fi6++ {
                v5 += p2
                p0 = 42
                var v7 int8 = 1
                _ = fi6
                _ = v7
            }
            _ = v5
        }
        if true {
            for fi8 := 0; fi8 < 8; fi8++ {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 uint64 = 457
                    p1 = -9223372036854775808
                    _ = fi9
                    _ = v10
                }
                v11 := []int{p1, -254, 32768}
                _ = fi8
                _ = v11
            }
            var v12 uint = p0
            _ = v12
        }
        v13 := []int{}
        _ = v13
    }
    p0 -= p0
    if false {
        for fi14 := 0; fi14 < 7; fi14++ {
            if false {
                fmt.Println(p2, p1, p2)
                var v15 int = (p1 >> ((p0) & 63))
                p1 += (-(v15))
                _ = v15
            }
            _ = fi14
        }
        p2 -= (p2 / ((p2) | 1))
        fmt.Println(p0, p2, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return 183
}

func main() {
    var v1 uint64 = 42
    var v2 float32 = 3.14159
    v3 := []string{}
    var v4 int8 = -128
    if false {
        v2 = 3.14159
        if len(v3) > 0 {
            v3[0] = "a"
        }
    } else {
        switch "a" {
        case "":
            fmt.Println(v2)
        case "a":
            switch (("foo bar" + (("a" + "a") + "a")) + "a") {
            case "\n\t":
                if (false && (false || false)) {
                    v16 := GenS1{F0: ("" + "foo bar"), F1: 296}
                    v17 := GenS2{F0: 930, F1: genFn0(323, 16, 1000), F2: false}
                    fmt.Println(v4)
                    _ = v16
                    _ = v17
                } else {
                    v4 = v4
                }
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                fmt.Println(v1)
            }
        default:
            v1 += 42
        }
        v18 := GenS2{F0: 127, F1: genFn0(2, int(v4), -239), F2: (true || true)}
        _ = v18
    }
    fmt.Println(v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (4294967295 <= 2) {
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 uint16 = 255
            _ = fi19
            _ = v20
        }
        var v21 int64 = -306
        switch 128 {
        case 0:
            v2 -= -2.25
        case -352:
            v21 += 0
        case 9223372036854775807:
            v2 -= (((3.14159 / (v2 + (v2 / v2))) + float32(-2.25)) / v2)
        default:
            if true {
                var v22 uint8 = 18
                for fi23 := 0; fi23 < 17; fi23++ {
                    v21 = v21
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v24 uint64 = v1
                    _ = fi23
                    _ = v24
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v22
            } else {
                switch 32767 {
                case 127:
                    switch -9223372036854775808 {
                    case -911:
                        v25 := GenS2{F0: 127, F1: 1, F2: true}
                        _ = v25
                    }
                }
            }
        }
        _ = v21
    }
    v2 -= 1e10
    v26 := GenS3{F0: 127, F1: 83}
    v27 := GenS2{F0: 2147483647, F1: 65535, F2: (507 == -9223372036854775808)}
    var v28 int64 = -991
    for fi29 := 0; fi29 < 20; fi29++ {
        v30 := GenS2{F0: 9223372036854775807, F1: 2147483648, F2: (847 <= 32767)}
        v31 := GenS4{F0: false, F1: fi29}
        switch "foo bar" {
        case "foo bar":
            v32 := GenS4{F0: false, F1: 32767}
            _ = v32
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = (("\n\t" + ("foo bar" + "a")) + (("" + "foo bar") + ("a" + "a")))
            }
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    v33 := GenS5{F0: 0, F1: 0}
    if true {
        var v34 int32 = 134
        v4 += (v4 >> ((1) & 63))
        _ = v34
    }
    var v35 string = ("\n\t" + "")
    var v36 float64 = 1.5
    _ = len(v3)
    v4 = (v4 ^ v4)
    var v37 uint8 = 255
    if true {
        v3 = append(v3, v35)
    } else {
        v38 := GenS6{F0: 9223372036854775807, F1: 0, F2: 318}
        _ = v38
    }
    var v39 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
}
