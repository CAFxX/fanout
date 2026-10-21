package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }

func main() {
    var v1 int8 = 1
    var v2 float64 = 0.1
    var v3 uint16 = 892
    var v4 int32 = 1
    var v5 float32 = 3.14159
    var v6 uint16 = v3
    var v7 uint16 = (v3 >> ((42) & 63))
    switch ("" + (("foo bar" + "a") + "a")) {
    case "":
        fmt.Println(v2)
    }
    var v8 uint64 = 2
    v9 := []int{32767, -9223372036854775808, int(v7), 127}
    if (v2 != (v2 + (3.14159 + (v2 + v2)))) {
        var v10 int = 2147483647
        var v11 int8 = -27
        _ = v10
        _ = v11
    } else {
        var v12 uint16 = (v6 * uint16(v8))
        if true {
            if len(v9) > 0 {
                v9[0] = 4294967295
            }
        } else {
            var v13 float64 = float64(v5)
            _ = v13
        }
        if (42 < 809) {
            if (true || false) {
                v1 = v1
                v9 = append(v9, 32767)
                var v14 uint64 = v8
                _ = v14
            } else {
                var v15 uint8 = 127
                _ = v15
            }
        }
        _ = v12
    }
    var v16 int16 = 32767
    v3 += v7
    v2 = (0.0 * float64(1e10))
    if (v6 > v3) {
        if ((3.14159 * -2.25) >= 1.5) {
            switch 4294967295 {
            case 42:
                v5 = (0.1 + (v5 + 1.5))
            case 239:
                for fi17 := 0; fi17 < 15; fi17++ {
                    var v18 uint32 = 2
                    fmt.Println(v2, v7, v5)
                    fmt.Println(v2)
                    _ = fi17
                    _ = v18
                }
            }
            var v19 int32 = (-(v4))
            _ = v19
        } else {
            if false {
                if true {
                    var v20 string = "foo bar"
                    _ = len(v9)
                    _ = v20
                } else {
                    var v21 uint16 = 546
                    _ = v21
                }
                v5 -= v5
                switch 0 {
                case 128:
                    for fi22 := 0; fi22 < 8; fi22++ {
                        fmt.Println(v16, v8)
                        _ = fi22
                    }
                }
            }
            v2 += v2
            _ = len(v9)
        }
        var v23 float32 = 1.5
        _ = v23
    } else {
        v4 -= v4
    }
    v24 := GenS1{F0: -78, F1: 185}
    v25 := GenS1{F0: 42, F1: 127}
    if true {
        if ((true || ((v4 >> ((76) & 63)) < (v4 << ((0) & 63)))) && (1.5 < -2.25)) {
            _ = len(v9)
            v26 := GenS1{F0: 0, F1: 0}
            _ = v26
        } else {
            var v27 float64 = float64(v5)
            var v28 uint16 = v7
            _ = v27
            _ = v28
        }
        if len(v9) > 0 {
            v9[0] = -1
        }
    } else {
        var v29 uint8 = 0
        _ = len(v9)
        _ = v29
    }
    v6 -= 40
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
