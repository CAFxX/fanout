package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int16 }
func genFn0() uint8 {
    var v5 int64 = 724
    v6 := GenS1{F0: 65535, F1: 782, F2: (-(v5))}
    _ = v5
    _ = v6
    return 128
}

func main() {
    v1 := []uint32{2147483648, 1}
    var v2 int16 = 32767
    var v3 int8 = 13
    var v4 int8 = (v3 >> ((255) & 63))
    v7 := GenS2{F0: 18446744073709551615, F1: 0}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (-909 > 1) {
        if (!(892 < 511)) {
            var v8 int8 = v4
            _ = v8
        }
        fmt.Println(v3, v3)
        v9 := GenS2{F0: 2, F1: -9223372036854775808}
        _ = v9
    } else {
        switch "\n\t" {
        case "a":
            if false {
                if true {
                    for fi10 := 0; fi10 < 13; fi10++ {
                        for fi11 := 0; fi11 < 13; fi11++ {
                            if false {
                                var v12 int64 = 9223372036854775807
                                switch (fi11 % (((fi10 % ((fi11) | 1))) | 1)) {
                                case 644:
                                    var v13 bool = true
                                    _ = v13
                                case 32768:
                                    v1 = append(v1, 4294967295)
                                    v4 += v4
                                case -149:
                                    var v14 int64 = v12
                                    _ = v14
                                }
                                _ = v12
                            } else {
                                var v15 int32 = -530
                                v4 += v4
                                v15 += (-(v15))
                                _ = v15
                            }
                            v3 += (v4 << ((988) & 63))
                            _ = fi11
                        }
                        v16 := GenS3{F0: 2147483648}
                        var v17 uint8 = 17
                        _ = fi10
                        _ = v16
                        _ = v17
                    }
                } else {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        if (((1.5 + 0.1) > float64(1.5)) || (!((("a" != "hello") && (-2147483648 < 251)) && (v3 == v4)))) {
                            _ = len(v1)
                            v2 += v2
                            if false {
                                v19 := GenS4{F0: 1, F1: v2}
                                _ = v19
                            } else {
                                v20 := GenS4{F0: (^(v2)), F1: v2}
                                switch ("hello" + ("" + "\n\t")) {
                                case "foo bar":
                                    v2 += 1000
                                case "a":
                                    v4 = (^(v4))
                                case "":
                                    var v21 float32 = -2.25
                                    _ = v21
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = 49
                                    }
                                }
                                _ = v20
                            }
                        } else {
                            for fi22 := 0; fi22 < 9; fi22++ {
                                _ = len(v1)
                                _ = fi22
                            }
                            v4 += (-(v4))
                            if true {
                                v23 := GenS3{F0: 458}
                                _ = v23
                            } else {
                                switch ("foo bar" + "hello") {
                                case "foo bar":
                                    v24 := GenS2{F0: 18446744073709551615, F1: 544}
                                    _ = v24
                                default:
                                    var v25 int16 = v2
                                    _ = v25
                                }
                            }
                        }
                        v3 += (v4 >> ((42) & 63))
                        _ = fi18
                    }
                }
            }
        default:
            v26 := GenS1{F0: 2, F1: 65535, F2: 0}
            _ = v26
        }
        if (42 != 0) {
            fmt.Println(v4)
        } else {
            v3 += (v4 - -1)
            v4 -= v3
        }
    }
    var v27 uint16 = 42
    if len(v1) > 0 {
        v1[0] = 0
    }
    v27 -= v27
    v28 := GenS1{F0: 0, F1: 51, F2: 9223372036854775807}
    if ((v3 / ((-1) | 1)) >= (-(v4))) {
        switch "hello" {
        case "":
            if len(v1) > 0 {
                v1[0] = 391
            }
        }
        v3 -= (v3 >> ((18446744073709551615) & 63))
        var v29 bool = false
        _ = v29
    } else {
        v30 := GenS3{F0: 65535}
        var v31 int16 = (v2 - (v2 & v2))
        _ = v30
        _ = v31
    }
    v32 := GenS3{F0: 1}
    v2 = -1
    var v33 int16 = (v2 * v2)
    for fi34 := 0; fi34 < 20; fi34++ {
        _ = len(v1)
        fmt.Println(v27, v3)
        fmt.Println(v27, fi34)
        _ = fi34
    }
    v27 = v27
    v35 := GenS3{F0: 2147483647}
    v36 := GenS2{F0: 0, F1: 245}
    fmt.Println(v2, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v4
    _ = v7
}
