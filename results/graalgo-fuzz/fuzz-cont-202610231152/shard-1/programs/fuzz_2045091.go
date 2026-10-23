package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int32 = 42
    v2 := make(map[string]int)
    var v3 int64 = -55
    fmt.Println(v1, v3)
    var v4 uint64 = 2
    v3 += (v3 >> ((545) & 63))
    v2[("\n\t" + "hello")] = 32768
    var v5 uint16 = 796
    if true {
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint16 = v5
            _ = fi6
            _ = v7
        }
        var v8 float32 = ((3.14159 / float32(3.14159)) + float32(((1.5 + float64(0.0)) - 0.1)))
        var v9 int16 = -32768
        _ = v8
        _ = v9
    } else {
        var v10 uint32 = 65535
        v11 := GenS1{F0: false}
        var v12 int = 32768
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 uint16 = v5
    var v14 bool = true
    var v15 bool = v14
    if (!(2 >= 622)) {
        for fi16 := 0; fi16 < 18; fi16++ {
            switch fi16 {
            case 705:
                delete(v2, "\n\t")
            }
            fmt.Println(fi16, v14, v14)
            delete(v2, ("a" + ("" + "a")))
            _ = fi16
        }
    } else {
        switch "a" {
        case "hello":
            switch -174 {
            case -877:
                v17 := []int{}
                _ = v17
            default:
                v2["hello"] = 127
            }
        }
        v18 := []int{}
        _ = v18
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
