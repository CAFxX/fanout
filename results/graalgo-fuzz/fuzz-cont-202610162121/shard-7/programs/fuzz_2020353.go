package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint }
type GenS5 struct { F0 int; F1 int; F2 bool }

func main() {
    var v1 bool = (4 < 255)
    var v2 int8 = -128
    v3 := make(map[uint32]int)
    var v4 uint8 = 42
    var v5 int32 = -2147483648
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int16 = 955
        if true {
            v8 := []int{}
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 float64 = 3.14159
                v5 = (v5 << ((255) & 63))
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: 128, F1: ("a" + ""), F2: 617}
            _ = v11
            _ = v8
        }
        if (0 < 9223372036854775807) {
            var v12 bool = ((((0.1 / 1e10) <= 3.14159) && (v1 && v1)) && (2 != 129))
            _ = v12
        }
        _ = fi6
        _ = v7
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 uint16 = 0
        for fi15 := 0; fi15 < 4; fi15++ {
            if (false || (2 == 2)) {
                if v1 {
                    v16 := GenS1{F0: v4, F1: "\n\t", F2: (v14 >> ((255) & 63))}
                    var v17 bool = (v2 < (v2 >> ((255) & 63)))
                    switch "\n\t" {
                    case "\n\t":
                        var v18 int = (fi13 % (((-(fi13))) | 1))
                        _ = v18
                    case "foo bar":
                        var v19 float64 = 0.1
                        _ = v19
                    default:
                        v1 = ((2 == 4294967295) || (!v17))
                    }
                    _ = v16
                    _ = v17
                } else {
                    v5 = (v5 + (^(v5)))
                    v20 := GenS1{F0: 127, F1: ("foo bar" + ("foo bar" + ((("hello" + ("foo bar" + "foo bar")) + "\n\t") + "a"))), F2: v14}
                    _ = v20
                }
                if v1 {
                    v14 = (^(v14))
                    switch (("hello" + "a") + "foo bar") {
                    case "foo bar":
                        v21 := []int{(fi15 & 58), int(v4)}
                        v14 = (v14 << ((4294967295) & 63))
                        _ = v21
                    case "":
                        v22 := GenS2{F0: v14, F1: 973, F2: 18446744073709551615}
                        _ = v22
                    }
                    switch "foo bar" {
                    case "hello":
                        if v1 {
                            v23 := []int{fi13, 308, (fi15 | 4294967295), 127}
                            var v24 bool = v1
                            var v25 uint8 = 0
                            if (!(true && (true || (!(!v1))))) {
                                fmt.Println(v25, fi13, fi15)
                                v26 := GenS3{F0: v2}
                                v24 = v1
                                _ = v26
                            }
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    case "":
                        switch fi13 {
                        case 9223372036854775807:
                            delete(v3, 2)
                        case -9223372036854775808:
                            v27 := GenS2{F0: (v14 >> ((65535) & 63)), F1: 912, F2: 710}
                            _ = v27
                        }
                    case "\n\t":
                        v28 := []int{(fi13 >> ((748) & 63))}
                        _ = v28
                    default:
                        _ = len(v3)
                    }
                } else {
                    if (231 < 2) {
                        v14 -= (v14 | (v14 / (((v14 | (v14 + v14))) | 1)))
                    }
                }
            } else {
                if v1 {
                    v4 = (^(v4))
                    var v29 uint8 = v4
                    _ = v29
                } else {
                    v2 -= 1
                    v4 = 128
                    fmt.Println(fi15)
                }
            }
            fmt.Println(v1)
            fmt.Println(v2, fi15, v1)
            delete(v3, 42)
            _ = fi15
        }
        v2 += v2
        _ = fi13
        _ = v14
    }
    var v30 uint16 = 2
    fmt.Println(v5, v2)
    var v31 uint16 = (v30 % (((v30 / (((v30 >> ((0) & 63))) | 1))) | 1))
    var v32 float32 = 3.14159
    switch ("hello" + "\n\t") {
    case "":
        if false {
            var v33 int8 = 42
            _ = v33
        } else {
            v5 += (v5 ^ (-(v5)))
            switch "" {
            case "":
                switch 160 {
                case -846:
                    for fi34 := 0; fi34 < 4; fi34++ {
                        var v35 int8 = 28
                        for fi36 := 0; fi36 < 1; fi36++ {
                            v31 = v31
                            _ = fi36
                        }
                        _ = fi34
                        _ = v35
                    }
                default:
                    v32 += 0.0
                }
            case "\n\t":
                v37 := GenS4{F0: true, F1: 355, F2: 326}
                _ = v37
            default:
                switch "a" {
                case "hello":
                    v1 = true
                case "\n\t":
                    switch "hello" {
                    case "":
                        for fi38 := 0; fi38 < 2; fi38++ {
                            v39 := GenS1{F0: (v4 * (^(v4))), F1: (("" + "\n\t") + ("hello" + "foo bar")), F2: v30}
                            _ = fi38
                            _ = v39
                        }
                    }
                default:
                    v30 = (v31 << ((4294967295) & 63))
                }
            }
        }
    case "hello":
        v2 = v2
    case "a":
        v3[1] = 32768
    default:
        switch -1 {
        case 9223372036854775807:
            v40 := v3[1]
            _ = v40
        case -671:
            var v41 uint8 = 153
            _ = v41
        case 381:
            v42 := GenS5{F0: 2147483647, F1: 337, F2: v1}
            _ = v42
        }
    }
    if (!(!v1)) {
        var v43 uint64 = 860
        var v44 string = "a"
        var v45 uint64 = v43
        _ = v43
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 17; fi46++ {
        v3[2] = (^(fi46))
        var v47 string = "\n\t"
        _ = fi46
        _ = v47
    }
    switch "a" {
    case "a":
        v48 := GenS4{F0: v1, F1: 42, F2: 0}
        _ = v48
    }
    switch 0 {
    case 42:
        v2 += v2
    case -1:
        v49 := GenS1{F0: v4, F1: "hello", F2: 2}
        _ = v49
    case 2147483647:
        var v50 uint8 = v4
        _ = v50
    default:
        fmt.Println(v5, v4, v1)
    }
    for fi51 := 0; fi51 < 13; fi51++ {
        if (true || (73 <= 2)) {
            switch "foo bar" {
            case "\n\t":
                var v52 int8 = (v2 & (v2 * v2))
                _ = v52
            case "hello":
                v1 = false
            }
            v5 = v5
            var v53 uint16 = (v30 | v31)
            _ = v53
        }
        switch (((("hello" + (("" + "") + "hello")) + "foo bar") + "hello") + (("foo bar" + ("\n\t" + "hello")) + "\n\t")) {
        case "foo bar":
            fmt.Println(v2, v2)
        }
        v54 := GenS4{F0: false, F1: 1, F2: 42}
        _ = fi51
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
