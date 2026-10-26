package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int8 = -59
    var v2 uint8 = 2
    var v3 int32 = 2147483647
    var v4 int = -9223372036854775808
    var v5 int64 = 1
    v4 = (^(v4))
    if true {
        switch 850 {
        case 4294967295:
            v6 := []int{(v4 & (^(v4)))}
            _ = v6
        default:
            var v7 int64 = v5
            _ = v7
        }
    } else {
        v8 := []int{(-(v4))}
        if true {
            switch 0 {
            case 128:
                switch v4 {
                case 1:
                    if (true && (0.1 > (float32(float64((-2.25 / (1.5 * 0.1)))) / 3.14159))) {
                        v1 += v1
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        var v9 int8 = -94
                        _ = v9
                    } else {
                        v2 -= (v2 << ((18446744073709551615) & 63))
                    }
                default:
                    fmt.Println(v1, v3, v3)
                }
            case 1:
                var v10 int8 = (v1 % ((127) | 1))
                _ = v10
            }
            v3 -= v3
        } else {
            v1 += v1
            v3 += (v3 >> ((0) & 63))
            v1 = (v1 & 42)
        }
        v1 = -19
        _ = v8
    }
    var v11 string = (("foo bar" + "") + "hello")
    if (((^(v1)) < v1) || true) {
        v12 := []int{(v4 - v4)}
        _ = v12
    } else {
        v13 := []int{v4, (v4 & (-(v4))), v4}
        fmt.Println(v2, v5, v4)
        v3 = v3
        _ = v13
    }
    if (0.0 != (3.14159 - 0.0)) {
        v14 := []int{803, v4}
        switch 1 {
        case 127:
            var v15 uint64 = 2
            _ = v15
        case -707:
            var v16 uint8 = (v2 << ((1) & 63))
            _ = v16
        case 118:
            var v17 int8 = v1
            _ = v17
        default:
            var v18 float64 = 1.5
            _ = v18
        }
        switch (^(v4)) {
        case -2147483648:
            v3 += -1
        case 670:
            v19 := GenS1{F0: 1, F1: 255, F2: 0}
            _ = v19
        case -9223372036854775808:
            var v20 bool = false
            _ = v20
        default:
            var v21 uint64 = 0
            _ = v21
        }
        _ = v14
    } else {
        if (!(3.14159 != 0.0)) {
            v22 := []int{}
            v22 = append(v22, 807)
            _ = v22
        }
        fmt.Println(v1)
        v11 = v11
    }
    if (42 <= int16(v4)) {
        var v23 int32 = -1
        var v24 int32 = v3
        _ = v23
        _ = v24
    } else {
        v25 := []int{(v4 | v4), v4, (v4 >> ((1) & 63)), v4}
        for fi26 := 0; fi26 < 10; fi26++ {
            v4 = 2147483647
            v27 := GenS2{F0: 534}
            if len(v25) > 0 {
                v25[0] = 9223372036854775807
            }
            _ = fi26
            _ = v27
        }
        v5 = v5
        _ = v25
    }
    switch (v4 / (((v4 - v4)) | 1)) {
    case 128:
        var v28 int16 = 1
        _ = v28
    case 2147483647:
        v11 = "foo bar"
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        switch (v11 + v11) {
        case "\n\t":
            switch v11 {
            case "":
                v30 := GenS1{F0: 32767, F1: 0, F2: 895}
                _ = v30
            case "a":
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v11)
                }
            }
        case "foo bar":
            var v31 uint64 = 524
            _ = v31
        case "hello":
            if false {
                v1 += (v1 | (v1 + 42))
            }
        }
        _ = fi29
    }
    fmt.Println(v5)
    if false {
        fmt.Println(v11)
    } else {
        fmt.Println(v5, v4)
        switch v4 {
        case -1:
            v11 = "foo bar"
        }
    }
    v32 := GenS2{F0: 2}
    v5 -= v5
    var v33 uint64 = 18446744073709551615
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
