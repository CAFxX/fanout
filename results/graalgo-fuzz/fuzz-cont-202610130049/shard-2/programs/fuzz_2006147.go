package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{}
    v3 := []int{-1}
    for fi4 := 0; fi4 < 15; fi4++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v5 int = (fi4 >> ((65535) & 63))
        _ = fi4
        _ = v5
    }
    v1 = v1
    v3 = append(v3, -1)
    var v6 uint = 389
    v3 = append(v3, -2147483648)
    if (897 <= 553) {
        var v7 int16 = 1
        _ = v7
    }
    v8 := GenS1{F0: 995, F1: (!true)}
    v6 -= 1
    if (true && (65535 <= (^(v6)))) {
        for fi9 := 0; fi9 < 2; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 bool = false
                if len(v2) > 0 {
                    v2[0] = ("hello" + ("" + (("\n\t" + "hello") + ("\n\t" + "\n\t"))))
                }
                v11 = v11
                if len(v2) > 0 {
                    v2[0] = ""
                }
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        var v12 float64 = ((1e10 + 0.0) / 1.5)
        _ = v12
    } else {
        switch 2147483647 {
        case 381:
            var v13 int32 = 0
            _ = v13
        case 127:
            v14 := GenS1{F0: -875, F1: (2 > 127)}
            _ = v14
        default:
            if len(v3) > 0 {
                v3[0] = 512
            }
        }
    }
    var v15 int32 = 1
    var v16 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v8
}
