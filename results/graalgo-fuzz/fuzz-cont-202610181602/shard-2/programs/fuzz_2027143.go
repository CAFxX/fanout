package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 string; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint32 = 274
    v2 := []int{}
    v3 := []uint32{v1, 109}
    switch -2147483648 {
    case 9223372036854775807:
        v1 += (v1 ^ (v1 >> ((4294967295) & 63)))
    }
    var v4 int16 = 0
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 float64 = float64(float32(0.0))
        _ = fi5
        _ = v6
    }
    var v7 int32 = -991
    v1 += v1
    var v8 uint64 = 0
    v7 = (v7 << ((42) & 63))
    var v9 float32 = 1e10
    if (v7 > (v7 + v7)) {
        v4 = (v4 - v4)
        var v10 int16 = 32767
        _ = v10
    } else {
        fmt.Println(v9)
        switch -9223372036854775808 {
        case -212:
            if ((v1 & v1) == (v1 & 681)) {
                switch 859 {
                case 127:
                    if (32767 == 532) {
                        v11 := GenS1{F0: ""}
                        _ = len(v2)
                        var v12 bool = (0 == 77)
                        _ = v11
                        _ = v12
                    }
                case 839:
                    if ((2 > 2) && true) {
                        for fi13 := 0; fi13 < 9; fi13++ {
                            for fi14 := 0; fi14 < 15; fi14++ {
                                var v15 int8 = int8(v7)
                                var v16 bool = (v8 >= v8)
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            _ = fi13
                        }
                        var v17 uint8 = 1
                        _ = v17
                    } else {
                        if (v8 > v8) {
                            if ((v7 / ((v7) | 1)) > int32(v8)) {
                                if (-1 >= 137) {
                                    fmt.Println(v1)
                                } else {
                                    v9 += 3.14159
                                    switch 456 {
                                    case -1:
                                        for fi18 := 0; fi18 < 18; fi18++ {
                                            v4 -= (v4 & (-(v4)))
                                            var v19 int = (fi18 >> ((18446744073709551615) & 63))
                                            for fi20 := 0; fi20 < 16; fi20++ {
                                                v7 = v7
                                                _ = fi20
                                            }
                                            _ = fi18
                                            _ = v19
                                        }
                                    case 32768:
                                        if true {
                                            fmt.Println(v9, v1, v1)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    case -239:
                                        if (("foo bar" + ("hello" + ((("a" + "foo bar") + "\n\t") + (((((((((("a" + (("hello" + "a") + "a")) + "") + ("\n\t" + (("hello" + "hello") + "a"))) + "hello") + "\n\t") + ("a" + "\n\t")) + ("a" + "foo bar")) + "") + ("hello" + "\n\t")) + "foo bar")))) < ("foo bar" + "\n\t")) {
                                            var v21 uint64 = (^(v8))
                                            _ = v21
                                        }
                                    }
                                }
                            }
                            v22 := GenS1{F0: (("\n\t" + "foo bar") + "")}
                            _ = v22
                        }
                        v23 := GenS1{F0: "\n\t"}
                        _ = v23
                    }
                case 4294967295:
                    for fi24 := 0; fi24 < 15; fi24++ {
                        if len(v2) > 0 {
                            v2[0] = fi24
                        }
                        _ = fi24
                    }
                }
                if len(v3) > 0 {
                    v3[0] = 1
                }
            }
        case 852:
            if true {
                var v25 float32 = v9
                v4 -= v4
                v2 = append(v2, 632)
                _ = v25
            }
        case 195:
            var v26 uint64 = v8
            _ = v26
        }
        if false {
            if false {
                fmt.Println(v1, v9, v9)
                var v27 uint64 = v8
                if true {
                    v28 := GenS2{F0: ((false && (-823 == 1)) || (false || false)), F1: 255, F2: v8}
                    _ = v28
                } else {
                    v29 := GenS3{F0: 761}
                    _ = v29
                }
                _ = v27
            } else {
                v7 = 1
            }
            v30 := GenS4{F0: "a"}
            if (42 > -2147483648) {
                var v31 uint32 = v1
                if (42 == 42) {
                    v31 += v31
                    v32 := GenS1{F0: "hello"}
                    _ = v32
                }
                _ = v31
            }
            _ = v30
        } else {
            v33 := GenS5{F0: 18446744073709551615, F1: "a", F2: (^(v1))}
            v4 = v4
            _ = v33
        }
    }
    if true {
        for fi34 := 0; fi34 < 11; fi34++ {
            _ = len(v3)
            for fi35 := 0; fi35 < 16; fi35++ {
                var v36 uint16 = 255
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
        v8 -= v8
        var v37 bool = true
        _ = v37
    }
    v38 := GenS6{F0: false, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
