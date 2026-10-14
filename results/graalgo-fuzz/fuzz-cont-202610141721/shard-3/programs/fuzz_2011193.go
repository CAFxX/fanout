package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int16; F1 int16 }
func genFn0() int8 {
    v7 := GenS1{F0: 121}
    var v8 bool = false
    _ = v7
    _ = v8
    return 127
}

func main() {
    v1 := []string{"hello"}
    var v2 float64 = (3.14159 - 0.0)
    var v3 int32 = -571
    var v4 string = (((("" + "foo bar") + ("a" + "hello")) + "\n\t") + "\n\t")
    v5 := []uint32{1, 314, 724}
    var v6 uint16 = 65535
    _ = len(v1)
    switch "a" {
    case "a":
        var v9 float64 = v2
        _ = v9
    }
    fmt.Println(v6)
    v6 = v6
    _ = len(v1)
    switch v4 {
    case "foo bar":
        v10 := GenS1{F0: 18446744073709551615}
        _ = v10
    case "\n\t":
        var v11 float32 = 0.1
        _ = v11
    case "":
        for fi12 := 0; fi12 < 14; fi12++ {
            if (0.1 >= float64((1e10 + float32((((3.14159 * -2.25) - v2) + v2))))) {
                v4 = v4
                v6 -= (v6 | (v6 ^ v6))
            } else {
                v6 = v6
                for fi13 := 0; fi13 < 15; fi13++ {
                    if ((^(fi12)) >= (fi12 & fi12)) {
                        fmt.Println(v2, v2, v6)
                        v6 -= (^(v6))
                        var v14 uint8 = 163
                        fmt.Println(fi12, fi13, v14)
                        v15 := GenS2{F0: 42, F1: ((-2.25 + 0.1) > (float32(v2) + 3.14159))}
                        _ = v14
                        _ = v15
                    }
                    _ = fi13
                }
            }
            if ((-848 < -9223372036854775808) && (0 < 99)) {
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 int = fi16
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi12
        }
    default:
        for fi18 := 0; fi18 < 5; fi18++ {
            v4 = ((((v4 + "\n\t") + "") + "") + "foo bar")
            if true {
                if (-9223372036854775808 <= 0) {
                    fmt.Println(v2)
                } else {
                    v19 := GenS2{F0: 4294967295, F1: true}
                    _ = v19
                }
            } else {
                v2 += float64((3.14159 - (0.0 / 3.14159)))
                switch 450 {
                case -941:
                    v2 -= 1.5
                case 127:
                    v6 = (v6 / ((515) | 1))
                }
            }
            fmt.Println(v6)
            _ = fi18
        }
    }
    v1 = append(v1, (v4 + v4))
    v6 += 0
    if (false && (601 == 240)) {
        if (((32768 == 9223372036854775807) && false) || false) {
            switch v4 {
            case "a":
                fmt.Println(v6, v6)
            case "hello":
                v3 = (v3 - (v3 >> ((490) & 63)))
            case "foo bar":
                var v20 uint64 = 42
                _ = v20
            default:
                var v21 float64 = (1.5 + -2.25)
                _ = v21
            }
            var v22 string = ("hello" + v4)
            var v23 int32 = (v3 + (-(v3)))
            _ = v22
            _ = v23
        } else {
            _ = len(v1)
            if (v3 >= -718) {
                v24 := GenS2{F0: 2, F1: ((-(v3)) >= 471)}
                var v25 float32 = 1.5
                _ = v24
                _ = v25
            } else {
                switch 128 {
                case 128:
                    for fi26 := 0; fi26 < 14; fi26++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            var v27 int32 = (v3 >> ((1) & 63))
                            if false {
                                v3 = (v3 % (((v3 & v27)) | 1))
                                fmt.Println(v2, v6)
                                var v28 int16 = 32767
                                _ = v28
                            } else {
                                v29 := GenS1{F0: 253}
                                if len(v1) > 0 {
                                    v1[0] = ((v4 + "hello") + "hello")
                                }
                                _ = v29
                            }
                            var v30 float32 = (-2.25 - -2.25)
                            _ = v27
                            _ = v30
                        }
                        var v31 int = (fi26 >> ((181) & 63))
                        _ = fi26
                        _ = v31
                    }
                case 9223372036854775807:
                    v6 += v6
                case 795:
                    for fi32 := 0; fi32 < 17; fi32++ {
                        fmt.Println(v3, v6)
                        _ = fi32
                    }
                }
                v33 := GenS1{F0: 324}
                v6 = v6
                _ = v33
            }
        }
    } else {
        _ = len(v1)
        v6 += v6
        v34 := GenS2{F0: 2, F1: (1e10 <= -2.25)}
        _ = v34
    }
    switch -1 {
    case 0:
        var v35 uint32 = 116
        _ = v35
    }
    v36 := GenS3{F0: 42, F1: 497}
    switch int(v3) {
    case 42:
        for fi37 := 0; fi37 < 3; fi37++ {
            if len(v5) > 0 {
                v5[0] = 660
            }
            _ = fi37
        }
    case 127:
        fmt.Println(v6, v6, v3)
    default:
        var v38 int64 = 42
        _ = v38
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
