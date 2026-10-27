package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint8 }
func genFn0() uint64 {
    if true {
        var v4 uint8 = 248
        v4 += v4
        fmt.Println(v4)
        _ = v4
    }
    var v5 int8 = 100
    var v6 string = "\n\t"
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 42
    v3 := []string{"\n\t", "a", "\n\t"}
    v2 = (v2 << ((730) & 63))
    if ((-1 < 100) && false) {
        v3 = append(v3, (("" + ("foo bar" + "foo bar")) + ("a" + ("foo bar" + "hello"))))
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 uint32 = 455
                switch (fi7 % ((-109) | 1)) {
                case -1:
                    v9 += (v9 / ((v9) | 1))
                }
                v2 += v2
                var v10 uint32 = (^(v9))
                v2 = v2
                fmt.Println(v10, fi8)
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := GenS1{F0: 849}
            for fi13 := 0; fi13 < 18; fi13++ {
                v2 -= 17
                v2 -= (v2 * 18446744073709551615)
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
    } else {
        v2 = (v2 & v2)
        v2 = (v2 << ((2) & 63))
    }
    fmt.Println(v2)
    var v14 uint64 = v2
    var v15 uint8 = 128
    if false {
        v16 := GenS2{F0: v2, F1: 0, F2: (v15 - v15)}
        var v17 uint32 = 617
        var v18 uint = 400
        _ = v16
        _ = v17
        _ = v18
    } else {
        if false {
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 uint64 = v2
                fmt.Println(fi19, fi19)
                _ = fi19
                _ = v20
            }
        } else {
            switch 2147483647 {
            case 567:
                v2 = (^(v2))
            case 258:
                if (v15 != (v15 - v15)) {
                    fmt.Println(v15)
                } else {
                    v15 = 127
                    _ = len(v3)
                    var v21 float32 = (0.1 - 1e10)
                    _ = v21
                }
            }
        }
        var v22 bool = false
        _ = v22
    }
    v2 = genFn0()
    for fi23 := 0; fi23 < 20; fi23++ {
        switch "\n\t" {
        case "\n\t":
            switch ("\n\t" + (("" + ("hello" + "a")) + "\n\t")) {
            case "foo bar":
                if false {
                    fmt.Println(v14, v14, v15)
                    fmt.Println(v2)
                }
            case "hello":
                switch "hello" {
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "foo bar":
                    if (65535 > 913) {
                        for fi24 := 0; fi24 < 15; fi24++ {
                            var v25 string = ("\n\t" + "foo bar")
                            _ = fi24
                            _ = v25
                        }
                        var v26 float32 = ((-2.25 + 0.1) - float32(((3.14159 - (float64(1e10) - 1.5)) + -2.25)))
                        fmt.Println(v14, v15, fi23)
                        _ = v26
                    }
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        default:
            _ = len(v3)
        }
        v27 := GenS2{F0: v14, F1: 469, F2: v15}
        _ = fi23
        _ = v27
    }
    if true {
        fmt.Println(v14, v14, v2)
    }
    if (true || (3.14159 >= 1e10)) {
        switch "" {
        case "foo bar":
            v2 = v2
        case "hello":
            for fi28 := 0; fi28 < 11; fi28++ {
                v14 = (v2 % (((^(v14))) | 1))
                if true {
                    var v29 string = "hello"
                    for fi30 := 0; fi30 < 9; fi30++ {
                        delete(v1, 856)
                        _ = fi30
                    }
                    _ = v29
                }
                _ = fi28
            }
        }
        if ((7 > 42) || (!(uint16(v15) < 525))) {
            var v31 uint8 = (v15 % ((v15) | 1))
            _ = v31
        }
        var v32 bool = true
        _ = v32
    } else {
        fmt.Println(v14, v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
