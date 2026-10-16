package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0() uint8 {
    var v4 int8 = 0
    var v5 int = 1
    fmt.Println(v5)
    _ = v4
    _ = v5
    return 127
}
func genFn1(p0 int8) uint32 {
    p0 = (^(p0))
    if true {
        p0 -= (p0 >> ((65535) & 63))
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint = 1
        v8 := []int{(fi6 << (((v7 / ((42) | 1))) & 63)), fi6}
        fmt.Println(v7, p0)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    return 471
}

func main() {
    var v1 uint8 = 104
    v2 := make(map[int]bool)
    var v3 string = ("a" + "a")
    v2[-119] = true
    var v9 int = 127
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            v1 = v1
            v2[32768] = false
            if ((0 != 200) && (125 == v9)) {
                if (-232 < 1) {
                    var v12 uint16 = 32767
                    _ = v12
                }
                v1 = (v1 * v1)
                switch (fi10 | (fi10 >> ((4294967295) & 63))) {
                case -713:
                    v1 += v1
                }
            }
            _ = fi11
        }
        v9 -= fi10
        switch ("hello" + "") {
        case "":
            if (true && false) {
                v3 = (v3 + ("hello" + v3))
                var v13 uint32 = 731
                for fi14 := 0; fi14 < 5; fi14++ {
                    switch "\n\t" {
                    case "":
                        if false {
                            var v15 float32 = (3.14159 / 3.14159)
                            switch fi10 {
                            case -310:
                                if (true && false) {
                                    v16 := GenS1{F0: 510, F1: (^(v13)), F2: 1}
                                    v17 := GenS2{F0: (v1 << ((233) & 63)), F1: v3}
                                    _ = v16
                                    _ = v17
                                } else {
                                    v2[(fi10 * (fi14 / ((v9) | 1)))] = true
                                }
                            default:
                                var v18 int32 = -1
                                _ = v18
                            }
                            fmt.Println(fi14)
                            _ = v15
                        }
                    }
                    v9 = (^(fi10))
                    v19 := GenS3{F0: genFn0()}
                    v9 += fi14
                    if false {
                        v20 := v2[-9223372036854775808]
                        v21 := []int{0}
                        _ = v20
                        _ = v21
                    } else {
                        var v22 uint8 = (v1 % ((v1) | 1))
                        v23 := GenS2{F0: genFn0(), F1: "\n\t"}
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS4{F0: 42, F1: 18446744073709551615}
                    _ = fi14
                    _ = v19
                    _ = v24
                }
                _ = v13
            } else {
                v25 := v2[fi10]
                for fi26 := 0; fi26 < 6; fi26++ {
                    v3 = v3
                    v27 := GenS4{F0: 99, F1: 0}
                    _ = fi26
                    _ = v27
                }
                v9 += (fi10 / (((fi10 / ((128) | 1))) | 1))
                _ = v25
            }
        default:
            var v28 uint16 = 255
            _ = v28
        }
        _ = fi10
    }
    if true {
        v29 := v2[v9]
        delete(v2, (v9 | 4294967295))
        _ = v29
    }
    _ = len(v2)
    for fi30 := 0; fi30 < 3; fi30++ {
        v3 = (v3 + (v3 + v3))
        var v31 uint = 42
        v32 := []int{-983}
        _ = fi30
        _ = v31
        _ = v32
    }
    v33 := []int{(v9 / ((v9) | 1)), (v9 ^ 1), -703, (v9 * v9)}
    for fi34 := 0; fi34 < 10; fi34++ {
        var v35 float32 = -2.25
        v36 := GenS4{F0: 79, F1: 913}
        v37 := GenS4{F0: 1, F1: 799}
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    v9 += -1
    fmt.Println(v3)
    var v38 int32 = -28
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v38:", v38)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v9
}
