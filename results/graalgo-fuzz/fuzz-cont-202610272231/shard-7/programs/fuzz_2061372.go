package main

import (
    "fmt"
)

func genFn0() int {
    var v7 bool = true
    v7 = v7
    v7 = v7
    v8 := []int{32768, 0}
    _ = v7
    _ = v8
    return 128
}

func main() {
    var v1 int64 = 14
    var v2 float32 = 1e10
    v3 := []string{("hello" + "")}
    var v4 int8 = 50
    var v5 uint8 = 217
    v6 := make(map[string]int)
    if (-1 >= v4) {
        v1 = 9223372036854775807
    }
    if false {
        v1 = v1
        var v9 uint32 = 65535
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int8 = v4
        var v12 bool = false
        _ = fi10
        _ = v11
        _ = v12
    }
    v5 -= (^(v5))
    fmt.Println(v1)
    switch (("foo bar" + "a") + "hello") {
    case "a":
        var v13 uint16 = 255
        _ = v13
    case "hello":
        if true {
            switch 935 {
            case 128:
                var v14 int = -46
                _ = v14
            case 2147483647:
                v5 = 121
            default:
                if false {
                    v5 += v5
                }
            }
            fmt.Println(v5, v1)
        }
    case "\n\t":
        var v15 int32 = 2147483647
        _ = v15
    }
    var v16 int16 = 1
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 int8 = 100
        var v19 uint32 = 520
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
