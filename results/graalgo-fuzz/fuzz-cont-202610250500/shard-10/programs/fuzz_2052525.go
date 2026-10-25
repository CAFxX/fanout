package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 int64; F1 uint64; F2 int16 }
type GenS7 struct { F0 string; F1 uint64 }
func genFn0() string {
    v5 := []int{}
    if true {
        switch "" {
        case "":
            var v6 int = 1
            _ = v6
        case "a":
            var v7 float64 = 3.14159
            _ = v7
        }
    } else {
        var v8 int64 = 42
        var v9 bool = (true || true)
        for fi10 := 0; fi10 < 4; fi10++ {
            switch "" {
            case "hello":
                var v11 uint64 = 42
                _ = v11
            case "\n\t":
                v12 := GenS1{F0: 4294967295, F1: 0, F2: 42}
                _ = v12
            }
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    switch ("foo bar" + "\n\t") {
    case "a":
        var v13 int = -592
        _ = v13
    default:
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint = 419
            v15 = (v15 * 42)
            _ = fi14
            _ = v15
        }
    }
    _ = v5
    return "hello"
}
func genFn1(p0 int, p1 int32, p2 bool) int {
    fmt.Println(p2, p0)
    v16 := GenS2{F0: 0}
    v17 := GenS2{F0: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return -1
}

func main() {
    var v1 uint = 114
    v2 := []int{9223372036854775807, 32768}
    var v3 float32 = (0.0 - ((-2.25 + 1.5) - ((-2.25 / 3.14159) * 3.14159)))
    var v4 uint8 = 128
    v4 -= 41
    for fi18 := 0; fi18 < 10; fi18++ {
        switch (fi18 | fi18) {
        case 32768:
            v19 := GenS3{F0: -1, F1: (-(fi18)), F2: 959}
            _ = v19
        }
        for fi20 := 0; fi20 < 13; fi20++ {
            switch fi18 {
            case 127:
                fmt.Println(fi18, v4)
            case -578:
                if (144 <= 42) {
                    fmt.Println(v4, fi20, v4)
                } else {
                    var v21 bool = true
                    _ = v21
                }
            }
            _ = fi20
        }
        var v22 uint8 = (v4 << ((255) & 63))
        _ = fi18
        _ = v22
    }
    v23 := GenS2{F0: (v4 + v4)}
    if (!true) {
        switch ("foo bar" + "a") {
        case "foo bar":
            switch 4294967295 {
            case 128:
                switch "hello" {
                case "a":
                    v3 = ((-2.25 / (1e10 - 0.1)) + float32(-2.25))
                case "hello":
                    v24 := GenS1{F0: 42, F1: (v4 << (((v1 + 793)) & 63)), F2: v4}
                    _ = v24
                case "":
                    v25 := GenS3{F0: 42, F1: 4294967295, F2: 1}
                    _ = v25
                }
            case -874:
                fmt.Println(v3, v4, v3)
            }
        case "a":
            v1 -= (^(v1))
        default:
            for fi26 := 0; fi26 < 15; fi26++ {
                for fi27 := 0; fi27 < 17; fi27++ {
                    v1 = (v1 << ((255) & 63))
                    v1 += (v1 * 2)
                    _ = fi27
                }
                v28 := GenS4{F0: -32768, F1: 0}
                if true {
                    if len(v2) > 0 {
                        v2[0] = 158
                    }
                }
                _ = fi26
                _ = v28
            }
        }
        v29 := GenS3{F0: -128, F1: 1, F2: 2147483647}
        v1 -= 1
        _ = v29
    }
    if len(v2) > 0 {
        v2[0] = 127
    }
    var v30 float32 = v3
    var v31 int8 = 100
    v1 += 18446744073709551615
    switch 107 {
    case 9223372036854775807:
        if len(v2) > 0 {
            v2[0] = 32767
        }
    }
    if (false && true) {
        if (65535 > 1) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v32 := GenS2{F0: 174}
            v3 -= v3
            _ = v32
        }
        if (0 != v1) {
            var v33 int16 = 0
            if true {
                v1 -= 42
                for fi34 := 0; fi34 < 18; fi34++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v35 uint8 = (v4 << (((v1 | v1)) & 63))
                    _ = fi34
                    _ = v35
                }
            } else {
                var v36 int8 = v31
                _ = v36
            }
            v1 = 0
            _ = v33
        }
        for fi37 := 0; fi37 < 17; fi37++ {
            switch fi37 {
            case 2147483647:
                fmt.Println(v3, v31)
            case 489:
                switch (("hello" + ("foo bar" + "")) + "hello") {
                case "hello":
                    var v38 int16 = -32768
                    _ = v38
                case "\n\t":
                    v3 += float32(1e10)
                default:
                    var v39 uint64 = 4294967295
                    _ = v39
                }
            case 127:
                if (true || false) {
                    for fi40 := 0; fi40 < 6; fi40++ {
                        v1 += (v1 ^ 255)
                        switch (((("\n\t" + "\n\t") + (("foo bar" + ("" + "foo bar")) + "foo bar")) + "foo bar") + "a") {
                        case "":
                            v41 := GenS3{F0: (v31 << ((1) & 63)), F1: fi37, F2: 170}
                            _ = v41
                        }
                        if false {
                            if true {
                                v42 := GenS1{F0: v1, F1: (v4 % ((v4) | 1)), F2: 127}
                                v1 = v1
                                var v43 string = "\n\t"
                                _ = v42
                                _ = v43
                            }
                            fmt.Println(v30, v31)
                        }
                        _ = fi40
                    }
                    v44 := GenS1{F0: 648, F1: 42, F2: 42}
                    _ = v44
                }
            default:
                var v45 float64 = float64(0.0)
                _ = v45
            }
            v3 += v30
            _ = fi37
        }
    } else {
        v1 = (v1 >> ((uint(v4)) & 63))
        switch -1 {
        case 42:
            if (!(0.1 >= float64((-2.25 * ((float32(0.0) / -2.25) / (v30 + float32(float64(((((((0.0 * 1e10) * 3.14159) - float32(1e10)) * v30) * (v3 / 1.0)) * v30))))))))) {
                for fi46 := 0; fi46 < 19; fi46++ {
                    var v47 uint64 = 0
                    fmt.Println(v3)
                    fmt.Println(v30, v31, v31)
                    _ = fi46
                    _ = v47
                }
                switch "a" {
                case "":
                    for fi48 := 0; fi48 < 3; fi48++ {
                        v49 := GenS5{F0: (v1 ^ 4294967295), F1: (679 == 560)}
                        v1 += v1
                        for fi50 := 0; fi50 < 11; fi50++ {
                            v3 = v3
                            _ = fi50
                        }
                        _ = fi48
                        _ = v49
                    }
                case "foo bar":
                    if true {
                        if (42 < 1) {
                            if true {
                                var v51 int16 = -462
                                v52 := GenS6{F0: 943, F1: 0, F2: -1}
                                _ = v51
                                _ = v52
                            } else {
                                fmt.Println(v30, v1)
                            }
                            v1 = (^(v1))
                        } else {
                            v2 = append(v2, -459)
                            v4 = v4
                            fmt.Println(v4, v3)
                        }
                        fmt.Println(v31)
                    }
                case "a":
                    v4 += v4
                default:
                    v2 = append(v2, 4294967295)
                }
            } else {
                v30 = v30
            }
        case 815:
            if ((v4 % ((100) | 1)) >= v4) {
                v4 = (v4 & 41)
            } else {
                if false {
                    v2 = append(v2, 1)
                }
            }
        case 0:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    for fi53 := 0; fi53 < 11; fi53++ {
        if len(v2) > 0 {
            v2[0] = (fi53 << ((v1) & 63))
        }
        _ = fi53
    }
    for fi54 := 0; fi54 < 11; fi54++ {
        v31 = v31
        var v55 int = int(v1)
        _ = fi54
        _ = v55
    }
    if true {
        if true {
            if false {
                if false {
                    v1 = (v1 | uint(v4))
                }
                switch "hello" {
                case "foo bar":
                    v56 := GenS7{F0: "hello", F1: 695}
                    _ = v56
                case "a":
                    if false {
                        v57 := GenS2{F0: (v4 & 128)}
                        var v58 uint16 = 2
                        fmt.Println(v30)
                        _ = v57
                        _ = v58
                    }
                }
                var v59 bool = (v31 != v31)
                _ = v59
            }
            _ = len(v2)
            fmt.Println(v3, v4)
        } else {
            if len(v2) > 0 {
                v2[0] = 9223372036854775807
            }
        }
    } else {
        if len(v2) > 0 {
            v2[0] = 128
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
