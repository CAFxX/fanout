package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []uint32{2147483647, 42, 2147483647}
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v4 int = -2147483648
    var v5 string = "foo bar"
    _ = len(v3)
    v4 = -9223372036854775808
    for fi6 := 0; fi6 < 7; fi6++ {
        switch fi6 {
        case 32767:
            if (-9223372036854775808 != -9223372036854775808) {
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v7 bool = false
                    v7 = false
                    _ = v7
                }
                v5 = (v5 + v5)
                v5 = ((v5 + v5) + "")
            } else {
                var v8 uint32 = 892
                _ = v8
            }
        case 0:
            if true {
                v9 := GenS1{F0: 18446744073709551615}
                _ = v9
            } else {
                switch fi6 {
                case 32768:
                    if (true || (221 >= 255)) {
                        fmt.Println(v5, fi6, v4)
                        v4 = (^(fi6))
                    }
                case 128:
                    var v10 uint32 = 4294967295
                    _ = v10
                default:
                    for fi11 := 0; fi11 < 18; fi11++ {
                        _ = len(v1)
                        _ = fi11
                    }
                }
                var v12 int = v4
                if true {
                    fmt.Println(fi6, v4)
                    _ = len(v3)
                    for fi13 := 0; fi13 < 1; fi13++ {
                        fmt.Println(v12, fi6)
                        _ = fi13
                    }
                } else {
                    v5 = v5
                }
                _ = v12
            }
        case 1:
            v14 := GenS1{F0: 2}
            _ = v14
        default:
            v5 = v5
        }
        _ = fi6
    }
    fmt.Println(v4, v5)
    v5 = "foo bar"
    v4 += (v4 | (v4 + v4))
    var v15 string = "hello"
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
