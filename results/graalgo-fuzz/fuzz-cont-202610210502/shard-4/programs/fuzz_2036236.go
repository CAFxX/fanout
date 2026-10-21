package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int32 }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]int)
    var v3 int = 128
    v4 := GenS1{F0: 2, F1: 131, F2: -556}
    var v5 int16 = 235
    fmt.Println(v3, v1)
    switch v3 {
    case -1:
        var v6 int64 = -23
        _ = v6
    case 9223372036854775807:
        var v7 string = ""
        _ = v7
    case 42:
        switch "\n\t" {
        case "foo bar":
            for fi8 := 0; fi8 < 12; fi8++ {
                v2["a"] = 9223372036854775807
                _ = fi8
            }
        case "":
            switch 127 {
            case 9223372036854775807:
                v9 := GenS2{F0: 255}
                _ = v9
            case 4294967295:
                switch "" {
                case "\n\t":
                    if false {
                        fmt.Println(v5, v1, v5)
                        if (!(!((v5 - (v5 | v5)) > (-(v5))))) {
                            var v10 int16 = (v5 / ((-234) | 1))
                            v11 := GenS3{F0: -2147483648}
                            _ = v10
                            _ = v11
                        }
                    }
                case "hello":
                    v12 := GenS1{F0: 2147483647, F1: -1, F2: (v5 * -32768)}
                    _ = v12
                case "":
                    fmt.Println(v1)
                default:
                    var v13 float32 = 3.14159
                    _ = v13
                }
            }
        }
    }
    var v14 int32 = -2147483648
    v14 = v14
    v5 += (v5 + 0)
    fmt.Println(v1)
    switch (("hello" + "\n\t") + "") {
    case "foo bar":
        var v15 int8 = -1
        _ = v15
    case "hello":
        v16 := []int{(v3 ^ v3), (v3 << (((v1 - v1)) & 63)), (v3 >> ((42) & 63))}
        _ = v16
    }
    if (-1 < 100) {
        var v17 int64 = 0
        v17 = (v17 - (v17 << (((v1 % (((v1 / ((0) | 1))) | 1))) & 63)))
        for fi18 := 0; fi18 < 3; fi18++ {
            for fi19 := 0; fi19 < 6; fi19++ {
                v3 = (fi19 % ((fi19) | 1))
                v20 := GenS1{F0: 2, F1: (^(v14)), F2: v5}
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        _ = v17
    }
    v14 = (v14 & (-(v14)))
    v21 := GenS4{F0: 1, F1: (-(v14))}
    v14 = 439
    v22 := []int{int(v1), v3}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
