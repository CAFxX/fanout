package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    var v3 uint8 = 127
    var v4 uint32 = 2147483647
    v5 := make(map[uint32]string)
    var v6 int32 = 1
    fmt.Println(v4, v4)
    fmt.Println(v3, v4, v3)
    if (608 == -392) {
        v7 := []int{4294967295, 128}
        var v8 float64 = -2.25
        _ = v7
        _ = v8
    } else {
        v9 := []int{-9223372036854775808, -2147483648}
        switch "foo bar" {
        case "\n\t":
            if (255 == 536) {
                for fi10 := 0; fi10 < 19; fi10++ {
                    v3 += (^(v3))
                    v11 := GenS1{F0: 127, F1: 349}
                    _ = fi10
                    _ = v11
                }
                v6 = v6
            } else {
                switch 242 {
                case 568:
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
            }
        default:
            v12 := GenS2{F0: 2}
            _ = v12
        }
        if (-1 > 0) {
            switch "foo bar" {
            case "":
                if (!(true && false)) {
                    fmt.Println(v6, v6, v6)
                    if false {
                        for fi13 := 0; fi13 < 1; fi13++ {
                            v14 := GenS1{F0: -48, F1: -9223372036854775808}
                            _ = fi13
                            _ = v14
                        }
                    }
                    v15 := GenS3{F0: false, F1: 1, F2: 38}
                    _ = v15
                } else {
                    v4 -= 73
                    var v16 uint8 = v3
                    _ = v16
                }
            default:
                switch -1 {
                case 774:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        switch ("hello" + "\n\t") {
                        case "":
                            switch "a" {
                            case "foo bar":
                                v18 := GenS2{F0: 2}
                                _ = v18
                            default:
                                var v19 string = ("hello" + "hello")
                                _ = v19
                            }
                        case "a":
                            v3 += (^(v3))
                        case "\n\t":
                            _ = len(v2)
                        }
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v4 += v4
                            var v21 float32 = 0.0
                            v21 = (v21 - v21)
                            _ = fi20
                            _ = v21
                        }
                        var v22 string = ("a" + "a")
                        _ = fi17
                        _ = v22
                    }
                case 42:
                    var v23 float64 = (-2.25 + (0.0 * (3.14159 * ((0.1 + -2.25) / ((-2.25 + 1e10) * 3.14159)))))
                    _ = v23
                case -361:
                    if (true && true) {
                        if (!false) {
                            for fi24 := 0; fi24 < 6; fi24++ {
                                v4 = uint32(v3)
                                fmt.Println(v6)
                                var v25 string = ("foo bar" + "hello")
                                _ = fi24
                                _ = v25
                            }
                            var v26 uint16 = 331
                            _ = v26
                        } else {
                            v6 += v6
                            var v27 int = 9223372036854775807
                            v3 += 0
                            _ = v27
                        }
                        switch "a" {
                        case "foo bar":
                            for fi28 := 0; fi28 < 7; fi28++ {
                                v9 = append(v9, int(v6))
                                switch fi28 {
                                case 127:
                                    v6 -= 42
                                case -860:
                                    fmt.Println(v3)
                                default:
                                    var v29 int = -223
                                    _ = v29
                                }
                                v6 -= -187
                                _ = fi28
                            }
                        default:
                            v30 := GenS3{F0: true, F1: 42, F2: 42}
                            _ = v30
                        }
                    }
                }
            }
            fmt.Println(v3)
        } else {
            var v31 int64 = 492
            var v32 float64 = (((0.1 + 1e10) * 3.14159) / -2.25)
            _ = v31
            _ = v32
        }
        _ = v9
    }
    var v33 uint64 = 4294967295
    var v34 int64 = 0
    var v35 int8 = 0
    switch -1 {
    case 390:
        v35 -= (v35 * 42)
    case 1:
        v36 := GenS1{F0: v35, F1: (v34 << ((694) & 63))}
        _ = v36
    default:
        var v37 uint64 = (v33 ^ (^(v33)))
        _ = v37
    }
    v38 := []int{-187, 32767}
    if len(v38) > 0 {
        v38[0] = -1
    }
    var v39 float32 = 0.1
    var v40 float32 = v39
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
