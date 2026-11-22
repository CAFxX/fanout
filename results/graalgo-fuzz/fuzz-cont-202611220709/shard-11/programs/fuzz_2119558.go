package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint16 }

func main() {
    var v1 uint32 = 42
    var v2 int = 42
    var v3 int = (v2 << ((21) & 63))
    v2 += (v3 / (((-(v2))) | 1))
    var v4 float64 = 0.1
    v5 := GenS1{F0: -2147483648, F1: 1, F2: 128}
    v6 := []int{(v3 - 1)}
    var v7 int = v2
    if false {
        v2 = 26
    }
    if true {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS2{F0: 1, F1: 241, F2: 1}
            v10 := GenS2{F0: -53, F1: 1, F2: 406}
            _ = fi8
            _ = v10
            _ = v9
        }
        if (127 >= -1) {
            if false {
                for fi11 := 0; fi11 < 4; fi11++ {
                    v12 := GenS1{F0: -1, F1: -2147483648, F2: 1}
                    switch ("a" + ("a" + "\n\t")) {
                    case "foo bar":
                        var v13 float32 = (1.5 / 1e10)
                        _ = v13
                    case "":
                        switch (v7 % ((fi11) | 1)) {
                        case 9223372036854775807:
                            if true {
                                var v14 int32 = -565
                                v15 := GenS2{F0: 0, F1: 42, F2: 65535}
                                _ = v14
                                _ = v15
                            }
                        case 127:
                            v16 := GenS2{F0: -40, F1: 0, F2: 735}
                            _ = v16
                        }
                    default:
                        for fi17 := 0; fi17 < 12; fi17++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v18 uint = 18446744073709551615
                            var v19 string = ("foo bar" + "foo bar")
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                    v20 := GenS2{F0: 127, F1: 255, F2: 65535}
                    _ = fi11
                    _ = v12
                    _ = v20
                }
            } else {
                var v21 uint = 571
                v3 = -39
                _ = v21
            }
        }
    }
    v4 = (v4 / 1.5)
    if len(v6) > 0 {
        v6[0] = 2147483647
    }
    switch ((((("a" + "a") + "hello") + "a") + "hello") + (("a" + ("a" + "foo bar")) + "hello")) {
    case "a":
        v22 := GenS1{F0: 1, F1: 2147483647, F2: 255}
        _ = v22
    default:
        var v23 int32 = -1
        _ = v23
    }
    v6 = append(v6, (v2 + (v7 ^ (v2 << ((65535) & 63)))))
    var v24 int64 = 0
    for fi25 := 0; fi25 < 14; fi25++ {
        v26 := GenS2{F0: 112, F1: 1, F2: 620}
        for fi27 := 0; fi27 < 9; fi27++ {
            v28 := GenS2{F0: 127, F1: 127, F2: 124}
            var v29 int8 = 100
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi25
        _ = v26
    }
    if len(v6) > 0 {
        v6[0] = v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
