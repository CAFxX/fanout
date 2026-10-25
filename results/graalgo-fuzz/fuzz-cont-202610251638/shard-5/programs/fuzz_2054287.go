package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []string{"", "\n\t", "hello"}
    var v2 uint64 = 1
    var v3 int16 = -1
    v4 := make(map[string]uint32)
    v5 := []uint32{}
    v2 = (v2 % (((^(v2))) | 1))
    var v6 int32 = -2147483648
    var v7 int32 = (v6 | -857)
    var v8 int64 = 9223372036854775807
    var v9 int8 = 10
    switch -9223372036854775808 {
    case 0:
        var v10 bool = false
        _ = v10
    case 32768:
        v6 = v6
    case 4294967295:
        v3 -= (v3 + v3)
    }
    var v11 uint16 = 956
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            switch "" {
            case "foo bar":
                v14 := GenS1{F0: "\n\t"}
                _ = v14
            case "hello":
                if len(v5) > 0 {
                    v5[0] = 65535
                }
            default:
                if true {
                    var v15 int64 = (v8 % ((42) | 1))
                    if true {
                        var v16 bool = ((^(v7)) != (v7 % ((-2147483648) | 1)))
                        v3 = (v3 ^ (v3 & (-(v3))))
                        v8 = v8
                        _ = v16
                    }
                    var v17 int64 = v15
                    var v18 int32 = -1
                    _ = v15
                    _ = v17
                    _ = v18
                }
            }
            v19 := GenS1{F0: ("a" + "a")}
            var v20 uint64 = 18446744073709551615
            _ = fi13
            _ = v19
            _ = v20
        }
        _ = fi12
    }
    fmt.Println(v7, v6, v7)
    var v21 int8 = -128
    v22 := GenS1{F0: ""}
    if ((v3 + v3) != v3) {
        v9 = (v21 & 127)
        v2 -= 18446744073709551615
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
