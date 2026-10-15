package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
func genFn0() uint8 {
    var v4 bool = (("\n\t" + ("\n\t" + "a")) >= ("" + "foo bar"))
    fmt.Println(v4)
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: v4}
        var v7 int8 = int8(fi5)
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = v4
    return 255
}

func main() {
    var v1 int32 = 384
    var v2 float64 = ((3.14159 + float64(3.14159)) + 0.1)
    var v3 float32 = (((float32(1e10) - (((1e10 + 0.0) * float32(3.14159)) - 1.5)) - 0.0) - 0.0)
    if (!(140 != 1)) {
        v1 = -845
        v8 := GenS2{F0: 32768}
        v3 -= v3
        _ = v8
    }
    var v9 uint32 = 4294967295
    var v10 uint = 2
    v9 = (v9 >> (((v10 ^ (v10 - 341))) & 63))
    if true {
        v9 = v9
        v2 = -2.25
        switch ("hello" + "foo bar") {
        case "\n\t":
            var v11 int8 = 0
            _ = v11
        }
    } else {
        var v12 int32 = (v1 / (((v1 + v1)) | 1))
        switch "a" {
        case "hello":
            v10 += v10
        case "a":
            v1 = -2147483648
        }
        _ = v12
    }
    var v13 uint64 = 2
    v14 := []int{1}
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(v1, v3)
        for fi16 := 0; fi16 < 13; fi16++ {
            v14 = append(v14, (-(fi16)))
            _ = len(v14)
            _ = fi16
        }
        _ = fi15
    }
    v9 = (v9 / ((v9) | 1))
    var v17 int16 = -1
    var v18 int16 = -1
    v19 := GenS1{F0: true}
    switch 508 {
    case 1:
        if len(v14) > 0 {
            v14[0] = 749
        }
    case 4294967295:
        v13 = (v13 | (v13 & v13))
    case -756:
        var v20 float64 = (v2 / 3.14159)
        _ = v20
    }
    fmt.Println(v10, v1, v13)
    var v21 int8 = -128
    v17 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v9
}
