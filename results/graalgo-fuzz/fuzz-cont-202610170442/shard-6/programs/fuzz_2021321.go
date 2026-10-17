package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint64) string {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint32 = 617
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = (p0 / (((p0 + (p0 + 734))) | 1))
    _ = p0
    return ("foo bar" + "")
}
func genFn1(p0 uint64, p1 int, p2 uint8) int32 {
    var v10 uint = 468
    var v11 int = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 431
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 1
    var v3 uint16 = 255
    v4 := make(map[uint32]string)
    var v5 string = "foo bar"
    var v6 string = v5
    var v12 int8 = 1
    var v13 uint = 0
    fmt.Println(v6, v3, v6)
    v1[2147483647] = -9223372036854775808
    _ = len(v4)
    v6 = (v5 + v5)
    v14 := GenS1{F0: 2147483647}
    v2 += v2
    if (false || true) {
        if (-2.25 <= ((float32(3.14159) - (1.5 * 3.14159)) + 1.5)) {
            fmt.Println(v6)
        }
    } else {
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 12; fi16++ {
                v13 = 0
                var v17 bool = (!false)
                fmt.Println(v13)
                var v18 uint32 = 2
                var v19 int8 = v12
                _ = len(v1)
                var v20 uint16 = (v3 & v3)
                var v21 float32 = (1.5 + 0.0)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 float64 = 3.14159
            _ = fi15
            _ = v22
        }
    }
    var v23 uint64 = 2
    var v24 int32 = -226
    v25 := GenS1{F0: 124}
    fmt.Println(v23, v2)
    v24 = v24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
