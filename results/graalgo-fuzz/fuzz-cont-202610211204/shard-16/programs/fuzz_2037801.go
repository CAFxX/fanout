package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int64 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 uint16, p1 int, p2 uint64) uint32 {
    var v4 int16 = 42
    p2 -= p2
    fmt.Println(v4, v4)
    switch "\n\t" {
    case "":
        switch 242 {
        case 89:
            p0 -= p0
        case -9223372036854775808:
            switch "\n\t" {
            case "":
                p0 = (p0 - p0)
            }
        case 0:
            v5 := GenS1{F0: 0, F1: 349, F2: -1}
            _ = v5
        }
    }
    v6 := GenS1{F0: 255, F1: 65535, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 2147483648
}
func genFn1(p0 uint16) uint {
    p0 = (p0 - (^(p0)))
    fmt.Println(p0)
    p0 = 2
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = 0
        _ = fi7
    }
    _ = p0
    return 255
}

func main() {
    var v1 int64 = -898
    var v2 uint64 = 0
    var v3 bool = (false || (-1 != 658))
    v8 := []int{int(v2), -635, 9223372036854775807, 128}
    v9 := GenS1{F0: 2, F1: 0, F2: 42}
    if v3 {
        switch "a" {
        case "a":
            var v10 uint = 65535
            _ = v10
        default:
            v2 = v2
        }
    }
    v11 := GenS1{F0: 127, F1: 18446744073709551615, F2: (v1 * (v1 - -1))}
    if len(v8) > 0 {
        v8[0] = -2147483648
    }
    fmt.Println(v1, v1)
    var v12 float32 = (1.5 - 0.0)
    v2 -= v2
    v13 := GenS1{F0: 127, F1: 255, F2: v1}
    fmt.Println(v12, v3)
    v3 = v3
    switch "\n\t" {
    case "\n\t":
        var v14 int64 = (v1 | v1)
        _ = v14
    case "":
        switch -728 {
        case 0:
            var v15 int64 = v1
            _ = v15
        }
    case "a":
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    default:
        if v3 {
            if (!v3) {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                switch "hello" {
                case "\n\t":
                    v2 += v2
                case "":
                    var v16 int64 = 175
                    _ = v16
                }
            }
        } else {
            fmt.Println(v3)
            switch 32767 {
            case 4294967295:
                if (!(250 >= 38)) {
                    if false {
                        v17 := GenS1{F0: 128, F1: 18446744073709551615, F2: v1}
                        var v18 string = ("foo bar" + ("a" + ("\n\t" + "foo bar")))
                        v19 := GenS1{F0: 187, F1: 18446744073709551615, F2: -1}
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        if v3 {
                            if len(v8) > 0 {
                                v8[0] = 42
                            }
                            var v20 int16 = 1
                            var v21 uint64 = v2
                            _ = v20
                            _ = v21
                        }
                        v1 = -1
                        var v22 int32 = -1
                        _ = v22
                    }
                    v23 := GenS1{F0: 28, F1: 598, F2: (v1 << ((669) & 63))}
                    v3 = ((!true) || false)
                    _ = v23
                }
            case 32768:
                var v24 bool = true
                _ = v24
            }
        }
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        v26 := GenS2{F0: 127, F1: 32767}
        _ = fi25
        _ = v26
    }
    if v3 {
        if false {
            v2 = 3
        } else {
            switch "hello" {
            case "hello":
                var v27 float64 = -2.25
                _ = v27
            }
            v28 := GenS1{F0: 2, F1: 615, F2: -769}
            v3 = (788 <= 0)
            _ = v28
        }
    } else {
        v1 = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
