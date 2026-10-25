package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }

func main() {
    var v1 int32 = 0
    var v2 string = (("hello" + "a") + "")
    var v3 string = v2
    var v4 int32 = (v1 >> ((65535) & 63))
    v5 := []int{505, 128, -2147483648, 1}
    fmt.Println(v1)
    for fi6 := 0; fi6 < 2; fi6++ {
        switch 327 {
        case 4294967295:
            if false {
                for fi7 := 0; fi7 < 7; fi7++ {
                    if (0 < (v1 ^ -1)) {
                        fmt.Println(fi7, v1)
                    }
                    fmt.Println(v2)
                    _ = fi7
                }
                if len(v5) > 0 {
                    v5[0] = fi6
                }
                var v8 float32 = 3.14159
                _ = v8
            }
        default:
            v9 := GenS1{F0: ((v1 >> ((18446744073709551615) & 63)) == v1), F1: ((0 >= 914) && false)}
            _ = v9
        }
        switch -97 {
        case -1:
            switch ("a" + "foo bar") {
            case "hello":
                v1 += (-(v4))
            case "\n\t":
                _ = len(v5)
            default:
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 string = v2
                    var v12 string = (v11 + v3)
                    v2 = v11
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        case 32768:
            var v13 string = v2
            _ = v13
        default:
            fmt.Println(v3)
        }
        _ = fi6
    }
    fmt.Println(v1, v2)
    var v14 uint8 = 127
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    var v15 int64 = -9223372036854775808
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int64 = v15
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(v3, v15)
        _ = fi18
    }
    var v19 int32 = (-(v4))
    var v20 int = -1
    v5 = append(v5, -56)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 += 0
    var v21 int16 = -32768
    v14 -= 121
    if (!((false && true) && (184 >= 1))) {
        v21 -= (-(v21))
        var v22 bool = true
        _ = v22
    } else {
        v23 := GenS1{F0: false, F1: true}
        v15 = (^(v15))
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
