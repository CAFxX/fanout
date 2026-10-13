package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 float32 = 0.1
    var v2 float32 = (1.5 * v1)
    v3 := make(map[uint32]int)
    var v4 int32 = 2147483647
    var v5 bool = true
    v6 := GenS1{F0: 0}
    v7 := GenS1{F0: 42}
    fmt.Println(v5, v4, v4)
    if v5 {
        var v8 float64 = 1e10
        v8 = 1e10
        _ = v8
    } else {
        v4 = 42
        var v9 string = "hello"
        v4 = v4
        _ = v9
    }
    v3[65535] = 127
    delete(v3, 2147483647)
    v1 = 1.5
    if true {
        v5 = false
        switch ("\n\t" + ("\n\t" + "hello")) {
        case "\n\t":
            v4 = v4
        case "a":
            fmt.Println(v2, v2)
        }
        v10 := []int{32767, -986, 42, -9223372036854775808}
        _ = v10
    } else {
        if (128 < 327) {
            fmt.Println(v2)
        }
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := v3[2]
        fmt.Println(v1, v4)
        fmt.Println(v1, v1)
        _ = fi11
        _ = v12
    }
    switch 32767 {
    case 829:
        var v13 int = -2147483648
        _ = v13
    case 4294967295:
        var v14 uint8 = 1
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
