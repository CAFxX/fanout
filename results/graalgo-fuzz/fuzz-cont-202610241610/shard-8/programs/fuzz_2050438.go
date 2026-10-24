package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint16 = 2
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    var v5 uint64 = 845
    v6 := GenS1{F0: 42, F1: "a", F2: 42}
    v7 := []int{}
    fmt.Println(v1, v1, v2)
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 float32 = -2.25
        for fi10 := 0; fi10 < 18; fi10++ {
            fmt.Println(v5)
            _ = fi10
        }
        fmt.Println(v5)
        _ = fi8
        _ = v9
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    delete(v4, (("foo bar" + "") + ""))
    var v11 uint64 = 42
    switch ((("hello" + "") + ("a" + "")) + "a") {
    case "a":
        v12 := v3["hello"]
        _ = v12
    }
    var v13 uint8 = 0
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS2{F0: 100, F1: -128, F2: 1}
        _ = fi14
        _ = v15
    }
    var v16 int8 = 100
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 bool = false
        _ = fi17
        _ = v18
    }
    switch "\n\t" {
    case "foo bar":
        v5 = v11
    case "\n\t":
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 bool = true
            for fi21 := 0; fi21 < 19; fi21++ {
                v22 := GenS3{F0: v13}
                _ = fi21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
    case "":
        v16 -= (^(v16))
    default:
        if len(v7) > 0 {
            v7[0] = 32767
        }
    }
    delete(v4, "")
    var v23 uint8 = uint8(v11)
    if ((v1 >> ((65535) & 63)) >= v1) {
        v5 = (v5 * 4294967295)
        var v24 int64 = v2
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v13
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
