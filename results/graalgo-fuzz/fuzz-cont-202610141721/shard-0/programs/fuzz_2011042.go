package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 bool) int8 {
    var v5 int32 = 42
    p0 = ((-173 == -686) && (!p0))
    _ = p0
    _ = v5
    return -28
}
func genFn1() uint8 {
    var v6 float64 = (3.14159 * (((1.5 - (((0.0 / 0.1) / 1.5) - 0.0)) - (1.5 - 3.14159)) / (((float64(float32((0.1 * ((float64(0.1) * 0.1) - 0.1)))) * float64(float32(1e10))) / -2.25) - 0.0)))
    switch 1 {
    case 42:
        if (42 <= 1) {
            v6 += v6
        }
    case 128:
        v6 += (v6 / ((0.0 - v6) * (v6 * v6)))
    case 32767:
        v7 := GenS1{F0: 1, F1: 78}
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: 100}
    _ = v6
    _ = v8
    return 128
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 4294967295
    v3 := make(map[string]int)
    var v4 uint16 = 65535
    v9 := v3["hello"]
    v10 := []int{v9, 9223372036854775807}
    if ((("\n\t" + ((("" + "a") + ("hello" + ("a" + "hello"))) + "")) >= "") || true) {
        var v11 float32 = 0.1
        v9 = 406
        _ = v11
    }
    if (true || (-1 == 425)) {
        if false {
            if len(v10) > 0 {
                v10[0] = (v9 | int(v2))
            }
            switch "a" {
            case "foo bar":
                v12 := GenS2{F0: 1}
                _ = v12
            default:
                var v13 int8 = int8(v2)
                _ = v13
            }
            var v14 float64 = 1.5
            _ = v14
        } else {
            var v15 float32 = float32(-2.25)
            fmt.Println(v2)
            _ = v15
        }
    }
    fmt.Println(v9, v9, v2)
    v2 = 933
    for fi16 := 0; fi16 < 19; fi16++ {
        switch 127 {
        case 128:
            switch (("foo bar" + "") + ("hello" + "")) {
            case "foo bar":
                var v17 int64 = 9223372036854775807
                _ = v17
            case "\n\t":
                for fi18 := 0; fi18 < 9; fi18++ {
                    var v19 int16 = 42
                    if (((1.5 + (3.14159 + (-2.25 + (float32(1.5) * (((1.5 / float32(1e10)) * 3.14159) + 1.5))))) <= ((1.5 + (((-2.25 / 1.5) / 1.5) / 3.14159)) / (float32((((1e10 - 0.0) * 1.5) + ((1.5 + (1e10 / 1.0)) - -2.25))) + (0.1 / 1.5)))) || (false && ((!(!true)) || true))) {
                        if len(v10) > 0 {
                            v10[0] = fi16
                        }
                        fmt.Println(v4)
                        v20 := v3["hello"]
                        _ = v20
                    }
                    v21 := GenS3{F0: 1}
                    v22 := GenS3{F0: -296}
                    var v23 bool = (813 > 513)
                    _ = fi18
                    _ = v19
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
        case 726:
            fmt.Println(v9, v4, fi16)
        default:
            v2 -= v2
        }
        _ = fi16
    }
    var v24 uint16 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
