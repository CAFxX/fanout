package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint }
type GenS2 struct { F0 int8; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint }
type GenS5 struct { F0 uint; F1 int; F2 string }
func genFn0(p0 string) uint {
    switch p0 {
    case "":
        fmt.Println(p0)
    }
    p0 = (p0 + "a")
    v7 := GenS1{F0: -32768, F1: 998, F2: 86}
    if (0.1 == -2.25) {
        var v8 uint8 = 109
        v9 := GenS2{F0: 1, F1: false, F2: 2}
        if true {
            v10 := []int{-2147483648, -1}
            v8 += (v8 >> ((0) & 63))
            _ = v10
        }
        _ = v8
        _ = v9
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS3{F0: 2, F1: -32768}
            v13 := []int{fi11, (-(fi11)), fi11}
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := []int{-913, -2147483648, 127}
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = v14
    }
    if (1 > 0) {
        fmt.Println(p0)
        p0 = p0
    }
    _ = p0
    _ = v7
    return 2
}

func main() {
    var v1 int16 = 1000
    v2 := []uint32{2147483648, 690}
    var v3 uint = 65535
    var v4 uint32 = 42
    v5 := make(map[string]bool)
    var v6 int = 9223372036854775807
    v1 += 1
    if ("\n\t" <= ("" + "hello")) {
        v6 += (v6 ^ (v6 / (((v6 + (v6 + v6))) | 1)))
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        if (1 >= 2147483647) {
            fmt.Println(fi15, v4)
            switch int(v3) {
            case -2147483648:
                _ = len(v5)
            }
        }
        _ = fi15
    }
    v6 -= int(v3)
    var v16 int8 = 100
    v3 = (^(v3))
    fmt.Println(v16, v6, v3)
    if (!((uint8(v4) > 0) || false)) {
        v5[("" + "a")] = false
    }
    v4 -= (v4 * (v4 * v4))
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 string = "foo bar"
            if ((!(((0.0 < (-2.25 * 3.14159)) || true) && (65535 <= 32767))) || ((("foo bar" + v19) + (v19 + ((v19 + v19) + v19))) >= v19)) {
                switch -283 {
                case -1:
                    fmt.Println(v19)
                }
                v5[(((v19 + ("hello" + v19)) + "foo bar") + v19)] = false
                if (-1 == 42) {
                    v16 += v16
                    if (1 <= -801) {
                        var v20 int64 = -898
                        v19 = v19
                        _ = v20
                    } else {
                        if len(v2) > 0 {
                            v2[0] = (^(v4))
                        }
                        v16 = (v16 + v16)
                    }
                } else {
                    v19 = v19
                    var v21 float64 = (float64(0.0) + 0.1)
                    _ = v21
                }
            } else {
                fmt.Println(v3, fi18, v6)
            }
            var v22 int32 = -75
            _ = fi18
            _ = v19
            _ = v22
        }
        v23 := GenS1{F0: (^(v1)), F1: (^(v3)), F2: (v3 ^ 151)}
        _ = fi17
        _ = v23
    }
    v24 := GenS3{F0: v4, F1: 32767}
    switch "a" {
    case "a":
        v25 := GenS4{F0: 42, F1: 2, F2: genFn0("foo bar")}
        _ = v25
    case "\n\t":
        v3 = 661
    case "foo bar":
        v2 = append(v2, v4)
    }
    v1 = v1
    fmt.Println(v3)
    v1 -= v1
    for fi26 := 0; fi26 < 13; fi26++ {
        if false {
            _ = len(v2)
            fmt.Println(v1, fi26, v3)
            v27 := GenS5{F0: genFn0(("hello" + "a")), F1: (v6 + fi26), F2: "\n\t"}
            _ = v27
        }
        v2 = append(v2, v4)
        _ = fi26
    }
    fmt.Println(v4, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
