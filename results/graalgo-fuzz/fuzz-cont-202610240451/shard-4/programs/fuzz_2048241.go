package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 int16, p1 string) uint16 {
    v6 := []int{-9223372036854775808, 42}
    v7 := GenS1{F0: 195, F1: 1, F2: 86}
    v8 := GenS1{F0: 255, F1: -2147483648, F2: 2}
    if true {
        v9 := GenS1{F0: 297, F1: -391, F2: 138}
        p0 = (p0 % ((p0) | 1))
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 4294967295
    v3 := make(map[string]bool)
    var v4 string = "\n\t"
    v5 := make(map[uint32]bool)
    var v10 uint8 = 0
    var v11 int64 = 540
    if false {
        switch -837 {
        case 798:
            v12 := GenS1{F0: 255, F1: -2147483648, F2: 209}
            _ = v12
        case 32767:
            if false {
                var v13 int = 32767
                delete(v5, 194)
                _ = v13
            } else {
                v4 = ("hello" + v4)
                v11 = v11
                v2 = (v2 << ((18446744073709551615) & 63))
            }
        case -900:
            v4 = v4
        }
        if ((0 > 540) && true) {
            if false {
                v2 -= 1
                for fi14 := 0; fi14 < 5; fi14++ {
                    v15 := []int{}
                    if (((v4 + "\n\t") + v4) <= (("\n\t" + "a") + v4)) {
                        v16 := v5[649]
                        _ = v16
                    }
                    fmt.Println(v11)
                    _ = fi14
                    _ = v15
                }
            } else {
                v4 = v4
                for fi17 := 0; fi17 < 4; fi17++ {
                    var v18 float32 = 1e10
                    _ = fi17
                    _ = v18
                }
                var v19 int = 2147483647
                _ = v19
            }
            if true {
                v11 = v11
            }
            var v20 uint64 = 18446744073709551615
            _ = v20
        } else {
            if (!((2147483647 <= 0) && false)) {
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 bool = (true && true)
                    _ = fi21
                    _ = v22
                }
                if false {
                    v23 := GenS1{F0: 2, F1: 42, F2: v10}
                    switch 2147483647 {
                    case 651:
                        _ = len(v3)
                    case 1:
                        fmt.Println(v10)
                    default:
                        v4 = v4
                    }
                    _ = v23
                } else {
                    if (v11 <= -452) {
                        var v24 uint32 = 0
                        if true {
                            v25 := GenS1{F0: 312, F1: 74, F2: v10}
                            _ = v25
                        }
                        v26 := []int{}
                        _ = v24
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 12; fi27++ {
                            var v28 int64 = (v11 << ((43) & 63))
                            if ((v2 | (v2 + (v2 - v2))) > (^(v2))) {
                                for fi29 := 0; fi29 < 14; fi29++ {
                                    switch fi29 {
                                    case 32:
                                        var v30 int64 = (v28 << ((255) & 63))
                                        _ = v30
                                    }
                                    v2 -= (v2 * v2)
                                    var v31 int32 = 0
                                    _ = fi29
                                    _ = v31
                                }
                                v10 += v10
                            }
                            var v32 float64 = (3.14159 + 1e10)
                            _ = fi27
                            _ = v28
                            _ = v32
                        }
                        var v33 int8 = 42
                        v34 := GenS1{F0: 65535, F1: 281, F2: v10}
                        _ = v33
                        _ = v34
                    }
                    v3[(v4 + v4)] = false
                    v4 = "a"
                }
            }
            if (v10 != v10) {
                v35 := GenS2{F0: false, F1: 907}
                _ = v35
            }
        }
        fmt.Println(v10, v11)
    }
    switch 0 {
    case 919:
        fmt.Println(v4)
    case 4294967295:
        v36 := GenS1{F0: 0, F1: 42, F2: v10}
        _ = v36
    case -333:
        if (v10 > (v10 | v10)) {
            v37 := v3[""]
            v38 := GenS1{F0: 2, F1: -2147483648, F2: (v10 % (((v10 << ((137) & 63))) | 1))}
            _ = v37
            _ = v38
        }
    default:
        var v39 uint16 = genFn0(1000, v4)
        _ = v39
    }
    v40 := GenS2{F0: false, F1: (v2 ^ 1)}
    fmt.Println(v10, v4, v2)
    v1[("foo bar" + v4)] = 4294967295
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
