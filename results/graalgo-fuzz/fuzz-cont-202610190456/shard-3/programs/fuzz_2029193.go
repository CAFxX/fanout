package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int16; F2 int64 }
func genFn0(p0 int32, p1 uint8) int64 {
    v7 := []int{0, 2147483647}
    v8 := GenS1{F0: 1}
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(fi10, p1, fi10)
            _ = fi10
        }
        _ = fi9
    }
    var v11 bool = false
    p0 += -2147483648
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 uint8 = 0
    var v5 bool = (true || (-2.25 < ((0.0 / (3.14159 - 0.1)) + float64(float32(float64(3.14159))))))
    var v6 int16 = 32767
    v4 -= (v4 + (v4 & v4))
    delete(v2, 0)
    v12 := GenS2{F0: 815, F1: (v6 | v6), F2: v1}
    v5 = (!(v5 && ((0.1 <= 3.14159) || (!true))))
    v5 = v5
    fmt.Println(v5, v6, v1)
    v13 := GenS1{F0: v1}
    if ((!v5) || (v5 || (!(3.14159 > 0.0)))) {
        var v14 int64 = v1
        v15 := v3[("foo bar" + "hello")]
        _ = v14
        _ = v15
    } else {
        var v16 int8 = 127
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        delete(v2, (fi17 ^ fi17))
        _ = fi17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 int = fi18
        if (!(!v5)) {
            v4 -= (v4 << ((970) & 63))
            var v20 uint64 = 18446744073709551615
            var v21 uint64 = v20
            _ = v20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 uint16 = 73
        if (v5 && v5) {
            switch fi22 {
            case 0:
                v6 -= v6
            default:
                for fi24 := 0; fi24 < 5; fi24++ {
                    v25 := GenS2{F0: 1, F1: (v6 + (v6 ^ (v6 / (((v6 - (v6 & (v6 + int16(v23))))) | 1)))), F2: 322}
                    var v26 int16 = v6
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            v27 := []int{}
            _ = v27
        }
        _ = fi22
        _ = v23
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
