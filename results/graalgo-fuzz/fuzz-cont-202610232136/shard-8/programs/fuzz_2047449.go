package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 float64 = -2.25
    var v4 int64 = -371
    if (v4 == 1) {
        v5 := GenS1{F0: 13}
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := GenS2{F0: fi6}
            _ = fi6
            _ = v7
        }
        var v8 int8 = 0
        _ = v5
        _ = v8
    } else {
        var v9 int = 127
        _ = v9
    }
    switch -174 {
    case 861:
        v3 -= v3
    }
    if false {
        var v10 string = ""
        _ = v10
    }
    var v11 int64 = v4
    for fi12 := 0; fi12 < 12; fi12++ {
        v4 = (v4 & (^(v4)))
        v13 := []int{fi12}
        var v14 float32 = 3.14159
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 bool = false
    v15 = (!(v15 || (!(0 != -1))))
    v16 := GenS3{F0: 0, F1: (("foo bar" + "\n\t") == "a")}
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
