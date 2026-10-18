package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint16; F2 string }

func main() {
    var v1 int64 = 1
    var v2 uint32 = 2
    v3 := []uint32{}
    var v4 uint8 = 127
    var v5 uint32 = 2
    var v6 uint8 = v4
    v5 = v5
    var v7 float64 = 0.0
    fmt.Println(v2)
    v7 += 3.14159
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint16 = 32767
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        if (1 == 885) {
            fmt.Println(v4)
        }
        v11 := GenS1{F0: v6}
        _ = fi10
        _ = v11
    }
    var v12 int64 = v1
    var v13 bool = true
    if (true || (true && false)) {
        v14 := GenS2{F0: -1, F1: 507, F2: ("hello" + ("" + (("a" + "\n\t") + ("foo bar" + ""))))}
        _ = v14
    }
    v15 := GenS1{F0: (v6 << ((1) & 63))}
    switch (((("hello" + "") + "a") + (("\n\t" + "") + (("a" + (((("hello" + "a") + ("foo bar" + "foo bar")) + ("" + "hello")) + ("\n\t" + "hello"))) + "foo bar"))) + "hello") {
    case "":
        v2 += (^(v2))
    case "a":
        v13 = v13
    case "hello":
        var v16 uint8 = (v6 - (^(v6)))
        _ = v16
    }
    v2 = v5
    for fi17 := 0; fi17 < 18; fi17++ {
        v18 := GenS1{F0: v6}
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
