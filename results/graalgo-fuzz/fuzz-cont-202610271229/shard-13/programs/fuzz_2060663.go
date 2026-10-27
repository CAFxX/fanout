package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int64 }

func main() {
    var v1 int16 = 42
    var v2 uint32 = 2147483647
    var v3 string = ""
    v4 := make(map[int]int)
    var v5 uint8 = 2
    fmt.Println(v5, v3)
    v6 := GenS1{F0: v1, F1: 769, F2: 0}
    switch v3 {
    case "hello":
        if true {
            var v7 float32 = -2.25
            _ = v7
        } else {
            if true {
                fmt.Println(v1, v2)
            } else {
                var v8 int32 = 16
                _ = v8
            }
        }
    default:
        var v9 uint8 = v5
        _ = v9
    }
    var v10 string = v3
    v5 = (v5 & (v5 << ((894) & 63)))
    for fi11 := 0; fi11 < 3; fi11++ {
        if (!true) {
            for fi12 := 0; fi12 < 9; fi12++ {
                v10 = (v3 + (v10 + v10))
                _ = fi12
            }
            switch "a" {
            case "a":
                if (false && (-9223372036854775808 != -1)) {
                    v13 := GenS1{F0: v1, F1: 2147483647, F2: 1}
                    v2 -= (v2 / ((v2) | 1))
                    v3 = (v3 + "")
                    _ = v13
                }
            case "":
                for fi14 := 0; fi14 < 7; fi14++ {
                    v10 = v3
                    var v15 uint16 = 2
                    v16 := []int{(fi14 | (fi14 & 0)), (fi11 - fi14), (fi11 - (fi11 ^ 912)), fi11}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "\n\t":
                for fi17 := 0; fi17 < 8; fi17++ {
                    fmt.Println(fi11, fi17)
                    switch v10 {
                    case "foo bar":
                        var v18 uint8 = v5
                        _ = v18
                    default:
                        var v19 uint32 = v2
                        _ = v19
                    }
                    v20 := GenS2{F0: 32767, F1: -1}
                    _ = fi17
                    _ = v20
                }
            default:
                v2 = (v2 + (^(v2)))
            }
            if (-128 != -51) {
                fmt.Println(fi11)
                if (v10 == ((v10 + "foo bar") + "")) {
                    v1 -= (v1 - (v1 ^ v1))
                } else {
                    var v21 int = (fi11 ^ (fi11 << ((200) & 63)))
                    _ = v21
                }
            }
        }
        v10 = ("hello" + v10)
        _ = fi11
    }
    delete(v4, 32767)
    switch v3 {
    case "":
        v22 := GenS2{F0: 42, F1: 41}
        _ = v22
    case "foo bar":
        var v23 int16 = v1
        _ = v23
    default:
        v3 = v3
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 uint = 0
        v26 := []int{32768, fi24, (-(fi24))}
        if (-1 >= -9223372036854775808) {
            if false {
                v27 := v4[fi24]
                _ = v27
            } else {
                var v28 float32 = -2.25
                _ = v28
            }
            fmt.Println(v5)
            if (0 != v5) {
                _ = len(v26)
                _ = len(v4)
                if len(v26) > 0 {
                    v26[0] = (fi24 | 2147483647)
                }
            } else {
                switch v10 {
                case "foo bar":
                    if (false || (!(v5 == (v5 ^ 128)))) {
                        var v29 int64 = -1
                        v3 = "a"
                        _ = v29
                    } else {
                        v2 += uint32(v5)
                        var v30 uint16 = 1
                        v10 = v3
                        _ = v30
                    }
                case "a":
                    var v31 uint32 = v2
                    _ = v31
                case "\n\t":
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                }
                var v32 uint = (v25 + (v25 ^ 255))
                _ = v32
            }
        } else {
            for fi33 := 0; fi33 < 13; fi33++ {
                var v34 uint8 = (v5 | (^(v5)))
                v10 = v3
                var v35 uint32 = v2
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    if false {
        fmt.Println(v3)
    }
    v5 -= 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
