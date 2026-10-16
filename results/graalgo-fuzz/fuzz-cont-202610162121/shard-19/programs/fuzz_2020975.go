package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint32 }

func main() {
    var v1 int32 = -1
    v2 := make(map[string]int)
    var v3 uint64 = 0
    v4 := GenS1{F0: (true && (204 == -1))}
    var v5 string = ""
    switch 2147483647 {
    case 4294967295:
        v1 = v1
    default:
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint16 = 255
            _ = fi6
            _ = v7
        }
    }
    var v8 int32 = 1
    switch v5 {
    case "hello":
        v5 = v5
    case "\n\t":
        v9 := GenS1{F0: (2 < 128)}
        _ = v9
    }
    switch (v5 + v5) {
    case "hello":
        for fi10 := 0; fi10 < 2; fi10++ {
            v8 = v8
            _ = fi10
        }
    case "":
        var v11 float32 = -2.25
        _ = v11
    default:
        var v12 string = v5
        _ = v12
    }
    v13 := []int{-1}
    if len(v13) > 0 {
        v13[0] = -2147483648
    }
    var v14 float64 = 1e10
    v15 := GenS1{F0: false}
    v13 = append(v13, 9223372036854775807)
    for fi16 := 0; fi16 < 3; fi16++ {
        if false {
            v17 := GenS2{F0: 65535}
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            if (v5 > v5) {
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 int = fi18
                    var v20 int = (-(fi18))
                    fmt.Println(v3)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
            _ = v17
        }
        _ = fi16
    }
    v21 := GenS3{F0: (v5 + "a"), F1: 2147483648}
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v5)
        _ = fi22
    }
    var v23 float64 = v14
    v24 := GenS1{F0: (373 < 255)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
