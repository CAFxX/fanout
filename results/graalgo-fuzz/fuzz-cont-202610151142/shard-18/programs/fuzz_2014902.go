package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 uint = 65535
        v6 := []int{fi4, fi4, (^(fi4))}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int8 = 1
    var v8 uint32 = 4294967295
    if (0 != 4294967295) {
        v8 = (^(v8))
        v7 = (v7 & (v7 ^ (v7 << ((65535) & 63))))
    }
    var v9 uint16 = 0
    switch "foo bar" {
    case "a":
        if true {
            var v10 uint32 = 0
            _ = v10
        }
    case "hello":
        var v11 uint = 4294967295
        _ = v11
    case "":
        if ((-1 >= 32767) || false) {
            var v12 string = ("" + (("hello" + "\n\t") + ("hello" + "")))
            v13 := []int{42, 494, 190, 383}
            fmt.Println(v7, v7)
            _ = v12
            _ = v13
        }
    }
    v14 := v3[(("" + ("" + ("" + ("a" + "foo bar")))) + (("a" + "a") + (("hello" + "a") + "foo bar")))]
    var v15 float32 = 0.1
    v14 = 32768
    var v16 uint = 18446744073709551615
    v17 := []int{}
    var v18 string = "hello"
    if false {
        var v19 int16 = -282
        var v20 int8 = -57
        _ = v19
        _ = v20
    }
    v21 := v3["a"]
    if true {
        var v22 float32 = v15
        _ = v22
    } else {
        v18 = ((v18 + "foo bar") + "foo bar")
        v8 = 0
    }
    v23 := GenS1{F0: (!false), F1: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
