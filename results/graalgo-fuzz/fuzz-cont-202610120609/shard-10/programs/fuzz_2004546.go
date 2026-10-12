package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint; F2 int64 }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint16, p1 uint16, p2 uint16) int8 {
    switch (("foo bar" + "\n\t") + (("" + ("a" + "")) + "a")) {
    case "hello":
        var v5 int16 = 32767
        _ = v5
    default:
        if true {
            var v6 int16 = 64
            v7 := GenS1{F0: 65535, F1: v6, F2: 0}
            _ = v6
            _ = v7
        } else {
            var v8 string = "hello"
            var v9 int = 128
            _ = v8
            _ = v9
        }
    }
    v10 := GenS2{F0: (p2 & (p1 / ((p0) | 1))), F1: 42, F2: -730}
    if false {
        for fi11 := 0; fi11 < 2; fi11++ {
            if true {
                if false {
                    fmt.Println(p0)
                    if ((311 >= 0) && false) {
                        switch -950 {
                        case 575:
                            for fi12 := 0; fi12 < 9; fi12++ {
                                fmt.Println(p1, p0, fi12)
                                fmt.Println(p0, p2)
                                _ = fi12
                            }
                        case 42:
                            var v13 int = (fi11 << ((420) & 63))
                            _ = v13
                        }
                        if ((1e10 + 3.14159) <= -2.25) {
                            v14 := []int{(fi11 & 563)}
                            switch fi11 {
                            case 705:
                                if (!false) {
                                    var v15 uint8 = 1
                                    _ = v15
                                }
                            default:
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            }
                            _ = v14
                        } else {
                            p2 -= p0
                        }
                        fmt.Println(p1)
                    }
                } else {
                    fmt.Println(p1, p1, p2)
                }
            } else {
                for fi16 := 0; fi16 < 13; fi16++ {
                    v17 := GenS2{F0: 32767, F1: 255, F2: 9223372036854775807}
                    switch "a" {
                    case "foo bar":
                        switch fi16 {
                        case -9223372036854775808:
                            var v18 int64 = 9223372036854775807
                            _ = v18
                        case 32768:
                            fmt.Println(p1, fi16)
                        case -2147483648:
                            fmt.Println(fi11)
                        default:
                            var v19 float64 = -2.25
                            _ = v19
                        }
                    default:
                        var v20 uint = uint(p1)
                        _ = v20
                    }
                    p0 = (^(p2))
                    p1 -= (p2 % (((p2 / ((p1) | 1))) | 1))
                    fmt.Println(p1)
                    switch (("a" + "\n\t") + "") {
                    case "\n\t":
                        v21 := []int{(fi11 - fi11)}
                        _ = v21
                    default:
                        v22 := []int{fi11}
                        _ = v22
                    }
                    p0 += 65535
                    _ = fi16
                    _ = v17
                }
                p0 -= (p2 % ((32767) | 1))
            }
            var v23 float64 = 3.14159
            _ = fi11
            _ = v23
        }
        v24 := GenS3{F0: 2147483647, F1: 854}
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 100
}
func genFn1(p0 uint32, p1 uint64, p2 uint16) bool {
    fmt.Println(p0, p2)
    for fi25 := 0; fi25 < 18; fi25++ {
        fmt.Println(p1)
        _ = fi25
    }
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    var v1 int = -2147483648
    var v2 int64 = 9223372036854775807
    var v3 int8 = -113
    v4 := make(map[uint32]int)
    fmt.Println(v1, v2, v3)
    delete(v4, 317)
    var v26 int16 = 1
    v3 = (v3 & 100)
    var v27 int64 = v2
    v3 = (-(v3))
    var v28 uint = 136
    fmt.Println(v1)
    for fi29 := 0; fi29 < 9; fi29++ {
        v30 := GenS2{F0: 27, F1: v28, F2: 9223372036854775807}
        v31 := GenS3{F0: 0, F1: 118}
        v32 := []int{9223372036854775807, (^(v1))}
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 uint32 = 4294967295
        var v35 bool = (true || false)
        v1 -= (fi33 | v1)
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println(v27, v2)
    fmt.Println(v2, v28)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
