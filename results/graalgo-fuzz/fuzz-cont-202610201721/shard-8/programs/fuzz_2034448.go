package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int) int32 {
    var v5 float32 = 1e10
    fmt.Println(v5, v5)
    var v6 uint64 = 18446744073709551615
    _ = p0
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 int64 = 42
    var v2 int8 = 100
    var v3 uint8 = 2
    v4 := make(map[uint32]int)
    if true {
        var v7 uint = 605
        _ = v7
    } else {
        switch 128 {
        case -592:
            var v8 uint16 = 52
            _ = v8
        case 2147483647:
            v3 = v3
        case 679:
            v3 = (v3 - 255)
        }
    }
    v9 := []int{42, 127, -2147483648, -2147483648}
    v2 -= v2
    v2 = (v2 >> ((2) & 63))
    v1 -= (v1 >> ((872) & 63))
    v3 += (v3 - (v3 * (v3 << ((4294967295) & 63))))
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int8 = 100
            if len(v9) > 0 {
                v9[0] = (-(fi11))
            }
            _ = fi11
            _ = v12
        }
        var v13 int16 = 0
        _ = fi10
        _ = v13
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        if false {
            v14 := GenS1{F0: 734, F1: 4294967295, F2: 4294967295}
            _ = v14
        }
    case "":
        fmt.Println(v1, v3, v1)
    case "hello":
        fmt.Println(v3, v1, v2)
    default:
        fmt.Println(v3, v3, v2)
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v15 int8 = int8(v3)
    fmt.Println(v3, v3, v15)
    switch -9223372036854775808 {
    case 9223372036854775807:
        if (4294967295 < 0) {
            var v16 uint8 = v3
            _ = v16
        } else {
            _ = len(v9)
        }
    case 0:
        var v17 bool = (!(!(!((-(v15)) == (v15 << ((65535) & 63))))))
        _ = v17
    case 154:
        var v18 uint8 = (v3 - v3)
        _ = v18
    }
    if true {
        if (false || (!(!(false || (((^(v3)) > 102) && true))))) {
            var v19 int32 = 279
            _ = v19
        } else {
            v3 -= v3
            switch 128 {
            case 9223372036854775807:
                var v20 bool = (1 >= -921)
                _ = v20
            case -9223372036854775808:
                var v21 int32 = 2147483647
                _ = v21
            case 4294967295:
                switch (("hello" + ("foo bar" + "foo bar")) + "foo bar") {
                case "foo bar":
                    if true {
                        if (false && false) {
                            var v22 int = 127
                            v1 = v1
                            _ = v22
                        } else {
                            v23 := GenS1{F0: 65535, F1: 42, F2: 1}
                            if (1 >= 2147483647) {
                                v24 := GenS2{F0: 791, F1: 329}
                                v1 = 362
                                v15 -= v2
                                _ = v24
                            } else {
                                v15 += (v2 - (^(v15)))
                            }
                            delete(v4, 65535)
                            _ = v23
                        }
                    }
                case "hello":
                    v2 -= (v2 ^ (v2 + (v2 % (((v15 * v2)) | 1))))
                }
            default:
                var v25 uint64 = 191
                _ = v25
            }
            var v26 int32 = genFn0(-2147483648)
            _ = v26
        }
        v15 -= (v15 | (v15 | (v15 & v2)))
        fmt.Println(v3, v15, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
