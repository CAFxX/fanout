package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 0.0
    var v3 uint32 = 2147483648
    v4 := []int{42, 1, 1}
    var v5 uint8 = 127
    v5 = 85
    for fi6 := 0; fi6 < 18; fi6++ {
        v5 = v5
        v3 = (v3 & (v3 << ((2) & 63)))
        var v7 float32 = float32((3.14159 * (3.14159 - float64(float32(0.0)))))
        _ = fi6
        _ = v7
    }
    var v8 int32 = 0
    v3 -= (v3 - v3)
    v9 := GenS1{F0: 18446744073709551615, F1: 255, F2: 0}
    switch (("foo bar" + ("a" + ("" + "\n\t"))) + ("\n\t" + "a")) {
    case "":
        if true {
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := GenS1{F0: 42, F1: 127, F2: (fi10 >> ((189) & 63))}
                _ = len(v1)
                var v12 float32 = float32(((3.14159 - 1e10) / (float64((v2 / 0.1)) * 0.0)))
                _ = fi10
                _ = v11
                _ = v12
            }
            if ((!((v2 - float32(((3.14159 - (-2.25 - 1e10)) - -2.25))) < ((1.5 - (v2 + -2.25)) / (-2.25 * (v2 - float32((((-2.25 + 1e10) + 3.14159) + 0.0))))))) || false) {
                v13 := GenS2{F0: int(v8)}
                _ = v13
            }
            var v14 int32 = -713
            _ = v14
        } else {
            switch -163 {
            case 2147483647:
                var v15 uint16 = 65535
                _ = v15
            case 1:
                var v16 string = "hello"
                _ = v16
            case 618:
                var v17 bool = true
                _ = v17
            }
        }
    default:
        v2 = (v2 * v2)
    }
    v18 := GenS2{F0: 42}
    v3 = v3
    if false {
        _ = len(v4)
        v19 := GenS2{F0: 0}
        v2 += v2
        _ = v19
    }
    var v20 float64 = 0.1
    v1[("a" + (("\n\t" + ("hello" + "\n\t")) + ("hello" + ("" + "hello"))))] = 775
    for fi21 := 0; fi21 < 11; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            v3 = 474
            _ = fi22
        }
        var v23 uint64 = 18446744073709551615
        _ = fi21
        _ = v23
    }
    v2 = float32(v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
