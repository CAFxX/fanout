package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8; F1 int32 }

func main() {
    var v1 uint16 = 914
    v2 := []uint32{}
    var v3 string = ("hello" + "\n\t")
    v4 := []uint32{}
    for fi5 := 0; fi5 < 15; fi5++ {
        if (int32(fi5) >= 2147483647) {
            for fi6 := 0; fi6 < 7; fi6++ {
                v3 = ((v3 + v3) + v3)
                if ((fi6 >> ((255) & 63)) > fi6) {
                    fmt.Println(fi5, v3)
                    var v7 float32 = (0.0 + 0.0)
                    _ = v7
                }
                _ = fi6
            }
        }
        switch (v3 + "hello") {
        case "hello":
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 uint16 = (v1 / ((v1) | 1))
                v10 := GenS1{F0: (fi5 ^ 32768), F1: 1}
                _ = fi8
                _ = v10
                _ = v9
            }
        default:
            for fi11 := 0; fi11 < 14; fi11++ {
                v1 = 32767
                _ = fi11
            }
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            v1 = (v1 / ((2) | 1))
            _ = fi12
        }
        _ = fi5
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint16 = v1
        if (1000 >= -32768) {
            v4 = append(v4, 137)
            v3 = (v3 + (((v3 + v3) + v3) + (v3 + v3)))
            fmt.Println(v3, fi13, v3)
        }
        _ = fi13
        _ = v14
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        _ = len(v4)
        _ = fi15
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        if true {
            if len(v2) > 0 {
                v2[0] = 0
            }
            fmt.Println(v3)
            v3 = "foo bar"
        } else {
            v17 := GenS2{F0: 1, F1: 1}
            v1 -= 402
            var v18 uint = 379
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    var v19 int32 = -2147483648
    fmt.Println(v19)
    var v20 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
