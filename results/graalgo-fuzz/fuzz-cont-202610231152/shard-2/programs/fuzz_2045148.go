package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() string {
    v4 := []int{1, 9223372036854775807, -1}
    var v5 string = "\n\t"
    v5 = (v5 + "\n\t")
    _ = v4
    _ = v5
    return v5
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{("" + "\n\t"), ((("a" + "a") + ((((((("hello" + "foo bar") + "a") + "a") + "hello") + (("hello" + "hello") + "hello")) + ("hello" + "")) + ("" + "foo bar"))) + "a"), "", ("foo bar" + "\n\t")}
    var v3 string = "hello"
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(v3, fi6)
        _ = fi6
    }
    var v7 int16 = 0
    if false {
        delete(v1, -52)
    }
    _ = len(v2)
    var v8 int = -679
    for fi9 := 0; fi9 < 17; fi9++ {
        v8 = (fi9 - (fi9 - fi9))
        for fi10 := 0; fi10 < 16; fi10++ {
            v1[(^(v8))] = 48
            _ = fi10
        }
        _ = fi9
    }
    v3 = (v3 + v3)
    v7 += 32767
    var v11 int16 = v7
    var v12 string = v3
    v13 := GenS1{F0: 394}
    v1[v8] = v8
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
