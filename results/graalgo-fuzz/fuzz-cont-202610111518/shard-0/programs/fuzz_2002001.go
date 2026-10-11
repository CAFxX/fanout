package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint }

func main() {
    var v1 uint64 = 425
    v2 := []string{("" + "")}
    var v3 uint16 = 255
    var v4 int32 = -2147483648
    v3 = (v3 ^ 42)
    fmt.Println(v1, v4, v4)
    var v5 int8 = 0
    v4 += (-(v4))
    if (-1 >= -1) {
        v5 -= v5
        v5 -= 42
    }
    v6 := GenS1{F0: 65535, F1: -9223372036854775808, F2: 1000}
    var v7 int32 = (-(v4))
    var v8 int16 = 446
    var v9 uint64 = 42
    if (("a" + "") <= "a") {
        switch 896 {
        case 0:
            switch -9223372036854775808 {
            case -9223372036854775808:
                var v10 string = (("foo bar" + "") + "a")
                _ = v10
            case 785:
                for fi11 := 0; fi11 < 4; fi11++ {
                    if true {
                        var v12 uint64 = 0
                        v13 := GenS1{F0: 0, F1: -9223372036854775808, F2: v8}
                        if false {
                            v9 = 0
                            for fi14 := 0; fi14 < 1; fi14++ {
                                v4 -= (v4 & int32(v1))
                                if ((v3 < 42) || (!(((127 <= 145) || (v5 != (v5 >> ((4294967295) & 63)))) && true))) {
                                    var v15 int = fi11
                                    var v16 int8 = (v5 ^ (v5 + (v5 - int8(v3))))
                                    _ = v15
                                    _ = v16
                                }
                                _ = fi14
                            }
                        } else {
                            v1 = (v1 & (^(v12)))
                            v17 := GenS2{F0: 0, F1: 18446744073709551615}
                            _ = v17
                        }
                        _ = v12
                        _ = v13
                    } else {
                        v18 := GenS2{F0: 80, F1: uint(v3)}
                        _ = v18
                    }
                    v9 -= v9
                    var v19 uint16 = v3
                    _ = fi11
                    _ = v19
                }
            case -187:
                var v20 int = 9223372036854775807
                _ = v20
            }
        case -9223372036854775808:
            v8 += v8
        case 4294967295:
            v9 = v9
        default:
            v7 -= 0
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v7 -= v7
    }
    fmt.Println(v8, v5, v7)
    if false {
        if false {
            v4 -= (v7 | int32(v8))
            var v21 bool = (int(v4) > int(v3))
            _ = v21
        }
        var v22 uint16 = (v3 ^ 362)
        switch 560 {
        case 32768:
            switch 32767 {
            case -9223372036854775808:
                switch "\n\t" {
                case "a":
                    v3 += v3
                case "hello":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v23 uint64 = (v9 * v1)
                    _ = v23
                }
            }
        default:
            var v24 bool = (1 >= 255)
            _ = v24
        }
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
