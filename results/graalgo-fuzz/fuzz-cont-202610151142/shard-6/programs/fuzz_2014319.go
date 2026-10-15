package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int16; F2 int16 }
type GenS3 struct { F0 int; F1 int64; F2 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 int8 }

func main() {
    var v1 int32 = -1
    v2 := []int{-762, 551, -100, -798}
    var v3 string = "foo bar"
    var v4 int32 = (v1 << ((1) & 63))
    if ((float32((3.14159 / (0.0 - ((1e10 / 1.0) * ((1.5 * ((float64((0.1 + (-2.25 - 3.14159))) - 0.1) - (1e10 / 0.1))) - 0.0))))) - -2.25) == (3.14159 - (-2.25 - 0.1))) {
        v4 = v4
    }
    v5 := GenS1{F0: 9223372036854775807}
    switch -926 {
    case 128:
        var v6 uint8 = 2
        _ = v6
    case -747:
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(v3)
            for fi8 := 0; fi8 < 15; fi8++ {
                _ = len(v2)
                var v9 int = fi7
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    default:
        if true {
            v1 -= (v4 << ((42) & 63))
            if len(v2) > 0 {
                v2[0] = -9223372036854775808
            }
        } else {
            fmt.Println(v3)
        }
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        v1 = v1
    case 127:
        v3 = (v3 + v3)
    }
    var v10 int16 = 42
    var v11 string = "hello"
    v12 := GenS1{F0: -9223372036854775808}
    switch (("a" + "a") + ((v3 + v11) + (("\n\t" + ((v11 + v11) + ((v11 + v3) + v11))) + ""))) {
    case "foo bar":
        if true {
            v11 = v11
            fmt.Println(v4, v3)
            if true {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v14 := GenS1{F0: -997}
                    v2 = append(v2, (fi13 % ((fi13) | 1)))
                    _ = fi13
                    _ = v14
                }
            } else {
                var v15 uint64 = 4294967295
                var v16 int8 = 1
                _ = v15
                _ = v16
            }
        } else {
            v4 = v4
            var v17 bool = (v3 > "foo bar")
            _ = v17
        }
    default:
        var v18 int64 = 0
        _ = v18
    }
    if ((3.14159 * 3.14159) <= (1e10 / -2.25)) {
        if (-128 <= -1) {
            var v19 float32 = float32(((1e10 + 3.14159) - 1e10))
            _ = v19
        }
    } else {
        v20 := GenS2{F0: 1, F1: v10, F2: v10}
        for fi21 := 0; fi21 < 17; fi21++ {
            if (v4 > v1) {
                v10 += -32768
                switch v3 {
                case "hello":
                    if (!(!false)) {
                        v22 := GenS2{F0: 197, F1: v10, F2: v10}
                        fmt.Println(v11, fi21, fi21)
                        _ = v22
                    }
                case "\n\t":
                    if (2 > 39) {
                        v23 := GenS3{F0: 578, F1: -9223372036854775808, F2: int8(v4)}
                        v3 = v3
                        v11 = v3
                        _ = v23
                    }
                default:
                    v24 := GenS4{F0: fi21}
                    _ = v24
                }
            } else {
                v25 := GenS1{F0: -9223372036854775808}
                _ = v25
            }
            v1 -= 2147483647
            switch fi21 {
            case 127:
                if (!((!(!true)) && (178 >= 679))) {
                    v10 -= -32768
                    v11 = (v11 + v3)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    v26 := GenS5{F0: v10, F1: -110}
                    v3 = ""
                    _ = v26
                }
            case 87:
                v4 = (v1 % ((v4) | 1))
            default:
                v3 = v3
            }
            _ = fi21
        }
        v27 := GenS3{F0: 42, F1: 42, F2: 1}
        _ = v20
        _ = v27
    }
    fmt.Println(v3)
    v28 := GenS3{F0: -9223372036854775808, F1: 0, F2: 108}
    var v29 uint64 = 621
    var v30 string = (("\n\t" + "") + v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
