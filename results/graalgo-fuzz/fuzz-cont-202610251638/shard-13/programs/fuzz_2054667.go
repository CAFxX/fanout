package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int; F2 int64 }
type GenS3 struct { F0 int64; F1 int16; F2 uint32 }
type GenS4 struct { F0 int16 }
func genFn0() int16 {
    var v5 int16 = -866
    fmt.Println(v5)
    _ = v5
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 32768
    var v3 int8 = -54
    var v4 uint32 = 1
    var v6 int32 = -484
    v3 -= -128
    for fi7 := 0; fi7 < 17; fi7++ {
        v3 = v3
        fmt.Println(fi7, v6, v4)
        if (fi7 > (-(fi7))) {
            fmt.Println(v6)
            if true {
                v4 = 0
                var v8 int8 = -1
                var v9 int32 = v6
                _ = v8
                _ = v9
            } else {
                if (4294967295 < 946) {
                    if (true || false) {
                        v10 := GenS1{F0: 2}
                        for fi11 := 0; fi11 < 18; fi11++ {
                            switch "" {
                            case "hello":
                                var v12 bool = true
                                _ = v12
                            case "\n\t":
                                switch fi11 {
                                case -9223372036854775808:
                                    v13 := GenS2{F0: 1, F1: fi11, F2: 42}
                                    _ = v13
                                }
                            default:
                                v3 -= 127
                                switch (-(fi11)) {
                                case 1:
                                    var v14 string = "\n\t"
                                    var v15 bool = (true || false)
                                    v1[int(v6)] = (fi11 * 42)
                                    _ = v14
                                    _ = v15
                                case 32768:
                                    var v16 float64 = 0.1
                                    _ = v16
                                case 32767:
                                    switch "hello" {
                                    case "a":
                                        v6 -= int32(v3)
                                    default:
                                        var v17 int16 = 1
                                        _ = v17
                                    }
                                }
                                v6 = v6
                            }
                            _ = fi11
                        }
                        v18 := []int{-151, (v2 & fi7), fi7, (fi7 * -115)}
                        _ = v10
                        _ = v18
                    } else {
                        v4 -= 1
                        if true {
                            v19 := []int{fi7, -1, fi7, fi7}
                            v1[fi7] = (v2 >> ((4294967295) & 63))
                            var v20 uint16 = 65535
                            _ = v19
                            _ = v20
                        }
                    }
                }
                v3 -= v3
            }
            var v21 string = ""
            _ = v21
        } else {
            var v22 int8 = 94
            v22 -= (v22 >> ((243) & 63))
            v22 += v3
            _ = v22
        }
        _ = fi7
    }
    v23 := []int{(v2 / ((9223372036854775807) | 1)), (v2 | 0), (v2 >> ((18446744073709551615) & 63))}
    for fi24 := 0; fi24 < 7; fi24++ {
        for fi25 := 0; fi25 < 19; fi25++ {
            v2 += 2147483647
            var v26 int8 = (v3 & v3)
            if false {
                var v27 float32 = -2.25
                if (2 != 492) {
                    v28 := GenS1{F0: 255}
                    v29 := GenS1{F0: 76}
                    var v30 int8 = (v3 + -59)
                    v31 := GenS2{F0: 4294967295, F1: (fi25 / (((v2 & (fi24 / ((fi25) | 1)))) | 1)), F2: 1}
                    fmt.Println(fi25, v3, fi25)
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    v26 = v3
                    v32 := GenS3{F0: -1, F1: -300, F2: (v4 - (v4 | v4))}
                    _ = v32
                }
                var v33 uint64 = 911
                _ = v27
                _ = v33
            }
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    if (false || true) {
        v34 := GenS4{F0: 42}
        if (true || false) {
            fmt.Println(v3, v3)
            if true {
                _ = len(v1)
                if false {
                    v2 -= -502
                    v4 += v4
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                }
            }
            var v35 bool = false
            _ = v35
        }
        switch (v2 << ((2) & 63)) {
        case 4294967295:
            v2 = (v2 & 4294967295)
        }
        _ = v34
    }
    if true {
        switch "foo bar" {
        case "a":
            delete(v1, v2)
        case "":
            fmt.Println(v6, v6, v6)
        case "hello":
            if false {
                if (-1 < -981) {
                    var v36 bool = (1 < 32767)
                    _ = v36
                }
            } else {
                v3 += (v3 + (^(v3)))
                var v37 uint64 = 151
                _ = v37
            }
        default:
            fmt.Println(v2, v3, v3)
        }
        if len(v23) > 0 {
            v23[0] = v2
        }
        v38 := v1[(-(v2))]
        _ = v38
    }
    fmt.Println(v3, v6, v3)
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v6
}
