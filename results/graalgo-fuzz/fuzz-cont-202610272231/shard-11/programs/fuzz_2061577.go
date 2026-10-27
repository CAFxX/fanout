package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int8 = 39
    var v3 uint8 = 127
    var v4 int = -464
    v5 := make(map[string]uint32)
    v1 += v1
    var v6 int32 = 233
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v3, fi7)
        _ = fi7
    }
    switch "foo bar" {
    case "foo bar":
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(v4)
            _ = fi8
        }
    case "a":
        if (!(993 >= 222)) {
            if ((("" >= "") || (int(v2) <= v4)) && (1.5 > 0.1)) {
                for fi9 := 0; fi9 < 16; fi9++ {
                    v2 += int8(v1)
                    _ = fi9
                }
                if ((v4 ^ 4294967295) == (v4 - v4)) {
                    var v10 bool = true
                    switch (v4 >> ((65535) & 63)) {
                    case 127:
                        var v11 int8 = v2
                        _ = v11
                    case -2147483648:
                        var v12 int8 = (v2 + v2)
                        _ = v12
                    }
                    if v10 {
                        if v10 {
                            v13 := GenS1{F0: 32767, F1: -32768}
                            v14 := []int{v4, (-(v4))}
                            _ = v13
                            _ = v14
                        } else {
                            v10 = (!(v10 && false))
                            fmt.Println(v3, v3)
                        }
                        if (!((v4 / ((v4) | 1)) > (-(v4)))) {
                            delete(v5, "foo bar")
                        }
                        _ = len(v5)
                    }
                    _ = v10
                }
                v3 = (v3 << ((609) & 63))
            } else {
                for fi15 := 0; fi15 < 7; fi15++ {
                    if true {
                        v16 := v5["hello"]
                        for fi17 := 0; fi17 < 10; fi17++ {
                            var v18 uint16 = 255
                            _ = fi17
                            _ = v18
                        }
                        _ = v16
                    }
                    switch v4 {
                    case 42:
                        var v19 uint16 = 457
                        _ = v19
                    case 32768:
                        for fi20 := 0; fi20 < 18; fi20++ {
                            v5[""] = 42
                            _ = fi20
                        }
                    default:
                        if ((float32(3.14159) >= 3.14159) && false) {
                            v21 := v5[("a" + "a")]
                            _ = v21
                        }
                    }
                    v3 -= 167
                    _ = fi15
                }
                v6 -= v6
                v1 = (v1 + v1)
            }
        } else {
            v4 -= v4
            switch (v4 ^ v4) {
            case -868:
                _ = len(v5)
            case 128:
                fmt.Println(v6)
            default:
                if true {
                    v3 += (^(v3))
                    switch 127 {
                    case 128:
                        v22 := []int{(v4 >> ((785) & 63)), (v4 % ((int(v6)) | 1))}
                        _ = v22
                    default:
                        fmt.Println(v3, v4)
                    }
                    var v23 int16 = 909
                    _ = v23
                }
            }
        }
    case "hello":
        var v24 uint16 = 193
        _ = v24
    }
    var v25 int8 = v2
    var v26 uint8 = 128
    v27 := GenS1{F0: 1, F1: 0}
    v4 = (v4 / ((v4) | 1))
    fmt.Println(v25)
    v28 := GenS1{F0: 927, F1: 1000}
    var v29 int32 = v6
    var v30 int16 = 42
    if true {
        _ = len(v5)
        v25 = 42
        v2 = (v25 + (v2 / (((-(v2))) | 1)))
    }
    v4 -= (v4 >> ((2) & 63))
    var v31 int = (v4 ^ (^(v4)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
