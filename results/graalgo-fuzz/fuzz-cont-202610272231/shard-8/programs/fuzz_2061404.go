package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint64; F1 int8; F2 int8 }

func main() {
    var v1 uint64 = 1
    var v2 uint16 = 65535
    v3 := []string{""}
    var v4 float64 = (0.1 + 0.0)
    v5 := GenS1{F0: 2, F1: 9223372036854775807}
    for fi6 := 0; fi6 < 10; fi6++ {
        if ((!(!true)) && (1 == 4294967295)) {
            v7 := GenS2{F0: 687}
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 int64 = -9223372036854775808
                v10 := GenS3{F0: v2, F1: true}
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        }
        var v11 float64 = v4
        var v12 float64 = float64(0.1)
        _ = fi6
        _ = v11
        _ = v12
    }
    switch -9223372036854775808 {
    case -505:
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 int16 = 0
            _ = fi13
            _ = v14
        }
    default:
        v2 = v2
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 bool = (!(fi15 >= fi15))
        var v17 int32 = int32(v1)
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 uint64 = (v1 << ((559) & 63))
        _ = fi18
        _ = v19
    }
    if true {
        if true {
            v20 := GenS2{F0: 602}
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 bool = true
        v23 := GenS4{F0: 47, F1: 1, F2: -1}
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v2, v1)
    v1 += (v1 >> ((18446744073709551615) & 63))
    for fi24 := 0; fi24 < 15; fi24++ {
        v2 = (v2 - v2)
        fmt.Println(v1)
        _ = fi24
    }
    switch -2147483648 {
    case -744:
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    case 1:
        switch -1 {
        case 32768:
            var v25 float32 = (((0.0 / float32(v4)) * ((0.1 + 0.1) * (0.0 / 1e10))) / 1.5)
            _ = v25
        case -201:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 94:
            if (2147483647 > -909) {
                for fi26 := 0; fi26 < 10; fi26++ {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        switch ((("\n\t" + "\n\t") + "\n\t") + "foo bar") {
                        case "foo bar":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        v28 := GenS1{F0: 938, F1: fi26}
                        _ = fi27
                        _ = v28
                    }
                    if len(v3) > 0 {
                        v3[0] = (("a" + "") + ("foo bar" + ("hello" + "hello")))
                    }
                    var v29 uint16 = v2
                    _ = fi26
                    _ = v29
                }
                v2 = v2
                for fi30 := 0; fi30 < 13; fi30++ {
                    for fi31 := 0; fi31 < 19; fi31++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v32 float64 = (0.0 * (v4 + ((v4 - v4) - 0.0)))
                        _ = fi31
                        _ = v32
                    }
                    _ = fi30
                }
            } else {
                v33 := GenS3{F0: v2, F1: true}
                var v34 uint8 = 1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v33
                _ = v34
            }
        default:
            fmt.Println(v2)
        }
    case 93:
        var v35 uint16 = v2
        _ = v35
    }
    v4 = v4
    v1 = 666
    for fi36 := 0; fi36 < 11; fi36++ {
        var v37 int = (fi36 % (((fi36 | (-(fi36)))) | 1))
        _ = fi36
        _ = v37
    }
    var v38 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
