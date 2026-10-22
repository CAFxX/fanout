package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int32, p1 int, p2 int16) uint {
    fmt.Println(p0, p2, p2)
    var v6 int64 = 1
    if true {
        if (!(((-1 <= (v6 ^ (v6 + v6))) && true) || true)) {
            var v7 float32 = 3.14159
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 uint8 = 127
                p2 += (p2 % ((-86) | 1))
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            v10 := GenS1{F0: (p2 / ((p2) | 1))}
            for fi11 := 0; fi11 < 1; fi11++ {
                switch "\n\t" {
                case "\n\t":
                    v12 := []int{(fi11 % ((201) | 1)), 32768, (fi11 + (^(fi11)))}
                    _ = v12
                }
                _ = fi11
            }
            _ = v10
        }
    }
    if true {
        var v13 float32 = 1e10
        _ = v13
    } else {
        p0 -= (p0 * p0)
        v14 := []int{p1, p1}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 663
}

func main() {
    v1 := []string{("" + "\n\t")}
    v2 := make(map[int]string)
    v3 := []uint32{240, 2147483647, 2}
    var v4 uint8 = 1
    var v5 int16 = 0
    var v15 bool = (false && true)
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := GenS2{F0: (v4 << ((1) & 63)), F1: 42}
        _ = fi16
        _ = v17
    }
    var v18 int = 9223372036854775807
    if (v15 && false) {
        var v19 float32 = 1e10
        if (0 <= 1) {
            var v20 int32 = -321
            for fi21 := 0; fi21 < 7; fi21++ {
                v22 := GenS1{F0: v5}
                for fi23 := 0; fi23 < 3; fi23++ {
                    switch 0 {
                    case 810:
                        v19 = (v19 * 1e10)
                    default:
                        v20 = 42
                        fmt.Println(v5)
                        v24 := GenS1{F0: v5}
                        _ = v24
                    }
                    _ = fi23
                }
                _ = fi21
                _ = v22
            }
            var v25 int8 = 26
            _ = v20
            _ = v25
        } else {
            if len(v1) > 0 {
                v1[0] = "a"
            }
        }
        delete(v2, v18)
        _ = v19
    }
    var v26 int64 = 0
    for fi27 := 0; fi27 < 3; fi27++ {
        fmt.Println(v15, v26, v15)
        var v28 int16 = (v5 * (v5 >> ((65535) & 63)))
        for fi29 := 0; fi29 < 13; fi29++ {
            v3 = append(v3, 2147483647)
            fmt.Println(fi29, v15)
            v28 = -1
            v15 = false
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    var v30 bool = (!v15)
    fmt.Println(v15, v5, v4)
    v31 := GenS1{F0: v5}
    v32 := GenS2{F0: v4, F1: 772}
    fmt.Println(v26, v5)
    if v30 {
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 int8 = 1
            v35 := GenS1{F0: (v5 / (((v5 * (v5 / (((^(v5))) | 1)))) | 1))}
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    v26 += (v26 & (v26 << ((8) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
