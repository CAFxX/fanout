package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint16; F2 bool }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS6 struct { F0 int; F1 int; F2 uint32 }
type GenS7 struct { F0 int64 }
func genFn0() bool {
    var v4 float32 = 1.5
    for fi5 := 0; fi5 < 1; fi5++ {
        switch (fi5 << ((4294967295) & 63)) {
        case 0:
            v6 := GenS1{F0: 1, F1: 0, F2: 127}
            _ = v6
        }
        _ = fi5
    }
    fmt.Println(v4)
    if (606 != 18446744073709551615) {
        fmt.Println(v4)
    }
    v7 := []int{42, 42, 176, 383}
    _ = v4
    _ = v7
    return false
}
func genFn1(p0 int32) int64 {
    var v8 string = ("hello" + "a")
    var v9 float64 = ((((0.1 + (((3.14159 - float64(float32(((float64(1.5) / 3.14159) / (((1.5 * (1e10 - 0.1)) + (1.5 - (0.0 - 0.1))) - 1e10))))) / 3.14159) - 1e10)) + ((float64(3.14159) + -2.25) - ((0.0 / 1.5) / 3.14159))) * 1.5) + 1e10)
    v10 := []int{32767, 752, 1}
    var v11 uint = 487
    v11 = (v11 << ((v11) & 63))
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-558}
    var v3 bool = (("a" + "foo bar") >= "a")
    fmt.Println(v3)
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            v2 = append(v2, (fi12 >> ((1) & 63)))
            v14 := GenS1{F0: 270, F1: 127, F2: 2}
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 uint8 = 255
    v2 = append(v2, 4294967295)
    for fi16 := 0; fi16 < 20; fi16++ {
        if v3 {
            _ = len(v2)
            for fi17 := 0; fi17 < 14; fi17++ {
                v15 = v15
                v2 = append(v2, fi17)
                _ = fi17
            }
            v18 := GenS1{F0: 4294967295, F1: 127, F2: (v15 * 127)}
            _ = v18
        } else {
            fmt.Println(fi16, v15, v3)
            if v3 {
                v19 := GenS1{F0: 2147483647, F1: -1, F2: (^(v15))}
                var v20 int64 = genFn1(785)
                var v21 int = (fi16 % ((-842) | 1))
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 int8 = 0
            _ = v22
        }
        var v23 uint16 = uint16(v15)
        _ = fi16
        _ = v23
    }
    var v24 uint32 = 58
    fmt.Println(v24)
    switch "hello" {
    case "\n\t":
        for fi25 := 0; fi25 < 10; fi25++ {
            v3 = (1.5 >= (0.0 + -2.25))
            v15 -= 127
            _ = fi25
        }
    case "a":
        switch ((("foo bar" + "a") + (("a" + "") + "foo bar")) + ("hello" + "a")) {
        case "\n\t":
            if (-2.25 != (-2.25 + 1.5)) {
                var v26 bool = v3
                _ = v26
            }
        case "":
            for fi27 := 0; fi27 < 14; fi27++ {
                v28 := GenS2{F0: int16(v24), F1: 32767, F2: ("\n\t" == "")}
                switch fi27 {
                case 9223372036854775807:
                    switch "hello" {
                    case "":
                        v29 := GenS3{F0: 0, F1: (^(v15))}
                        _ = v29
                    case "\n\t":
                        var v30 uint32 = (^(v24))
                        _ = v30
                    case "foo bar":
                        var v31 uint32 = 2147483647
                        _ = v31
                    }
                case 32768:
                    fmt.Println(fi27)
                default:
                    var v32 bool = true
                    _ = v32
                }
                _ = fi27
                _ = v28
            }
        case "a":
            for fi33 := 0; fi33 < 10; fi33++ {
                if v3 {
                    if v3 {
                        if (genFn1(331) != 200) {
                            for fi34 := 0; fi34 < 13; fi34++ {
                                if v3 {
                                    if ((^(fi34)) == 2147483647) {
                                        var v35 bool = false
                                        _ = v35
                                    } else {
                                        v36 := GenS4{F0: (fi34 / ((-9223372036854775808) | 1)), F1: "\n\t"}
                                        _ = v36
                                    }
                                    var v37 int64 = 9223372036854775807
                                    _ = v37
                                }
                                switch fi33 {
                                case 2147483647:
                                    if v3 {
                                        var v38 int32 = 42
                                        fmt.Println(v38)
                                        v24 = 0
                                        v39 := GenS2{F0: int16(fi33), F1: 305, F2: v3}
                                        var v40 float32 = -2.25
                                        var v41 int64 = -1
                                        v2 = append(v2, (fi34 % (((^(fi33))) | 1)))
                                        v42 := GenS5{F0: -1, F1: (v24 | v24), F2: 0}
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                case -137:
                                    if (!true) {
                                        var v43 int32 = -883
                                        v44 := GenS6{F0: (fi34 | (-(fi33))), F1: (^(fi34)), F2: (v24 + (v24 & (v24 | v24)))}
                                        _ = v43
                                        _ = v44
                                    }
                                }
                                _ = len(v2)
                                delete(v1, "a")
                                _ = fi34
                            }
                            for fi45 := 0; fi45 < 6; fi45++ {
                                v15 = (v15 - (v15 - (v15 << ((0) & 63))))
                                v15 = v15
                                _ = fi45
                            }
                        } else {
                            v24 += 1
                            _ = len(v2)
                            var v46 int8 = -61
                            _ = v46
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        var v47 bool = true
                        var v48 int64 = 0
                        if v3 {
                            if false {
                                if v47 {
                                    v3 = v47
                                    for fi49 := 0; fi49 < 12; fi49++ {
                                        var v50 string = ("foo bar" + "foo bar")
                                        var v51 bool = v47
                                        _ = fi49
                                        _ = v50
                                        _ = v51
                                    }
                                }
                            } else {
                                v15 -= (v15 % (((v15 + (v15 | v15))) | 1))
                            }
                            fmt.Println(v47)
                            if (("hello" + (("" + "foo bar") + "")) >= (("hello" + ("hello" + "\n\t")) + "hello")) {
                                if (0.0 > 1.5) {
                                    v52 := GenS7{F0: (^(v48))}
                                    _ = v52
                                }
                                v1[(("hello" + (("foo bar" + "foo bar") + ("hello" + ("foo bar" + "hello")))) + "")] = (-(fi33))
                                var v53 int = fi33
                                _ = v53
                            }
                        }
                        _ = v47
                        _ = v48
                    }
                }
                if (!(!v3)) {
                    _ = len(v2)
                }
                var v54 uint64 = 194
                _ = fi33
                _ = v54
            }
        }
    case "":
        v3 = v3
    default:
        v15 = 208
    }
    v24 = 542
    v55 := GenS2{F0: -176, F1: 1, F2: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v55
}
