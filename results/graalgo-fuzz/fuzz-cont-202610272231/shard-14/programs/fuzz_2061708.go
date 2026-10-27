package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 bool; F2 int32 }
type GenS3 struct { F0 uint; F1 int8; F2 uint }
func genFn0(p0 int16, p1 uint16, p2 uint16) uint64 {
    var v4 int = 32768
    v4 = (v4 | 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}

func main() {
    var v1 int16 = -808
    var v2 int32 = -1
    v3 := make(map[uint32]bool)
    if false {
        v2 = (v2 % (((v2 * v2)) | 1))
        for fi5 := 0; fi5 < 3; fi5++ {
            v6 := []int{-882, fi5, fi5, (^(fi5))}
            _ = fi5
            _ = v6
        }
        if false {
            v7 := []int{4294967295, -2147483648}
            v2 += v2
            fmt.Println(v2)
            _ = v7
        }
    } else {
        v8 := GenS1{F0: 162}
        for fi9 := 0; fi9 < 16; fi9++ {
            if (18446744073709551615 < 42) {
                for fi10 := 0; fi10 < 3; fi10++ {
                    v11 := GenS2{F0: 42, F1: (1.5 >= 1.5), F2: -2147483648}
                    _ = fi10
                    _ = v11
                }
                v1 -= (v1 ^ (v1 * 649))
                v12 := GenS1{F0: 936}
                _ = v12
            } else {
                var v13 string = ("a" + "hello")
                _ = v13
            }
            var v14 uint64 = 2
            _ = fi9
            _ = v14
        }
        fmt.Println(v2, v1)
        _ = v8
    }
    if true {
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 uint8 = 39
                v2 += -2147483648
                _ = fi16
                _ = v17
            }
            if false {
                for fi18 := 0; fi18 < 19; fi18++ {
                    var v19 int32 = v2
                    var v20 int32 = v19
                    v21 := GenS3{F0: 773, F1: -128, F2: 65535}
                    v19 -= (v20 - v19)
                    if true {
                        v19 += (v2 + v20)
                        v19 += 0
                        v1 += 441
                        var v22 uint32 = 901
                        _ = v22
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                v1 += (v1 << ((42) & 63))
                for fi23 := 0; fi23 < 19; fi23++ {
                    if false {
                        v24 := GenS2{F0: 32767, F1: (826 != 32767), F2: 42}
                        var v25 int = fi15
                        switch ("hello" + "\n\t") {
                        case "":
                            var v26 uint64 = 4294967295
                            _ = v26
                        case "a":
                            fmt.Println(v25)
                        default:
                            var v27 uint32 = 0
                            _ = v27
                        }
                        _ = v24
                        _ = v25
                    } else {
                        if true {
                            v2 = v2
                        } else {
                            if true {
                                switch (fi15 - fi15) {
                                case 511:
                                    v2 = 940
                                case 1:
                                    var v28 string = ("a" + "a")
                                    _ = v28
                                }
                                v2 -= 0
                                v1 -= v1
                            } else {
                                fmt.Println(fi23, fi23, v2)
                                if (2147483647 <= 42) {
                                    var v29 uint32 = 704
                                    _ = v29
                                }
                            }
                            var v30 bool = (65535 >= 926)
                            _ = v30
                        }
                    }
                    v1 += (v1 * (-(v1)))
                    v1 = -250
                    _ = fi23
                }
            }
            _ = fi15
        }
    }
    var v31 int8 = 0
    v32 := v3[2147483648]
    v32 = (!(((1e10 + 1.5) / 3.14159) != (float32((-2.25 * 3.14159)) / 1.0)))
    var v33 uint = 42
    v32 = false
    v34 := GenS3{F0: 708, F1: -1, F2: 1}
    v35 := []int{-1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v35:", len(v35), v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
