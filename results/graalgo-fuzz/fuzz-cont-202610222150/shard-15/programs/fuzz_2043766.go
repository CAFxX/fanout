package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint }

func main() {
    var v1 float64 = float64((0.0 / 1.0))
    var v2 int = 32768
    var v3 string = "\n\t"
    v4 := make(map[uint32]bool)
    v5 := []int{v2, v2, (v2 * v2)}
    var v6 int16 = -235
    v2 = v2
    var v7 uint = 0
    var v8 string = ("" + v3)
    var v9 uint32 = 827
    v6 -= v6
    v10 := GenS1{F0: 1, F1: 32767}
    if len(v5) > 0 {
        v5[0] = -443
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        v7 = v7
        _ = fi11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        if false {
            _ = len(v5)
            v7 -= 255
        } else {
            if true {
                if (-838 < 1) {
                    v5 = append(v5, 0)
                    var v13 int8 = 27
                    var v14 int32 = 985
                    _ = v13
                    _ = v14
                } else {
                    _ = len(v5)
                }
                v1 -= v1
            } else {
                v7 = 65535
                v4[v9] = false
                for fi15 := 0; fi15 < 5; fi15++ {
                    v1 += ((v1 - v1) / float64(1.5))
                    var v16 uint8 = 0
                    if (!true) {
                        if true {
                            if (3.14159 <= (((float32(v1) * (float32(v1) * float32(v1))) + (0.0 + (0.1 / 0.1))) * ((1.5 * 3.14159) + 3.14159))) {
                                switch 4294967295 {
                                case 9223372036854775807:
                                    _ = len(v4)
                                }
                            } else {
                                v4[(v9 / ((798) | 1))] = ((v2 << (((v7 / ((42) | 1))) & 63)) == (fi12 | fi12))
                                var v17 uint = v7
                                var v18 string = ("foo bar" + "a")
                                _ = v17
                                _ = v18
                            }
                            _ = len(v5)
                        } else {
                            v19 := GenS1{F0: 255, F1: 2}
                            _ = v19
                        }
                    } else {
                        fmt.Println(fi15, v8, v3)
                        _ = len(v5)
                        v20 := GenS2{F0: ("a" + v8)}
                        _ = v20
                    }
                    _ = fi15
                    _ = v16
                }
            }
            if true {
                var v21 uint64 = 120
                v21 += v21
                v21 -= v21
                _ = v21
            } else {
                fmt.Println(v8)
                v22 := GenS1{F0: 0, F1: 217}
                if false {
                    v23 := GenS2{F0: ("hello" + v3)}
                    var v24 uint = (v7 & (^(v7)))
                    switch ((((v8 + "") + v8) + "\n\t") + v8) {
                    case "foo bar":
                        v7 = 255
                    case "a":
                        switch "hello" {
                        case "foo bar":
                            v4[(v9 | 42)] = (false && true)
                        default:
                            fmt.Println(v7)
                        }
                    default:
                        var v25 int8 = 1
                        _ = v25
                    }
                    _ = v23
                    _ = v24
                } else {
                    var v26 string = ""
                    for fi27 := 0; fi27 < 16; fi27++ {
                        if (1 == 4294967295) {
                            v8 = (v26 + v3)
                            v3 = ("hello" + v26)
                        } else {
                            var v28 int = fi27
                            switch (v26 + ((v8 + (("foo bar" + v8) + v26)) + "")) {
                            case "\n\t":
                                v4[v9] = (0 < -9223372036854775808)
                            case "":
                                if (!false) {
                                    var v29 int8 = 1
                                    v30 := GenS3{F0: 255, F1: 18446744073709551615, F2: v7}
                                    _ = v29
                                    _ = v30
                                }
                            }
                            _ = v28
                        }
                        v2 += fi12
                        v6 = v6
                        _ = fi27
                    }
                    _ = v26
                }
                _ = v22
            }
            v31 := GenS2{F0: "\n\t"}
            _ = v31
        }
        switch (v2 % ((fi12) | 1)) {
        case 1:
            for fi32 := 0; fi32 < 8; fi32++ {
                v9 = 65535
                var v33 int16 = v6
                _ = fi32
                _ = v33
            }
        default:
            v8 = ""
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi12
    }
    if (false || false) {
        var v34 uint = (v7 * uint(v9))
        _ = v34
    } else {
        var v35 int64 = -9223372036854775808
        _ = v35
    }
    v7 += 0
    if len(v5) > 0 {
        v5[0] = (v2 * v2)
    }
    if (v1 == v1) {
        v36 := GenS2{F0: "\n\t"}
        _ = v36
    }
    fmt.Println(v2)
    v2 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
