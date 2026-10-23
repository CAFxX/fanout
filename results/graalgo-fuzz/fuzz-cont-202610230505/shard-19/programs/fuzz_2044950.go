package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint64 }

func main() {
    var v1 int = -453
    var v2 uint16 = 32767
    var v3 int32 = 42
    v4 := []uint32{42, 736, 216, 4294967295}
    var v5 string = "foo bar"
    v5 = ((v5 + "foo bar") + v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        v3 += v3
    }
    v3 -= (v3 | (v3 - (v3 >> ((315) & 63))))
    v4 = append(v4, 2147483648)
    var v6 int16 = -519
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint64 = 18446744073709551615
        if true {
            v9 := GenS1{F0: 4294967295}
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 uint = 751
                v5 = v5
                if (0 <= 862) {
                    var v12 float64 = (0.1 * 3.14159)
                    _ = v12
                } else {
                    var v13 bool = true
                    fmt.Println(v13, v8, v6)
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
            v3 = (-(v3))
            _ = v9
        } else {
            v4 = append(v4, 42)
        }
        v14 := GenS2{F0: true}
        _ = fi7
        _ = v14
        _ = v8
    }
    if false {
        v15 := GenS1{F0: 2147483647}
        var v16 float32 = (1.5 / 3.14159)
        if false {
            var v17 int = v1
            _ = v17
        }
        _ = v15
        _ = v16
    }
    if (v2 <= (v2 & v2)) {
        v5 = ((v5 + "") + "a")
        var v18 int32 = v3
        _ = v18
    } else {
        var v19 int8 = -67
        _ = v19
    }
    v20 := GenS3{F0: 0, F1: 42, F2: 298}
    if len(v4) > 0 {
        v4[0] = 108
    }
    v1 = -868
    v6 -= v6
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v21 := GenS3{F0: 4294967295, F1: 42, F2: 2}
    switch "foo bar" {
    case "foo bar":
        v22 := GenS1{F0: 824}
        _ = v22
    case "\n\t":
        if ((127 == 249) && false) {
            v6 -= v6
            var v23 uint8 = 127
            for fi24 := 0; fi24 < 6; fi24++ {
                v3 -= 2147483647
                _ = fi24
            }
            _ = v23
        } else {
            v6 = (v6 | (v6 ^ -259))
        }
    case "":
        if (!(-125 <= 42)) {
            fmt.Println(v3, v1, v1)
            fmt.Println(v3, v2, v6)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
