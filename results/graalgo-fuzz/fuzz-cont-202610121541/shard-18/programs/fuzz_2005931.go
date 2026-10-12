package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int }
func genFn0() bool {
    var v5 int64 = 9223372036854775807
    var v6 int64 = -340
    for fi7 := 0; fi7 < 2; fi7++ {
        switch fi7 {
        case 128:
            fmt.Println(fi7, fi7)
        case 32767:
            v8 := GenS1{F0: ("foo bar" + "")}
            _ = v8
        default:
            v9 := GenS1{F0: "\n\t"}
            _ = v9
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            if true {
                switch ((("foo bar" + "foo bar") + "\n\t") + ("a" + "hello")) {
                case "hello":
                    var v11 uint8 = 128
                    _ = v11
                default:
                    v12 := GenS1{F0: "hello"}
                    _ = v12
                }
                switch fi7 {
                case 32767:
                    var v13 bool = (!false)
                    _ = v13
                case 128:
                    v5 -= (v6 >> ((2) & 63))
                case 32768:
                    var v14 int8 = 127
                    _ = v14
                }
            } else {
                var v15 uint = 74
                v16 := []int{fi10, (-(fi10))}
                v16 = append(v16, fi10)
                _ = v15
                _ = v16
            }
            fmt.Println(fi7, v5, v6)
            v17 := []int{fi7}
            v5 = -9223372036854775808
            var v18 int8 = 127
            _ = fi10
            _ = v17
            _ = v18
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return false
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -1
    v3 := []uint32{228, 2, 2147483648, 2147483647}
    var v4 bool = true
    v2 = (v2 + (v2 << ((4294967295) & 63)))
    var v19 int32 = v2
    if (0 != 42) {
        for fi20 := 0; fi20 < 15; fi20++ {
            v2 = (v2 + (v19 % ((-73) | 1)))
            var v21 bool = (74 == 2)
            _ = fi20
            _ = v21
        }
    } else {
        v1[4294967295] = -2147483648
        v22 := GenS2{F0: -156, F1: 998}
        v23 := GenS2{F0: -213, F1: -784}
        _ = v22
        _ = v23
    }
    v4 = (1.5 <= (1.5 / 1e10))
    var v24 int32 = v19
    switch "foo bar" {
    case "a":
        if ((!v4) || (!(v24 > (v2 << ((1) & 63))))) {
            v25 := GenS1{F0: ""}
            v26 := GenS2{F0: 1, F1: -9223372036854775808}
            _ = v25
            _ = v26
        } else {
            v27 := GenS2{F0: 0, F1: -1}
            for fi28 := 0; fi28 < 10; fi28++ {
                v4 = v4
                fmt.Println(v2)
                if ("foo bar" < "hello") {
                    v29 := GenS1{F0: "foo bar"}
                    _ = v29
                } else {
                    v4 = false
                    for fi30 := 0; fi30 < 17; fi30++ {
                        v2 -= (-(v19))
                        _ = fi30
                    }
                    var v31 int32 = (v24 << ((18446744073709551615) & 63))
                    _ = v31
                }
                _ = fi28
            }
            _ = v27
        }
    case "":
        for fi32 := 0; fi32 < 4; fi32++ {
            for fi33 := 0; fi33 < 4; fi33++ {
                v2 += v2
                _ = fi33
            }
            _ = fi32
        }
    case "hello":
        var v34 int32 = v2
        _ = v34
    }
    v3 = append(v3, 162)
    var v35 uint = 42
    fmt.Println(v24)
    v1[-674] = 4294967295
    fmt.Println(v35, v4)
    var v36 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
