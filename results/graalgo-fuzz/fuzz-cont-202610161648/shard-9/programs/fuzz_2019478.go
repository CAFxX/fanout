package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint; F1 int64; F2 uint }
type GenS6 struct { F0 int }

func main() {
    v1 := []int{}
    v2 := []uint32{277, 277}
    var v3 uint8 = 127
    v4 := []int{-2147483648, 127, 233, 9223372036854775807}
    var v5 bool = false
    fmt.Println(v5, v3)
    if ((65535 < 2147483647) || (0.0 > 0.0)) {
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
        fmt.Println(v5, v5)
        if (!v5) {
            var v6 float64 = 0.0
            switch "" {
            case "foo bar":
                var v7 int32 = -546
                _ = v7
            case "a":
                var v8 int = -839
                _ = v8
            case "":
                var v9 uint = 455
                _ = v9
            default:
                if (v5 || v5) {
                    v1 = append(v1, -1)
                    v5 = v5
                    v3 += v3
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            _ = v6
        }
    }
    if (v5 || (0 >= uint16(v3))) {
        v5 = v5
    } else {
        if (1 == 16) {
            _ = len(v2)
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := GenS1{F0: 42, F1: 65535, F2: fi10}
                _ = fi10
                _ = v11
            }
        }
        _ = len(v2)
    }
    if (42 >= -751) {
        v5 = v5
        var v12 bool = v5
        v3 = v3
        _ = v12
    } else {
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 float64 = (0.1 * ((1e10 - ((0.1 + 1.5) + -2.25)) + -2.25))
            _ = fi13
            _ = v14
        }
    }
    switch "foo bar" {
    case "a":
        v5 = v5
    case "hello":
        v5 = v5
    default:
        switch -644 {
        case 42:
            v5 = (!(24 >= 0))
        case -1:
            if v5 {
                var v15 int8 = 0
                v15 = 6
                _ = v15
            }
        case 162:
            if len(v4) > 0 {
                v4[0] = -1
            }
        }
    }
    v16 := GenS2{F0: -224}
    var v17 string = "a"
    v3 -= (v3 + (v3 | 255))
    v2 = append(v2, 4294967295)
    v3 = (v3 << ((18446744073709551615) & 63))
    if (v17 > v17) {
        switch 4294967295 {
        case 4294967295:
            if v5 {
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                var v18 int = -417
                _ = v18
            }
        case 127:
            var v19 int8 = -67
            _ = v19
        case 131:
            v20 := GenS1{F0: 899, F1: 156, F2: -2147483648}
            _ = v20
        }
        if (v3 == v3) {
            if v5 {
                for fi21 := 0; fi21 < 13; fi21++ {
                    var v22 uint16 = 882
                    v17 = v17
                    v23 := GenS2{F0: 1}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v3 -= 128
                v24 := GenS3{F0: 0, F1: 65535, F2: 632}
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            fmt.Println(v3)
            v17 = (v17 + "\n\t")
            v5 = (float32(0.1) > (0.0 + (0.0 * 0.1)))
            _ = fi25
        }
    } else {
        if len(v4) > 0 {
            v4[0] = 128
        }
    }
    v26 := GenS4{F0: 2147483648, F1: 97, F2: v3}
    v3 = (v3 ^ v3)
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 uint8 = 128
        var v29 float32 = (0.1 * (3.14159 - (-2.25 * 1e10)))
        for fi30 := 0; fi30 < 19; fi30++ {
            v1 = append(v1, (fi30 | (^(fi27))))
            v31 := GenS3{F0: 32767, F1: 99, F2: 32767}
            v32 := GenS5{F0: 4294967295, F1: 0, F2: 2}
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    if ((!(1 >= 42)) || v5) {
        fmt.Println(v17, v3)
        fmt.Println(v3, v17, v5)
        var v33 uint64 = 18446744073709551615
        _ = v33
    } else {
        switch -481 {
        case 2147483647:
            var v34 int8 = -128
            _ = v34
        default:
            var v35 int = -1
            _ = v35
        }
        for fi36 := 0; fi36 < 11; fi36++ {
            if (111 >= 255) {
                v37 := GenS6{F0: fi36}
                for fi38 := 0; fi38 < 7; fi38++ {
                    switch (v17 + "") {
                    case "foo bar":
                        _ = len(v1)
                    case "a":
                        var v39 uint8 = v3
                        _ = v39
                    }
                    if len(v2) > 0 {
                        v2[0] = 603
                    }
                    _ = fi38
                }
                var v40 int32 = 2147483647
                _ = v37
                _ = v40
            }
            _ = fi36
        }
    }
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
