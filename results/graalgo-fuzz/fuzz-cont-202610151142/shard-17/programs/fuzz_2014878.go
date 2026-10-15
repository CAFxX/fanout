package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint8 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []uint32{}
    var v2 uint16 = 32767
    v3 := []string{}
    v4 := make(map[uint32]string)
    var v5 int32 = -1
    _ = len(v4)
    v2 = v2
    var v6 uint16 = v2
    if (0 >= 32767) {
        for fi7 := 0; fi7 < 7; fi7++ {
            if false {
                v8 := GenS1{F0: 127, F1: (!true), F2: 42}
                _ = v8
            } else {
                if ((v5 >> ((65535) & 63)) >= v5) {
                    var v9 string = "hello"
                    switch -85 {
                    case 376:
                        _ = len(v4)
                    }
                    _ = v9
                } else {
                    v5 += (v5 % ((v5) | 1))
                    for fi10 := 0; fi10 < 1; fi10++ {
                        v11 := GenS2{F0: v5, F1: -745}
                        v2 = 1
                        _ = len(v4)
                        _ = fi10
                        _ = v11
                    }
                }
            }
            _ = fi7
        }
        v12 := GenS1{F0: 31, F1: (-1 < 0), F2: 128}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 15; fi13++ {
            v5 += -189
            fmt.Println(v5)
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS3{F0: uint32(v6)}
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 uint16 = (v2 | 2)
            v2 = (v6 % ((255) | 1))
            fmt.Println(fi16)
            _ = fi16
            _ = v17
        }
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        v2 = v6
    case 22:
        v2 = (v2 | v6)
    }
    var v18 uint8 = 115
    v3 = append(v3, "foo bar")
    for fi19 := 0; fi19 < 16; fi19++ {
        fmt.Println(v6, v6)
        _ = fi19
    }
    v5 += v5
    var v20 uint = 4294967295
    v4[4294967295] = "\n\t"
    if true {
        var v21 bool = true
        v22 := GenS3{F0: 0}
        _ = v21
        _ = v22
    }
    var v23 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
