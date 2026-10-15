package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }

func main() {
    var v1 uint = 2
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    v3["hello"] = 2
    v2[127] = ("" + ((((("" + "\n\t") + "hello") + ((((("hello" + "\n\t") + (((((("a" + "foo bar") + ("" + ("hello" + "foo bar"))) + "\n\t") + (("a" + "") + "")) + "hello") + ("hello" + ("hello" + ("hello" + ("\n\t" + ("" + "\n\t"))))))) + "") + ("a" + "a")) + "foo bar")) + "hello") + ""))
    switch 32768 {
    case 128:
        for fi4 := 0; fi4 < 13; fi4++ {
            var v5 int32 = 2147483647
            _ = fi4
            _ = v5
        }
    case 394:
        _ = len(v2)
    default:
        v1 = v1
    }
    switch 2147483647 {
    case 42:
        fmt.Println(v1)
    default:
        fmt.Println(v1)
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        _ = len(v3)
        if true {
            v7 := []int{fi6, fi6, (fi6 & (fi6 % ((fi6) | 1))), 32767}
            _ = v7
        }
        _ = fi6
    }
    delete(v2, 127)
    v1 = (^(v1))
    v1 = (v1 / ((2) | 1))
    fmt.Println(v1)
    v1 = v1
    v1 += (v1 + 1)
    v1 = (v1 % (((v1 * (v1 / ((v1) | 1)))) | 1))
    for fi8 := 0; fi8 < 1; fi8++ {
        switch (fi8 / (((fi8 | 127)) | 1)) {
        case 128:
            v9 := GenS1{F0: 749, F1: 0, F2: 2}
            _ = v9
        }
        _ = fi8
    }
    if ((0 != 18446744073709551615) || true) {
        for fi10 := 0; fi10 < 13; fi10++ {
            v2[fi10] = ("" + "hello")
            var v11 int64 = 42
            _ = fi10
            _ = v11
        }
    } else {
        v12 := GenS1{F0: 88, F1: 42, F2: 1}
        var v13 uint = (v1 - (v1 % ((4) | 1)))
        var v14 uint = 255
        _ = v12
        _ = v13
        _ = v14
    }
    _ = len(v2)
    var v15 bool = true
    fmt.Println(v1, v15)
    v15 = (0 <= 42)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
}
