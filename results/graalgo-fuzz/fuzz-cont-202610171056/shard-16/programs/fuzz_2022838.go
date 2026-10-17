package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint }
type GenS3 struct { F0 string }

func main() {
    v1 := []string{"hello", "", ("foo bar" + (("" + "") + "foo bar"))}
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    var v5 int64 = 42
    fmt.Println(v5)
    var v6 float64 = float64(0.0)
    if (988 <= 2) {
        v6 = v6
        v5 += v5
        v5 += v5
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        v5 = (v5 & v5)
        v8 := GenS1{F0: 1}
        fmt.Println(v6)
        _ = fi7
        _ = v8
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2["foo bar"] = ("\n\t" + ("foo bar" + (("foo bar" + "\n\t") + "hello")))
    if ((589 == 0) && true) {
        v9 := GenS1{F0: -1}
        _ = len(v1)
        _ = v9
    } else {
        v1 = append(v1, ("\n\t" + ("" + "a")))
        v10 := GenS2{F0: 1, F1: 65535, F2: 65535}
        v5 += (v5 & (v5 / ((v5) | 1)))
        _ = v10
    }
    v11 := GenS1{F0: 1}
    if (true && (737 < 0)) {
        v6 = (v6 + v6)
        v6 -= v6
        if false {
            for fi12 := 0; fi12 < 12; fi12++ {
                for fi13 := 0; fi13 < 10; fi13++ {
                    v14 := GenS2{F0: 1, F1: 576, F2: 859}
                    v6 += 0.1
                    var v15 int = 247
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 uint16 = 255
                _ = fi12
                _ = v16
            }
            v17 := GenS1{F0: -1}
            var v18 int = -40
            _ = v17
            _ = v18
        } else {
            fmt.Println(v6, v6)
            var v19 uint = 18446744073709551615
            v20 := GenS3{F0: "a"}
            _ = v19
            _ = v20
        }
    }
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
