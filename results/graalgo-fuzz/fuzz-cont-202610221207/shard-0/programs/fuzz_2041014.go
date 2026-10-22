package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int32; F1 int64 }

func main() {
    var v1 bool = (-2147483648 == 385)
    v2 := []string{""}
    var v3 bool = v1
    var v4 int32 = -1
    v5 := []uint32{}
    var v6 string = "a"
    v7 := GenS1{F0: 1, F1: 738}
    if true {
        v4 += 1
        v1 = false
        if v3 {
            switch "a" {
            case "\n\t":
                var v8 float32 = 3.14159
                _ = v8
            }
            if true {
                v9 := GenS2{F0: v4, F1: -1}
                _ = v9
            } else {
                fmt.Println(v3)
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                if len(v5) > 0 {
                    v5[0] = 1
                }
                _ = fi10
            }
        } else {
            switch ("a" + (v6 + v6)) {
            case "hello":
                for fi11 := 0; fi11 < 11; fi11++ {
                    fmt.Println(v4, v1, v3)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi11
                }
            case "":
                if false {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    fmt.Println(v4)
                    if v3 {
                        v12 := GenS2{F0: 219, F1: 777}
                        _ = len(v2)
                        _ = v12
                    } else {
                        fmt.Println(v3, v1, v6)
                        for fi13 := 0; fi13 < 11; fi13++ {
                            v14 := GenS2{F0: 2147483647, F1: 9223372036854775807}
                            _ = fi13
                            _ = v14
                        }
                        var v15 uint16 = 0
                        _ = v15
                    }
                    v16 := GenS2{F0: v4, F1: -1}
                    _ = v16
                }
            }
        }
    }
    var v17 uint8 = 202
    var v18 float64 = 3.14159
    if false {
        if len(v5) > 0 {
            v5[0] = 1
        }
        switch v6 {
        case "\n\t":
            v6 = ((((v6 + "hello") + v6) + (v6 + v6)) + v6)
        case "":
            var v19 float32 = -2.25
            _ = v19
        default:
            var v20 int = 128
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi21
    }
    var v22 int64 = 768
    fmt.Println(v6)
    v2 = append(v2, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
