package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int32; F2 int8 }

func main() {
    v1 := []string{}
    var v2 uint32 = 42
    var v3 int32 = -1
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]uint32)
    v2 += v2
    v6 := GenS1{F0: 65535, F1: 197}
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            switch "foo bar" {
            case "\n\t":
                var v9 int16 = -1
                _ = v9
            case "a":
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            default:
                v5[v2] = (v2 / ((v2) | 1))
            }
            var v10 int32 = v3
            _ = fi8
            _ = v10
        }
        v11 := GenS2{F0: 32767}
        _ = fi7
        _ = v11
    }
    var v12 bool = false
    var v13 float64 = 1e10
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int16 = 32767
        _ = fi14
        _ = v15
    }
    v1 = append(v1, "foo bar")
    var v16 float32 = float32(v13)
    var v17 uint32 = v2
    var v18 uint = 644
    var v19 int16 = 0
    if len(v1) > 0 {
        v1[0] = ("foo bar" + ("hello" + ""))
    }
    var v20 int = 32767
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v21 uint16 = 42
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 uint16 = 603
        v19 -= (v19 / (((^(v19))) | 1))
        _ = fi22
        _ = v23
    }
    switch v20 {
    case -899:
        if v12 {
            fmt.Println(v2, v21)
            if (v12 && (v13 != 0.1)) {
                if true {
                    v24 := GenS2{F0: v21}
                    v25 := GenS2{F0: (v21 & v21)}
                    _ = v24
                    _ = v25
                }
            } else {
                var v26 int32 = v3
                switch ("a" + ("\n\t" + ("hello" + "hello"))) {
                case "hello":
                    fmt.Println(v19, v12, v19)
                case "\n\t":
                    var v27 uint8 = 103
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 19; fi28++ {
                        var v29 float32 = (float32((v13 - 0.1)) / 1e10)
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        v30 := GenS2{F0: v21}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 int16 = v19
                    if true {
                        fmt.Println(v16)
                    } else {
                        switch 634 {
                        case 4294967295:
                            var v33 bool = false
                            _ = v33
                        default:
                            var v34 int8 = 100
                            _ = v34
                        }
                        delete(v4, (-(fi31)))
                        var v35 uint8 = 2
                        _ = v35
                    }
                    v36 := GenS3{F0: (v18 & v18), F1: 42, F2: 78}
                    _ = fi31
                    _ = v32
                    _ = v36
                }
                _ = v26
            }
        }
    default:
        var v37 int32 = -957
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
