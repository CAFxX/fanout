package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 uint32 }

func main() {
    v1 := []string{"a", ("\n\t" + "\n\t"), ((("foo bar" + "foo bar") + ((("a" + "foo bar") + (("a" + ("" + "a")) + "hello")) + ("\n\t" + "foo bar"))) + ("" + ("a" + "hello"))), ""}
    v2 := make(map[int]bool)
    var v3 int32 = 0
    var v4 float64 = (1e10 / float64((((-2.25 - (1.5 * float32((3.14159 / ((1.5 / 0.1) * 1.5))))) + (float32(1e10) + 3.14159)) / (1e10 / 1.5))))
    var v5 uint = 2
    v3 = (v3 * v3)
    var v6 float64 = v4
    v4 += float64((float32((((0.1 - v6) + v6) / 1.5)) + 3.14159))
    delete(v2, 51)
    for fi7 := 0; fi7 < 16; fi7++ {
        v6 -= float64((3.14159 + float32(((1.5 + 3.14159) * v4))))
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        if ((v5 == v5) && false) {
            for fi9 := 0; fi9 < 9; fi9++ {
                v10 := GenS1{F0: (v3 & -1), F1: v3, F2: true}
                _ = fi9
                _ = v10
            }
        } else {
            var v11 string = "foo bar"
            _ = v11
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 float64 = v4
        _ = fi12
        _ = v13
    }
    fmt.Println(v6, v4)
    v14 := GenS2{F0: 1}
    var v15 int8 = -26
    v15 += (v15 << ((v5) & 63))
    v16 := GenS3{F0: (^(v3)), F1: 4294967295}
    v15 += (^(v15))
    var v17 int32 = v3
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
