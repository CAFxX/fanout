package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int8; F1 int8; F2 int16 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint = 1
    v2 := []uint32{4294967295, 1}
    var v3 int16 = 0
    var v4 uint8 = 31
    var v5 int64 = 0
    v6 := make(map[uint32]uint32)
    v3 = (v3 >> ((v1) & 63))
    if false {
        if true {
            for fi7 := 0; fi7 < 20; fi7++ {
                switch ("\n\t" + "foo bar") {
                case "hello":
                    for fi8 := 0; fi8 < 6; fi8++ {
                        v4 = v4
                        var v9 uint64 = 18446744073709551615
                        _ = fi8
                        _ = v9
                    }
                case "a":
                    var v10 string = ("hello" + (((("hello" + (((("\n\t" + "") + (("\n\t" + "foo bar") + ("\n\t" + "a"))) + "\n\t") + "")) + "\n\t") + "\n\t") + ("\n\t" + "hello")))
                    _ = v10
                }
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 uint32 = 2
                    var v13 float32 = ((1e10 - 0.0) + 0.1)
                    v6[v12] = v12
                    var v14 uint16 = 1
                    v4 -= uint8(v1)
                    v12 = v12
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi7
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 uint64 = 1
                v16 = (v16 & 1)
                _ = fi15
                _ = v16
            }
            var v17 uint16 = 32767
            _ = v17
        }
    } else {
        var v18 float32 = (0.1 - 0.1)
        v1 -= 0
        _ = v18
    }
    v3 = v3
    if ((float32((1e10 + -2.25)) - ((0.1 * 0.0) * 3.14159)) != (-2.25 - 0.0)) {
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 int64 = -1
            _ = fi19
            _ = v20
        }
        var v21 uint = (v1 + v1)
        var v22 uint = (v1 & 18446744073709551615)
        _ = v21
        _ = v22
    } else {
        var v23 float64 = 1e10
        v24 := GenS1{F0: 456, F1: 236}
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: int8(v5), F1: -85, F2: (v3 ^ (v3 >> (((v1 / (((v1 ^ v1)) | 1))) & 63)))}
    var v26 uint16 = 0
    switch 32767 {
    case 9223372036854775807:
        for fi27 := 0; fi27 < 20; fi27++ {
            if true {
                var v28 int16 = (v3 & 0)
                _ = v28
            }
            _ = fi27
        }
    }
    if (255 <= v1) {
        switch 1 {
        case 9223372036854775807:
            v3 = 1
        case 127:
            v29 := GenS2{F0: -1, F1: 112, F2: (v3 % ((-64) | 1))}
            _ = v29
        case -64:
            var v30 int = -2147483648
            _ = v30
        }
        switch -273 {
        case 0:
            if true {
                for fi31 := 0; fi31 < 1; fi31++ {
                    var v32 uint64 = 2
                    v4 = (^(v4))
                    v33 := GenS3{F0: fi31}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                switch 1 {
                case 9223372036854775807:
                    v34 := GenS2{F0: -128, F1: 113, F2: 42}
                    _ = v34
                case -502:
                    var v35 uint8 = (^(v4))
                    _ = v35
                case -9223372036854775808:
                    v1 = v1
                default:
                    var v36 uint16 = 42
                    _ = v36
                }
                for fi37 := 0; fi37 < 20; fi37++ {
                    var v38 int = (^(fi37))
                    if false {
                        v4 -= v4
                    }
                    var v39 int8 = 0
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
        case 128:
            fmt.Println(v3, v1)
        case 2147483647:
            fmt.Println(v5, v3)
        }
        v2 = append(v2, 2)
    }
    fmt.Println(v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
