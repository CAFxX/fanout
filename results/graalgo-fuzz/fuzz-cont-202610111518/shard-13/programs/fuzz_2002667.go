package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := []uint32{65535}
    v4 := make(map[uint32]bool)
    var v5 uint32 = 65535
    v6 := []uint32{2147483647, v5, v5}
    switch "foo bar" {
    case "foo bar":
        if false {
            switch "" {
            case "hello":
                var v7 int8 = -67
                _ = v7
            default:
                if true {
                    v8 := GenS1{F0: 681}
                    _ = v8
                } else {
                    v5 -= v5
                    for fi9 := 0; fi9 < 6; fi9++ {
                        var v10 uint = 22
                        _ = fi9
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 9; fi11++ {
                        v5 -= 1
                        var v12 string = ("a" + "hello")
                        _ = fi11
                        _ = v12
                    }
                }
            }
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int32 = 975
                v5 += 0
                _ = fi13
                _ = v14
            }
        }
    case "":
        if true {
            var v15 bool = true
            v16 := GenS1{F0: 986}
            v17 := GenS1{F0: 42}
            _ = v15
            _ = v16
            _ = v17
        }
    case "\n\t":
        var v18 int16 = -1
        _ = v18
    default:
        var v19 int64 = -729
        _ = v19
    }
    var v20 uint = 74
    var v21 int16 = -856
    v21 += v21
    fmt.Println(v21, v21, v21)
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS1{F0: 18446744073709551615}
        var v24 int8 = 12
        _ = fi22
        _ = v23
        _ = v24
    }
    v21 -= 427
    for fi25 := 0; fi25 < 14; fi25++ {
        fmt.Println(v21, v20)
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
