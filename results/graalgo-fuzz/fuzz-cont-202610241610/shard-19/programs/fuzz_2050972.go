package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 bool, p1 uint) uint16 {
    v4 := GenS1{F0: "a"}
    if p0 {
        v5 := []int{-253, 42}
        _ = v5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 uint = p1
            fmt.Println(v8, fi6, fi6)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := []int{4294967295, 0}
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return 0
}
func genFn1() uint16 {
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 string = ("" + ("a" + ("foo bar" + ("foo bar" + "\n\t"))))
        var v12 bool = true
        v13 := []int{fi10, (fi10 - fi10), (fi10 % ((-77) | 1)), (fi10 + (-(fi10)))}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    if false {
        v14 := GenS2{F0: ("\n\t" + "hello"), F1: -34}
        _ = v14
    }
    switch "a" {
    case "a":
        var v15 bool = true
        _ = v15
    }
    var v16 uint64 = 42
    _ = v16
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 1e10
    v3 := make(map[string]bool)
    v17 := GenS3{F0: 0}
    var v18 int64 = -902
    var v19 float32 = (3.14159 * float32(0.1))
    var v20 uint64 = 18446744073709551615
    for fi21 := 0; fi21 < 17; fi21++ {
        v2 = 0.0
        if (!((!false) || (1e10 == float64(1.5)))) {
            v18 += 0
            var v22 float32 = ((v19 - (v19 + -2.25)) * v19)
            _ = v22
        }
        _ = fi21
    }
    var v23 int = 127
    delete(v3, "\n\t")
    for fi24 := 0; fi24 < 13; fi24++ {
        v23 += -9223372036854775808
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
