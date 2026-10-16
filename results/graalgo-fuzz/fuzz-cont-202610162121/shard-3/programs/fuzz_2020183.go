package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int16; F2 int8 }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 uint32, p1 int, p2 int16) uint8 {
    fmt.Println(p1, p1, p0)
    p0 = (p0 / (((^(p0))) | 1))
    p2 += 42
    switch "" {
    case "a":
        var v4 float32 = (1.5 - (((1.5 - 3.14159) + (0.1 / (-2.25 / ((-2.25 - (float32(1.5) * (-2.25 * ((float32((1.5 * 0.1)) * 1.5) / ((1e10 / 0.1) / 3.14159))))) / (float32(1.5) * 3.14159))))) * -2.25))
        _ = v4
    default:
        for fi5 := 0; fi5 < 3; fi5++ {
            p2 = -1
            if (!(false && (false || (0 > 639)))) {
                v6 := []int{(-(fi5)), -572, int(p2)}
                v7 := GenS1{F0: p0}
                _ = v6
                _ = v7
            } else {
                if true {
                    if (true || ((float64(3.14159) + 1e10) == float64(0.1))) {
                        var v8 float64 = 0.0
                        _ = v8
                    } else {
                        var v9 int8 = -5
                        var v10 uint16 = 42
                        for fi11 := 0; fi11 < 4; fi11++ {
                            v12 := GenS1{F0: p0}
                            v13 := GenS1{F0: p0}
                            p2 += (p2 ^ 1)
                            var v14 int32 = 2147483647
                            fmt.Println(v10, p2, v14)
                            v9 = v9
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        _ = v10
                        _ = v9
                    }
                    var v15 int8 = 1
                    _ = v15
                }
            }
            var v16 uint = 1
            _ = fi5
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 int64 = -9223372036854775808
    v17 := GenS1{F0: 4294967295}
    switch ("hello" + "hello") {
    case "\n\t":
        if (!false) {
            for fi18 := 0; fi18 < 10; fi18++ {
                if true {
                    v3 -= (v3 & (v3 - v3))
                    v19 := []int{940, fi18}
                    switch "foo bar" {
                    case "a":
                        v3 += (v3 & (v3 * 100))
                    case "hello":
                        for fi20 := 0; fi20 < 16; fi20++ {
                            v3 = 9223372036854775807
                            _ = fi20
                        }
                    }
                    _ = v19
                }
                if (18446744073709551615 <= 42) {
                    v3 = (v3 >> ((42) & 63))
                    v21 := []int{fi18, fi18, (-(fi18))}
                    v3 += 629
                    _ = v21
                }
                _ = fi18
            }
        } else {
            fmt.Println(v3)
            var v22 uint64 = 1
            _ = v22
        }
    case "foo bar":
        fmt.Println(v3)
    case "hello":
        v3 = v3
    }
    v23 := GenS2{F0: 65535, F1: 30, F2: 100}
    v24 := GenS2{F0: 42, F1: -32768, F2: 42}
    var v25 float64 = ((1e10 * (0.1 - ((((-2.25 + ((float64((0.0 - 1.5)) - -2.25) / 1.0)) * 0.1) - ((float64(0.0) / 1.0) / 1e10)) - (0.0 + float64(1.5))))) + 0.1)
    v25 = v25
    var v26 uint = 974
    if (true && true) {
        fmt.Println(v25)
        for fi27 := 0; fi27 < 12; fi27++ {
            if (!(53 != 65535)) {
                v26 -= (v26 + (v26 & 18446744073709551615))
            } else {
                v28 := GenS2{F0: (^(v26)), F1: 0, F2: 0}
                _ = v28
            }
            fmt.Println(v26, v25, v26)
            var v29 int16 = 1000
            _ = fi27
            _ = v29
        }
        var v30 bool = true
        _ = v30
    } else {
        fmt.Println(v26, v25)
        switch "" {
        case "\n\t":
            v31 := GenS3{F0: -2147483648, F1: (("\n\t" + ("" + "\n\t")) <= "foo bar")}
            _ = v31
        }
        switch -9223372036854775808 {
        case 1:
            for fi32 := 0; fi32 < 15; fi32++ {
                v26 -= (v26 + (v26 << (((v26 | 65535)) & 63)))
                switch "" {
                case "foo bar":
                    v33 := []int{fi32, fi32}
                    _ = v33
                case "hello":
                    _ = len(v1)
                case "a":
                    var v34 int64 = 9223372036854775807
                    _ = v34
                default:
                    _ = len(v2)
                }
                _ = fi32
            }
        default:
            v35 := []int{}
            _ = v35
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
