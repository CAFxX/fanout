package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int8 {
    if (!true) {
        switch -104 {
        case -1:
            var v6 float32 = 0.1
            _ = v6
        case -207:
            var v7 int64 = 9223372036854775807
            _ = v7
        case -735:
            var v8 string = "\n\t"
            _ = v8
        default:
            var v9 float64 = (((-2.25 - 0.1) + 1.5) / 0.1)
            _ = v9
        }
        switch "foo bar" {
        case "hello":
            switch 9223372036854775807 {
            case 9223372036854775807:
                var v10 uint64 = 445
                _ = v10
            }
        case "a":
            var v11 bool = (14 <= -1)
            _ = v11
        case "\n\t":
            var v12 float64 = ((1.5 * ((float64((-2.25 - float32(0.0))) / 1.5) - -2.25)) / 1.0)
            _ = v12
        default:
            if ((166 < 2) && true) {
                var v13 bool = true
                fmt.Println(v13)
                v14 := GenS1{F0: -2147483648}
                _ = v13
                _ = v14
            }
        }
        v15 := []int{}
        _ = v15
    } else {
        var v16 uint64 = 2
        for fi17 := 0; fi17 < 14; fi17++ {
            v16 = (v16 << ((4294967295) & 63))
            v18 := GenS1{F0: 2147483647}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 3; fi19++ {
            for fi20 := 0; fi20 < 5; fi20++ {
                switch "" {
                case "hello":
                    var v21 bool = false
                    _ = v21
                default:
                    v16 += v16
                }
                v16 += 4294967295
                _ = fi20
            }
            _ = fi19
        }
        _ = v16
    }
    var v22 int64 = 42
    _ = v22
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 uint8 = 128
    v4 := make(map[uint32]bool)
    var v5 uint = 18446744073709551615
    var v23 bool = false
    v1["foo bar"] = "a"
    switch 42 {
    case 0:
        if v23 {
            for fi24 := 0; fi24 < 11; fi24++ {
                var v25 uint32 = 230
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 1; fi26++ {
                v3 -= (v3 * 67)
                for fi27 := 0; fi27 < 6; fi27++ {
                    var v28 uint16 = 65535
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
        } else {
            if true {
                switch 1 {
                case 42:
                    var v29 uint8 = (v3 + (v3 << (((v5 << (((v5 / ((255) | 1))) & 63))) & 63)))
                    _ = v29
                case 127:
                    for fi30 := 0; fi30 < 14; fi30++ {
                        for fi31 := 0; fi31 < 17; fi31++ {
                            v23 = (4294967295 != 2)
                            var v32 uint64 = 2
                            _ = fi31
                            _ = v32
                        }
                        _ = fi30
                    }
                case 247:
                    for fi33 := 0; fi33 < 18; fi33++ {
                        var v34 string = "\n\t"
                        v3 += uint8(v5)
                        for fi35 := 0; fi35 < 14; fi35++ {
                            var v36 string = (v34 + "a")
                            _ = len(v4)
                            _ = fi35
                            _ = v36
                        }
                        _ = fi33
                        _ = v34
                    }
                }
            }
        }
    case -2147483648:
        v23 = (((!((18446744073709551615 <= 2) && (!(1 > -490)))) && v23) && false)
    case 42:
        for fi37 := 0; fi37 < 14; fi37++ {
            v38 := GenS1{F0: 2147483647}
            _ = fi37
            _ = v38
        }
    default:
        for fi39 := 0; fi39 < 17; fi39++ {
            fmt.Println(fi39, v23)
            _ = fi39
        }
    }
    var v40 uint64 = 4294967295
    v3 -= 118
    if len(v2) > 0 {
        v2[0] = ""
    }
    switch ("foo bar" + ("" + "\n\t")) {
    case "foo bar":
        v4[513] = (!(2 < 2))
    case "\n\t":
        v40 = (^(v40))
    }
    v2 = append(v2, "hello")
    v23 = (v3 > (v3 * (v3 << (((v5 >> ((808) & 63))) & 63))))
    v23 = (2147483647 <= int(v40))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
