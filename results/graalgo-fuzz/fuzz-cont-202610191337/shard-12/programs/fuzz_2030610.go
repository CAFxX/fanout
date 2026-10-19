package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int }

func main() {
    var v1 uint16 = 32767
    v2 := []int{}
    v3 := []string{("" + ("hello" + "\n\t"))}
    var v4 uint = 581
    v5 := make(map[uint32]bool)
    v5[831] = true
    v5[4294967295] = (198 == 44)
    v1 -= (v1 + v1)
    var v6 int8 = -1
    v6 = v6
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint = v4
        for fi9 := 0; fi9 < 16; fi9++ {
            _ = len(v5)
            v10 := GenS1{F0: 2, F1: fi7}
            fmt.Println(v6)
            if (42 <= (v6 >> ((v8) & 63))) {
                var v11 int32 = 546
                v12 := GenS1{F0: (^(v1)), F1: (fi7 + (fi9 & int(v4)))}
                var v13 string = ("\n\t" + (("\n\t" + "") + "foo bar"))
                var v14 int8 = -66
                _ = v11
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        v15 := GenS1{F0: v1, F1: (^(fi7))}
        _ = fi7
        _ = v15
        _ = v8
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v4 = v4
        fmt.Println(fi16, v1)
        for fi17 := 0; fi17 < 13; fi17++ {
            v4 = (v4 / (((v4 % ((1) | 1))) | 1))
            _ = fi17
        }
        _ = fi16
    }
    v18 := GenS2{F0: 128}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
