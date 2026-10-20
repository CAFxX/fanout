package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int }
type GenS2 struct { F0 uint16; F1 string; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int16; F2 int }

func main() {
    var v1 bool = (255 < 32767)
    var v2 int8 = 42
    var v3 int8 = (^(v2))
    switch -2147483648 {
    case 355:
        v4 := GenS1{F0: -516, F1: 2147483647, F2: 9223372036854775807}
        _ = v4
    case 2147483647:
        v2 = v3
    }
    v5 := GenS2{F0: 42, F1: "", F2: -2147483648}
    switch ("foo bar" + ("\n\t" + ("hello" + ("hello" + ("hello" + ""))))) {
    case "hello":
        if (!(217 > 32767)) {
            fmt.Println(v1)
            v6 := GenS1{F0: -781, F1: 436, F2: 42}
            _ = v6
        }
    case "\n\t":
        switch 0 {
        case 790:
            v7 := []int{246, 351}
            _ = v7
        default:
            switch 128 {
            case 42:
                v8 := GenS2{F0: 65535, F1: ("" + "hello"), F2: 4294967295}
                _ = v8
            case 128:
                for fi9 := 0; fi9 < 6; fi9++ {
                    if v1 {
                        fmt.Println(v1, v1)
                    } else {
                        v2 -= (v2 % ((1) | 1))
                    }
                    _ = fi9
                }
            case 4294967295:
                switch -1 {
                case -1:
                    v10 := GenS3{F0: 2}
                    _ = v10
                case 1:
                    v1 = (!(-1 == 42))
                case -2147483648:
                    v2 = v2
                }
            default:
                var v11 int32 = 634
                _ = v11
            }
        }
    }
    v12 := GenS4{F0: 42, F1: -32768, F2: 533}
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := []int{969, (fi13 - (fi13 / ((fi13) | 1)))}
        var v15 float32 = ((0.0 / -2.25) + -2.25)
        var v16 uint8 = 27
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 int64 = -9223372036854775808
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 uint = 511
        v20 := []int{(fi18 % (((fi18 ^ fi18)) | 1))}
        _ = fi18
        _ = v19
        _ = v20
    }
    v2 = 31
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v5
}
