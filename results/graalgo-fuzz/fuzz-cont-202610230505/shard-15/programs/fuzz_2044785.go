package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64, p1 uint8, p2 bool) int32 {
    v5 := []int{128, 464, 42}
    p0 = 18446744073709551615
    fmt.Println(p2)
    fmt.Println(p0, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2147483647
}

func main() {
    v1 := []uint32{2147483647, 1}
    var v2 int8 = 78
    var v3 int64 = -9223372036854775808
    v4 := []int{127, 472, 2147483647, 42}
    var v6 int = 42
    v2 = v2
    v6 -= 9223372036854775807
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 -= v6
        var v8 uint32 = 919
        v1 = append(v1, 193)
        _ = fi7
        _ = v8
    }
    var v9 float64 = -2.25
    var v10 uint32 = 2
    var v11 int = 2147483647
    v3 -= -9223372036854775808
    v12 := GenS1{F0: -681}
    v1 = append(v1, v10)
    var v13 uint = 1
    if true {
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 5; fi15++ {
                v6 += 9223372036854775807
                _ = fi15
            }
            _ = fi14
        }
        var v16 int8 = (v2 << ((65535) & 63))
        var v17 uint32 = 42
        _ = v16
        _ = v17
    }
    switch (v11 & int(v3)) {
    case 42:
        v10 += (v10 / (((v10 - (v10 >> ((v13) & 63)))) | 1))
    case 127:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 int16 = 42
                fmt.Println(v2, v20, fi18)
                fmt.Println(v10)
                v21 := GenS2{F0: 991, F1: 0}
                var v22 int = 127
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            var v23 int64 = (v3 >> (((^(v13))) & 63))
            v10 = v10
            _ = fi18
            _ = v23
        }
    }
    v6 = v6
    v24 := GenS1{F0: 249}
    v3 += -9223372036854775808
    if (1 == 632) {
        v11 = v11
        for fi25 := 0; fi25 < 12; fi25++ {
            v11 += (v6 * (v6 + int(v3)))
            _ = fi25
        }
        switch (v6 ^ 0) {
        case 615:
            v26 := GenS3{F0: 1}
            _ = v26
        case -2147483648:
            if true {
                fmt.Println(v11)
                switch "" {
                case "hello":
                    var v27 int32 = 520
                    _ = v27
                default:
                    v13 = 671
                }
                for fi28 := 0; fi28 < 17; fi28++ {
                    for fi29 := 0; fi29 < 20; fi29++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi29
                    }
                    var v30 string = ("a" + (("foo bar" + "hello") + "hello"))
                    v6 = v11
                    _ = fi28
                    _ = v30
                }
            } else {
                var v31 int8 = (-(v2))
                _ = v31
            }
        default:
            var v32 string = "\n\t"
            _ = v32
        }
    } else {
        for fi33 := 0; fi33 < 16; fi33++ {
            v6 = (v6 - (fi33 ^ v6))
            _ = fi33
        }
        if len(v4) > 0 {
            v4[0] = (^(v6))
        }
        var v34 uint64 = 4294967295
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
