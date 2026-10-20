package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    v3 := []string{}
    var v4 int64 = 0
    v4 += v4
    v4 += (-(v4))
    if false {
        v5 := v2[411]
        _ = v5
    } else {
        switch 2147483647 {
        case -9223372036854775808:
            v4 += v4
        case 551:
            var v6 string = "foo bar"
            _ = v6
        default:
            fmt.Println(v4)
        }
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 string = "foo bar"
            v9 := GenS1{F0: 0, F1: fi7, F2: 32767}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v4 += v4
        _ = fi10
    }
    switch 656 {
    case 128:
        for fi11 := 0; fi11 < 11; fi11++ {
            if (1e10 >= float32(float64((1e10 * float32((((0.1 * 0.1) - 1e10) / 1.0)))))) {
                var v12 bool = false
                v13 := GenS1{F0: 32767, F1: 127, F2: fi11}
                v14 := GenS1{F0: 991, F1: 32767, F2: fi11}
                _ = v12
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 19; fi15++ {
                    var v16 uint64 = 674
                    if false {
                        if (-704 <= (v4 << ((18446744073709551615) & 63))) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v16 += (v16 % (((^(v16))) | 1))
                            v4 = v4
                        }
                    } else {
                        v4 += v4
                    }
                    var v17 uint = 18446744073709551615
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = fi11
        }
    case -9223372036854775808:
        for fi18 := 0; fi18 < 17; fi18++ {
            if (false || (!((v4 << ((2) & 63)) != v4))) {
                for fi19 := 0; fi19 < 5; fi19++ {
                    switch (fi18 << ((212) & 63)) {
                    case 32767:
                        v4 += 9223372036854775807
                    case 9223372036854775807:
                        var v20 int32 = -1
                        _ = v20
                    case 4294967295:
                        v4 -= v4
                    default:
                        v4 += v4
                    }
                    var v21 string = ("hello" + "a")
                    if (!(!(128 <= 732))) {
                        var v22 string = "a"
                        _ = v22
                    } else {
                        var v23 int16 = -32768
                        v23 = 42
                        if false {
                            v24 := GenS1{F0: 2, F1: 41, F2: 1}
                            _ = v24
                        }
                        _ = v23
                    }
                    _ = fi19
                    _ = v21
                }
                for fi25 := 0; fi25 < 1; fi25++ {
                    v3 = append(v3, "")
                    _ = fi25
                }
            }
            v26 := GenS1{F0: 109, F1: (fi18 >> ((0) & 63)), F2: fi18}
            _ = fi18
            _ = v26
        }
    case 0:
        var v27 float64 = float64(float32(3.14159))
        _ = v27
    default:
        if len(v3) > 0 {
            v3[0] = ""
        }
    }
    delete(v2, 1)
    v4 -= v4
    if true {
        if (true || (-1 < 42)) {
            if true {
                v1["foo bar"] = true
            }
            v4 += v4
        }
    } else {
        v4 = v4
        var v28 uint = 315
        var v29 uint8 = 2
        _ = v28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
