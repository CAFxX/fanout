package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint8 = 42
    var v2 int64 = -786
    var v3 uint = 0
    var v4 uint32 = 388
    v5 := []int{0, -9223372036854775808}
    v6 := GenS1{F0: (^(v1)), F1: 0, F2: (^(v2))}
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: v1, F1: 9223372036854775807, F2: v2}
        switch "a" {
        case "":
            _ = len(v5)
        }
        if true {
            if (1 > 1) {
                v1 = uint8(v3)
            } else {
                v9 := GenS2{F0: 42}
                _ = v9
            }
            var v10 float32 = 0.0
            _ = v10
        } else {
            var v11 int = 545
            v4 -= v4
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        switch fi12 {
        case 9223372036854775807:
            for fi13 := 0; fi13 < 4; fi13++ {
                if len(v5) > 0 {
                    v5[0] = (^(fi13))
                }
                if true {
                    v2 -= (v2 | v2)
                } else {
                    var v14 uint = (v3 % (((v3 % ((673) | 1))) | 1))
                    var v15 uint64 = 2
                    _ = v14
                    _ = v15
                }
                if (!true) {
                    v4 = (^(v4))
                }
                _ = len(v5)
                fmt.Println(v3, fi13)
                if false {
                    v4 = v4
                    var v16 int16 = 0
                    v17 := GenS1{F0: (v1 ^ (v1 / (((v1 & 210)) | 1))), F1: (v2 + v2), F2: (v2 & v2)}
                    _ = v16
                    _ = v17
                }
                switch (("a" + "") + "hello") {
                case "\n\t":
                    v5 = append(v5, 909)
                case "a":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "hello":
                    var v18 uint32 = (v4 + v4)
                    _ = v18
                default:
                    var v19 float32 = (1e10 + (3.14159 * (1.5 * 3.14159)))
                    _ = v19
                }
                _ = fi13
            }
        case 128:
            for fi20 := 0; fi20 < 2; fi20++ {
                v3 -= (v3 >> (((v3 | v3)) & 63))
                v1 -= v1
                var v21 string = "foo bar"
                _ = fi20
                _ = v21
            }
        case -764:
            v4 -= 2147483647
        }
        _ = fi12
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        v5 = append(v5, fi22)
        _ = fi22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        switch "" {
        case "hello":
            v4 -= v4
        }
        var v24 uint = 427
        _ = fi23
        _ = v24
    }
    v3 = (v3 & (^(v3)))
    var v25 float32 = 1e10
    v3 -= (^(v3))
    if (843 >= 57) {
        for fi26 := 0; fi26 < 19; fi26++ {
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 uint64 = 4294967295
                if false {
                    var v29 int32 = 0
                    var v30 bool = false
                    _ = v29
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
        v25 = 0.1
    }
    switch "a" {
    case "":
        var v31 uint32 = v4
        _ = v31
    case "hello":
        v1 -= 128
    }
    v1 += v1
    var v32 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
