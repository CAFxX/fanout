package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int16 }
func genFn0(p0 int8, p1 int32) uint64 {
    v5 := []int{32768, 4294967295, -9223372036854775808}
    if true {
        var v6 int32 = p1
        var v7 float32 = (1e10 / 1.0)
        _ = v6
        _ = v7
    }
    if (true && ((!true) || (false && (!(2147483648 >= 2147483648))))) {
        switch -629 {
        case 4294967295:
            var v8 float32 = float32((((1.5 + 3.14159) / (1e10 + (((1.5 / 1.0) * float64(float32(-2.25))) - 1.5))) - float64(0.1)))
            _ = v8
        case 32767:
            var v9 int16 = 1
            _ = v9
        case -392:
            var v10 uint64 = 771
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 int32 = 889
    v2 := make(map[int]int)
    v3 := []uint32{42, 2, 4294967295, 4294967295}
    var v4 int = 412
    var v11 bool = true
    v12 := GenS1{F0: 9223372036854775807}
    fmt.Println(v11)
    for fi13 := 0; fi13 < 18; fi13++ {
        if (3.14159 <= float32(0.0)) {
            v11 = (v11 || (872 != 218))
            var v14 int16 = -530
            v11 = (!true)
            _ = v14
        }
        if ("foo bar" == ("\n\t" + ("hello" + ""))) {
            for fi15 := 0; fi15 < 20; fi15++ {
                v1 = (v1 ^ v1)
                fmt.Println(v4, fi13, fi15)
                v1 -= 0
                _ = fi15
            }
            v11 = (((v1 ^ v1) >= v1) && v11)
            var v16 bool = (v11 && (!(42 <= 32767)))
            _ = v16
        }
        v4 = (fi13 << ((65535) & 63))
        _ = fi13
    }
    v11 = (v11 && (127 != -127))
    if v11 {
        v4 = (v4 & (-(v4)))
        var v17 uint64 = 0
        _ = v17
    }
    fmt.Println(v4, v11, v4)
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            fmt.Println(fi18, v1, fi19)
            if ((!(9223372036854775807 < (fi19 << ((200) & 63)))) && v11) {
                v20 := GenS1{F0: 457}
                fmt.Println(fi19)
                _ = v20
            }
            switch v4 {
            case 4294967295:
                v2[32768] = (-(v4))
            default:
                v1 -= v1
                _ = len(v3)
            }
            _ = fi19
        }
        var v21 float32 = (3.14159 * float32(float64((-2.25 / 0.1))))
        v22 := GenS2{F0: (v4 - (v4 * (v4 - -1))), F1: 525}
        _ = fi18
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
