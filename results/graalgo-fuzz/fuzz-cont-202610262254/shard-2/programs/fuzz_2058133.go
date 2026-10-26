package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 847
    var v3 uint32 = 42
    fmt.Println(v3, v3)
    if ((^(v3)) != 2147483647) {
        for fi4 := 0; fi4 < 16; fi4++ {
            for fi5 := 0; fi5 < 16; fi5++ {
                var v6 int16 = 1000
                _ = fi5
                _ = v6
            }
            v7 := GenS1{F0: -1, F1: 0, F2: "a"}
            _ = fi4
            _ = v7
        }
    }
    if (1.5 >= 0.0) {
        var v8 uint8 = 42
        _ = v8
    } else {
        for fi9 := 0; fi9 < 17; fi9++ {
            if ((("hello" + ("foo bar" + ("\n\t" + ("hello" + ((("hello" + "foo bar") + "") + ("\n\t" + ("" + "foo bar"))))))) + "a") == "\n\t") {
                var v10 string = ""
                v11 := []int{}
                fmt.Println(v3, v3, v10)
                _ = v10
                _ = v11
            } else {
                var v12 bool = false
                var v13 bool = (127 <= 1)
                _ = v12
                _ = v13
            }
            v2 += (v2 & (v2 + v2))
            _ = fi9
        }
    }
    var v14 uint = uint(v3)
    if true {
        var v15 uint32 = uint32(v2)
        switch ("foo bar" + "hello") {
        case "a":
            v2 = 42
        case "\n\t":
            v3 = 4294967295
        case "hello":
            fmt.Println(v15, v15, v3)
        }
        fmt.Println(v2)
        _ = v15
    } else {
        v16 := GenS2{F0: (!(!(!false))), F1: (v3 * (v3 % ((v3) | 1)))}
        _ = v16
    }
    fmt.Println(v3)
    v17 := GenS2{F0: ((1 < 1) && true), F1: 2147483648}
    v14 = (v14 >> ((255) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
}
