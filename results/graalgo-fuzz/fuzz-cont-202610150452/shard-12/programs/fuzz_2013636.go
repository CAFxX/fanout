package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int }
type GenS2 struct { F0 int8; F1 bool }
func genFn0(p0 int) int32 {
    v6 := []int{p0}
    for fi7 := 0; fi7 < 2; fi7++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(p0, fi7)
        v6 = append(v6, (p0 - 359))
        _ = fi7
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v6
    return -2147483648
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]string)
    var v3 int16 = -276
    v4 := make(map[int]int)
    var v5 string = "\n\t"
    switch -200 {
    case 42:
        v5 = v5
    case 9223372036854775807:
        var v8 int8 = 127
        _ = v8
    default:
        v5 = v5
    }
    v9 := GenS1{F0: 2, F1: 32768, F2: 9223372036854775807}
    switch ((v5 + ((((v5 + v5) + (v5 + ("a" + ""))) + "a") + (("\n\t" + "") + (v5 + v5)))) + v5) {
    case "\n\t":
        v2[771] = ("" + v5)
    case "":
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 float64 = 3.14159
            if (-9223372036854775808 <= 32767) {
                v5 = "\n\t"
                if (1 != -42) {
                    v11 -= ((v11 + v11) / 0.1)
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v13 := []int{4294967295, 297}
                        v1 += 255
                        _ = fi12
                        _ = v13
                    }
                    v14 := GenS2{F0: -42, F1: (!true)}
                    _ = v14
                }
            } else {
                v1 -= (^(v1))
                var v15 uint = 18446744073709551615
                _ = v15
            }
            if (!false) {
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 int16 = v3
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi10
            _ = v11
        }
    case "a":
        var v18 int16 = (v3 ^ -346)
        _ = v18
    }
    var v19 uint = uint(v1)
    var v20 string = (("" + "\n\t") + v5)
    switch "\n\t" {
    case "\n\t":
        var v21 uint8 = 1
        _ = v21
    }
    v22 := GenS1{F0: v1, F1: 849, F2: 127}
    v1 -= 605
    v23 := []int{-2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
