package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int16 }

func main() {
    v1 := []string{("foo bar" + ("foo bar" + "")), "hello", "foo bar"}
    var v2 uint8 = 1
    var v3 float32 = (-2.25 + 1.5)
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: (v2 << ((18446744073709551615) & 63))}
    v2 += (v2 << ((255) & 63))
    var v8 int64 = -35
    v8 += (v8 / ((v8) | 1))
    switch "hello" {
    case "hello":
        switch "hello" {
        case "foo bar":
            var v9 float32 = v3
            _ = v9
        case "\n\t":
            v10 := GenS2{F0: "", F1: -678}
            _ = v10
        case "":
            if false {
                if (v2 <= (v2 ^ 1)) {
                    var v11 uint64 = 0
                    v12 := GenS2{F0: "foo bar", F1: 760}
                    for fi13 := 0; fi13 < 7; fi13++ {
                        if ((false && (0 <= -32768)) || false) {
                            if (v2 >= (v2 << ((2) & 63))) {
                                v14 := GenS1{F0: v2}
                                v15 := v6[uint32(v11)]
                                _ = v14
                                _ = v15
                            }
                            delete(v4, ("hello" + "foo bar"))
                            fmt.Println(fi13, v2, v3)
                        }
                        _ = fi13
                    }
                    _ = v11
                    _ = v12
                }
            } else {
                v16 := GenS1{F0: (v2 & (v2 ^ (v2 >> ((65535) & 63))))}
                v8 -= -908
                _ = v16
            }
        default:
            v4["a"] = 537
        }
    case "a":
        _ = len(v4)
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS1{F0: v2}
        _ = fi17
        _ = v18
    }
    fmt.Println(v2)
    switch "" {
    case "hello":
        for fi19 := 0; fi19 < 5; fi19++ {
            v5[("hello" + "foo bar")] = ""
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v1)
            _ = fi19
        }
    case "\n\t":
        var v20 int16 = 32767
        _ = v20
    }
    fmt.Println(v8, v8, v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
