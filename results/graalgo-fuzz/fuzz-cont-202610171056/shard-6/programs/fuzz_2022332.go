package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]int)
    var v3 uint16 = 516
    v4 := make(map[string]int)
    var v5 bool = (false && false)
    for fi6 := 0; fi6 < 2; fi6++ {
        _ = len(v2)
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS1{F0: v5, F1: 0, F2: 548}
            v5 = (true && (true && (fi6 == (fi7 - (fi6 + fi6)))))
            switch fi6 {
            case -1:
                fmt.Println(v1, v1)
                if (127 == int8(v1)) {
                    if v5 {
                        switch 127 {
                        case 1:
                            fmt.Println(v5)
                        default:
                            var v9 int32 = -1
                            _ = v9
                        }
                        if ("\n\t" < "hello") {
                            var v10 uint16 = (^(v3))
                            _ = v10
                        } else {
                            switch (("" + "") + ("\n\t" + "hello")) {
                            case "":
                                _ = len(v2)
                            case "a":
                                var v11 int8 = -50
                                _ = v11
                            case "hello":
                                v12 := v4["hello"]
                                _ = v12
                            default:
                                var v13 float64 = 0.0
                                _ = v13
                            }
                            v5 = v5
                            var v14 int16 = int16(v3)
                            _ = v14
                        }
                        var v15 float32 = 0.1
                        var v16 uint = 899
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS1{F0: (707 <= 1), F1: 2, F2: 255}
                        var v18 uint32 = uint32(v1)
                        v1 = 249
                        _ = v17
                        _ = v18
                    }
                    v3 = v3
                    delete(v4, "a")
                }
            default:
                v19 := v4["a"]
                _ = v19
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    _ = len(v2)
    v5 = v5
    if (18446744073709551615 == 0) {
        v3 = (v3 >> ((uint(v1)) & 63))
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := GenS1{F0: v5, F1: 2147483647, F2: 65535}
            _ = len(v4)
            var v22 uint64 = 42
            _ = fi20
            _ = v21
            _ = v22
        }
        switch "hello" {
        case "a":
            _ = len(v4)
        case "foo bar":
            var v23 int16 = 860
            _ = v23
        }
    }
    fmt.Println(v3, v3, v5)
    fmt.Println(v1, v3, v3)
    for fi24 := 0; fi24 < 9; fi24++ {
        v25 := GenS1{F0: (v5 && v5), F1: 2147483647, F2: 577}
        var v26 uint16 = v3
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int64 = 42
    if (false || v5) {
        fmt.Println(v5, v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
