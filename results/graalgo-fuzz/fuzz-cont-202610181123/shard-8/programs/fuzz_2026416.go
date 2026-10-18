package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint; F1 int16 }

func main() {
    var v1 int64 = -462
    v2 := make(map[string]bool)
    var v3 int64 = v1
    var v4 int8 = -82
    var v5 int32 = -983
    var v6 uint64 = 0
    if false {
        v4 = 127
        var v7 int64 = v3
        _ = v7
    }
    fmt.Println(v4)
    switch 128 {
    case 35:
        v3 = 1
    case 128:
        switch 4294967295 {
        case 0:
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 int32 = v5
                _ = fi8
                _ = v9
            }
        }
    case 1:
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 uint16 = 32767
            _ = fi10
            _ = v11
        }
    default:
        v6 = v6
    }
    v6 += v6
    if false {
        v6 = (v6 * (v6 >> ((2) & 63)))
    } else {
        v12 := GenS1{F0: (v6 & 379), F1: 1}
        if true {
            if ("a" == "a") {
                var v13 float32 = float32((1.5 / 0.1))
                v4 += (v4 >> ((4294967295) & 63))
                var v14 float64 = 1.5
                _ = v13
                _ = v14
            } else {
                var v15 int64 = (v3 << ((4294967295) & 63))
                if ((v4 + (v4 - (v4 >> ((0) & 63)))) <= 127) {
                    for fi16 := 0; fi16 < 13; fi16++ {
                        for fi17 := 0; fi17 < 17; fi17++ {
                            v18 := []int{(fi17 | (fi17 | (fi17 + 127))), (-(fi17)), 4294967295}
                            v19 := GenS1{F0: 1, F1: v15}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        v20 := []int{(fi16 * (fi16 << ((2) & 63))), 32767}
                        if (!(85 != v3)) {
                            v21 := GenS1{F0: 2, F1: v15}
                            if (false && (!true)) {
                                var v22 uint16 = 255
                                if (!false) {
                                    v22 = v22
                                } else {
                                    v5 = 42
                                }
                                var v23 int = (fi16 ^ (-(fi16)))
                                _ = v22
                                _ = v23
                            }
                            _ = v21
                        } else {
                            v2["foo bar"] = false
                            switch 0 {
                            case 4294967295:
                                v1 += (v1 ^ (^(v1)))
                            case 1:
                                v24 := GenS2{F0: 2, F1: 32767}
                                _ = v24
                            }
                            var v25 int32 = (-(v5))
                            _ = v25
                        }
                        _ = fi16
                        _ = v20
                    }
                    for fi26 := 0; fi26 < 5; fi26++ {
                        v2[("a" + "a")] = false
                        v1 = (v15 >> ((1) & 63))
                        _ = fi26
                    }
                }
                var v27 int64 = v3
                _ = v15
                _ = v27
            }
        } else {
            var v28 uint32 = 32
            _ = v28
        }
        v6 -= (v6 | 2)
        _ = v12
    }
    if (-9223372036854775808 >= -708) {
        for fi29 := 0; fi29 < 3; fi29++ {
            v1 -= (v1 / (((v1 / ((v1) | 1))) | 1))
            _ = fi29
        }
        v2["\n\t"] = true
        v1 = v1
    }
    var v30 uint = 4294967295
    v31 := v2[("hello" + "")]
    var v32 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
