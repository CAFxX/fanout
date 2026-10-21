package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }

func main() {
    var v1 float32 = (float32(3.14159) * -2.25)
    var v2 float32 = 3.14159
    var v3 uint = 255
    var v4 uint8 = 1
    var v5 float64 = -2.25
    if (128 != int(v4)) {
        for fi6 := 0; fi6 < 17; fi6++ {
            for fi7 := 0; fi7 < 10; fi7++ {
                fmt.Println(v5)
                v3 -= v3
                _ = fi7
            }
            fmt.Println(v1)
            _ = fi6
        }
    }
    var v8 bool = false
    switch -2147483648 {
    case 42:
        var v9 uint16 = 654
        _ = v9
    case 0:
        if v8 {
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := []int{fi10, 128, (fi10 - (fi10 - 0)), (fi10 / ((fi10) | 1))}
                _ = fi10
                _ = v11
            }
            fmt.Println(v1, v8, v3)
        } else {
            switch ((("\n\t" + "hello") + "foo bar") + "foo bar") {
            case "a":
                if v8 {
                    var v12 float32 = 0.0
                    v1 += float32(v5)
                    fmt.Println(v3)
                    _ = v12
                }
            }
            fmt.Println(v8, v8)
            v13 := []int{-132}
            _ = v13
        }
    }
    v14 := GenS1{F0: 4294967295, F1: -1}
    v15 := []int{}
    switch "\n\t" {
    case "":
        var v16 int64 = 0
        _ = v16
    }
    switch -2147483648 {
    case 0:
        var v17 string = "foo bar"
        _ = v17
    default:
        var v18 string = ("hello" + "foo bar")
        _ = v18
    }
    switch 32767 {
    case 0:
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := GenS2{F0: 902, F1: int64(fi19), F2: 255}
            _ = fi19
            _ = v20
        }
    case 42:
        switch "hello" {
        case "hello":
            v1 = v1
        case "\n\t":
            var v21 bool = v8
            _ = v21
        }
    case 1000:
        v5 = ((v5 + v5) / v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
