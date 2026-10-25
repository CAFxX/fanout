package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int16 }
func genFn0(p0 string) int {
    var v7 float32 = (float32(3.14159) * 0.0)
    var v8 int = 32768
    p0 = (p0 + (("hello" + "\n\t") + p0))
    p0 = "foo bar"
    _ = p0
    _ = v7
    _ = v8
    return -755
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"", "\n\t", "foo bar", "\n\t"}
    var v3 float64 = 0.1
    var v4 int8 = 1
    v5 := []string{"a", ""}
    v6 := make(map[string]uint32)
    var v9 int = 0
    if false {
        switch 2147483647 {
        case 32767:
            switch "" {
            case "":
                v9 -= (v9 << ((704) & 63))
            case "hello":
                v10 := GenS1{F0: v9, F1: 65535, F2: int16(v9)}
                _ = v10
            default:
                v11 := GenS1{F0: (v9 / (((^(v9))) | 1)), F1: 0, F2: -1}
                _ = v11
            }
        default:
            v9 += genFn0("hello")
        }
    } else {
        var v12 uint64 = 4294967295
        if ((-(v9)) > v9) {
            v13 := v6["\n\t"]
            fmt.Println(v4)
            _ = v13
        }
        v14 := GenS1{F0: (-(v9)), F1: 150, F2: 32767}
        _ = v12
        _ = v14
    }
    switch "a" {
    case "":
        var v15 int16 = 439
        _ = v15
    case "foo bar":
        for fi16 := 0; fi16 < 19; fi16++ {
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 float32 = float32(0.0)
                if (411 > 1) {
                    _ = len(v6)
                    fmt.Println(v4, v9, v4)
                }
                _ = fi17
                _ = v18
            }
            v4 = v4
            v9 += genFn0("")
            _ = fi16
        }
    default:
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 float32 = 0.1
            v20 = (v20 * 3.14159)
            v4 += int8(v9)
            _ = fi19
            _ = v20
        }
    }
    v4 = 42
    v21 := GenS1{F0: v9, F1: 1, F2: -1}
    v9 -= v9
    v22 := GenS1{F0: genFn0(("hello" + "hello")), F1: 254, F2: -1}
    var v23 float64 = float64(1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
