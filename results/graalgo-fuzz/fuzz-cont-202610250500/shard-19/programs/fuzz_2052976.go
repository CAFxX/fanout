package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int; F1 int64; F2 int }

func main() {
    var v1 int64 = -442
    v2 := make(map[string]uint32)
    v3 := []string{("\n\t" + "\n\t")}
    var v4 string = "hello"
    var v5 uint = 772
    fmt.Println(v4, v4)
    v6 := GenS1{F0: (false || ((v5 & (v5 / ((0) | 1))) <= 255)), F1: 128}
    var v7 float32 = 0.0
    delete(v2, v4)
    if (0 != 745) {
        v3 = append(v3, v4)
        v8 := GenS1{F0: true, F1: 32768}
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS1{F0: false, F1: fi9}
            delete(v2, v4)
            if false {
                v11 := GenS1{F0: true, F1: fi9}
                v5 -= 65535
                v12 := GenS1{F0: (!((true || true) || (((true && (!false)) || ((v1 & v1) > v1)) && ((3.14159 - (float64((((v7 - float32(0.0)) + 3.14159) * 3.14159)) + 1e10)) != 3.14159)))), F1: (fi9 & (fi9 | fi9))}
                _ = v11
                _ = v12
            } else {
                if true {
                    v4 = v4
                }
                v13 := GenS1{F0: true, F1: 32767}
                _ = v13
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        v14 := GenS2{F0: 32768, F1: int64(v5), F2: 4294967295}
        _ = v14
    }
    fmt.Println(v5)
    var v15 int64 = -785
    fmt.Println(v4, v4)
    var v16 bool = true
    var v17 float64 = 1e10
    v2[""] = 2
    var v18 float32 = ((v7 / v7) / v7)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "\n\t" {
    case "\n\t":
        var v19 float64 = 0.0
        _ = v19
    case "":
        var v20 bool = (!false)
        _ = v20
    default:
        for fi21 := 0; fi21 < 5; fi21++ {
            if false {
                var v22 int32 = -82
                _ = v22
            } else {
                v3 = append(v3, v4)
                fmt.Println(v7)
            }
            _ = fi21
        }
    }
    v23 := GenS2{F0: -9223372036854775808, F1: v15, F2: -248}
    var v24 uint8 = 128
    var v25 int32 = 793
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
