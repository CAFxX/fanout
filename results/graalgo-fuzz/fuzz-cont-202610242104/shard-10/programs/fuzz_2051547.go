package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint8 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 string = "hello"
    var v3 int16 = 0
    v4 := []uint32{65535, 0, 916}
    for fi5 := 0; fi5 < 20; fi5++ {
        switch v2 {
        case "\n\t":
            if true {
                var v6 int8 = 88
                v1 -= (v1 & (^(v1)))
                _ = v6
            }
        case "hello":
            if (!true) {
                var v7 uint64 = 1
                v8 := GenS1{F0: 2147483648}
                for fi9 := 0; fi9 < 18; fi9++ {
                    v10 := GenS2{F0: 884, F1: int8(v3)}
                    _ = fi9
                    _ = v10
                }
                _ = v7
                _ = v8
            }
        case "foo bar":
            v4 = append(v4, 0)
        default:
            var v11 uint16 = 0
            _ = v11
        }
        var v12 int32 = 2147483647
        _ = fi5
        _ = v12
    }
    v13 := GenS1{F0: 2147483648}
    if len(v4) > 0 {
        v4[0] = 42
    }
    if true {
        v1 += -724
    }
    var v14 int8 = -84
    switch v2 {
    case "\n\t":
        if true {
            if (3.14159 <= ((1.5 * 1e10) + float32(1.5))) {
                var v15 int8 = v14
                switch ("a" + v2) {
                case "hello":
                    v15 += v14
                case "":
                    for fi16 := 0; fi16 < 7; fi16++ {
                        var v17 float32 = 0.0
                        switch fi16 {
                        case 244:
                            var v18 bool = true
                            _ = v18
                        default:
                            if false {
                                v14 = -46
                            }
                        }
                        _ = fi16
                        _ = v17
                    }
                }
                var v19 float64 = 0.0
                _ = v15
                _ = v19
            } else {
                var v20 int8 = (v14 * (v14 ^ (v14 & (v14 & -111))))
                _ = v20
            }
            var v21 uint64 = 2
            if (0 > 32767) {
                var v22 uint = 0
                _ = v22
            } else {
                if true {
                    v3 -= 1000
                    v2 = v2
                } else {
                    var v23 int16 = 96
                    var v24 int32 = 2147483647
                    for fi25 := 0; fi25 < 17; fi25++ {
                        fmt.Println(v2, fi25)
                        _ = fi25
                    }
                    _ = v23
                    _ = v24
                }
                fmt.Println(v2, v2, v1)
            }
            _ = v21
        } else {
            if (true || ((0 == 127) || true)) {
                var v26 uint = 4294967295
                for fi27 := 0; fi27 < 18; fi27++ {
                    fmt.Println(v3, v1)
                    var v28 uint8 = 157
                    _ = fi27
                    _ = v28
                }
                if true {
                    switch "\n\t" {
                    case "\n\t":
                        for fi29 := 0; fi29 < 5; fi29++ {
                            var v30 uint32 = 996
                            v31 := GenS1{F0: (v30 ^ (v30 - v30))}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    default:
                        v2 = ("\n\t" + "foo bar")
                    }
                    fmt.Println(v1)
                } else {
                    for fi32 := 0; fi32 < 5; fi32++ {
                        if (false && false) {
                            var v33 string = (("foo bar" + v2) + v2)
                            _ = len(v4)
                            _ = v33
                        } else {
                            fmt.Println(v2)
                        }
                        _ = fi32
                    }
                }
                _ = v26
            }
            if false {
                v34 := GenS3{F0: 2}
                var v35 int = 32767
                v36 := GenS4{F0: (v1 ^ (v1 / ((9223372036854775807) | 1))), F1: 1}
                _ = v34
                _ = v35
                _ = v36
            }
        }
    case "foo bar":
        v3 += (v3 * (v3 & v3))
    case "":
        v37 := GenS1{F0: 0}
        _ = v37
    default:
        v38 := GenS1{F0: 65535}
        _ = v38
    }
    var v39 uint8 = 244
    if (v39 > 1) {
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 int8 = v14
            v39 = v39
            _ = fi40
            _ = v41
        }
        var v42 int64 = (v1 + (-(v1)))
        _ = v42
    } else {
        var v43 uint8 = (v39 | 128)
        var v44 int64 = v1
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
}
