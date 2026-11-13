package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int64 }
type GenS2 struct { F0 int; F1 uint }
func genFn0(p0 int, p1 uint8, p2 int32) int32 {
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 int16 = 1
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        switch 32767 {
        case 32768:
            p2 += (p2 & p2)
        }
        _ = fi10
    }
    var v11 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (^(p2))
}
func genFn1(p0 int8, p1 uint16) uint32 {
    var v12 float64 = 3.14159
    var v13 bool = false
    var v14 float32 = 0.1
    switch 127 {
    case 127:
        for fi15 := 0; fi15 < 2; fi15++ {
            switch "foo bar" {
            case "hello":
                for fi16 := 0; fi16 < 14; fi16++ {
                    v17 := []int{fi16, (-(fi15))}
                    _ = fi16
                    _ = v17
                }
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                v19 := []int{fi18, (fi18 * (fi15 + 32767)), fi18, (fi15 & 9223372036854775807)}
                if (p1 == p1) {
                    switch "\n\t" {
                    case "a":
                        _ = len(v19)
                    case "foo bar":
                        var v20 int16 = 682
                        _ = v20
                    }
                    fmt.Println(p0)
                    v12 = (v12 + 3.14159)
                }
                _ = fi18
                _ = v19
            }
            _ = fi15
        }
    case -891:
        var v21 uint8 = 0
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return 686
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{9223372036854775807}
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    var v5 float64 = float64((((1.5 / 1.0) / 1.5) / 1.5))
    v6 := make(map[uint32]uint32)
    var v22 bool = (184 != -32768)
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if (-1 != -9223372036854775808) {
        v2 = append(v2, 571)
        v23 := GenS1{F0: (v1 << ((247) & 63)), F1: 4294967295, F2: v1}
        _ = v23
    } else {
        var v24 int = 425
        var v25 string = "\n\t"
        v24 = 32767
        _ = v24
        _ = v25
    }
    v1 += -65
    v5 += (v5 / v5)
    if true {
        v22 = true
    } else {
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 int8 = 127
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 uint16 = 263
                _ = fi28
                _ = v29
            }
            _ = fi26
            _ = v27
        }
    }
    var v30 int8 = 0
    for fi31 := 0; fi31 < 8; fi31++ {
        if (!(fi31 > fi31)) {
            switch -820 {
            case 411:
                for fi32 := 0; fi32 < 13; fi32++ {
                    v33 := GenS1{F0: v1, F1: 334, F2: v1}
                    v5 -= (v5 + 0.0)
                    _ = fi32
                    _ = v33
                }
            case -1:
                v30 = v30
            case 940:
                if (false && true) {
                    var v34 int8 = (v30 + v30)
                    v35 := GenS2{F0: 32767, F1: 65535}
                    v34 += v34
                    _ = v34
                    _ = v35
                }
            }
        } else {
            delete(v3, (fi31 - 32767))
            if (255 > 32767) {
                var v36 int16 = 32767
                _ = v36
            } else {
                var v37 uint64 = 548
                switch fi31 {
                case 9223372036854775807:
                    v37 += 2
                case 42:
                    var v38 uint64 = v37
                    _ = v38
                case 32767:
                    v1 = (v1 + (v1 | 307))
                default:
                    v39 := GenS2{F0: (-(fi31)), F1: 42}
                    _ = v39
                }
                _ = v37
            }
        }
        _ = fi31
    }
    delete(v4, "hello")
    var v40 int32 = 1
    var v41 float64 = float64(1e10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
