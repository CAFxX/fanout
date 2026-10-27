package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int }
type GenS2 struct { F0 int16; F1 int8; F2 bool }
type GenS3 struct { F0 int64; F1 int32 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]bool)
    v3 := []int{32768, -798}
    v1 -= (^(v1))
    for fi4 := 0; fi4 < 20; fi4++ {
        if len(v3) > 0 {
            v3[0] = (fi4 / (((fi4 << ((65535) & 63))) | 1))
        }
        for fi5 := 0; fi5 < 13; fi5++ {
            v1 -= v1
            _ = len(v2)
            var v6 uint8 = 64
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v7 := GenS1{F0: 4294967295, F1: 234, F2: 128}
    var v8 uint64 = 1
    var v9 string = ""
    var v10 float64 = 1e10
    switch 4294967295 {
    case 2147483647:
        var v11 int16 = 1000
        _ = v11
    case -1:
        v9 = "a"
    default:
        if (((!((v1 + (v1 + v1)) >= v1)) || (1 > 218)) && (32768 < 42)) {
            fmt.Println(v8, v9, v10)
            v3 = append(v3, 2147483647)
            if len(v3) > 0 {
                v3[0] = 0
            }
        }
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2[2147483647] = ((v8 <= 2) && true)
            var v14 int = 9223372036854775807
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: 2, F1: (v8 - 93), F2: fi12}
        delete(v2, 1)
        _ = fi12
        _ = v15
    }
    if (!true) {
        fmt.Println(v1)
    } else {
        if false {
            v16 := GenS2{F0: 824, F1: 99, F2: true}
            _ = v16
        } else {
            if true {
                delete(v2, 369)
                var v17 float64 = v10
                var v18 bool = (v10 > -2.25)
                _ = v17
                _ = v18
            } else {
                for fi19 := 0; fi19 < 14; fi19++ {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v8 = 0
                        _ = fi20
                    }
                    _ = fi19
                }
                delete(v2, 65535)
            }
            fmt.Println(v9)
            switch ((v9 + v9) + ((((v9 + "a") + (v9 + "\n\t")) + (v9 + v9)) + v9)) {
            case "foo bar":
                for fi21 := 0; fi21 < 3; fi21++ {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        var v23 int8 = -62
                        v24 := GenS3{F0: 0, F1: 42}
                        v25 := v2[uint32(v8)]
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    v8 += (^(v8))
                    _ = fi21
                }
            case "hello":
                var v26 uint32 = 2
                _ = v26
            default:
                var v27 uint = 577
                _ = v27
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
