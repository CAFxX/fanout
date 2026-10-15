package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 bool = false
    var v4 bool = (!(!v3))
    v5 := make(map[string]int)
    v4 = v4
    v6 := v2["\n\t"]
    var v7 int16 = -628
    var v8 float32 = 0.1
    switch (v6 << ((0) & 63)) {
    case 0:
        v4 = v3
    case -9223372036854775808:
        var v9 uint32 = 65535
        _ = v9
    default:
        switch v6 {
        case 855:
            if ((v8 < v8) || (464 <= 0)) {
                var v10 int = (^(v6))
                v11 := GenS1{F0: 0, F1: 0}
                v12 := GenS1{F0: 42, F1: 2}
                _ = v10
                _ = v11
                _ = v12
            } else {
                if true {
                    var v13 int64 = -982
                    v4 = ((v6 != (v6 << ((65535) & 63))) || v3)
                    _ = v13
                } else {
                    var v14 int = v6
                    var v15 uint32 = 42
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 7; fi16++ {
                    if (1000 >= (v7 * (v7 % (((v7 ^ (v7 % ((721) | 1)))) | 1)))) {
                        switch "\n\t" {
                        case "":
                            v4 = false
                        case "\n\t":
                            var v17 uint = 2
                            _ = v17
                        case "a":
                            var v18 bool = v4
                            _ = v18
                        }
                        v19 := GenS1{F0: -1, F1: 1}
                        v20 := v2["foo bar"]
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 20; fi21++ {
                        if true {
                            switch (fi16 / (((fi16 + (v6 >> ((0) & 63)))) | 1)) {
                            case 490:
                                v22 := GenS1{F0: 1, F1: 757}
                                _ = v22
                            }
                            v23 := GenS1{F0: -435, F1: 299}
                            _ = v23
                        } else {
                            v24 := []int{32767, (fi21 * v6)}
                            v6 -= fi16
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(v3, v3, v3)
                            default:
                                var v25 int8 = -105
                                _ = v25
                            }
                            _ = v24
                        }
                        _ = fi21
                    }
                    for fi26 := 0; fi26 < 17; fi26++ {
                        v3 = v4
                        var v27 int16 = (v7 << ((0) & 63))
                        _ = fi26
                        _ = v27
                    }
                    _ = fi16
                }
            }
        case 0:
            v28 := GenS1{F0: 1, F1: 18446744073709551615}
            _ = v28
        default:
            v29 := []int{}
            _ = v29
        }
    }
    if true {
        v30 := GenS1{F0: -670, F1: 743}
        var v31 uint64 = 0
        var v32 string = "a"
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 float64 = 1.5
    v34 := GenS1{F0: -9223372036854775808, F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
