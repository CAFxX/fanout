package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 bool }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int }

func main() {
    var v1 uint8 = 127
    var v2 string = (("a" + "a") + "foo bar")
    var v3 uint16 = 1
    var v4 uint16 = (v3 << ((255) & 63))
    var v5 float64 = float64(0.1)
    v6 := []int{127, 43}
    v7 := GenS1{F0: 42, F1: 694, F2: (-2147483648 != -974)}
    if (false || (9223372036854775807 > 0)) {
        v6 = append(v6, 32768)
    } else {
        fmt.Println(v5, v3, v1)
    }
    switch v2 {
    case "\n\t":
        fmt.Println(v1, v2)
    case "":
        fmt.Println(v2)
    }
    switch v2 {
    case "\n\t":
        var v8 uint16 = v4
        _ = v8
    case "":
        var v9 float64 = v5
        _ = v9
    case "hello":
        switch 860 {
        case 128:
            if (42 < (v4 >> ((255) & 63))) {
                var v10 int = -1
                if true {
                    switch ("foo bar" + ((v2 + (("hello" + v2) + (v2 + v2))) + "foo bar")) {
                    case "hello":
                        if false {
                            var v11 uint8 = (^(v1))
                            _ = v11
                        }
                    case "":
                        v12 := GenS1{F0: (^(v1)), F1: (v10 << ((4294967295) & 63)), F2: (613 <= -1)}
                        _ = v12
                    }
                } else {
                    var v13 string = v2
                    v1 = (v1 >> ((0) & 63))
                    _ = v13
                }
                v14 := GenS1{F0: (v1 * (^(v1))), F1: (v10 + 42), F2: (42 == 735)}
                _ = v10
                _ = v14
            }
        case -87:
            v3 -= v3
        case 9223372036854775807:
            fmt.Println(v1, v2)
        }
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    if true {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 float32 = (-2.25 - 0.0)
            if len(v6) > 0 {
                v6[0] = fi15
            }
            _ = fi15
            _ = v16
        }
        switch v2 {
        case "":
            v17 := GenS2{F0: (false && (!false)), F1: (((v2 + v2) + v2) + v2)}
            _ = v17
        default:
            v18 := GenS3{F0: -886}
            _ = v18
        }
    }
    fmt.Println(v2, v4)
    v19 := GenS2{F0: false, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
