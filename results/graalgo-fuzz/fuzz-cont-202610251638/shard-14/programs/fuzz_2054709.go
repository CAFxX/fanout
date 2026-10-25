package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[int]bool)
    var v2 string = ""
    var v3 float64 = 0.0
    v4 := make(map[int]string)
    v5 := []int{-1, 215, -434}
    fmt.Println(v3, v3)
    if true {
        var v6 float32 = 3.14159
        _ = v6
    } else {
        var v7 uint8 = 255
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int = -476
            switch v9 {
            case 999:
                delete(v4, (v9 >> ((642) & 63)))
            case -1:
                for fi10 := 0; fi10 < 9; fi10++ {
                    v11 := GenS1{F0: "\n\t"}
                    if ((2147483647 <= 0) || false) {
                        var v12 uint = 448
                        var v13 int32 = 638
                        fmt.Println(v7)
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: v2}
                    var v15 uint32 = 2
                    v16 := GenS1{F0: v2}
                    _ = fi10
                    _ = v11
                    _ = v14
                    _ = v15
                    _ = v16
                }
            default:
                fmt.Println(fi8, v7, v7)
            }
            var v17 uint8 = 27
            _ = fi8
            _ = v17
            _ = v9
        }
        v3 = v3
        _ = v7
    }
    var v18 uint64 = 0
    v18 = (v18 - 0)
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    var v19 uint64 = 1
    for fi20 := 0; fi20 < 9; fi20++ {
        v5 = append(v5, fi20)
        _ = fi20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 float32 = (0.0 / 1e10)
        if true {
            if false {
                switch v2 {
                case "":
                    switch v2 {
                    case "a":
                        for fi23 := 0; fi23 < 9; fi23++ {
                            fmt.Println(fi21, v18, v22)
                            v2 = v2
                            _ = fi23
                        }
                    }
                case "a":
                    var v24 float64 = v3
                    _ = v24
                }
                v25 := GenS1{F0: "\n\t"}
                delete(v4, fi21)
                _ = v25
            } else {
                v3 = v3
            }
            v26 := GenS1{F0: (v2 + "hello")}
            v27 := GenS1{F0: "\n\t"}
            _ = v26
            _ = v27
        } else {
            fmt.Println(v18)
            if len(v5) > 0 {
                v5[0] = fi21
            }
        }
        v28 := v1[fi21]
        _ = fi21
        _ = v22
        _ = v28
    }
    v3 = float64((0.0 - (-2.25 - 0.1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
