package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint64; F1 int32; F2 uint8 }
func genFn0() uint {
    var v7 int8 = -89
    v7 += v7
    _ = v7
    return 1
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]string)
    var v3 uint = 2
    var v4 uint32 = 4294967295
    var v5 uint16 = 0
    v6 := make(map[uint32]string)
    if true {
        v3 -= 0
        if false {
            if (340 == 2147483648) {
                v3 = (v3 % ((980) | 1))
                v5 += (^(v5))
                var v8 float64 = 0.1
                _ = v8
            } else {
                switch ("\n\t" + "") {
                case "hello":
                    if false {
                        var v9 string = "hello"
                        _ = v9
                    }
                case "":
                    v3 = (v3 % (((^(v3))) | 1))
                case "a":
                    for fi10 := 0; fi10 < 19; fi10++ {
                        var v11 int16 = int16(v4)
                        _ = fi10
                        _ = v11
                    }
                }
                fmt.Println(v3)
                v12 := []int{-1, 127}
                _ = v12
            }
            v3 -= (^(v3))
        }
        v13 := GenS1{F0: 1, F1: 9223372036854775807}
        _ = v13
    }
    var v14 string = ("\n\t" + ("" + "\n\t"))
    for fi15 := 0; fi15 < 4; fi15++ {
        fmt.Println(v14, v5)
        v16 := GenS1{F0: 683, F1: -332}
        _ = fi15
        _ = v16
    }
    delete(v2, ((v14 + "\n\t") + v14))
    for fi17 := 0; fi17 < 15; fi17++ {
        v4 += v4
        _ = fi17
    }
    fmt.Println(v5, v14, v1)
    v18 := GenS2{F0: v4, F1: v3}
    v6[1] = v14
    v19 := []int{}
    v20 := GenS2{F0: v4, F1: 1}
    if false {
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 uint32 = (^(v4))
            v1 = (v1 - v1)
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 19; fi23++ {
            if false {
                v5 = (^(v5))
                fmt.Println(fi23, v3)
                var v24 int32 = int32(v3)
                _ = v24
            }
            if false {
                fmt.Println(v4, v14, v4)
                for fi25 := 0; fi25 < 2; fi25++ {
                    _ = len(v19)
                    _ = fi25
                }
            } else {
                switch fi23 {
                case -868:
                    v26 := GenS2{F0: uint32(v5), F1: (v3 * v3)}
                    _ = v26
                default:
                    var v27 int16 = -108
                    _ = v27
                }
            }
            _ = fi23
        }
    } else {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        v28 := GenS3{F0: 18446744073709551615, F1: 0, F2: 206}
        _ = v28
    }
    v3 = 108
    for fi29 := 0; fi29 < 1; fi29++ {
        v5 += v5
        switch ("" + v14) {
        case "foo bar":
            if false {
                for fi30 := 0; fi30 < 8; fi30++ {
                    switch -945 {
                    case -2147483648:
                        fmt.Println(fi30, fi29, v1)
                    case 341:
                        if (!true) {
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        }
                    }
                    if true {
                        v14 = (v14 + v14)
                        v31 := GenS2{F0: 878, F1: (^(v3))}
                        if (!((111 < -127) || true)) {
                            switch int(v3) {
                            case 2147483647:
                                _ = len(v2)
                            case 394:
                                var v32 int = fi30
                                _ = v32
                            }
                            v14 = "hello"
                            v4 += v4
                        }
                        _ = v31
                    }
                    fmt.Println(fi29)
                    v14 = v14
                    var v33 string = v14
                    delete(v2, v33)
                    _ = fi30
                    _ = v33
                }
            }
        case "hello":
            var v34 float32 = v1
            _ = v34
        case "":
            v5 = 32767
        }
        _ = fi29
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        var v36 uint16 = 0
        fmt.Println(v14, v14, v1)
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
