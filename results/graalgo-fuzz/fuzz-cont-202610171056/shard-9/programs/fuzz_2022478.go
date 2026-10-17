package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 bool }

func main() {
    var v1 int16 = 953
    v2 := make(map[string]int)
    v3 := []int{127, 2147483647, -1}
    for fi4 := 0; fi4 < 14; fi4++ {
        if true {
            v1 += v1
            v1 = (v1 % (((v1 / ((v1) | 1))) | 1))
            delete(v2, "hello")
        }
        _ = fi4
    }
    v5 := GenS1{F0: (((("a" + ("a" + "hello")) + "hello") + "a") < ("" + ("hello" + ((("\n\t" + "\n\t") + (("a" + "a") + "foo bar")) + "hello")))), F1: 128, F2: (((228 < 127) && false) || (!true))}
    if (1 < 381) {
        fmt.Println(v1)
    } else {
        v6 := GenS1{F0: (true && (!true)), F1: 0, F2: true}
        v1 += (v1 >> ((4294967295) & 63))
        fmt.Println(v1)
        _ = v6
    }
    delete(v2, "\n\t")
    if (127 < 236) {
        var v7 uint = 4294967295
        _ = v7
    }
    if false {
        if false {
            var v8 int8 = 1
            _ = v8
        }
        v2["hello"] = 9223372036854775807
        v1 -= v1
    }
    v1 += 473
    var v9 string = "a"
    fmt.Println(v9, v9)
    v1 = (v1 - 723)
    v9 = v9
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v10 int = 9223372036854775807
    v1 -= 1000
    var v11 int = 128
    if (v1 <= (v1 | v1)) {
        v10 -= v10
        _ = len(v2)
    }
    var v12 int = v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
