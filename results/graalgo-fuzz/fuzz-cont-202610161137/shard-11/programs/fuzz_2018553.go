package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint64; F1 int64; F2 int16 }
type GenS6 struct { F0 int32; F1 uint8 }

func main() {
    v1 := []uint32{756, 42, 452}
    var v2 uint64 = 1
    var v3 int32 = -89
    var v4 int32 = v3
    var v5 int64 = 476
    v6 := make(map[uint32]string)
    v3 = 1
    if false {
        if false {
            fmt.Println(v5)
            var v7 uint8 = 127
            _ = v7
        } else {
            v3 -= (-(v3))
            fmt.Println(v5)
        }
        v2 += v2
    } else {
        v8 := GenS1{F0: 42}
        v9 := v6[304]
        v10 := GenS1{F0: 100}
        _ = v10
        _ = v8
        _ = v9
    }
    switch ("hello" + "foo bar") {
    case "hello":
        v11 := GenS1{F0: 100}
        _ = v11
    default:
        if false {
            for fi12 := 0; fi12 < 13; fi12++ {
                v3 += (v4 - v3)
                fmt.Println(v5, v4)
                for fi13 := 0; fi13 < 8; fi13++ {
                    switch "hello" {
                    case "":
                        if false {
                            var v14 float32 = (float32(0.1) / 1.0)
                            var v15 bool = (("" + ("hello" + ("foo bar" + ("a" + "foo bar")))) != "a")
                            var v16 uint8 = 42
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                        delete(v6, 65535)
                        var v17 int16 = -384
                        _ = v17
                    default:
                        if ((!true) && (v2 != (v2 - 42))) {
                            var v18 float32 = 0.1
                            v19 := GenS2{F0: -118, F1: 65535}
                            var v20 int64 = 42
                            if (true && true) {
                                var v21 float64 = 0.1
                                _ = v21
                            } else {
                                var v22 int16 = 1000
                                var v23 float64 = 0.1
                                _ = v22
                                _ = v23
                            }
                            _ = v18
                            _ = v19
                            _ = v20
                        } else {
                            var v24 uint = 4294967295
                            if true {
                                v3 -= 1
                                var v25 uint16 = 1
                                var v26 uint = (^(v24))
                                _ = v25
                                _ = v26
                            }
                            _ = v24
                        }
                    }
                    _ = len(v6)
                    _ = fi13
                }
                _ = fi12
            }
        }
    }
    v27 := GenS2{F0: -44, F1: 1}
    var v28 int = 32768
    fmt.Println(v3, v5)
    v29 := GenS1{F0: 42}
    for fi30 := 0; fi30 < 6; fi30++ {
        for fi31 := 0; fi31 < 8; fi31++ {
            v4 = v3
            v28 = (fi30 / (((fi31 >> ((4294967295) & 63))) | 1))
            v32 := GenS3{F0: 42}
            var v33 float32 = (3.14159 * float32(1e10))
            var v34 uint64 = (v2 ^ 991)
            switch "a" {
            case "a":
                fmt.Println(v5, fi30)
                fmt.Println(v4)
                fmt.Println(v5)
            case "\n\t":
                v5 = -1
            }
            v5 += -1
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        if true {
            if true {
                fmt.Println(v5, v28)
                if true {
                    v35 := GenS2{F0: 1, F1: 974}
                    v36 := GenS4{F0: 65535, F1: 917, F2: v4}
                    _ = v35
                    _ = v36
                } else {
                    if false {
                        v37 := GenS5{F0: v2, F1: (v5 << ((18446744073709551615) & 63)), F2: 32767}
                        v38 := GenS1{F0: -127}
                        _ = v37
                        _ = v38
                    }
                    v2 = (v2 | (v2 << ((4294967295) & 63)))
                    if (1 < -1) {
                        v28 += (fi30 << ((0) & 63))
                    }
                }
            }
            switch "hello" {
            case "a":
                for fi39 := 0; fi39 < 1; fi39++ {
                    v40 := GenS5{F0: v2, F1: -9223372036854775808, F2: -858}
                    v4 -= v4
                    _ = len(v6)
                    _ = fi39
                    _ = v40
                }
            case "":
                _ = len(v1)
            }
            if false {
                switch (fi30 & (fi30 - v28)) {
                case 4294967295:
                    var v41 uint = 18446744073709551615
                    _ = v41
                }
            } else {
                v42 := v6[2]
                for fi43 := 0; fi43 < 12; fi43++ {
                    switch 32768 {
                    case 9223372036854775807:
                        var v44 uint64 = (v2 % ((1) | 1))
                        fmt.Println(v2)
                        switch -1 {
                        case -1:
                            fmt.Println(v5)
                        default:
                            fmt.Println(fi43, v4)
                        }
                        _ = v44
                    default:
                        v2 -= 2
                    }
                    _ = len(v6)
                    if ((v2 & v2) < 18446744073709551615) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v45 := GenS5{F0: (v2 | v2), F1: 1, F2: 42}
                        _ = v45
                    } else {
                        var v46 int = 900
                        v47 := GenS6{F0: (v4 % ((-1) | 1)), F1: 1}
                        _ = v46
                        _ = v47
                    }
                    var v48 float64 = ((1e10 / float64(-2.25)) * 3.14159)
                    _ = fi43
                    _ = v48
                }
                v3 += v3
                _ = v42
            }
        }
        for fi49 := 0; fi49 < 6; fi49++ {
            v4 -= (v3 % ((-1) | 1))
            fmt.Println(fi30, fi30, v4)
            var v50 float64 = (-2.25 - (-2.25 - (1e10 - ((float64(float32(1.5)) + -2.25) * (1e10 - -2.25)))))
            _ = fi49
            _ = v50
        }
        _ = fi30
    }
    var v51 int64 = int64(v28)
    var v52 uint16 = 0
    var v53 int = v28
    v3 = (v3 | 897)
    if true {
        var v54 float64 = -2.25
        _ = v54
    } else {
        v55 := GenS1{F0: 0}
        v52 = (v52 + (^(v52)))
        v56 := GenS2{F0: 42, F1: 925}
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 8; fi57++ {
        v4 -= (-(v4))
        var v58 uint32 = 2147483648
        _ = fi57
        _ = v58
    }
    var v59 int64 = (v51 / (((v5 | int64(v53))) | 1))
    fmt.Println(v59)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v59
    _ = v6
}
