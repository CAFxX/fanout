package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint32, p1 uint) int {
    v7 := GenS1{F0: "foo bar", F1: (p1 ^ p1), F2: 407}
    p1 = (p1 / ((646) | 1))
    _ = p0
    _ = p1
    _ = v7
    return -984
}

func main() {
    var v1 uint8 = 149
    var v2 int16 = 32767
    v3 := make(map[int]string)
    var v4 uint8 = (v1 & 2)
    var v5 uint = 255
    v6 := make(map[string]uint32)
    var v8 uint32 = 2147483648
    if true {
        v9 := v6[""]
        switch -743 {
        case 4294967295:
            for fi10 := 0; fi10 < 19; fi10++ {
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 uint16 = 306
                    v13 := []int{(fi11 * fi11), -195, (fi11 ^ fi11), 127}
                    var v14 string = ("" + ("hello" + ((("a" + "hello") + "\n\t") + "foo bar")))
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v15 uint16 = 542
                v15 -= (^(v15))
                _ = fi10
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 8; fi16++ {
                if (!true) {
                    v2 = 657
                    var v17 uint32 = v8
                    for fi18 := 0; fi18 < 4; fi18++ {
                        v8 = v9
                        var v19 uint64 = 18446744073709551615
                        v20 := []int{-9223372036854775808, genFn0((v17 + uint32(v19)), 144), (-(fi16))}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    _ = v17
                }
                v9 -= (^(v9))
                _ = fi16
            }
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 uint8 = v4
            if (true || (!(!false))) {
                v23 := GenS1{F0: "a", F1: (v5 << ((v5) & 63)), F2: 42}
                fmt.Println(fi21, v9, v22)
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        _ = v9
    }
    var v24 float64 = (((-2.25 * (3.14159 * (0.1 - ((1e10 - -2.25) + 1.5)))) - ((-2.25 * (3.14159 * (1e10 - -2.25))) + (float64((-2.25 + (0.0 / 1.0))) / 3.14159))) - float64(1.5))
    for fi25 := 0; fi25 < 17; fi25++ {
        switch fi25 {
        case -386:
            if (v4 == 219) {
                v2 -= -32768
                var v26 int32 = 1
                _ = v26
            }
        case -257:
            _ = len(v3)
        default:
            v27 := []int{genFn0(v8, 4294967295), genFn0((^(v8)), (v5 + (v5 | v5))), (fi25 << ((uint(v8)) & 63)), 0}
            _ = v27
        }
        _ = fi25
    }
    if true {
        v4 = (v1 * (v4 ^ 2))
        for fi28 := 0; fi28 < 6; fi28++ {
            v4 = v1
            _ = fi28
        }
    } else {
        delete(v6, "")
    }
    v29 := GenS2{F0: v4, F1: 107}
    v2 = -1
    if true {
        switch "" {
        case "foo bar":
            var v30 int8 = 100
            _ = v30
        case "\n\t":
            v31 := GenS1{F0: ("hello" + "hello"), F1: 116, F2: 1}
            _ = v31
        default:
            for fi32 := 0; fi32 < 8; fi32++ {
                v2 = v2
                _ = fi32
            }
        }
        v2 -= (^(v2))
        v8 = v8
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
