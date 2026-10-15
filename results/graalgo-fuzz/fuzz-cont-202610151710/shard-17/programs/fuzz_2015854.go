package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint8; F2 bool }
type GenS4 struct { F0 int16; F1 uint8; F2 int64 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 int }

func main() {
    var v1 bool = (0 < 233)
    v2 := []string{}
    var v3 uint64 = 1
    v4 := make(map[int]bool)
    var v5 int64 = 928
    var v6 bool = v1
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 10; fi8++ {
        if (42 <= v3) {
            var v9 int64 = v5
            v10 := GenS1{F0: -2147483648, F1: 206}
            v3 -= 1
            _ = v10
            _ = v9
        }
        if (!v6) {
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := GenS1{F0: 1, F1: 128}
                fmt.Println(v3, fi8, v3)
                _ = fi11
                _ = v12
            }
            v5 = v5
            for fi13 := 0; fi13 < 2; fi13++ {
                fmt.Println(fi13)
                var v14 uint64 = 330
                _ = fi13
                _ = v14
            }
        }
        if false {
            fmt.Println(fi8, v5)
            v3 -= 42
        } else {
            v7 = ("a" + ("a" + ("" + v7)))
            fmt.Println(v5)
        }
        _ = fi8
    }
    if (18446744073709551615 > 65535) {
        fmt.Println(v3)
        fmt.Println(v7, v3)
    } else {
        if (v1 || v1) {
            if (((0 <= 149) && v1) || (184 < 65535)) {
                v3 = (^(v3))
            } else {
                var v15 int16 = 1000
                for fi16 := 0; fi16 < 18; fi16++ {
                    for fi17 := 0; fi17 < 6; fi17++ {
                        if v6 {
                            v5 -= (v5 | (v5 >> ((65535) & 63)))
                        }
                        var v18 int8 = 42
                        v5 = (v5 & int64(v3))
                        _ = fi17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v20 := GenS2{F0: 42, F1: true, F2: 65535}
                        v4[(fi16 % ((fi19) | 1))] = (v1 || (867 != 42))
                        if (!v1) {
                            if (-868 != -308) {
                                v21 := GenS2{F0: 128, F1: v1, F2: 65535}
                                v15 += v15
                                if ((2 <= 42) || v1) {
                                    fmt.Println(v6, v7, fi19)
                                    fmt.Println(v15, v3)
                                    var v22 uint64 = v3
                                    _ = v22
                                } else {
                                    v7 = v7
                                }
                                _ = v21
                            } else {
                                var v23 int = fi16
                                _ = v23
                            }
                            v1 = v1
                        } else {
                            if (int8(v3) < 0) {
                                _ = len(v2)
                            }
                        }
                        _ = fi19
                        _ = v20
                    }
                    _ = fi16
                }
                if (((!false) && true) || false) {
                    v24 := GenS2{F0: 127, F1: ((((1.5 * 0.1) + -2.25) * -2.25) <= 3.14159), F2: 971}
                    _ = v24
                }
                _ = v15
            }
            if (!(!(int8(v5) <= 20))) {
                var v25 uint8 = 255
                _ = v25
            } else {
                v26 := GenS3{F0: 0, F1: 127, F2: v1}
                _ = len(v2)
                _ = v26
            }
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 uint64 = (v3 ^ v3)
                v5 -= v5
                _ = fi27
                _ = v28
            }
        }
        v7 = "\n\t"
        v5 += (v5 << ((4294967295) & 63))
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 int64 = v5
            if (!(42 == -509)) {
                var v32 int8 = -1
                v33 := GenS3{F0: -1, F1: 2, F2: (1000 > -32768)}
                _ = v32
                _ = v33
            } else {
                var v34 int16 = 1000
                v35 := GenS4{F0: 0, F1: 127, F2: -9223372036854775808}
                _ = v34
                _ = v35
            }
            v4[fi29] = false
            _ = fi30
            _ = v31
        }
        var v36 uint = 1
        delete(v4, (fi29 - fi29))
        _ = fi29
        _ = v36
    }
    v37 := GenS3{F0: 423, F1: 88, F2: v1}
    v38 := GenS5{F0: int32(v5), F1: 24}
    var v39 bool = v1
    var v40 uint8 = 1
    fmt.Println(v3, v6)
    v40 = 2
    fmt.Println(v40, v7, v40)
    var v41 uint32 = 0
    v3 += (v3 >> ((255) & 63))
    switch ("a" + v7) {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        v42 := GenS3{F0: -32768, F1: (v40 << ((881) & 63)), F2: (!false)}
        _ = v42
    default:
        fmt.Println(v40, v1, v7)
    }
    switch ((v7 + "hello") + "\n\t") {
    case "a":
        switch int(v41) {
        case 1:
            v43 := GenS6{F0: -1}
            _ = v43
        }
    }
    var v44 int16 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
