package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int64 }
func genFn0(p0 int8, p1 int8, p2 int8) uint16 {
    var v4 uint32 = 2
    var v5 uint = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 32767
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = -2.25
    var v3 uint8 = 127
    if true {
        v6 := []int{217, -87, -2147483648, -528}
        for fi7 := 0; fi7 < 6; fi7++ {
            if false {
                var v8 int = fi7
                if true {
                    for fi9 := 0; fi9 < 10; fi9++ {
                        var v10 float32 = 3.14159
                        v3 = (v3 / ((v3) | 1))
                        _ = fi9
                        _ = v10
                    }
                }
                if true {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        v2 += (v2 + 0.0)
                    case "a":
                        for fi11 := 0; fi11 < 9; fi11++ {
                            if (!false) {
                                var v12 bool = (true && (!(!(18446744073709551615 <= 2))))
                                if (!v12) {
                                    v8 = (v8 * (fi11 / ((487) | 1)))
                                } else {
                                    var v13 bool = v12
                                    _ = v13
                                }
                                _ = v12
                            } else {
                                _ = len(v1)
                                v14 := GenS1{F0: 81, F1: 255, F2: 1}
                                _ = v14
                            }
                            if true {
                                v2 += v2
                                var v15 int8 = 111
                                v8 = fi7
                                var v16 uint = 65535
                                _ = v15
                                _ = v16
                            }
                            v17 := GenS1{F0: 65535, F1: 0, F2: 1}
                            _ = fi11
                            _ = v17
                        }
                    case "hello":
                        if true {
                            fmt.Println(fi7, v3, v3)
                            var v18 uint64 = 773
                            _ = v18
                        }
                    }
                }
                _ = v8
            } else {
                v19 := GenS1{F0: 116, F1: 696, F2: 9223372036854775807}
                var v20 string = "foo bar"
                _ = v19
                _ = v20
            }
            _ = fi7
        }
        switch "" {
        case "\n\t":
            v21 := v1[513]
            _ = v21
        case "hello":
            v2 = v2
        case "a":
            v22 := v1[2147483647]
            _ = v22
        default:
            if len(v6) > 0 {
                v6[0] = 32767
            }
        }
        _ = v6
    } else {
        var v23 int = 98
        var v24 float32 = (1e10 + -2.25)
        v3 += (^(v3))
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        delete(v1, 2147483647)
        _ = fi25
    }
    if true {
        var v26 float64 = v2
        _ = v26
    } else {
        for fi27 := 0; fi27 < 14; fi27++ {
            for fi28 := 0; fi28 < 7; fi28++ {
                v2 = float64(float32(v2))
                _ = fi28
            }
            _ = fi27
        }
    }
    switch "" {
    case "a":
        v3 = (^(v3))
    case "foo bar":
        var v29 int16 = 42
        _ = v29
    case "hello":
        for fi30 := 0; fi30 < 20; fi30++ {
            v2 = 1e10
            _ = fi30
        }
    }
    v31 := v1[2147483648]
    switch 0 {
    case 4294967295:
        v3 -= v3
    case 2147483647:
        v32 := GenS1{F0: 2147483647, F1: 255, F2: 1}
        _ = v32
    default:
        for fi33 := 0; fi33 < 4; fi33++ {
            for fi34 := 0; fi34 < 7; fi34++ {
                v31 = (v31 + (v31 + (v31 + v31)))
                var v35 uint64 = 2
                _ = fi34
                _ = v35
            }
            v31 = v31
            v36 := v1[416]
            _ = fi33
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        v38 := GenS1{F0: 2147483648, F1: 4294967295, F2: -331}
        v39 := GenS1{F0: 2147483648, F1: 18446744073709551615, F2: 1}
        _ = fi37
        _ = v38
        _ = v39
    }
    v31 = v31
    if (3.14159 > (float64(0.0) / -2.25)) {
        v40 := GenS1{F0: 2, F1: 42, F2: 9223372036854775807}
        switch 2147483647 {
        case 71:
            switch "" {
            case "\n\t":
                var v41 uint = 829
                _ = v41
            }
        default:
            fmt.Println(v2, v31)
        }
        _ = v40
    }
    v3 = v3
    for fi42 := 0; fi42 < 18; fi42++ {
        fmt.Println(fi42, v3, v2)
        fmt.Println(v3, fi42, v3)
        _ = fi42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
}
