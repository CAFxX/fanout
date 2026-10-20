package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int32 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            var v8 bool = (1e10 <= ((0.0 + 0.1) + (1e10 + (0.1 * -2.25))))
            v8 = v8
            _ = v8
        } else {
            v9 := []int{fi7, 32767}
            _ = v9
        }
        switch ("" + ("foo bar" + ("" + "hello"))) {
        case "hello":
            var v10 uint8 = 2
            _ = v10
        }
        var v11 string = "\n\t"
        _ = fi7
        _ = v11
    }
    var v12 int = 4294967295
    if (0 > 638) {
        for fi13 := 0; fi13 < 8; fi13++ {
            v12 -= v12
            v12 += (fi13 & v12)
            _ = fi13
        }
        v14 := []int{}
        _ = v14
    } else {
        if true {
            v15 := GenS1{F0: (42 <= 586), F1: 255}
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int8 = 1
            _ = fi16
            _ = v17
        }
    }
    fmt.Println(v12)
    v12 = (v12 + 619)
    _ = v12
    return 65535
}

func main() {
    var v1 int = -552
    v2 := make(map[string]int)
    var v3 int = (^(v1))
    v4 := make(map[uint32]bool)
    var v5 int32 = 42
    var v6 int64 = 42
    if (((3.14159 * (((0.1 + (0.0 / (0.1 + (-2.25 - 0.1)))) / 1.0) + (1.5 - (0.0 + float32(0.1))))) != 3.14159) && (!false)) {
        v18 := []int{(v1 * v3)}
        _ = v18
    }
    v1 = 9223372036854775807
    var v19 uint = 4294967295
    _ = len(v2)
    v6 = (v6 << ((170) & 63))
    switch 32767 {
    case 9223372036854775807:
        fmt.Println(v3, v5)
    case 4294967295:
        fmt.Println(v5)
    case 0:
        var v20 int32 = (v5 << ((v19) & 63))
        _ = v20
    }
    v1 = v1
    if false {
        for fi21 := 0; fi21 < 8; fi21++ {
            v22 := GenS2{F0: 9223372036854775807}
            v23 := GenS3{F0: (("" + "foo bar") + "foo bar")}
            v24 := []int{32767}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        if true {
            v25 := GenS4{F0: -126, F1: (!(2 < 584))}
            _ = v25
        } else {
            if true {
                if ((1e10 / 1.5) >= (-2.25 * 1e10)) {
                    switch v3 {
                    case -988:
                        v26 := GenS5{F0: 255, F1: (v19 >> (((v19 | (v19 | 18446744073709551615))) & 63))}
                        _ = v26
                    case 32768:
                        v2["a"] = -2147483648
                    case 127:
                        var v27 uint64 = 1
                        _ = v27
                    }
                    var v28 uint16 = 0
                    v19 += 0
                    _ = v28
                } else {
                    switch "" {
                    case "hello":
                        v29 := []int{(v3 >> ((18446744073709551615) & 63)), -2147483648}
                        _ = v29
                    case "a":
                        var v30 uint16 = 402
                        _ = v30
                    case "foo bar":
                        var v31 int = 935
                        _ = v31
                    default:
                        fmt.Println(v1)
                    }
                    v32 := GenS6{F0: v5}
                    _ = v32
                }
                for fi33 := 0; fi33 < 9; fi33++ {
                    v6 = (-(v6))
                    _ = fi33
                }
            } else {
                var v34 int = 254
                if (!(18446744073709551615 > 1)) {
                    v34 += v34
                    var v35 uint16 = 32767
                    var v36 int32 = (v5 & v5)
                    _ = v35
                    _ = v36
                }
                _ = v34
            }
            v4[65535] = false
        }
        if (false && true) {
            if false {
                v37 := GenS4{F0: 127, F1: false}
                if (!true) {
                    v6 -= 9223372036854775807
                    if (!(v19 == v19)) {
                        v38 := GenS5{F0: (v19 & v19), F1: 136}
                        var v39 float64 = 1e10
                        v5 = -1
                        _ = v38
                        _ = v39
                    }
                } else {
                    for fi40 := 0; fi40 < 7; fi40++ {
                        v1 -= -538
                        _ = fi40
                    }
                }
                _ = len(v4)
                _ = v37
            } else {
                fmt.Println(v6)
            }
            v41 := []int{int(v19), 127, (^(v3))}
            _ = v41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
