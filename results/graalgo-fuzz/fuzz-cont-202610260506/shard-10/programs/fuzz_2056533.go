package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int64 = 670
    v2 := []uint32{853}
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    v1 = v1
    var v5 float64 = 3.14159
    if (true || (1 != 42)) {
        v1 += (v1 + -1)
        if false {
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 uint32 = 41
                v8 := GenS1{F0: 2}
                _ = fi6
                _ = v7
                _ = v8
            }
        } else {
            v9 := v4[(("foo bar" + "foo bar") + "a")]
            _ = v9
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v10 uint = 918
        v11 := v4[("\n\t" + "a")]
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 255}
    var v13 float32 = (1e10 - -2.25)
    if len(v2) > 0 {
        v2[0] = 916
    }
    var v14 int32 = 367
    switch 127 {
    case 987:
        v13 = (v13 + v13)
    case 32767:
        var v15 int64 = -1
        _ = v15
    default:
        var v16 int8 = 1
        _ = v16
    }
    v14 = v14
    v14 -= (^(v14))
    v2 = append(v2, 65535)
    if true {
        var v17 string = "\n\t"
        _ = v17
    }
    var v18 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
