package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int8 }

func main() {
    var v1 int32 = -566
    var v2 int64 = 581
    var v3 uint = 0
    var v4 uint8 = 150
    var v5 int = 2147483647
    var v6 string = "hello"
    var v7 float32 = 0.1
    v8 := []int{(^(v5)), v5, (-(v5))}
    v9 := GenS1{F0: 1, F1: (v2 >= (v2 << (((v3 & v3)) & 63)))}
    switch (-(v5)) {
    case 4294967295:
        v10 := GenS2{F0: (v4 | (v4 % (((v4 / ((127) | 1))) | 1))), F1: 42, F2: 2}
        _ = v10
    case 32768:
        for fi11 := 0; fi11 < 14; fi11++ {
            if len(v8) > 0 {
                v8[0] = -1
            }
            switch int(v4) {
            case -195:
                var v12 int = (fi11 + 752)
                _ = v12
            case 127:
                switch -1 {
                case 128:
                    v2 = v2
                case 605:
                    v1 = v1
                default:
                    var v13 uint8 = v4
                    _ = v13
                }
            case 32767:
                v8 = append(v8, (^(v5)))
            }
            fmt.Println(v6, v5, v2)
            _ = fi11
        }
    case 128:
        if true {
            if (-749 < -1) {
                switch v6 {
                case "\n\t":
                    var v14 uint8 = uint8(v3)
                    _ = v14
                case "hello":
                    v2 += (v2 * (v2 + 662))
                case "a":
                    var v15 int8 = 42
                    _ = v15
                }
                v16 := GenS2{F0: uint8(v3), F1: 1000, F2: 2}
                var v17 uint64 = 2
                _ = v16
                _ = v17
            } else {
                v7 = float32(-2.25)
            }
        }
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v6 = v6
    if len(v8) > 0 {
        v8[0] = (-(v5))
    }
    var v18 int64 = 1
    var v19 int64 = v18
    fmt.Println(v1, v1)
    v8 = append(v8, v5)
    if false {
        v2 += 9223372036854775807
    } else {
        v7 += v7
        var v20 uint32 = 2147483648
        fmt.Println(v18)
        _ = v20
    }
    if false {
        var v21 uint64 = 889
        switch "foo bar" {
        case "":
            for fi22 := 0; fi22 < 8; fi22++ {
                if false {
                    for fi23 := 0; fi23 < 1; fi23++ {
                        var v24 int = fi23
                        var v25 int8 = 1
                        v26 := GenS3{F0: 0, F1: -128}
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    fmt.Println(fi22, v5, v3)
                } else {
                    for fi27 := 0; fi27 < 5; fi27++ {
                        v21 = 0
                        fmt.Println(v7, v18, fi22)
                        fmt.Println(v7, v3)
                        _ = fi27
                    }
                }
                _ = fi22
            }
        }
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
