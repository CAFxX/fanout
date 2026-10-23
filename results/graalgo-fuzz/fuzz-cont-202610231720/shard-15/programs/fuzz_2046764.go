package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int; F1 int16 }
func genFn0(p0 string) int8 {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 127
}

func main() {
    var v1 int = -98
    var v2 string = "a"
    v3 := make(map[string]string)
    v4 := []uint32{65535, 4294967295, 2}
    var v5 float64 = ((3.14159 - 0.0) + 0.1)
    var v6 uint64 = 0
    var v7 uint = 181
    var v8 int64 = 9223372036854775807
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint16 = 973
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(v8, v5, fi9)
            if ((-(v8)) != 0) {
                var v12 uint16 = 45
                _ = v12
            }
            var v13 int8 = 100
            switch (("hello" + v2) + (v2 + (("" + "\n\t") + "hello"))) {
            case "foo bar":
                v14 := GenS1{F0: (v10 * uint16(v7)), F1: 240}
                _ = v14
            case "":
                if len(v4) > 0 {
                    v4[0] = 2147483648
                }
            case "a":
                v8 = (v8 ^ 1)
            }
            v15 := GenS1{F0: (^(v10)), F1: (v6 & v6)}
            _ = fi11
            _ = v13
            _ = v15
        }
        var v16 int8 = genFn0((v2 + "a"))
        _ = fi9
        _ = v10
        _ = v16
    }
    var v17 uint16 = 920
    fmt.Println(v6, v7, v6)
    v18 := GenS1{F0: (v17 & 1), F1: 18446744073709551615}
    if (false && false) {
        for fi19 := 0; fi19 < 3; fi19++ {
            switch v2 {
            case "a":
                v7 -= (v7 ^ (v7 - v7))
            }
            for fi20 := 0; fi20 < 11; fi20++ {
                if (((v17 | (v17 ^ (^(v17)))) >= v17) && true) {
                    fmt.Println(v1, v1, v7)
                    var v21 float32 = -2.25
                    var v22 float64 = v5
                    _ = v21
                    _ = v22
                }
                switch v1 {
                case -9223372036854775808:
                    if true {
                        v1 -= (v1 + fi19)
                        fmt.Println(v2, v5)
                    } else {
                        var v23 uint8 = 42
                        v24 := GenS2{F0: (^(fi19)), F1: -720}
                        _ = v23
                        _ = v24
                    }
                }
                var v25 uint32 = 65535
                _ = fi20
                _ = v25
            }
            v26 := GenS1{F0: (^(v17)), F1: 29}
            _ = fi19
            _ = v26
        }
    } else {
        if (0 < 293) {
            for fi27 := 0; fi27 < 14; fi27++ {
                delete(v3, ("foo bar" + ""))
                _ = fi27
            }
        } else {
            if false {
                if false {
                    for fi28 := 0; fi28 < 1; fi28++ {
                        var v29 float32 = -2.25
                        fmt.Println(v6, fi28, v2)
                        _ = fi28
                        _ = v29
                    }
                } else {
                    if len(v4) > 0 {
                        v4[0] = 2147483648
                    }
                    for fi30 := 0; fi30 < 14; fi30++ {
                        if false {
                            v2 = v2
                            for fi31 := 0; fi31 < 8; fi31++ {
                                var v32 uint64 = (v6 | v6)
                                _ = fi31
                                _ = v32
                            }
                        } else {
                            fmt.Println(fi30)
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi30
                    }
                }
            } else {
                delete(v3, "\n\t")
                if len(v4) > 0 {
                    v4[0] = 2
                }
                var v33 uint64 = v6
                _ = v33
            }
        }
    }
    fmt.Println(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
