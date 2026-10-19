package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 bool }
type GenS2 struct { F0 int }

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 128
    var v3 string = ((("foo bar" + "a") + (("" + "\n\t") + "\n\t")) + "")
    v4 := GenS1{F0: (^(v1)), F1: v3, F2: false}
    var v5 int32 = -1
    var v6 string = v3
    if (32767 >= 574) {
        v7 := GenS1{F0: v1, F1: v6, F2: false}
        var v8 uint32 = 1
        if false {
            fmt.Println(v6, v5)
            v8 += (v8 * uint32(v2))
        } else {
            v9 := GenS1{F0: (v1 ^ v1), F1: "hello", F2: (v8 == 7)}
            _ = v9
        }
        _ = v7
        _ = v8
    }
    if (false || ((v5 << ((1) & 63)) <= v5)) {
        fmt.Println(v3, v6)
    }
    var v10 int64 = 139
    v11 := []int{128}
    if true {
        var v12 uint = 65535
        _ = v12
    } else {
        if false {
            var v13 int32 = 1
            _ = v13
        } else {
            v5 = (v5 ^ 2147483647)
        }
        if true {
            _ = len(v11)
            if len(v11) > 0 {
                v11[0] = 9223372036854775807
            }
        }
    }
    v14 := GenS2{F0: -2147483648}
    if (42 < 537) {
        switch ("a" + v3) {
        case "":
            if (true && (65535 > 18446744073709551615)) {
                _ = len(v11)
            }
        case "foo bar":
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 int64 = v10
                fmt.Println(v6, v16, v1)
                switch 2147483647 {
                case 93:
                    v16 += (v16 + (v16 % (((v10 % ((-603) | 1))) | 1)))
                case 9223372036854775807:
                    v17 := GenS2{F0: fi15}
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
        case "hello":
            for fi18 := 0; fi18 < 8; fi18++ {
                v3 = ((v3 + v6) + v6)
                var v19 int64 = (v10 << ((347) & 63))
                v20 := GenS1{F0: (v1 % ((4294967295) | 1)), F1: "\n\t", F2: (-1 <= -9223372036854775808)}
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    }
    var v21 int16 = 0
    switch 0 {
    case 156:
        var v22 bool = (!false)
        _ = v22
    case 919:
        switch -2147483648 {
        case 0:
            fmt.Println(v2, v6)
        case 42:
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        case -9223372036854775808:
            for fi23 := 0; fi23 < 5; fi23++ {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                if false {
                    v11 = append(v11, (fi23 * fi23))
                }
                _ = fi23
            }
        }
    case 128:
        if ((v21 - v21) > int16(v10)) {
            v3 = ((("\n\t" + ((v3 + v6) + (v6 + v3))) + "") + (("hello" + (((("" + "a") + (v3 + v6)) + "\n\t") + (v3 + (v6 + (v3 + ((v6 + v6) + v6)))))) + "a"))
            if (true && false) {
                fmt.Println(v10, v5)
            } else {
                switch v6 {
                case "":
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                }
            }
            var v24 bool = true
            _ = v24
        } else {
            v5 = (v5 * -1)
            var v25 float32 = 1.5
            _ = v25
        }
    }
    var v26 float64 = 0.1
    v27 := GenS2{F0: 108}
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 uint64 = 18446744073709551615
        _ = fi28
        _ = v29
    }
    var v30 uint8 = uint8(v1)
    v31 := GenS2{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
