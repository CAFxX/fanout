package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 string }
type GenS3 struct { F0 int64; F1 int8; F2 string }
type GenS4 struct { F0 int32; F1 int32; F2 uint }
func genFn0() int {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 uint64 = 1
        var v7 float64 = (float64(((1.5 - float32(0.0)) - (1.5 + 0.0))) * (-2.25 / 1e10))
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        v8 := GenS1{F0: 0}
        var v9 uint16 = 1
        var v10 uint32 = 42
        _ = v10
        _ = v8
        _ = v9
    } else {
        v11 := GenS2{F0: 1, F1: 42, F2: "hello"}
        var v12 int32 = 344
        _ = v11
        _ = v12
    }
    return 145
}

func main() {
    v1 := []uint32{33}
    var v2 int64 = -1
    var v3 uint32 = 2147483648
    var v4 int8 = 42
    var v13 float64 = ((1e10 + 0.0) * 0.0)
    var v14 uint64 = 18446744073709551615
    v15 := GenS3{F0: (v2 >> ((857) & 63)), F1: 42, F2: ""}
    v1 = append(v1, v3)
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int64 = (v2 >> ((84) & 63))
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            v14 = v14
            _ = fi19
        }
        _ = fi18
    }
    var v20 int16 = 32767
    if (v3 >= v3) {
        v3 = v3
        v21 := GenS1{F0: int8(v3)}
        if (74 > 255) {
            v1 = append(v1, 2147483648)
            v22 := GenS2{F0: 42, F1: (v2 ^ v2), F2: ""}
            _ = v22
        }
        _ = v21
    }
    v3 = uint32(v14)
    v2 += (^(v2))
    if true {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 uint64 = v14
            if false {
                v1 = append(v1, v3)
                switch "a" {
                case "hello":
                    for fi25 := 0; fi25 < 5; fi25++ {
                        switch (fi23 << ((255) & 63)) {
                        case 2147483647:
                            v26 := GenS2{F0: 65535, F1: 9223372036854775807, F2: "foo bar"}
                            _ = v26
                        }
                        if len(v1) > 0 {
                            v1[0] = (v3 * v3)
                        }
                        _ = fi25
                    }
                }
            }
            for fi27 := 0; fi27 < 10; fi27++ {
                if true {
                    v4 += v4
                    v13 += v13
                } else {
                    var v28 uint32 = (v3 % (((^(v3))) | 1))
                    _ = v28
                }
                var v29 int32 = -2147483648
                _ = fi27
                _ = v29
            }
            _ = fi23
            _ = v24
        }
        fmt.Println(v2)
        v13 = 3.14159
    }
    if false {
        _ = len(v1)
        var v30 uint16 = 341
        _ = v30
    } else {
        var v31 uint16 = 49
        _ = v31
    }
    v14 += v14
    var v32 int64 = v2
    if ((63 >= 202) && true) {
        v33 := GenS2{F0: 1, F1: -476, F2: "hello"}
        for fi34 := 0; fi34 < 13; fi34++ {
            v2 = -477
            var v35 uint32 = 74
            _ = fi34
            _ = v35
        }
        _ = v33
    } else {
        fmt.Println(v3)
    }
    switch 821 {
    case 127:
        if (true && false) {
            var v36 uint = 470
            var v37 float64 = v13
            _ = v36
            _ = v37
        } else {
            v38 := GenS4{F0: 2147483647, F1: -1, F2: 265}
            v14 += 18446744073709551615
            for fi39 := 0; fi39 < 3; fi39++ {
                var v40 bool = true
                v41 := GenS1{F0: 100}
                if (!v40) {
                    for fi42 := 0; fi42 < 1; fi42++ {
                        v2 = (v32 << ((18446744073709551615) & 63))
                        switch fi42 {
                        case 78:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            var v43 int8 = (v4 - int8(v32))
                            _ = v43
                        }
                        _ = fi42
                    }
                } else {
                    v44 := GenS1{F0: (-(v4))}
                    v1 = append(v1, (v3 << ((1) & 63)))
                    if v40 {
                        fmt.Println(v3, v4)
                    }
                    _ = v44
                }
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = v38
        }
    case -2147483648:
        if false {
            _ = len(v1)
            v4 = v4
        } else {
            var v45 uint16 = 295
            var v46 uint16 = v45
            _ = v45
            _ = v46
        }
    default:
        for fi47 := 0; fi47 < 15; fi47++ {
            for fi48 := 0; fi48 < 12; fi48++ {
                v3 -= (v3 - v3)
                if false {
                    v3 = (v3 - 2147483647)
                }
                v3 += (^(v3))
                _ = fi48
            }
            v32 -= v32
            _ = fi47
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v4
}
