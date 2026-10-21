package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string; F2 uint16 }
func genFn0(p0 bool) uint16 {
    var v4 int = 128
    fmt.Println(v4)
    var v5 string = "hello"
    _ = p0
    _ = v4
    _ = v5
    return 454
}

func main() {
    var v1 int = 128
    v2 := make(map[int]int)
    var v3 float64 = 3.14159
    switch "\n\t" {
    case "a":
        v1 += v1
    case "hello":
        fmt.Println(v3, v3)
    case "":
        v1 += v1
    default:
        var v6 uint16 = 996
        _ = v6
    }
    switch ("hello" + "a") {
    case "foo bar":
        var v7 uint32 = 0
        _ = v7
    case "a":
        v3 -= 0.1
    case "":
        v8 := []int{v1, 2147483647, (v1 << ((238) & 63))}
        _ = v8
    }
    if (false || ((genFn0((18446744073709551615 <= 0)) > 42) && (42 == 0))) {
        v9 := GenS1{F0: 0}
        var v10 int16 = 32767
        _ = v10
        _ = v9
    } else {
        v1 = 127
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 uint8 = 45
            fmt.Println(v1)
            if true {
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 int8 = int8(v12)
                    _ = fi13
                    _ = v14
                }
                var v15 uint8 = v12
                _ = v15
            }
            _ = fi11
            _ = v12
        }
        if (-1 < 1) {
            var v16 bool = ((1000 < 560) || true)
            _ = v16
        }
    }
    _ = len(v2)
    v17 := GenS2{F0: 0, F1: ("hello" + ("\n\t" + "")), F2: 0}
    v1 += v1
    var v18 uint64 = 42
    v19 := GenS1{F0: -11}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
