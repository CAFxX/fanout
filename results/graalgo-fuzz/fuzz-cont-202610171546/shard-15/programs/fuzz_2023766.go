package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint64 }
type GenS2 struct { F0 uint; F1 bool; F2 int64 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint32; F1 int8; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 644
    v3 := []string{}
    var v4 float32 = (3.14159 + 0.0)
    var v5 uint64 = 4294967295
    for fi6 := 0; fi6 < 2; fi6++ {
        v4 = (float32(((0.1 * (0.1 + (((float64((float32(1.5) - (v4 - 3.14159))) / (float64(float32(1.5)) - ((((0.0 + 1e10) * (0.0 / (1.5 * (0.0 - 1e10)))) / float64(v4)) / (0.1 * (0.1 + float64(v4)))))) * (3.14159 / 3.14159)) / (1.5 + (0.0 + (0.0 / 1.0)))))) / (1e10 + 0.0))) - (((((v4 * 1.5) * float32(((-2.25 - 3.14159) * 0.1))) * v4) - -2.25) - (float32(((1e10 * (0.0 - -2.25)) - 1e10)) / (v4 - v4))))
        var v7 int64 = 1
        _ = fi6
        _ = v7
    }
    fmt.Println(v5, v4, v5)
    if false {
        v8 := GenS1{F0: 216, F1: "\n\t", F2: (v5 | v5)}
        fmt.Println(v5, v5, v4)
        _ = v8
    }
    var v9 int64 = -9223372036854775808
    if (42 < 0) {
        var v10 uint8 = 165
        for fi11 := 0; fi11 < 12; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch ("a" + "") {
                case "\n\t":
                    var v13 uint32 = 786
                    _ = v13
                }
                delete(v1, fi11)
                _ = fi12
            }
            v14 := GenS2{F0: (v2 | (v2 >> ((v2) & 63))), F1: true, F2: v9}
            if (!false) {
                for fi15 := 0; fi15 < 7; fi15++ {
                    var v16 string = (("foo bar" + "") + "foo bar")
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v4, v4, v2)
                    _ = fi15
                    _ = v16
                }
            } else {
                var v17 float32 = ((((v4 - float32(0.0)) / v4) / v4) * v4)
                _ = v17
            }
            _ = fi11
            _ = v14
        }
        v18 := GenS1{F0: (^(v10)), F1: "a", F2: uint64(v2)}
        _ = v10
        _ = v18
    } else {
        v19 := GenS3{F0: -1, F1: (("hello" + ("hello" + "hello")) + "foo bar")}
        var v20 int8 = 0
        var v21 bool = false
        _ = v19
        _ = v20
        _ = v21
    }
    v4 = (float32(3.14159) * (v4 + v4))
    var v22 int32 = 131
    _ = len(v1)
    v23 := GenS4{F0: 0, F1: 100, F2: 4294967295}
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
