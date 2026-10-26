package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int64 }
func genFn0() int {
    var v4 uint8 = 255
    for fi5 := 0; fi5 < 9; fi5++ {
        if true {
            var v6 int16 = 427
            v6 += v6
            _ = v6
        }
        var v7 float64 = (-2.25 * 0.0)
        _ = fi5
        _ = v7
    }
    var v8 float64 = 0.0
    v9 := []int{2147483647, -2147483648, 127, 0}
    fmt.Println(v8, v8)
    _ = v4
    _ = v8
    _ = v9
    return -168
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]bool)
    var v3 uint8 = 43
    v1 = (false && (false && (!(42 == 0))))
    var v10 uint16 = 32767
    v11 := []int{32767, 42, 196}
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint8 = (v3 ^ (v3 >> ((80) & 63)))
            switch 2147483647 {
            case 2147483647:
                v10 = v10
            case 321:
                switch fi12 {
                case 0:
                    fmt.Println(v14)
                default:
                    fmt.Println(v10, fi12)
                }
            default:
                if ((-(fi12)) > fi12) {
                    v15 := GenS1{F0: "foo bar", F1: 32767, F2: int64(v10)}
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    if ("a" > ("\n\t" + ("hello" + "foo bar"))) {
                        switch fi12 {
                        case 117:
                            v3 = v3
                        }
                        var v16 int32 = 247
                        fmt.Println(fi12)
                        v3 = 1
                        fmt.Println(v10, fi13, fi12)
                        _ = v16
                    }
                    _ = v15
                }
            }
            var v17 int16 = 24
            var v18 int16 = 42
            _ = fi13
            _ = v14
            _ = v17
            _ = v18
        }
        var v19 float32 = float32(1e10)
        _ = fi12
        _ = v19
    }
    v2[0] = v1
    v3 = (v3 << ((304) & 63))
    v20 := v2[1]
    v21 := GenS1{F0: (("\n\t" + "hello") + "a"), F1: -727, F2: 42}
    if (-60 != 11) {
        fmt.Println(v3)
        v20 = false
    } else {
        if (v1 && (!false)) {
            _ = len(v11)
        } else {
            if false {
                fmt.Println(v20)
            }
            if v1 {
                var v22 int8 = -1
                _ = v22
            } else {
                v10 = v10
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            }
            v10 -= v10
        }
        var v23 int32 = 2147483647
        var v24 int64 = -9223372036854775808
        _ = v23
        _ = v24
    }
    if v1 {
        switch "\n\t" {
        case "":
            v1 = v20
        }
    } else {
        v20 = (-1 != 0)
    }
    switch "hello" {
    case "":
        v25 := GenS1{F0: "", F1: -1, F2: 1}
        _ = v25
    }
    v26 := GenS1{F0: "hello", F1: 738, F2: 9223372036854775807}
    fmt.Println(v1, v3)
    var v27 int32 = -976
    v28 := GenS1{F0: "hello", F1: 963, F2: 42}
    if ("" == (("" + "\n\t") + "foo bar")) {
        v29 := GenS1{F0: "foo bar", F1: 0, F2: 1}
        _ = v29
    }
    fmt.Println(v27, v10, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
