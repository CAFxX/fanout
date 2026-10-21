package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int64 }

func main() {
    var v1 int16 = -1
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 uint8 = 0
        _ = fi4
        _ = v5
    }
    var v6 float64 = -2.25
    var v7 float64 = v6
    v6 -= float64(0.1)
    var v8 string = (((("a" + "hello") + ("" + (("hello" + "foo bar") + "a"))) + ((("a" + "hello") + "") + "\n\t")) + ("" + (("a" + ("\n\t" + ("\n\t" + (("foo bar" + ("\n\t" + "\n\t")) + "\n\t")))) + "hello")))
    v6 -= -2.25
    v9 := GenS1{F0: 9223372036854775807, F1: 2147483647, F2: -913}
    fmt.Println(v7, v8, v8)
    v10 := GenS1{F0: 127, F1: 1, F2: 42}
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            if (!true) {
                v13 := []int{}
                if true {
                    v14 := v3[fi12]
                    _ = v14
                }
                v13 = append(v13, (fi11 + (fi11 - fi12)))
                _ = v13
            }
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
