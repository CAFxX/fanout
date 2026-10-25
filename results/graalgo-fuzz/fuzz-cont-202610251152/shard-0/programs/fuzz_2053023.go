package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int; F1 uint64 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    v4 := []uint32{961, 2147483648}
    var v5 uint16 = 284
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            fmt.Println(fi7, v5, fi8)
            var v9 float32 = 1.5
            v5 += (v5 >> ((0) & 63))
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 int64 = -1
    switch 128 {
    case 404:
        var v11 int = -547
        _ = v11
    }
    v12 := GenS1{F0: (true && (0.0 != -2.25)), F1: 42}
    var v13 uint = 0
    if ((v5 - 785) > 233) {
        if len(v4) > 0 {
            v4[0] = 266
        }
        switch -9223372036854775808 {
        case 32768:
            fmt.Println(v5, v13, v10)
        default:
            switch "foo bar" {
            case "hello":
                v14 := GenS1{F0: (true && true), F1: 4294967295}
                _ = v14
            }
        }
    } else {
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := GenS2{F0: (("\n\t" + (("a" + "foo bar") + ("hello" + ""))) + "\n\t"), F1: fi15}
            switch fi15 {
            case 9223372036854775807:
                var v17 uint8 = 225
                _ = v17
            case 42:
                switch (fi15 / (((^(fi15))) | 1)) {
                case -221:
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi18
                    }
                }
            case 462:
                for fi19 := 0; fi19 < 13; fi19++ {
                    if false {
                        if (false || ("\n\t" == "foo bar")) {
                            delete(v1, 32768)
                            if ((false && ((true && true) || (-1 >= -680))) || (!false)) {
                                fmt.Println(v5, v10)
                                _ = len(v3)
                                if ((("a" + "a") + (("foo bar" + "") + ("" + ((("\n\t" + "") + "\n\t") + "")))) == "") {
                                    fmt.Println(v13, fi19)
                                    _ = len(v2)
                                }
                            }
                        } else {
                            v10 -= 1
                            v20 := GenS3{F0: 65535}
                            var v21 int = 215
                            _ = v20
                            _ = v21
                        }
                        if ("hello" != "\n\t") {
                            fmt.Println(fi15, fi15, v13)
                            v5 = (v5 >> ((2) & 63))
                        }
                        fmt.Println(v5, fi19, v5)
                    }
                    _ = fi19
                }
            default:
                v5 = v5
            }
            _ = fi15
            _ = v16
        }
        if false {
            v22 := GenS4{F0: 10, F1: v13}
            _ = v22
        } else {
            v13 += 42
            fmt.Println(v10)
            if ("a" <= ("" + "a")) {
                v10 -= (v10 / (((v10 * int64(v5))) | 1))
                for fi23 := 0; fi23 < 12; fi23++ {
                    var v24 uint = uint(v5)
                    v24 = 946
                    if ((167 >= 1) && (804 < 42)) {
                        var v25 uint16 = 757
                        for fi26 := 0; fi26 < 7; fi26++ {
                            _ = len(v4)
                            _ = fi26
                        }
                        if (4294967295 == 895) {
                            v27 := GenS4{F0: fi23, F1: v24}
                            v1[-460] = false
                            v13 = v24
                            _ = v27
                        }
                        _ = v25
                    } else {
                        v13 -= (v13 + v24)
                    }
                    _ = fi23
                    _ = v24
                }
            }
        }
    }
    var v28 bool = (!("" < "hello"))
    if (9223372036854775807 > 2147483647) {
        v13 -= v13
        v29 := v2[42]
        _ = v29
    }
    var v30 uint8 = 2
    var v31 int32 = -645
    v32 := GenS5{F0: int(v5), F1: 42}
    v13 -= (v13 << (((v13 | v13)) & 63))
    _ = len(v4)
    v33 := GenS2{F0: ("a" + "a"), F1: 2147483647}
    var v34 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
