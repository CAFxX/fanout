package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint; F2 int32 }
func genFn0() int16 {
    var v5 uint = 2
    var v6 int16 = 648
    v6 = v6
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 int64) int16 {
    v7 := []int{-757, 4294967295, 128, -1}
    v8 := GenS1{F0: -128}
    _ = p0
    _ = v7
    _ = v8
    return -627
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{42, 2, 250}
    var v3 uint16 = 0
    v4 := make(map[uint32]bool)
    var v9 int8 = -128
    var v10 float64 = float64((1.5 / 1.0))
    var v11 uint64 = 4294967295
    for fi12 := 0; fi12 < 10; fi12++ {
        v13 := v1[("a" + "")]
        _ = fi12
        _ = v13
    }
    var v14 int16 = -389
    var v15 float32 = (3.14159 + 0.0)
    v16 := GenS1{F0: -62}
    v9 = 127
    _ = len(v2)
    switch "" {
    case "a":
        v17 := GenS1{F0: (v9 & v9)}
        _ = v17
    }
    if true {
        if (!((("" + (("foo bar" + ("hello" + "\n\t")) + "foo bar")) < "a") && false)) {
            var v18 uint16 = 675
            if true {
                v19 := GenS2{F0: (v9 >> ((616) & 63)), F1: 4, F2: -819}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v9, v15, v11)
                _ = v19
            }
            var v20 int64 = 0
            _ = v18
            _ = v20
        }
        v21 := GenS1{F0: (v9 + (v9 - int8(v3)))}
        _ = v21
    }
    v14 = genFn1(321)
    v4[650] = true
    _ = len(v2)
    v22 := GenS2{F0: int8(v11), F1: 42, F2: -1}
    v2 = append(v2, 4294967295)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
