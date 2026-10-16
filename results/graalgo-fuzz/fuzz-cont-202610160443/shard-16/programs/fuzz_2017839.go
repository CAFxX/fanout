package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 string; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 42
    var v3 float32 = (1e10 / 1e10)
    v4 := make(map[uint32]string)
    v5 := make(map[int]int)
    v2 = v2
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int64 = -1
        fmt.Println(v2, v1, v3)
        fmt.Println(v7)
        _ = fi6
        _ = v7
    }
    if false {
        fmt.Println(v3)
        v3 += v3
    }
    fmt.Println(v2, v1)
    var v8 string = "a"
    v1 += v1
    switch 402 {
    case -498:
        v8 = v8
    default:
        v9 := GenS1{F0: 2147483647, F1: 187}
        _ = v9
    }
    if true {
        switch 128 {
        case -2147483648:
            if true {
                v2 -= (v2 << ((716) & 63))
                for fi10 := 0; fi10 < 1; fi10++ {
                    v11 := []int{(fi10 ^ fi10), (fi10 ^ (fi10 * fi10))}
                    var v12 bool = (!true)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 int8 = -1
                    var v15 bool = true
                    if (fi13 != int(v1)) {
                        var v16 uint = 4294967295
                        v16 -= (v16 + (v16 ^ v16))
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            } else {
                v2 = v2
            }
        default:
            v2 += v2
        }
    } else {
        v2 -= (v2 & v2)
    }
    var v17 int32 = 719
    if false {
        v2 = v2
        if false {
            for fi18 := 0; fi18 < 16; fi18++ {
                v19 := []int{fi18}
                _ = fi18
                _ = v19
            }
            v1 = (v1 + v1)
            v1 = (v1 | v1)
        } else {
            fmt.Println(v8)
            v3 += (1.5 * 0.1)
        }
        if (0 <= 46) {
            if ((18446744073709551615 == 1) || true) {
                if (!(v17 != v17)) {
                    fmt.Println(v3, v1, v8)
                } else {
                    var v20 float64 = ((0.1 + float64(v3)) + 0.0)
                    v21 := GenS2{F0: 839, F1: (v8 + v8), F2: 42}
                    _ = v20
                    _ = v21
                }
            } else {
                for fi22 := 0; fi22 < 6; fi22++ {
                    if (1e10 >= 0.0) {
                        v17 = (v17 * (v17 & 1))
                    }
                    var v23 uint16 = v2
                    v24 := GenS2{F0: 0, F1: (v8 + v8), F2: 4294967295}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v8 = ((("a" + v8) + "a") + v8)
                _ = len(v4)
            }
        } else {
            for fi25 := 0; fi25 < 4; fi25++ {
                v26 := GenS1{F0: -682, F1: 42}
                _ = fi25
                _ = v26
            }
            if true {
                for fi27 := 0; fi27 < 5; fi27++ {
                    fmt.Println(v2, fi27, fi27)
                    var v28 uint8 = 242
                    _ = fi27
                    _ = v28
                }
                if false {
                    v29 := GenS1{F0: -2147483648, F1: 2}
                    _ = v29
                } else {
                    v30 := []int{451}
                    for fi31 := 0; fi31 < 18; fi31++ {
                        fmt.Println(v8, v2, v3)
                        _ = fi31
                    }
                    _ = v30
                }
                if (-2.25 <= v3) {
                    if false {
                        v8 = ("\n\t" + "a")
                        var v32 int8 = 42
                        v17 -= v17
                        _ = v32
                    }
                } else {
                    v33 := GenS3{F0: 402, F1: 127}
                    for fi34 := 0; fi34 < 4; fi34++ {
                        if false {
                            v35 := []int{(fi34 * (fi34 | (fi34 << ((1) & 63)))), fi34, 9223372036854775807, fi34}
                            _ = v35
                        }
                        for fi36 := 0; fi36 < 16; fi36++ {
                            fmt.Println(v8, fi34)
                            var v37 int16 = (^(v1))
                            _ = fi36
                            _ = v37
                        }
                        _ = fi34
                    }
                    _ = v33
                }
            }
            v38 := []int{-644, 32767, 2147483647, -897}
            _ = v38
        }
    } else {
        for fi39 := 0; fi39 < 8; fi39++ {
            v40 := v5[2147483647]
            fmt.Println(v8, v1)
            _ = fi39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
