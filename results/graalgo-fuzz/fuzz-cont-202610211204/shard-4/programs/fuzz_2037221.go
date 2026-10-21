package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int32 }
type GenS2 struct { F0 int8; F1 uint8; F2 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 65535
    var v3 uint = v2
    var v4 uint32 = 42
    v5 := []string{"foo bar", ("hello" + "a"), ("a" + "a")}
    if true {
        fmt.Println(v3, v2)
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 int8 = 42
            _ = fi6
            _ = v7
        }
        if false {
            var v8 int8 = -16
            v3 = 650
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := v1["hello"]
                _ = fi9
                _ = v10
            }
            _ = v8
        }
    }
    v4 += 65535
    fmt.Println(v3)
    if (float32(0.0) == 1e10) {
        fmt.Println(v2, v4)
        for fi11 := 0; fi11 < 1; fi11++ {
            _ = len(v5)
            _ = fi11
        }
        var v12 bool = (false || ("\n\t" >= ""))
        _ = v12
    } else {
        v4 = v4
        v3 = v3
        if true {
            var v13 float32 = float32(-2.25)
            _ = v13
        }
    }
    switch 127 {
    case -9223372036854775808:
        var v14 uint64 = 18446744073709551615
        _ = v14
    case -2147483648:
        var v15 uint16 = 255
        _ = v15
    default:
        v2 -= (v2 - (v3 & (v2 | (v2 & v3))))
    }
    v16 := GenS1{F0: -1, F1: ("a" + ("" + ("a" + (("hello" + "\n\t") + "foo bar")))), F2: 0}
    var v17 int16 = 42
    v4 = v4
    for fi18 := 0; fi18 < 1; fi18++ {
        _ = len(v5)
        v2 -= (v2 + v3)
        if (42 == 1) {
            var v19 float64 = 0.0
            v4 -= 28
            _ = v19
        }
        _ = fi18
    }
    var v20 uint32 = v4
    var v21 int16 = 258
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS2{F0: 100, F1: 119, F2: (!false)}
        _ = fi22
        _ = v23
    }
    var v24 uint8 = 1
    v24 -= 209
    v3 = 512
    if false {
        if true {
            if true {
                v3 += v2
            }
            var v25 bool = (42 > -9223372036854775808)
            var v26 uint64 = 862
            _ = v25
            _ = v26
        } else {
            v24 = v24
            for fi27 := 0; fi27 < 4; fi27++ {
                fmt.Println(v2, fi27)
                _ = fi27
            }
            v3 = uint(v20)
        }
        if true {
            fmt.Println(v2)
            if true {
                if (2 == 65535) {
                    v3 -= (v2 + v2)
                    if true {
                        _ = len(v1)
                        if (1 > 42) {
                            v28 := GenS1{F0: 1, F1: "", F2: -330}
                            if (71 != -128) {
                                var v29 uint64 = 4294967295
                                for fi30 := 0; fi30 < 7; fi30++ {
                                    var v31 uint64 = (v29 * v29)
                                    _ = fi30
                                    _ = v31
                                }
                                _ = v29
                            } else {
                                switch "a" {
                                case "a":
                                    switch "\n\t" {
                                    case "a":
                                        v32 := GenS2{F0: 1, F1: (v24 + v24), F2: (2 >= uint64(v2))}
                                        _ = v32
                                    case "hello":
                                        var v33 uint16 = 566
                                        _ = v33
                                    default:
                                        _ = len(v5)
                                    }
                                case "foo bar":
                                    var v34 int16 = v17
                                    _ = v34
                                }
                            }
                            _ = v28
                        }
                        v17 -= 42
                    }
                    for fi35 := 0; fi35 < 16; fi35++ {
                        var v36 uint8 = (v24 << (((^(v2))) & 63))
                        _ = fi35
                        _ = v36
                    }
                } else {
                    switch 2147483647 {
                    case -2147483648:
                        if (0 <= 672) {
                            if len(v5) > 0 {
                                v5[0] = "foo bar"
                            }
                            for fi37 := 0; fi37 < 13; fi37++ {
                                var v38 int16 = v21
                                if len(v5) > 0 {
                                    v5[0] = "a"
                                }
                                _ = fi37
                                _ = v38
                            }
                        } else {
                            var v39 int16 = v17
                            _ = v39
                        }
                    case 32768:
                        v40 := GenS1{F0: 2147483647, F1: "\n\t", F2: 2147483647}
                        _ = v40
                    case -1:
                        var v41 bool = false
                        _ = v41
                    default:
                        for fi42 := 0; fi42 < 14; fi42++ {
                            fmt.Println(v3, v21, v4)
                            for fi43 := 0; fi43 < 7; fi43++ {
                                v44 := v1["a"]
                                switch (fi42 >> (((^(v3))) & 63)) {
                                case 128:
                                    if v44 {
                                        v24 += 1
                                        v45 := GenS1{F0: -2147483648, F1: ("\n\t" + (("a" + "a") + "\n\t")), F2: -667}
                                        v21 += -12
                                        if v44 {
                                            fmt.Println(v21, v17)
                                            v17 = v17
                                            var v46 int32 = 42
                                            _ = v46
                                        }
                                        _ = v45
                                    } else {
                                        if true {
                                            v20 = 1
                                            var v47 int8 = 73
                                            v1[("hello" + "hello")] = v44
                                            _ = v47
                                        } else {
                                            if (true && v44) {
                                                v4 += v20
                                            } else {
                                                var v48 int32 = 42
                                                _ = v48
                                            }
                                            switch "" {
                                            case "hello":
                                                var v49 int32 = -1
                                                _ = v49
                                            case "\n\t":
                                                v44 = (1 != v17)
                                            }
                                        }
                                    }
                                case 127:
                                    v3 -= v2
                                default:
                                    v3 = 65535
                                }
                                _ = fi43
                                _ = v44
                            }
                            v20 += (^(v20))
                            _ = fi42
                        }
                    }
                    v50 := GenS2{F0: int8(v21), F1: v24, F2: true}
                    _ = v50
                }
                var v51 int8 = 1
                _ = v51
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
