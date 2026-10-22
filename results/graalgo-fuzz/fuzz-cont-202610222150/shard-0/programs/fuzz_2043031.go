package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -445
    var v3 int64 = 926
    v4 := []int{v2, 128, (v2 << ((1) & 63)), (-(v2))}
    v5 := []int{(v2 + (-(v2))), 0, v2}
    v6 := []string{""}
    var v7 int32 = 0
    v2 -= (v2 >> ((2) & 63))
    var v8 int = v2
    if false {
        v9 := GenS1{F0: 268}
        v7 = v7
        v7 += v7
        _ = v9
    } else {
        switch v2 {
        case 765:
            var v10 uint = 42
            _ = v10
        case 32768:
            fmt.Println(v2)
        case -1:
            if true {
                fmt.Println(v8, v7, v8)
                fmt.Println(v7)
                switch "" {
                case "hello":
                    switch v2 {
                    case 42:
                        switch "foo bar" {
                        case "foo bar":
                            if true {
                                v11 := GenS1{F0: 151}
                                var v12 int16 = 0
                                _ = v11
                                _ = v12
                            } else {
                                for fi13 := 0; fi13 < 16; fi13++ {
                                    var v14 int = (v8 + v8)
                                    v4 = append(v4, v8)
                                    _ = fi13
                                    _ = v14
                                }
                            }
                        case "\n\t":
                            v2 = v8
                        default:
                            v15 := GenS1{F0: 202}
                            _ = v15
                        }
                    default:
                        switch (v2 + (v8 % (((v8 << ((42) & 63))) | 1))) {
                        case -915:
                            for fi16 := 0; fi16 < 16; fi16++ {
                                var v17 string = "\n\t"
                                _ = fi16
                                _ = v17
                            }
                        }
                    }
                case "a":
                    var v18 int16 = -1
                    _ = v18
                case "foo bar":
                    var v19 float64 = 1.5
                    _ = v19
                default:
                    var v20 uint = 42
                    _ = v20
                }
            }
        }
    }
    var v21 uint16 = 32767
    for fi22 := 0; fi22 < 6; fi22++ {
        if (true && false) {
            if (!(287 < 1)) {
                v23 := GenS2{F0: v7}
                _ = v23
            }
            v2 = v8
            v24 := GenS1{F0: 2}
            _ = v24
        }
        fmt.Println(v3, v21, v2)
        delete(v1, "foo bar")
        _ = fi22
    }
    v21 -= v21
    if ((((-(v7)) <= v7) || true) || false) {
        _ = len(v4)
        v7 -= 42
    } else {
        switch 32768 {
        case 9223372036854775807:
            v25 := GenS2{F0: (v7 & (v7 * -1))}
            _ = v25
        case -2147483648:
            var v26 bool = false
            _ = v26
        case 2147483647:
            var v27 uint64 = 18446744073709551615
            _ = v27
        }
    }
    v28 := v1["\n\t"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
