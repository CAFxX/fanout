package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int16 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint; F1 int; F2 int64 }

func main() {
    var v1 uint = 923
    v2 := []string{"\n\t", "hello", "foo bar", ("hello" + "\n\t")}
    var v3 uint8 = 0
    var v4 int32 = 980
    if (!false) {
        v3 -= (^(v3))
    }
    var v5 int = 280
    v1 = 0
    _ = len(v2)
    fmt.Println(v3, v5)
    v6 := GenS1{F0: 32767, F1: v5, F2: 0}
    _ = len(v2)
    fmt.Println(v1, v4)
    var v7 uint8 = (v3 * v3)
    v8 := GenS2{F0: 78, F1: "hello"}
    var v9 bool = false
    var v10 string = "foo bar"
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int = v5
        _ = fi11
        _ = v12
    }
    if (v9 || ((v3 & 42) <= v7)) {
        if v9 {
            var v13 int = 0
            _ = v13
        } else {
            v14 := GenS3{F0: "", F1: "a"}
            v15 := GenS4{F0: v1, F1: (v5 % ((-677) | 1)), F2: -621}
            _ = v14
            _ = v15
        }
    } else {
        var v16 float64 = 3.14159
        v1 += (^(v1))
        v2 = append(v2, "hello")
        _ = v16
    }
    v7 += 231
    var v17 uint32 = 497
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
