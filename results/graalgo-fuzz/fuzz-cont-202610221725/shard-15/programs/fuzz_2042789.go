package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 614
    var v3 uint8 = 2
    var v4 int = 190
    var v5 float32 = (1.5 / -2.25)
    v3 -= 42
    v6 := GenS1{F0: (("foo bar" + "hello") + "\n\t")}
    fmt.Println(v4, v5)
    v5 -= v5
    fmt.Println(v5, v2, v5)
    if (731 != -9223372036854775808) {
        fmt.Println(v2, v2)
    } else {
        v3 += v3
        var v7 uint = 65535
        _ = v7
    }
    v5 += (float32(1.5) - v5)
    v5 = -2.25
    v8 := GenS1{F0: "\n\t"}
    for fi9 := 0; fi9 < 1; fi9++ {
        v3 -= (v3 ^ 0)
        switch "hello" {
        case "a":
            for fi10 := 0; fi10 < 5; fi10++ {
                v3 = v3
                if (false && false) {
                    switch (v4 << ((42) & 63)) {
                    case -851:
                        var v11 uint8 = v3
                        _ = v11
                    case -70:
                        v12 := []int{(^(v4)), (fi9 % (((^(v4))) | 1)), fi10}
                        _ = v12
                    case 32768:
                        fmt.Println(v5, v4, v4)
                    default:
                        var v13 uint32 = 2147483648
                        _ = v13
                    }
                    switch 0 {
                    case 32768:
                        v14 := v1[("hello" + ("a" + "\n\t"))]
                        _ = v14
                    case 1:
                        v15 := GenS2{F0: (fi9 / (((fi10 ^ 1)) | 1)), F1: 32767}
                        _ = v15
                    default:
                        var v16 uint64 = (^(v2))
                        fmt.Println(v2, v3)
                        _ = v16
                    }
                    if true {
                        v2 = (^(v2))
                        fmt.Println(v5, v5, v3)
                    }
                    v5 -= (v5 - v5)
                }
                v4 = v4
                _ = fi10
            }
        case "":
            for fi17 := 0; fi17 < 8; fi17++ {
                v18 := []int{(v4 | 1), fi9, (fi9 & fi17)}
                _ = fi17
                _ = v18
            }
        }
        v5 += (v5 / (float32(1e10) / v5))
        _ = fi9
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 uint64 = (v2 & v2)
        _ = fi19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
