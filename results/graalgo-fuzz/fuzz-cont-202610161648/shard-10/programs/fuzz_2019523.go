package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int8 {
    var v7 uint64 = 18446744073709551615
    v7 = v7
    _ = v7
    return -128
}

func main() {
    var v1 string = "a"
    v2 := []int{42, -2147483648, -774, -137}
    var v3 uint = 2
    var v4 int = 127
    v5 := make(map[string]uint32)
    var v6 uint64 = 2
    if true {
        var v8 float32 = (0.1 / -2.25)
        var v9 int16 = -1
        var v10 uint16 = 65535
        _ = v10
        _ = v8
        _ = v9
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v6 = (v6 - 4294967295)
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int32 = -1
        v14 := GenS1{F0: v3}
        _ = fi12
        _ = v13
        _ = v14
    }
    v6 -= v6
    var v15 int16 = 0
    delete(v5, ("\n\t" + (("\n\t" + "foo bar") + ("\n\t" + "foo bar"))))
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int16 = (v15 & (v15 << (((^(v3))) & 63)))
            fmt.Println(v1, v3)
            var v19 bool = (444 < 42)
            if ((fi17 % (((v4 - fi17)) | 1)) <= fi16) {
                var v20 float32 = (-2.25 * float32(((1.5 - float64(((((((float32(-2.25) / -2.25) / (1e10 / 1.5)) * 1e10) - 1e10) / 1e10) + 3.14159))) / -2.25)))
                _ = v20
            } else {
                if (498 == v3) {
                    var v21 uint32 = 1
                    _ = v21
                }
                v2 = append(v2, (^(v4)))
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        switch ((v1 + v1) + "") {
        case "a":
            v5[(v1 + v1)] = 4294967295
        default:
            fmt.Println(v1)
        }
        fmt.Println(fi16, fi16)
        _ = fi16
    }
    v22 := GenS1{F0: (v3 * 1)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
