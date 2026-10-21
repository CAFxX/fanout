package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int16 = -1
    var v2 uint = 357
    var v3 float64 = 1.5
    v4 := []int{-9223372036854775808, -2147483648}
    fmt.Println(v2, v1)
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(fi5)
        fmt.Println(v1)
        v2 = 255
        _ = fi5
    }
    if (!true) {
        if true {
            switch 561 {
            case 42:
                v1 = v1
            }
            switch 9223372036854775807 {
            case -148:
                v2 = 2
            case 676:
                if false {
                    v4 = append(v4, -1)
                }
            case 42:
                v2 = 18446744073709551615
            default:
                var v6 int32 = -751
                _ = v6
            }
            v3 -= v3
        } else {
            var v7 int32 = -1
            v8 := GenS1{F0: 77, F1: 9223372036854775807}
            if false {
                for fi9 := 0; fi9 < 12; fi9++ {
                    switch "" {
                    case "\n\t":
                        var v10 float32 = 1e10
                        _ = v10
                    case "foo bar":
                        v1 = v1
                    default:
                        _ = len(v4)
                    }
                    _ = fi9
                }
                switch ("\n\t" + "hello") {
                case "foo bar":
                    var v11 uint32 = 2147483647
                    _ = v11
                case "":
                    v1 -= (^(v1))
                case "a":
                    v12 := GenS1{F0: 0, F1: 9223372036854775807}
                    _ = v12
                default:
                    var v13 uint16 = 65535
                    _ = v13
                }
            }
            _ = v7
            _ = v8
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch 2147483647 {
    case -364:
        v1 -= v1
    case 743:
        var v14 int = -322
        _ = v14
    default:
        var v15 bool = (123 <= 4294967295)
        _ = v15
    }
    var v16 int = 42
    if false {
        var v17 int8 = 127
        _ = v17
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if ((!(false || true)) && ((!false) || false)) {
        for fi18 := 0; fi18 < 16; fi18++ {
            if (v1 > (v1 & v1)) {
                if false {
                    fmt.Println(v16, fi18, v3)
                }
                v16 = -2147483648
            }
            if (0 <= -9223372036854775808) {
                switch -2147483648 {
                case 127:
                    v1 -= v1
                default:
                    for fi19 := 0; fi19 < 14; fi19++ {
                        v20 := GenS1{F0: 0, F1: 121}
                        if ((!(331 <= 42)) && false) {
                            v4 = append(v4, fi18)
                        }
                        v16 = 2147483647
                        fmt.Println(fi19, v16)
                        var v21 int8 = 100
                        if false {
                            var v22 uint16 = 2
                            v22 = v22
                            _ = v22
                        } else {
                            v2 = v2
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                for fi23 := 0; fi23 < 4; fi23++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi23
                }
            }
            _ = fi18
        }
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 bool = false
            for fi26 := 0; fi26 < 2; fi26++ {
                v27 := GenS2{F0: 18446744073709551615, F1: 0}
                var v28 uint32 = 65535
                var v29 uint = (v2 | (v2 >> ((65535) & 63)))
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            _ = fi24
            _ = v25
        }
    }
    if len(v4) > 0 {
        v4[0] = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
