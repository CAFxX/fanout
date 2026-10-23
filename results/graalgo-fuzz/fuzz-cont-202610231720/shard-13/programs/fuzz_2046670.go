package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 bool; F2 string }
func genFn0() uint {
    var v6 float64 = float64(0.0)
    v6 = v6
    _ = v6
    return 1
}
func genFn1() int {
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    var v8 int64 = 42
    fmt.Println(v8)
    var v9 string = (("hello" + "\n\t") + "\n\t")
    _ = v8
    _ = v9
    return 32768
}

func main() {
    var v1 uint64 = 42
    var v2 uint32 = 2147483647
    var v3 uint16 = 0
    v4 := []int{4294967295, -1, int(v2), 924}
    var v5 uint32 = v2
    switch "\n\t" {
    case "hello":
        if (true || (!(0.1 > -2.25))) {
            v1 = (^(v1))
        } else {
            var v10 uint16 = (^(v3))
            if len(v4) > 0 {
                v4[0] = 491
            }
            _ = v10
        }
    }
    v2 = (v5 * v2)
    switch -2147483648 {
    case 127:
        var v11 int16 = -1
        _ = v11
    }
    var v12 bool = false
    var v13 bool = (v12 || (-766 > 1))
    if false {
        v1 -= (v1 >> ((65535) & 63))
        v1 -= (^(v1))
        var v14 uint8 = 42
        _ = v14
    }
    switch 2147483647 {
    case 2147483647:
        v3 = v3
    case 42:
        v5 += (v5 | v2)
    }
    if false {
        v2 = v5
        v4 = append(v4, 2147483647)
    } else {
        fmt.Println(v2, v13)
        v5 -= v5
        for fi15 := 0; fi15 < 1; fi15++ {
            if (-32768 == -474) {
                for fi16 := 0; fi16 < 10; fi16++ {
                    v17 := GenS1{F0: 0, F1: uint8(v3)}
                    switch (("hello" + "a") + "foo bar") {
                    case "\n\t":
                        v18 := GenS2{F0: 116, F1: (!(v13 || ((("hello" + ((("hello" + "\n\t") + (("" + "hello") + (("foo bar" + ("\n\t" + ((("foo bar" + "\n\t") + "a") + "a"))) + ("hello" + "")))) + ("a" + "foo bar"))) + "\n\t") < "foo bar"))), F2: "\n\t"}
                        _ = v18
                    case "hello":
                        var v19 int16 = 0
                        _ = v19
                    default:
                        switch ("foo bar" + "hello") {
                        case "a":
                            var v20 uint16 = 122
                            _ = v20
                        case "\n\t":
                            var v21 float64 = float64((((1e10 / (1e10 - (3.14159 - 0.1))) + (((1.5 * -2.25) - ((float32(1.5) + 3.14159) / 3.14159)) + (0.1 + (((1e10 * 3.14159) * 3.14159) * (((3.14159 * 3.14159) / 1.5) * (0.0 / 1.0)))))) - -2.25))
                            _ = v21
                        default:
                            v13 = v12
                        }
                    }
                    _ = len(v4)
                    _ = fi16
                    _ = v17
                }
                v1 = v1
                if v13 {
                    v5 = 1
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v22 := GenS2{F0: v5, F1: (false || v12), F2: "foo bar"}
                    _ = v22
                }
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 string = ""
                _ = fi23
                _ = v24
            }
            _ = fi15
        }
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    var v25 uint16 = v3
    fmt.Println(v25)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
