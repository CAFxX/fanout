package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
func genFn0() int8 {
    var v4 uint8 = 0
    v5 := GenS1{F0: 842, F1: 1, F2: 0}
    v4 -= (v4 % ((v4) | 1))
    v6 := []int{-9223372036854775808, -2147483648}
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 bool = true
    v2 := make(map[string]int)
    var v3 int8 = 127
    fmt.Println(v3, v3)
    switch "foo bar" {
    case "":
        switch -9223372036854775808 {
        case 0:
            v1 = (v1 && (!(3.14159 == 0.0)))
        default:
            v3 += genFn0()
        }
    case "a":
        for fi7 := 0; fi7 < 17; fi7++ {
            v1 = (1e10 != (0.1 + 0.1))
            v1 = (32767 == 0)
            v3 -= 127
            _ = fi7
        }
    default:
        if false {
            var v8 uint32 = 2147483648
            switch 0 {
            case -317:
                var v9 int32 = -934
                _ = v9
            }
            fmt.Println(v1)
            _ = v8
        } else {
            for fi10 := 0; fi10 < 17; fi10++ {
                if v1 {
                    if v1 {
                        var v11 float32 = (((float32(0.1) - 1.5) / float32(1.5)) / 1.0)
                        var v12 uint = 18446744073709551615
                        _ = v11
                        _ = v12
                    }
                }
                _ = fi10
            }
            v3 -= (^(v3))
        }
    }
    v13 := v2["\n\t"]
    v14 := []int{v13, v13}
    for fi15 := 0; fi15 < 4; fi15++ {
        v13 = fi15
        if v1 {
            v16 := GenS1{F0: 0, F1: (v13 | v13), F2: 42}
            _ = v16
        }
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v1, v1, v3)
        case "":
            v17 := GenS1{F0: 255, F1: v13, F2: 2}
            _ = v17
        }
        _ = fi15
    }
    v14 = append(v14, -95)
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS1{F0: 663, F1: (v13 << ((2) & 63)), F2: 1}
        v2[(("a" + "foo bar") + "a")] = 0
        switch -1 {
        case 32767:
            fmt.Println(v3, v1)
        }
        _ = fi18
        _ = v19
    }
    v20 := v2["foo bar"]
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 uint = 255
        _ = fi21
        _ = v22
    }
    var v23 bool = v1
    v23 = v23
    v20 += (v13 | (v13 - 127))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
