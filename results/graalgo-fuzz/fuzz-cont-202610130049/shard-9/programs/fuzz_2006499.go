package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 uint; F2 uint8 }

func main() {
    var v1 uint8 = 127
    var v2 string = ("foo bar" + "hello")
    var v3 int32 = 349
    v4 := []string{v2, "hello"}
    v5 := make(map[uint32]uint32)
    v3 = (v3 / ((v3) | 1))
    v6 := GenS1{F0: 52, F1: 333}
    var v7 uint8 = 1
    v4 = append(v4, v2)
    var v8 int16 = 1
    var v9 uint16 = 298
    if (!((4294967295 != 42) && false)) {
        if true {
            var v10 float32 = (1e10 * -2.25)
            _ = v10
        } else {
            switch v2 {
            case "":
                switch 128 {
                case 4294967295:
                    if (true && false) {
                        for fi11 := 0; fi11 < 3; fi11++ {
                            v12 := GenS2{F0: 1, F1: (v9 + v9)}
                            _ = fi11
                            _ = v12
                        }
                    } else {
                        var v13 uint8 = v7
                        var v14 int32 = (^(v3))
                        _ = v13
                        _ = v14
                    }
                case 604:
                    fmt.Println(v8, v2)
                default:
                    var v15 uint16 = 65535
                    _ = v15
                }
            case "a":
                for fi16 := 0; fi16 < 19; fi16++ {
                    delete(v5, 2147483647)
                    switch fi16 {
                    case 227:
                        v17 := GenS3{F0: 0, F1: int16(v1)}
                        _ = v17
                    case 128:
                        var v18 int8 = 117
                        _ = v18
                    }
                    _ = fi16
                }
            }
        }
        switch ("hello" + (v2 + ("a" + "a"))) {
        case "a":
            var v19 int = 2147483647
            _ = v19
        case "foo bar":
            if true {
                for fi20 := 0; fi20 < 16; fi20++ {
                    v21 := GenS2{F0: 239, F1: (v9 % ((280) | 1))}
                    _ = fi20
                    _ = v21
                }
                var v22 uint16 = v9
                _ = v22
            }
        }
    } else {
        if false {
            for fi23 := 0; fi23 < 12; fi23++ {
                v24 := GenS1{F0: v1, F1: 0}
                var v25 string = v2
                for fi26 := 0; fi26 < 17; fi26++ {
                    if ((v3 / ((0) | 1)) != (v3 & v3)) {
                        switch ("\n\t" + ((v25 + "\n\t") + "foo bar")) {
                        case "foo bar":
                            v9 = (v9 & (v9 * v9))
                            v9 -= 32767
                            var v27 uint = 521
                            v7 += uint8(v9)
                            v28 := GenS4{F0: (v3 - int32(v8))}
                            _ = v27
                            _ = v28
                        case "\n\t":
                            delete(v5, 894)
                        }
                    } else {
                        fmt.Println(v8, v1)
                        var v29 int16 = v8
                        v2 = v2
                        _ = v29
                    }
                    _ = fi26
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            v9 = (v9 ^ (v9 ^ v9))
            fmt.Println(v2)
        }
        for fi30 := 0; fi30 < 2; fi30++ {
            fmt.Println(fi30)
            _ = fi30
        }
    }
    v8 = (-(v8))
    v3 -= v3
    var v31 int32 = (v3 + v3)
    var v32 float64 = float64((((1e10 * ((-2.25 - (1e10 + 0.1)) - 1e10)) / 3.14159) + float32((((0.0 + float64(1.5)) + 1.5) / (3.14159 / 1e10)))))
    var v33 int8 = 1
    var v34 int = -19
    v3 -= 1
    v35 := GenS5{F0: v3, F1: 2, F2: (v7 * (^(v1)))}
    if (v33 > v33) {
        var v36 uint = 4294967295
        v37 := GenS4{F0: v3}
        _ = v36
        _ = v37
    }
    var v38 int = v34
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
