package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint16 }
type GenS2 struct { F0 int8; F1 int64; F2 int8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int8, p1 uint16) int8 {
    v5 := GenS1{F0: uint(p1), F1: 32767, F2: p1}
    p1 += (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = v5
    return (p0 << ((2) & 63))
}
func genFn1(p0 int16) uint8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(p0, fi6)
        _ = fi6
    }
    if true {
        p0 += p0
    }
    var v7 int = 32767
    switch "hello" {
    case "\n\t":
        if false {
            v8 := GenS1{F0: 2, F1: p0, F2: 2}
            v9 := []int{(v7 & v7), int(p0)}
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 string = "hello"
                var v12 uint16 = 2
                _ = fi10
                _ = v11
                _ = v12
            }
            v13 := GenS2{F0: 14, F1: int64(v7), F2: 42}
            v14 := GenS1{F0: 975, F1: (p0 | (-(p0))), F2: 573}
            _ = v13
            _ = v14
        }
    case "a":
        p0 = (p0 / (((p0 | (p0 & (^(p0))))) | 1))
    default:
        v15 := GenS2{F0: 100, F1: 985, F2: -18}
        _ = v15
    }
    v7 = 127
    _ = p0
    _ = v7
    return 128
}

func main() {
    var v1 int32 = 0
    var v2 string = "\n\t"
    var v3 int = 4294967295
    v4 := make(map[uint32]bool)
    v16 := GenS3{F0: "", F1: 42}
    v17 := GenS4{F0: int32(v3), F1: v1}
    var v18 int16 = -274
    v2 = v2
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 float32 = ((-2.25 / (1.5 - -2.25)) - 0.1)
        _ = fi19
        _ = v20
    }
    switch v3 {
    case -9223372036854775808:
        for fi21 := 0; fi21 < 2; fi21++ {
            if false {
                var v22 int64 = -589
                var v23 uint8 = genFn1(819)
                switch v3 {
                case -9223372036854775808:
                    v24 := GenS5{F0: 10}
                    _ = v24
                case 2147483647:
                    v25 := []int{}
                    _ = v25
                case -1:
                    var v26 uint = 65535
                    _ = v26
                }
                _ = v22
                _ = v23
            } else {
                for fi27 := 0; fi27 < 9; fi27++ {
                    if true {
                        v18 += 1
                        var v28 uint64 = 4
                        switch v2 {
                        case "foo bar":
                            if (true || (v3 == fi27)) {
                                var v29 int32 = (v1 << ((4294967295) & 63))
                                _ = v29
                            }
                        }
                        _ = v28
                    } else {
                        var v30 uint = 42
                        _ = v30
                    }
                    var v31 bool = false
                    fmt.Println(v1, fi21, v2)
                    var v32 int = v3
                    v18 += v18
                    _ = fi27
                    _ = v31
                    _ = v32
                }
                var v33 uint32 = 264
                _ = v33
            }
            v2 = v2
            for fi34 := 0; fi34 < 14; fi34++ {
                if (2 > 2) {
                    var v35 int32 = (v1 - (v1 ^ v1))
                    v36 := []int{}
                    _ = len(v36)
                    _ = v35
                    _ = v36
                }
                _ = fi34
            }
            _ = fi21
        }
    }
    v37 := GenS5{F0: 2}
    if true {
        for fi38 := 0; fi38 < 2; fi38++ {
            var v39 bool = (true && (((v2 + v2) + "") < ((v2 + "foo bar") + v2)))
            v40 := GenS6{F0: 1, F1: 65535}
            _ = fi38
            _ = v39
            _ = v40
        }
    } else {
        var v41 uint16 = 255
        if true {
            fmt.Println(v41)
        }
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
