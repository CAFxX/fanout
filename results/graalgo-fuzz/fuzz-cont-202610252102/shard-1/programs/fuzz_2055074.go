package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint8; F1 uint }

func main() {
    var v1 float32 = float32(1.5)
    var v2 float64 = 0.0
    var v3 bool = true
    v4 := GenS1{F0: v3}
    if true {
        var v5 uint64 = 302
        _ = v5
    }
    var v6 uint16 = 602
    fmt.Println(v2, v6, v2)
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 float32 = 0.1
        fmt.Println(v1, fi7, v3)
        _ = fi7
        _ = v8
    }
    var v9 float32 = -2.25
    var v10 int64 = 9223372036854775807
    v11 := []int{32768, 32768}
    _ = len(v11)
    v3 = v3
    if len(v11) > 0 {
        v11[0] = 4294967295
    }
    switch "a" {
    case "foo bar":
        if (v3 && (42 >= 161)) {
            v9 -= ((-2.25 / float32(v2)) / v9)
        } else {
            if (("foo bar" > "a") && true) {
                _ = len(v11)
                v2 -= 1e10
            }
            var v12 uint64 = 4294967295
            _ = v12
        }
    case "a":
        fmt.Println(v6, v6)
    case "":
        if len(v11) > 0 {
            v11[0] = 2147483647
        }
    default:
        v11 = append(v11, -446)
    }
    v11 = append(v11, 9223372036854775807)
    var v13 uint8 = 127
    if true {
        if (100 >= 24) {
            if (859 <= -1) {
                var v14 uint8 = (^(v13))
                _ = v14
            }
            switch 1 {
            case 42:
                if false {
                    if v3 {
                        v6 += (v6 | 2)
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v16 := GenS1{F0: (("a" + ("hello" + "\n\t")) <= "foo bar")}
                            fmt.Println(v6)
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        v17 := GenS2{F0: -1, F1: 0, F2: 706}
                        _ = v17
                    }
                    _ = len(v11)
                } else {
                    if (4294967295 <= 716) {
                        if (v13 == (v13 * 135)) {
                            var v18 int16 = 1000
                            var v19 float32 = v1
                            _ = v18
                            _ = v19
                        }
                        fmt.Println(v10)
                        v10 -= 1
                    }
                }
            case 127:
                v3 = (!(v3 || (42 == 65535)))
            default:
                if v3 {
                    v6 = v6
                } else {
                    var v20 int16 = -1
                    v21 := GenS1{F0: (100 <= 42)}
                    _ = v20
                    _ = v21
                }
            }
        } else {
            var v22 int = int(v13)
            v23 := GenS1{F0: v3}
            _ = v22
            _ = v23
        }
        v3 = false
        fmt.Println(v2, v9)
    }
    v3 = (-5 == -101)
    var v24 string = "hello"
    if true {
        var v25 int16 = -1
        _ = v25
    } else {
        v26 := GenS2{F0: v10, F1: 562, F2: -32768}
        v27 := GenS3{F0: 200, F1: uint(v13)}
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
