package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
func genFn0() bool {
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{fi6, 2147483647}
        var v8 uint64 = 886
        _ = len(v7)
        _ = fi6
        _ = v7
        _ = v8
    }
    if (127 != 32767) {
        var v9 uint32 = 330
        _ = v9
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            if (!true) {
                fmt.Println(fi10)
                if (1 >= 42) {
                    var v11 int8 = -1
                    _ = v11
                } else {
                    var v12 float64 = (1e10 + 3.14159)
                    v13 := []int{}
                    fmt.Println(v12, v12)
                    _ = v12
                    _ = v13
                }
                var v14 int32 = 194
                _ = v14
            } else {
                v15 := []int{(-(fi10)), fi10, (^(fi10))}
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 uint8 = 144
                    var v18 int16 = -733
                    v17 += v17
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = v15
            }
            if ((42 != 1) || (9223372036854775807 < 1)) {
                fmt.Println(fi10)
                v19 := GenS1{F0: 42, F1: 65}
                _ = v19
            } else {
                v20 := []int{(fi10 << ((18446744073709551615) & 63)), 32767, (fi10 / ((4294967295) | 1))}
                _ = v20
            }
            var v21 bool = false
            _ = fi10
            _ = v21
        }
    }
    v22 := GenS1{F0: 9223372036854775807, F1: -1}
    for fi23 := 0; fi23 < 19; fi23++ {
        if (((32768 <= 2147483647) && false) || true) {
            v24 := GenS1{F0: 9223372036854775807, F1: 55}
            v25 := []int{-1, fi23}
            v25 = append(v25, fi23)
            _ = v24
            _ = v25
        } else {
            var v26 int32 = -948
            var v27 int16 = 1
            _ = v26
            _ = v27
        }
        var v28 int16 = -344
        for fi29 := 0; fi29 < 19; fi29++ {
            v28 += v28
            _ = fi29
        }
        _ = fi23
        _ = v28
    }
    var v30 int8 = 100
    _ = v22
    _ = v30
    return false
}

func main() {
    v1 := []int{-669}
    var v2 float32 = 1e10
    var v3 float32 = (v2 / v2)
    var v4 string = ("\n\t" + "\n\t")
    var v5 float64 = (0.0 + 1e10)
    v31 := GenS1{F0: -1, F1: 909}
    for fi32 := 0; fi32 < 5; fi32++ {
        fmt.Println(fi32, v3)
        _ = fi32
    }
    var v33 uint32 = 248
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 1 {
    case 668:
        for fi34 := 0; fi34 < 13; fi34++ {
            for fi35 := 0; fi35 < 3; fi35++ {
                v1 = append(v1, fi34)
                var v36 int = (fi34 >> ((0) & 63))
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
    case 2147483647:
        v5 -= 1.5
    case 32768:
        v5 -= ((float64(v2) - v5) / 0.1)
    }
    fmt.Println(v4)
    v5 += v5
    v2 -= (v2 / (v2 * v2))
    switch -759 {
    case 42:
        var v37 float32 = v3
        _ = v37
    case -2147483648:
        if false {
            var v38 string = ""
            for fi39 := 0; fi39 < 12; fi39++ {
                if (!((!(1 == 65535)) || ((false && (true && true)) || true))) {
                    v4 = v4
                    var v40 int8 = -85
                    v40 = v40
                    _ = v40
                }
                v5 += (0.0 - v5)
                _ = fi39
            }
            _ = v38
        } else {
            var v41 int8 = 127
            _ = v41
        }
    case 32768:
        for fi42 := 0; fi42 < 7; fi42++ {
            v2 -= v2
            fmt.Println(v3)
            switch ("foo bar" + v4) {
            case "":
                var v43 int = fi42
                _ = v43
            }
            _ = fi42
        }
    }
    if (false && (894 != (v33 - (v33 / (((v33 << ((727) & 63))) | 1))))) {
        if false {
            v44 := GenS1{F0: 9223372036854775807, F1: -464}
            _ = v44
        }
        v4 = ((((v4 + "a") + "\n\t") + v4) + v4)
        v45 := GenS1{F0: -1, F1: 998}
        _ = v45
    }
    v33 = (v33 + (v33 + (v33 - 42)))
    fmt.Println(v4, v3)
    for fi46 := 0; fi46 < 8; fi46++ {
        var v47 uint = 2
        if true {
            if len(v1) > 0 {
                v1[0] = int(v33)
            }
            v3 -= (0.1 - (v2 / (3.14159 + (v3 / ((v3 + v2) - (3.14159 * v2))))))
        } else {
            for fi48 := 0; fi48 < 3; fi48++ {
                var v49 float64 = 0.1
                v1 = append(v1, (fi46 - fi48))
                _ = fi48
                _ = v49
            }
        }
        _ = fi46
        _ = v47
    }
    v33 -= (v33 % ((2147483648) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v5
}
