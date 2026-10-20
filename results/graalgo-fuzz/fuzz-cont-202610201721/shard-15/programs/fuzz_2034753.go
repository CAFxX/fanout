package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int32 = 174
    var v2 uint8 = 128
    var v3 int64 = -1
    var v4 float64 = 1e10
    var v5 int16 = 722
    v6 := make(map[string]bool)
    v5 -= v5
    v2 = (v2 - v2)
    if false {
        v7 := []int{-433, 127, -9223372036854775808}
        var v8 int = 127
        var v9 uint16 = 65535
        _ = v7
        _ = v8
        _ = v9
    }
    v3 += (v3 >> ((1) & 63))
    v4 = 1e10
    if false {
        v5 = 429
        var v10 int16 = 1
        _ = v10
    }
    v5 = (^(v5))
    switch 227 {
    case 633:
        switch -1 {
        case 42:
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int64 = v3
                _ = fi11
                _ = v12
            }
        case 1:
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 uint64 = uint64(v2)
                _ = fi13
                _ = v14
            }
        }
    }
    if (0 != 1) {
        fmt.Println(v2)
        fmt.Println(v2, v2, v4)
    } else {
        v2 = v2
        var v15 int16 = (v5 / (((v5 | 1)) | 1))
        var v16 float32 = (float32(1.5) * ((1.5 / 3.14159) - ((1.5 / float32(((v4 + 0.0) + 3.14159))) / 3.14159)))
        _ = v15
        _ = v16
    }
    switch -9223372036854775808 {
    case -1:
        if true {
            v17 := GenS1{F0: -55, F1: (v1 & v1)}
            v5 -= 939
            _ = v17
        } else {
            if true {
                var v18 string = ("hello" + "a")
                if (361 == 738) {
                    fmt.Println(v18, v3)
                    v3 -= (-(v3))
                }
                _ = v18
            } else {
                for fi19 := 0; fi19 < 2; fi19++ {
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v21 := []int{}
                        _ = fi20
                        _ = v21
                    }
                    _ = fi19
                }
                var v22 int8 = -12
                _ = v22
            }
            var v23 uint16 = 57
            fmt.Println(v2)
            _ = v23
        }
    case 42:
        if (1 <= -128) {
            if ((!(!(false && (4294967295 != 988)))) || (475 >= 452)) {
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 bool = (556 != 674)
                    v26 := GenS1{F0: (^(v3)), F1: 1}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                switch "a" {
                case "a":
                    for fi27 := 0; fi27 < 7; fi27++ {
                        var v28 uint32 = 4294967295
                        _ = fi27
                        _ = v28
                    }
                default:
                    v2 = (v2 / ((v2) | 1))
                }
                v5 = -302
            } else {
                switch "a" {
                case "\n\t":
                    for fi29 := 0; fi29 < 2; fi29++ {
                        if ("hello" <= (("foo bar" + "foo bar") + "\n\t")) {
                            if (!(1 >= -128)) {
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    var v31 string = "hello"
                                    _ = fi30
                                    _ = v31
                                }
                                v32 := GenS2{F0: (v5 / (((-(v5))) | 1))}
                                _ = v32
                            } else {
                                v3 -= int64(v2)
                                var v33 int8 = 63
                                _ = v33
                            }
                            v34 := GenS3{F0: v3}
                            var v35 int8 = 100
                            _ = v34
                            _ = v35
                        }
                        v36 := GenS1{F0: 0, F1: (-(v1))}
                        _ = fi29
                        _ = v36
                    }
                case "":
                    if false {
                        fmt.Println(v4, v3)
                        var v37 float32 = (1.5 + (1e10 + ((1.5 + (float32(0.0) + (0.0 + 0.1))) - 1.5)))
                        _ = v37
                    } else {
                        v38 := GenS1{F0: v3, F1: -371}
                        _ = v38
                    }
                }
                v39 := GenS1{F0: (^(v3)), F1: v1}
                _ = v39
            }
        } else {
            var v40 int64 = (v3 + v3)
            _ = v40
        }
    case -2147483648:
        var v41 int64 = (v3 >> ((65535) & 63))
        _ = v41
    default:
        var v42 int16 = (v5 << ((2) & 63))
        _ = v42
    }
    for fi43 := 0; fi43 < 7; fi43++ {
        v1 += -427
        _ = fi43
    }
    var v44 float32 = float32((float64((-2.25 * (3.14159 - (-2.25 - float32(v4))))) + 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
