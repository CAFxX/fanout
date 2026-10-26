package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 bool }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16; F1 bool; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint8; F2 int8 }

func main() {
    var v1 uint64 = 814
    var v2 uint32 = 42
    var v3 int16 = -32768
    v4 := []string{"a"}
    var v5 uint = 972
    var v6 int16 = (-(v3))
    switch ("hello" + "") {
    case "a":
        if (-1 != int(v2)) {
            switch 9223372036854775807 {
            case -862:
                var v7 uint16 = 2
                _ = v7
            default:
                v1 = 1
            }
            v1 += v1
            v2 += v2
        } else {
            if true {
                var v8 uint32 = v2
                if false {
                    v9 := GenS1{F0: v8, F1: 42, F2: (v5 != (^(v5)))}
                    fmt.Println(v6, v3, v8)
                    _ = v9
                }
                _ = v8
            }
            var v10 uint32 = 2147483648
            _ = v10
        }
    case "\n\t":
        if (!false) {
            fmt.Println(v1, v6, v5)
            v1 = v1
        }
    case "":
        if false {
            fmt.Println(v3)
        }
    }
    var v11 string = (("foo bar" + "foo bar") + "hello")
    v4 = append(v4, v11)
    switch 32768 {
    case -1:
        var v12 float64 = ((((0.0 / float64(((float32(0.0) - ((((float32((0.1 - 0.1)) / (float32(0.1) - 0.0)) / 0.1) + -2.25) / (0.0 + 1e10))) + 3.14159))) / 1.0) / 1.0) - (3.14159 / 1.0))
        _ = v12
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 127 {
    case 9223372036854775807:
        for fi13 := 0; fi13 < 18; fi13++ {
            switch (fi13 * 32768) {
            case -641:
                v11 = (v11 + v11)
            case 9223372036854775807:
                for fi14 := 0; fi14 < 13; fi14++ {
                    v15 := GenS2{F0: 69, F1: v5}
                    var v16 int32 = -1
                    v16 = v16
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 18; fi17++ {
                    if (!(false || (128 != uint8(v2)))) {
                        if true {
                            fmt.Println(v11, v2)
                            var v18 uint = v5
                            _ = v18
                        }
                        switch v11 {
                        case "foo bar":
                            v19 := GenS1{F0: v2, F1: -903, F2: false}
                            _ = v19
                        case "a":
                            fmt.Println(v5, v6)
                        }
                        v20 := GenS1{F0: v2, F1: 1, F2: (false && (((!(true || true)) || false) || false))}
                        _ = v20
                    } else {
                        if ((("a" + v11) + (v11 + ("hello" + "foo bar"))) == v11) {
                            fmt.Println(v3, fi13, v11)
                            var v21 int16 = (-(v3))
                            if false {
                                v5 -= v5
                            }
                            _ = v21
                        } else {
                            var v22 bool = true
                            if true {
                                v1 += (v1 - (^(v1)))
                            } else {
                                v23 := GenS2{F0: 1, F1: uint(v2)}
                                _ = v23
                            }
                            _ = v22
                        }
                    }
                    var v24 uint32 = (v2 - v2)
                    _ = fi17
                    _ = v24
                }
            }
            _ = fi13
        }
    case -9223372036854775808:
        if true {
            for fi25 := 0; fi25 < 16; fi25++ {
                switch (fi25 | (fi25 * fi25)) {
                case 32767:
                    v1 = 122
                }
                _ = fi25
            }
        }
    case 927:
        switch -2147483648 {
        case -61:
            var v26 uint32 = (^(v2))
            _ = v26
        default:
            for fi27 := 0; fi27 < 3; fi27++ {
                for fi28 := 0; fi28 < 18; fi28++ {
                    v11 = (v11 + (v11 + ""))
                    v29 := GenS3{F0: v3, F1: false, F2: -128}
                    switch ((v11 + (v11 + ((((v11 + v11) + v11) + v11) + v11))) + "foo bar") {
                    case "hello":
                        v2 += (v2 / (((v2 * v2)) | 1))
                    case "":
                        v30 := GenS3{F0: v3, F1: true, F2: 42}
                        _ = v30
                    default:
                        v31 := GenS3{F0: (v6 << ((v5) & 63)), F1: (fi28 == (fi27 + 32767)), F2: -80}
                        _ = v31
                    }
                    _ = fi28
                    _ = v29
                }
                _ = fi27
            }
        }
    default:
        for fi32 := 0; fi32 < 10; fi32++ {
            var v33 uint32 = 135
            if (false || (1 < 42)) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (!true) {
                    v1 += v1
                }
            } else {
                if true {
                    v34 := GenS4{F0: 53, F1: 154, F2: int8(v6)}
                    _ = v34
                }
                var v35 uint64 = v1
                _ = v35
            }
            _ = fi32
            _ = v33
        }
    }
    v5 = (v5 * (^(v5)))
    v2 = (v2 & (v2 | (^(v2))))
    v11 = ((v11 + v11) + "a")
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
