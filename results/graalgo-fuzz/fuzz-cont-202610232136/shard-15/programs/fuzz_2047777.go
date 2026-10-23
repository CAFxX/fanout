package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 bool }
func genFn0() uint64 {
    v6 := GenS1{F0: 42, F1: "a", F2: true}
    var v7 int16 = -440
    var v8 uint32 = 1
    v9 := GenS1{F0: uint64(v8), F1: "\n\t", F2: false}
    if (!false) {
        v10 := []int{0}
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 uint) uint8 {
    if true {
        p0 = 804
        for fi11 := 0; fi11 < 7; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                p0 = (p0 % ((1) | 1))
                _ = fi12
            }
            switch "\n\t" {
            case "a":
                for fi13 := 0; fi13 < 11; fi13++ {
                    p0 = p0
                    if (("foo bar" + "hello") != "a") {
                        p0 = p0
                        var v14 int8 = 42
                        _ = v14
                    }
                    fmt.Println(fi11, fi13, fi11)
                    _ = fi13
                }
            default:
                var v15 float32 = 3.14159
                _ = v15
            }
            fmt.Println(p0, p0)
            _ = fi11
        }
    }
    p0 += (p0 & 65535)
    var v16 string = ("\n\t" + "a")
    _ = p0
    _ = v16
    return 132
}

func main() {
    v1 := []uint32{2147483647}
    v2 := make(map[string]uint32)
    var v3 bool = (!true)
    v4 := []int{587, -787, 330}
    var v5 int64 = -9223372036854775808
    var v17 int32 = 4
    v17 -= (v17 >> ((223) & 63))
    for fi18 := 0; fi18 < 13; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            fmt.Println(v17, v5, v5)
            var v20 int = fi19
            v21 := GenS1{F0: 4294967295, F1: "foo bar", F2: (2147483648 != 42)}
            _ = fi19
            _ = v20
            _ = v21
        }
        v2["foo bar"] = 999
        var v22 bool = v3
        _ = fi18
        _ = v22
    }
    v23 := GenS1{F0: 538, F1: "a", F2: v3}
    var v24 float64 = float64(3.14159)
    v25 := GenS1{F0: 830, F1: "", F2: (v3 || ((-2.25 / 3.14159) <= 3.14159))}
    fmt.Println(v24)
    v5 += -1
    v26 := GenS1{F0: 42, F1: "\n\t", F2: (!false)}
    v17 -= (v17 * (v17 >> ((4294967295) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
