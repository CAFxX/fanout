package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 4294967295
    var v3 uint = (v2 & (^(v2)))
    var v4 int32 = 1
    var v5 string = "a"
    if true {
        v6 := []int{594, 1, 32768, -196}
        _ = v6
    }
    v3 -= 18446744073709551615
    delete(v1, 182)
    delete(v1, 2)
    if true {
        _ = len(v1)
        var v7 uint = v3
        _ = len(v1)
        _ = v7
    } else {
        _ = len(v1)
        if true {
            fmt.Println(v2, v2)
        }
    }
    v8 := []int{-2147483648, -1, -18, -430}
    switch v5 {
    case "":
        fmt.Println(v5, v2, v5)
    default:
        if true {
            v5 = v5
        } else {
            v9 := GenS1{F0: v2}
            fmt.Println(v4)
            _ = v9
        }
    }
    if (!true) {
        v10 := GenS2{F0: 42, F1: 42}
        var v11 uint8 = 128
        _ = v10
        _ = v11
    }
    if true {
        v5 = (v5 + (v5 + ""))
        switch -9223372036854775808 {
        case 127:
            var v12 bool = (true && (!false))
            _ = v12
        default:
            switch v5 {
            case "":
                var v13 uint16 = 574
                _ = v13
            case "a":
                var v14 int16 = 42
                _ = v14
            default:
                switch "" {
                case "hello":
                    var v15 bool = true
                    _ = v15
                case "":
                    if ((v5 + v5) < ((v5 + v5) + (((v5 + (((v5 + v5) + ("foo bar" + v5)) + (v5 + v5))) + v5) + v5))) {
                        fmt.Println(v5)
                    }
                case "\n\t":
                    var v16 int8 = 127
                    _ = v16
                }
            }
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            switch v5 {
            case "a":
                var v18 uint16 = 42
                _ = v18
            case "foo bar":
                v8 = append(v8, (fi17 ^ 2147483647))
            }
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(fi19)
                if (false && false) {
                    fmt.Println(v5)
                    v5 = ("hello" + ((v5 + ("\n\t" + ("\n\t" + "hello"))) + "hello"))
                } else {
                    if false {
                        v4 += (v4 % (((v4 >> ((v2) & 63))) | 1))
                        if false {
                            if true {
                                v3 += v3
                            } else {
                                v20 := GenS1{F0: (v2 ^ (v2 ^ (v3 | 18446744073709551615)))}
                                if len(v8) > 0 {
                                    v8[0] = (^(fi17))
                                }
                                if len(v8) > 0 {
                                    v8[0] = fi19
                                }
                                _ = v20
                            }
                            if (42 <= 605) {
                                fmt.Println(v4, v3, fi17)
                            } else {
                                var v21 float64 = 1e10
                                _ = v21
                            }
                        }
                    }
                    fmt.Println(v4, v4)
                    v2 -= v2
                    var v22 uint64 = 18446744073709551615
                    _ = v22
                }
                _ = fi19
            }
            if (!false) {
                delete(v1, 65535)
            }
            _ = fi17
        }
    } else {
        var v23 int32 = (v4 * v4)
        _ = v23
    }
    switch -9223372036854775808 {
    case 32768:
        v24 := GenS2{F0: 2, F1: 205}
        _ = v24
    case 0:
        v25 := GenS1{F0: (^(v3))}
        _ = v25
    }
    if (255 == 502) {
        v26 := GenS2{F0: 42, F1: 42}
        fmt.Println(v3)
        switch v5 {
        case "foo bar":
            var v27 float64 = 3.14159
            _ = v27
        case "\n\t":
            v2 = (v2 << ((18446744073709551615) & 63))
        case "hello":
            fmt.Println(v2, v3, v5)
        }
        _ = v26
    } else {
        fmt.Println(v3)
        var v28 bool = false
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
