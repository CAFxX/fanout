package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint64 = 4294967295
    v3 := make(map[int]string)
    v4 := GenS1{F0: 381, F1: -128}
    v5 := []int{}
    var v6 uint16 = 255
    var v7 uint32 = uint32(v6)
    if len(v5) > 0 {
        v5[0] = 32767
    }
    v2 += v2
    v8 := GenS2{F0: 182}
    v2 -= 18446744073709551615
    v1 = 9223372036854775807
    for fi9 := 0; fi9 < 11; fi9++ {
        v7 = v7
        var v10 uint8 = 1
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 float32 = (3.14159 + (1e10 / 1.0))
            v13 := v3[32767]
            switch (v13 + "a") {
            case "":
                var v14 uint = 255
                _ = v14
            case "a":
                if (599 > 2147483647) {
                    var v15 bool = true
                    switch v13 {
                    case "\n\t":
                        switch v13 {
                        case "hello":
                            v5 = append(v5, (fi9 % (((fi9 | 128)) | 1)))
                        default:
                            v6 = 32767
                            v16 := GenS2{F0: fi9}
                            _ = v16
                        }
                    default:
                        v17 := GenS3{F0: (v15 || (!(9223372036854775807 < (fi9 * (-(fi11))))))}
                        _ = v17
                    }
                    _ = v15
                }
            }
            var v18 uint64 = (v2 ^ 0)
            _ = fi11
            _ = v12
            _ = v13
            _ = v18
        }
        _ = fi9
        _ = v10
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v6 = (v6 | (v6 << ((455) & 63)))
        for fi20 := 0; fi20 < 10; fi20++ {
            v2 = (v2 + (^(v2)))
            v1 = (-(v1))
            _ = fi20
        }
        v1 += 9223372036854775807
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
