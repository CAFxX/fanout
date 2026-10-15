package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }

func main() {
    var v1 int8 = 100
    var v2 uint = 0
    var v3 float32 = 3.14159
    switch -1 {
    case 4294967295:
        v3 += v3
    }
    for fi4 := 0; fi4 < 12; fi4++ {
        if false {
            v5 := []int{(-(fi4)), fi4, (^(fi4))}
            if true {
                switch "" {
                case "\n\t":
                    if false {
                        for fi6 := 0; fi6 < 18; fi6++ {
                            if (!((float64((v3 * 0.1)) > (-2.25 * float64(v3))) || (!((0.1 > float64(((v3 + -2.25) / v3))) || ((v3 - v3) >= v3))))) {
                                if len(v5) > 0 {
                                    v5[0] = (fi4 - (fi4 << ((v2) & 63)))
                                }
                            }
                            v5 = append(v5, (fi4 * (fi4 - fi4)))
                            _ = fi6
                        }
                    } else {
                        switch "hello" {
                        case "a":
                            v5 = append(v5, (fi4 - -1))
                        case "hello":
                            v7 := GenS1{F0: 1, F1: 2147483647}
                            _ = v7
                        default:
                            v8 := GenS2{F0: 1, F1: 4294967295}
                            _ = v8
                        }
                        fmt.Println(v3, v1, v3)
                    }
                case "hello":
                    var v9 uint8 = 128
                    _ = v9
                case "a":
                    for fi10 := 0; fi10 < 11; fi10++ {
                        v11 := GenS3{F0: v1}
                        _ = fi10
                        _ = v11
                    }
                }
            }
            _ = v5
        } else {
            v3 = 0.0
            var v12 uint32 = 0
            v13 := GenS2{F0: 42, F1: (^(v2))}
            _ = v12
            _ = v13
        }
        v14 := GenS4{F0: "hello"}
        fmt.Println(v1, v2, v1)
        _ = fi4
        _ = v14
    }
    if false {
        fmt.Println(v3, v2)
    }
    v15 := []int{}
    var v16 int32 = 2147483647
    if true {
        fmt.Println(v16, v3)
    }
    v17 := GenS2{F0: v16, F1: (v2 ^ (v2 | v2))}
    var v18 bool = (1 < -576)
    v18 = v18
    for fi19 := 0; fi19 < 10; fi19++ {
        _ = len(v15)
        var v20 uint = v2
        _ = fi19
        _ = v20
    }
    v3 = float32(1e10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
