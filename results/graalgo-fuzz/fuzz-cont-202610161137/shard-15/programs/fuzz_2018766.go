package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS2 struct { F0 bool; F1 int16 }

func main() {
    var v1 uint8 = 58
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    var v4 uint16 = 204
    fmt.Println(v4, v4)
    v5 := GenS1{F0: v1, F1: false, F2: 0}
    for fi6 := 0; fi6 < 1; fi6++ {
        v1 -= v1
        v7 := GenS1{F0: v1, F1: ((3.14159 < 1.5) || true), F2: (v1 % ((uint8(v4)) | 1))}
        _ = fi6
        _ = v7
    }
    switch -76 {
    case 32768:
        v1 -= v1
    }
    v8 := GenS2{F0: true, F1: 536}
    var v9 uint8 = v1
    fmt.Println(v9, v9)
    switch (("" + "foo bar") + "\n\t") {
    case "":
        for fi10 := 0; fi10 < 18; fi10++ {
            for fi11 := 0; fi11 < 11; fi11++ {
                if (1 >= (fi11 - (-(fi10)))) {
                    v9 = v9
                    var v12 int16 = -1
                    if false {
                        if true {
                            v9 = (v9 % (((v9 >> ((2) & 63))) | 1))
                            v13 := []int{fi11}
                            _ = v13
                        }
                    }
                    v4 = (^(v4))
                    _ = v12
                }
                _ = fi11
            }
            _ = fi10
        }
    case "hello":
        fmt.Println(v1, v4, v4)
    }
    fmt.Println(v4, v4, v4)
    var v14 int = 1
    v15 := []int{441, 32768, (v14 >> ((1) & 63))}
    var v16 string = ""
    var v17 uint32 = 502
    switch v14 {
    case 2147483647:
        var v18 bool = (!true)
        _ = v18
    default:
        fmt.Println(v9, v1, v9)
    }
    var v19 bool = false
    fmt.Println(v4, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
