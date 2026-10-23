package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint }

func main() {
    var v1 float32 = 0.1
    var v2 int32 = -2147483648
    v3 := make(map[string]bool)
    fmt.Println(v2, v2)
    for fi4 := 0; fi4 < 12; fi4++ {
        if (!true) {
            v5 := GenS1{F0: 2, F1: "foo bar", F2: 42}
            _ = v5
        } else {
            var v6 uint64 = 334
            if (-1 < -455) {
                switch ("" + "a") {
                case "foo bar":
                    v7 := GenS1{F0: 255, F1: "\n\t", F2: 42}
                    _ = v7
                case "a":
                    v6 = 0
                case "\n\t":
                    v8 := v3[""]
                    _ = v8
                default:
                    var v9 float64 = (float64((float32((1.5 / 3.14159)) + ((float32(0.0) - v1) - (float32(3.14159) - v1)))) * float64(v1))
                    _ = v9
                }
                for fi10 := 0; fi10 < 5; fi10++ {
                    switch ((((("a" + "") + "foo bar") + (("foo bar" + "foo bar") + "a")) + "a") + "hello") {
                    case "a":
                        v11 := GenS1{F0: 812, F1: "a", F2: 18446744073709551615}
                        _ = v11
                    case "":
                        v6 = v6
                    default:
                        var v12 bool = true
                        _ = v12
                    }
                    _ = fi10
                }
            } else {
                for fi13 := 0; fi13 < 9; fi13++ {
                    var v14 uint8 = 33
                    v1 = float32(0.1)
                    var v15 int8 = -128
                    v14 += v14
                    var v16 float64 = (0.1 + 3.14159)
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
            v1 -= -2.25
            _ = v6
        }
        v2 -= (v2 >> ((83) & 63))
        v17 := []int{0}
        _ = fi4
        _ = v17
    }
    if true {
        v1 -= (float32(0.1) / 1.0)
        switch 32768 {
        case 32767:
            if (true && (!false)) {
                var v18 int = 127
                var v19 uint64 = 715
                _ = v18
                _ = v19
            } else {
                fmt.Println(v2, v2)
            }
        case -9223372036854775808:
            v20 := []int{1}
            _ = v20
        }
        v1 += v1
    } else {
        v21 := GenS1{F0: 18446744073709551615, F1: ("a" + "\n\t"), F2: 18446744073709551615}
        var v22 uint = 4294967295
        var v23 int64 = 42
        _ = v21
        _ = v22
        _ = v23
    }
    switch 32768 {
    case -2147483648:
        var v24 int8 = -1
        _ = v24
    case 4294967295:
        fmt.Println(v2)
    case 844:
        v25 := GenS1{F0: 65535, F1: "", F2: 668}
        _ = v25
    default:
        fmt.Println(v2, v2)
    }
    v2 = v2
    fmt.Println(v1, v1)
    var v26 uint32 = 423
    if false {
        fmt.Println(v2)
    }
    v3["foo bar"] = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
}
