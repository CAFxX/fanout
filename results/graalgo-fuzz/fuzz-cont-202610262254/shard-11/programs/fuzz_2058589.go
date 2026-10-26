package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 bool }
type GenS2 struct { F0 bool; F1 int64 }

func main() {
    v1 := []string{("hello" + ("hello" + ("a" + (("hello" + "hello") + ("\n\t" + "hello"))))), "hello"}
    var v2 uint64 = 1
    var v3 uint32 = 4294967295
    var v4 string = "\n\t"
    v5 := make(map[string]string)
    switch v4 {
    case "":
        v6 := GenS1{F0: v2, F1: 533, F2: false}
        _ = v6
    case "foo bar":
        var v7 uint32 = v3
        _ = v7
    default:
        if len(v1) > 0 {
            v1[0] = (("" + ((((v4 + (("" + v4) + ((v4 + (v4 + ((v4 + v4) + v4))) + ((v4 + v4) + v4)))) + (("" + (("foo bar" + ("a" + "foo bar")) + v4)) + ("foo bar" + v4))) + v4) + ((v4 + (((v4 + ((("\n\t" + v4) + v4) + ("\n\t" + (v4 + ("foo bar" + "a"))))) + (v4 + v4)) + ((v4 + "foo bar") + "\n\t"))) + v4))) + "foo bar")
        }
    }
    var v8 uint64 = 18446744073709551615
    v9 := GenS1{F0: v8, F1: 255, F2: false}
    v5[(("a" + "hello") + "")] = v4
    var v10 string = ""
    if len(v1) > 0 {
        v1[0] = v4
    }
    var v11 bool = true
    if (false || (696 < 764)) {
        if v11 {
            for fi12 := 0; fi12 < 8; fi12++ {
                _ = len(v1)
                _ = fi12
            }
        }
    } else {
        if len(v1) > 0 {
            v1[0] = (v4 + (v10 + ((v4 + (v4 + v4)) + v4)))
        }
        for fi13 := 0; fi13 < 16; fi13++ {
            v1 = append(v1, "hello")
            v14 := GenS1{F0: 200, F1: 65535, F2: v11}
            if ((true && v11) && (v3 == 2147483648)) {
                _ = len(v5)
                if (0 == 42) {
                    v8 = (v8 + 1)
                }
                for fi15 := 0; fi15 < 8; fi15++ {
                    v1 = append(v1, (v4 + v4))
                    delete(v5, v10)
                    v2 = (v2 >> ((50) & 63))
                    _ = fi15
                }
            }
            _ = fi13
            _ = v14
        }
        var v16 uint = 500
        _ = v16
    }
    v3 = v3
    _ = len(v5)
    var v17 int16 = 1
    if (563 < (v3 / (((v3 % ((73) | 1))) | 1))) {
        v17 = v17
    } else {
        fmt.Println(v8)
        if v11 {
            v10 = "hello"
            _ = len(v5)
            if v11 {
                var v18 int16 = v17
                _ = v18
            }
        }
        v19 := GenS2{F0: (v11 && ((0.1 >= (3.14159 + -2.25)) && v11)), F1: -9223372036854775808}
        _ = v19
    }
    var v20 int32 = -2147483648
    switch "a" {
    case "foo bar":
        var v21 uint16 = 42
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
