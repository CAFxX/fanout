package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint16; F1 int16 }

func main() {
    v1 := []int{128, 127, 32767, 0}
    var v2 int = 127
    var v3 int32 = 1
    v4 := make(map[uint32]int)
    v5 := make(map[string]bool)
    v6 := GenS1{F0: 118, F1: false}
    for fi7 := 0; fi7 < 3; fi7++ {
        v1 = append(v1, v2)
        for fi8 := 0; fi8 < 18; fi8++ {
            v2 = v2
            _ = fi8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint16 = 2
        _ = fi9
        _ = v10
    }
    fmt.Println(v2, v3)
    for fi11 := 0; fi11 < 3; fi11++ {
        v2 -= (v2 - (fi11 - (fi11 >> ((42) & 63))))
        fmt.Println(v2)
        _ = fi11
    }
    v12 := GenS1{F0: 231, F1: (!false)}
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            switch (-(v2)) {
            case 878:
                fmt.Println(fi13)
            default:
                fmt.Println(v3, v3, fi13)
            }
            var v15 int64 = -9223372036854775808
            if len(v1) > 0 {
                v1[0] = int(v3)
            }
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    if false {
        v16 := v5["hello"]
        v17 := GenS1{F0: 0, F1: (!(((v3 >> ((255) & 63)) <= (v3 * v3)) && (42 == 861)))}
        _ = v16
        _ = v17
    } else {
        var v18 string = ("" + "a")
        if false {
            if false {
                var v19 uint16 = 65535
                var v20 bool = true
                var v21 int = (v2 & v2)
                _ = v19
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 6; fi22++ {
                    switch "" {
                    case "foo bar":
                        var v23 int16 = 1
                        _ = v23
                    case "\n\t":
                        delete(v4, 2147483647)
                    default:
                        if (42 > -625) {
                            switch (-(v2)) {
                            case 128:
                                var v24 string = "foo bar"
                                _ = v24
                            default:
                                var v25 bool = ((!(true && (!(9223372036854775807 != (^(v2)))))) && (0 < 19))
                                _ = v25
                            }
                        }
                    }
                    v18 = v18
                    _ = fi22
                }
                v26 := GenS2{F0: 65535, F1: 853}
                _ = v26
            }
            var v27 uint8 = 13
            if (0 == 0) {
                var v28 uint32 = 197
                _ = v28
            }
            _ = v27
        } else {
            _ = len(v4)
        }
        switch (v2 + 127) {
        case 124:
            if false {
                if true {
                    var v29 bool = ((2147483647 > 4294967295) || false)
                    var v30 float64 = (float64(1.5) * 1e10)
                    _ = v29
                    _ = v30
                }
                v3 = v3
                for fi31 := 0; fi31 < 8; fi31++ {
                    if true {
                        for fi32 := 0; fi32 < 10; fi32++ {
                            v33 := GenS2{F0: 32767, F1: 32767}
                            fmt.Println(v18, fi32)
                            _ = fi32
                            _ = v33
                        }
                        var v34 uint32 = 42
                        _ = v34
                    }
                    v18 = (v18 + (v18 + v18))
                    var v35 float32 = (3.14159 + 0.1)
                    _ = fi31
                    _ = v35
                }
            }
        case 4294967295:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v18
    }
    fmt.Println(v2, v2)
    v36 := GenS2{F0: 0, F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
