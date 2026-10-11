package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int16) string {
    for fi7 := 0; fi7 < 16; fi7++ {
        switch ("hello" + ("a" + ("a" + "a"))) {
        case "":
            switch fi7 {
            case 42:
                var v8 float64 = -2.25
                _ = v8
            }
        case "foo bar":
            if (0.1 > 0.0) {
                fmt.Println(fi7, p1)
            } else {
                p0 += (p0 << ((42) & 63))
                fmt.Println(p1)
                v9 := GenS1{F0: 0, F1: 206, F2: int16(p0)}
                _ = v9
            }
        default:
            p1 = p1
        }
        v10 := GenS1{F0: 4294967295, F1: 2147483648, F2: (p1 ^ p1)}
        var v11 int64 = 368
        _ = fi7
        _ = v10
        _ = v11
    }
    p0 -= 2
    var v12 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = v12
    return "a"
}

func main() {
    var v1 bool = (true || true)
    var v2 uint32 = 411
    v3 := []string{"\n\t", "a", "hello", (("hello" + "hello") + "a")}
    var v4 uint = 0
    var v5 int16 = -32768
    var v6 uint = 2
    var v13 int16 = (v5 ^ (v5 | -896))
    v3 = append(v3, ("a" + "\n\t"))
    var v14 int64 = 9223372036854775807
    var v15 float32 = (0.0 * 1.5)
    v15 = v15
    v5 = (v5 + -733)
    if false {
        var v16 uint = (v4 << ((18446744073709551615) & 63))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v16
    } else {
        var v17 float32 = 0.1
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(v2, v2)
            v3 = append(v3, "a")
            _ = fi18
        }
        _ = v17
    }
    var v19 int32 = -2147483648
    for fi20 := 0; fi20 < 12; fi20++ {
        fmt.Println(v5)
        if (3.14159 >= 1e10) {
            switch (("" + "a") + "hello") {
            case "hello":
                v21 := GenS1{F0: v2, F1: 0, F2: 0}
                _ = v21
            default:
                if (("" + "a") == "hello") {
                    if v1 {
                        v3 = append(v3, "hello")
                        fmt.Println(v14, v15, v13)
                        var v22 uint8 = 0
                        _ = v22
                    }
                    v13 = v5
                    var v23 uint8 = 2
                    _ = v23
                }
            }
            v2 = v2
        }
        v24 := GenS2{F0: 1}
        _ = fi20
        _ = v24
    }
    _ = len(v3)
    v25 := GenS1{F0: (v2 * v2), F1: (v2 - v2), F2: v5}
    var v26 int32 = -2147483648
    var v27 uint = v4
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
