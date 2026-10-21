package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint8 }

func main() {
    var v1 float32 = float32(-2.25)
    var v2 int64 = 0
    var v3 uint = 255
    v4 := []int{}
    v3 = (v3 + 65535)
    fmt.Println(v1, v1, v3)
    _ = len(v4)
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    if len(v4) > 0 {
        v4[0] = -849
    }
    if len(v4) > 0 {
        v4[0] = -735
    }
    if (false && ("foo bar" != "a")) {
        switch 1 {
        case 32768:
            var v5 bool = (!true)
            _ = v5
        default:
            if true {
                v3 += (v3 % ((42) | 1))
                for fi6 := 0; fi6 < 1; fi6++ {
                    var v7 int64 = (-(v2))
                    _ = fi6
                    _ = v7
                }
                v8 := GenS1{F0: 2, F1: ("foo bar" + ("a" + "")), F2: 42}
                _ = v8
            }
        }
    }
    v4 = append(v4, -2147483648)
    var v9 uint32 = 65535
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS1{F0: 559, F1: "hello", F2: 128}
        v12 := GenS1{F0: 344, F1: ("a" + "hello"), F2: 127}
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 uint8 = 119
            v3 = (^(v3))
            _ = fi13
            _ = v14
        }
        var v15 string = "\n\t"
        _ = v15
    } else {
        for fi16 := 0; fi16 < 2; fi16++ {
            if (!false) {
                var v17 int64 = (v2 << (((^(v3))) & 63))
                var v18 string = "hello"
                fmt.Println(v18, v3)
                _ = v17
                _ = v18
            }
            _ = fi16
        }
        if len(v4) > 0 {
            v4[0] = 42
        }
    }
    var v19 int = 2147483647
    var v20 uint16 = 2
    v1 += v1
    var v21 int8 = 126
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
