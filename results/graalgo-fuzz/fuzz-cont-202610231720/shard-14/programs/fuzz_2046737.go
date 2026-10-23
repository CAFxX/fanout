package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 bool }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 bool) int16 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := []int{fi4, (fi4 % ((0) | 1))}
        v6 := GenS1{F0: (!false), F1: 1, F2: (1 == 103)}
        if false {
            if (!p0) {
                var v7 bool = (!(!true))
                v7 = (0 < 0)
                _ = v7
            } else {
                if ((128 == 252) || p0) {
                    var v8 bool = false
                    fmt.Println(fi4)
                    var v9 uint = 42
                    _ = v8
                    _ = v9
                } else {
                    v10 := GenS1{F0: (((2 <= 2) || true) && p0), F1: -79, F2: true}
                    switch -322 {
                    case 127:
                        for fi11 := 0; fi11 < 3; fi11++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            p0 = false
                            _ = fi11
                        }
                    case 0:
                        for fi12 := 0; fi12 < 8; fi12++ {
                            switch "" {
                            case "a":
                                var v13 int64 = -1
                                if true {
                                    var v14 int64 = v13
                                    v15 := GenS1{F0: true, F1: 127, F2: true}
                                    _ = v14
                                    _ = v15
                                } else {
                                    p0 = (218 < 1)
                                }
                                if p0 {
                                    if (true && (!(-9223372036854775808 <= fi4))) {
                                        var v16 uint = 1
                                        switch "" {
                                        case "a":
                                            switch ("" + ((("foo bar" + "\n\t") + "foo bar") + ("hello" + "a"))) {
                                            case "\n\t":
                                                v17 := GenS1{F0: p0, F1: 42, F2: (1 < 127)}
                                                _ = v17
                                            }
                                        }
                                        switch (fi12 >> (((v16 + (v16 * v16))) & 63)) {
                                        case 585:
                                            if len(v5) > 0 {
                                                v5[0] = fi4
                                            }
                                        case -9223372036854775808:
                                            v16 -= v16
                                        case -2147483648:
                                            v18 := GenS1{F0: (p0 && (!p0)), F1: -63, F2: p0}
                                            _ = v18
                                        }
                                        p0 = (!(!(p0 || p0)))
                                        _ = v16
                                    }
                                    v5 = append(v5, (-(fi12)))
                                } else {
                                    if ((-2147483648 < 123) || false) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v13 += 0
                                    } else {
                                        var v19 int8 = 0
                                        _ = v19
                                    }
                                    var v20 uint16 = 0
                                    var v21 int32 = 42
                                    if false {
                                        fmt.Println(fi12)
                                    }
                                    var v22 uint16 = v20
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                                _ = v13
                            }
                            var v23 uint32 = 65535
                            _ = fi12
                            _ = v23
                        }
                    }
                    _ = v10
                }
                fmt.Println(fi4, fi4)
            }
            var v24 int32 = 2147483647
            _ = v24
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    _ = p0
    return 1
}

func main() {
    var v1 int16 = 731
    v2 := make(map[uint32]int)
    var v3 uint32 = 2147483647
    for fi25 := 0; fi25 < 14; fi25++ {
        v26 := GenS2{F0: 4294967295, F1: 255}
        var v27 int32 = 1
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 bool = true
    switch -9223372036854775808 {
    case 0:
        var v29 uint64 = 2
        _ = v29
    case 32768:
        var v30 uint32 = v3
        _ = v30
    default:
        v31 := GenS1{F0: v28, F1: -58, F2: (0 < 42)}
        _ = v31
    }
    v1 = v1
    var v32 int = 763
    var v33 uint32 = v3
    v3 = 42
    switch "" {
    case "a":
        fmt.Println(v32)
    case "foo bar":
        var v34 uint8 = 1
        _ = v34
    case "":
        switch (v32 ^ (v32 + v32)) {
        case 32767:
            fmt.Println(v1, v3, v33)
        case 2147483647:
            switch "a" {
            case "hello":
                if (749 == 0) {
                    switch (v32 * v32) {
                    case -905:
                        var v35 float32 = 3.14159
                        _ = v35
                    default:
                        fmt.Println(v33, v1)
                    }
                }
            case "a":
                var v36 float64 = 3.14159
                _ = v36
            case "foo bar":
                for fi37 := 0; fi37 < 3; fi37++ {
                    var v38 int64 = 496
                    var v39 uint64 = 768
                    fmt.Println(v39)
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            default:
                v40 := []int{(-(v32))}
                _ = v40
            }
        default:
            var v41 int = (-(v32))
            _ = v41
        }
    }
    var v42 int32 = -1
    v43 := v2[(v3 + (v3 << ((42) & 63)))]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v42
    _ = v43
}
