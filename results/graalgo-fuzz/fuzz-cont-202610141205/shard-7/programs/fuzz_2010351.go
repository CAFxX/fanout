package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
func genFn0(p0 uint16) uint {
    v4 := []int{-9223372036854775808, 32767, -1, 9223372036854775807}
    v5 := GenS1{F0: 0, F1: -1, F2: -1}
    if len(v4) > 0 {
        v4[0] = -313
    }
    v6 := GenS2{F0: 1, F1: uint32(p0), F2: 42}
    if true {
        var v7 string = "hello"
        v7 = (v7 + v7)
        var v8 uint64 = 0
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 195
}

func main() {
    var v1 int16 = -32768
    var v2 int8 = 1
    var v3 int = 4294967295
    var v9 float64 = (3.14159 / 3.14159)
    switch (v3 << ((18446744073709551615) & 63)) {
    case 127:
        v3 += (v3 >> ((867) & 63))
    default:
        fmt.Println(v1, v3)
    }
    v1 = (v1 % ((v1) | 1))
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 float64 = (v9 / -2.25)
        if false {
            fmt.Println(v11, v3)
            v2 = (v2 - v2)
        } else {
            switch fi10 {
            case -45:
                v9 -= v9
            }
            v2 = -51
        }
        fmt.Println(v11)
        _ = fi10
        _ = v11
    }
    var v12 uint64 = 2
    if ((float64(-2.25) < float64(-2.25)) || true) {
        switch "" {
        case "\n\t":
            v3 = 88
        case "":
            switch 32767 {
            case 2147483647:
                var v13 uint32 = 4294967295
                _ = v13
            default:
                if (!(false || true)) {
                    var v14 uint8 = 167
                    for fi15 := 0; fi15 < 9; fi15++ {
                        switch "foo bar" {
                        case "foo bar":
                            var v16 uint64 = v12
                            _ = v16
                        case "":
                            for fi17 := 0; fi17 < 7; fi17++ {
                                v18 := []int{(fi15 >> ((42) & 63)), (^(fi15))}
                                v1 -= v1
                                _ = fi17
                                _ = v18
                            }
                        }
                        _ = fi15
                    }
                    for fi19 := 0; fi19 < 9; fi19++ {
                        if (v1 <= 42) {
                            var v20 int32 = 80
                            v2 = (v2 << ((653) & 63))
                            _ = v20
                        } else {
                            var v21 bool = true
                            var v22 int = (v3 & (^(v3)))
                            _ = v21
                            _ = v22
                        }
                        _ = fi19
                    }
                    _ = v14
                }
            }
        default:
            switch "\n\t" {
            case "a":
                var v23 uint16 = 42
                _ = v23
            case "":
                for fi24 := 0; fi24 < 5; fi24++ {
                    if true {
                        v9 = ((v9 + v9) * v9)
                        var v25 string = ""
                        _ = v25
                    } else {
                        fmt.Println(v3)
                        v26 := []int{(fi24 ^ (fi24 << ((628) & 63))), 143}
                        _ = v26
                    }
                    fmt.Println(v2, v12)
                    _ = fi24
                }
            case "\n\t":
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    for fi27 := 0; fi27 < 8; fi27++ {
                        v28 := []int{}
                        _ = fi27
                        _ = v28
                    }
                }
            }
        }
        var v29 int8 = v2
        _ = v29
    }
    if false {
        v30 := []int{v3, (^(v3)), (v3 >> ((42) & 63)), (v3 * (v3 / ((9223372036854775807) | 1)))}
        _ = v30
    } else {
        for fi31 := 0; fi31 < 2; fi31++ {
            if (935 >= 32767) {
                v2 = 42
                fmt.Println(v12, v3)
            }
            _ = fi31
        }
        var v32 uint16 = 65535
        _ = v32
    }
    var v33 float32 = 1.5
    v2 = (^(v2))
    var v34 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v9
}
