package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint32; F2 int64 }
type GenS3 struct { F0 int; F1 string; F2 int32 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint8, p1 int16, p2 int32) int64 {
    var v5 int64 = int64(p2)
    v6 := GenS1{F0: p0}
    v7 := GenS1{F0: p0}
    if (("hello" <= (("foo bar" + ("hello" + "")) + ("" + "hello"))) && (!false)) {
        v5 = (^(v5))
        var v8 float64 = (1e10 / float64(0.1))
        var v9 int32 = p2
        _ = v8
        _ = v9
    } else {
        var v10 int64 = v5
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (v5 << ((42) & 63))
}
func genFn1(p0 int8, p1 int8) uint {
    v11 := GenS2{F0: 0, F1: 2147483647, F2: 1}
    v12 := []int{32767, 0, 32767}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 int64 = 1
    v2 := []string{"a", ""}
    v3 := []string{"a", "a", ("" + "\n\t")}
    var v4 int = -9223372036854775808
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v13 := GenS3{F0: (v4 / (((v4 | (v4 + (-(v4))))) | 1)), F1: "hello", F2: -296}
    var v14 int = 42
    if len(v2) > 0 {
        v2[0] = ((("foo bar" + "") + ("hello" + ("a" + "a"))) + "")
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        if true {
            v4 = (fi15 >> ((4294967295) & 63))
            v4 -= v14
            v1 -= v1
        }
        if (true && false) {
            v14 += fi15
            for fi16 := 0; fi16 < 12; fi16++ {
                v17 := GenS1{F0: 1}
                _ = fi16
                _ = v17
            }
            if false {
                if (156 > 855) {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v19 := GenS3{F0: fi15, F1: "\n\t", F2: 636}
                        _ = fi18
                        _ = v19
                    }
                    v20 := GenS2{F0: v14, F1: 536, F2: v1}
                    v4 -= v14
                    _ = v20
                }
                v14 -= v14
                v21 := GenS4{F0: -642, F1: ((!(("hello" + "foo bar") > (("foo bar" + ("a" + ((("hello" + "\n\t") + "") + "foo bar"))) + ("a" + "")))) && (1 < 98))}
                _ = v21
            } else {
                if (1 != 255) {
                    v22 := GenS5{F0: (v1 + v1), F1: 1}
                    _ = v22
                }
                v23 := GenS6{F0: 255}
                _ = v23
            }
        } else {
            if (1e10 == 3.14159) {
                _ = len(v3)
                var v24 uint16 = 65535
                if ((^(v4)) >= v14) {
                    for fi25 := 0; fi25 < 19; fi25++ {
                        if true {
                            var v26 uint = 4294967295
                            switch v4 {
                            case 128:
                                if true {
                                    v24 += (v24 << ((2) & 63))
                                    if false {
                                        v27 := GenS3{F0: fi25, F1: ("" + "foo bar"), F2: -1}
                                        _ = v27
                                    } else {
                                        var v28 uint = (v26 % ((v26) | 1))
                                        var v29 uint16 = 872
                                        _ = v28
                                        _ = v29
                                    }
                                }
                            case 32768:
                                v30 := GenS1{F0: 187}
                                _ = v30
                            case 9223372036854775807:
                                if true {
                                    var v31 int8 = 1
                                    v26 += 286
                                    var v32 int = 127
                                    switch (("a" + "hello") + "\n\t") {
                                    case "hello":
                                        var v33 int16 = -1
                                        fmt.Println(v1)
                                        _ = v33
                                    default:
                                        fmt.Println(v14)
                                    }
                                    v34 := GenS5{F0: (^(v1)), F1: 0}
                                    _ = v31
                                    _ = v32
                                    _ = v34
                                } else {
                                    var v35 string = ""
                                    v26 += (v26 - 18446744073709551615)
                                    _ = v35
                                }
                            }
                            _ = v26
                        }
                        var v36 int16 = int16(v24)
                        _ = fi25
                        _ = v36
                    }
                } else {
                    var v37 uint8 = 42
                    _ = v37
                }
                _ = v24
            } else {
                if len(v3) > 0 {
                    v3[0] = ""
                }
                v14 = 2147483647
            }
        }
        v4 = (v4 & 128)
        _ = fi15
    }
    var v38 int64 = (v1 | 42)
    var v39 int32 = 99
    fmt.Println(v1, v14, v1)
    var v40 uint16 = 65535
    v41 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
