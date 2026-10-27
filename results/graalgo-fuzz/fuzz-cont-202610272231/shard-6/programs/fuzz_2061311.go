package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 string; F2 int }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint8, p2 string) int32 {
    v7 := GenS1{F0: 184}
    v8 := GenS1{F0: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int64 = 0
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 float64 = -2.25
    var v5 float64 = v4
    var v6 uint32 = 2147483647
    var v9 uint64 = 667
    var v10 bool = (v6 == v6)
    v11 := GenS2{F0: -2147483648, F1: "\n\t", F2: 9223372036854775807}
    v5 += 1.5
    for fi12 := 0; fi12 < 20; fi12++ {
        v9 += (v9 ^ 862)
        _ = fi12
    }
    v13 := GenS3{F0: 1000}
    var v14 uint = 1
    v4 += (((((3.14159 / v5) + 1.5) + v5) / ((((0.0 * (((v4 + (0.0 * 0.1)) * v4) - (float64(1e10) * v4))) / v5) / float64(1.5)) - (3.14159 - float64(3.14159)))) + v4)
    var v15 uint8 = 251
    delete(v3, (("\n\t" + ("" + ("hello" + "a"))) + "hello"))
    var v16 int32 = 1
    v17 := GenS1{F0: 2147483648}
    var v18 uint8 = (v15 & uint8(v9))
    switch "hello" {
    case "\n\t":
        switch "hello" {
        case "":
            v19 := []int{127, 0, -74}
            _ = v19
        case "\n\t":
            for fi20 := 0; fi20 < 14; fi20++ {
                switch "a" {
                case "hello":
                    v21 := []int{}
                    _ = v21
                case "a":
                    for fi22 := 0; fi22 < 3; fi22++ {
                        var v23 uint = v14
                        var v24 int32 = genFn0(v18, v15, "a")
                        switch ((("a" + "") + "foo bar") + "a") {
                        case "hello":
                            var v25 int = -818
                            _ = v25
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                default:
                    var v26 uint = uint(v15)
                    _ = v26
                }
                _ = fi20
            }
        }
    case "a":
        var v27 uint16 = 247
        _ = v27
    default:
        delete(v3, "foo bar")
    }
    v18 -= 0
    v9 = 1
    var v28 uint16 = 715
    for fi29 := 0; fi29 < 10; fi29++ {
        delete(v3, "a")
        v30 := []int{(fi29 << ((v14) & 63))}
        var v31 int8 = 127
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
