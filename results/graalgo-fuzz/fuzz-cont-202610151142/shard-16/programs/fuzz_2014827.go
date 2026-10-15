package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint }
func genFn0(p0 uint16, p1 uint64) uint {
    p0 = (p0 % (((p0 + 797)) | 1))
    var v4 uint = 200
    _ = p0
    _ = p1
    _ = v4
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"a", "hello", ("\n\t" + "\n\t")}
    v3 := make(map[uint32]int)
    if (75 == 42) {
        v5 := GenS1{F0: 1, F1: 0}
        var v6 uint64 = 0
        _ = v5
        _ = v6
    }
    var v7 uint32 = 2
    v7 = (^(v7))
    if (!(!((1.5 >= (0.0 + (-2.25 / (0.1 * 1e10)))) || ("a" == "")))) {
        for fi8 := 0; fi8 < 3; fi8++ {
            switch (fi8 >> ((4294967295) & 63)) {
            case -715:
                fmt.Println(fi8)
            case 98:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 136:
                for fi9 := 0; fi9 < 3; fi9++ {
                    if false {
                        v1[(^(fi9))] = v7
                        switch (fi8 ^ (fi8 / ((-1) | 1))) {
                        case 261:
                            v7 = 1
                        case 128:
                            _ = len(v2)
                        case 9223372036854775807:
                            if true {
                                v2 = append(v2, (("\n\t" + "\n\t") + ("hello" + ("hello" + ((("foo bar" + ("" + ("hello" + ("foo bar" + "")))) + "foo bar") + "a")))))
                            }
                        default:
                            fmt.Println(fi9, fi8, v7)
                        }
                    } else {
                        var v10 uint64 = 42
                        v7 += v7
                        _ = v10
                    }
                    v7 = v7
                    _ = fi9
                }
            default:
                v7 = (v7 * 1)
            }
            var v11 int16 = 526
            var v12 int32 = 0
            _ = fi8
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint = 18446744073709551615
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: 1, F1: 42}
        _ = v15
    } else {
        if true {
            if false {
                switch "" {
                case "hello":
                    v7 -= v7
                default:
                    switch 9223372036854775807 {
                    case -382:
                        fmt.Println(v7)
                    case -2147483648:
                        v7 = (v7 >> ((4294967295) & 63))
                    default:
                        fmt.Println(v7)
                    }
                }
                delete(v1, 42)
            } else {
                var v16 int64 = 9223372036854775807
                var v17 uint8 = 33
                _ = v16
                _ = v17
            }
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int8 = 42
        _ = fi18
        _ = v19
    }
    var v20 int64 = -9223372036854775808
    switch 2147483647 {
    case -22:
        if false {
            var v21 int = 378
            var v22 int16 = -32768
            _ = v21
            _ = v22
        }
    case 2147483647:
        var v23 float32 = 0.0
        _ = v23
    case 9223372036854775807:
        v24 := GenS2{F0: ("a" + ("hello" + ((("\n\t" + ("a" + "a")) + ("hello" + "\n\t")) + "\n\t"))), F1: false}
        _ = v24
    default:
        _ = len(v2)
    }
    var v25 int16 = 42
    v26 := GenS3{F0: 42, F1: 93, F2: 42}
    v7 = (v7 >> ((4294967295) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v7
}
