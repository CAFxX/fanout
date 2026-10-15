package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := []uint32{1}
    var v2 int = 32767
    v3 := []int{v2, (-(v2)), v2}
    var v4 int32 = 901
    v1 = append(v1, 2147483647)
    v4 -= -1
    var v5 uint = 2
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int16 = -1
            fmt.Println(v2, v2, v8)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if true {
        var v9 uint16 = 0
        _ = v9
    } else {
        v2 = v2
        for fi10 := 0; fi10 < 6; fi10++ {
            v4 = v4
            fmt.Println(v4, v2, v5)
            if true {
                switch fi10 {
                case 32767:
                    fmt.Println(fi10)
                case 1:
                    if ((!(!false)) || false) {
                        v11 := GenS1{F0: false}
                        fmt.Println(v4, v5)
                        _ = v11
                    }
                }
            }
            _ = fi10
        }
        v12 := GenS2{F0: (65535 >= 42)}
        _ = v12
    }
    if len(v1) > 0 {
        v1[0] = 659
    }
    fmt.Println(v4)
    fmt.Println(v5, v2)
    v2 -= (v2 ^ v2)
    switch "" {
    case "foo bar":
        switch (v2 * 0) {
        case 4294967295:
            if (((v2 | 127) <= (v2 % ((127) | 1))) || (0.1 < (1e10 / 1e10))) {
                v13 := GenS2{F0: false}
                v2 = (v2 | (v2 + (v2 & v2)))
                v5 -= v5
                _ = v13
            }
        case 2147483647:
            if len(v1) > 0 {
                v1[0] = 352
            }
        case -395:
            var v14 int = v2
            _ = v14
        }
    }
    if true {
        v15 := GenS3{F0: 42, F1: 860}
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        if (!true) {
            v5 = v5
            v4 += (v4 % (((^(v4))) | 1))
            v4 += -781
        } else {
            fmt.Println(v5, v4)
        }
        v4 -= 0
        var v17 string = "\n\t"
        _ = fi16
        _ = v17
    }
    var v18 int16 = 1000
    v1 = append(v1, 2147483648)
    var v19 int = v2
    if len(v3) > 0 {
        v3[0] = v19
    }
    v1 = append(v1, 416)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
