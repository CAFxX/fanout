package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int8; F2 uint64 }
func genFn0() int16 {
    var v6 uint16 = 255
    var v7 uint32 = 674
    fmt.Println(v7, v6)
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 bool, p1 uint32) int32 {
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: 2, F1: 0}
        switch -2147483648 {
        case 812:
            v10 := []int{fi8, fi8, (fi8 ^ (fi8 & (fi8 & 2147483647)))}
            _ = v10
        case 32768:
            switch ("" + "") {
            case "a":
                p0 = (("" + "a") > ("hello" + "hello"))
            default:
                var v11 int = int(p1)
                _ = v11
            }
        case -1:
            v12 := []int{fi8}
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    var v13 int64 = -234
    _ = p0
    _ = p1
    _ = v13
    return -2147483648
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -32768
    var v3 uint64 = 428
    v4 := make(map[uint32]bool)
    var v5 string = "\n\t"
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 string = (v5 + v5)
        _ = fi14
        _ = v15
    }
    fmt.Println(v3)
    if true {
        v16 := GenS1{F0: 1, F1: (v3 ^ 2)}
        _ = v16
    } else {
        if false {
            for fi17 := 0; fi17 < 15; fi17++ {
                fmt.Println(v2, fi17)
                var v18 int16 = v2
                if false {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        switch fi17 {
                        case -2147483648:
                            switch fi17 {
                            case -9223372036854775808:
                                v2 = (v18 << ((354) & 63))
                            case -966:
                                v20 := GenS1{F0: 18446744073709551615, F1: (^(v3))}
                                delete(v1, (fi17 + 127))
                                _ = v20
                            case -304:
                                var v21 uint64 = v3
                                _ = v21
                            default:
                                v18 = v2
                            }
                        default:
                            var v22 float32 = (1.5 - 0.0)
                            _ = v22
                        }
                        v5 = "foo bar"
                        _ = fi19
                    }
                }
                _ = fi17
                _ = v18
            }
            _ = len(v1)
            var v23 int8 = -11
            _ = v23
        }
    }
    if false {
        v24 := GenS2{F0: 255, F1: v3}
        _ = v24
    } else {
        switch v5 {
        case "a":
            v2 -= -189
        case "hello":
            switch 583 {
            case 1:
                v1[32768] = (100 >= 127)
            default:
                v25 := GenS1{F0: 65535, F1: 42}
                _ = v25
            }
        case "foo bar":
            for fi26 := 0; fi26 < 13; fi26++ {
                fmt.Println(v5, fi26, v5)
                fmt.Println(v3)
                switch fi26 {
                case -2147483648:
                    switch v5 {
                    case "\n\t":
                        var v27 uint16 = 0
                        _ = v27
                    default:
                        switch (fi26 ^ -724) {
                        case 4294967295:
                            v5 = v5
                        case 608:
                            v28 := v4[745]
                            _ = v28
                        }
                    }
                case -934:
                    if (false || ((255 == 53) && (false && (!true)))) {
                        v29 := GenS1{F0: 4294967295, F1: 397}
                        _ = v29
                    }
                }
                _ = fi26
            }
        }
        v30 := GenS1{F0: 350, F1: (v3 * (v3 & (^(v3))))}
        _ = v30
    }
    v5 = v5
    v31 := GenS1{F0: 42, F1: (v3 >> ((4294967295) & 63))}
    v32 := GenS3{F0: 1}
    fmt.Println(v5, v5)
    v33 := []int{-629, 32768}
    v34 := GenS4{F0: 76, F1: -1, F2: (v3 % ((798) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
