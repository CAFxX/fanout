package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int16; F1 int32; F2 string }
func genFn0(p0 uint32) uint64 {
    var v7 uint32 = p0
    v8 := []int{}
    for fi9 := 0; fi9 < 3; fi9++ {
        if (42 <= 695) {
            fmt.Println(v7)
        }
        if len(v8) > 0 {
            v8[0] = fi9
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            switch ("hello" + "hello") {
            case "\n\t":
                var v11 float32 = 0.1
                _ = v11
            case "a":
                switch fi9 {
                case 127:
                    if len(v8) > 0 {
                        v8[0] = (fi9 * 2147483647)
                    }
                case 128:
                    if (18446744073709551615 >= 663) {
                        v8 = append(v8, (fi9 * fi9))
                    } else {
                        v7 -= v7
                    }
                default:
                    v8 = append(v8, 32767)
                }
            case "foo bar":
                p0 = (p0 ^ 4294967295)
            default:
                var v12 uint8 = 2
                _ = v12
            }
            var v13 bool = (182 < 4)
            _ = fi10
            _ = v13
        }
        _ = fi9
    }
    if (((!((-2147483648 > -1) || (100 != 18))) && ((890 != 156) || (-9223372036854775808 < -51))) || true) {
        v14 := GenS1{F0: (v7 | (p0 >> ((437) & 63))), F1: 2}
        p0 = (p0 >> ((42) & 63))
        _ = v14
    } else {
        var v15 float64 = 0.1
        _ = v15
    }
    v7 = v7
    _ = p0
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 32767
    var v2 uint16 = (v1 & v1)
    var v3 int32 = -2147483648
    v4 := make(map[uint32]bool)
    var v5 int = 0
    var v6 bool = false
    fmt.Println(v1, v5)
    fmt.Println(v3)
    fmt.Println(v2)
    var v16 int16 = 1000
    v17 := []int{(v5 + 32768), int(v3), (-(v5)), v5}
    v18 := GenS1{F0: 42, F1: 18446744073709551615}
    v1 = v2
    var v19 int = v5
    v20 := GenS1{F0: 1, F1: 42}
    v4[42] = v6
    var v21 uint = 65535
    switch "\n\t" {
    case "foo bar":
        for fi22 := 0; fi22 < 1; fi22++ {
            v4[65535] = ((v3 + (v3 | (v3 / (((v3 % (((-(v3))) | 1))) | 1)))) == (v3 * (v3 | (v3 | -1))))
            for fi23 := 0; fi23 < 11; fi23++ {
                v6 = v6
                _ = fi23
            }
            _ = fi22
        }
    }
    var v24 uint32 = 0
    v25 := GenS2{F0: int16(v21), F1: 1, F2: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
