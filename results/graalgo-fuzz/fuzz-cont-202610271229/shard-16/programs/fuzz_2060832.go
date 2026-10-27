package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int64, p1 int64) uint16 {
    if false {
        v7 := []int{2147483647, 163}
        var v8 bool = true
        _ = v7
        _ = v8
    } else {
        fmt.Println(p0, p1)
        fmt.Println(p0)
    }
    var v9 int16 = -32768
    v10 := GenS1{F0: 42, F1: (p1 / ((p0) | 1))}
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS2{F0: 0}
        var v13 int = (fi11 | (fi11 ^ 2147483647))
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := []string{}
    v2 := make(map[string]uint32)
    var v3 float32 = (1e10 / float32(-2.25))
    var v4 uint16 = 0
    var v5 int8 = 127
    var v6 float64 = 3.14159
    v14 := GenS1{F0: 223, F1: -9223372036854775808}
    v4 = (v4 - (v4 + 687))
    v15 := GenS1{F0: -941, F1: -9223372036854775808}
    v3 += 0.0
    if true {
        var v16 bool = (42 <= 930)
        var v17 int16 = 670
        _ = v16
        _ = v17
    } else {
        v3 = v3
        if (2147483647 == 707) {
            if (((18446744073709551615 != 4294967295) || (0 == 4294967295)) && false) {
                v5 = (v5 >> ((1) & 63))
            } else {
                v3 -= 1e10
                if (!false) {
                    fmt.Println(v5, v6)
                    v18 := GenS1{F0: 32767, F1: -225}
                    _ = v18
                } else {
                    if (v3 >= 3.14159) {
                        switch 2147483647 {
                        case -88:
                            switch "foo bar" {
                            case "a":
                                v4 = genFn0(42, -151)
                            default:
                                var v19 int16 = 32767
                                _ = v19
                            }
                        }
                        var v20 int32 = -340
                        fmt.Println(v20)
                        _ = v20
                    }
                    if false {
                        v1 = append(v1, ("a" + "hello"))
                        var v21 float32 = v3
                        var v22 int32 = -1
                        _ = v21
                        _ = v22
                    }
                    var v23 int32 = 0
                    _ = v23
                }
            }
            var v24 uint = 4294967295
            v5 -= -1
            _ = v24
        }
    }
    v25 := GenS1{F0: 0, F1: -9223372036854775808}
    for fi26 := 0; fi26 < 12; fi26++ {
        delete(v2, (("\n\t" + (("foo bar" + (("" + "hello") + "a")) + "\n\t")) + "\n\t"))
        fmt.Println(v3)
        var v27 uint32 = uint32(v4)
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        if false {
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 uint32 = 2147483648
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v31 := GenS1{F0: 2147483647, F1: -1}
                v5 = (v5 << ((255) & 63))
                _ = fi29
                _ = v30
                _ = v31
            }
            var v32 uint32 = 1
            v4 = genFn0(1, -9223372036854775808)
            _ = v32
        }
        var v33 int = 9223372036854775807
        for fi34 := 0; fi34 < 16; fi34++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v35 uint8 = 127
            _ = fi34
            _ = v35
        }
        _ = fi28
        _ = v33
    }
    v5 = (v5 & 100)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
