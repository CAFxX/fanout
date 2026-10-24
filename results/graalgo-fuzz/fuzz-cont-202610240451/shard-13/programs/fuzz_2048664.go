package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }

func main() {
    var v1 uint = 42
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    var v4 string = "a"
    var v5 string = ""
    fmt.Println(v1, v4, v1)
    v6 := v3[32767]
    v7 := GenS1{F0: "", F1: -754}
    if (2 < 198) {
        v5 = (v6 + v5)
    }
    v6 = (v6 + ((("foo bar" + ((v5 + ("foo bar" + "\n\t")) + v4)) + "\n\t") + v4))
    v5 = (v4 + "foo bar")
    if (true || (!true)) {
        _ = len(v3)
        var v8 uint64 = 978
        var v9 uint = (v1 / ((v1) | 1))
        _ = v8
        _ = v9
    } else {
        v5 = v5
        var v10 string = ("" + v4)
        _ = v10
    }
    v1 += (^(v1))
    var v11 int32 = 1
    fmt.Println(v5, v1)
    v12 := GenS1{F0: v4, F1: (v11 | (v11 - -314))}
    v2[4294967295] = 715
    v1 -= (v1 >> (((v1 << (((^(v1))) & 63))) & 63))
    v2[2147483647] = 0
    if true {
        for fi13 := 0; fi13 < 16; fi13++ {
            switch (v5 + v4) {
            case "foo bar":
                if (0 <= -1) {
                    var v14 uint = (^(v1))
                    switch v6 {
                    case "foo bar":
                        v4 = v6
                    case "\n\t":
                        var v15 uint = (v14 >> ((v14) & 63))
                        _ = v15
                    case "":
                        var v16 uint = v14
                        _ = v16
                    default:
                        for fi17 := 0; fi17 < 2; fi17++ {
                            var v18 float32 = 1.5
                            v19 := GenS1{F0: ((v6 + ("" + (v5 + v4))) + v4), F1: 13}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                    for fi20 := 0; fi20 < 15; fi20++ {
                        v6 = ("foo bar" + ((v4 + "") + v4))
                        _ = fi20
                    }
                    _ = v14
                }
            case "hello":
                v21 := GenS1{F0: (v4 + v5), F1: 2147483647}
                _ = v21
            }
            for fi22 := 0; fi22 < 13; fi22++ {
                fmt.Println(v5, fi13, v6)
                _ = fi22
            }
            _ = fi13
        }
        v4 = "hello"
        var v23 uint16 = 0
        _ = v23
    } else {
        v24 := GenS1{F0: "foo bar", F1: (^(v11))}
        var v25 int16 = 780
        v26 := []int{-1, 127, -9223372036854775808, 42}
        _ = v24
        _ = v25
        _ = v26
    }
    v5 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
