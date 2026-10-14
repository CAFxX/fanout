package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool) uint16 {
    p0 = true
    v5 := GenS1{F0: 1000, F1: (238 > 42), F2: 0}
    _ = p0
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 int = -2147483648
    var v3 float32 = 0.1
    var v4 int32 = -1
    v4 = (v4 | (v4 - v4))
    v4 += 2147483647
    if false {
        switch ("a" + "") {
        case "a":
            if (!false) {
                for fi6 := 0; fi6 < 8; fi6++ {
                    for fi7 := 0; fi7 < 7; fi7++ {
                        var v8 string = "foo bar"
                        v9 := GenS1{F0: int16(v4), F1: false, F2: (v4 * v4)}
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    v3 = float32(3.14159)
                    _ = fi6
                }
            }
        }
        v2 = (v2 >> ((478) & 63))
    }
    if false {
        if (-618 <= 1000) {
            v4 = v4
            v2 = -416
            for fi10 := 0; fi10 < 8; fi10++ {
                fmt.Println(fi10)
                _ = fi10
            }
        }
        v2 = v2
        v11 := GenS1{F0: -32768, F1: true, F2: (v4 / (((v4 + -2147483648)) | 1))}
        _ = v11
    } else {
        var v12 string = ""
        if (false || false) {
            var v13 int64 = 0
            fmt.Println(v2, v2)
            _ = v13
        } else {
            _ = len(v1)
        }
        fmt.Println(v3, v12)
        _ = v12
    }
    v14 := GenS2{F0: 32767}
    if false {
        v4 -= (v4 << ((255) & 63))
        if (!false) {
            var v15 int = (-(v2))
            _ = v15
        }
        var v16 float64 = -2.25
        _ = v16
    } else {
        if (1 < 65535) {
            delete(v1, ("" + "hello"))
            fmt.Println(v3, v2)
            v17 := []int{(-(v2)), v2, (v2 % (((-(v2))) | 1)), (v2 << ((42) & 63))}
            _ = v17
        } else {
            v18 := []int{(-(v2)), v2, 32767}
            var v19 uint16 = 672
            _ = v18
            _ = v19
        }
    }
    v4 += v4
    if (2 == 1) {
        for fi20 := 0; fi20 < 14; fi20++ {
            switch 361 {
            case 368:
                v21 := []int{}
                _ = v21
            case 128:
                fmt.Println(v3)
            case -9223372036854775808:
                v22 := GenS2{F0: 586}
                _ = v22
            default:
                v4 = (v4 + (v4 | v4))
            }
            _ = fi20
        }
    } else {
        var v23 int16 = 1000
        v24 := []int{v2, (^(v2)), (v2 % ((v2) | 1))}
        var v25 int64 = 1
        _ = v23
        _ = v24
        _ = v25
    }
    delete(v1, "hello")
    if true {
        switch (("" + ("hello" + ("\n\t" + ("\n\t" + "\n\t")))) + "") {
        case "":
            fmt.Println(v4, v2, v4)
        case "hello":
            v2 = 32767
        case "foo bar":
            v3 = 1.5
        }
        var v26 int8 = 1
        _ = v26
    } else {
        var v27 float64 = (-2.25 * float64(v3))
        fmt.Println(v4)
        if false {
            v4 = -1
            for fi28 := 0; fi28 < 2; fi28++ {
                v29 := GenS1{F0: 1000, F1: false, F2: (v4 ^ v4)}
                if false {
                    for fi30 := 0; fi30 < 19; fi30++ {
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v4)
                        case "foo bar":
                            v31 := GenS1{F0: 754, F1: (v4 <= v4), F2: v4}
                            _ = v31
                        case "hello":
                            if (false || true) {
                                var v32 uint64 = 0
                                _ = v32
                            }
                        }
                        v27 = v27
                        _ = fi30
                    }
                }
                v27 -= float64(v3)
                _ = fi28
                _ = v29
            }
            fmt.Println(v3, v3)
        } else {
            var v33 uint64 = 1
            var v34 int = v2
            _ = v33
            _ = v34
        }
        _ = v27
    }
    v3 = ((float32(1.5) / 3.14159) + v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
