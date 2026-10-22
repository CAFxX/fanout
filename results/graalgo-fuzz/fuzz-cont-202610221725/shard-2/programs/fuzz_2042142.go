package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []int{127, 4294967295, 4294967295}
    var v2 string = (("a" + "hello") + "foo bar")
    v3 := []uint32{218, 2147483647, 4294967295, 65535}
    var v4 uint8 = 177
    var v5 string = (v2 + (v2 + "\n\t"))
    var v6 uint8 = (v4 % (((v4 / ((v4) | 1))) | 1))
    var v7 int16 = 0
    switch 4294967295 {
    case -1:
        var v8 bool = (false && false)
        _ = v8
    case -759:
        if false {
            v5 = v2
        }
    case 2147483647:
        v4 -= (v4 & (v4 & v6))
    default:
        var v9 int = 1
        _ = v9
    }
    if true {
        var v10 float32 = -2.25
        v4 = v6
        _ = v10
    }
    var v11 int8 = -103
    v6 = v4
    v4 = v6
    if ((v11 + (v11 & v11)) >= v11) {
        v5 = "foo bar"
        var v12 int16 = v7
        if ((v11 >> ((255) & 63)) != (v11 / (((-(v11))) | 1))) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            switch v5 {
            case "foo bar":
                _ = len(v1)
            case "hello":
                if len(v3) > 0 {
                    v3[0] = 42
                }
            default:
                var v13 float32 = 3.14159
                _ = v13
            }
            v14 := GenS1{F0: true}
            _ = v14
        }
        _ = v12
    }
    var v15 uint32 = 65535
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v16 := GenS1{F0: true}
    var v17 float64 = 0.0
    switch -173 {
    case 128:
        v18 := GenS2{F0: 295}
        _ = v18
    case 0:
        if ((((v5 + (v2 + v2)) + v2) + v5) < "") {
            v3 = append(v3, v15)
            v19 := GenS1{F0: ((!false) || true)}
            var v20 uint64 = 4294967295
            _ = v19
            _ = v20
        }
    case 42:
        if false {
            if ((v11 | -30) >= v11) {
                switch v2 {
                case "hello":
                    v15 = 0
                }
                var v21 int16 = (v7 << ((1) & 63))
                if false {
                    if false {
                        v17 = 1.5
                        switch 128 {
                        case -103:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 601:
                            fmt.Println(v5, v11)
                        }
                    }
                    v22 := GenS1{F0: (18446744073709551615 < 0)}
                    _ = v22
                } else {
                    switch v2 {
                    case "foo bar":
                        v5 = v5
                    }
                }
                _ = v21
            } else {
                v4 += 128
            }
        } else {
            var v23 string = (v5 + v2)
            _ = v23
        }
    }
    var v24 uint = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
