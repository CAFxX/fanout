package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int64 }
func genFn0() int {
    var v4 string = (("foo bar" + ((("hello" + "a") + ("" + "foo bar")) + ("hello" + "a"))) + ("hello" + ""))
    switch 32767 {
    case -422:
        v4 = ""
    case -841:
        var v5 uint32 = 2147483648
        _ = v5
    default:
        v4 = (("foo bar" + v4) + v4)
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(fi7, fi6, v4)
            var v8 uint = 18446744073709551615
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 uint32 = 700
    _ = v4
    _ = v9
    return -2147483648
}
func genFn1(p0 string) uint8 {
    p0 = p0
    var v10 int64 = -1
    _ = p0
    _ = v10
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 1
    v3 := make(map[uint32]int)
    var v11 uint32 = (v2 & 2147483647)
    v2 -= (v2 | 165)
    v12 := GenS1{F0: ("" == "hello"), F1: false, F2: 0}
    var v13 uint8 = 1
    _ = len(v1)
    switch "\n\t" {
    case "\n\t":
        v13 = (v13 << ((4294967295) & 63))
    case "":
        v2 = v11
    case "foo bar":
        var v14 int16 = 0
        _ = v14
    default:
        v13 = v13
    }
    v15 := []int{4294967295, genFn0(), 720, 4294967295}
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(v13, fi16)
        var v17 int = fi16
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 float32 = 3.14159
            if ((680 > 581) && (32768 < fi18)) {
                if true {
                    if true {
                        v20 := GenS2{F0: 326}
                        v21 := GenS2{F0: 4294967295}
                        v17 = -1
                        _ = v20
                        _ = v21
                    }
                    switch (((("a" + "a") + "hello") + "hello") + "\n\t") {
                    case "":
                        fmt.Println(v2, v13, v19)
                    }
                } else {
                    v11 += (v11 * (v11 << ((4294967295) & 63)))
                    var v22 uint64 = 911
                    _ = v22
                }
                var v23 uint = 0
                _ = v23
            } else {
                if false {
                    var v24 float32 = 0.1
                    if true {
                        fmt.Println(v11, v13, v2)
                        if ((fi16 % ((fi16) | 1)) < 917) {
                            v13 = v13
                        } else {
                            if false {
                                var v25 int16 = 32767
                                fmt.Println(v25, v24)
                                _ = v25
                            }
                        }
                        v15 = append(v15, -9223372036854775808)
                    }
                    v26 := GenS3{F0: -9223372036854775808, F1: (v13 / ((v13) | 1))}
                    v19 = v24
                    var v27 int32 = 1
                    var v28 int32 = (v27 + v27)
                    _ = v24
                    _ = v26
                    _ = v27
                    _ = v28
                }
                v29 := GenS1{F0: true, F1: false, F2: 96}
                var v30 string = "a"
                _ = v29
                _ = v30
            }
            v31 := GenS4{F0: -32768, F1: 249}
            v2 = v2
            _ = fi18
            _ = v19
            _ = v31
        }
        _ = fi16
        _ = v17
    }
    if (true && ((!(-32768 <= 32767)) || true)) {
        v11 -= (v2 | (v11 >> ((1) & 63)))
        var v32 float32 = 1.5
        switch "a" {
        case "foo bar":
            switch -181 {
            case 128:
                for fi33 := 0; fi33 < 12; fi33++ {
                    var v34 int64 = 9223372036854775807
                    switch fi33 {
                    case -491:
                        if true {
                            var v35 float64 = -2.25
                            v34 = 1
                            _ = v35
                        } else {
                            if false {
                                var v36 uint = 78
                                _ = v36
                            } else {
                                for fi37 := 0; fi37 < 13; fi37++ {
                                    var v38 int16 = -1
                                    _ = fi37
                                    _ = v38
                                }
                                v2 += v11
                                v34 = (v34 >> ((487) & 63))
                            }
                            for fi39 := 0; fi39 < 12; fi39++ {
                                if true {
                                    var v40 bool = (128 > v13)
                                    var v41 uint8 = genFn1(("\n\t" + ((("\n\t" + ("\n\t" + "foo bar")) + (((("" + "") + "foo bar") + ("hello" + ("hello" + "\n\t"))) + "\n\t")) + ("\n\t" + ("hello" + "hello")))))
                                    v41 = genFn1("foo bar")
                                    _ = v40
                                    _ = v41
                                } else {
                                    fmt.Println(fi39)
                                    switch "\n\t" {
                                    case "foo bar":
                                        v32 = ((v32 + v32) / v32)
                                    case "hello":
                                        var v42 bool = (1 != -1)
                                        _ = v42
                                    case "\n\t":
                                        v2 += (v11 ^ 42)
                                    default:
                                        switch (fi39 ^ (fi33 % ((-1) | 1))) {
                                        case 9223372036854775807:
                                            v2 += v2
                                        case 495:
                                            v11 = 65535
                                        }
                                    }
                                }
                                _ = fi39
                            }
                            var v43 uint32 = (v2 - v2)
                            _ = v43
                        }
                    case 9223372036854775807:
                        for fi44 := 0; fi44 < 19; fi44++ {
                            switch (^(fi44)) {
                            case 73:
                                v13 = (v13 << ((255) & 63))
                            case 2147483647:
                                v1[(^(v2))] = "foo bar"
                                var v45 uint16 = 42
                                _ = v45
                            case -1:
                                var v46 int64 = v34
                                _ = v46
                            default:
                                v47 := GenS4{F0: -889, F1: -921}
                                _ = v47
                            }
                            _ = fi44
                        }
                    }
                    _ = fi33
                    _ = v34
                }
            case 827:
                v48 := GenS5{F0: 929}
                _ = v48
            case -2147483648:
                v2 += 337
            }
        default:
            switch 9223372036854775807 {
            case 32768:
                var v49 uint = 4294967295
                _ = v49
            case -2147483648:
                for fi50 := 0; fi50 < 13; fi50++ {
                    v51 := v1[2147483647]
                    _ = fi50
                    _ = v51
                }
            case 127:
                v2 += (v2 - 2147483647)
            }
        }
        _ = v32
    }
    var v52 uint16 = 445
    switch genFn0() {
    case 4294967295:
        for fi53 := 0; fi53 < 9; fi53++ {
            for fi54 := 0; fi54 < 11; fi54++ {
                v13 -= 127
                _ = fi54
            }
            switch "a" {
            case "a":
                var v55 uint = 18446744073709551615
                _ = v55
            }
            _ = fi53
        }
    case -620:
        var v56 uint16 = v52
        _ = v56
    case -2147483648:
        switch ("\n\t" + (((("" + (("hello" + ("foo bar" + "hello")) + "a")) + "\n\t") + ("" + ("hello" + ("\n\t" + "hello")))) + ("hello" + "\n\t"))) {
        case "foo bar":
            _ = len(v15)
        case "hello":
            fmt.Println(v2)
        default:
            fmt.Println(v2, v13)
        }
    default:
        var v57 uint64 = 42
        _ = v57
    }
    v58 := GenS3{F0: 42, F1: 78}
    v13 = v13
    v13 = (v13 | 127)
    for fi59 := 0; fi59 < 8; fi59++ {
        v13 += v13
        _ = fi59
    }
    fmt.Println(v13, v52, v2)
    switch "" {
    case "\n\t":
        fmt.Println(v52)
    case "a":
        var v60 uint32 = (v11 ^ (v11 / ((v2) | 1)))
        _ = v60
    default:
        var v61 uint32 = (v2 / ((381) | 1))
        _ = v61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v52
    _ = v58
}
