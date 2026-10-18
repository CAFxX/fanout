package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }
func genFn0() uint16 {
    if ((0 > -396) && ((!false) || true)) {
        var v5 float64 = 1e10
        _ = v5
    } else {
        v6 := []int{-268, 190, 177}
        for fi7 := 0; fi7 < 16; fi7++ {
            if true {
                var v8 int = 9223372036854775807
                _ = v8
            }
            _ = fi7
        }
        var v9 bool = ("a" < "hello")
        _ = v6
        _ = v9
    }
    switch 127 {
    case 2147483647:
        v10 := GenS1{F0: 152}
        _ = v10
    case 127:
        if false {
            var v11 int16 = 0
            var v12 bool = (3.14159 > ((3.14159 + 1e10) / 1e10))
            _ = v11
            _ = v12
        } else {
            var v13 uint = 4294967295
            v14 := []int{185, -813, -9223372036854775808, 564}
            for fi15 := 0; fi15 < 20; fi15++ {
                switch (("a" + "\n\t") + "a") {
                case "\n\t":
                    var v16 float64 = ((3.14159 + (float64(((0.0 + float32(1.5)) - ((1.5 / 1.0) - ((float32(1e10) * 0.0) * -2.25)))) - 0.1)) * float64(1.5))
                    _ = v16
                default:
                    v17 := GenS2{F0: 18446744073709551615, F1: 779}
                    _ = v17
                }
                v13 = (v13 | 1)
                v13 -= (v13 >> ((v13) & 63))
                _ = fi15
            }
            _ = v13
            _ = v14
        }
    default:
        var v18 float32 = float32(1.5)
        _ = v18
    }
    var v19 uint = 2
    _ = v19
    return 58
}
func genFn1() uint32 {
    var v20 uint8 = 42
    var v21 uint8 = 127
    _ = v20
    _ = v21
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 34
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v22 int8 = 84
    switch "foo bar" {
    case "\n\t":
        v23 := []int{-9223372036854775808, -1, 1, -2147483648}
        _ = v23
    case "foo bar":
        v24 := []int{78, -635, -2147483648, 128}
        _ = v24
    default:
        var v25 uint64 = 42
        _ = v25
    }
    var v26 int8 = (v22 * (v22 * v22))
    var v27 int = -57
    if ((1 <= 112) && true) {
        if (42 < 32767) {
            v28 := []int{(v27 & v27), (v27 << ((307) & 63)), (v27 - (v27 - (v27 << ((65535) & 63)))), v27}
            v2 = v2
            _ = v28
        }
    }
    v29 := GenS2{F0: 2, F1: 172}
    if true {
        fmt.Println(v26)
    } else {
        var v30 int32 = 2147483647
        _ = v30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        v22 -= v26
        var v32 float32 = 1e10
        v2 = (v2 / ((2) | 1))
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        var v34 uint8 = 2
        var v35 string = ("\n\t" + "foo bar")
        _ = fi33
        _ = v34
        _ = v35
    }
    v27 = 4294967295
    switch (v27 ^ v27) {
    case 32767:
        var v36 uint32 = (v2 / (((v2 | 0)) | 1))
        _ = v36
    default:
        if (true && (22 < -32768)) {
            var v37 int16 = -32768
            _ = v37
        }
    }
    var v38 float64 = (0.1 * (1e10 * (1.5 / 3.14159)))
    var v39 int = 127
    fmt.Println(v2, v38, v2)
    v22 -= (v26 * (-(v22)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
