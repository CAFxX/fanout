package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS5 struct { F0 string; F1 uint64; F2 uint16 }
type GenS6 struct { F0 int16; F1 int64 }
func genFn0(p0 int16, p1 uint) int8 {
    var v7 bool = (false && true)
    p1 = p1
    switch -1 {
    case -599:
        p1 += p1
    case -773:
        v8 := GenS1{F0: -2147483648}
        _ = v8
    case 1:
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 11; fi10++ {
                fmt.Println(p0, fi9, p1)
                var v11 uint16 = 2
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    default:
        var v12 int = 0
        _ = v12
    }
    v13 := GenS2{F0: 473}
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS1{F0: -1}
        for fi16 := 0; fi16 < 2; fi16++ {
            fmt.Println(fi14, v7)
            _ = fi16
        }
        v17 := GenS2{F0: 2}
        _ = fi14
        _ = v15
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return 65
}

func main() {
    var v1 bool = ((1 <= 42) || (0 < 1))
    var v2 float32 = 3.14159
    v3 := []uint32{2, 0, 1}
    var v4 int = -9223372036854775808
    var v5 int8 = 120
    v6 := []string{"", ("a" + ""), "\n\t"}
    switch 127 {
    case 411:
        switch v4 {
        case 32768:
            var v18 int8 = 1
            _ = v18
        case 9223372036854775807:
            v5 += (^(v5))
        default:
            var v19 float64 = 0.1
            _ = v19
        }
    default:
        switch (v4 + v4) {
        case 0:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 11:
            v20 := GenS2{F0: 1}
            _ = v20
        default:
            for fi21 := 0; fi21 < 5; fi21++ {
                if len(v3) > 0 {
                    v3[0] = 0
                }
                v22 := GenS2{F0: 277}
                v23 := GenS1{F0: 2147483647}
                _ = fi21
                _ = v22
                _ = v23
            }
        }
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 uint8 = 0
        var v26 uint32 = 186
        _ = fi24
        _ = v25
        _ = v26
    }
    if v1 {
        switch ("\n\t" + "") {
        case "hello":
            v2 = v2
        case "foo bar":
            v2 = (0.0 / v2)
        case "\n\t":
            v1 = false
        }
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 int8 = (v5 - v5)
            _ = fi27
            _ = v28
        }
        _ = len(v3)
    } else {
        if len(v6) > 0 {
            v6[0] = ("a" + "foo bar")
        }
        v4 -= v4
    }
    var v29 uint16 = 32767
    v5 += v5
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (32768 <= v4) {
        _ = len(v3)
        v30 := GenS2{F0: 42}
        var v31 uint64 = 1
        _ = v30
        _ = v31
    }
    v2 = ((((-2.25 / float32(3.14159)) - v2) * (v2 - (1.5 + v2))) + v2)
    if v1 {
        for fi32 := 0; fi32 < 11; fi32++ {
            v1 = v1
            v1 = true
            if false {
                for fi33 := 0; fi33 < 12; fi33++ {
                    if (v1 && (!(1e10 == 0.1))) {
                        v34 := GenS3{F0: v1, F1: 42}
                        fmt.Println(v29, v29, fi33)
                        var v35 bool = (42 >= 0)
                        _ = v34
                        _ = v35
                    } else {
                        fmt.Println(v2, fi33, v29)
                        v4 = 32767
                        v36 := GenS1{F0: 42}
                        if (true || (958 < 1)) {
                            if len(v6) > 0 {
                                v6[0] = ("\n\t" + "a")
                            }
                            switch "a" {
                            case "\n\t":
                                switch "\n\t" {
                                case "foo bar":
                                    v37 := GenS2{F0: 726}
                                    _ = v37
                                case "":
                                    v38 := GenS4{F0: v29, F1: -1, F2: 2}
                                    _ = v38
                                case "hello":
                                    var v39 int32 = 0
                                    _ = v39
                                }
                            case "foo bar":
                                if ("a" > ("foo bar" + "hello")) {
                                    v40 := GenS2{F0: 2}
                                    switch "hello" {
                                    case "a":
                                        var v41 int8 = (^(v5))
                                        _ = v41
                                    case "hello":
                                        fmt.Println(fi32, v29, v5)
                                    }
                                    v1 = ((^(v29)) != v29)
                                    _ = v40
                                }
                            case "a":
                                v29 -= (v29 >> ((0) & 63))
                            default:
                                v3 = append(v3, 765)
                            }
                            var v42 int = -1
                            _ = v42
                        } else {
                            v43 := GenS5{F0: "", F1: 960, F2: 780}
                            _ = v43
                        }
                        _ = v36
                    }
                    fmt.Println(fi33)
                    _ = fi33
                }
            }
            _ = fi32
        }
        var v44 int = v4
        _ = v44
    } else {
        v1 = ((0 != -509) || (v1 || false))
    }
    v1 = v1
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v45 := GenS6{F0: -1, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
