package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int8; F1 int64; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 int64; F2 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }

func main() {
    v1 := []string{"foo bar", (("a" + ("foo bar" + ("hello" + "\n\t"))) + "hello")}
    var v2 uint16 = 2
    var v3 int = -1
    v4 := GenS1{F0: -634, F1: -128}
    if (-1 <= 120) {
        v5 := GenS1{F0: v3, F1: -128}
        if (42 <= 802) {
            if (42 > -369) {
                fmt.Println(v2)
            } else {
                var v6 int8 = 42
                if len(v1) > 0 {
                    v1[0] = ((("a" + "") + (("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))) + "hello")) + "foo bar")
                }
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                _ = v6
            }
            if (true || (0 != 18446744073709551615)) {
                if (-2147483648 >= v3) {
                    var v7 bool = ((^(v3)) != (v3 % ((-1) | 1)))
                    _ = v7
                }
                v8 := GenS2{F0: -78, F1: -1, F2: (v3 / ((v3) | 1))}
                for fi9 := 0; fi9 < 12; fi9++ {
                    v2 = v2
                    fmt.Println(fi9)
                    v3 += fi9
                    _ = fi9
                }
                _ = v8
            }
            v10 := GenS1{F0: 32768, F1: 1}
            _ = v10
        }
        var v11 int16 = -725
        _ = v11
        _ = v5
    } else {
        fmt.Println(v2, v2)
    }
    if true {
        if false {
            fmt.Println(v2, v2)
            fmt.Println(v2, v3)
            v3 = -9223372036854775808
        } else {
            if true {
                var v12 int = v3
                _ = v12
            }
        }
    } else {
        v3 = (v3 ^ 4294967295)
        v13 := GenS1{F0: v3, F1: -124}
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 int64 = 9223372036854775807
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    v16 := GenS1{F0: (v3 | v3), F1: -128}
    if ((!true) || (true || false)) {
        if (false || true) {
            if ((("foo bar" + "") + "a") < ("\n\t" + ((("hello" + "hello") + "foo bar") + ("foo bar" + ("\n\t" + "hello"))))) {
                fmt.Println(v2)
                v3 += -1
                fmt.Println(v3, v2)
            } else {
                if true {
                    var v17 int16 = -1
                    _ = v17
                }
                var v18 int = -832
                var v19 float32 = ((3.14159 + 0.0) + float32(((0.1 + (1e10 / (((0.1 * float64((float32(1e10) - (float32(float64(1.5)) / (0.0 + 0.1))))) * float64((3.14159 * 0.1))) + -2.25))) - (1.5 + float64((3.14159 + -2.25))))))
                _ = v18
                _ = v19
            }
            v20 := GenS2{F0: -128, F1: int64(v3), F2: (-(v3))}
            var v21 uint16 = 65535
            _ = v20
            _ = v21
        } else {
            var v22 uint32 = 2147483647
            _ = v22
        }
        v3 -= 1
    } else {
        v23 := GenS3{F0: 1}
        var v24 bool = (673 <= -9223372036854775808)
        _ = v23
        _ = v24
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            var v25 uint8 = 1
            for fi26 := 0; fi26 < 3; fi26++ {
                switch (fi26 << ((2) & 63)) {
                case -2147483648:
                    for fi27 := 0; fi27 < 9; fi27++ {
                        v25 = v25
                        v1 = append(v1, ("foo bar" + ""))
                        _ = fi27
                    }
                }
                v28 := GenS1{F0: (fi26 % (((^(v3))) | 1)), F1: 42}
                fmt.Println(v2, v2, v3)
                _ = fi26
                _ = v28
            }
            v3 -= (v3 >> ((839) & 63))
            _ = v25
        }
        var v29 uint64 = 293
        _ = v29
    }
    if (2 >= 0) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch "" {
        case "foo bar":
            if false {
                for fi30 := 0; fi30 < 4; fi30++ {
                    fmt.Println(v2, v2, v3)
                    fmt.Println(v2, fi30)
                    _ = fi30
                }
                var v31 string = "a"
                _ = v31
            } else {
                if len(v1) > 0 {
                    v1[0] = ("a" + ("foo bar" + (("\n\t" + ("" + "a")) + "hello")))
                }
            }
        case "":
            for fi32 := 0; fi32 < 18; fi32++ {
                for fi33 := 0; fi33 < 6; fi33++ {
                    v34 := GenS4{F0: 0, F1: -1, F2: 131}
                    switch ("hello" + "\n\t") {
                    case "a":
                        if true {
                            fmt.Println(fi33, fi32)
                        }
                    }
                    if (592 > (v3 % (((^(fi33))) | 1))) {
                        v3 -= fi33
                        var v35 uint32 = 782
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v35
                    } else {
                        var v36 uint16 = v2
                        _ = v36
                    }
                    _ = fi33
                    _ = v34
                }
                for fi37 := 0; fi37 < 18; fi37++ {
                    var v38 uint16 = v2
                    fmt.Println(fi32, fi32)
                    _ = fi37
                    _ = v38
                }
                _ = fi32
            }
        default:
            if ((1.5 + (1.5 * 1e10)) <= (-2.25 / 1.0)) {
                if (!false) {
                    if false {
                        var v39 uint8 = 127
                        fmt.Println(v3, v3, v3)
                        fmt.Println(v2, v2, v39)
                        _ = v39
                    } else {
                        var v40 uint64 = 276
                        for fi41 := 0; fi41 < 11; fi41++ {
                            var v42 int16 = 42
                            var v43 int = v3
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        switch (v3 % (((v3 + v3)) | 1)) {
                        case 236:
                            var v44 uint64 = v40
                            _ = v44
                        default:
                            fmt.Println(v40, v3)
                        }
                        _ = v40
                    }
                    if (601 < 2) {
                        var v45 uint64 = 42
                        _ = v45
                    }
                } else {
                    for fi46 := 0; fi46 < 16; fi46++ {
                        if (42 != 42) {
                            fmt.Println(v2, fi46)
                            var v47 uint8 = 1
                            v48 := GenS2{F0: 0, F1: 9223372036854775807, F2: -9223372036854775808}
                            _ = v47
                            _ = v48
                        }
                        _ = fi46
                    }
                }
                fmt.Println(v2)
                var v49 uint64 = 42
                _ = v49
            }
        }
    } else {
        switch ("hello" + ("a" + "")) {
        case "foo bar":
            v50 := GenS5{F0: (((("hello" + "\n\t") + "a") + "") + ("\n\t" + "\n\t"))}
            _ = v50
        }
        for fi51 := 0; fi51 < 2; fi51++ {
            v52 := GenS6{F0: 0}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v53 uint = uint(v2)
            _ = fi51
            _ = v52
            _ = v53
        }
    }
    fmt.Println(v2)
    _ = len(v1)
    var v54 uint64 = 1
    fmt.Println(v3, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v54
}
