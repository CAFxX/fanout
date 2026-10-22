package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 bool }
func genFn0(p0 int8) uint16 {
    v7 := []int{32767}
    v8 := GenS1{F0: 42, F1: false}
    var v9 uint8 = 127
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            switch fi11 {
            case -298:
                switch fi11 {
                case 9223372036854775807:
                    if true {
                        switch (("\n\t" + "\n\t") + "a") {
                        case "hello":
                            var v12 int64 = -141
                            _ = v12
                        }
                        v9 = (v9 * v9)
                    }
                default:
                    var v13 int = fi11
                    p0 += (p0 % ((p0) | 1))
                    v13 -= v13
                    _ = v13
                }
            case 2147483647:
                v14 := GenS1{F0: v9, F1: (-2147483648 < 0)}
                _ = v14
            default:
                fmt.Println(fi11)
            }
            _ = fi11
        }
        v9 -= v9
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS2{F0: 1}
            _ = fi15
            _ = v16
        }
        _ = fi10
    }
    fmt.Println(v9, v9)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []int{1, 122, 381}
    var v2 int8 = 100
    v3 := []string{"", ""}
    var v4 int8 = v2
    v5 := make(map[int]int)
    var v6 uint = 18446744073709551615
    v4 = v2
    v5[94] = 915
    var v17 uint = (v6 * (v6 - 255))
    fmt.Println(v2, v4)
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 uint8 = 217
        v20 := GenS1{F0: v19, F1: (12 != (v2 + (v4 * v4)))}
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 string = "foo bar"
    if (1 < 1000) {
        v2 = v4
        v2 = -1
    } else {
        _ = len(v3)
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := GenS3{F0: ((32767 < 32767) || true)}
            _ = fi22
            _ = v23
        }
        v4 = (v2 << ((4294967295) & 63))
    }
    var v24 float64 = (-2.25 * (float64(-2.25) * (-2.25 * 0.0)))
    switch 1 {
    case 166:
        v25 := v5[-9223372036854775808]
        _ = v25
    }
    var v26 uint8 = 1
    v4 = v2
    _ = len(v3)
    delete(v5, -9223372036854775808)
    if (668 <= -32768) {
        v6 -= v6
        v27 := GenS4{F0: -77, F1: true}
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
