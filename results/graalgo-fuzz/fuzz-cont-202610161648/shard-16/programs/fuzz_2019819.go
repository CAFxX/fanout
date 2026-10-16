package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }
type GenS2 struct { F0 string; F1 int32; F2 bool }
type GenS3 struct { F0 uint64; F1 int32; F2 uint }

func main() {
    var v1 bool = (2 < 131)
    v2 := []int{}
    var v3 uint64 = 42
    var v4 int8 = 127
    v5 := make(map[int]int)
    v6 := make(map[uint32]uint32)
    v3 -= (v3 + v3)
    v7 := GenS1{F0: 0, F1: v1, F2: -32768}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v8 := GenS2{F0: ("" + "a"), F1: -537, F2: true}
    var v9 uint = 65535
    if len(v2) > 0 {
        v2[0] = -750
    }
    fmt.Println(v4)
    switch 32768 {
    case 246:
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    case -506:
        delete(v6, 1)
    default:
        if v1 {
            v4 = (-(v4))
            switch 372 {
            case -583:
                v9 = v9
            case -5:
                switch -136 {
                case 2147483647:
                    v4 = (v4 * 127)
                case 972:
                    fmt.Println(v9, v4)
                default:
                    v10 := GenS1{F0: -159, F1: (!v1), F2: int16(v3)}
                    _ = v10
                }
            case 4294967295:
                v3 = v3
            default:
                v11 := GenS1{F0: 0, F1: v1, F2: -32768}
                _ = v11
            }
        }
    }
    var v12 string = "a"
    _ = len(v6)
    var v13 uint16 = 65535
    switch 1 {
    case 32768:
        var v14 float32 = float32((0.0 - ((0.0 - -2.25) + (0.1 * 1.5))))
        _ = v14
    case 1:
        for fi15 := 0; fi15 < 9; fi15++ {
            if false {
                v16 := GenS3{F0: 42, F1: 2147483647, F2: 255}
                _ = v16
            } else {
                var v17 int = -211
                v17 += -1
                _ = v17
            }
            switch v12 {
            case "hello":
                v4 = (v4 % (((v4 >> (((v9 | 245)) & 63))) | 1))
            case "a":
                if len(v2) > 0 {
                    v2[0] = -2147483648
                }
            }
            for fi18 := 0; fi18 < 18; fi18++ {
                v9 = (^(v9))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v19 int = 9223372036854775807
                v4 = 1
                v1 = (3.14159 == ((0.0 - 0.1) * (((1.5 * 0.1) / 1.0) + 0.0)))
                v13 = v13
                _ = fi18
                _ = v19
            }
            _ = fi15
        }
    case 128:
        if len(v2) > 0 {
            v2[0] = 128
        }
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        switch (fi20 - (fi20 / (((fi20 ^ (fi20 * (fi20 - fi20)))) | 1))) {
        case 0:
            var v21 uint32 = 2147483648
            _ = v21
        case 2147483647:
            var v22 int32 = 42
            _ = v22
        case -716:
            v23 := GenS1{F0: -1, F1: false, F2: -1}
            _ = v23
        }
        _ = fi20
    }
    v9 = (v9 >> ((2) & 63))
    if false {
        if (!v1) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v24 := v6[1]
        _ = v24
    } else {
        switch 9223372036854775807 {
        case 32767:
            fmt.Println(v1)
        case 2147483647:
            if v1 {
                switch (((v12 + ("" + v12)) + "hello") + "hello") {
                case "hello":
                    var v25 int8 = v4
                    _ = v25
                default:
                    v26 := GenS1{F0: 967, F1: (!v1), F2: 1000}
                    _ = v26
                }
            }
        default:
            var v27 int32 = 298
            _ = v27
        }
    }
    v12 = ("hello" + v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
