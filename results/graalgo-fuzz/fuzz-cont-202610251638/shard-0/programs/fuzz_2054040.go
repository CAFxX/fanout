package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint16; F1 int32 }
func genFn0(p0 int16, p1 int32, p2 int8) uint8 {
    var v4 int = 1
    var v5 float64 = 0.0
    var v6 uint16 = 1
    v4 = (v4 << ((65535) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int32 = 1
    v2 := []int{4294967295, 0}
    var v3 string = ""
    fmt.Println(v1)
    if ((float32((float64(-2.25) + (1e10 * 0.1))) > ((1e10 - (-2.25 + (0.0 * float32(1.5)))) / float32(-2.25))) && ((!true) || (0 <= -1))) {
        for fi7 := 0; fi7 < 14; fi7++ {
            v3 = v3
            var v8 int8 = 100
            _ = fi7
            _ = v8
        }
        var v9 int32 = -270
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int16 = 42
            v11 += -699
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 int = fi12
            v14 := GenS1{F0: 59}
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            if false {
                if true {
                    if len(v2) > 0 {
                        v2[0] = fi16
                    }
                }
            }
            fmt.Println(fi16, v3, fi16)
            var v17 bool = (!((42 < 104) && true))
            switch (fi15 & (-(fi15))) {
            case -392:
                if ((^(fi16)) == 127) {
                    fmt.Println(v3)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    switch 0 {
                    case 32768:
                        v18 := GenS2{F0: v3}
                        _ = v18
                    case 2147483647:
                        if (42 != 2147483647) {
                            if v17 {
                                fmt.Println(v3)
                                if (v17 && v17) {
                                    v19 := GenS2{F0: v3}
                                    v20 := GenS3{F0: 1}
                                    var v21 int = 0
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                } else {
                                    var v22 string = "foo bar"
                                    fmt.Println(v22, v22, fi15)
                                    var v23 int32 = 0
                                    _ = v22
                                    _ = v23
                                }
                            }
                            v17 = (1.5 <= 1e10)
                            v24 := GenS4{F0: -847}
                            v3 = v3
                            var v25 int16 = -32768
                            _ = v24
                            _ = v25
                        } else {
                            var v26 int32 = v1
                            _ = v26
                        }
                    }
                    var v27 int64 = 803
                    _ = v27
                }
            default:
                v28 := GenS4{F0: 1}
                _ = v28
            }
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v3 = v3
    switch 77 {
    case 9223372036854775807:
        v29 := GenS5{F0: -401, F1: 18446744073709551615, F2: -1}
        _ = v29
    case 32768:
        v3 = (v3 + v3)
    case 127:
        if false {
            _ = len(v2)
            fmt.Println(v3, v3)
            v30 := GenS5{F0: 626, F1: 18446744073709551615, F2: 42}
            _ = v30
        }
    }
    if false {
        for fi31 := 0; fi31 < 18; fi31++ {
            v32 := GenS6{F0: 370, F1: v1}
            _ = fi31
            _ = v32
        }
        if (42 != -927) {
            fmt.Println(v1)
        } else {
            var v33 string = v3
            _ = v33
        }
        fmt.Println(v1, v3)
    } else {
        switch ("" + v3) {
        case "\n\t":
            v1 = v1
        default:
            v3 = "a"
        }
        if (1e10 != ((float64((-2.25 + -2.25)) / 1e10) * ((((1.5 - 1e10) + 1e10) + 3.14159) * (0.1 * 1.5)))) {
            var v34 int = -9223372036854775808
            _ = v34
        }
    }
    v2 = append(v2, -2147483648)
    switch v3 {
    case "hello":
        if ((^(v1)) >= v1) {
            fmt.Println(v3)
            v2 = append(v2, 4294967295)
            v3 = v3
        }
    default:
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
