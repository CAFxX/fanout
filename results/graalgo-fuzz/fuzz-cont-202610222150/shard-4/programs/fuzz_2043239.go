package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    var v1 uint64 = 2
    var v2 bool = false
    v3 := make(map[string]uint32)
    var v4 float32 = (0.1 / -2.25)
    v5 := make(map[string]string)
    var v6 uint = 65535
    var v7 bool = (true && ((^(v1)) > (v1 + 18446744073709551615)))
    v1 += (v1 - 4294967295)
    if ((v7 || ((-32768 > -1) && v7)) || v2) {
        fmt.Println(v2, v4)
    } else {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint16 = 59
            _ = fi8
            _ = v9
        }
    }
    v7 = (42 != (v1 & 4294967295))
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 bool = v2
        _ = fi10
        _ = v11
    }
    if false {
        switch "foo bar" {
        case "foo bar":
            var v12 uint32 = 0
            _ = v12
        case "hello":
            switch (("" + ("\n\t" + "hello")) + ("" + ("\n\t" + "a"))) {
            case "\n\t":
                var v13 uint = 510
                _ = v13
            default:
                v7 = true
            }
        case "a":
            v14 := []int{0, int(v6), 42, 1}
            _ = v14
        }
        var v15 int64 = 1
        _ = v15
    }
    v4 -= -2.25
    switch 32767 {
    case 42:
        v16 := GenS1{F0: 65535, F1: -972}
        _ = v16
    case 32767:
        if ("a" != "a") {
            delete(v3, (("hello" + "foo bar") + ""))
        }
    case 2147483647:
        v6 -= 18446744073709551615
    }
    _ = len(v5)
    v17 := GenS1{F0: 236, F1: 0}
    switch "foo bar" {
    case "a":
        for fi18 := 0; fi18 < 17; fi18++ {
            v6 = (v6 * (v6 << (((v6 * 116)) & 63)))
            _ = fi18
        }
    case "":
        v6 = 4294967295
    }
    fmt.Println(v1, v6)
    v19 := []int{1}
    v4 = ((-2.25 - v4) / ((1.5 - float32(3.14159)) + ((v4 / (v4 / 3.14159)) * v4)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
