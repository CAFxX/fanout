package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint) int {
    var v4 int = 32767
    v4 += (v4 & -2147483648)
    v4 = v4
    fmt.Println(p0, p0)
    v5 := GenS1{F0: 2147483647}
    _ = p0
    _ = v4
    _ = v5
    return v4
}

func main() {
    var v1 uint = 0
    var v2 float64 = 0.0
    v3 := []uint32{35, 884, 765}
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint64 = 0
        _ = fi6
        _ = v7
    }
    var v8 bool = ((v2 * (v2 - (v2 * 1.5))) == (((1e10 - v2) - v2) * (v2 + 1e10)))
    if (!false) {
        if v8 {
            _ = len(v3)
            var v9 string = (("" + "hello") + "a")
            _ = v9
        }
        var v10 uint16 = 0
        v3 = append(v3, 4294967295)
        _ = v10
    }
    switch 9223372036854775807 {
    case -1:
        v2 = float64(1.5)
    case 781:
        v2 += (1.5 * v2)
    }
    v1 += v1
    v8 = (923 == 18446744073709551615)
    var v11 uint8 = 128
    v12 := GenS1{F0: 42}
    v8 = ((v8 || (((0 == 974) || v8) && (!(1e10 == 3.14159)))) || v8)
    v3 = append(v3, 42)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
}
