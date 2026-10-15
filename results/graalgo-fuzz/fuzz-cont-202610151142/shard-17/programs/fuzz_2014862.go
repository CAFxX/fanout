package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 bool }
func genFn0(p0 int) uint8 {
    v6 := GenS1{F0: -55, F1: 1, F2: 2147483648}
    var v7 uint32 = 570
    if true {
        if true {
            for fi8 := 0; fi8 < 19; fi8++ {
                fmt.Println(v7, fi8, fi8)
                var v9 int32 = 284
                if false {
                    var v10 uint8 = 222
                    var v11 uint = 427
                    _ = v10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            var v12 int8 = -1
            _ = v12
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 bool, p1 uint64) int16 {
    v13 := []int{42, 544}
    p1 = p1
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 float32 = (0.0 * (3.14159 * ((1e10 + 0.1) + (1e10 * 0.1))))
        switch fi14 {
        case 32768:
            _ = len(v13)
        case 883:
            v16 := GenS1{F0: 100, F1: 0, F2: 42}
            _ = v16
        case -1:
            var v17 uint32 = 2147483647
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v13
    return 1
}

func main() {
    var v1 uint8 = 56
    var v2 float64 = -2.25
    var v3 uint = 42
    var v4 uint64 = 1
    var v5 string = (((("\n\t" + (("foo bar" + "hello") + "foo bar")) + (("\n\t" + "foo bar") + "foo bar")) + ("\n\t" + "\n\t")) + "")
    v2 += v2
    if ((v1 % ((v1) | 1)) == v1) {
        var v18 bool = true
        v18 = v18
        _ = v18
    }
    var v19 uint32 = 65535
    if false {
        if (!false) {
            if true {
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 bool = (!false)
                    v22 := GenS1{F0: 127, F1: 616, F2: (v19 - (^(v19)))}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                var v23 uint64 = (v4 | v4)
                if false {
                    v24 := GenS1{F0: -1, F1: -1, F2: v19}
                    v25 := GenS2{F0: 32768, F1: (v19 % ((1) | 1))}
                    var v26 bool = false
                    _ = v24
                    _ = v25
                    _ = v26
                }
                for fi27 := 0; fi27 < 1; fi27++ {
                    var v28 uint8 = (v1 - (^(v1)))
                    if (((genFn0(0) == 1) || false) || false) {
                        for fi29 := 0; fi29 < 19; fi29++ {
                            v30 := GenS3{F0: 255, F1: v1, F2: (^(v1))}
                            if (2 >= 779) {
                                var v31 bool = true
                                v31 = (true || v31)
                                _ = v31
                            } else {
                                v19 -= v19
                                v32 := []int{}
                                fmt.Println(v2)
                                _ = v32
                            }
                            var v33 uint16 = 500
                            _ = fi29
                            _ = v30
                            _ = v33
                        }
                        var v34 string = "\n\t"
                        _ = v34
                    } else {
                        v35 := GenS4{F0: v1, F1: 65535}
                        _ = v35
                    }
                    var v36 int = fi27
                    _ = fi27
                    _ = v28
                    _ = v36
                }
                _ = v23
            }
        }
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 uint8 = v1
            fmt.Println(v38, v4)
            _ = fi37
            _ = v38
        }
    }
    if true {
        v5 = "a"
        v3 -= (v3 ^ uint(v1))
    } else {
        v39 := GenS5{F0: true}
        v40 := GenS4{F0: v1, F1: v3}
        _ = v39
        _ = v40
    }
    fmt.Println(v2, v5)
    v41 := []int{}
    var v42 uint32 = v19
    v4 -= (v4 * v4)
    fmt.Println(v42)
    if len(v41) >= 2 {
        v41 = v41[1:len(v41)]
    }
    var v43 float64 = v2
    fmt.Println(v19, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
