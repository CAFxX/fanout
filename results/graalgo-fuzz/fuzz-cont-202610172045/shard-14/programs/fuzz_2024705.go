package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
func genFn0() uint16 {
    var v7 int16 = 589
    v8 := []int{4294967295}
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 2
    var v3 uint8 = 127
    var v4 string = "hello"
    var v5 uint8 = v3
    v6 := make(map[string]string)
    if (true && (!(739 == 42))) {
        v4 = (((((v4 + "") + ("hello" + "a")) + v4) + v4) + ("a" + ((v4 + (("foo bar" + ((v4 + v4) + ("a" + (v4 + v4)))) + "a")) + (v4 + ((("hello" + ("foo bar" + v4)) + "hello") + (((("hello" + (((v4 + ((v4 + v4) + v4)) + "") + (((("foo bar" + v4) + "foo bar") + v4) + v4))) + ((v4 + (("hello" + ((v4 + v4) + "hello")) + v4)) + v4)) + v4) + "\n\t"))))))
        v9 := GenS1{F0: v4, F1: 255}
        _ = v9
    } else {
        var v10 bool = false
        var v11 float32 = -2.25
        _ = v10
        _ = v11
    }
    switch "hello" {
    case "foo bar":
        var v12 bool = true
        _ = v12
    case "\n\t":
        v13 := []int{-31, 2147483647, 9223372036854775807, -243}
        _ = v13
    default:
        var v14 float64 = (float64(3.14159) * float64(3.14159))
        _ = v14
    }
    switch -1 {
    case -2147483648:
        v6["hello"] = v4
    case 219:
        v2 += 703
    case 128:
        v15 := []int{867}
        _ = v15
    }
    v16 := GenS1{F0: v4, F1: 32767}
    v17 := GenS1{F0: v4, F1: 440}
    _ = len(v6)
    var v18 int8 = -128
    _ = len(v6)
    var v19 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
