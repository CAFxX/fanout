package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }

func main() {
    var v1 uint16 = 32767
    var v2 uint32 = 1
    var v3 int32 = -1
    for fi4 := 0; fi4 < 7; fi4++ {
        for fi5 := 0; fi5 < 13; fi5++ {
            var v6 string = "a"
            var v7 int32 = 42
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 string = ((("a" + "a") + ("" + ("hello" + ("" + "foo bar")))) + (("a" + ((((("a" + "\n\t") + "foo bar") + (("hello" + "a") + ((("\n\t" + "hello") + "a") + "hello"))) + ("\n\t" + "")) + ("\n\t" + "\n\t"))) + "\n\t"))
        if true {
            v1 += uint16(v2)
            v9 := GenS1{F0: ((v8 + "") + v8), F1: false, F2: (-(fi4))}
            fmt.Println(v2, v2, v8)
            _ = v9
        }
        _ = fi4
        _ = v8
    }
    v10 := GenS1{F0: ("" + "\n\t"), F1: (("hello" != "\n\t") && (87 >= 42)), F2: 1}
    v11 := []int{}
    if len(v11) > 0 {
        v11[0] = -430
    }
    fmt.Println(v2, v3, v3)
    v2 -= 2
    var v12 int64 = 1
    v2 += v2
    _ = len(v11)
    v11 = append(v11, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
