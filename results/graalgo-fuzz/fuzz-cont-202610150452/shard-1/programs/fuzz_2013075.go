package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int32 {
    var v7 float64 = (1e10 / (-2.25 - float64(float32(((float64(((((0.0 * 0.1) + ((1e10 * (float32(0.1) - 0.1)) * ((((1.5 + 0.0) / 1.0) / 1.0) + ((0.1 - ((0.1 - 0.0) * float32(-2.25))) * float32((float64(3.14159) * 0.1)))))) * ((3.14159 + 0.0) - float32(-2.25))) + 0.0)) * 1.5) - 1e10)))))
    v7 = v7
    v7 += 1e10
    v8 := GenS1{F0: 1}
    _ = v7
    _ = v8
    return 358
}

func main() {
    var v1 uint = 630
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    var v5 uint8 = 42
    var v6 int8 = -29
    v9 := GenS1{F0: v5}
    v10 := GenS1{F0: (v5 / (((v5 | 127)) | 1))}
    if (v6 >= v6) {
        v11 := GenS1{F0: v5}
        v1 -= (v1 + v1)
        for fi12 := 0; fi12 < 2; fi12++ {
            v5 = 1
            _ = fi12
        }
        _ = v11
    }
    fmt.Println(v5, v5, v1)
    v13 := []int{9223372036854775807, 0, 4294967295, 0}
    if ((v1 < v1) || true) {
        v14 := GenS1{F0: (^(v5))}
        _ = v14
    } else {
        var v15 uint16 = 42
        _ = v15
    }
    if ((229 > -1) && (32767 >= 42)) {
        if true {
            v1 -= 1
            if (!true) {
                switch 42 {
                case 252:
                    var v16 bool = false
                    _ = v16
                }
            } else {
                var v17 uint16 = 2
                for fi18 := 0; fi18 < 15; fi18++ {
                    v19 := v4["hello"]
                    v20 := GenS1{F0: v5}
                    var v21 uint16 = uint16(v5)
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                var v22 bool = false
                _ = v17
                _ = v22
            }
            switch ("a" + "foo bar") {
            case "hello":
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 uint = (v1 ^ 176)
                    _ = fi23
                    _ = v24
                }
            default:
                fmt.Println(v5)
            }
        }
        v4["\n\t"] = (!false)
        v6 -= v6
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 int16 = -661
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
