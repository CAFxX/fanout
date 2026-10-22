package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int8; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint }

func main() {
    var v1 int16 = 1000
    var v2 string = "a"
    var v3 bool = (-1 <= 0)
    var v4 uint64 = 619
    var v5 float32 = 3.14159
    v4 = 18446744073709551615
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 0, F1: 605, F2: 709}
        _ = fi6
        _ = v7
    }
    v8 := []int{-817, -722}
    v3 = ((v4 & (v4 / (((v4 + v4)) | 1))) < (v4 >> ((1) & 63)))
    if (253 == 100) {
        var v9 uint = 18446744073709551615
        for fi10 := 0; fi10 < 10; fi10++ {
            switch ("\n\t" + "\n\t") {
            case "hello":
                v11 := GenS2{F0: (v4 * v4), F1: (v9 & v9)}
                _ = v11
            case "":
                switch "hello" {
                case "a":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                case "":
                    if len(v8) > 0 {
                        v8[0] = 42
                    }
                }
            }
            switch (((v2 + v2) + v2) + (v2 + ((v2 + v2) + "foo bar"))) {
            case "":
                v12 := GenS1{F0: 0, F1: 42, F2: 2}
                _ = v12
            case "a":
                v5 = v5
            }
            _ = fi10
        }
        var v13 uint16 = 32767
        _ = v13
        _ = v9
    }
    var v14 float32 = ((v5 / v5) - v5)
    v15 := GenS3{F0: 100, F1: 0, F2: 32767}
    v16 := GenS4{F0: uint(v4)}
    v17 := GenS3{F0: -92, F1: -26, F2: 255}
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(fi18, v14)
        v14 += ((0.1 - -2.25) + (float32(1.5) / v5))
        v5 = ((v5 / float32(1.5)) + -2.25)
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
