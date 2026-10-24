package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 int8 = -33
    var v2 int32 = 0
    var v3 int16 = 1
    if (0 != 9223372036854775807) {
        fmt.Println(v3)
        v2 -= 2147483647
    }
    v4 := GenS1{F0: 42}
    var v5 string = "hello"
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int32 = v2
        v8 := []int{}
        for fi9 := 0; fi9 < 3; fi9++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v2)
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            switch 9223372036854775807 {
            case 2147483647:
                if (((2 == 38) && true) && false) {
                    var v12 uint8 = 0
                    v13 := []int{(fi10 ^ (fi11 | fi11))}
                    _ = v12
                    _ = v13
                } else {
                    switch -1 {
                    case -2147483648:
                        v14 := GenS1{F0: 4294967295}
                        _ = v14
                    case -1:
                        fmt.Println(v5)
                    case 9223372036854775807:
                        switch fi11 {
                        case 4294967295:
                            var v15 uint16 = 0
                            _ = v15
                        default:
                            v3 += int16(v1)
                        }
                    default:
                        v16 := GenS2{F0: 128, F1: 2147483647}
                        _ = v16
                    }
                }
            case -200:
                v5 = ((v5 + v5) + v5)
            case -788:
                var v17 bool = true
                _ = v17
            }
            _ = fi11
        }
        _ = fi10
    }
    v18 := []int{128, -116, 874, -9223372036854775808}
    for fi19 := 0; fi19 < 14; fi19++ {
        v18 = append(v18, 814)
        v20 := GenS1{F0: 4294967295}
        v3 = v3
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
