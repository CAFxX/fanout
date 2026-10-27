package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint8 = 127
    var v2 int8 = 1
    v3 := make(map[uint32]int)
    var v4 float32 = 0.0
    var v5 int64 = 0
    var v6 float64 = float64(v4)
    v1 = v1
    for fi7 := 0; fi7 < 2; fi7++ {
        v1 = (^(v1))
        if false {
            var v8 string = "foo bar"
            _ = v8
        } else {
            if (!false) {
                var v9 bool = (v2 != v2)
                v10 := GenS1{F0: 1, F1: 32767}
                _ = v10
                _ = v9
            } else {
                if false {
                    if (!("" < ("a" + (("" + ((((("\n\t" + "hello") + "foo bar") + "hello") + "foo bar") + "a")) + "")))) {
                        var v11 int32 = 0
                        if true {
                            var v12 uint32 = 0
                            v3[2] = fi7
                            _ = v12
                        } else {
                            v13 := []int{fi7, (fi7 ^ (fi7 & (fi7 & (fi7 >> ((399) & 63))))), -1, fi7}
                            v11 += 42
                            _ = v13
                        }
                        _ = v11
                    } else {
                        v14 := []int{(fi7 * int(v2)), (fi7 + (^(fi7))), (fi7 | (fi7 / ((fi7) | 1))), (fi7 * fi7)}
                        _ = v14
                    }
                }
            }
            fmt.Println(v6, v2)
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 float32 = (v4 - (float32((float64((v4 * v4)) + ((float64((-2.25 - v4)) / 1.5) / (3.14159 + float64(-2.25))))) * (v4 + v4)))
                if (18446744073709551615 == 4294967295) {
                    if (v2 != (^(v2))) {
                        if ((792 <= 572) || false) {
                            fmt.Println(v16, v4)
                        }
                    }
                    v1 -= (v1 << ((775) & 63))
                } else {
                    v6 = -2.25
                    v17 := []int{(fi7 | fi7)}
                    v5 -= v5
                    _ = v17
                }
                _ = len(v3)
                _ = fi15
                _ = v16
            }
        }
        var v18 uint16 = 255
        _ = fi7
        _ = v18
    }
    var v19 float32 = v4
    if false {
        var v20 bool = (!true)
        _ = v20
    } else {
        v2 = v2
        var v21 int = -37
        var v22 uint8 = 183
        _ = v21
        _ = v22
    }
    var v23 string = ("hello" + "hello")
    switch "\n\t" {
    case "a":
        v19 += v4
    case "foo bar":
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 uint16 = 2
            fmt.Println(fi24, v4)
            if ((-100 == 127) || (((v6 - 0.1) < (v6 - v6)) || (32767 >= 598))) {
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 string = ("a" + "a")
                    if (false && (!false)) {
                        fmt.Println(fi24, v19)
                    } else {
                        v28 := GenS1{F0: 0, F1: 811}
                        switch fi26 {
                        case -1:
                            var v29 uint64 = 1
                            _ = v29
                        case 127:
                            if (uint64(v25) > 42) {
                                var v30 float32 = (v4 * (v19 * ((0.0 - v4) * 0.0)))
                                var v31 uint64 = 4294967295
                                var v32 bool = true
                                _ = v30
                                _ = v31
                                _ = v32
                            } else {
                                v33 := GenS2{F0: -1, F1: 554}
                                v3[4294967295] = fi24
                                _ = v33
                            }
                        }
                        _ = v28
                    }
                    _ = fi26
                    _ = v27
                }
            }
            _ = fi24
            _ = v25
        }
    }
    v5 -= (v5 | v5)
    var v34 string = "hello"
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 int32 = 903
        _ = fi35
        _ = v36
    }
    v3[2147483647] = 32767
    var v37 uint64 = 1
    v38 := []int{376, 0, 2147483647, -9223372036854775808}
    for fi39 := 0; fi39 < 1; fi39++ {
        var v40 int32 = 2147483647
        v41 := GenS3{F0: v2}
        _ = fi39
        _ = v40
        _ = v41
    }
    v1 -= 128
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
