package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 string; F1 int64; F2 int32 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 uint8 }
func genFn0() int8 {
    v5 := GenS1{F0: ("hello" + ""), F1: 6}
    switch 9223372036854775807 {
    case 4294967295:
        v6 := []int{-2147483648, 128, 1, -511}
        _ = v6
    case 0:
        if false {
            if false {
                switch ("" + "") {
                case "\n\t":
                    v7 := []int{-607}
                    _ = v7
                }
            }
        } else {
            var v8 bool = true
            var v9 uint32 = 1
            if true {
                var v10 int64 = 761
                var v11 float64 = (0.1 * 1e10)
                _ = v10
                _ = v11
            }
            _ = v8
            _ = v9
        }
    case -9223372036854775808:
        var v12 uint8 = 0
        _ = v12
    default:
        var v13 string = ""
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 uint64 = 1
        var v16 int32 = 42
        v15 -= v15
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = v5
    return 0
}
func genFn1() int8 {
    if (3.14159 == float64((1.5 * 0.1))) {
        var v17 int = -9223372036854775808
        v18 := GenS1{F0: "\n\t", F1: 32767}
        _ = v17
        _ = v18
    } else {
        if ((true || (336 != 4294967295)) || true) {
            var v19 uint64 = 2
            _ = v19
        } else {
            var v20 int = 4294967295
            fmt.Println(v20)
            _ = v20
        }
        v21 := GenS2{F0: "\n\t", F1: -1, F2: 0}
        v22 := GenS3{F0: "a", F1: -1}
        _ = v21
        _ = v22
    }
    if false {
        var v23 uint16 = 255
        var v24 uint32 = 65535
        var v25 uint16 = v23
        _ = v23
        _ = v24
        _ = v25
    }
    return -128
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 1
    v3 := []string{"", ("foo bar" + ((("" + "") + "\n\t") + "hello")), "\n\t"}
    v4 := []uint32{1, 2147483647, 42}
    v2 -= v2
    var v26 uint = 2
    switch -624 {
    case 192:
        var v27 uint32 = 977
        _ = v27
    case 2147483647:
        v26 += v26
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 bool = false
        fmt.Println(v26, v2, v2)
        if (127 > 128) {
            v26 -= (v26 << ((v26) & 63))
            var v30 uint16 = v2
            for fi31 := 0; fi31 < 6; fi31++ {
                var v32 float32 = 3.14159
                var v33 int32 = -1
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = v30
        } else {
            v2 = 0
        }
        _ = fi28
        _ = v29
    }
    switch "" {
    case "a":
        v34 := GenS2{F0: "a", F1: 42, F2: 2147483647}
        _ = v34
    default:
        v26 = (v26 * uint(v2))
    }
    v2 = v2
    for fi35 := 0; fi35 < 1; fi35++ {
        v36 := GenS1{F0: "a", F1: 0}
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        for fi38 := 0; fi38 < 10; fi38++ {
            fmt.Println(fi37, fi38)
            if ((2 > 128) || true) {
                var v39 string = ("\n\t" + "\n\t")
                _ = v39
            }
            _ = fi38
        }
        var v40 int64 = 42
        switch "a" {
        case "\n\t":
            switch 598 {
            case 0:
                for fi41 := 0; fi41 < 6; fi41++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch "" {
                    case "\n\t":
                        v1[1] = 870
                    case "":
                        switch "hello" {
                        case "":
                            var v42 uint8 = 253
                            _ = v42
                        default:
                            switch fi37 {
                            case 32768:
                                v26 = (v26 % (((v26 * 4294967295)) | 1))
                            case -52:
                                v40 -= v40
                            case 0:
                                v2 = v2
                            default:
                                v43 := GenS4{F0: ("foo bar" + "hello")}
                                _ = v43
                            }
                        }
                    }
                    _ = fi41
                }
            case -1:
                v40 += -1
            }
        case "a":
            var v44 int16 = -1
            _ = v44
        case "":
            v2 += (v2 ^ (^(v2)))
        default:
            switch "foo bar" {
            case "hello":
                v45 := GenS5{F0: (v26 & (v26 ^ (v26 ^ (v26 << (((v26 / (((^(v26))) | 1))) & 63))))), F1: 147}
                _ = v45
            default:
                v26 = v26
            }
        }
        _ = fi37
        _ = v40
    }
    if (0.1 == float32((((0.1 * float64(((1.5 - 0.0) + 0.0))) * 1.5) - 3.14159))) {
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        if (genFn1() <= -1) {
            var v46 string = ("a" + ("hello" + "a"))
            v46 = "\n\t"
            _ = v46
        }
    } else {
        if len(v3) > 0 {
            v3[0] = (("hello" + ("hello" + "\n\t")) + "a")
        }
        v26 -= v26
        var v47 int = 42
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
