package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0(p0 int16, p1 int8, p2 int64) string {
    p1 -= p1
    v5 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ""
}

func main() {
    var v1 float32 = 3.14159
    var v2 string = "foo bar"
    var v3 int64 = 923
    var v4 int8 = -1
    var v6 int16 = 32767
    if true {
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(v4, v3)
            var v8 uint = 42
            var v9 float64 = -2.25
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v1 = (float32(((-2.25 / 1.0) + 0.1)) / (0.1 * 3.14159))
    v6 -= 0
    v3 += int64(v6)
    if (-9223372036854775808 != 343) {
        v1 += v1
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 int8 = v4
            var v12 int64 = v3
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := GenS2{F0: 32767, F1: 0, F2: -1}
        _ = v13
    } else {
        fmt.Println(v4)
        var v14 uint64 = 1
        switch "" {
        case "hello":
            var v15 bool = false
            _ = v15
        default:
            v16 := []int{-9223372036854775808, 9223372036854775807}
            _ = v16
        }
        _ = v14
    }
    var v17 int64 = v3
    fmt.Println(v4, v17, v2)
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := []int{(fi18 + fi18), -1}
        var v20 uint16 = 1
        _ = fi18
        _ = v19
        _ = v20
    }
    switch 0 {
    case 978:
        if (65535 >= 103) {
            fmt.Println(v3, v2)
        } else {
            switch 4294967295 {
            case 4294967295:
                var v21 int16 = (v6 + 221)
                _ = v21
            case -462:
                var v22 uint32 = 4294967295
                _ = v22
            case 536:
                v23 := GenS1{F0: false}
                _ = v23
            }
        }
    case -744:
        v24 := GenS1{F0: false}
        _ = v24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        fmt.Println(fi25, fi25)
        if false {
            v2 = "\n\t"
            var v26 string = v2
            v27 := GenS3{F0: (v2 + "foo bar"), F1: 42}
            _ = v26
            _ = v27
        } else {
            v28 := []int{42, fi25, fi25, fi25}
            _ = v28
        }
        if (!false) {
            v29 := []int{(fi25 & (fi25 & 438)), (fi25 % (((fi25 * -2147483648)) | 1))}
            _ = v29
        } else {
            if true {
                fmt.Println(v4, fi25, fi25)
            } else {
                fmt.Println(v17, v4, v2)
                var v30 float32 = 3.14159
                _ = v30
            }
        }
        _ = fi25
    }
    var v31 uint32 = 678
    v1 = v1
    switch v2 {
    case "hello":
        v32 := []int{128, 348, -9223372036854775808}
        _ = v32
    case "foo bar":
        fmt.Println(v6)
    }
    var v33 int16 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v6
}
