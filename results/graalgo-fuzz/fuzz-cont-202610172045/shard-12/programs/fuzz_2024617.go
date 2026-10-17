package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int }
type GenS2 struct { F0 uint32; F1 uint }
func genFn0(p0 uint64) int {
    for fi6 := 0; fi6 < 7; fi6++ {
        p0 = p0
        if true {
            p0 = (^(p0))
            for fi7 := 0; fi7 < 2; fi7++ {
                p0 -= p0
                var v8 int16 = -32768
                if (-128 < 25) {
                    switch fi7 {
                    case 42:
                        p0 = 202
                    default:
                        p0 -= (p0 / ((p0) | 1))
                    }
                }
                _ = fi7
                _ = v8
            }
            var v9 float64 = 3.14159
            _ = v9
        } else {
            switch "" {
            case "foo bar":
                var v10 uint32 = 65535
                _ = v10
            case "a":
                v11 := GenS1{F0: 0, F1: -128, F2: fi6}
                _ = v11
            }
            p0 = p0
            for fi12 := 0; fi12 < 2; fi12++ {
                p0 -= p0
                v13 := GenS1{F0: 42, F1: 1, F2: int(p0)}
                p0 += 278
                _ = fi12
                _ = v13
            }
        }
        fmt.Println(p0, p0)
        _ = fi6
    }
    p0 = p0
    switch 9223372036854775807 {
    case 4294967295:
        p0 += (p0 - (p0 % (((p0 / ((p0) | 1))) | 1)))
    case -9223372036854775808:
        p0 += (p0 & p0)
    }
    _ = p0
    return 42
}

func main() {
    var v1 string = ("a" + ("" + ("hello" + "foo bar")))
    var v2 string = v1
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    v5 := make(map[int]string)
    for fi14 := 0; fi14 < 11; fi14++ {
        v1 = v1
        var v15 uint16 = 65535
        v16 := GenS2{F0: 866, F1: 42}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint64 = 42
    for fi18 := 0; fi18 < 16; fi18++ {
        if false {
            if (0.1 > 1e10) {
                _ = len(v3)
            } else {
                for fi19 := 0; fi19 < 5; fi19++ {
                    v20 := GenS1{F0: -2147483648, F1: 0, F2: (fi18 / ((fi18) | 1))}
                    if true {
                        delete(v4, 65535)
                    }
                    _ = fi19
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 5; fi21++ {
                if false {
                    var v22 uint8 = 128
                    v23 := []int{}
                    if false {
                        var v24 uint8 = (v22 >> ((42) & 63))
                        if (-500 > 42) {
                            v22 = (^(v24))
                            v25 := v5[fi21]
                            v22 = 9
                            _ = v25
                        }
                        _ = v24
                    }
                    _ = v22
                    _ = v23
                }
                _ = fi21
            }
        } else {
            v17 -= v17
            switch 127 {
            case 0:
                switch "hello" {
                case "\n\t":
                    var v26 float32 = (1e10 - 0.0)
                    _ = v26
                case "foo bar":
                    fmt.Println(v1, v2)
                default:
                    delete(v3, 42)
                }
            case 42:
                var v27 int8 = 100
                _ = v27
            case 128:
                v17 = v17
            }
            for fi28 := 0; fi28 < 10; fi28++ {
                var v29 string = ("hello" + v2)
                v30 := GenS2{F0: 65535, F1: 4294967295}
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        _ = fi18
    }
    var v31 int64 = 42
    switch 127 {
    case -1:
        if true {
            v31 = v31
            v2 = (v2 + v2)
        } else {
            v17 -= (v17 & 4294967295)
        }
    case 2147483647:
        var v32 uint = 795
        _ = v32
    default:
        var v33 uint = 65535
        _ = v33
    }
    v2 = ("foo bar" + v1)
    if (284 == 4294967295) {
        fmt.Println(v17)
        v31 = (-(v31))
        delete(v5, -1)
    } else {
        fmt.Println(v31)
    }
    v17 = v17
    var v34 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
