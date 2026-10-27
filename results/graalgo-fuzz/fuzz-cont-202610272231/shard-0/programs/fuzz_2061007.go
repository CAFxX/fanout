package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16, p1 int) int64 {
    switch (-(p1)) {
    case -583:
        switch "" {
        case "a":
            v5 := GenS1{F0: 42}
            _ = v5
        case "\n\t":
            p0 = int16(p1)
        case "foo bar":
            p1 += p1
        }
    case 128:
        var v6 bool = ((-128 < 42) && false)
        _ = v6
    case 9223372036854775807:
        v7 := []int{p1, 446, 1}
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        p1 = (fi8 << ((529) & 63))
        _ = fi8
    }
    _ = p0
    _ = p1
    return int64(p0)
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]int)
    var v3 float32 = 0.1
    v4 := []int{-9223372036854775808}
    if (!false) {
        var v9 int = -730
        var v10 uint16 = 42
        fmt.Println(v3, v9, v3)
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 127}
    v12 := GenS1{F0: -1}
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    if (2 > 0) {
        for fi13 := 0; fi13 < 8; fi13++ {
            fmt.Println(v1)
            _ = fi13
        }
        fmt.Println(v3, v1)
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        v1 = (((((((((v1 / float64(v3)) / v1) - (float64(0.1) - v1)) / v1) / (v1 - (v1 * 1.5))) - (v1 - (3.14159 / v1))) * (v1 + v1)) - 0.0) / float64(v3))
        fmt.Println(v3)
        _ = fi14
    }
    var v15 uint8 = 255
    if false {
        fmt.Println(v3, v3, v1)
        v3 += v3
    }
    var v16 float32 = v3
    v16 = (float32((v1 - 1e10)) + (((-2.25 * (v16 + (1.5 - (((3.14159 - (v3 / 1e10)) - float32(1.5)) * v3)))) * (v3 + (1.5 - (float32(v1) - (v3 + ((-2.25 - 1e10) * ((0.0 / v3) / v16))))))) * 1e10))
    v2["hello"] = 2147483647
    if true {
        var v17 float64 = 0.0
        _ = v17
    }
    var v18 int16 = int16(v15)
    v3 += v16
    if true {
        switch (("foo bar" + (("a" + ("foo bar" + "a")) + ("" + (("\n\t" + ("\n\t" + ("\n\t" + "foo bar"))) + "hello")))) + (("a" + "\n\t") + (("foo bar" + "foo bar") + "\n\t"))) {
        case "a":
            v19 := v2["foo bar"]
            _ = v19
        default:
            for fi20 := 0; fi20 < 3; fi20++ {
                v1 = ((v1 / 1.0) / float64(v3))
                fmt.Println(v1, v16)
                _ = fi20
            }
        }
    } else {
        for fi21 := 0; fi21 < 6; fi21++ {
            v3 = v3
            v22 := GenS1{F0: 0}
            for fi23 := 0; fi23 < 12; fi23++ {
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
                fmt.Println(v15, v18)
                var v24 uint = 65535
                _ = fi23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            if (0.1 < (v3 / v16)) {
                v1 += float64(v3)
                var v26 float64 = (v1 / -2.25)
                _ = v26
            } else {
                v27 := GenS1{F0: 80}
                v28 := GenS1{F0: 1}
                _ = v27
                _ = v28
            }
            v16 += ((v16 / v3) - (0.0 * v3))
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 float64 = ((v1 + 0.0) + float64(0.1))
                var v31 uint8 = v15
                v15 += v15
                var v32 uint64 = 1
                var v33 bool = true
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
                _ = v33
            }
            _ = fi25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
