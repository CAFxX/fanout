package main

import (
    "fmt"
)


func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]int)
    var v3 uint8 = 128
    var v4 uint32 = 305
    v5 := []int{-986, 32768, 32768, 806}
    v6 := []string{}
    _ = len(v6)
    var v7 int32 = 1
    var v8 uint16 = 32767
    var v9 string = "a"
    for fi10 := 0; fi10 < 13; fi10++ {
        _ = len(v2)
        var v11 uint32 = (v4 + (v4 - v4))
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint8 = 128
            var v15 uint16 = 0
            if true {
                fmt.Println(v9, v3, v8)
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        switch fi12 {
        case 225:
            fmt.Println(v8, v8, v8)
        }
        _ = len(v5)
        _ = fi12
    }
    var v16 uint8 = 2
    v17 := v2[-9223372036854775808]
    var v18 string = ((v9 + v9) + (v9 + "a"))
    v16 -= (v16 % (((v16 << ((65535) & 63))) | 1))
    var v19 uint8 = (v3 & v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
