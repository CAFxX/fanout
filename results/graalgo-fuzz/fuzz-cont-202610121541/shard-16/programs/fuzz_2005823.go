package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int; F1 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 int16 = -32768
    var v4 int8 = -98
    var v5 uint16 = 466
    v6 := GenS1{F0: "", F1: "foo bar"}
    var v7 int = -924
    v7 = v7
    _ = len(v1)
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int = 32768
        var v10 bool = (true && (1.5 <= 1e10))
        v11 := GenS2{F0: 1, F1: ""}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(v7)
    switch (v7 ^ -2147483648) {
    case 9223372036854775807:
        switch "a" {
        case "":
            if (true && ((-(v4)) >= 0)) {
                if (2 != 2147483647) {
                    var v12 uint32 = 65535
                    _ = v12
                }
            }
        default:
            if false {
                v5 += (v5 >> ((183) & 63))
            } else {
                var v13 uint32 = 42
                var v14 bool = true
                _ = v13
                _ = v14
            }
        }
    case -605:
        if ((v7 << ((65535) & 63)) > (-(v7))) {
            v15 := []int{}
            switch v7 {
            case 241:
                switch v7 {
                case 42:
                    v4 += v4
                default:
                    if false {
                        if true {
                            v5 = 255
                            fmt.Println(v5)
                            var v16 int32 = 714
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 7; fi17++ {
                            fmt.Println(fi17, fi17, v7)
                            var v18 int = (v7 << ((2) & 63))
                            var v19 uint64 = 387
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                }
            default:
                if false {
                    v7 = v7
                    fmt.Println(v3, v3)
                    var v20 float32 = ((1e10 / -2.25) - ((float32((-2.25 + 3.14159)) + -2.25) * 0.0))
                    _ = v20
                }
            }
            _ = v15
        }
    case -9223372036854775808:
        v21 := v2[65535]
        _ = v21
    default:
        var v22 uint64 = 18446744073709551615
        _ = v22
    }
    switch "" {
    case "a":
        var v23 uint8 = 0
        _ = v23
    case "hello":
        v5 += (^(v5))
    case "\n\t":
        v24 := []int{v7, (^(v7))}
        _ = v24
    default:
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 string = ("" + "hello")
            if ((!(!false)) || false) {
                var v27 uint32 = 536
                for fi28 := 0; fi28 < 15; fi28++ {
                    var v29 float32 = (float32(float64((1.5 - 0.0))) + (1.5 - float32(0.1)))
                    v30 := GenS3{F0: (-(v7)), F1: fi28}
                    v31 := GenS1{F0: v26, F1: v26}
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                _ = v27
            } else {
                if (-1 >= -1) {
                    var v32 int16 = (v3 - (v3 | v3))
                    _ = v32
                }
                v2[1] = true
            }
            _ = fi25
            _ = v26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
