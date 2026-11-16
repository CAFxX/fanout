package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0() uint8 {
    var v4 string = "foo bar"
    v5 := []int{32768, -702, 2147483647}
    _ = v4
    _ = v5
    return 82
}

func main() {
    v1 := []int{32767}
    var v2 float64 = ((1.5 / 1.0) - -2.25)
    v3 := make(map[int]string)
    if false {
        v6 := GenS1{F0: 42, F1: -1, F2: 128}
        var v7 bool = (42 <= 65535)
        v2 -= ((((v2 - 3.14159) / (v2 * (v2 - 1e10))) * v2) / v2)
        _ = v6
        _ = v7
    } else {
        v2 -= float64(0.0)
    }
    var v8 uint32 = 2147483648
    var v9 float64 = 1e10
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 int64 = 793
            var v13 string = (("" + ("hello" + "a")) + "a")
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := GenS2{F0: 0}
        v2 = (v2 - ((float64((-2.25 / 3.14159)) * float64((0.1 * (-2.25 * ((0.1 + 3.14159) / 1e10))))) / 1.0))
        _ = fi10
        _ = v14
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v8, v2, v8)
    if len(v1) > 0 {
        v1[0] = 127
    }
    var v15 bool = (!(-2147483648 >= -9223372036854775808))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
