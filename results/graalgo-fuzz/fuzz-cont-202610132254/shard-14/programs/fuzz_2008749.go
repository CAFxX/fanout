package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }

func main() {
    v1 := []uint32{122, 42}
    v2 := make(map[int]bool)
    var v3 uint = 415
    v4 := make(map[int]string)
    var v5 uint32 = 42
    var v6 float32 = float32(1.5)
    v3 = 255
    v5 = v5
    v7 := GenS1{F0: "\n\t", F1: (("\n\t" + ("a" + ("a" + "a"))) + ((("foo bar" + "a") + ("foo bar" + (("hello" + "\n\t") + ("\n\t" + "\n\t")))) + "hello"))}
    var v8 int8 = 100
    fmt.Println(v8, v8, v6)
    var v9 string = "a"
    v3 = (^(v3))
    v6 -= 0.1
    v3 = v3
    var v10 int = -505
    if ((1.5 + 1.5) == 1e10) {
        if false {
            if true {
                fmt.Println(v5, v5, v3)
                fmt.Println(v5, v10, v6)
                v11 := GenS1{F0: ("a" + v9), F1: v9}
                _ = v11
            } else {
                if (!true) {
                    v8 -= 100
                    delete(v2, -545)
                    for fi12 := 0; fi12 < 3; fi12++ {
                        var v13 float64 = (1e10 / 1e10)
                        _ = fi12
                        _ = v13
                    }
                } else {
                    delete(v4, (v10 * v10))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v10 -= (-(v10))
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                v3 = (v3 + 0)
                var v15 bool = true
                if (!v15) {
                    v16 := GenS1{F0: (v9 + "a"), F1: v9}
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            var v17 uint64 = 1
            fmt.Println(v9, v9, v6)
            _ = v17
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
