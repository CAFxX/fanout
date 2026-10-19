package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint64; F1 int; F2 int16 }
type GenS3 struct { F0 int16; F1 uint16; F2 int32 }
func genFn0(p0 int64, p1 int64) int16 {
    var v6 int = 581
    for fi7 := 0; fi7 < 3; fi7++ {
        switch fi7 {
        case 32768:
            p1 -= p0
        case 4294967295:
            v8 := GenS1{F0: -1, F1: ("\n\t" + "")}
            _ = v8
        case 9223372036854775807:
            v9 := GenS2{F0: 52, F1: 127, F2: 32767}
            _ = v9
        default:
            v10 := GenS3{F0: -556, F1: 539, F2: 0}
            _ = v10
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return int16(v6)
}

func main() {
    v1 := []uint32{2147483648}
    v2 := []string{"\n\t", ""}
    var v3 int16 = -1
    var v4 uint64 = 18446744073709551615
    v5 := []uint32{}
    var v11 uint = 917
    v12 := GenS1{F0: 1, F1: (((("hello" + ("a" + "")) + ("" + ("a" + ("\n\t" + "\n\t")))) + "a") + "hello")}
    v3 = v3
    fmt.Println(v4)
    switch 128 {
    case -9223372036854775808:
        fmt.Println(v3)
    case 412:
        var v13 uint8 = 128
        _ = v13
    case 0:
        if false {
            v14 := GenS1{F0: 42, F1: "\n\t"}
            _ = v14
        }
    default:
        if false {
            var v15 uint64 = v4
            _ = v15
        }
    }
    v4 = v4
    v3 = v3
    switch 9223372036854775807 {
    case 128:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 0:
        if (1e10 < -2.25) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v11 += 1
            for fi16 := 0; fi16 < 20; fi16++ {
                fmt.Println(v11, v11)
                _ = fi16
            }
            v4 = v4
        }
    case -1:
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 int32 = 0
            switch "" {
            case "hello":
                if (true && (-1 <= -1)) {
                    var v19 int16 = v3
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v19
                }
            case "a":
                v3 += v3
            case "foo bar":
                v1 = append(v1, 600)
            }
            _ = fi17
            _ = v18
        }
    default:
        var v20 uint = v11
        _ = v20
    }
    fmt.Println(v11, v3, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
