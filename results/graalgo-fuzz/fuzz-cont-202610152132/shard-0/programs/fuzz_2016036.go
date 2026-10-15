package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int }

func main() {
    var v1 int8 = 1
    var v2 int32 = 1
    var v3 uint64 = 42
    v4 := []string{("hello" + (("a" + ("a" + ("foo bar" + ("hello" + "a")))) + (((("\n\t" + "foo bar") + "a") + "") + ("" + (("\n\t" + "") + ("" + "\n\t"))))))}
    var v5 string = "\n\t"
    var v6 int16 = 32767
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v7 bool = (546 < 42)
    fmt.Println(v5)
    var v8 int = -662
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    v7 = v7
    var v9 string = v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint64 = v3
        v12 := GenS1{F0: -42, F1: v6, F2: (v8 & fi10)}
        var v13 float32 = 0.1
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v2 = (v2 % (((v2 | (v2 & v2))) | 1))
    v14 := GenS1{F0: int(v6), F1: v6, F2: 58}
    var v15 string = "foo bar"
    v8 += 32768
    v4 = append(v4, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
