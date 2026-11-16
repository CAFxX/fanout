package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 3.14159
    v3 := []uint32{126, 4294967295, 1}
    var v4 string = "\n\t"
    var v5 float32 = -2.25
    v2 = (1.5 / 1.0)
    v4 = (("foo bar" + (("\n\t" + "\n\t") + "\n\t")) + "hello")
    switch "foo bar" {
    case "foo bar":
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint = 2
            v5 = (((v5 + 1.5) + v5) + 0.0)
            switch -2147483648 {
            case 128:
                var v8 int16 = 32767
                _ = v8
            case 9223372036854775807:
                _ = len(v3)
            case 1:
                var v9 int64 = -1
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    case "\n\t":
        v3 = append(v3, 695)
    default:
        var v10 uint = 42
        _ = v10
    }
    var v11 int8 = 1
    v12 := GenS1{F0: -683}
    switch -734 {
    case 2147483647:
        if true {
            switch -847 {
            case -1:
                switch (("a" + "hello") + (v4 + v4)) {
                case "foo bar":
                    for fi13 := 0; fi13 < 2; fi13++ {
                        var v14 uint8 = 128
                        _ = fi13
                        _ = v14
                    }
                case "hello":
                    v15 := v1[0]
                    _ = v15
                case "":
                    v5 -= v5
                default:
                    for fi16 := 0; fi16 < 17; fi16++ {
                        var v17 int32 = 608
                        _ = fi16
                        _ = v17
                    }
                }
            case 128:
                for fi18 := 0; fi18 < 15; fi18++ {
                    if (42 == 42) {
                        v2 = float64(((3.14159 / v5) - v5))
                        v5 += -2.25
                        var v19 uint16 = 0
                        _ = v19
                    }
                    var v20 int = (fi18 ^ fi18)
                    _ = fi18
                    _ = v20
                }
            case -467:
                var v21 bool = false
                _ = v21
            default:
                var v22 int32 = -2147483648
                _ = v22
            }
            v5 += 3.14159
            var v23 uint32 = 65535
            _ = v23
        }
    case -356:
        var v24 int = 2147483647
        _ = v24
    case -9223372036854775808:
        var v25 int64 = 38
        _ = v25
    default:
        v11 = (v11 | v11)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
