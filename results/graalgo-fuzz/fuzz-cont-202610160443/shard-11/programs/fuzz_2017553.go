package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -9223372036854775808
    v3 := make(map[uint32]bool)
    var v4 uint32 = 0
    v5 := make(map[string]uint32)
    v6 := GenS1{F0: 2147483647, F1: 0, F2: v4}
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            switch 1 {
            case 42:
                var v8 uint32 = (^(v4))
                _ = v8
            case -1:
                var v9 float64 = (0.0 - 3.14159)
                _ = v9
            }
        }
        var v10 uint64 = 0
        var v11 int = 32767
        _ = fi7
        _ = v10
        _ = v11
    }
    v12 := v5["\n\t"]
    for fi13 := 0; fi13 < 8; fi13++ {
        if ((!(673 != 255)) && (1 == 922)) {
            for fi14 := 0; fi14 < 8; fi14++ {
                v15 := GenS2{F0: 127}
                _ = fi14
                _ = v15
            }
            if (v4 >= 65535) {
                fmt.Println(v12, v4, v4)
            } else {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := GenS2{F0: 19}
                    switch "\n\t" {
                    case "foo bar":
                        _ = len(v1)
                    default:
                        v18 := GenS3{F0: 255, F1: 230}
                        _ = v18
                    }
                    v19 := GenS4{F0: (v2 - (v2 - -1)), F1: 231}
                    _ = fi16
                    _ = v17
                    _ = v19
                }
                for fi20 := 0; fi20 < 13; fi20++ {
                    v21 := GenS3{F0: 0, F1: 15}
                    if ("foo bar" < "hello") {
                        switch "a" {
                        case "\n\t":
                            v2 += v2
                            var v22 int64 = 0
                            _ = v22
                        default:
                            v23 := v1[0]
                            _ = v23
                        }
                        switch (-(fi13)) {
                        case 246:
                            fmt.Println(fi20, fi13)
                        }
                        var v24 uint64 = 18446744073709551615
                        _ = v24
                    }
                    _ = fi20
                    _ = v21
                }
                v2 = (v2 ^ v2)
            }
        }
        v2 += -716
        var v25 uint16 = 926
        _ = fi13
        _ = v25
    }
    var v26 uint64 = 4294967295
    var v27 uint16 = 368
    v28 := GenS5{F0: ("foo bar" + "hello")}
    var v29 int64 = (v2 | v2)
    v12 = v12
    if false {
        var v30 int = 42
        _ = v30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 float64 = 1.5
        fmt.Println(v4, v12, v4)
        v33 := []int{}
        _ = fi31
        _ = v32
        _ = v33
    }
    v27 = v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
