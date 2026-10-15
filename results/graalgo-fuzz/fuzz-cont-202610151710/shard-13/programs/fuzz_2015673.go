package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int8 }
func genFn0(p0 int64) int {
    v7 := []int{127, -2147483648, -2147483648}
    fmt.Println(p0)
    var v8 int8 = 0
    _ = p0
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1() uint32 {
    switch 32767 {
    case -2147483648:
        var v9 uint64 = 4294967295
        _ = v9
    case 4294967295:
        var v10 uint8 = 255
        _ = v10
    case 128:
        var v11 string = "\n\t"
        _ = v11
    default:
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := []int{genFn0(-87), (fi12 ^ genFn0(-9223372036854775808))}
            var v14 int64 = 243
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    var v15 float64 = ((3.14159 - float64(0.0)) - ((float64((-2.25 * 0.0)) + (0.0 * 0.0)) + 3.14159))
    v15 -= (0.1 + v15)
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 int16 = -373
        _ = fi16
        _ = v17
    }
    fmt.Println(v15)
    _ = v15
    return 2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 1
    var v3 string = "\n\t"
    var v4 float64 = float64((0.1 - (-2.25 - 0.1)))
    var v5 uint = 2
    var v6 int64 = 1
    for fi18 := 0; fi18 < 10; fi18++ {
        for fi19 := 0; fi19 < 11; fi19++ {
            v20 := GenS1{F0: (v2 | v2), F1: 498, F2: 42}
            v21 := []int{fi18, int(v6), fi18, (^(fi19))}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    if true {
        var v22 uint8 = 146
        var v23 string = v3
        if (255 > 255) {
            v1[(("foo bar" + (v3 + v23)) + v3)] = 65535
            v24 := v1[(v3 + v3)]
            _ = v24
        } else {
            var v25 int64 = 226
            v26 := GenS2{F0: 58, F1: 255, F2: 127}
            var v27 uint32 = 65535
            _ = v25
            _ = v26
            _ = v27
        }
        _ = v22
        _ = v23
    }
    if (((v5 < v5) || true) && false) {
        v1["foo bar"] = 65535
        var v28 int32 = 1
        switch 128 {
        case -9223372036854775808:
            for fi29 := 0; fi29 < 11; fi29++ {
                v30 := []int{}
                _ = fi29
                _ = v30
            }
        case 364:
            v31 := GenS1{F0: v2, F1: 4294967295, F2: 32767}
            _ = v31
        default:
            if false {
                v28 = (v28 << ((v5) & 63))
            } else {
                if (1000 <= 42) {
                    var v32 int16 = 775
                    fmt.Println(v32)
                    v28 -= v28
                    _ = v32
                }
                v33 := GenS2{F0: 98, F1: (^(v2)), F2: 100}
                _ = v33
            }
        }
        _ = v28
    }
    v6 -= 1
    _ = len(v1)
    _ = len(v1)
    delete(v1, ((((((v3 + v3) + v3) + (v3 + "")) + ("" + "")) + "\n\t") + v3))
    v5 = 18446744073709551615
    v34 := v1[("hello" + v3)]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
