package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint8) string {
    p0 = p0
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := []int{fi5, 9223372036854775807, 593, -692}
        _ = fi5
        _ = v6
    }
    switch (("foo bar" + (((("\n\t" + "\n\t") + "\n\t") + "hello") + "hello")) + "\n\t") {
    case "a":
        var v7 uint8 = 255
        _ = v7
    default:
        fmt.Println(p0)
    }
    var v8 uint32 = 2147483647
    _ = p0
    _ = v8
    return "a"
}
func genFn1(p0 int64) int32 {
    fmt.Println(p0)
    p0 -= (p0 >> ((255) & 63))
    _ = p0
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    var v4 int16 = 1
    v9 := []int{945, 2147483647, 813, 127}
    var v10 bool = (!true)
    switch 9223372036854775807 {
    case -1:
        if (-1 == -31) {
            for fi11 := 0; fi11 < 3; fi11++ {
                fmt.Println(v4)
                fmt.Println(v2, v2, v10)
                _ = fi11
            }
            v12 := GenS1{F0: 18446744073709551615}
            _ = v12
        }
    case 709:
        if (42 > 18446744073709551615) {
            var v13 int8 = 127
            _ = v13
        } else {
            v4 -= (v4 - v4)
        }
    case 973:
        var v14 uint16 = 65535
        _ = v14
    default:
        _ = len(v9)
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        v4 += (v4 + 32767)
        for fi16 := 0; fi16 < 19; fi16++ {
            if v10 {
                var v17 int8 = -115
                _ = v17
            }
            v4 -= v4
            v3[694] = (-(fi15))
            _ = fi16
        }
        v10 = (863 != 908)
        _ = fi15
    }
    switch 935 {
    case -52:
        fmt.Println(v2)
    case 2147483647:
        var v18 uint64 = 607
        _ = v18
    case -851:
        switch 32768 {
        case 9223372036854775807:
            var v19 int64 = -39
            _ = v19
        }
    }
    var v20 uint = 2
    for fi21 := 0; fi21 < 6; fi21++ {
        v10 = v10
        _ = fi21
    }
    if false {
        if v10 {
            v2 -= 1
            var v22 uint = v20
            _ = v22
        }
    }
    switch "a" {
    case "\n\t":
        if (1 <= 32767) {
            v20 += (v20 & (v20 | (v20 % ((18446744073709551615) | 1))))
        }
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 bool = ((false && v10) && v10)
        fmt.Println(v10)
        _ = fi23
        _ = v24
    }
    v2 += 67
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
