package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint8 }
type GenS2 struct { F0 uint; F1 string; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint16; F1 int }

func main() {
    var v1 uint8 = 0
    var v2 int32 = 0
    var v3 int32 = (v2 - (^(v2)))
    switch 127 {
    case 9223372036854775807:
        switch -2147483648 {
        case 1:
            if (((("a" + "") >= "") || (1 < 127)) || true) {
                v4 := GenS1{F0: (v1 % (((v1 / ((20) | 1))) | 1)), F1: "hello", F2: (v1 & 255)}
                _ = v4
            }
        case 9223372036854775807:
            if true {
                switch "" {
                case "":
                    v5 := GenS1{F0: 42, F1: (("hello" + "foo bar") + ""), F2: v1}
                    _ = v5
                default:
                    var v6 int32 = (v2 - v2)
                    _ = v6
                }
                fmt.Println(v2, v2)
            }
        case -1:
            var v7 uint16 = 0
            _ = v7
        }
    case 32767:
        var v8 string = ((("hello" + "a") + "foo bar") + ("\n\t" + (("hello" + ("foo bar" + "foo bar")) + "")))
        _ = v8
    }
    v9 := GenS1{F0: v1, F1: (("\n\t" + "hello") + (("hello" + (("" + "foo bar") + (("hello" + "foo bar") + (("\n\t" + ("hello" + "a")) + (("\n\t" + ("a" + "\n\t")) + "hello"))))) + "")), F2: v1}
    var v10 bool = (!(!true))
    v3 = -904
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS1{F0: (v1 << ((927) & 63)), F1: ("foo bar" + ""), F2: 42}
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v10 = (!v10)
        v14 := GenS2{F0: 0, F1: ("\n\t" + "foo bar"), F2: -9223372036854775808}
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 42, F1: "hello", F2: 608}
    v16 := []int{32767, -243, 0}
    _ = len(v16)
    _ = len(v16)
    if ((1 != 42) && (-279 <= -1)) {
        switch "\n\t" {
        case "":
            var v17 int8 = 115
            _ = v17
        }
        v10 = (!(-9223372036854775808 <= 127))
        v18 := GenS1{F0: (^(v1)), F1: (("hello" + "hello") + ("\n\t" + "foo bar")), F2: 1}
        _ = v18
    } else {
        v3 += v3
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 bool = (v1 != v1)
        v1 -= (v1 ^ (^(v1)))
        _ = fi19
        _ = v20
    }
    _ = len(v16)
    v10 = v10
    if (v10 || (0 >= -32768)) {
        if (!false) {
            var v21 bool = false
            _ = v21
        } else {
            var v22 int = -368
            _ = v22
        }
    } else {
        switch 1 {
        case 42:
            for fi23 := 0; fi23 < 1; fi23++ {
                if ((!(906 > -9223372036854775808)) || (-1 > 30)) {
                    fmt.Println(v1)
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                } else {
                    var v24 uint32 = 509
                    v25 := GenS3{F0: 42}
                    if true {
                        v1 += (v1 << ((310) & 63))
                    }
                    _ = v24
                    _ = v25
                }
                v2 = -743
                _ = fi23
            }
        case -2147483648:
            if v10 {
                _ = len(v16)
                v10 = true
                var v26 int32 = 2147483647
                _ = v26
            }
        case 2147483647:
            v1 = (v1 % ((1) | 1))
        }
        for fi27 := 0; fi27 < 12; fi27++ {
            if v10 {
                var v28 int = (^(fi27))
                if len(v16) > 0 {
                    v16[0] = fi27
                }
                _ = v28
            }
            if (768 >= 404) {
                fmt.Println(v2)
            } else {
                var v29 int16 = 0
                if (409 >= 42) {
                    for fi30 := 0; fi30 < 6; fi30++ {
                        v29 += 0
                        v16 = append(v16, fi30)
                        switch (fi30 >> ((255) & 63)) {
                        case 2147483647:
                            v16 = append(v16, 377)
                        case 4294967295:
                            v31 := GenS4{F0: 1}
                            var v32 uint64 = 0
                            var v33 string = "a"
                            _ = v31
                            _ = v32
                            _ = v33
                        default:
                            var v34 float64 = 1.5
                            _ = v34
                        }
                        _ = fi30
                    }
                }
                if v10 {
                    switch fi27 {
                    case -1:
                        _ = len(v16)
                    case 0:
                        var v35 uint16 = 983
                        _ = v35
                    case 948:
                        v36 := GenS4{F0: 42}
                        _ = v36
                    default:
                        var v37 bool = (!((v29 <= (v29 * (v29 >> ((383) & 63)))) && v10))
                        _ = v37
                    }
                } else {
                    if true {
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    }
                    var v38 bool = v10
                    v1 -= (v1 >> ((105) & 63))
                    _ = v38
                }
                _ = v29
            }
            fmt.Println(v1)
            _ = fi27
        }
    }
    switch "a" {
    case "a":
        v2 += v3
    case "foo bar":
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    default:
        v2 = -1
    }
    if (456 < 827) {
        fmt.Println(v2, v2, v2)
        if (true || (32768 <= -2147483648)) {
            switch "" {
            case "hello":
                v2 += (v3 * (-(v3)))
            }
        } else {
            for fi39 := 0; fi39 < 2; fi39++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                for fi40 := 0; fi40 < 5; fi40++ {
                    if true {
                        var v41 float32 = 3.14159
                        _ = v41
                    } else {
                        v16 = append(v16, fi39)
                        var v42 uint16 = 255
                        v43 := GenS5{F0: 4294967295, F1: "\n\t"}
                        v1 -= v1
                        var v44 uint16 = (v42 * 0)
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                    _ = fi40
                }
                for fi45 := 0; fi45 < 2; fi45++ {
                    v10 = (!(fi45 > -399))
                    if v10 {
                        v10 = (v10 && true)
                        switch ("\n\t" + ("foo bar" + ("\n\t" + ("a" + "a")))) {
                        case "":
                            fmt.Println(v1, v1, v10)
                        case "foo bar":
                            v3 -= v2
                        case "a":
                            var v46 float64 = (-2.25 * ((float64(3.14159) / 3.14159) - 0.0))
                            _ = v46
                        default:
                            v47 := GenS6{F0: 1, F1: fi39}
                            _ = v47
                        }
                        v1 -= 1
                        if len(v16) > 0 {
                            v16[0] = (fi39 | 128)
                        }
                        var v48 uint64 = 639
                        _ = v48
                    }
                    if len(v16) > 0 {
                        v16[0] = 4294967295
                    }
                    _ = fi45
                }
                _ = fi39
            }
        }
    }
    v49 := GenS2{F0: 2, F1: "", F2: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v49
    _ = v9
}
