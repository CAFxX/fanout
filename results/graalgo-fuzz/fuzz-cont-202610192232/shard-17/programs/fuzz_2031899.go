package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint16 = 42
    var v2 float32 = 1.5
    v3 := make(map[string]bool)
    var v4 float32 = 0.1
    v5 := make(map[int]uint32)
    switch (("\n\t" + "hello") + ("foo bar" + "foo bar")) {
    case "\n\t":
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := GenS1{F0: "\n\t", F1: 0}
            switch "\n\t" {
            case "\n\t":
                if false {
                    if true {
                        v2 = 1e10
                    }
                } else {
                    v8 := GenS1{F0: (("foo bar" + "\n\t") + (((("\n\t" + "") + ("" + ("hello" + ""))) + "") + "foo bar")), F1: 32767}
                    _ = v8
                }
            case "a":
                fmt.Println(v2, v4)
            case "":
                for fi9 := 0; fi9 < 13; fi9++ {
                    v4 += v2
                    _ = fi9
                }
            }
            _ = fi6
            _ = v7
        }
    case "":
        var v10 int64 = 1
        _ = v10
    default:
        v4 += (0.0 * float32((0.0 / 1.0)))
    }
    var v11 uint = 2
    var v12 uint16 = (^(v1))
    v13 := GenS1{F0: "\n\t", F1: 42}
    var v14 int8 = -1
    fmt.Println(v12, v4, v12)
    var v15 int32 = -28
    if (false || false) {
        var v16 uint16 = v1
        fmt.Println(v16, v2)
        var v17 int8 = 0
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(v14)
        if false {
            switch 2147483647 {
            case 127:
                fmt.Println(v14, fi18, v4)
            case 9223372036854775807:
                for fi19 := 0; fi19 < 19; fi19++ {
                    v20 := []int{fi18, 32768, fi18}
                    _ = fi19
                    _ = v20
                }
            case 318:
                var v21 float64 = 0.0
                _ = v21
            }
        }
        _ = fi18
    }
    v22 := GenS2{F0: v12}
    v23 := v5[2147483647]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
