package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
func genFn0(p0 string) uint32 {
    var v7 float64 = float64(0.0)
    v8 := []int{0, 0}
    _ = p0
    _ = v7
    _ = v8
    return 729
}

func main() {
    var v1 uint16 = 32767
    v2 := []string{"a", ("foo bar" + "foo bar"), (("a" + ("hello" + "foo bar")) + "foo bar"), ("foo bar" + "a")}
    v3 := make(map[string]string)
    var v4 uint = 42
    var v5 uint64 = 738
    v6 := make(map[int]uint32)
    v1 = (^(v1))
    for fi9 := 0; fi9 < 16; fi9++ {
        switch "hello" {
        case "":
            fmt.Println(v1)
        default:
            fmt.Println(fi9, fi9)
        }
        _ = fi9
    }
    _ = len(v6)
    v4 += (^(v4))
    var v10 int = 128
    v11 := GenS1{F0: -1, F1: 32767}
    if (4294967295 <= 42) {
        var v12 uint32 = 216
        fmt.Println(v10)
        _ = v12
    }
    v2 = append(v2, "a")
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            fmt.Println(v4, v1, fi14)
            v5 += (v5 << ((v4) & 63))
            v6[(v10 >> (((^(v4))) & 63))] = 2147483648
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            if len(v2) > 0 {
                v2[0] = ("" + "foo bar")
            }
            _ = fi14
        }
        var v15 int16 = 492
        _ = fi13
        _ = v15
    }
    fmt.Println(v1, v10, v10)
    v16 := GenS1{F0: 1, F1: -772}
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := v6[(fi17 ^ 127)]
        var v19 uint16 = v1
        v1 = (v19 << ((v4) & 63))
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: -9223372036854775808, F1: (v10 & 128)}
    v4 += uint(v1)
    for fi21 := 0; fi21 < 2; fi21++ {
        if (!(false && (77 == 128))) {
            if (((v1 * v1) < v1) || (!false)) {
                v5 = (v5 | 42)
                v4 = v4
                var v22 int16 = -32768
                _ = v22
            }
            for fi23 := 0; fi23 < 17; fi23++ {
                v4 = v4
                _ = fi23
            }
            if len(v2) > 0 {
                v2[0] = ("\n\t" + "a")
            }
        }
        _ = fi21
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
