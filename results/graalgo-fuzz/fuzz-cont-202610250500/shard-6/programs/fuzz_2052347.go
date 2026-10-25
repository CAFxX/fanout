package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 bool }
func genFn0() uint64 {
    switch "hello" {
    case "foo bar":
        switch 32768 {
        case 42:
            v4 := []int{}
            _ = v4
        case 0:
            v5 := GenS1{F0: -1}
            _ = v5
        case 4294967295:
            v6 := GenS2{F0: 0, F1: true}
            _ = v6
        default:
            var v7 int8 = -17
            _ = v7
        }
    default:
        var v8 bool = true
        _ = v8
    }
    var v9 float64 = 3.14159
    fmt.Println(v9)
    var v10 int64 = 1
    _ = v10
    _ = v9
    return 4294967295
}

func main() {
    var v1 int64 = -1
    var v2 uint64 = 0
    var v3 int64 = v1
    v11 := GenS1{F0: 1000}
    switch ("\n\t" + ("\n\t" + (((((("a" + "hello") + ("a" + (("hello" + "foo bar") + "a"))) + ("foo bar" + ("a" + "a"))) + "\n\t") + "a") + ""))) {
    case "":
        var v12 string = "a"
        _ = v12
    }
    switch 2147483647 {
    case 4294967295:
        v1 -= v3
    default:
        var v13 int32 = 0
        _ = v13
    }
    v2 = 233
    fmt.Println(v3)
    var v14 string = "foo bar"
    for fi15 := 0; fi15 < 19; fi15++ {
        if true {
            switch fi15 {
            case 1:
                fmt.Println(v3, v1, v14)
            }
            v2 += (v2 - 4294967295)
        }
        fmt.Println(v3, fi15)
        _ = fi15
    }
    fmt.Println(v1, v14)
    v1 += 0
    v16 := []int{-9223372036854775808, 0}
    if (true && (("hello" + "\n\t") <= "")) {
        v16 = append(v16, 72)
    } else {
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    v3 = -1
    v17 := GenS2{F0: 129, F1: true}
    v18 := GenS1{F0: -209}
    switch v14 {
    case "\n\t":
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS1{F0: int16(v2)}
            _ = fi19
            _ = v20
        }
    case "":
        var v21 uint16 = 2
        _ = v21
    default:
        _ = len(v16)
    }
    _ = len(v16)
    var v22 bool = true
    var v23 uint8 = 47
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
