package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint; F1 int16; F2 int16 }
func genFn0(p0 string, p1 bool) int16 {
    v7 := GenS1{F0: 76, F1: -2147483648}
    p1 = (262 > 1)
    var v8 int = -9223372036854775808
    v8 = 4294967295
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -32768
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    var v3 uint = 4294967295
    var v4 uint16 = 725
    var v5 int32 = -1
    var v6 string = (("foo bar" + (("\n\t" + "\n\t") + "")) + "")
    for fi9 := 0; fi9 < 7; fi9++ {
        if (true && false) {
            v5 += (v5 & v5)
        }
        _ = fi9
    }
    v10 := GenS1{F0: 440, F1: v5}
    v6 = "a"
    v1[2147483647] = (false && true)
    v11 := GenS2{F0: v3, F1: genFn0("\n\t", ((((0.1 - (-2.25 * 1e10)) - 1.5) + 0.1) >= 3.14159)), F2: 0}
    if true {
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 20; fi13++ {
                fmt.Println(fi13, v4)
                switch ((v6 + "hello") + v6) {
                case "hello":
                    fmt.Println(v4)
                case "foo bar":
                    var v14 float64 = (0.0 - (1e10 + float64(3.14159)))
                    _ = v14
                case "":
                    v4 -= (v4 - v4)
                default:
                    v15 := GenS2{F0: v3, F1: 23, F2: -546}
                    _ = v15
                }
                _ = fi13
            }
            _ = fi12
        }
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := []int{fi16, (fi16 / ((fi16) | 1)), fi16}
        switch fi16 {
        case -880:
            var v18 int16 = 42
            _ = v18
        case 894:
            if (4294967295 >= 65535) {
                fmt.Println(fi16, v6, fi16)
            }
        case 32768:
            fmt.Println(fi16)
        default:
            fmt.Println(v3, v3)
        }
        _ = fi16
        _ = v17
    }
    var v19 uint16 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
