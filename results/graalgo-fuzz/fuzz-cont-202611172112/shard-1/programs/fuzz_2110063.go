package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int }
type GenS2 struct { F0 int32; F1 int16; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int32; F2 int8 }
type GenS5 struct { F0 int; F1 int32 }

func main() {
    var v1 int16 = 14
    v2 := []string{"\n\t", "\n\t", "\n\t", ("foo bar" + "foo bar")}
    var v3 uint16 = 32767
    var v4 bool = (34 != 653)
    v3 += v3
    switch "\n\t" {
    case "\n\t":
        v1 = v1
    case "":
        v4 = (127 == 729)
    case "foo bar":
        if v4 {
            var v5 int8 = -128
            _ = v5
        }
    default:
        if false {
            if (!((!false) || v4)) {
                var v6 int32 = -1
                _ = v6
            }
        } else {
            v3 -= (v3 ^ v3)
        }
    }
    switch "a" {
    case "a":
        var v7 int16 = (v1 >> ((0) & 63))
        _ = v7
    case "":
        if true {
            v8 := GenS1{F0: 619, F1: 42, F2: 128}
            for fi9 := 0; fi9 < 10; fi9++ {
                v2 = append(v2, "\n\t")
                _ = fi9
            }
            _ = v8
        }
    default:
        v3 = (v3 | v3)
    }
    v10 := GenS1{F0: -9223372036854775808, F1: -317, F2: 9223372036854775807}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v4 {
        if len(v2) > 0 {
            v2[0] = ((("" + "hello") + "\n\t") + (("foo bar" + "foo bar") + "a"))
        }
        v3 = v3
        var v11 uint8 = 1
        _ = v11
    }
    if (-9223372036854775808 != 0) {
        if (!true) {
            var v12 int16 = -1
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            _ = v12
        } else {
            _ = len(v2)
            var v13 uint64 = 18446744073709551615
            if (true && ((927 < 4294967295) || (0 > 128))) {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS2{F0: -1, F1: (v1 << ((400) & 63)), F2: "hello"}
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v1 -= 1
                        _ = fi16
                    }
                    switch "a" {
                    case "":
                        for fi17 := 0; fi17 < 11; fi17++ {
                            if v4 {
                                v2 = append(v2, (("foo bar" + "") + "a"))
                            } else {
                                v18 := GenS3{F0: 2}
                                _ = len(v2)
                                v3 += v3
                                _ = v18
                            }
                            var v19 uint32 = 42
                            _ = fi17
                            _ = v19
                        }
                    default:
                        for fi20 := 0; fi20 < 3; fi20++ {
                            var v21 int32 = 2147483647
                            _ = fi20
                            _ = v21
                        }
                    }
                    _ = fi14
                    _ = v15
                }
            } else {
                v4 = (2 == 18446744073709551615)
                if (!(("\n\t" + (("a" + "hello") + "foo bar")) >= "\n\t")) {
                    fmt.Println(v4, v3, v4)
                    v1 = v1
                    var v22 float32 = (float32(0.0) * 0.0)
                    _ = v22
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (false && true) {
                        switch 42 {
                        case -1:
                            fmt.Println(v1)
                        case 42:
                            v23 := GenS4{F0: v4, F1: -733, F2: -59}
                            _ = v23
                        case 947:
                            fmt.Println(v1, v4, v13)
                        }
                        v24 := GenS5{F0: 42, F1: 2147483647}
                        var v25 int16 = v1
                        _ = v24
                        _ = v25
                    }
                }
                if v4 {
                    v26 := GenS4{F0: (v4 && (v4 || v4)), F1: -904, F2: 121}
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 7; fi27++ {
                        var v28 int = fi27
                        _ = fi27
                        _ = v28
                    }
                    var v29 uint = 18446744073709551615
                    _ = v29
                }
            }
            _ = v13
        }
        if v4 {
            var v30 uint8 = 0
            v1 -= (^(v1))
            _ = v30
        } else {
            for fi31 := 0; fi31 < 1; fi31++ {
                if true {
                    if (v4 || (!(!(!v4)))) {
                        var v32 uint64 = 952
                        v3 += v3
                        _ = v32
                    }
                    var v33 int16 = v1
                    v3 = (v3 << ((855) & 63))
                    _ = v33
                } else {
                    if (!(610 != 4294967295)) {
                        v4 = ((!(v3 > (v3 ^ (v3 + 65535)))) || false)
                        v34 := GenS1{F0: 613, F1: 9223372036854775807, F2: (fi31 ^ fi31)}
                        _ = v34
                    }
                    switch fi31 {
                    case 866:
                        var v35 uint16 = (v3 + 255)
                        _ = v35
                    }
                }
                _ = fi31
            }
            var v36 int64 = 42
            _ = v36
        }
    } else {
        v4 = v4
    }
    var v37 int = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
