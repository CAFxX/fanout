package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    var v3 int64 = 42
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: 65535}
        v3 -= v3
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: 82}
            v8 := GenS2{F0: 1}
            v9 := []int{(fi6 & fi4), (fi6 % (((fi6 - 1)) | 1)), (fi4 & fi6)}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        _ = fi4
        _ = v5
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int16 = 390
        _ = fi10
        _ = v11
    }
    delete(v1, 302)
    var v12 int = -2147483648
    if true {
        v3 -= (v3 | (v3 + v3))
        fmt.Println(v3, v3)
    }
    fmt.Println(v12, v12)
    if false {
        v13 := GenS1{F0: 2147483648}
        _ = v13
    }
    v3 = 9223372036854775807
    v14 := GenS1{F0: 147}
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
}
