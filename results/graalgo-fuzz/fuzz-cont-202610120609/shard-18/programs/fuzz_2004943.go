package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int8; F1 string }
func genFn0(p0 int16, p1 string, p2 int8) int64 {
    if true {
        switch (p1 + "") {
        case "foo bar":
            var v4 float32 = -2.25
            _ = v4
        default:
            p2 = p2
        }
        for fi5 := 0; fi5 < 7; fi5++ {
            for fi6 := 0; fi6 < 6; fi6++ {
                var v7 string = ""
                p2 = p2
                _ = fi6
                _ = v7
            }
            switch p1 {
            case "foo bar":
                fmt.Println(p2, fi5)
            }
            if false {
                var v8 int16 = p0
                switch ("" + ((p1 + p1) + ("hello" + ((p1 + ("a" + "\n\t")) + "foo bar")))) {
                case "hello":
                    p2 += (p2 | 70)
                case "\n\t":
                    p0 = p0
                default:
                    fmt.Println(v8)
                }
                _ = v8
            } else {
                v9 := []int{fi5, fi5, (fi5 >> ((292) & 63))}
                var v10 int16 = p0
                v10 = (p0 / (((p0 & p0)) | 1))
                _ = v10
                _ = v9
            }
            _ = fi5
        }
        v11 := GenS1{F0: 44}
        _ = v11
    }
    fmt.Println(p2, p0)
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 int = fi12
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: 0}
        _ = fi12
        _ = v15
    }
    v16 := GenS2{F0: p1, F1: 60, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 449
}

func main() {
    var v1 float32 = 3.14159
    v2 := []string{"foo bar", "foo bar", ("" + "a")}
    v3 := make(map[uint32]bool)
    if true {
        var v17 float64 = 0.0
        fmt.Println(v1)
        _ = v17
    }
    v18 := GenS1{F0: 42}
    v1 = float32(((1e10 - 0.1) - 1.5))
    v1 -= v1
    v3[65535] = false
    v1 = (v1 - (((((v1 * float32((0.1 * 1.5))) * (v1 * 0.1)) - (((v1 / -2.25) * 0.0) / ((0.1 - (((v1 - v1) * 0.1) + -2.25)) / ((v1 / float32(1e10)) - (0.1 * v1))))) + (v1 + v1)) * 0.0))
    v19 := GenS3{F0: 127, F1: "hello"}
    var v20 float64 = 1e10
    v21 := GenS3{F0: -128, F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
