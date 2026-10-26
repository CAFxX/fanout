package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int8; F1 string; F2 uint }
func genFn0(p0 uint32, p1 int64) uint64 {
    if (2 != 874) {
        v7 := GenS1{F0: "hello", F1: 32767}
        _ = v7
    }
    if (0.0 == (float32(3.14159) - (-2.25 - 1.5))) {
        switch ((("" + "") + "a") + (("a" + "") + "\n\t")) {
        case "foo bar":
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := []int{}
                _ = fi8
                _ = v9
            }
        case "a":
            var v10 int32 = -74
            _ = v10
        case "":
            v11 := GenS1{F0: "a", F1: 752}
            _ = v11
        default:
            if false {
                var v12 int8 = -32
                var v13 int32 = 22
                _ = v12
                _ = v13
            } else {
                p0 += (p0 + (p0 - 666))
            }
        }
        var v14 bool = false
        _ = v14
    }
    if (-1 != 1) {
        for fi15 := 0; fi15 < 6; fi15++ {
            p1 += (p1 + 42)
            _ = fi15
        }
    }
    p1 = (-(p1))
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint64 = 171
    var v2 int32 = -2147483648
    var v3 string = "\n\t"
    v4 := []uint32{492}
    v5 := make(map[string]string)
    var v6 uint16 = 32767
    fmt.Println(v6, v2, v6)
    if (!(!true)) {
        var v16 int = -528
        _ = v16
    }
    v3 = "\n\t"
    switch 127 {
    case 32768:
        var v17 float64 = ((float64(1.5) + 3.14159) - (-2.25 / 1.0))
        _ = v17
    case -875:
        switch "" {
        case "hello":
            var v18 float64 = (0.0 / 1.0)
            _ = v18
        case "foo bar":
            v6 = v6
        case "":
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS1{F0: ("foo bar" + v3), F1: 776}
                switch (((v3 + v3) + (v3 + v3)) + v3) {
                case "\n\t":
                    fmt.Println(v3, v1)
                }
                _ = fi19
                _ = v20
            }
        default:
            v21 := GenS2{F0: -1, F1: v3, F2: 1}
            _ = v21
        }
    }
    if false {
        var v22 float32 = 1.5
        _ = v22
    }
    _ = len(v4)
    v23 := v5[v3]
    var v24 int32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
