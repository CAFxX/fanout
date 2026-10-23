package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 2
    var v3 int32 = 0
    v4 := make(map[int]string)
    v5 := []int{4294967295, 0, 4294967295, -2147483648}
    var v6 int8 = -128
    v6 += int8(v2)
    v2 = (v2 >> ((65535) & 63))
    v5 = append(v5, -833)
    fmt.Println(v6)
    if true {
        for fi7 := 0; fi7 < 15; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := GenS1{F0: true, F1: (fi8 + (fi8 - fi7))}
                v6 = (v6 & (v6 / ((100) | 1)))
                v10 := GenS1{F0: (2 > 18446744073709551615), F1: 128}
                _ = fi8
                _ = v10
                _ = v9
            }
            v5 = append(v5, fi7)
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(v6, v2)
                v12 := GenS1{F0: (-2.25 == 3.14159), F1: 4294967295}
                v2 -= v2
                _ = fi11
                _ = v12
            }
            _ = fi7
        }
    }
    if (855 >= 65535) {
        var v13 uint = 255
        _ = v13
    } else {
        v1[("\n\t" + ("" + ""))] = "a"
        if ((v2 >= v2) && (false || ((!false) && (127 != 77)))) {
            if true {
                v14 := GenS2{F0: 65535}
                fmt.Println(v3)
                var v15 string = "hello"
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 8; fi16++ {
                switch (-(fi16)) {
                case 9223372036854775807:
                    var v17 bool = (v6 != (-(v6)))
                    _ = v17
                default:
                    var v18 int32 = (v3 % ((v3) | 1))
                    _ = v18
                }
                var v19 int32 = -258
                _ = fi16
                _ = v19
            }
        } else {
            v2 -= v2
        }
        switch 32767 {
        case -473:
            var v20 uint8 = 184
            _ = v20
        default:
            v3 -= (v3 * v3)
        }
    }
    v3 = v3
    v6 = 42
    var v21 int = 359
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch -1 {
    case 32767:
        v22 := v1["hello"]
        _ = v22
    case -2147483648:
        v6 = (v6 - v6)
    case 645:
        _ = len(v1)
    default:
        var v23 uint8 = 0
        _ = v23
    }
    var v24 int = v21
    fmt.Println(v21, v24, v24)
    for fi25 := 0; fi25 < 3; fi25++ {
        fmt.Println(fi25)
        if true {
            var v26 int32 = v3
            _ = v26
        }
        _ = fi25
    }
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
