package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0() int32 {
    for fi4 := 0; fi4 < 11; fi4++ {
        if (3.14159 <= float32((0.0 / 1e10))) {
            var v5 int64 = 0
            switch "hello" {
            case "foo bar":
                fmt.Println(fi4, fi4)
            case "":
                var v6 int8 = 119
                _ = v6
            }
            var v7 bool = (42 != 0)
            _ = v5
            _ = v7
        } else {
            switch "\n\t" {
            case "hello":
                if false {
                    v8 := []int{(-(fi4)), 634, fi4}
                    v8 = append(v8, (fi4 - 0))
                    if (1e10 <= -2.25) {
                        var v9 uint64 = 42
                        _ = v9
                    }
                    _ = v8
                }
            case "":
                for fi10 := 0; fi10 < 20; fi10++ {
                    fmt.Println(fi4, fi10)
                    switch "a" {
                    case "":
                        v11 := []int{}
                        _ = v11
                    }
                    var v12 float64 = float64(0.1)
                    switch fi10 {
                    case 4294967295:
                        fmt.Println(fi10, fi4)
                    default:
                        switch ("hello" + ("\n\t" + "")) {
                        case "a":
                            v12 += v12
                        case "\n\t":
                            var v13 uint16 = 197
                            _ = v13
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                v12 = ((v12 - v12) + ((v12 * 1.5) * (v12 - 1.5)))
                            case "\n\t":
                                if true {
                                    v14 := []int{fi10, fi4}
                                    _ = v14
                                } else {
                                    var v15 uint64 = 4294967295
                                    var v16 uint32 = 647
                                    _ = v15
                                    _ = v16
                                }
                            }
                        }
                    }
                    v12 += v12
                    _ = fi10
                    _ = v12
                }
            case "a":
                fmt.Println(fi4)
            }
            v17 := []int{(fi4 * (fi4 % (((fi4 ^ fi4)) | 1)))}
            v18 := GenS1{F0: 100, F1: 0, F2: -828}
            _ = v17
            _ = v18
        }
        _ = fi4
    }
    v19 := []int{-1, 265, -2147483648, 2147483647}
    var v20 float32 = -2.25
    _ = v19
    _ = v20
    return -1
}

func main() {
    var v1 int64 = 587
    var v2 string = ("foo bar" + (("\n\t" + "a") + ("\n\t" + (("foo bar" + (("hello" + "") + "a")) + ""))))
    v3 := []uint32{415, 2147483648, 4294967295}
    v2 = (v2 + (v2 + "foo bar"))
    v21 := GenS1{F0: 100, F1: 4294967295, F2: 28}
    v22 := GenS1{F0: 127, F1: 1, F2: -1}
    v2 = v2
    v23 := GenS1{F0: 0, F1: 59, F2: 1}
    v24 := GenS2{F0: 78, F1: 0}
    var v25 bool = ((-2.25 - 3.14159) != -2.25)
    switch v2 {
    case "":
        v25 = true
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        for fi26 := 0; fi26 < 10; fi26++ {
            v1 -= v1
            if len(v3) > 0 {
                v3[0] = 2
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi26
        }
    }
    fmt.Println(v2, v1)
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
