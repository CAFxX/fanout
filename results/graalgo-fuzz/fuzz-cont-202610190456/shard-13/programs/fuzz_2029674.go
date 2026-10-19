package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }

func main() {
    var v1 uint = 65535
    v2 := []uint32{2147483648}
    var v3 int64 = 1
    v4 := make(map[int]int)
    v5 := []uint32{4294967295}
    v6 := GenS1{F0: v1, F1: 2147483647}
    switch 2147483647 {
    case -386:
        if true {
            if false {
                var v7 int16 = 0
                var v8 uint8 = 255
                _ = v7
                _ = v8
            }
        }
    case 998:
        var v9 int16 = 0
        _ = v9
    case 2147483647:
        var v10 bool = true
        _ = v10
    default:
        v1 = 18446744073709551615
    }
    v1 = v1
    v1 = (v1 / (((v1 * v1)) | 1))
    if (!true) {
        if true {
            for fi11 := 0; fi11 < 5; fi11++ {
                switch "a" {
                case "hello":
                    if true {
                        v3 = 42
                        var v12 uint16 = 42
                        v13 := GenS1{F0: 2, F1: 915}
                        _ = v12
                        _ = v13
                    }
                case "foo bar":
                    fmt.Println(v3, fi11, v3)
                }
                if false {
                    v1 += (v1 + (v1 / ((v1) | 1)))
                    var v14 bool = true
                    v15 := GenS1{F0: v1, F1: 0}
                    _ = v14
                    _ = v15
                } else {
                    if false {
                        var v16 uint16 = 1
                        _ = v16
                    } else {
                        v17 := GenS1{F0: 447, F1: 80}
                        v18 := GenS2{F0: (fi11 ^ (^(fi11)))}
                        switch (fi11 ^ (fi11 * -9223372036854775808)) {
                        case -1:
                            v3 -= v3
                        default:
                            v3 -= 0
                        }
                        _ = v17
                        _ = v18
                    }
                }
                for fi19 := 0; fi19 < 15; fi19++ {
                    var v20 uint16 = 0
                    v20 = v20
                    v5 = append(v5, 483)
                    var v21 uint8 = 32
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = fi11
            }
            fmt.Println(v3, v1)
            var v22 uint64 = 1
            _ = v22
        } else {
            fmt.Println(v3)
        }
    }
    _ = len(v4)
    switch 128 {
    case 0:
        v23 := GenS2{F0: -754}
        _ = v23
    case 42:
        switch "" {
        case "foo bar":
            var v24 float32 = (3.14159 * 0.0)
            _ = v24
        default:
            v3 = v3
        }
    case -1:
        v25 := GenS3{F0: ("\n\t" + "a")}
        _ = v25
    }
    var v26 int8 = 100
    fmt.Println(v26, v26)
    fmt.Println(v1, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
