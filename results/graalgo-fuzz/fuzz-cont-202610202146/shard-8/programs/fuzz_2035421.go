package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 float64 = -2.25
    if ("a" < ("a" + "foo bar")) {
        v4 := v1[2147483647]
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 int32 = 476
            _ = fi5
            _ = v6
        }
        var v7 string = (("a" + "") + "hello")
        _ = v4
        _ = v7
    } else {
        v8 := GenS1{F0: 1}
        switch "foo bar" {
        case "\n\t":
            v9 := []int{32767, -2147483648}
            _ = v9
        }
        _ = v8
    }
    var v10 string = ""
    switch v10 {
    case "\n\t":
        v10 = ((v10 + (v10 + "")) + (v10 + v10))
    case "a":
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 uint64 = 42
            _ = fi11
            _ = v12
        }
    }
    if true {
        switch (v10 + "\n\t") {
        case "foo bar":
            v3 -= float64(1e10)
        default:
            v3 -= (v3 + v3)
        }
        v13 := []int{-1, 32767, 128, 257}
        _ = v13
    }
    var v14 float32 = -2.25
    v15 := GenS1{F0: -271}
    v3 = float64(v14)
    for fi16 := 0; fi16 < 18; fi16++ {
        v14 = v14
        v17 := []int{42}
        _ = fi16
        _ = v17
    }
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
