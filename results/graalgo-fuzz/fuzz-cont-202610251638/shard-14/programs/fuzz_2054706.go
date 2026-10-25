package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 int8; F1 uint64; F2 int16 }

func main() {
    var v1 string = "a"
    var v2 uint32 = 2147483647
    var v3 int16 = 0
    var v4 uint32 = v2
    var v5 float32 = 3.14159
    var v6 uint64 = 42
    v6 = v6
    switch v1 {
    case "":
        var v7 int64 = 0
        _ = v7
    case "\n\t":
        fmt.Println(v5)
    }
    if (-639 <= 1) {
        if (0.1 > float32(-2.25)) {
            if false {
                if ((0.0 * 1e10) != 1.5) {
                    var v8 float64 = 3.14159
                    _ = v8
                }
                fmt.Println(v3, v5)
                var v9 uint32 = v2
                _ = v9
            } else {
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 int16 = (v3 % ((v3) | 1))
                    _ = fi10
                    _ = v11
                }
                v3 += 1
                if (!(2 != 4294967295)) {
                    switch "hello" {
                    case "a":
                        v12 := GenS1{F0: true, F1: 394, F2: 185}
                        _ = v12
                    case "\n\t":
                        v3 += v3
                    }
                }
            }
            v3 += int16(v4)
        } else {
            if false {
                v1 = v1
            } else {
                v13 := []int{4294967295, 9223372036854775807, -690, -678}
                var v14 bool = true
                fmt.Println(v5)
                _ = v13
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            v5 = v5
            _ = fi15
        }
        v16 := GenS1{F0: false, F1: 42, F2: -894}
        _ = v16
    } else {
        v17 := GenS2{F0: 102, F1: v6, F2: v3}
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS1{F0: true, F1: uint(v2), F2: 2147483647}
            var v20 uint32 = (v2 + uint32(v6))
            var v21 uint8 = 1
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        _ = v17
    }
    switch v1 {
    case "a":
        var v22 int32 = 42
        _ = v22
    case "foo bar":
        switch v1 {
        case "a":
            for fi23 := 0; fi23 < 7; fi23++ {
                v4 = (^(v2))
                var v24 bool = (!(true || true))
                var v25 uint8 = 42
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    default:
        for fi26 := 0; fi26 < 6; fi26++ {
            v27 := []int{(fi26 - fi26)}
            _ = fi26
            _ = v27
        }
    }
    var v28 float32 = v5
    fmt.Println(v5, v3, v1)
    v29 := []int{82, 32768, 32768, -9223372036854775808}
    if (4294967295 < 379) {
        switch -9223372036854775808 {
        case 42:
            if true {
                switch -2147483648 {
                case 42:
                    var v30 int64 = -1
                    _ = v30
                case -1:
                    if (!((v6 % ((971) | 1)) > v6)) {
                        var v31 int = 2147483647
                        _ = v31
                    } else {
                        v6 = 1
                        var v32 uint8 = 1
                        _ = v32
                    }
                }
                v6 = (v6 << ((2) & 63))
            }
        case -236:
            var v33 int = 0
            _ = v33
        case -799:
            if len(v29) > 0 {
                v29[0] = 1
            }
        default:
            var v34 float32 = v28
            _ = v34
        }
        v35 := GenS2{F0: 42, F1: v6, F2: (v3 ^ (v3 >> ((1) & 63)))}
        _ = v35
    }
    v36 := GenS1{F0: true, F1: 255, F2: 42}
    var v37 string = (v1 + "")
    switch 2147483647 {
    case -1:
        fmt.Println(v37, v28)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
