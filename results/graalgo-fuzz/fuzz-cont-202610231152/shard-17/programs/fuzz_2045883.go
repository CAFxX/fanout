package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }
func genFn0() uint8 {
    v6 := GenS1{F0: true, F1: 255, F2: 18446744073709551615}
    var v7 int8 = 42
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 int64 = -725
    v2 := make(map[int]bool)
    var v3 int32 = -1
    var v4 uint16 = 0
    var v5 int64 = (v1 >> ((685) & 63))
    v8 := GenS2{F0: 2, F1: 1}
    switch ("\n\t" + ("" + "a")) {
    case "\n\t":
        if (true && (!(!true))) {
            var v9 uint8 = 128
            v4 += v4
            _ = v9
        } else {
            if false {
                v4 = 339
                v4 += (v4 >> ((295) & 63))
            } else {
                var v10 int8 = 127
                _ = v10
            }
            for fi11 := 0; fi11 < 7; fi11++ {
                for fi12 := 0; fi12 < 13; fi12++ {
                    if false {
                        if ((-(v1)) < v5) {
                            v1 += (v5 & 0)
                            var v13 uint64 = uint64(v4)
                            _ = v13
                        } else {
                            if (0 > -1) {
                                v2[fi11] = false
                                if (1 == 4294967295) {
                                    v1 += 1
                                } else {
                                    if true {
                                        if (1 > 42) {
                                            if true {
                                                v14 := GenS2{F0: 278, F1: 267}
                                                switch -2147483648 {
                                                case 0:
                                                    var v15 int = (fi11 * (fi12 - int(v3)))
                                                    _ = v15
                                                }
                                                _ = v14
                                            } else {
                                                var v16 bool = true
                                                fmt.Println(v3, v1, fi11)
                                                v5 += (v5 >> ((65535) & 63))
                                                _ = v16
                                            }
                                            var v17 uint32 = 4294967295
                                            _ = v17
                                        }
                                    } else {
                                        v3 += (v3 >> ((0) & 63))
                                        v5 -= -9223372036854775808
                                    }
                                }
                                if (!(false && false)) {
                                    fmt.Println(v5)
                                    v5 = v1
                                }
                            } else {
                                delete(v2, -322)
                                v18 := []int{fi11}
                                _ = v18
                            }
                            var v19 int64 = v1
                            _ = v19
                        }
                        delete(v2, fi12)
                    }
                    v20 := v2[(fi12 % ((715) | 1))]
                    switch "\n\t" {
                    case "hello":
                        var v21 uint64 = 342
                        _ = v21
                    }
                    _ = fi12
                    _ = v20
                }
                var v22 int32 = 0
                var v23 float64 = 1.5
                _ = fi11
                _ = v22
                _ = v23
            }
            v1 = (v1 % ((v1) | 1))
        }
    }
    v24 := GenS1{F0: false, F1: 2, F2: 523}
    fmt.Println(v1)
    v3 += (^(v3))
    v25 := GenS3{F0: 128}
    for fi26 := 0; fi26 < 15; fi26++ {
        _ = len(v2)
        if true {
            v1 = (v5 % ((v5) | 1))
            var v27 float32 = (3.14159 / 0.1)
            _ = v27
        } else {
            switch fi26 {
            case 2147483647:
                var v28 float64 = float64(1.5)
                _ = v28
            case 169:
                fmt.Println(v3, v1, v5)
            case -1:
                v4 = v4
            default:
                if ((439 >= 42) && ((!(595 == 2)) || false)) {
                    v29 := GenS4{F0: 32767}
                    _ = v29
                }
            }
        }
        v2[(fi26 * fi26)] = (105 == 2)
        _ = fi26
    }
    switch 4294967295 {
    case 128:
        var v30 uint64 = 18446744073709551615
        _ = v30
    }
    if false {
        var v31 uint64 = 2
        v32 := GenS2{F0: 2147483648, F1: 42}
        var v33 uint8 = 82
        _ = v31
        _ = v32
        _ = v33
    } else {
        var v34 int64 = v1
        _ = v34
    }
    fmt.Println(v5, v5, v5)
    v4 += 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
