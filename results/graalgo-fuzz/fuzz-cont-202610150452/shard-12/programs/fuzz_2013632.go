package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint32) int16 {
    var v4 int = -1
    p1 -= 2147483648
    for fi5 := 0; fi5 < 17; fi5++ {
        p0 = (!(true || p0))
        p0 = (4294967295 <= 42)
        p1 = 1
        _ = fi5
    }
    v6 := GenS1{F0: 100, F1: 1, F2: 4294967295}
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 663
}

func main() {
    v1 := []string{"hello", ("" + ("\n\t" + (("" + "a") + "\n\t")))}
    v2 := make(map[uint32]string)
    var v3 int16 = 32767
    v3 += (v3 & v3)
    if true {
        v3 = (v3 * 221)
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v3 += (v3 - v3)
        if true {
            v8 := GenS1{F0: 1, F1: 32767, F2: 409}
            _ = v8
        }
        _ = fi7
    }
    v3 = -1
    for fi9 := 0; fi9 < 14; fi9++ {
        v3 = (v3 >> ((255) & 63))
        var v10 int16 = 124
        _ = fi9
        _ = v10
    }
    switch 0 {
    case 32768:
        v3 -= (v3 * v3)
    default:
        if (-712 < -2147483648) {
            v11 := GenS1{F0: -1, F1: 897, F2: 2}
            _ = v11
        } else {
            v12 := GenS2{F0: 32767}
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := GenS1{F0: 0, F1: 564, F2: 124}
                var v15 bool = (-339 != genFn0(true, 2147483648))
                var v16 int32 = 42
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            v3 -= 1
            _ = v12
        }
    }
    v3 -= (v3 / ((v3) | 1))
    var v17 string = ("hello" + "a")
    if (((float32(((1e10 / (-2.25 * 1.5)) + 1e10)) * (float32(((-2.25 / 3.14159) * 3.14159)) - ((float32(0.0) * 3.14159) + (0.1 * 3.14159)))) / (float32(1.5) * 1.5)) < 1e10) {
        var v18 float32 = 1e10
        _ = v18
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
}
