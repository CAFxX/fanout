package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 string }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int64 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int16 = -32768
    v3 := make(map[int]int)
    v4 := []int{-9223372036854775808, 4294967295}
    var v5 uint32 = 402
    fmt.Println(v2, v2, v2)
    fmt.Println(v2, v2)
    var v6 uint8 = uint8(v5)
    v7 := GenS1{F0: (^(v1)), F1: 255, F2: ""}
    for fi8 := 0; fi8 < 18; fi8++ {
        if true {
            var v9 uint8 = v6
            v10 := GenS2{F0: (1.5 < 1e10), F1: 255}
            switch 570 {
            case 1:
                for fi11 := 0; fi11 < 12; fi11++ {
                    fmt.Println(v2, v6)
                    v12 := GenS3{F0: 4294967295, F1: 2, F2: (^(v1))}
                    _ = fi11
                    _ = v12
                }
            case 198:
                if (!(1 >= 230)) {
                    v2 -= v2
                    var v13 uint8 = 205
                    var v14 bool = false
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS1{F0: -964, F1: 0, F2: "a"}
                    var v16 int32 = 922
                    v5 = 473
                    _ = v15
                    _ = v16
                }
            case 32768:
                for fi17 := 0; fi17 < 16; fi17++ {
                    if false {
                        var v18 int8 = 0
                        _ = v18
                    }
                    _ = fi17
                }
            default:
                switch "a" {
                case "hello":
                    v9 = v6
                }
            }
            _ = v10
            _ = v9
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 float32 = (0.1 * (0.1 - 0.1))
            fmt.Println(v2, v6, v1)
            _ = fi19
            _ = v20
        }
        _ = fi8
    }
    if (!false) {
        v1 = v1
    }
    if ((!true) || true) {
        v5 += v5
    } else {
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch ("a" + ("foo bar" + "hello")) {
            case "hello":
                _ = len(v4)
            case "a":
                var v21 int32 = 42
                _ = v21
            case "foo bar":
                var v22 int16 = v2
                _ = v22
            }
        }
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 uint16 = 255
        _ = fi23
        _ = v24
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
