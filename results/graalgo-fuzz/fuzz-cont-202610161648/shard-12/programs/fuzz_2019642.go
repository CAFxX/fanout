package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS3 struct { F0 int16 }
func genFn0() int32 {
    if true {
        var v6 bool = false
        if false {
            for fi7 := 0; fi7 < 7; fi7++ {
                switch "foo bar" {
                case "hello":
                    var v8 string = (((("foo bar" + "hello") + (("\n\t" + "\n\t") + "a")) + ("\n\t" + ("\n\t" + (("" + ("a" + "\n\t")) + ((("" + ((("foo bar" + (("a" + "hello") + "\n\t")) + "foo bar") + "hello")) + (("" + "\n\t") + "\n\t")) + "foo bar"))))) + "a")
                    _ = v8
                }
                v9 := []int{-1, (-(fi7))}
                _ = fi7
                _ = v9
            }
            var v10 bool = (v6 || (v6 || (!(!v6))))
            v11 := GenS1{F0: 0, F1: 255, F2: 42}
            _ = v10
            _ = v11
        } else {
            if (0 < 1) {
                for fi12 := 0; fi12 < 15; fi12++ {
                    v13 := GenS2{F0: 127, F1: 1, F2: 2}
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v6 = true
                        var v15 bool = v6
                        _ = fi14
                        _ = v15
                    }
                    fmt.Println(fi12, v6)
                    _ = fi12
                    _ = v13
                }
            } else {
                v6 = v6
                if v6 {
                    v16 := GenS3{F0: 1000}
                    var v17 uint = 1
                    _ = v16
                    _ = v17
                } else {
                    var v18 float32 = (3.14159 * (-2.25 + 1.5))
                    fmt.Println(v18, v18)
                    _ = v18
                }
                v19 := []int{-1}
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := []int{127, (fi20 >> ((75) & 63))}
            v6 = true
            _ = fi20
            _ = v21
        }
        _ = v6
    }
    v22 := []int{124, 4294967295, 32767, 2147483647}
    v23 := GenS2{F0: 105, F1: 42, F2: 794}
    for fi24 := 0; fi24 < 17; fi24++ {
        fmt.Println(fi24)
        var v25 uint = 208
        v25 -= 4294967295
        _ = fi24
        _ = v25
    }
    _ = v22
    _ = v23
    return -351
}
func genFn1(p0 int, p1 int64, p2 int16) int16 {
    var v26 float32 = 1e10
    fmt.Println(v26, p1, p2)
    v27 := []int{32767, -410}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return p2
}

func main() {
    var v1 float64 = 1.5
    v2 := []int{}
    var v3 uint16 = 2
    var v4 uint32 = 203
    var v5 uint = 42
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 uint32 = v4
        _ = fi28
        _ = v29
    }
    v1 -= float64(-2.25)
    var v30 float64 = (((float64(-2.25) * v1) / v1) * v1)
    fmt.Println(v4)
    v5 += (v5 << (((v5 * (v5 / (((v5 - (v5 & 1))) | 1)))) & 63))
    var v31 uint64 = 18446744073709551615
    fmt.Println(v3, v4)
    var v32 uint8 = 61
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
