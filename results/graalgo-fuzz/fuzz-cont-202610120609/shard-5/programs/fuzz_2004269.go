package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 bool, p1 int) int32 {
    switch ("a" + ((("foo bar" + "") + "hello") + ("a" + "a"))) {
    case "hello":
        fmt.Println(p0)
    case "a":
        for fi7 := 0; fi7 < 2; fi7++ {
            p1 -= p1
            p0 = p0
            fmt.Println(p1)
            _ = fi7
        }
    case "\n\t":
        switch "a" {
        case "a":
            if p0 {
                fmt.Println(p0, p1)
            }
        default:
            var v8 uint16 = 0
            _ = v8
        }
    default:
        v9 := GenS1{F0: "\n\t"}
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        p0 = (262 <= 42)
        var v11 int32 = 967
        _ = fi10
        _ = v11
    }
    var v12 int = p1
    p0 = false
    if false {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 int16 = 1000
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 float32 = (0.0 - (-2.25 * -2.25))
            if ((!(100 != -1)) || true) {
                p0 = p0
                v17 := []int{v12, (^(fi15))}
                v16 = 0.1
                _ = v17
            }
            _ = fi15
            _ = v16
        }
    } else {
        var v18 uint32 = 2147483647
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v12
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{}
    var v3 uint8 = 67
    v4 := make(map[uint32]uint32)
    var v5 float64 = (1.5 + -2.25)
    var v6 string = "\n\t"
    if (false || (false || false)) {
        v3 += (^(v3))
        switch 500 {
        case 42:
            v19 := GenS1{F0: "a"}
            _ = v19
        default:
            var v20 uint16 = 32767
            _ = v20
        }
        if (42 <= 9223372036854775807) {
            v21 := GenS1{F0: v6}
            _ = v21
        }
    } else {
        var v22 string = "hello"
        var v23 float64 = ((1.5 / v5) * v5)
        _ = len(v2)
        _ = v22
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = 289
    }
    v5 -= float64(0.0)
    v3 += v3
    v6 = v6
    var v24 uint = 323
    if false {
        var v25 int16 = -573
        v5 = (float64(float32(v5)) * 1e10)
        _ = v25
    }
    if (v3 != v3) {
        switch -858 {
        case -2147483648:
            var v26 int32 = 0
            _ = v26
        }
    }
    switch (v6 + v6) {
    case "\n\t":
        v2 = append(v2, 4294967295)
    }
    var v27 uint16 = 65535
    delete(v1, 65535)
    v28 := GenS1{F0: v6}
    var v29 float32 = 0.1
    v30 := GenS1{F0: (v6 + "\n\t")}
    v29 = v29
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
