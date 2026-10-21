package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int16; F2 bool }
type GenS3 struct { F0 uint64; F1 int32; F2 int32 }
func genFn0() int32 {
    v4 := GenS1{F0: 127}
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int8 = 1
        v7 := GenS2{F0: -6, F1: int16(fi5), F2: (false && (!true))}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = v4
    return -1
}
func genFn1() int16 {
    var v8 int8 = 0
    v9 := []int{}
    var v10 int16 = -32768
    _ = v10
    _ = v8
    _ = v9
    return (v10 << ((843) & 63))
}

func main() {
    var v1 uint32 = 65535
    v2 := []uint32{}
    var v3 string = ("a" + "\n\t")
    var v11 uint = 65535
    for fi12 := 0; fi12 < 2; fi12++ {
        if true {
            if (-1 < 9223372036854775807) {
                var v13 string = "hello"
                _ = v13
            }
        }
        var v14 int64 = 42
        fmt.Println(v11)
        _ = fi12
        _ = v14
    }
    switch v3 {
    case "a":
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
    }
    if len(v2) > 0 {
        v2[0] = v1
    }
    var v15 int16 = 1000
    fmt.Println(v1, v3)
    if true {
        var v16 int64 = 0
        _ = v16
    }
    v11 = 543
    switch 9223372036854775807 {
    case -1:
        v15 -= -328
    case 1:
        for fi17 := 0; fi17 < 6; fi17++ {
            if false {
                for fi18 := 0; fi18 < 10; fi18++ {
                    if false {
                        v11 = v11
                    }
                    var v19 int32 = -2147483648
                    var v20 string = ((v3 + "") + (v3 + v3))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v21 uint32 = v1
                _ = v21
            } else {
                for fi22 := 0; fi22 < 3; fi22++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v3 = ("" + "\n\t")
                    _ = fi22
                }
            }
            _ = fi17
        }
    default:
        for fi23 := 0; fi23 < 17; fi23++ {
            for fi24 := 0; fi24 < 5; fi24++ {
                fmt.Println(v1, v3)
                _ = fi24
            }
            _ = fi23
        }
    }
    v25 := GenS1{F0: 128}
    switch 127 {
    case 454:
        if false {
            v3 = (("" + v3) + ("foo bar" + v3))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi26 := 0; fi26 < 12; fi26++ {
                var v27 int8 = 42
                _ = fi26
                _ = v27
            }
        }
    case 42:
        v3 = (v3 + v3)
    default:
        var v28 float64 = (1.5 - (((((0.0 - (1.5 + (1.5 * (3.14159 - 1e10)))) * 3.14159) - (((-2.25 / -2.25) / (float64(0.1) / 1.0)) * 0.1)) / 3.14159) / float64(1e10)))
        _ = v28
    }
    if len(v2) > 0 {
        v2[0] = v1
    }
    var v29 string = ("foo bar" + v3)
    switch 2147483647 {
    case -2147483648:
        for fi30 := 0; fi30 < 19; fi30++ {
            v2 = append(v2, 2147483648)
            _ = fi30
        }
    case 4294967295:
        _ = len(v2)
    case 0:
        if true {
            var v31 uint = v11
            _ = v31
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    default:
        var v32 int8 = -1
        _ = v32
    }
    var v33 int8 = 42
    if len(v2) > 0 {
        v2[0] = v1
    }
    switch v29 {
    case "foo bar":
        for fi34 := 0; fi34 < 3; fi34++ {
            for fi35 := 0; fi35 < 16; fi35++ {
                switch fi35 {
                case -2147483648:
                    var v36 int8 = -128
                    _ = v36
                case -9223372036854775808:
                    v37 := GenS3{F0: 254, F1: 2147483647, F2: 1}
                    _ = v37
                default:
                    var v38 uint8 = 221
                    _ = v38
                }
                v15 -= -32768
                var v39 uint64 = 122
                v29 = v3
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v40 uint64 = 1
                _ = fi35
                _ = v39
                _ = v40
            }
            var v41 uint8 = 1
            _ = fi34
            _ = v41
        }
    case "hello":
        v2 = append(v2, (^(v1)))
    case "\n\t":
        var v42 int = 32768
        _ = v42
    default:
        var v43 uint16 = 624
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v33
}
