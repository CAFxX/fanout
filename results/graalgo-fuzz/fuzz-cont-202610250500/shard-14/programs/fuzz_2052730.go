package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint16; F2 bool }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 int32 = 1
    v5 := []string{("" + "hello"), ("" + "hello"), ("\n\t" + "a"), "foo bar"}
    var v6 int64 = -9223372036854775808
    var v7 float64 = 3.14159
    v6 -= (-(v6))
    if (255 == 997) {
        v6 -= (v6 | v6)
        for fi8 := 0; fi8 < 7; fi8++ {
            if true {
                v9 := GenS1{F0: fi8}
                for fi10 := 0; fi10 < 7; fi10++ {
                    fmt.Println(fi8, v6, fi10)
                    var v11 bool = (2 > 2)
                    _ = fi10
                    _ = v11
                }
                var v12 float32 = float32(float64(3.14159))
                _ = v12
                _ = v9
            } else {
                v13 := v1[(("foo bar" + "hello") + "a")]
                for fi14 := 0; fi14 < 8; fi14++ {
                    switch "hello" {
                    case "hello":
                        var v15 float32 = float32(v7)
                        _ = v15
                    case "":
                        var v16 uint = 18446744073709551615
                        _ = v16
                    case "a":
                        v17 := GenS1{F0: (fi14 & (fi8 >> ((42) & 63)))}
                        _ = v17
                    default:
                        var v18 float64 = 0.0
                        _ = v18
                    }
                    v4 = v4
                    _ = fi14
                }
                v6 += int64(v4)
                _ = v13
            }
            _ = fi8
        }
    } else {
        var v19 bool = false
        _ = v19
    }
    fmt.Println(v6, v4, v7)
    v20 := GenS1{F0: 4294967295}
    if false {
        switch -686 {
        case 9223372036854775807:
            if true {
                if (174 != 0) {
                    v21 := GenS1{F0: 0}
                    var v22 int = 292
                    _ = v21
                    _ = v22
                }
                if (2 != 18446744073709551615) {
                    var v23 int8 = 127
                    var v24 bool = (false && false)
                    v7 = v7
                    _ = v23
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 9; fi25++ {
                        if false {
                            var v26 int64 = (v6 / ((v6) | 1))
                            var v27 uint16 = 972
                            _ = v26
                            _ = v27
                        } else {
                            v5 = append(v5, "a")
                            var v28 float32 = 1e10
                            _ = v28
                        }
                        var v29 bool = false
                        _ = fi25
                        _ = v29
                    }
                }
            } else {
                if true {
                    var v30 uint64 = 4294967295
                    var v31 float64 = v7
                    var v32 uint64 = v30
                    _ = v30
                    _ = v31
                    _ = v32
                }
                v33 := GenS2{F0: 42, F1: 65535, F2: false}
                _ = v33
            }
        case 798:
            var v34 int16 = 32767
            _ = v34
        default:
            v6 = (v6 & v6)
        }
    }
    v7 += v7
    for fi35 := 0; fi35 < 19; fi35++ {
        v7 = (v7 - v7)
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
