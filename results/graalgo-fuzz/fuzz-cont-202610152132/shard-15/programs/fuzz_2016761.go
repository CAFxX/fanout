package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint; F2 string }
func genFn0() int8 {
    v4 := []int{}
    var v5 uint = 0
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(v5)
        v5 = v5
        _ = fi6
    }
    v5 += (v5 * v5)
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 uint = 255
    var v2 uint32 = 2
    v3 := make(map[string]uint32)
    var v7 float64 = 1.5
    v3["\n\t"] = v2
    var v8 uint = (^(v1))
    switch 4294967295 {
    case 128:
        fmt.Println(v7, v7, v1)
    case -756:
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint32 = (v2 << (((v8 | (v1 % ((18446744073709551615) | 1)))) & 63))
            v11 := GenS1{F0: 0}
            v12 := GenS1{F0: 0}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    default:
        if false {
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := GenS2{F0: 42, F1: (^(v1)), F2: "hello"}
                if true {
                    switch ("foo bar" + "a") {
                    case "":
                        var v15 int8 = 127
                        _ = v15
                    }
                } else {
                    v8 = v8
                    if true {
                        switch ("" + "") {
                        case "a":
                            var v16 int16 = -32768
                            _ = v16
                        }
                    }
                }
                _ = fi13
                _ = v14
            }
            var v17 string = "a"
            delete(v3, v17)
            _ = v17
        }
    }
    switch 0 {
    case 0:
        fmt.Println(v7)
    case 127:
        v18 := []int{}
        _ = v18
    default:
        if (((!false) && false) || (((((0.1 + ((0.0 + (0.0 / (v7 - v7))) - v7)) * 3.14159) / v7) - v7) <= float64(1.5))) {
            switch "foo bar" {
            case "a":
                var v19 int8 = 42
                _ = v19
            }
        }
    }
    fmt.Println(v2)
    v3[(("a" + "") + "\n\t")] = uint32(v8)
    var v20 int8 = 42
    v8 += v1
    fmt.Println(v7, v2, v20)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v8
}
