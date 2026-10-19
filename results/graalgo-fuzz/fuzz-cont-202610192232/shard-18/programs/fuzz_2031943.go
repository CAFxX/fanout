package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int8 = 45
    var v2 uint = 42
    v3 := make(map[uint32]string)
    var v4 uint8 = 173
    v5 := make(map[uint32]bool)
    if true {
        v1 += (v1 - v1)
    } else {
        if (65535 < 65535) {
            switch 2147483647 {
            case 2147483647:
                var v6 int8 = 109
                _ = v6
            }
            v7 := GenS1{F0: 694, F1: 0, F2: 0}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := GenS1{F0: 631, F1: 2147483647, F2: fi8}
                v4 -= (v4 / (((v4 & (v4 << ((4294967295) & 63)))) | 1))
                _ = fi8
                _ = v9
            }
            fmt.Println(v1, v4, v4)
            var v10 int64 = 9223372036854775807
            _ = v10
        }
        v2 += (v2 / (((v2 ^ (v2 * (^(v2))))) | 1))
    }
    if false {
        v1 = v1
        var v11 int64 = 1
        if ((!false) && (uint16(v4) > 32767)) {
            v4 -= v4
            v12 := GenS1{F0: (v2 | (v2 + 4294967295)), F1: 32768, F2: 9223372036854775807}
            _ = v12
        } else {
            v11 -= v11
            v4 -= 0
            for fi13 := 0; fi13 < 8; fi13++ {
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(v1, fi14, v11)
                    v4 -= (v4 << ((v2) & 63))
                    var v15 string = "a"
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
        }
        _ = v11
    }
    v16 := []int{4294967295, -9223372036854775808}
    switch 332 {
    case 1:
        fmt.Println(v4, v1, v2)
    case 0:
        v16 = append(v16, 187)
    default:
        var v17 int64 = 9223372036854775807
        _ = v17
    }
    if (0 < -106) {
        if (true && true) {
            var v18 int64 = 1
            switch 9223372036854775807 {
            case 9223372036854775807:
                var v19 float64 = (3.14159 - (float64(float32(1e10)) * -2.25))
                _ = v19
            case 2147483647:
                if (133 != 128) {
                    var v20 uint = 290
                    _ = v20
                }
            case -773:
                v21 := GenS1{F0: 4, F1: -100, F2: 4294967295}
                _ = v21
            }
            v1 += (v1 + (v1 - (v1 - 127)))
            _ = v18
        }
        if false {
            v4 += v4
            var v22 int8 = v1
            _ = v22
        } else {
            var v23 uint32 = 2147483648
            var v24 int32 = 214
            _ = v23
            _ = v24
        }
        _ = len(v3)
    }
    if true {
        var v25 int = 4294967295
        v1 = (v1 & (v1 & v1))
        var v26 float64 = (float64(0.1) - float64(0.1))
        _ = v25
        _ = v26
    } else {
        _ = len(v5)
        if ((1e10 < 0.0) || (false && (v2 == (v2 | v2)))) {
            v27 := GenS1{F0: 18446744073709551615, F1: 128, F2: -308}
            if len(v16) > 0 {
                v16[0] = 9223372036854775807
            }
            _ = v27
        }
    }
    if (0.0 <= -2.25) {
        var v28 int16 = -71
        v4 += (v4 & (v4 / (((v4 ^ v4)) | 1)))
        _ = v28
    } else {
        v29 := GenS1{F0: v2, F1: -9223372036854775808, F2: 42}
        v16 = append(v16, 32767)
        _ = v29
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        var v31 uint8 = 0
        v1 = (v1 % ((v1) | 1))
        v1 += (v1 * (v1 | v1))
        _ = fi30
        _ = v31
    }
    switch 128 {
    case 32768:
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    case 0:
        v32 := GenS1{F0: (v2 % ((1) | 1)), F1: 32768, F2: 128}
        _ = v32
    default:
        switch -1 {
        case 372:
            v4 -= (v4 | (v4 / ((v4) | 1)))
        case 42:
            v4 = v4
        case 127:
            v4 += 175
        }
    }
    switch "foo bar" {
    case "foo bar":
        v4 = 128
    case "\n\t":
        v33 := GenS2{F0: -35}
        _ = v33
    case "a":
        if (true && true) {
            if (!false) {
                var v34 float64 = 0.0
                _ = v34
            } else {
                var v35 int16 = 32767
                var v36 float32 = 1e10
                _ = v35
                _ = v36
            }
            v37 := GenS3{F0: v4}
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
