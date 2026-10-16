package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int, p1 uint16) int16 {
    p0 += (^(p0))
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := []int{p0, (fi6 - (fi6 - fi6)), (p0 >> ((42) & 63))}
        if len(v7) > 0 {
            v7[0] = p0
        }
        if true {
            var v8 int8 = -128
            _ = v8
        } else {
            var v9 uint = 717
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    p0 = (p0 / (((p0 << ((4294967295) & 63))) | 1))
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int) bool {
    p0 += 127
    p0 -= (p0 % ((p0) | 1))
    v10 := []int{p0, 156, -2147483648}
    _ = p0
    _ = v10
    return true
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 1.5
    v3 := []uint32{65535, 2147483648, 2, 465}
    v4 := make(map[uint32]int)
    var v5 uint32 = 2
    for fi11 := 0; fi11 < 10; fi11++ {
        if (!true) {
            v5 = (v5 | v5)
        } else {
            for fi12 := 0; fi12 < 9; fi12++ {
                switch (("" + "\n\t") + "\n\t") {
                case "hello":
                    fmt.Println(v5)
                case "a":
                    v5 = (^(v5))
                case "\n\t":
                    if true {
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v13 := v4[(^(v5))]
                            var v14 int = 32768
                            _ = v13
                            _ = v14
                        } else {
                            if len(v3) > 0 {
                                v3[0] = (v5 >> ((65535) & 63))
                            }
                        }
                        v15 := GenS1{F0: 32767, F1: 485, F2: -1}
                        switch (fi11 * (fi11 / ((-46) | 1))) {
                        case -2147483648:
                            v5 += (^(v5))
                        case -1:
                            var v16 string = "foo bar"
                            _ = v16
                        case -9223372036854775808:
                            fmt.Println(fi11)
                        default:
                            var v17 int64 = 1
                            _ = v17
                        }
                        v18 := v4[v5]
                        _ = v15
                        _ = v18
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(v2, fi12)
                        switch (fi11 | -1) {
                        case 4294967295:
                            fmt.Println(fi11, v5, fi12)
                        case -33:
                            v2 -= -2.25
                            fmt.Println(fi12)
                        case 42:
                            var v19 int64 = 720
                            _ = v19
                        }
                    }
                }
                var v20 bool = (!false)
                _ = fi12
                _ = v20
            }
            delete(v1, (((("" + (("" + "\n\t") + "\n\t")) + "\n\t") + "\n\t") + ("hello" + "\n\t")))
        }
        _ = fi11
    }
    v3 = append(v3, v5)
    var v21 int16 = 133
    if false {
        fmt.Println(v5, v2, v2)
        v21 += v21
    }
    switch "foo bar" {
    case "a":
        v21 += (^(v21))
    case "foo bar":
        if false {
            v22 := GenS1{F0: 2, F1: v5, F2: 100}
            v5 = v5
            _ = v22
        } else {
            var v23 uint16 = 905
            delete(v4, (^(v5)))
            _ = v23
        }
    }
    var v24 bool = false
    for fi25 := 0; fi25 < 10; fi25++ {
        v26 := GenS1{F0: 406, F1: v5, F2: 0}
        _ = fi25
        _ = v26
    }
    switch 865 {
    case -2147483648:
        v2 -= ((v2 * (1e10 * ((-2.25 * float64(1e10)) * (float64(0.0) * 1e10)))) - 0.0)
    case 72:
        switch 127 {
        case -431:
            for fi27 := 0; fi27 < 6; fi27++ {
                for fi28 := 0; fi28 < 8; fi28++ {
                    v3 = append(v3, v5)
                    _ = fi28
                }
                v29 := GenS1{F0: 32767, F1: (v5 | v5), F2: 1}
                var v30 uint64 = 4294967295
                _ = fi27
                _ = v29
                _ = v30
            }
        }
    }
    if ((787 <= 0) || (false && v24)) {
        var v31 uint = 65535
        _ = v31
    }
    var v32 int16 = (^(v21))
    if v24 {
        var v33 bool = v24
        v2 = (v2 * (v2 * v2))
        _ = v33
    }
    if (false && (793 <= 9223372036854775807)) {
        v34 := GenS2{F0: -47}
        v35 := GenS1{F0: 893, F1: (v5 + v5), F2: 0}
        _ = v34
        _ = v35
    }
    fmt.Println(v5, v2)
    v24 = v24
    v36 := GenS1{F0: 184, F1: v5, F2: 2}
    for fi37 := 0; fi37 < 11; fi37++ {
        v38 := GenS2{F0: 100}
        for fi39 := 0; fi39 < 17; fi39++ {
            v21 = v21
            fmt.Println(v24, fi39)
            v2 -= v2
            switch (^(fi39)) {
            case 4294967295:
                fmt.Println(v32, v5)
            }
            _ = fi39
        }
        _ = fi37
        _ = v38
    }
    v32 += (v21 - int16(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
}
