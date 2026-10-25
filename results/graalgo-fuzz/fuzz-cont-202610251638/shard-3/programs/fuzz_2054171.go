package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint64, p1 uint8, p2 uint64) uint8 {
    p2 -= p0
    v5 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 uint = 65535
    var v6 uint32 = 2
    var v7 uint16 = 145
    if (2147483647 >= 42) {
        if (!(uint8(v4) <= 217)) {
            v2[949] = true
            if (false || true) {
                if false {
                    if false {
                        switch "hello" {
                        case "\n\t":
                            var v8 uint = v4
                            _ = v8
                        }
                        fmt.Println(v6, v7, v6)
                    } else {
                        var v9 float32 = 1.5
                        _ = v9
                    }
                    v10 := GenS1{F0: (v4 & v4)}
                    _ = v10
                }
            } else {
                fmt.Println(v7, v6)
            }
            v11 := GenS1{F0: v4}
            _ = v11
        } else {
            _ = len(v3)
        }
        v12 := GenS1{F0: 42}
        _ = v12
    } else {
        v7 += (^(v7))
    }
    switch 32768 {
    case 127:
        v13 := []int{-9223372036854775808, -2147483648}
        _ = v13
    case -622:
        v4 -= (v4 << ((v4) & 63))
    case 128:
        v7 += (v7 & (v7 << ((v4) & 63)))
    default:
        if (!(847 != 1)) {
            v6 = v6
            switch "hello" {
            case "a":
                v14 := GenS1{F0: (^(v4))}
                _ = v14
            case "hello":
                var v15 uint16 = v7
                _ = v15
            case "":
                fmt.Println(v4, v7)
            }
        } else {
            v16 := []int{128, 899, 600, 32767}
            fmt.Println(v4, v4, v7)
            switch "\n\t" {
            case "hello":
                if ((true && (true || true)) || (10 == 42)) {
                    v6 = v6
                    fmt.Println(v7, v4, v7)
                } else {
                    delete(v1, -1)
                }
            case "a":
                var v17 uint32 = v6
                _ = v17
            default:
                v6 -= v6
            }
            _ = v16
        }
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 int64 = 711
        _ = fi18
        _ = v19
    }
    v7 = (v7 % ((v7) | 1))
    switch ("foo bar" + (("hello" + "hello") + "\n\t")) {
    case "a":
        v4 += (v4 | v4)
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 string = ""
        var v22 int16 = 32767
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint64 = uint64(v7)
        v4 -= (v4 >> (((^(v4))) & 63))
        if (-2147483648 > 42) {
            v7 = v7
        }
        _ = fi23
        _ = v24
    }
    if false {
        if true {
            if true {
                if false {
                    fmt.Println(v7, v4, v6)
                }
            }
            var v25 int = -1
            v4 = (^(v4))
            _ = v25
        } else {
            v26 := []int{}
            _ = v26
        }
        for fi27 := 0; fi27 < 2; fi27++ {
            v28 := []int{fi27, fi27, (fi27 * 127), 4294967295}
            if false {
                var v29 int16 = 42
                _ = v29
            } else {
                switch "" {
                case "\n\t":
                    v7 = (v7 << ((v4) & 63))
                case "a":
                    v6 -= v6
                default:
                    if (42 <= 0) {
                        v6 = uint32(v4)
                        if true {
                            for fi30 := 0; fi30 < 8; fi30++ {
                                v31 := GenS1{F0: v4}
                                _ = fi30
                                _ = v31
                            }
                        }
                    }
                }
            }
            _ = fi27
            _ = v28
        }
    }
    v4 -= v4
    v32 := GenS1{F0: (v4 + 18446744073709551615)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v6
    _ = v7
}
