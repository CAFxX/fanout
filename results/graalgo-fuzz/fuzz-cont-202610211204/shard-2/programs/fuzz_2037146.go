package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 bool = (42 >= 227)
    var v3 int16 = -32768
    v4 := GenS1{F0: -419, F1: 209, F2: 42}
    v1[(("foo bar" + ("" + ("" + ("a" + (("" + ("hello" + ((("foo bar" + ("\n\t" + "\n\t")) + "\n\t") + ""))) + "hello"))))) + "foo bar")] = -498
    v2 = (!(!true))
    v5 := []int{9223372036854775807, 9223372036854775807, 127}
    for fi6 := 0; fi6 < 18; fi6++ {
        fmt.Println(v3, v2, v2)
        _ = fi6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS2{F0: 0}
        v2 = (1 == 2)
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    if (v2 && v2) {
        switch 1 {
        case -38:
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                v5 = append(v5, -400)
            case "a":
                fmt.Println(v3)
            case "hello":
                v3 += -32768
            default:
                v9 := v1["foo bar"]
                _ = v9
            }
        case -9223372036854775808:
            v3 -= (v3 + 1000)
        default:
            _ = len(v5)
        }
        if false {
            fmt.Println(v3)
        } else {
            fmt.Println(v2, v2)
            switch 32767 {
            case 2147483647:
                for fi10 := 0; fi10 < 19; fi10++ {
                    v11 := GenS3{F0: true}
                    var v12 uint32 = 619
                    var v13 uint8 = 1
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            }
            fmt.Println(v3, v2)
        }
        fmt.Println(v3)
    } else {
        v14 := GenS4{F0: 255, F1: (^(v3))}
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(v3)
        v1["\n\t"] = (fi15 & (fi15 << ((0) & 63)))
        v1["a"] = (fi15 & (fi15 * (fi15 - fi15)))
        _ = fi15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
