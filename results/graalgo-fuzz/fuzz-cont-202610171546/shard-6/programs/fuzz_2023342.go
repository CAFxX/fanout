package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int8 }
type GenS2 struct { F0 int16 }
func genFn0() int32 {
    switch ("\n\t" + "\n\t") {
    case "hello":
        if (!false) {
            v6 := []int{9223372036854775807}
            var v7 string = (("a" + ("hello" + "\n\t")) + "")
            v8 := GenS1{F0: 0, F1: 0, F2: -52}
            _ = v6
            _ = v7
            _ = v8
        } else {
            var v9 uint8 = 128
            v10 := GenS1{F0: 890, F1: -1, F2: -1}
            v9 = v9
            _ = v10
            _ = v9
        }
    default:
        if false {
            if false {
                var v11 uint = 42
                v12 := GenS2{F0: -85}
                v13 := []int{}
                _ = v11
                _ = v12
                _ = v13
            } else {
                var v14 uint8 = 2
                fmt.Println(v14)
                _ = v14
            }
        }
    }
    var v15 int16 = 0
    fmt.Println(v15)
    _ = v15
    return 1
}
func genFn1(p0 int8) int8 {
    p0 += 42
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := GenS2{F0: -32768}
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    _ = p0
    return (p0 - p0)
}

func main() {
    var v1 uint8 = 2
    var v2 float64 = 3.14159
    v3 := []string{}
    var v4 uint64 = 706
    var v5 uint16 = 255
    var v19 uint64 = (v4 + (v4 + 1))
    fmt.Println(v4, v2, v1)
    var v20 uint64 = v19
    fmt.Println(v1, v2)
    var v21 uint = 65535
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v5 -= v5
    v22 := GenS1{F0: 437, F1: -1, F2: 51}
    _ = len(v3)
    for fi23 := 0; fi23 < 13; fi23++ {
        _ = len(v3)
        v2 = (((0.0 - float64(1e10)) * v2) / v2)
        _ = fi23
    }
    v5 = (v5 << (((v21 * v21)) & 63))
    if true {
        if (2147483647 < -900) {
            if false {
                v20 -= (v19 >> ((uint(v20)) & 63))
                switch ("\n\t" + (("" + ((("hello" + "a") + "\n\t") + ("" + "\n\t"))) + "\n\t")) {
                case "foo bar":
                    v3 = append(v3, "")
                case "hello":
                    v3 = append(v3, "a")
                default:
                    for fi24 := 0; fi24 < 20; fi24++ {
                        v3 = append(v3, ("a" + "hello"))
                        switch 42 {
                        case -310:
                            v3 = append(v3, "foo bar")
                        }
                        _ = fi24
                    }
                }
            } else {
                v25 := GenS1{F0: 925, F1: -2147483648, F2: -128}
                fmt.Println(v19, v20, v20)
                _ = v25
            }
            fmt.Println(v21, v4, v1)
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        }
    } else {
        v20 -= (v4 ^ (^(v20)))
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v26 int64 = -518
    var v27 float32 = (1e10 - 3.14159)
    fmt.Println(v27, v5, v19)
    var v28 bool = false
    v26 -= (v26 ^ (v26 % ((0) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
