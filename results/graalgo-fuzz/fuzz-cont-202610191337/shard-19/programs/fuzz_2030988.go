package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{1}
    v2 := make(map[uint32]int)
    v3 := []int{1, 9223372036854775807, 128}
    v4 := make(map[int]uint32)
    var v5 int = 873
    v6 := make(map[int]bool)
    v1 = append(v1, 4294967295)
    fmt.Println(v5)
    v5 = (v5 + v5)
    if (32767 <= 65535) {
        fmt.Println(v5)
        switch "a" {
        case "\n\t":
            switch ("hello" + "\n\t") {
            case "foo bar":
                var v7 uint16 = 130
                _ = v7
            case "hello":
                fmt.Println(v5)
            case "a":
                if false {
                    v5 = 0
                } else {
                    v5 = -84
                    for fi8 := 0; fi8 < 4; fi8++ {
                        for fi9 := 0; fi9 < 4; fi9++ {
                            if (!(-32768 != 284)) {
                                _ = len(v6)
                            } else {
                                v5 = -1
                                v5 += 32767
                                fmt.Println(fi9, v5)
                            }
                            _ = fi9
                        }
                        _ = fi8
                    }
                }
            default:
                var v10 uint16 = 1
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 uint64 = 186
            for fi13 := 0; fi13 < 5; fi13++ {
                v12 += v12
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
    } else {
        v5 = v5
        v14 := v6[v5]
        _ = v14
    }
    fmt.Println(v5)
    v5 -= (v5 << ((18446744073709551615) & 63))
    v3 = append(v3, (-(v5)))
    v5 = v5
    if true {
        v1 = append(v1, 65535)
        if (3.14159 < (1e10 + 1.5)) {
            fmt.Println(v5)
        } else {
            var v15 int = 1
            delete(v6, (^(v5)))
            _ = v15
        }
    } else {
        var v16 int16 = 622
        _ = v16
    }
    v5 = v5
    v5 += -2147483648
    v17 := v6[(v5 + v5)]
    fmt.Println(v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
