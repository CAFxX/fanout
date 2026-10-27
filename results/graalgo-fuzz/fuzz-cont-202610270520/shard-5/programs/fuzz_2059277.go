package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
func genFn0(p0 uint8) int {
    var v4 uint32 = 2
    var v5 bool = false
    _ = p0
    _ = v4
    _ = v5
    return 32768
}

func main() {
    var v1 int16 = 32767
    var v2 uint = 0
    v3 := []string{("a" + "a")}
    v1 += v1
    v1 -= (v1 >> ((280) & 63))
    if (65535 == 2147483647) {
        v1 = (v1 * -1)
        if true {
            var v6 uint32 = 65535
            _ = v6
        } else {
            var v7 int8 = -20
            _ = v7
        }
    } else {
        var v8 int32 = 319
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        fmt.Println(v2, v1, v1)
        for fi10 := 0; fi10 < 12; fi10++ {
            v2 = 968
            if true {
                v3 = append(v3, "foo bar")
                v1 = -32768
                v11 := GenS1{F0: genFn0(245), F1: -128, F2: 1}
                _ = v11
            }
            _ = fi10
        }
        _ = fi9
    }
    v2 = (v2 - (v2 >> ((413) & 63)))
    var v12 uint8 = 42
    v1 = v1
    switch (("" + ("hello" + "")) + "a") {
    case "foo bar":
        fmt.Println(v12, v1, v12)
    case "hello":
        var v13 uint8 = 119
        _ = v13
    case "a":
        v12 = (v12 + v12)
    }
    var v14 bool = (-408 >= 0)
    v3 = append(v3, "foo bar")
    if false {
        v2 += (v2 - 65535)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
}
