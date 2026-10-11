package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint32 }

func main() {
    var v1 int = 128
    var v2 float32 = 1e10
    var v3 int = (^(v1))
    v4 := []int{v3, (v1 << ((1) & 63)), v1, (v1 | v1)}
    var v5 int = v1
    for fi6 := 0; fi6 < 12; fi6++ {
        v3 = (v1 + -697)
        if (((1e10 / 1e10) != v2) && (2 > 2)) {
            if (true && (2147483647 >= -1)) {
                for fi7 := 0; fi7 < 14; fi7++ {
                    if (true && false) {
                        var v8 uint = 42
                        _ = len(v4)
                        fmt.Println(v5, fi6)
                        if (-16 >= 123) {
                            fmt.Println(v8, v3, v2)
                        } else {
                            var v9 int16 = -11
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if len(v4) > 0 {
                                v4[0] = -9223372036854775808
                            }
                            _ = v9
                        }
                        _ = v8
                    }
                    _ = fi7
                }
                for fi10 := 0; fi10 < 9; fi10++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v11 := GenS1{F0: -128}
                    v12 := GenS2{F0: true}
                    fmt.Println(v5, v2)
                    fmt.Println(v5, fi10, v2)
                    v5 = (fi10 << ((2) & 63))
                    if (9223372036854775807 == 1) {
                        v13 := GenS2{F0: false}
                        fmt.Println(fi10, v5, v5)
                        _ = v13
                    } else {
                        if (791 <= 21) {
                            v4 = append(v4, (fi10 | fi10))
                            var v14 int = v5
                            fmt.Println(v1, v14)
                            fmt.Println(v3, fi6, v3)
                            v3 = (v5 ^ fi6)
                            _ = v14
                        }
                        v4 = append(v4, (v3 + (fi6 ^ (v5 - -2147483648))))
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v15 bool = (1e10 >= (v2 / (v2 * v2)))
                _ = v15
            } else {
                v5 = v1
                v2 -= v2
            }
            switch "\n\t" {
            case "hello":
                v1 = v1
            default:
                v1 = v3
            }
            switch v3 {
            case 32768:
                var v16 uint32 = 4294967295
                _ = v16
            case 127:
                var v17 uint = 65535
                _ = v17
            case -641:
                v5 = 0
            default:
                if (true && true) {
                    for fi18 := 0; fi18 < 6; fi18++ {
                        switch ("hello" + ("a" + ((("hello" + ("hello" + "foo bar")) + "hello") + ""))) {
                        case "":
                            var v19 int8 = 1
                            _ = v19
                        case "foo bar":
                            v20 := GenS2{F0: (508 != 2147483647)}
                            _ = v20
                        default:
                            if len(v4) > 0 {
                                v4[0] = -257
                            }
                        }
                        var v21 int64 = 846
                        v21 -= v21
                        _ = fi18
                        _ = v21
                    }
                    var v22 int16 = -1
                    _ = v22
                } else {
                    switch v1 {
                    case 128:
                        if (1 > 0) {
                            var v23 bool = ("\n\t" < ("\n\t" + ("\n\t" + "hello")))
                            if len(v4) > 0 {
                                v4[0] = v5
                            }
                            _ = v23
                        }
                    case -528:
                        if (false && false) {
                            v5 = (v5 & fi6)
                            fmt.Println(fi6)
                            v1 = (v1 - (^(v3)))
                        }
                    case 127:
                        if false {
                            v2 += v2
                            var v24 uint = 255
                            var v25 int16 = 1
                            _ = v24
                            _ = v25
                        } else {
                            fmt.Println(v3, v1, v3)
                            var v26 uint32 = 65535
                            _ = v26
                        }
                    }
                    var v27 uint16 = 32767
                    switch fi6 {
                    case 42:
                        if (120 < 127) {
                            v28 := GenS3{F0: 169}
                            v29 := GenS1{F0: -128}
                            v30 := GenS2{F0: true}
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    }
                    _ = v27
                }
            }
        } else {
            v2 -= v2
            v31 := GenS4{F0: 465, F1: 42, F2: 2147483647}
            _ = v31
        }
        _ = fi6
    }
    _ = len(v4)
    if len(v4) > 0 {
        v4[0] = v1
    }
    var v32 uint64 = 867
    var v33 uint64 = (^(v32))
    var v34 int16 = -526
    var v35 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
