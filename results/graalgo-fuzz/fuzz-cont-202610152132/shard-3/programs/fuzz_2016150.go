package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 190
    v3 := []uint32{1, 42, 2147483647}
    v4 := make(map[int]int)
    if len(v3) > 0 {
        v3[0] = 0
    }
    if false {
        delete(v4, 32767)
        if true {
            fmt.Println(v2)
            switch 0 {
            case -9223372036854775808:
                var v5 float64 = -2.25
                _ = v5
            default:
                v2 = (v2 >> ((65535) & 63))
            }
        }
        if true {
            for fi6 := 0; fi6 < 15; fi6++ {
                var v7 uint16 = 760
                _ = fi6
                _ = v7
            }
        } else {
            v2 -= v2
            var v8 uint64 = 0
            if (!true) {
                v9 := GenS1{F0: 312}
                var v10 uint8 = v2
                v11 := GenS2{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: 647, F2: 9223372036854775807}
                _ = v10
                _ = v11
                _ = v9
            }
            _ = v8
        }
    }
    if ((1.5 + 3.14159) < (-2.25 - 3.14159)) {
        switch ("\n\t" + "a") {
        case "a":
            v12 := GenS2{F0: "\n\t", F1: 18446744073709551615, F2: -1}
            _ = v12
        case "hello":
            v2 += (v2 * v2)
        default:
            var v13 int16 = 32767
            _ = v13
        }
        var v14 int64 = 1
        _ = v14
    }
    switch 9223372036854775807 {
    case 4294967295:
        _ = len(v3)
    case -200:
        var v15 uint16 = 1
        _ = v15
    case 128:
        var v16 bool = (2147483647 > 2147483648)
        _ = v16
    default:
        var v17 int32 = 0
        _ = v17
    }
    v2 += v2
    for fi18 := 0; fi18 < 1; fi18++ {
        if true {
            var v19 int8 = 42
            _ = v19
        } else {
            if len(v3) > 0 {
                v3[0] = 622
            }
            if (!(false && true)) {
                v20 := GenS2{F0: "\n\t", F1: 0, F2: 0}
                var v21 string = "foo bar"
                v2 += v2
                _ = v20
                _ = v21
            } else {
                switch 42 {
                case 42:
                    v2 = (v2 - 25)
                default:
                    var v22 float32 = (0.1 - 3.14159)
                    _ = v22
                }
                var v23 uint8 = 154
                _ = v23
            }
            v2 += v2
        }
        if (false || (9223372036854775807 < -1)) {
            var v24 uint16 = 724
            if (!true) {
                switch fi18 {
                case 0:
                    if (true || (1.5 > ((3.14159 + 0.1) - 0.1))) {
                        v2 -= 233
                    }
                default:
                    v24 = v24
                }
                fmt.Println(v24)
            }
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 int16 = -32768
                v2 -= 255
                _ = fi25
                _ = v26
            }
            _ = v24
        } else {
            if true {
                v27 := GenS3{F0: 889}
                if false {
                    fmt.Println(v2, fi18)
                    v2 -= (v2 * 113)
                }
                _ = v27
            }
            var v28 float32 = (3.14159 / 1.5)
            fmt.Println(v28, v28, v2)
            _ = v28
        }
        v2 = 127
        _ = fi18
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        v2 -= (^(v2))
        _ = fi29
    }
    v2 += 145
    v2 = v2
    v30 := GenS4{F0: 282}
    v1[8] = uint32(v2)
    v31 := GenS1{F0: -9223372036854775808}
    v2 += 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
