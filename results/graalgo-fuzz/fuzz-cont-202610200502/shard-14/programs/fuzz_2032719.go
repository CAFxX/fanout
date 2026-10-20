package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 102
    var v3 float32 = 0.1
    v4 := []string{((("\n\t" + "\n\t") + "hello") + "foo bar")}
    v5 := make(map[int]int)
    v6 := []string{"foo bar", (("" + "\n\t") + ""), "a", "hello"}
    v3 += 1e10
    v2 -= v2
    v1[32768] = 32767
    v2 -= v2
    v2 += (v2 % ((v2) | 1))
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint8 = 19
        _ = fi7
        _ = v8
    }
    v3 = 0.0
    if (!(!(-1 != -16))) {
        v2 = -2147483648
    }
    v3 += v3
    var v9 uint = 18446744073709551615
    if (true && false) {
        fmt.Println(v3)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        v2 += -2147483648
        var v10 uint64 = 1
        _ = v10
    } else {
        switch 128 {
        case 993:
            var v11 float64 = ((0.1 + float64(float32(((((-2.25 + 0.1) + float64(float32(1.5))) + 0.0) - (3.14159 / 1e10))))) - 1.5)
            _ = v11
        case 32767:
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 bool = true
                var v14 uint = (^(v9))
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        v15 := GenS1{F0: 1, F1: 1, F2: -454}
        for fi16 := 0; fi16 < 3; fi16++ {
            v3 = v3
            v1[(fi16 % ((fi16) | 1))] = (fi16 / ((fi16) | 1))
            if false {
                v17 := GenS2{F0: 42}
                _ = v17
            }
            _ = fi16
        }
        _ = v15
    }
    var v18 int = 937
    if true {
        v18 = (v18 / ((v18) | 1))
    } else {
        v2 = (v2 / ((v2) | 1))
        var v19 float32 = 0.0
        fmt.Println(v3)
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
