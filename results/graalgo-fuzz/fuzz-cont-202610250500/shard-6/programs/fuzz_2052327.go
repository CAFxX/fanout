package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]bool)
    var v3 uint16 = 42
    var v4 bool = (13 >= 47)
    if (v4 && (true || false)) {
        v4 = ((!v4) || (!v4))
        var v5 string = ""
        _ = v5
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int32 = -1
        switch "" {
        case "a":
            v8 := []int{}
            _ = v8
        case "hello":
            v9 := GenS1{F0: 1}
            _ = v9
        }
        v10 := []int{32768, 2147483647, (fi6 - fi6)}
        _ = fi6
        _ = v10
        _ = v7
    }
    var v11 int64 = 121
    v12 := GenS1{F0: v11}
    var v13 string = "a"
    v13 = ("hello" + v13)
    v14 := GenS1{F0: (v11 % (((v11 + v11)) | 1))}
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    v2[711] = v4
    var v17 string = (v13 + (v13 + (v13 + v13)))
    fmt.Println(v4, v13)
    if v4 {
        if false {
            v18 := v2[343]
            v19 := GenS1{F0: 156}
            _ = v18
            _ = v19
        } else {
            v17 = "\n\t"
        }
        var v20 int64 = v11
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            v13 = (v17 + "hello")
            fmt.Println(v3)
            _ = fi22
        }
        v1 -= -213
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
