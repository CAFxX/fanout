package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int64; F1 int; F2 int16 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 int8; F1 uint8; F2 int }
type GenS6 struct { F0 string; F1 int8 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]bool)
    var v3 int32 = -1
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 13; fi5++ {
        v1 -= v1
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := GenS1{F0: "hello", F1: v1}
            var v8 int = -249
            v8 = (fi6 >> ((4294967295) & 63))
            v1 -= (v1 ^ (v1 * 356))
            var v9 float32 = ((3.14159 / 1.5) / 1e10)
            _ = len(v4)
            v10 := GenS1{F0: "", F1: v1}
            _ = fi6
            _ = v10
            _ = v7
            _ = v8
            _ = v9
        }
        var v11 uint64 = (v1 * (v1 / ((v1) | 1)))
        _ = fi5
        _ = v11
    }
    switch 1 {
    case 375:
        fmt.Println(v1, v3)
    default:
        for fi12 := 0; fi12 < 4; fi12++ {
            v3 = (v3 << ((255) & 63))
            v13 := GenS1{F0: "a", F1: v1}
            _ = fi12
            _ = v13
        }
    }
    var v14 uint32 = 2
    v3 += (v3 ^ v3)
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 string = "foo bar"
        if false {
            if true {
                fmt.Println(v3)
                var v17 int16 = 0
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 int32 = -1
                    v20 := GenS1{F0: "\n\t", F1: (v1 % ((18446744073709551615) | 1))}
                    v21 := GenS1{F0: "a", F1: (^(v1))}
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = v17
            }
            v22 := GenS2{F0: (v3 + v3), F1: v1, F2: 94}
            _ = v22
        } else {
            if (false && ((v1 * (v1 - v1)) != 4294967295)) {
                for fi23 := 0; fi23 < 4; fi23++ {
                    if true {
                        var v24 int = 457
                        if true {
                            var v25 string = v16
                            _ = v25
                        }
                        if false {
                            v3 = v3
                            if ((127 <= 2) || (!(v3 > 963))) {
                                switch fi23 {
                                case 1:
                                    v26 := GenS3{F0: -9223372036854775808, F1: int(v1), F2: -875}
                                    _ = v26
                                case 32768:
                                    var v27 string = (v16 + v16)
                                    _ = v27
                                default:
                                    v28 := GenS4{F0: v14, F1: 42}
                                    _ = v28
                                }
                                var v29 uint32 = (^(v14))
                                v30 := GenS5{F0: -1, F1: 128, F2: (fi15 & (fi15 * fi15))}
                                _ = v29
                                _ = v30
                            } else {
                                if true {
                                    v31 := GenS5{F0: -128, F1: 255, F2: int(v1)}
                                    if true {
                                        v24 -= fi15
                                    }
                                    var v32 uint8 = 255
                                    _ = v31
                                    _ = v32
                                }
                                var v33 bool = (!(v16 < (v16 + "\n\t")))
                                v14 -= v14
                                _ = v33
                            }
                        } else {
                            v34 := GenS2{F0: v3, F1: v1, F2: 91}
                            var v35 bool = true
                            v35 = (v35 || v35)
                            var v36 uint16 = 1
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        _ = v24
                    }
                    _ = fi23
                }
                for fi37 := 0; fi37 < 18; fi37++ {
                    if ((v1 * (v1 / ((v1) | 1))) == v1) {
                        v1 = 433
                        fmt.Println(v14, v3, v1)
                        if false {
                            switch (v16 + v16) {
                            case "":
                                v38 := GenS6{F0: ("" + v16), F1: 42}
                                _ = v38
                            case "a":
                                fmt.Println(v3)
                            case "\n\t":
                                fmt.Println(v3)
                            default:
                                var v39 string = v16
                                v39 = v16
                                v39 = ""
                                _ = v39
                            }
                            fmt.Println(v16, v1)
                        } else {
                            var v40 uint16 = 2
                            _ = v40
                        }
                    }
                    v2[(("foo bar" + v16) + v16)] = (127 >= 127)
                    _ = fi37
                }
            }
            for fi41 := 0; fi41 < 2; fi41++ {
                if (true && ((fi41 <= (fi15 | -3)) || false)) {
                    fmt.Println(v14, v3)
                } else {
                    if (false && (false || true)) {
                        fmt.Println(v3, v3, v1)
                        var v42 uint32 = 248
                        _ = v42
                    } else {
                        v3 = v3
                    }
                }
                var v43 uint = 4294967295
                var v44 int64 = -610
                _ = fi41
                _ = v43
                _ = v44
            }
            var v45 string = ((((("foo bar" + ("hello" + (v16 + v16))) + ("\n\t" + v16)) + ("hello" + v16)) + (v16 + v16)) + v16)
            _ = v45
        }
        var v46 int16 = -1
        _ = fi15
        _ = v16
        _ = v46
    }
    v14 = uint32(v1)
    v2["foo bar"] = (true || (1.5 >= 3.14159))
    for fi47 := 0; fi47 < 18; fi47++ {
        for fi48 := 0; fi48 < 13; fi48++ {
            if (false || (int64(v3) == -1)) {
                var v49 bool = (v14 >= (v14 - v14))
                _ = v49
            }
            _ = fi48
        }
        _ = fi47
    }
    v14 = v14
    fmt.Println(v14, v14)
    var v50 int32 = v3
    v3 -= (v3 / (((v3 / ((-2147483648) | 1))) | 1))
    v51 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
}
