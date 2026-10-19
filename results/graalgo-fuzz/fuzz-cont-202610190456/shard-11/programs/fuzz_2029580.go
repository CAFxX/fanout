package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int16; F2 uint64 }
type GenS3 struct { F0 int; F1 string; F2 int64 }
type GenS4 struct { F0 int64; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 int32; F1 bool; F2 int }
type GenS9 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 1e10
    v3 := make(map[int]string)
    v4 := []int{32767, -2147483648, 163}
    v5 := make(map[int]bool)
    v2 -= float64(3.14159)
    if (-32768 >= -868) {
        if (!(-2147483648 == 2147483647)) {
            v2 += v2
        }
    } else {
        switch 2147483647 {
        case 1:
            var v6 float32 = 1.5
            _ = v6
        case 779:
            switch ("hello" + "hello") {
            case "foo bar":
                for fi7 := 0; fi7 < 12; fi7++ {
                    switch -249 {
                    case 0:
                        var v8 int = fi7
                        _ = v8
                    case -2147483648:
                        for fi9 := 0; fi9 < 5; fi9++ {
                            v2 = -2.25
                            _ = fi9
                        }
                    case -792:
                        v2 += v2
                    default:
                        v10 := GenS1{F0: 0}
                        _ = v10
                    }
                    v11 := GenS1{F0: 2}
                    fmt.Println(v2, fi7)
                    _ = fi7
                    _ = v11
                }
            default:
                v12 := GenS2{F0: ("\n\t" > ("foo bar" + "hello")), F1: 32767, F2: 0}
                _ = v12
            }
        case 188:
            v2 = float64(1e10)
        default:
            v2 = (0.0 * 1.5)
        }
        fmt.Println(v2)
        if ("a" > "") {
            switch ("a" + "a") {
            case "hello":
                v13 := GenS1{F0: 65535}
                _ = v13
            case "":
                switch -1 {
                case 42:
                    v2 -= v2
                case 9223372036854775807:
                    if true {
                        var v14 uint8 = 0
                        v2 = -2.25
                        _ = v14
                    } else {
                        if true {
                            if (-1 >= -1) {
                                v2 -= float64(((1.5 - float32(v2)) - 3.14159))
                                for fi15 := 0; fi15 < 9; fi15++ {
                                    v16 := v5[fi15]
                                    var v17 float64 = 1.5
                                    if v16 {
                                        v18 := GenS3{F0: (fi15 % ((4294967295) | 1)), F1: "a", F2: 0}
                                        switch "a" {
                                        case "a":
                                            v16 = (0.0 != ((1.5 / (float32(v2) + 0.1)) - 1e10))
                                        case "foo bar":
                                            v4 = append(v4, fi15)
                                        }
                                        fmt.Println(v17, v16, v16)
                                        _ = v18
                                    }
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                }
                                fmt.Println(v2)
                            } else {
                                var v19 int16 = -237
                                var v20 uint16 = 2
                                _ = v19
                                _ = v20
                            }
                            var v21 int = 42
                            _ = v21
                        }
                    }
                default:
                    fmt.Println(v2)
                }
            case "\n\t":
                var v22 float32 = -2.25
                _ = v22
            default:
                var v23 float32 = -2.25
                _ = v23
            }
            var v24 bool = (false && false)
            _ = v24
        } else {
            for fi25 := 0; fi25 < 10; fi25++ {
                v4 = append(v4, (^(fi25)))
                _ = fi25
            }
            var v26 int = 720
            if true {
                fmt.Println(v26, v26)
            }
            _ = v26
        }
    }
    v27 := GenS2{F0: ((!true) || (!true)), F1: 647, F2: 42}
    v28 := GenS3{F0: 42, F1: "foo bar", F2: 1}
    var v29 uint64 = 42
    v30 := GenS2{F0: false, F1: 377, F2: v29}
    for fi31 := 0; fi31 < 11; fi31++ {
        switch -175 {
        case -228:
            var v32 float64 = ((v2 - 1e10) - (-2.25 + (1e10 / -2.25)))
            _ = v32
        }
        if true {
            switch "hello" {
            case "\n\t":
                for fi33 := 0; fi33 < 8; fi33++ {
                    v34 := GenS3{F0: (fi31 + (-(fi31))), F1: "", F2: 9223372036854775807}
                    v29 -= v29
                    _ = fi33
                    _ = v34
                }
            default:
                var v35 bool = (v2 == ((float64(-2.25) * ((float64(float32((0.0 - ((1.5 + 1e10) + float64(float32(v2)))))) / v2) * 3.14159)) + v2))
                _ = v35
            }
            v36 := GenS2{F0: false, F1: 1000, F2: 4294967295}
            _ = v36
        }
        _ = fi31
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        _ = len(v5)
        _ = fi37
    }
    if true {
        var v38 bool = false
        _ = v38
    } else {
        for fi39 := 0; fi39 < 16; fi39++ {
            if true {
                if false {
                    if (255 <= 42) {
                        if false {
                            v40 := GenS4{F0: -861, F1: 1000, F2: 360}
                            _ = v40
                        } else {
                            _ = len(v1)
                            if false {
                                v29 = 2
                            }
                        }
                    } else {
                        v2 = (3.14159 * float64((1.5 + 1.5)))
                    }
                    v41 := GenS5{F0: 65535, F1: 42}
                    _ = v41
                } else {
                    v2 = (v2 + (((-2.25 * 1.5) * (v2 / -2.25)) / ((1.5 * v2) / 1.5)))
                    v42 := GenS6{F0: -1}
                    _ = v42
                }
                var v43 int = (fi39 & 32767)
                if (false || (v43 != v43)) {
                    v44 := GenS7{F0: (v29 ^ v29), F1: -9223372036854775808}
                    _ = v44
                }
                _ = v43
            } else {
                fmt.Println(v2)
                var v45 uint64 = (v29 << ((2) & 63))
                _ = v45
            }
            _ = fi39
        }
        v29 = 1
        for fi46 := 0; fi46 < 15; fi46++ {
            v47 := GenS8{F0: 2147483647, F1: ((("foo bar" + ("\n\t" + ("a" + ""))) + (("" + ("hello" + (("a" + "a") + "hello"))) + ("foo bar" + ("\n\t" + "a")))) < "foo bar"), F2: -2147483648}
            fmt.Println(v29, v29, fi46)
            v48 := v1[(fi46 + fi46)]
            _ = fi46
            _ = v47
            _ = v48
        }
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v2 = v2
    default:
        switch -855 {
        case -385:
            v29 += (v29 << ((137) & 63))
        case -1:
            fmt.Println(v29, v29)
        }
    }
    fmt.Println(v29, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v49 := GenS1{F0: 208}
    v50 := GenS1{F0: 65535}
    for fi51 := 0; fi51 < 15; fi51++ {
        v2 = float64(1.5)
        v4 = append(v4, fi51)
        var v52 uint = 0
        _ = fi51
        _ = v52
    }
    for fi53 := 0; fi53 < 12; fi53++ {
        v2 = v2
        for fi54 := 0; fi54 < 20; fi54++ {
            v55 := GenS9{F0: 0, F1: ""}
            switch (fi53 - int(v29)) {
            case 549:
                fmt.Println(fi53, v2, v29)
            case 1:
                v56 := GenS1{F0: 874}
                _ = v56
            case 676:
                _ = len(v4)
            default:
                _ = len(v4)
            }
            _ = fi54
            _ = v55
        }
        v2 += v2
        _ = fi53
    }
    var v57 uint = 578
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v57
}
