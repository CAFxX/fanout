package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int16 }

func main() {
    var v1 uint16 = 439
    v2 := []int{}
    var v3 int16 = 32767
    var v4 uint64 = 92
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: (0.1 < 1e10)}
        _ = fi5
        _ = v6
    }
    var v7 string = "foo bar"
    v7 = (("" + "hello") + v7)
    v8 := GenS1{F0: (!false)}
    _ = len(v2)
    fmt.Println(v7, v1, v7)
    v3 = v3
    v1 = v1
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v9 uint16 = (v1 % ((v1) | 1))
    if true {
        v9 -= (v1 ^ (v1 + v9))
        if (!false) {
            for fi10 := 0; fi10 < 13; fi10++ {
                if true {
                    v9 += v9
                    v3 += (v3 / ((-1) | 1))
                    if true {
                        switch ("\n\t" + v7) {
                        case "hello":
                            var v11 bool = (-2.25 <= 1e10)
                            _ = v11
                        }
                        v12 := GenS1{F0: (!(160 <= 0))}
                        _ = v12
                    }
                }
                fmt.Println(v4, fi10)
                v13 := GenS2{F0: fi10, F1: 42}
                _ = fi10
                _ = v13
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                for fi15 := 0; fi15 < 13; fi15++ {
                    fmt.Println(v7)
                    _ = fi15
                }
                _ = fi14
            }
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            if true {
                v17 := GenS3{F0: 202, F1: 0}
                fmt.Println(v3, v1, v9)
                _ = v17
            }
            _ = fi16
        }
    } else {
        if false {
            var v18 int64 = 849
            _ = v18
        } else {
            switch 1 {
            case 9223372036854775807:
                var v19 float32 = -2.25
                _ = v19
            case 519:
                var v20 int = 824
                _ = v20
            default:
                v9 += 321
            }
            var v21 int16 = (v3 % (((v3 & v3)) | 1))
            _ = v21
        }
    }
    if (-1 != 824) {
        for fi22 := 0; fi22 < 11; fi22++ {
            v4 += (v4 ^ (v4 % ((0) | 1)))
            _ = fi22
        }
    }
    v4 = v4
    for fi23 := 0; fi23 < 17; fi23++ {
        if len(v2) > 0 {
            v2[0] = (fi23 % (((fi23 >> ((0) & 63))) | 1))
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
