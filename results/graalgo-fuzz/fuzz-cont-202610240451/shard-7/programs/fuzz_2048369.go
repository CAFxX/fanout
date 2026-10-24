package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 uint32 = 769
    var v2 int32 = 1
    v3 := []string{"\n\t", ("" + ("foo bar" + "")), "hello", (("a" + "a") + ("a" + "foo bar"))}
    var v4 int8 = -128
    v1 -= v1
    switch "\n\t" {
    case "":
        v5 := GenS1{F0: 2}
        _ = v5
    case "foo bar":
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 uint8 = 0
            if (!(1 <= 770)) {
                v3 = append(v3, "\n\t")
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v8 int32 = (v2 | (v2 & (v2 - v2)))
            _ = fi6
            _ = v7
            _ = v8
        }
    case "a":
        var v9 int64 = -9223372036854775808
        _ = v9
    }
    var v10 int64 = 0
    var v11 int16 = int16(v4)
    var v12 int64 = 405
    v13 := GenS1{F0: 2}
    var v14 float64 = (float64(1e10) + 0.1)
    var v15 string = ((("hello" + "foo bar") + "\n\t") + "\n\t")
    v16 := GenS1{F0: 0}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v17 int8 = v4
    var v18 uint64 = 91
    var v19 uint = 65535
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 uint16 = 255
        var v22 uint32 = uint32(v18)
        v23 := GenS1{F0: (v21 + 65535)}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
