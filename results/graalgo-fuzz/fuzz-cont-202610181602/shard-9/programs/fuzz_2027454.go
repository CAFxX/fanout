package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0() int8 {
    var v6 string = "hello"
    var v7 string = v6
    var v8 int8 = 100
    _ = v6
    _ = v7
    _ = v8
    return (v8 - -1)
}
func genFn1() int32 {
    v9 := []int{}
    var v10 float64 = 1e10
    if true {
        for fi11 := 0; fi11 < 16; fi11++ {
            v10 -= (v10 - (float64(((((1.5 - (float32(1e10) - ((3.14159 + float32((0.1 - (v10 / ((3.14159 - v10) - v10))))) + 0.0))) + (-2.25 * 0.0)) / (1.5 / -2.25)) * (((1e10 + ((1e10 - (-2.25 * 0.1)) + 0.0)) - (3.14159 / float32(v10))) / 0.1))) - 0.1))
            v10 += 3.14159
            _ = fi11
        }
        if (true || (true || true)) {
            if len(v9) > 0 {
                v9[0] = 1
            }
            if true {
                var v12 string = "a"
                for fi13 := 0; fi13 < 3; fi13++ {
                    fmt.Println(v12, v10, fi13)
                    v10 = (v10 / v10)
                    _ = fi13
                }
                _ = v12
            } else {
                v9 = append(v9, -9223372036854775808)
                if (42 <= 9223372036854775807) {
                    v14 := GenS1{F0: 0}
                    v10 += (3.14159 + 0.1)
                    _ = v14
                } else {
                    v10 = (-2.25 * (0.1 - float64((-2.25 * (1.5 / 1.5)))))
                    v10 = 0.1
                }
            }
            v15 := GenS1{F0: 1}
            _ = v15
        }
    }
    var v16 uint32 = 2
    fmt.Println(v16, v10)
    _ = v10
    _ = v16
    _ = v9
    return -921
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 253
    var v3 uint64 = 42
    var v4 int = 804
    var v5 int = v4
    fmt.Println(v3)
    for fi17 := 0; fi17 < 1; fi17++ {
        switch (^(v5)) {
        case 128:
            if true {
                v18 := []int{(fi17 >> ((65535) & 63))}
                var v19 uint32 = 246
                v20 := GenS2{F0: 4294967295, F1: v3, F2: v3}
                _ = v18
                _ = v19
                _ = v20
            }
        case -2147483648:
            var v21 string = "foo bar"
            _ = v21
        case -9223372036854775808:
            v3 += (v3 >> ((65535) & 63))
        default:
            v22 := []int{}
            _ = v22
        }
        v23 := GenS1{F0: v3}
        switch (fi17 - (v5 + (v4 / ((v5) | 1)))) {
        case 127:
            for fi24 := 0; fi24 < 12; fi24++ {
                v25 := GenS1{F0: (v3 + (v3 << ((255) & 63)))}
                var v26 uint32 = 1
                _ = fi24
                _ = v25
                _ = v26
            }
        case 151:
            var v27 string = "foo bar"
            _ = v27
        }
        _ = fi17
        _ = v23
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        for fi29 := 0; fi29 < 18; fi29++ {
            v4 = (fi29 - -858)
            _ = fi29
        }
        _ = fi28
    }
    if (!false) {
        v30 := GenS1{F0: v3}
        v5 -= v4
        switch "foo bar" {
        case "a":
            fmt.Println(v2, v3, v3)
        case "\n\t":
            var v31 int64 = -950
            _ = v31
        }
        _ = v30
    }
    v32 := GenS3{F0: v3, F1: -712}
    v5 = v4
    var v33 uint64 = (v3 >> ((195) & 63))
    if (18446744073709551615 != 65535) {
        fmt.Println(v33, v4)
        switch v5 {
        case 128:
            for fi34 := 0; fi34 < 9; fi34++ {
                v3 = v3
                _ = fi34
            }
        default:
            if ((3.14159 - 0.0) <= 0.1) {
                v35 := []int{(v4 >> ((255) & 63)), v5, (v5 & v4)}
                var v36 int64 = 42
                _ = v35
                _ = v36
            }
        }
    } else {
        v2 += (v2 >> ((3) & 63))
        fmt.Println(v5)
    }
    var v37 int8 = 127
    var v38 int8 = genFn0()
    for fi39 := 0; fi39 < 11; fi39++ {
        v4 = v5
        if (65535 == 2) {
            fmt.Println(v5)
        } else {
            v33 += (v33 / (((v3 / ((v33) | 1))) | 1))
            v4 = 1
            v2 = v2
        }
        v5 -= (v4 * (v4 & (v5 - v4)))
        _ = fi39
    }
    var v40 int = (v4 / ((v5) | 1))
    var v41 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
