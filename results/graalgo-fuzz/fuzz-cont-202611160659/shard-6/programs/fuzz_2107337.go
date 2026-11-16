package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []string{"a", "a", "\n\t", "hello"}
    var v2 int32 = 971
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v5 int16 = 843
    var v6 uint16 = 42
    fmt.Println(v6, v5, v5)
    var v7 uint16 = (v6 / (((v6 & v6)) | 1))
    if false {
        for fi8 := 0; fi8 < 4; fi8++ {
            _ = len(v1)
            if (-1 == 109) {
                var v9 int32 = (v2 - (-(v2)))
                var v10 int64 = 0
                v1 = append(v1, ("hello" + ""))
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 uint64 = 0
                _ = fi11
                _ = v12
            }
            _ = fi8
        }
        v3[4294967295] = 32767
    }
    v13 := GenS1{F0: (396 <= -1)}
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            v7 = v7
            _ = fi15
        }
        _ = fi14
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint64 = 192
        switch fi16 {
        case -135:
            for fi18 := 0; fi18 < 15; fi18++ {
                switch "" {
                case "\n\t":
                    switch (fi16 + 9223372036854775807) {
                    case -9223372036854775808:
                        switch "hello" {
                        case "":
                            v2 += v2
                        case "\n\t":
                            var v19 int16 = (v5 ^ 32767)
                            _ = v19
                        default:
                            var v20 int16 = (v5 >> ((1) & 63))
                            _ = v20
                        }
                    case -2147483648:
                        switch 384 {
                        case 1:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 32767:
                            if (!true) {
                                var v21 int16 = v5
                                var v22 int = fi16
                                var v23 float64 = (-2.25 + 1.5)
                                v1 = append(v1, "")
                                v24 := GenS1{F0: ((v17 | (v17 & v17)) <= 362)}
                                _ = v21
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        case 32768:
                            fmt.Println(v2, v5, v5)
                        default:
                            fmt.Println(fi16, v6)
                        }
                    }
                default:
                    if true {
                        if len(v1) > 0 {
                            v1[0] = (("" + "hello") + ("hello" + "a"))
                        }
                    }
                }
                _ = fi18
            }
        }
        switch ("hello" + "foo bar") {
        case "":
            var v25 int32 = 1
            _ = v25
        case "a":
            var v26 uint16 = v7
            _ = v26
        }
        _ = fi16
        _ = v17
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        v6 = (v7 >> ((199) & 63))
        _ = fi27
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        v2 = v2
        var v29 uint32 = 0
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        v7 += v6
        _ = fi30
    }
    if true {
        var v31 uint32 = 0
        v2 -= -513
        _ = v31
    }
    fmt.Println(v2, v2, v6)
    var v32 float64 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
