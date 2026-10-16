package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int32 {
    var v7 bool = (true && (-2147483648 == -1))
    var v8 int32 = 2147483647
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"hello"}
    var v3 string = "hello"
    var v4 int8 = 100
    var v5 int16 = 842
    var v6 uint8 = 42
    switch "" {
    case "hello":
        v3 = "hello"
    case "\n\t":
        v3 = "a"
    default:
        var v9 float64 = 3.14159
        _ = v9
    }
    v10 := GenS1{F0: 1}
    var v11 uint8 = (^(v6))
    switch 311 {
    case -455:
        var v12 uint16 = 0
        _ = v12
    case 15:
        v4 -= (v4 ^ v4)
    case 552:
        if false {
            for fi13 := 0; fi13 < 15; fi13++ {
                if false {
                    v6 += (v6 % ((v6) | 1))
                    if len(v2) > 0 {
                        v2[0] = ((v3 + ("a" + ("" + v3))) + (v3 + ("a" + v3)))
                    }
                } else {
                    fmt.Println(fi13)
                }
                v2 = append(v2, v3)
                for fi14 := 0; fi14 < 13; fi14++ {
                    switch "a" {
                    case "\n\t":
                        switch "hello" {
                        case "foo bar":
                            var v15 int32 = 42
                            _ = v15
                        case "":
                            v3 = ("" + v3)
                        case "\n\t":
                            v16 := GenS1{F0: 2}
                            _ = v16
                        default:
                            var v17 int8 = (v4 - (v4 * v4))
                            _ = v17
                        }
                    case "hello":
                        v18 := v1[v3]
                        switch v3 {
                        case "":
                            if false {
                                v5 -= (v5 | (v5 | (-(v5))))
                                _ = len(v2)
                                v11 = (v6 | 2)
                                var v19 int16 = 32767
                                _ = v19
                            } else {
                                var v20 uint64 = 42
                                fmt.Println(v18, v5, v5)
                                _ = v20
                            }
                        default:
                            var v21 uint16 = 42
                            _ = v21
                        }
                        _ = v18
                    default:
                        v4 = v4
                    }
                    v3 = ((v3 + v3) + v3)
                    _ = fi14
                }
                _ = fi13
            }
            v22 := GenS1{F0: 2}
            _ = v22
        }
    }
    v3 = (("a" + (v3 + "foo bar")) + (v3 + (v3 + ("" + (v3 + v3)))))
    if (42 >= 32768) {
        var v23 uint8 = (v6 - (^(v11)))
        var v24 int16 = (v5 % (((v5 | (-(v5)))) | 1))
        _ = v23
        _ = v24
    }
    var v25 int8 = (v4 >> ((4294967295) & 63))
    if (!false) {
        delete(v1, "foo bar")
        for fi26 := 0; fi26 < 8; fi26++ {
            switch v3 {
            case "\n\t":
                for fi27 := 0; fi27 < 12; fi27++ {
                    fmt.Println(v6)
                    v4 = 42
                    _ = fi27
                }
            case "foo bar":
                var v28 int = (fi26 | fi26)
                _ = v28
            case "hello":
                for fi29 := 0; fi29 < 2; fi29++ {
                    fmt.Println(v11)
                    v3 = ("a" + "a")
                    _ = fi29
                }
            default:
                for fi30 := 0; fi30 < 18; fi30++ {
                    switch fi30 {
                    case 9223372036854775807:
                        v11 = 229
                    }
                    _ = fi30
                }
            }
            _ = fi26
        }
    }
    fmt.Println(v6, v5, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
