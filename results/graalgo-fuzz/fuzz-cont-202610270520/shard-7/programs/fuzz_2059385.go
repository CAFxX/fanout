package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint64; F2 bool }
type GenS4 struct { F0 uint16; F1 int64; F2 uint16 }
func genFn0() uint8 {
    v6 := GenS1{F0: 445}
    if ((("a" + (((("foo bar" + ("foo bar" + "foo bar")) + "a") + ("\n\t" + "\n\t")) + (("" + "a") + (("hello" + ("hello" + ("\n\t" + (("" + "a") + "foo bar")))) + "\n\t")))) + "hello") > "foo bar") {
        var v7 int8 = 71
        _ = v7
    } else {
        var v8 float64 = 1.5
        _ = v8
    }
    if (700 == 32767) {
        var v9 uint64 = 104
        for fi10 := 0; fi10 < 2; fi10++ {
            if true {
                switch (^(fi10)) {
                case 1:
                    switch -2147483648 {
                    case 127:
                        if ("a" < "foo bar") {
                            if true {
                                v11 := GenS2{F0: 1, F1: 60, F2: 32767}
                                v9 += (^(v9))
                                _ = v11
                            } else {
                                for fi12 := 0; fi12 < 9; fi12++ {
                                    v9 += (^(v9))
                                    var v13 float32 = (1.5 * 3.14159)
                                    fmt.Println(fi10, v9)
                                    _ = fi12
                                    _ = v13
                                }
                            }
                            v9 = 42
                            fmt.Println(fi10, fi10)
                        }
                    case -1:
                        switch -9223372036854775808 {
                        case 670:
                            var v14 string = ("hello" + "a")
                            _ = v14
                        case -2147483648:
                            v15 := []int{}
                            _ = v15
                        default:
                            v9 -= v9
                        }
                    case 651:
                        v16 := []int{fi10}
                        _ = v16
                    }
                case 42:
                    var v17 bool = (!true)
                    _ = v17
                }
            }
            _ = fi10
        }
        var v18 uint8 = 101
        _ = v18
        _ = v9
    }
    _ = v6
    return 167
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{65535, 6, 42}
    var v3 uint16 = 1
    v4 := make(map[int]uint32)
    var v5 int64 = 9223372036854775807
    v5 = v5
    if (!((((!(!true)) && true) || true) || false)) {
        for fi19 := 0; fi19 < 14; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                v1[65535] = (fi20 << ((134) & 63))
                fmt.Println(v3)
                _ = fi20
            }
            v5 += (v5 << ((18446744073709551615) & 63))
            var v21 uint64 = 42
            _ = fi19
            _ = v21
        }
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        v22 := GenS2{F0: -87, F1: 100, F2: (^(v3))}
        _ = v22
    } else {
        fmt.Println(v5, v5)
    }
    fmt.Println(v3, v5)
    if len(v2) > 0 {
        v2[0] = 1
    }
    var v23 float64 = (1e10 * -2.25)
    var v24 uint32 = 721
    fmt.Println(v3)
    v23 = 3.14159
    var v25 int64 = v5
    switch -715 {
    case -1:
        for fi26 := 0; fi26 < 14; fi26++ {
            if (1 > 128) {
                v5 += v25
            } else {
                fmt.Println(v3, fi26, v3)
                if false {
                    fmt.Println(fi26, v24)
                }
            }
            var v27 uint32 = v24
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi26
            _ = v27
        }
    case -377:
        v3 = (^(v3))
    default:
        v5 += (v25 & v5)
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        v5 = 0
        v29 := GenS3{F0: v25, F1: uint64(v3), F2: false}
        _ = fi28
        _ = v29
    }
    if len(v2) > 0 {
        v2[0] = v24
    }
    fmt.Println(v3, v23, v23)
    switch 1 {
    case 127:
        if false {
            v30 := GenS4{F0: v3, F1: (v25 ^ 42), F2: v3}
            switch 9223372036854775807 {
            case -2147483648:
                switch "hello" {
                case "\n\t":
                    v2 = append(v2, (^(v24)))
                }
            }
            _ = v30
        } else {
            var v31 uint16 = 0
            _ = v31
        }
    default:
        v5 += v25
    }
    fmt.Println(v3)
    var v32 float32 = (0.1 - ((-2.25 - (((((1e10 + ((0.0 - -2.25) / 0.1)) / 3.14159) * 1e10) - (0.1 - (-2.25 - 1e10))) + 0.1)) - 3.14159))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
