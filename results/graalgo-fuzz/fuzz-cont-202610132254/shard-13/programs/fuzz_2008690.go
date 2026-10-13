package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 uint = 65535
    var v4 float32 = (1.5 / 0.1)
    v5 := []string{}
    var v6 float32 = v4
    v3 += (v3 + (v3 >> ((1) & 63)))
    fmt.Println(v6)
    v6 = (v6 - v6)
    for fi7 := 0; fi7 < 4; fi7++ {
        if (0 != 2) {
            switch "a" {
            case "foo bar":
                v4 = float32(-2.25)
            case "\n\t":
                v3 += (v3 * 255)
            default:
                if (656 == 0) {
                    v3 = (v3 - (v3 << (((v3 - 18446744073709551615)) & 63)))
                    v8 := v2[""]
                    v3 = (v3 % (((v3 - (^(v3)))) | 1))
                    _ = v8
                } else {
                    if true {
                        v3 = 215
                        fmt.Println(v4, v3)
                    }
                }
            }
            switch "" {
            case "a":
                v9 := GenS1{F0: (551 != 2)}
                _ = v9
            case "\n\t":
                v3 = (v3 >> ((v3) & 63))
            case "hello":
                if false {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(fi7, v6)
                    v2["a"] = (fi7 >> (((v3 ^ (v3 / ((v3) | 1)))) & 63))
                }
            default:
                v6 = v4
            }
            var v10 uint16 = 255
            _ = v10
        } else {
            delete(v1, 2147483648)
            v3 -= v3
        }
        var v11 int16 = -1
        _ = fi7
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS1{F0: false}
        _ = fi12
        _ = v13
    }
    v3 = v3
    v3 += (v3 + (v3 - 2))
    if false {
        var v14 string = "a"
        _ = v14
    } else {
        var v15 uint64 = 788
        _ = v15
    }
    delete(v2, "hello")
    if true {
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(v3, v6)
            _ = fi16
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if true {
            var v17 uint = (v3 ^ v3)
            v5 = append(v5, "foo bar")
            _ = v17
        } else {
            switch 4294967295 {
            case 128:
                v18 := GenS1{F0: true}
                _ = v18
            case 2147483647:
                v5 = append(v5, "")
            default:
                v6 += v4
            }
        }
        if false {
            v6 -= v6
        } else {
            var v19 int8 = -74
            var v20 string = "foo bar"
            if len(v5) > 0 {
                v5[0] = v20
            }
            _ = v19
            _ = v20
        }
    }
    v21 := GenS1{F0: true}
    switch "" {
    case "hello":
        if false {
            var v22 int32 = 0
            var v23 int32 = 1
            _ = v22
            _ = v23
        }
    default:
        switch 2147483647 {
        case -9223372036854775808:
            var v24 string = ""
            _ = v24
        default:
            if (2 > 861) {
                v4 = (0.0 + 1e10)
                fmt.Println(v6, v4, v3)
            }
        }
    }
    v3 = v3
    var v25 uint16 = 42
    fmt.Println(v25, v6)
    v26 := GenS2{F0: -280}
    for fi27 := 0; fi27 < 15; fi27++ {
        if ((528 > -9223372036854775808) && (v3 >= v3)) {
            v3 = (v3 - (v3 / ((v3) | 1)))
            if false {
                v3 += (v3 >> ((0) & 63))
                var v28 int64 = 592
                _ = v28
            }
            fmt.Println(v3, v3, v6)
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
