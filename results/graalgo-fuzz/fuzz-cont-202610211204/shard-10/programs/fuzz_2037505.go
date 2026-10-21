package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0() uint16 {
    var v5 int8 = 114
    v6 := GenS1{F0: "", F1: false}
    v5 = (v5 << ((0) & 63))
    for fi7 := 0; fi7 < 2; fi7++ {
        v5 += -4
        fmt.Println(v5)
        var v8 uint64 = 51
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: "foo bar", F1: (2 > 0)}
    _ = v5
    _ = v6
    _ = v9
    return 255
}

func main() {
    var v1 int16 = 0
    v2 := make(map[string]uint32)
    var v3 uint32 = 0
    var v4 uint8 = 128
    v1 = v1
    if false {
        v10 := []int{-9223372036854775808}
        var v11 bool = (!true)
        _ = v10
        _ = v11
    }
    v3 = v3
    fmt.Println(v4, v4)
    fmt.Println(v1, v3, v1)
    var v12 uint16 = 255
    var v13 uint32 = v3
    switch -2147483648 {
    case 9223372036854775807:
        if false {
            if true {
                if false {
                    v4 = (^(v4))
                } else {
                    v3 = (^(v13))
                    for fi14 := 0; fi14 < 9; fi14++ {
                        fmt.Println(v12)
                        if ((uint64(v12) <= 18446744073709551615) || true) {
                            var v15 int8 = -1
                            _ = v15
                        } else {
                            v1 = 53
                        }
                        _ = fi14
                    }
                }
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    fmt.Println(v3, v3, v1)
                case "hello":
                    v16 := v2["hello"]
                    _ = v16
                default:
                    v1 = (v1 ^ -359)
                }
            } else {
                v1 += (v1 - v1)
            }
            v17 := GenS1{F0: (("" + ("" + ("hello" + (("foo bar" + "foo bar") + ("\n\t" + "hello"))))) + ("hello" + "a")), F1: (true || false)}
            _ = v17
        }
    default:
        v4 = (v4 - v4)
    }
    v18 := []int{-2147483648, 42, 2147483647}
    v12 += (v12 >> ((255) & 63))
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 int = 4294967295
        var v21 float32 = 1e10
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := GenS1{F0: ("" + ("\n\t" + "hello")), F1: true}
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
