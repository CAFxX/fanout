package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 int; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint; F1 int64; F2 int8 }
type GenS7 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS8 struct { F0 int64; F1 int32 }
type GenS9 struct { F0 uint8; F1 int16; F2 int64 }
type GenS10 struct { F0 int8; F1 int64 }
type GenS11 struct { F0 int }

func main() {
    var v1 uint32 = 65535
    var v2 int = 32768
    v3 := make(map[int]string)
    v4 := make(map[uint32]bool)
    v5 := []uint32{65535, (v1 % ((119) | 1))}
    v2 += (v2 ^ (v2 ^ (^(v2))))
    var v6 int64 = -9223372036854775808
    v7 := GenS1{F0: 141, F1: "foo bar", F2: "foo bar"}
    v6 += (v6 + (v6 * v6))
    var v8 uint8 = 172
    v6 = v6
    v9 := v4[v1]
    _ = len(v4)
    v10 := GenS1{F0: 0, F1: "a", F2: ""}
    v11 := GenS1{F0: 32767, F1: ("\n\t" + ("" + (("" + "") + ((("\n\t" + ("a" + ("" + "hello"))) + "foo bar") + "\n\t")))), F2: ""}
    for fi12 := 0; fi12 < 14; fi12++ {
        switch (v2 - v2) {
        case -2147483648:
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS2{F0: fi12, F1: 2}
                _ = fi13
                _ = v14
            }
        case 2147483647:
            v1 -= (v1 % (((v1 - 267)) | 1))
        case 245:
            delete(v3, fi12)
        default:
            if v9 {
                if v9 {
                    v15 := GenS3{F0: 883, F1: 2}
                    if (v9 && (0.0 != 1e10)) {
                        switch (fi12 >> ((65535) & 63)) {
                        case 127:
                            var v16 uint32 = 2
                            _ = v16
                        default:
                            if (0 == 32767) {
                                fmt.Println(v8, v6)
                                var v17 int32 = 0
                                for fi18 := 0; fi18 < 1; fi18++ {
                                    if v9 {
                                        var v19 uint = 255
                                        fmt.Println(v2, fi12)
                                        v20 := GenS4{F0: "hello", F1: 128}
                                        _ = v19
                                        _ = v20
                                    }
                                    var v21 int16 = -928
                                    _ = fi18
                                    _ = v21
                                }
                                _ = v17
                            } else {
                                for fi22 := 0; fi22 < 2; fi22++ {
                                    v23 := GenS5{F0: 32768, F1: 128, F2: 42}
                                    _ = fi22
                                    _ = v23
                                }
                                switch fi12 {
                                case 1:
                                    var v24 uint32 = (v1 / ((706) | 1))
                                    _ = v24
                                case 32767:
                                    _ = len(v5)
                                case 979:
                                    var v25 int8 = 42
                                    _ = v25
                                default:
                                    if (true || (!v9)) {
                                        v9 = v9
                                        v1 += (v1 ^ v1)
                                        var v26 int64 = (v6 | v6)
                                        _ = v26
                                    } else {
                                        v27 := GenS5{F0: (fi12 | 9223372036854775807), F1: v8, F2: -65}
                                        _ = v27
                                    }
                                }
                                v28 := GenS6{F0: 255, F1: 9223372036854775807, F2: 100}
                                _ = v28
                            }
                        }
                        if true {
                            v6 = (-(v6))
                            if (!true) {
                                var v29 float32 = 1e10
                                _ = v29
                            }
                            if (v9 || (-1 >= -1)) {
                                _ = len(v5)
                            }
                        } else {
                            fmt.Println(v1, v8, v2)
                            var v30 uint8 = v8
                            for fi31 := 0; fi31 < 18; fi31++ {
                                v2 = 1
                                var v32 uint32 = 997
                                var v33 uint64 = 42
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            _ = v30
                        }
                        var v34 int64 = v6
                        _ = v34
                    } else {
                        var v35 int16 = -32768
                        _ = v35
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v15
                }
                switch (v2 % (((-(v2))) | 1)) {
                case 2147483647:
                    fmt.Println(v2)
                case 4294967295:
                    v36 := GenS2{F0: v2, F1: 4294967295}
                    _ = v36
                default:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
        }
        for fi37 := 0; fi37 < 1; fi37++ {
            v38 := GenS7{F0: (false && true), F1: 0, F2: 0}
            _ = fi37
            _ = v38
        }
        if v9 {
            for fi39 := 0; fi39 < 5; fi39++ {
                v2 += -381
                v40 := GenS8{F0: int64(v8), F1: int32(v6)}
                if len(v5) > 0 {
                    v5[0] = 42
                }
                _ = fi39
                _ = v40
            }
            v8 -= v8
            v5 = append(v5, v1)
        }
        _ = fi12
    }
    if ((2147483647 != -662) || false) {
        switch "" {
        case "a":
            for fi41 := 0; fi41 < 5; fi41++ {
                v42 := GenS9{F0: 128, F1: 32767, F2: 9223372036854775807}
                for fi43 := 0; fi43 < 15; fi43++ {
                    v8 += v8
                    _ = fi43
                }
                var v44 bool = v9
                _ = fi41
                _ = v42
                _ = v44
            }
        case "hello":
            v45 := GenS10{F0: -128, F1: v6}
            _ = v45
        }
        var v46 float32 = 0.0
        var v47 int32 = 568
        _ = v46
        _ = v47
    } else {
        v1 += v1
    }
    if (6 > -1) {
        if (-2147483648 == -1) {
            var v48 uint = 18446744073709551615
            _ = v48
        } else {
            if (-2147483648 > 2147483647) {
                var v49 uint64 = 0
                v50 := GenS11{F0: v2}
                v2 = v2
                _ = v49
                _ = v50
            }
            fmt.Println(v6)
        }
        v9 = (!(true || (-577 <= -108)))
        fmt.Println(v6)
    }
    fmt.Println(v2, v1)
    v6 -= (v6 % ((9223372036854775807) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
