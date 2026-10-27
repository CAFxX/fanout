package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint8; F2 bool }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -32768
    var v3 int8 = 71
    var v4 uint64 = 2
    var v5 string = "hello"
    if true {
        if (!true) {
            if ((!((true || false) && ((0.0 - (0.1 * -2.25)) > (0.0 * 3.14159)))) || (false && true)) {
                for fi6 := 0; fi6 < 19; fi6++ {
                    for fi7 := 0; fi7 < 5; fi7++ {
                        if true {
                            if (671 <= 42) {
                                v5 = v5
                                v8 := GenS1{F0: v4}
                                v9 := GenS2{F0: ((v3 << ((2) & 63)) <= v3), F1: 1, F2: (100 == (v3 % ((127) | 1)))}
                                _ = v8
                                _ = v9
                            }
                        }
                        if true {
                            v3 -= (v3 * -128)
                        }
                        fmt.Println(v2, v3)
                        v10 := GenS2{F0: ((62 == (v3 * (v3 & 127))) && false), F1: 179, F2: true}
                        v1["hello"] = ("foo bar" + v5)
                        _ = fi7
                        _ = v10
                    }
                    v2 -= (v2 & (v2 << ((uint(v4)) & 63)))
                    v11 := GenS1{F0: (v4 + (v4 % ((v4) | 1)))}
                    _ = fi6
                    _ = v11
                }
                if (127 > 32767) {
                    if true {
                        fmt.Println(v3)
                        var v12 int = -609
                        v13 := GenS3{F0: v4}
                        _ = v12
                        _ = v13
                    }
                    var v14 int64 = 1
                    _ = v14
                } else {
                    var v15 float32 = 3.14159
                    if ((!(255 != 404)) && true) {
                        var v16 float64 = (3.14159 + (float64(0.0) / ((3.14159 + 0.0) - (0.0 - (-2.25 / 1.0)))))
                        _ = v16
                    } else {
                        var v17 int16 = int16(v3)
                        v3 += (v3 - v3)
                        v2 = 588
                        _ = v17
                    }
                    fmt.Println(v2, v3, v4)
                    _ = v15
                }
                var v18 uint16 = 65535
                _ = v18
            } else {
                for fi19 := 0; fi19 < 12; fi19++ {
                    if (v5 > "a") {
                        var v20 uint64 = (^(v4))
                        _ = v20
                    } else {
                        fmt.Println(fi19)
                    }
                    _ = fi19
                }
                var v21 int32 = 614
                _ = v21
            }
            delete(v1, "a")
        }
    }
    v4 = (v4 / ((v4) | 1))
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 float64 = 3.14159
            v5 = "a"
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    v25 := []int{-9223372036854775808, 1, 127}
    v2 -= 1
    if false {
        if (false || (!true)) {
            var v26 uint8 = uint8(v4)
            switch int(v4) {
            case 2147483647:
                if len(v25) > 0 {
                    v25[0] = 4294967295
                }
            case 674:
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
            }
            _ = v26
        }
        v2 -= (v2 >> ((1) & 63))
    }
    v27 := GenS2{F0: (0 != int32(v4)), F1: 0, F2: false}
    v5 = "hello"
    v5 = v5
    for fi28 := 0; fi28 < 5; fi28++ {
        fmt.Println(v5, v5, v2)
        _ = fi28
    }
    v4 = (v4 * v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
