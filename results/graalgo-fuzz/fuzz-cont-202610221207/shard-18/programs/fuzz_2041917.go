package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64 }
func genFn0() int8 {
    var v7 uint64 = 1
    fmt.Println(v7)
    v7 -= 1
    _ = v7
    return -1
}

func main() {
    v1 := []uint32{}
    var v2 string = ("a" + "hello")
    var v3 uint = 171
    v4 := []int{9223372036854775807, 32768, 1}
    var v5 int8 = 0
    var v6 int64 = -328
    switch v2 {
    case "hello":
        v2 = "hello"
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        v5 += genFn0()
        var v9 float64 = ((1e10 - (3.14159 / -2.25)) * 1.5)
        _ = fi8
        _ = v9
    }
    v4 = append(v4, -667)
    fmt.Println(v3, v5, v3)
    for fi10 := 0; fi10 < 7; fi10++ {
        v6 -= -1
        _ = fi10
    }
    v6 = (v6 | v6)
    var v11 int = -787
    if ((uint8(v3) != 255) && false) {
        for fi12 := 0; fi12 < 7; fi12++ {
            fmt.Println(v3)
            v11 += fi12
            _ = fi12
        }
        fmt.Println(v11, v6, v2)
        if (42 >= 157) {
            v13 := GenS1{F0: -820, F1: -792}
            _ = v13
        }
    } else {
        fmt.Println(v11, v3)
        v4 = append(v4, (v11 ^ (v11 ^ (v11 * v11))))
        var v14 int8 = (v5 >> ((0) & 63))
        _ = v14
    }
    v6 = v6
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS2{F0: v6}
        _ = fi15
        _ = v16
    }
    var v17 uint64 = 18446744073709551615
    v6 = (-(v6))
    v11 -= v11
    v3 = (v3 * 255)
    if ((true && false) || false) {
        if ("hello" <= "foo bar") {
            if len(v4) > 0 {
                v4[0] = -1
            }
        } else {
            switch (-(v11)) {
            case 892:
                for fi18 := 0; fi18 < 18; fi18++ {
                    v6 -= v6
                    var v19 int64 = v6
                    if (((255 == 22) && (!(true || true))) || ((!false) && false)) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v20 int8 = (-(v5))
                        if true {
                            switch ("hello" + v2) {
                            case "":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "hello":
                                v3 -= v3
                            default:
                                _ = len(v4)
                            }
                            var v21 uint64 = 462
                            _ = v21
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = v20
                    } else {
                        switch v2 {
                        case "a":
                            switch "hello" {
                            case "foo bar":
                                var v22 string = (v2 + v2)
                                _ = v22
                            case "":
                                v1 = append(v1, 2)
                            }
                        case "foo bar":
                            fmt.Println(v11, v11, v6)
                        case "":
                            v2 = "foo bar"
                        }
                    }
                    _ = fi18
                    _ = v19
                }
            }
            v5 += (^(v5))
        }
        v5 = 118
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v11 += (v11 + (v11 * (v11 ^ v11)))
    fmt.Println(v5, v17, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
