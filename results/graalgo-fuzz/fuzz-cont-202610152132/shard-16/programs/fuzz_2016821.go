package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0(p0 int16, p1 int) int64 {
    v6 := []int{(p1 | 42), -9223372036854775808}
    fmt.Println(p1)
    var v7 bool = ("a" < "a")
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := []int{128, -318, -2147483648, -280}
    v2 := []int{615, -1}
    var v3 uint = 65535
    var v4 uint16 = 255
    var v5 uint = v3
    var v8 uint8 = 64
    for fi9 := 0; fi9 < 8; fi9++ {
        v5 -= 42
        if ((true && false) && (false && false)) {
            v3 = 986
            if (3.14159 != -2.25) {
                v10 := GenS1{F0: genFn0(-32768, fi9), F1: v4}
                var v11 uint32 = 2
                _ = v10
                _ = v11
            }
            if true {
                v5 -= v3
                var v12 uint16 = v4
                _ = v12
            }
        } else {
            var v13 int16 = 42
            var v14 uint = v3
            _ = v13
            _ = v14
        }
        var v15 string = "\n\t"
        _ = fi9
        _ = v15
    }
    if true {
        var v16 uint32 = 816
        fmt.Println(v16, v16, v16)
        fmt.Println(v16, v3, v3)
        _ = v16
    } else {
        for fi17 := 0; fi17 < 3; fi17++ {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 float64 = -2.25
                var v20 uint32 = 2
                _ = fi18
                _ = v19
                _ = v20
            }
            v8 += v8
            _ = fi17
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS1{F0: -1, F1: v4}
            var v23 uint16 = (^(v4))
            v24 := GenS1{F0: 0, F1: 65535}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        v8 += (^(v8))
    }
    v1 = append(v1, 32767)
    switch -799 {
    case 77:
        var v25 uint8 = uint8(v5)
        _ = v25
    case 9223372036854775807:
        var v26 uint32 = 978
        _ = v26
    default:
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 bool = (18446744073709551615 <= 4294967295)
            var v29 uint32 = 2
            v5 = v3
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    v3 -= (v3 | v5)
    var v30 uint64 = uint64(v8)
    for fi31 := 0; fi31 < 4; fi31++ {
        fmt.Println(v8)
        _ = fi31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        v2 = append(v2, -823)
        if (!true) {
            if true {
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
                v30 += v30
            }
        } else {
            v30 += v30
            var v33 int16 = 42
            _ = v33
        }
        if len(v2) > 0 {
            v2[0] = (fi32 & fi32)
        }
        _ = fi32
    }
    var v34 uint8 = v8
    var v35 uint64 = uint64(v4)
    if (!false) {
        var v36 uint = v3
        for fi37 := 0; fi37 < 12; fi37++ {
            fmt.Println(v5, v35)
            _ = fi37
        }
        _ = v36
    } else {
        for fi38 := 0; fi38 < 11; fi38++ {
            for fi39 := 0; fi39 < 3; fi39++ {
                var v40 int64 = 42
                _ = fi39
                _ = v40
            }
            if len(v1) > 0 {
                v1[0] = fi38
            }
            _ = fi38
        }
        v41 := GenS1{F0: genFn0(1000, 577), F1: (v4 >> ((v5) & 63))}
        var v42 uint = (v3 % ((v3) | 1))
        _ = v41
        _ = v42
    }
    fmt.Println(v35, v30)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v8
}
