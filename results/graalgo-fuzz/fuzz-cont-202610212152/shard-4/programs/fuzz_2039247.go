package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := []int{1, 4294967295, 32767, 4294967295}
    v2 := make(map[int]int)
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]int)
    var v5 uint = 1
    v6 := []int{367, 128, 1, -9223372036854775808}
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    var v7 uint = (v5 % ((v5) | 1))
    _ = len(v1)
    v8 := GenS1{F0: 42, F1: 32767}
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 float64 = 1.5
        _ = fi9
        _ = v10
    }
    var v11 bool = true
    for fi12 := 0; fi12 < 20; fi12++ {
        v5 = (v5 | 65535)
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS2{F0: -752}
            _ = fi13
            _ = v14
        }
        v5 = (v5 ^ (v5 + v7))
        _ = fi12
    }
    var v15 uint = 711
    if false {
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(fi16)
            var v17 uint = (v5 - (v7 << ((255) & 63)))
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            v5 += (^(v15))
            if v11 {
                var v20 int64 = 9223372036854775807
                delete(v2, (fi19 ^ (-(fi18))))
                var v21 uint32 = 0
                _ = v20
                _ = v21
            }
            _ = fi19
        }
        fmt.Println(v11, fi18)
        v7 -= 2
        _ = fi18
    }
    if len(v1) > 0 {
        v1[0] = 127
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        delete(v2, (-(fi22)))
        v3 = v3
        switch -764 {
        case -669:
            for fi23 := 0; fi23 < 3; fi23++ {
                if ((1e10 + float64((1e10 + 1.5))) > (3.14159 * (1.5 + 1.5))) {
                    switch 9223372036854775807 {
                    case 1:
                        switch fi22 {
                        case 634:
                            var v24 float32 = 1.5
                            _ = v24
                        }
                    }
                } else {
                    v7 = 2
                    var v25 int64 = -1
                    v3 = v3
                    v26 := GenS3{F0: 0, F1: 187}
                    _ = v25
                    _ = v26
                }
                switch "foo bar" {
                case "a":
                    fmt.Println(v7, v7)
                }
                v27 := v4[9223372036854775807]
                _ = fi23
                _ = v27
            }
        default:
            v28 := GenS1{F0: 671, F1: 42}
            _ = v28
        }
        _ = fi22
    }
    switch "" {
    case "hello":
        for fi29 := 0; fi29 < 14; fi29++ {
            v3 += (v3 & (^(v3)))
            _ = fi29
        }
    case "a":
        if true {
            v7 = 42
            var v30 uint = 794
            _ = v30
        }
    case "\n\t":
        var v31 int = 4294967295
        _ = v31
    }
    v3 = v3
    v32 := v2[259]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
