package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint32; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int32 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]bool)
    var v3 int16 = 32767
    var v4 uint64 = 313
    if true {
        var v5 bool = false
        var v6 uint8 = 2
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 748}
    v1 = v1
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint32 = 2
        v3 = v3
        v10 := v2[65535]
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 uint16 = 0
        _ = fi11
        _ = v12
    }
    v1 = 0.1
    switch 42 {
    case -1:
        fmt.Println(v4, v1)
    case 128:
        fmt.Println(v3, v4, v4)
    case 9223372036854775807:
        if (true || (v1 == 1e10)) {
            switch "a" {
            case "\n\t":
                v3 -= (v3 + 256)
            case "hello":
                switch -2147483648 {
                case 195:
                    var v13 uint = 809
                    _ = v13
                default:
                    switch 32767 {
                    case 42:
                        var v14 int16 = (v3 + -32768)
                        _ = v14
                    case 32768:
                        v4 = (^(v4))
                    }
                }
            case "":
                for fi15 := 0; fi15 < 13; fi15++ {
                    v3 += (v3 & v3)
                    v16 := GenS2{F0: ("" + "\n\t"), F1: 2, F2: (true || true)}
                    switch "\n\t" {
                    case "foo bar":
                        v17 := []int{}
                        _ = v17
                    case "a":
                        for fi18 := 0; fi18 < 1; fi18++ {
                            v2[341] = (937 != -1)
                            var v19 int = fi15
                            var v20 bool = false
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    default:
                        var v21 uint64 = v4
                        _ = v21
                    }
                    _ = fi15
                    _ = v16
                }
            default:
                v4 -= (^(v4))
            }
            if ((((3.14159 - -2.25) - (3.14159 * float64(v1))) <= (((0.1 - 3.14159) + (0.0 / 1.5)) / 0.1)) && ("\n\t" <= "")) {
                var v22 int8 = 100
                var v23 uint64 = (v4 | (v4 & 4294967295))
                _ = v22
                _ = v23
            } else {
                _ = len(v2)
            }
            var v24 uint32 = 921
            _ = v24
        } else {
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 int32 = -721
                _ = fi25
                _ = v26
            }
            v27 := GenS3{F0: 65535}
            v4 = (^(v4))
            _ = v27
        }
    default:
        v28 := GenS1{F0: v3}
        _ = v28
    }
    v29 := []int{363, 4294967295}
    v30 := GenS4{F0: -1, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v7
}
