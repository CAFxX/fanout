package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 bool }
func genFn0(p0 int32, p1 int) int {
    var v7 float32 = (0.1 + 1.5)
    p1 = 581
    var v8 int8 = -12
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p1 * p1)
}
func genFn1(p0 int) bool {
    if true {
        var v9 int32 = -2147483648
        v10 := []int{genFn0(v9, (p0 % ((p0) | 1))), (p0 ^ (p0 | p0))}
        _ = v10
        _ = v9
    }
    var v11 int32 = -1
    v12 := GenS1{F0: 1}
    _ = p0
    _ = v11
    _ = v12
    return true
}

func main() {
    var v1 int64 = 168
    var v2 int = 32768
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v5 float32 = 0.0
    v6 := make(map[string]int)
    v5 -= v5
    v5 -= v5
    v13 := []int{(v2 ^ 4294967295), v2}
    v5 -= (v5 / (1.5 - ((3.14159 + float32(float64(0.0))) / float32(1.5))))
    v3[v2] = genFn0(1, -9223372036854775808)
    v1 += v1
    switch (("hello" + "a") + ("foo bar" + "hello")) {
    case "\n\t":
        if false {
            var v14 string = ("a" + "foo bar")
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 int8 = 106
                _ = len(v13)
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 17; fi17++ {
                if (!false) {
                    var v18 uint = 280
                    _ = v18
                }
                for fi19 := 0; fi19 < 13; fi19++ {
                    switch "hello" {
                    case "":
                        fmt.Println(v5)
                    }
                    v2 += v2
                    _ = fi19
                }
                v20 := GenS2{F0: 20, F1: (4294967295 >= 2147483647)}
                _ = fi17
                _ = v20
            }
            _ = v14
        } else {
            v21 := GenS1{F0: (v1 + (^(v1)))}
            _ = v21
        }
    case "":
        for fi22 := 0; fi22 < 19; fi22++ {
            for fi23 := 0; fi23 < 18; fi23++ {
                v5 -= v5
                v5 += (((v5 - (v5 * v5)) / ((-2.25 + (1.5 / 3.14159)) + v5)) * 1e10)
                _ = fi23
            }
            if (!(1.5 != (1e10 + 0.1))) {
                switch ("foo bar" + "foo bar") {
                case "hello":
                    var v24 int64 = v1
                    _ = v24
                case "\n\t":
                    var v25 uint = 792
                    _ = v25
                default:
                    v2 += -1
                }
                var v26 uint8 = 2
                _ = v26
            }
            v5 = ((v5 + (v5 - float32(0.1))) * v5)
            _ = fi22
        }
    case "hello":
        for fi27 := 0; fi27 < 14; fi27++ {
            v1 -= v1
            _ = fi27
        }
    }
    fmt.Println(v2, v1, v1)
    var v28 int = genFn0(-1, 127)
    var v29 uint64 = 1
    fmt.Println(v5)
    var v30 uint = 792
    var v31 bool = true
    v32 := v3[v2]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
