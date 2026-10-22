package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
func genFn0(p0 int32) string {
    fmt.Println(p0)
    v4 := []int{451, 414}
    for fi5 := 0; fi5 < 11; fi5++ {
        switch fi5 {
        case -1:
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 uint64 = 158
                fmt.Println(p0)
                p0 = (p0 * p0)
                fmt.Println(fi6, fi6, fi5)
                _ = fi6
                _ = v7
            }
        }
        var v8 float64 = -2.25
        var v9 uint = 42
        _ = fi5
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v4
    return "a"
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 42
    var v3 float32 = 1e10
    var v10 float64 = float64((v3 - (v3 * (float32(3.14159) * ((v3 - (v3 - v3)) * 0.0)))))
    v10 = (float64(float32(v10)) / v10)
    switch ("a" + ("\n\t" + "a")) {
    case "a":
        var v11 bool = (!(!(!true)))
        _ = v11
    case "foo bar":
        v12 := []int{32768}
        _ = v12
    case "\n\t":
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := GenS1{F0: 4294967295, F1: 18446744073709551615}
            v2 = 255
            var v15 uint32 = 4294967295
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v10 -= v10
    v10 -= v10
    fmt.Println(v2)
    v1[1] = true
    _ = len(v1)
    if (false || false) {
        v10 += v10
    }
    if (-32768 < 421) {
        v10 = 1e10
    }
    v3 += float32(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
