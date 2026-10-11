package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int, p2 uint32) int16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 42, F1: 65535, F2: 402}
        v9 := []int{-1, p1, (^(fi7)), (p1 ^ (^(fi7)))}
        p1 -= int(p2)
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = 2
    var v10 uint64 = uint64(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 0
}
func genFn1(p0 bool, p1 int8, p2 int64) int {
    fmt.Println(p0, p0)
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    var v3 int16 = 32767
    var v4 float32 = (1.5 - 0.1)
    var v5 int32 = -1
    var v6 uint = 1
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    v4 -= v4
    var v13 int64 = 850
    for fi14 := 0; fi14 < 7; fi14++ {
        if (2 <= 4294967295) {
            for fi15 := 0; fi15 < 10; fi15++ {
                v13 = (v13 & v13)
                v5 -= (v5 | (v5 | (^(v5))))
                if (uint64(v6) <= 892) {
                    var v16 int16 = v3
                    _ = v16
                }
                _ = fi15
            }
        } else {
            v17 := GenS2{F0: 142}
            var v18 int16 = genFn0(100, (fi14 + (^(fi14))), 0)
            _ = v17
            _ = v18
        }
        v6 += 2
        _ = fi14
    }
    v5 += v5
    if ((32767 <= 255) || false) {
        var v19 int = genFn1(false, 92, (v13 >> ((255) & 63)))
        fmt.Println(v3, v13)
        var v20 int32 = v5
        _ = v19
        _ = v20
    } else {
        v6 = 1
    }
    v21 := v2["a"]
    switch 0 {
    case -91:
        v4 = v4
    case -641:
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 uint = v6
            _ = fi22
            _ = v23
        }
    case 127:
        v24 := GenS1{F0: v3, F1: v6, F2: 42}
        _ = v24
    default:
        fmt.Println(v4, v5)
    }
    v3 = v3
    v4 -= 1.5
    for fi25 := 0; fi25 < 6; fi25++ {
        v1["\n\t"] = v21
        v4 += v4
        _ = fi25
    }
    if true {
        if (float64(1e10) != (1.5 + 3.14159)) {
            if true {
                var v26 string = "\n\t"
                _ = v26
            }
        } else {
            switch genFn1(true, -128, v13) {
            case -1:
                fmt.Println(v21, v4)
            case 2147483647:
                var v27 int = 32768
                _ = v27
            case 32768:
                delete(v1, ("hello" + "foo bar"))
            default:
                fmt.Println(v3)
            }
        }
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 int8 = 100
            v30 := GenS2{F0: 255}
            v31 := v2[(("foo bar" + "") + "\n\t")]
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
    }
    v32 := GenS2{F0: 1}
    switch genFn1(false, 1, 1) {
    case 0:
        v4 = v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
