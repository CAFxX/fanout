package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 string; F1 uint16 }

func main() {
    var v1 string = ("\n\t" + "\n\t")
    v2 := make(map[string]uint32)
    var v3 uint = 2
    var v4 string = v1
    var v5 float32 = 3.14159
    var v6 bool = true
    var v7 int = 2147483647
    var v8 uint = v3
    v9 := GenS1{F0: 1, F1: ((0 != 255) && (v6 || (!v6)))}
    v6 = false
    v4 = (v4 + ("\n\t" + (("a" + v1) + "foo bar")))
    fmt.Println(v3, v6)
    if (v7 > (v7 ^ 581)) {
        var v10 uint64 = 177
        _ = v10
    }
    switch "a" {
    case "hello":
        v11 := []int{-1}
        _ = v11
    case "a":
        if (794 <= 65535) {
            v12 := GenS2{F0: "", F1: 42}
            if v6 {
                if v6 {
                    v1 = ("hello" + v4)
                    v6 = (106 <= 239)
                }
            } else {
                v6 = v6
            }
            _ = v12
        }
    case "":
        fmt.Println(v1)
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS2{F0: "a", F1: 42}
        _ = fi13
        _ = v14
    }
    switch "hello" {
    case "":
        fmt.Println(v4, v5)
    case "foo bar":
        v6 = true
    case "a":
        var v15 string = "foo bar"
        _ = v15
    default:
        for fi16 := 0; fi16 < 1; fi16++ {
            if (1 == 0) {
                v6 = (v3 > 18446744073709551615)
                v3 = (v3 / ((v3) | 1))
                v17 := GenS1{F0: -1, F1: false}
                _ = v17
            } else {
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := GenS2{F0: v1, F1: 65535}
                    v6 = v6
                    v6 = (true && false)
                    _ = fi18
                    _ = v19
                }
                if (!v6) {
                    var v20 int16 = 32767
                    _ = v20
                }
            }
            switch ((v1 + (v4 + (v1 + v1))) + (v1 + (v4 + (v4 + (("" + "\n\t") + v1))))) {
            case "foo bar":
                v8 += v3
            case "\n\t":
                var v21 uint = 557
                _ = v21
            case "a":
                var v22 uint32 = 2147483648
                _ = v22
            }
            _ = fi16
        }
    }
    _ = len(v2)
    _ = len(v2)
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 uint = v3
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
