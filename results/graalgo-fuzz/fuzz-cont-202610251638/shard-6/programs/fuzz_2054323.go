package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint {
    var v7 uint = 75
    switch 859 {
    case 127:
        switch 127 {
        case 142:
            if true {
                fmt.Println(v7)
            }
        case 697:
            v8 := []int{42, 420, 0, 4294967295}
            _ = v8
        case -791:
            v9 := GenS1{F0: 0}
            _ = v9
        default:
            var v10 uint = (v7 % ((v7) | 1))
            _ = v10
        }
    }
    v7 += 0
    v7 -= (v7 << ((v7) & 63))
    fmt.Println(v7)
    _ = v7
    return 2
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    var v3 int16 = 32767
    var v4 uint16 = 79
    var v5 float32 = 0.1
    var v6 int16 = 1000
    switch "" {
    case "a":
        fmt.Println(v6)
    case "":
        v5 = -2.25
    case "foo bar":
        v3 = (^(v6))
    }
    fmt.Println(v5, v3, v6)
    for fi11 := 0; fi11 < 15; fi11++ {
        fmt.Println(v3, v3, fi11)
        v3 += v6
        fmt.Println(v4, fi11)
        _ = fi11
    }
    var v12 int = 42
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint8 = 1
        var v15 int64 = 42
        _ = fi13
        _ = v14
        _ = v15
    }
    v12 += v12
    if false {
        switch 128 {
        case -2147483648:
            if false {
                var v16 int32 = -886
                v5 = v5
                _ = v16
            }
        case 9223372036854775807:
            if ((218 <= 1) || ((^(v4)) > v4)) {
                switch ("" + ("foo bar" + "a")) {
                case "hello":
                    switch ("foo bar" + ("\n\t" + ("foo bar" + "foo bar"))) {
                    case "\n\t":
                        for fi17 := 0; fi17 < 9; fi17++ {
                            var v18 int8 = -96
                            v5 += (float32(-2.25) * 0.1)
                            v19 := []int{(v12 % (((-(fi17))) | 1)), fi17, (fi17 << ((uint(v4)) & 63)), (v12 << ((42) & 63))}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    case "hello":
                        var v20 uint32 = 194
                        _ = v20
                    }
                }
                v21 := []int{}
                _ = v21
            } else {
                var v22 uint8 = 236
                var v23 int32 = 42
                v12 -= -262
                _ = v22
                _ = v23
            }
        }
    }
    delete(v2, 2147483647)
    if ((false || false) || true) {
        if false {
            var v24 string = "a"
            v5 = 0.1
            v5 -= ((float32((1e10 / 1e10)) + v5) + v5)
            _ = v24
        }
        v25 := GenS1{F0: 18446744073709551615}
        var v26 int32 = 42
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS1{F0: 0}
        if (v12 > -245) {
            v5 -= float32(((-2.25 * 0.0) / ((float64(((v5 / float32(3.14159)) + 0.1)) / (0.1 / 1.0)) - (1e10 - -2.25))))
        } else {
            v12 = 128
        }
        _ = fi27
        _ = v28
    }
    v29 := v2[2147483648]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
