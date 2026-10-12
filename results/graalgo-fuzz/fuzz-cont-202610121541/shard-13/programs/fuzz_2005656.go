package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() int64 {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int = (-(fi6))
        switch fi6 {
        case 606:
            var v8 int = (v7 - fi6)
            _ = v8
        case 9223372036854775807:
            v7 = (fi6 + (fi6 >> ((541) & 63)))
        case -273:
            fmt.Println(fi6, v7)
        default:
            if false {
                v7 = fi6
            } else {
                var v9 int16 = -301
                if false {
                    v9 = 900
                    fmt.Println(v7)
                }
                _ = v9
            }
        }
        v10 := []int{}
        _ = fi6
        _ = v10
        _ = v7
    }
    var v11 string = "foo bar"
    _ = v11
    return -989
}

func main() {
    var v1 uint32 = 65535
    var v2 string = ((((("foo bar" + "foo bar") + "foo bar") + "hello") + "foo bar") + "a")
    v3 := make(map[uint32]string)
    var v4 int16 = 1
    var v5 float64 = 3.14159
    v12 := GenS1{F0: 4294967295}
    v4 = (v4 - 42)
    fmt.Println(v1, v4)
    if true {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint64 = 42
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: int(v1)}
        _ = v15
    } else {
        switch v2 {
        case "\n\t":
            v16 := GenS1{F0: 508}
            _ = v16
        case "hello":
            var v17 uint16 = 32767
            _ = v17
        }
    }
    v18 := v3[v1]
    var v19 uint16 = 32767
    var v20 bool = false
    var v21 float32 = (0.0 - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
