package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint8; F1 bool; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 bool; F2 uint8 }
type GenS8 struct { F0 int64 }
func genFn0() uint {
    v7 := []int{9223372036854775807, 313, -983}
    var v8 int16 = 32767
    switch "hello" {
    case "\n\t":
        v9 := GenS1{F0: 0, F1: 42, F2: 1}
        _ = v9
    case "":
        if (606 < 566) {
            var v10 int64 = -1
            var v11 uint16 = 786
            _ = v10
            _ = v11
        }
    case "foo bar":
        v12 := GenS1{F0: 1, F1: -368, F2: 4294967295}
        _ = v12
    default:
        v13 := GenS1{F0: 701, F1: 127, F2: 2}
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        if (436 >= 1) {
            switch "hello" {
            case "foo bar":
                if ((true && true) || (65535 <= 4294967295)) {
                    v8 += (v8 >> ((2) & 63))
                }
            case "a":
                v8 -= v8
            }
            v8 -= v8
        } else {
            var v15 uint64 = 980
            _ = v15
        }
        _ = fi14
    }
    if true {
        if (!(((true || ((0 == 42) || true)) || false) || true)) {
            var v16 int16 = 42
            if true {
                fmt.Println(v8, v16)
                var v17 uint8 = 193
                _ = v17
            }
            var v18 float64 = 0.1
            _ = v16
            _ = v18
        } else {
            _ = len(v7)
            var v19 int32 = -1
            v20 := GenS2{F0: 255, F1: "a"}
            _ = v19
            _ = v20
        }
        var v21 float32 = -2.25
        if (-2147483648 < 509) {
            if ("" <= ("hello" + "")) {
                v21 -= -2.25
            }
            switch "" {
            case "foo bar":
                var v22 uint32 = 1
                _ = v22
            case "\n\t":
                for fi23 := 0; fi23 < 10; fi23++ {
                    if true {
                        var v24 int32 = -762
                        v25 := GenS2{F0: 2, F1: "hello"}
                        var v26 int32 = -1
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v21 += 0.1
                    var v27 uint32 = 2
                    _ = fi23
                    _ = v27
                }
            case "a":
                switch ("\n\t" + "a") {
                case "":
                    var v28 uint = 65535
                    _ = v28
                default:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
            }
        }
        _ = v21
    } else {
        var v29 int8 = 0
        _ = len(v7)
        fmt.Println(v29, v8)
        _ = v29
    }
    _ = v7
    _ = v8
    return 734
}

func main() {
    var v1 string = "\n\t"
    var v2 int32 = -1
    v3 := make(map[uint32]bool)
    var v4 bool = true
    var v5 uint = 0
    var v6 uint32 = 394
    fmt.Println(v1)
    var v30 uint32 = (v6 + (v6 & 42))
    var v31 uint32 = 2147483647
    var v32 float32 = 3.14159
    for fi33 := 0; fi33 < 8; fi33++ {
        if v4 {
            v34 := GenS3{F0: 128, F1: v4, F2: -35}
            var v35 uint64 = 18446744073709551615
            _ = v34
            _ = v35
        }
        _ = fi33
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        var v37 uint64 = 42
        for fi38 := 0; fi38 < 18; fi38++ {
            switch "" {
            case "foo bar":
                fmt.Println(fi38)
            case "hello":
                v3[uint32(v37)] = (2147483648 == (v31 - (v6 ^ 418)))
            case "\n\t":
                if ((int(v2) == 32768) || (100 >= -1)) {
                    v39 := GenS2{F0: 255, F1: ("" + "")}
                    v32 -= ((v32 + ((v32 * v32) / 1.0)) / v32)
                    if v4 {
                        var v40 int64 = 42
                        v6 = v31
                        v37 = v37
                        _ = v40
                    }
                    _ = v39
                }
            }
            var v41 int = (fi38 << ((v5) & 63))
            var v42 uint16 = 5
            switch v41 {
            case 971:
                v43 := GenS1{F0: (^(v42)), F1: (fi38 / ((fi38) | 1)), F2: uint32(v42)}
                _ = v43
            case 1:
                var v44 uint8 = uint8(v37)
                _ = v44
            case 2147483647:
                v45 := GenS4{F0: (v5 ^ (v5 * 586))}
                _ = v45
            default:
                if v4 {
                    fmt.Println(v4, fi36)
                    v46 := []int{}
                    fmt.Println(v37, v37)
                    _ = v46
                } else {
                    switch "a" {
                    case "":
                        var v47 int8 = 1
                        _ = v47
                    case "hello":
                        var v48 uint32 = v31
                        _ = v48
                    case "a":
                        v1 = v1
                    default:
                        var v49 int8 = 100
                        _ = v49
                    }
                    v42 = 301
                    var v50 float32 = (v32 - 0.0)
                    _ = v50
                }
            }
            _ = fi38
            _ = v41
            _ = v42
        }
        _ = fi36
        _ = v37
    }
    for fi51 := 0; fi51 < 11; fi51++ {
        v32 += ((0.0 + ((v32 * 0.0) - 3.14159)) / v32)
        v52 := GenS5{F0: 33, F1: -128}
        if true {
            fmt.Println(v32, v31)
        } else {
            v30 = v30
            v53 := GenS6{F0: -1}
            if true {
                for fi54 := 0; fi54 < 13; fi54++ {
                    switch ("foo bar" + "a") {
                    case "hello":
                        var v55 uint = (v5 ^ 255)
                        _ = v55
                    default:
                        if (!v4) {
                            switch fi54 {
                            case 42:
                                if (!false) {
                                    if true {
                                        v56 := GenS7{F0: (v2 & v2), F1: (1.5 >= 1.5), F2: 185}
                                        v57 := GenS4{F0: (^(v5))}
                                        _ = v56
                                        _ = v57
                                    }
                                    fmt.Println(v5, v5)
                                    v30 -= 2147483648
                                    fmt.Println(v6, v1, v4)
                                    delete(v3, v30)
                                }
                            case 128:
                                var v58 uint16 = 65535
                                _ = v58
                            }
                            var v59 float64 = float64(1e10)
                            if v4 {
                                v5 = (v5 % ((394) | 1))
                            }
                            _ = v59
                        }
                    }
                    v60 := GenS4{F0: v5}
                    _ = fi54
                    _ = v60
                }
                var v61 int16 = 1000
                _ = len(v3)
                _ = v61
            }
            _ = v53
        }
        _ = fi51
        _ = v52
    }
    switch v1 {
    case "":
        switch v1 {
        case "":
            if (v2 != 553) {
                var v62 uint = (^(v5))
                v4 = (1 != -563)
                _ = v62
            }
        case "foo bar":
            var v63 uint64 = 42
            _ = v63
        case "a":
            for fi64 := 0; fi64 < 9; fi64++ {
                v65 := GenS8{F0: -9223372036854775808}
                _ = fi64
                _ = v65
            }
        default:
            switch ("a" + v1) {
            case "":
                var v66 uint32 = 4294967295
                _ = v66
            }
        }
    default:
        var v67 uint16 = 0
        _ = v67
    }
    v30 = v30
    var v68 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v68
}
