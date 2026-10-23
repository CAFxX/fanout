package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int32 {
    if (!(false || false)) {
        var v4 bool = true
        _ = v4
    }
    if (!(((3.14159 + 3.14159) * 1e10) >= (1e10 / -2.25))) {
        var v5 int8 = -128
        _ = v5
    } else {
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 int32 = 579
            _ = fi6
            _ = v7
        }
    }
    var v8 uint32 = 458
    v8 = 65535
    _ = v8
    return -1
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]uint32)
    v3 := []string{"", ("" + "foo bar"), "foo bar"}
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS1{F0: 42}
        v11 := GenS1{F0: -85}
        if (42 > 1) {
            if (1 != -1) {
                if true {
                    var v12 uint8 = 126
                    var v13 uint = 2
                    _ = v12
                    _ = v13
                }
                var v14 uint8 = 128
                var v15 uint16 = 1
                _ = v14
                _ = v15
            }
        } else {
            if ((((-2.25 * 1.5) * 0.1) + 3.14159) >= 1e10) {
                fmt.Println(fi9, fi9)
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := GenS1{F0: 23}
                    _ = len(v2)
                    v1 = v1
                    v18 := v2[550]
                    switch "a" {
                    case "a":
                        fmt.Println(v1, fi16, fi9)
                        v18 = uint32(v1)
                        v1 -= (v1 & 2)
                        switch "" {
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                        }
                        var v19 int64 = 105
                        if false {
                            switch "hello" {
                            case "":
                                if false {
                                    var v20 int16 = int16(v19)
                                    var v21 int = (^(fi16))
                                    v22 := GenS1{F0: 127}
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                            default:
                                switch "" {
                                case "a":
                                    var v23 float64 = float64(3.14159)
                                    _ = v23
                                }
                            }
                        }
                        _ = v19
                    default:
                        if (!(!true)) {
                            fmt.Println(fi16)
                            v18 = 9
                        }
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            } else {
                var v24 float32 = (1e10 * 1.5)
                _ = len(v3)
                _ = v24
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v25 int8 = 126
    _ = len(v3)
    var v26 int = -9223372036854775808
    var v27 float64 = (float64(3.14159) / float64((0.1 * (float32(-2.25) - (0.1 + (1e10 / -2.25))))))
    v25 += (v25 * 100)
    if true {
        for fi28 := 0; fi28 < 14; fi28++ {
            v1 = 115
            v29 := GenS1{F0: (v25 << ((18446744073709551615) & 63))}
            _ = fi28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        switch (fi30 << ((210) & 63)) {
        case 2147483647:
            switch (("hello" + "foo bar") + ("foo bar" + ("" + "a"))) {
            case "hello":
                var v31 float64 = -2.25
                _ = v31
            case "foo bar":
                for fi32 := 0; fi32 < 17; fi32++ {
                    v27 = v27
                    _ = fi32
                }
            case "":
                fmt.Println(fi30, fi30)
            default:
                v27 = -2.25
            }
        default:
            var v33 float64 = 1.5
            _ = v33
        }
        fmt.Println(v27, v26, v26)
        v1 -= (^(v1))
        _ = fi30
    }
    fmt.Println(v26, v26, v25)
    var v34 uint16 = uint16(v1)
    switch (v26 << ((1) & 63)) {
    case 4294967295:
        fmt.Println(v26, v25)
    case 128:
        v27 = float64(3.14159)
    case 234:
        for fi35 := 0; fi35 < 20; fi35++ {
            switch "a" {
            case "hello":
                var v36 float32 = 1e10
                _ = v36
            case "a":
                fmt.Println(v26)
            case "\n\t":
                fmt.Println(v27, v25)
            default:
                var v37 uint = 132
                _ = v37
            }
            v1 -= v1
            _ = fi35
        }
    default:
        if len(v3) > 0 {
            v3[0] = "a"
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
}
