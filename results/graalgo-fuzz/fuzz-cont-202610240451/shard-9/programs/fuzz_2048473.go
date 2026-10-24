package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 float64 = (1.5 - 3.14159)
    var v2 float32 = 0.1
    v3 := []uint32{2}
    var v4 int = 906
    v1 -= (0.0 / ((float64(v2) + v1) + 1.5))
    if (4294967295 < 854) {
        v5 := GenS1{F0: 119}
        _ = v5
    } else {
        var v6 uint64 = 4294967295
        v4 -= (v4 * (v4 | 391))
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int = (v4 / ((v4) | 1))
        fmt.Println(fi7)
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(v1, fi9, fi7)
            var v10 uint16 = 32767
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v2)
    if len(v3) > 0 {
        v3[0] = 1
    }
    switch v4 {
    case 4294967295:
        if true {
            v1 += v1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            fmt.Println(v4, v1)
        }
    case -2147483648:
        var v11 float32 = v2
        _ = v11
    case -1:
        _ = len(v3)
    default:
        v2 = v2
    }
    if false {
        _ = len(v3)
        v1 -= v1
        for fi12 := 0; fi12 < 12; fi12++ {
            if false {
                v13 := GenS1{F0: 255}
                fmt.Println(fi12, v1, v4)
                v14 := GenS1{F0: 127}
                _ = v13
                _ = v14
            }
            v3 = append(v3, 42)
            _ = fi12
        }
    }
    v2 += float32(v1)
    v15 := GenS2{F0: 251}
    var v16 bool = (true && false)
    fmt.Println(v4, v4, v16)
    if (42 <= 299) {
        if v16 {
            switch ("" + "hello") {
            case "\n\t":
                switch 128 {
                case 32767:
                    fmt.Println(v1, v16, v16)
                default:
                    var v17 int64 = -1
                    _ = v17
                }
            }
        }
        for fi18 := 0; fi18 < 7; fi18++ {
            v4 -= (fi18 & (fi18 ^ (fi18 & -244)))
            var v19 bool = (1 > -1)
            for fi20 := 0; fi20 < 5; fi20++ {
                if v16 {
                    v19 = ((-1 != 2147483647) && v16)
                } else {
                    v1 = v1
                    v21 := GenS3{F0: 2147483647}
                    v22 := GenS3{F0: 4294967295}
                    _ = v21
                    _ = v22
                }
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
    }
    switch "a" {
    case "foo bar":
        if (682 == 18446744073709551615) {
            var v23 int = -1
            fmt.Println(v1)
            _ = v23
        } else {
            var v24 uint = 42
            _ = v24
        }
    case "hello":
        fmt.Println(v1, v1, v4)
    case "a":
        var v25 int = v4
        _ = v25
    default:
        v26 := GenS1{F0: 128}
        _ = v26
    }
    var v27 float32 = 3.14159
    switch (("\n\t" + ("" + "a")) + (("a" + "hello") + "foo bar")) {
    case "foo bar":
        if false {
            var v28 uint = 0
            _ = v28
        }
    case "\n\t":
        var v29 uint8 = 128
        _ = v29
    case "":
        if (127 <= 147) {
            var v30 uint8 = 42
            _ = v30
        }
    default:
        v31 := GenS2{F0: 42}
        _ = v31
    }
    fmt.Println(v2, v2)
    v1 -= 3.14159
    var v32 int = (v4 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
}
