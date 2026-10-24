package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int32, p1 string) uint8 {
    p1 = p1
    v4 := []int{0, -438, 127, -2147483648}
    _ = p0
    _ = p1
    _ = v4
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 113
    var v3 int64 = 1
    v5 := v1[2]
    if ((^(v2)) > (v2 * v2)) {
        v6 := GenS1{F0: 255}
        v7 := []int{848}
        for fi8 := 0; fi8 < 5; fi8++ {
            v1[v5] = (v5 % ((v5) | 1))
            switch "foo bar" {
            case "":
                v1[v5] = (v5 - (v5 << ((65535) & 63)))
            case "foo bar":
                v9 := GenS1{F0: 255}
                _ = v9
            }
            _ = fi8
        }
        _ = v6
        _ = v7
    }
    v10 := []int{993, 32768, -826}
    if (1.5 <= (float64(0.1) / (((0.0 - 0.1) * -2.25) * 1.5))) {
        for fi11 := 0; fi11 < 16; fi11++ {
            v3 = v3
            var v12 int8 = v2
            _ = fi11
            _ = v12
        }
    }
    v5 -= v5
    for fi13 := 0; fi13 < 17; fi13++ {
        v3 -= int64(fi13)
        _ = fi13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS1{F0: 942}
        switch "\n\t" {
        case "\n\t":
            v1[v5] = 2147483648
        }
        var v16 uint64 = 144
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 bool = (486 != 4294967295)
    for fi18 := 0; fi18 < 19; fi18++ {
        if v17 {
            var v19 int16 = 0
            if (v17 || v17) {
                fmt.Println(v3, v5, fi18)
            } else {
                for fi20 := 0; fi20 < 16; fi20++ {
                    v10 = append(v10, (fi18 | (fi18 ^ fi20)))
                    if v17 {
                        v10 = append(v10, (fi20 % (((fi20 >> ((486) & 63))) | 1)))
                        if v17 {
                            v2 = v2
                        }
                    } else {
                        v3 = -9223372036854775808
                        var v21 bool = false
                        _ = v21
                    }
                    switch 32768 {
                    case 0:
                        var v22 uint = 255
                        _ = v22
                    default:
                        var v23 int32 = 321
                        _ = v23
                    }
                    _ = fi20
                }
            }
            for fi24 := 0; fi24 < 17; fi24++ {
                v1[(^(v5))] = 342
                v25 := GenS1{F0: 255}
                _ = fi24
                _ = v25
            }
            _ = v19
        } else {
            if (false || (!false)) {
                v5 = v5
                v17 = v17
            }
        }
        _ = fi18
    }
    v26 := GenS1{F0: 255}
    v27 := GenS1{F0: 42}
    v3 = (v3 >> ((1) & 63))
    var v28 int64 = v3
    v1[(v5 % (((v5 & v5)) | 1))] = (v5 + v5)
    v29 := GenS1{F0: uint16(v5)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v5
}
