package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }
func genFn0() uint {
    v6 := GenS1{F0: 1, F1: 0, F2: 255}
    v7 := GenS1{F0: 163, F1: 0, F2: 65535}
    switch "\n\t" {
    case "a":
        var v8 int = 293
        _ = v8
    case "hello":
        for fi9 := 0; fi9 < 9; fi9++ {
            if (!true) {
                var v10 float32 = 0.1
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 int32 = 472
                    var v13 int64 = -9223372036854775808
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = v10
            }
            _ = fi9
        }
    }
    v14 := []int{-2147483648, 263, 4294967295, 42}
    _ = v14
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := []int{}
    var v2 int8 = 1
    var v3 float32 = (1e10 * -2.25)
    var v4 int32 = 949
    var v5 uint64 = 42
    v15 := GenS1{F0: 103, F1: 0, F2: 912}
    for fi16 := 0; fi16 < 19; fi16++ {
        if (!false) {
            v17 := GenS1{F0: 1, F1: 1, F2: 929}
            if false {
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 int64 = 1
                    var v20 bool = ((int16(v4) <= 916) && (65535 == 2147483648))
                    switch (("foo bar" + "\n\t") + "a") {
                    case "foo bar":
                        var v21 int8 = -1
                        _ = v21
                    case "a":
                        var v22 uint = 42
                        _ = v22
                    default:
                        v4 -= (v4 - (v4 / ((0) | 1)))
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                fmt.Println(v3)
                switch fi16 {
                case 2147483647:
                    var v23 uint64 = v5
                    _ = v23
                case 0:
                    v2 -= (v2 - v2)
                default:
                    v3 = (v3 + 0.0)
                }
            } else {
                v5 = (v5 >> ((65535) & 63))
                var v24 int64 = 42
                v5 -= 4294967295
                _ = v24
            }
            _ = v17
        } else {
            var v25 int64 = 1
            _ = v25
        }
        v26 := GenS1{F0: 1, F1: 32767, F2: 65535}
        _ = fi16
        _ = v26
    }
    v4 = (v4 << ((4294967295) & 63))
    var v27 int = 2147483647
    var v28 uint16 = 255
    for fi29 := 0; fi29 < 9; fi29++ {
        v27 -= fi29
        _ = fi29
    }
    fmt.Println(v5)
    for fi30 := 0; fi30 < 13; fi30++ {
        v31 := GenS1{F0: 1, F1: v28, F2: 42}
        _ = fi30
        _ = v31
    }
    if false {
        v28 = v28
    }
    switch ("foo bar" + ("hello" + "a")) {
    case "\n\t":
        var v32 uint16 = 42
        _ = v32
    case "":
        var v33 uint16 = v28
        _ = v33
    }
    v27 += (v27 + 9223372036854775807)
    v34 := GenS1{F0: 0, F1: v28, F2: (v28 ^ (^(v28)))}
    v4 += (v4 ^ v4)
    if false {
        v2 = v2
        for fi35 := 0; fi35 < 11; fi35++ {
            if (((1.5 + (v3 / 1.0)) / v3) < ((((v3 + 1.5) - 0.0) - 1.5) + float32((3.14159 - ((float64(1e10) / -2.25) - 1.5))))) {
                var v36 float32 = ((v3 + float32(1e10)) - 1e10)
                v5 -= 1
                var v37 uint32 = 4294967295
                _ = v36
                _ = v37
            } else {
                fmt.Println(v4)
                var v38 bool = false
                v39 := GenS1{F0: -309, F1: (v28 - v28), F2: (v28 | v28)}
                _ = v38
                _ = v39
            }
            _ = fi35
        }
    }
    if len(v1) > 0 {
        v1[0] = (^(v27))
    }
    var v40 int64 = int64(v28)
    fmt.Println(v4, v4, v40)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v40
    _ = v5
}
