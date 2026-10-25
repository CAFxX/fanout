package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }

func main() {
    var v1 int16 = 32767
    v2 := []string{}
    v3 := []string{("a" + "\n\t"), "\n\t", ("foo bar" + "foo bar"), ""}
    var v4 int = 32767
    var v5 int8 = 1
    fmt.Println(v5, v4)
    v1 = 73
    if true {
        if len(v3) > 0 {
            v3[0] = (("a" + ("hello" + "")) + "\n\t")
        }
    }
    var v6 float64 = (-2.25 * 1e10)
    v7 := GenS1{F0: v5, F1: 42}
    if false {
        v6 = v6
    } else {
        v1 = v1
        var v8 int64 = 1
        v4 -= v4
        _ = v8
    }
    var v9 bool = true
    var v10 float64 = ((v6 * 3.14159) * ((v6 + (float64(1e10) + v6)) * v6))
    fmt.Println(v6, v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 uint16 = 65535
    v1 = 42
    if v9 {
        switch v4 {
        case 4294967295:
            fmt.Println(v4, v10)
        case -41:
            v12 := GenS1{F0: 0, F1: -2}
            _ = v12
        case 42:
            switch (-(v4)) {
            case 128:
                fmt.Println(v10)
            default:
                v9 = ((!(!true)) && false)
            }
        }
    } else {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 uint32 = 2
            v15 := GenS1{F0: (-(v5)), F1: 2147483647}
            switch (^(v4)) {
            case 1:
                var v16 float32 = -2.25
                _ = v16
            case -9223372036854775808:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 32768:
                v17 := GenS1{F0: -1, F1: 42}
                _ = v17
            default:
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 uint8 = 255
                    var v20 uint16 = (v11 ^ uint16(v14))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        var v21 int64 = 9223372036854775807
        _ = v21
    }
    var v22 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
