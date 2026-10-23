package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }

func main() {
    v1 := []int{32767}
    var v2 uint = 773
    var v3 int64 = -9223372036854775808
    var v4 int32 = int32(v3)
    for fi5 := 0; fi5 < 8; fi5++ {
        v4 += (v4 & v4)
        _ = fi5
    }
    if (((42 > -1) && true) && (128 > 128)) {
        var v6 int = 0
        var v7 int32 = v4
        v2 = v2
        _ = v6
        _ = v7
    }
    v2 = v2
    v8 := GenS1{F0: uint32(v2), F1: 1}
    for fi9 := 0; fi9 < 2; fi9++ {
        v1 = append(v1, (fi9 & fi9))
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
        _ = fi9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint32 = 1
        v2 = v2
        var v12 string = ((("a" + "a") + ("foo bar" + "hello")) + ((("" + "\n\t") + "hello") + "foo bar"))
        _ = fi10
        _ = v11
        _ = v12
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = (v4 - (v4 >> ((1) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
