package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 bool; F1 uint32; F2 string }
func genFn0() int64 {
    var v4 string = "hello"
    v4 = "hello"
    if false {
        var v5 int = 9223372036854775807
        v4 = v4
        var v6 int32 = 0
        _ = v5
        _ = v6
    } else {
        v4 = ("a" + v4)
        var v7 int16 = 42
        _ = v7
    }
    v4 = ""
    var v8 uint8 = 255
    _ = v4
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 459
    var v3 uint16 = 300
    switch "foo bar" {
    case "a":
        switch -580 {
        case -937:
            for fi9 := 0; fi9 < 3; fi9++ {
                v2 += v2
                _ = fi9
            }
        }
    }
    var v10 int32 = -375
    v11 := []int{-639, 1, 32767}
    v3 -= (v3 + v3)
    switch "hello" {
    case "hello":
        var v12 uint = 255
        _ = v12
    case "foo bar":
        for fi13 := 0; fi13 < 2; fi13++ {
            switch "hello" {
            case "\n\t":
                if ((fi13 / (((fi13 << ((814) & 63))) | 1)) < (fi13 | 128)) {
                    var v14 int8 = 1
                    v15 := GenS1{F0: 0, F1: 2, F2: v2}
                    _ = v14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        var v17 int64 = -9223372036854775808
                        _ = fi16
                        _ = v17
                    }
                    v18 := GenS2{F0: -128, F1: (v3 >> ((18446744073709551615) & 63))}
                    var v19 int16 = (v2 / ((0) | 1))
                    _ = v18
                    _ = v19
                }
            case "foo bar":
                var v20 uint32 = 467
                _ = v20
            default:
                var v21 uint8 = 42
                _ = v21
            }
            v22 := GenS1{F0: -72, F1: 4294967295, F2: 1}
            if true {
                if false {
                    for fi23 := 0; fi23 < 1; fi23++ {
                        v2 -= (v2 % (((v2 + (v2 >> ((440) & 63)))) | 1))
                        var v24 int = -2147483648
                        v1[502] = ("" + "hello")
                        var v25 int = fi23
                        v26 := GenS2{F0: 1, F1: v3}
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v10 += v10
                }
            }
            _ = fi13
            _ = v22
        }
    default:
        var v27 uint = 65535
        _ = v27
    }
    v11 = append(v11, 9223372036854775807)
    for fi28 := 0; fi28 < 13; fi28++ {
        v29 := GenS3{F0: 18446744073709551615, F1: int32(v3)}
        v30 := v1[156]
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        if true {
            switch "\n\t" {
            case "a":
                for fi32 := 0; fi32 < 10; fi32++ {
                    v33 := GenS3{F0: 255, F1: (v10 - -2147483648)}
                    if (1 > 42) {
                        v10 = v10
                        var v34 string = "\n\t"
                        if len(v11) > 0 {
                            v11[0] = fi32
                        }
                        _ = v34
                    } else {
                        v35 := GenS2{F0: 106, F1: 51}
                        _ = v35
                    }
                    v36 := GenS4{F0: (true && (("foo bar" == ((((("\n\t" + (("a" + ("hello" + "a")) + (("a" + "") + ("hello" + ((("foo bar" + "\n\t") + "") + (("\n\t" + ("a" + (("\n\t" + "\n\t") + "\n\t"))) + "a")))))) + "\n\t") + "\n\t") + "foo bar") + ("" + "\n\t"))) && true)), F1: 2, F2: "\n\t"}
                    _ = fi32
                    _ = v33
                    _ = v36
                }
            case "":
                if (9223372036854775807 < 957) {
                    if (true && true) {
                        var v37 int16 = v2
                        fmt.Println(v10)
                        v10 = (^(v10))
                        _ = v37
                    } else {
                        v2 -= v2
                        v38 := v1[2147483648]
                        _ = v38
                    }
                }
            case "foo bar":
                var v39 int64 = 504
                _ = v39
            default:
                for fi40 := 0; fi40 < 3; fi40++ {
                    var v41 uint64 = 674
                    fmt.Println(v2)
                    _ = fi40
                    _ = v41
                }
            }
            v42 := GenS3{F0: 1, F1: v10}
            var v43 int8 = 73
            _ = v42
            _ = v43
        } else {
            var v44 int16 = (v2 + v2)
            var v45 uint64 = 4294967295
            _ = v44
            _ = v45
        }
        for fi46 := 0; fi46 < 19; fi46++ {
            if (-2.25 == (1.5 * 3.14159)) {
                var v47 uint16 = (v3 - (v3 + v3))
                switch 127 {
                case -38:
                    v2 -= -1
                case -151:
                    var v48 int16 = v2
                    _ = v48
                case 32767:
                    var v49 int64 = 0
                    _ = v49
                }
                _ = v47
            }
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi46
        }
        v3 = (v3 - v3)
        _ = fi31
    }
    fmt.Println(v10, v10, v10)
    for fi50 := 0; fi50 < 12; fi50++ {
        for fi51 := 0; fi51 < 15; fi51++ {
            v11 = append(v11, -1)
            var v52 int = 0
            _ = fi51
            _ = v52
        }
        if (42 != -107) {
            v53 := GenS4{F0: (1.5 >= ((0.0 + 0.0) - 1.5)), F1: 4294967295, F2: ""}
            _ = v53
        }
        var v54 uint64 = uint64(v3)
        _ = fi50
        _ = v54
    }
    if ((859 != 42) && true) {
        fmt.Println(v3)
    } else {
        v55 := GenS4{F0: (!(true || false)), F1: 65535, F2: "\n\t"}
        if (1 <= 255) {
            v2 = (v2 | 32767)
            for fi56 := 0; fi56 < 18; fi56++ {
                v57 := GenS1{F0: 127, F1: 42, F2: v2}
                switch (fi56 >> ((447) & 63)) {
                case 32768:
                    switch "foo bar" {
                    case "hello":
                        v10 += -901
                    case "foo bar":
                        var v58 int32 = 42
                        _ = v58
                    default:
                        v10 += 468
                    }
                case -715:
                    delete(v1, 2)
                case 0:
                    v59 := GenS1{F0: -1, F1: 42, F2: v2}
                    _ = v59
                }
                v10 -= 1
                _ = fi56
                _ = v57
            }
        } else {
            v3 += (v3 * v3)
        }
        for fi60 := 0; fi60 < 10; fi60++ {
            if true {
                var v61 uint = 71
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                fmt.Println(v2, v2, v61)
                _ = v61
            }
            _ = fi60
        }
        _ = v55
    }
    if (!false) {
        v3 -= (v3 << ((970) & 63))
    }
    for fi62 := 0; fi62 < 4; fi62++ {
        v10 -= v10
        _ = fi62
    }
    var v63 int64 = -1
    v2 -= v2
    fmt.Println(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v63
}
