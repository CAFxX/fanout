package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() string {
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    var v6 uint16 = 32767
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (((0.0 / 0.1) - (0.1 + 1.5)) / 1.0)
    var v3 uint8 = 155
    v4 := []int{259, 0, 2147483647, -1}
    v2 -= v2
    if len(v4) > 0 {
        v4[0] = -215
    }
    if ((false || false) && (!(9223372036854775807 >= 42))) {
        var v7 string = (("hello" + "\n\t") + "")
        switch 410 {
        case -230:
            for fi8 := 0; fi8 < 7; fi8++ {
                v7 = v7
                delete(v1, "foo bar")
                var v9 float64 = (v2 - v2)
                _ = fi8
                _ = v9
            }
        case 889:
            if (true && false) {
                if true {
                    var v10 string = ("\n\t" + v7)
                    v1["a"] = 0
                    _ = v10
                }
                for fi11 := 0; fi11 < 15; fi11++ {
                    v12 := GenS1{F0: 91}
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 int32 = -2147483648
                    v3 = 1
                    _ = fi13
                    _ = v14
                }
            } else {
                v1[v7] = -194
                switch v7 {
                case "\n\t":
                    var v15 int8 = 127
                    _ = v15
                case "foo bar":
                    if true {
                        v3 = (v3 << ((uint(v3)) & 63))
                        switch (v7 + "\n\t") {
                        case "hello":
                            switch "\n\t" {
                            case "hello":
                                switch -933 {
                                case 32767:
                                    var v16 float64 = 1e10
                                    _ = v16
                                default:
                                    v2 = (v2 * (v2 * v2))
                                }
                            case "":
                                if true {
                                    var v17 string = "a"
                                    fmt.Println(v17)
                                    _ = v17
                                }
                            }
                        }
                    } else {
                        v18 := GenS1{F0: 100}
                        _ = v18
                    }
                default:
                    fmt.Println(v2, v7)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        default:
            fmt.Println(v2, v2, v3)
        }
        v7 = ((((((("\n\t" + v7) + (v7 + v7)) + "foo bar") + "") + ((("" + (v7 + v7)) + v7) + "a")) + v7) + "foo bar")
        _ = v7
    } else {
        _ = len(v1)
        var v19 uint8 = 127
        _ = v19
    }
    v20 := GenS1{F0: 42}
    var v21 uint64 = 42
    var v22 uint16 = 42
    if false {
        _ = len(v1)
        var v23 uint16 = uint16(v3)
        _ = v23
    } else {
        for fi24 := 0; fi24 < 13; fi24++ {
            _ = len(v4)
            _ = fi24
        }
    }
    if true {
        var v25 bool = (42 > 2147483647)
        delete(v1, "")
        v26 := GenS1{F0: 42}
        _ = v25
        _ = v26
    } else {
        switch 128 {
        case 179:
            switch 42 {
            case 1:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                v21 -= (v21 * v21)
            }
        default:
            v3 -= 255
        }
        fmt.Println(v22, v22, v2)
        v1[("a" + "")] = -711
    }
    var v27 uint16 = (v22 << ((2) & 63))
    v27 -= (v22 | (v27 - (v27 / ((v27) | 1))))
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
}
