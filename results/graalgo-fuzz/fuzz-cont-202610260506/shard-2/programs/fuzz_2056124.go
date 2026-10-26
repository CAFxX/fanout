package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint32) int8 {
    for fi5 := 0; fi5 < 17; fi5++ {
        p0 -= p0
        var v6 float64 = (3.14159 * 0.1)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        if ("foo bar" < ((("a" + "a") + "\n\t") + (("" + (("foo bar" + "hello") + "\n\t")) + "foo bar"))) {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 int64 = -1
                _ = fi8
                _ = v9
            }
        } else {
            var v10 int32 = -218
            var v11 int32 = v10
            v12 := GenS1{F0: 1, F1: 1}
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    p0 -= (p0 | (p0 ^ 2147483648))
    if true {
        var v13 int = 2147483647
        _ = v13
    }
    _ = p0
    return -17
}

func main() {
    v1 := []string{}
    var v2 uint32 = 0
    var v3 uint = 600
    v4 := []int{}
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v3, v3)
        _ = fi14
    }
    var v15 int = 222
    var v16 uint16 = 32767
    v2 = v2
    fmt.Println(v15)
    if ((v16 / (((^(v16))) | 1)) == (v16 ^ (v16 >> ((v3) & 63)))) {
        if ("hello" > (("\n\t" + "a") + "foo bar")) {
            for fi17 := 0; fi17 < 19; fi17++ {
                var v18 int16 = -32768
                _ = fi17
                _ = v18
            }
        }
        v4 = append(v4, v15)
    } else {
        for fi19 := 0; fi19 < 3; fi19++ {
            switch v15 {
            case 2147483647:
                var v20 int32 = -2147483648
                _ = v20
            }
            v21 := GenS1{F0: 0, F1: 42}
            v22 := GenS1{F0: 255, F1: -9223372036854775808}
            _ = fi19
            _ = v21
            _ = v22
        }
        var v23 uint = (^(v3))
        _ = v23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        v15 -= fi24
        _ = len(v4)
        _ = fi24
    }
    v25 := GenS1{F0: 1, F1: 9223372036854775807}
    fmt.Println(v3, v3)
    v3 -= v3
    var v26 uint = (v3 * v3)
    v3 += v3
    var v27 uint32 = v2
    v1 = append(v1, "")
    var v28 int32 = 2147483647
    if true {
        var v29 uint16 = 514
        if ((1e10 / 1.0) < 0.0) {
            v1 = append(v1, "a")
        }
        _ = v29
    } else {
        var v30 int = v15
        for fi31 := 0; fi31 < 10; fi31++ {
            var v32 uint8 = 2
            _ = fi31
            _ = v32
        }
        _ = v30
    }
    fmt.Println(v2, v28)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
