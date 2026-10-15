package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int = 49
    v2 := make(map[uint32]string)
    var v3 bool = (!false)
    v4 := make(map[uint32]bool)
    v5 := []uint32{2147483647, 2}
    var v6 int64 = 0
    v7 := GenS1{F0: 669}
    if false {
        switch "hello" {
        case "hello":
            for fi8 := 0; fi8 < 16; fi8++ {
                fmt.Println(v3, v1, v1)
                var v9 int32 = 2147483647
                _ = fi8
                _ = v9
            }
        case "foo bar":
            switch ("" + ("\n\t" + "hello")) {
            case "":
                if (2 <= 571) {
                    switch (^(v1)) {
                    case 2147483647:
                        var v10 bool = v3
                        _ = v10
                    case 42:
                        v1 = -277
                    case -9223372036854775808:
                        switch "hello" {
                        case "foo bar":
                            var v11 uint32 = 1
                            _ = v11
                        }
                    }
                }
            case "foo bar":
                fmt.Println(v6, v1)
            default:
                if (!v3) {
                    fmt.Println(v6, v3)
                }
            }
        }
        fmt.Println(v6)
        v5 = append(v5, 292)
    }
    v3 = (0.0 != 0.0)
    for fi12 := 0; fi12 < 20; fi12++ {
        v3 = v3
        _ = fi12
    }
    v6 -= (v6 * v6)
    v4[42] = v3
    switch "" {
    case "a":
        v3 = v3
    case "foo bar":
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 int16 = -171
            _ = fi13
            _ = v14
        }
    case "\n\t":
        v1 = v1
    }
    _ = len(v2)
    v1 -= (v1 >> ((18446744073709551615) & 63))
    v3 = ((v6 > (v6 % (((v6 >> ((807) & 63))) | 1))) && (!(!(!true))))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
