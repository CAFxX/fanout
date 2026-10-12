package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 78
    var v3 uint8 = 255
    var v4 uint8 = 2
    v5 := []uint32{}
    if (38 < 32767) {
        var v6 uint16 = 0
        _ = v6
    } else {
        var v7 int32 = 42
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v8 bool = true
        _ = v7
        _ = v8
    }
    var v9 string = "foo bar"
    v9 = "a"
    v3 = v3
    var v10 string = v9
    switch v9 {
    case "foo bar":
        if (2147483647 <= 881) {
            fmt.Println(v2, v4, v2)
        } else {
            switch 32768 {
            case 9223372036854775807:
                if (v2 < v2) {
                    v10 = v10
                } else {
                    if (42 == 2) {
                        if false {
                            for fi11 := 0; fi11 < 19; fi11++ {
                                v10 = ""
                                _ = fi11
                            }
                            var v12 uint16 = 2
                            _ = v12
                        } else {
                            v13 := v1[4294967295]
                            _ = v13
                        }
                    } else {
                        for fi14 := 0; fi14 < 3; fi14++ {
                            v10 = v9
                            v15 := GenS1{F0: 18446744073709551615}
                            _ = fi14
                            _ = v15
                        }
                    }
                }
            case -720:
                fmt.Println(v4, v9)
            case -9223372036854775808:
                v16 := v1[781]
                _ = v16
            default:
                fmt.Println(v4, v10)
            }
            var v17 uint64 = 73
            _ = v17
        }
    case "hello":
        switch 128 {
        case -603:
            v4 = (v4 << ((18446744073709551615) & 63))
        default:
            _ = len(v5)
        }
    }
    var v18 uint32 = 42
    if len(v5) > 0 {
        v5[0] = (v18 >> ((0) & 63))
    }
    v10 = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
