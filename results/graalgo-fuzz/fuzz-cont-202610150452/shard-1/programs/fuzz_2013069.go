package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 bool }
func genFn0() bool {
    v7 := []int{128, -1, 209, 4294967295}
    var v8 uint16 = 848
    _ = v7
    _ = v8
    return ((true || true) && true)
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    var v5 bool = false
    var v6 int32 = -230
    var v9 int8 = 127
    v6 -= 0
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 bool = false
        v12 := GenS1{F0: ("\n\t" + "")}
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: ("" + "")}
    v9 += -1
    v6 += (v6 << ((480) & 63))
    switch "foo bar" {
    case "foo bar":
        if v5 {
            v14 := GenS2{F0: ((("foo bar" + "foo bar") + "") + ("hello" + "")), F1: v5}
            fmt.Println(v6)
            _ = v14
        } else {
            v6 -= (v6 % (((v6 + 1)) | 1))
        }
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 uint64 = 919
        _ = fi15
        _ = v16
    }
    fmt.Println(v5)
    var v17 float64 = 0.1
    v4[int(v6)] = (int64(v1) != 1)
    if (v5 && v5) {
        if (v9 != v9) {
            for fi18 := 0; fi18 < 12; fi18++ {
                if (!(("" + "\n\t") >= "\n\t")) {
                    v2[4294967295] = int(v6)
                    fmt.Println(v17, v5, fi18)
                } else {
                    var v19 int32 = v6
                    _ = v19
                }
                var v20 uint64 = 839
                for fi21 := 0; fi21 < 9; fi21++ {
                    v22 := []int{}
                    fmt.Println(v20)
                    _ = fi21
                    _ = v22
                }
                _ = fi18
                _ = v20
            }
            if (!true) {
                var v23 int32 = 2147483647
                var v24 uint = 255
                if ((!v5) || (v6 < (^(v23)))) {
                    v6 = v6
                    fmt.Println(v24)
                    switch "a" {
                    case "foo bar":
                        switch ("" + (("" + ("" + "")) + "hello")) {
                        case "foo bar":
                            v23 += v23
                        case "a":
                            var v25 bool = (1 <= 9223372036854775807)
                            _ = v25
                        }
                    case "\n\t":
                        fmt.Println(v6, v17)
                    default:
                        v4[127] = (127 > 1)
                    }
                }
                _ = v23
                _ = v24
            } else {
                var v26 uint = 255
                v26 -= 0
                _ = v26
            }
        }
        v27 := []int{-40, 32768, 9223372036854775807, 42}
        fmt.Println(v1, v1, v6)
        _ = v27
    }
    if ("\n\t" != "hello") {
        var v28 float32 = (0.0 * 0.1)
        var v29 float32 = (v28 + -2.25)
        _ = v28
        _ = v29
    }
    if false {
        _ = len(v2)
        v17 = v17
    } else {
        for fi30 := 0; fi30 < 13; fi30++ {
            delete(v2, 1)
            v1 = (v1 << ((1) & 63))
            _ = fi30
        }
        v17 -= 0.0
        if (255 == 255) {
            v9 += v9
        }
    }
    var v31 int8 = v9
    v31 = v9
    fmt.Println(v1, v17, v17)
    var v32 int8 = v31
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
