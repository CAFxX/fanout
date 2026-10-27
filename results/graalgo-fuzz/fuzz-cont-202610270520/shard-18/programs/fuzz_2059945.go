package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]uint32)
    var v3 int64 = 1
    var v4 uint8 = (^(v1))
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(v3)
        switch "a" {
        case "a":
            var v6 int8 = -128
            _ = v6
        case "":
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 int32 = 553
                v3 -= v3
                v4 = (v4 & v1)
                fmt.Println(fi5, v1, v3)
                _ = fi7
                _ = v8
            }
        case "foo bar":
            for fi9 := 0; fi9 < 7; fi9++ {
                v4 -= v1
                _ = fi9
            }
        default:
            v4 -= 42
        }
        _ = fi5
    }
    if false {
        for fi10 := 0; fi10 < 2; fi10++ {
            fmt.Println(v4, v4)
            var v11 int32 = 1
            v12 := []int{fi10, fi10, 4294967295, (fi10 & -878)}
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    if (511 != 240) {
        switch "" {
        case "foo bar":
            var v13 int64 = (v3 << ((141) & 63))
            _ = v13
        case "a":
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 bool = false
                v16 := GenS1{F0: v1}
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            v17 := GenS1{F0: (v1 ^ v1)}
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v2[fi18] = 2
        _ = fi18
    }
    v4 = (v1 - v4)
    switch "" {
    case "foo bar":
        var v19 int = 2147483647
        _ = v19
    }
    var v20 string = ("foo bar" + "")
    v21 := []int{}
    if (59 > 0) {
        for fi22 := 0; fi22 < 14; fi22++ {
            if (690 <= 661) {
                var v23 int = -9223372036854775808
                _ = v23
            } else {
                var v24 uint16 = 2
                _ = v24
            }
            var v25 int64 = (-(v3))
            _ = fi22
            _ = v25
        }
        if (1.5 >= 1.5) {
            if len(v21) > 0 {
                v21[0] = 9223372036854775807
            }
            fmt.Println(v3, v20)
            for fi26 := 0; fi26 < 4; fi26++ {
                fmt.Println(fi26, fi26)
                if ((103 < 42) && (true && true)) {
                    v27 := v2[fi26]
                    v28 := GenS1{F0: (v1 ^ v1)}
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 10; fi29++ {
                    v30 := GenS1{F0: 193}
                    v31 := GenS1{F0: 1}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = fi26
            }
        }
    }
    switch 128 {
    case -9223372036854775808:
        var v32 uint8 = (v1 & (v1 + (v1 / ((v1) | 1))))
        _ = v32
    case -527:
        switch 0 {
        case 0:
            var v33 uint16 = 65535
            _ = v33
        case -1:
            fmt.Println(v4)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
