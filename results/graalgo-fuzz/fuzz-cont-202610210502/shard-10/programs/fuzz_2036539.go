package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }

func main() {
    v1 := []int{4294967295, 128}
    v2 := make(map[int]string)
    var v3 int32 = -2147483648
    v4 := make(map[uint32]string)
    var v5 int8 = 100
    var v6 int8 = v5
    v7 := v4[2]
    if (341 >= 42) {
        switch v7 {
        case "a":
            if ((168 != 1) && true) {
                v7 = (v7 + (v7 + "a"))
                if (v3 != v3) {
                    v5 -= 1
                    switch v7 {
                    case "foo bar":
                        v7 = "foo bar"
                    case "\n\t":
                        fmt.Println(v3)
                    default:
                        v7 = "\n\t"
                    }
                } else {
                    var v8 int16 = 32767
                    var v9 uint = 255
                    for fi10 := 0; fi10 < 12; fi10++ {
                        var v11 int64 = 92
                        fmt.Println(v3)
                        var v12 bool = true
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = v8
                    _ = v9
                }
            }
        case "":
            v5 += v5
        default:
            var v13 string = v7
            _ = v13
        }
        v14 := v4[1]
        _ = v14
    } else {
        var v15 int8 = (v5 >> ((4294967295) & 63))
        _ = v15
    }
    v16 := GenS1{F0: 128, F1: v6}
    var v17 bool = false
    v1 = append(v1, 32767)
    if true {
        var v18 string = ((("" + v7) + v7) + v7)
        _ = v18
    } else {
        v2[2147483647] = ("\n\t" + (v7 + ((v7 + v7) + (("a" + v7) + v7))))
        v7 = (v7 + "")
        fmt.Println(v7, v7, v5)
    }
    v19 := GenS1{F0: 42, F1: v5}
    var v20 int8 = 95
    var v21 uint = 255
    v17 = v17
    fmt.Println(v21, v20, v6)
    v22 := GenS1{F0: 3, F1: v6}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
