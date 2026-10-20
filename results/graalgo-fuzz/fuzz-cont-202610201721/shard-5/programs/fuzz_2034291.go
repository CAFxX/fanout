package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int; F1 uint }
func genFn0(p0 uint16, p1 int) uint8 {
    if false {
        fmt.Println(p0, p1)
        v5 := GenS1{F0: p1, F1: 4294967295, F2: 536}
        _ = v5
    } else {
        switch p1 {
        case -605:
            v6 := []int{p1, 0, (^(p1))}
            _ = v6
        default:
            v7 := GenS1{F0: -207, F1: 963, F2: 42}
            _ = v7
        }
    }
    v8 := []int{p1, (p1 * (p1 ^ 32767)), p1}
    switch "a" {
    case "\n\t":
        var v9 int64 = 9223372036854775807
        _ = v9
    default:
        if (true || ((3.14159 - 1.5) < ((3.14159 / (3.14159 / 0.1)) * (((-2.25 - 3.14159) + 1.5) - 3.14159)))) {
            for fi10 := 0; fi10 < 12; fi10++ {
                if true {
                    p1 += (-(p1))
                    p1 += 128
                } else {
                    fmt.Println(p0, fi10)
                }
                var v11 bool = false
                _ = fi10
                _ = v11
            }
            if (false && true) {
                for fi12 := 0; fi12 < 3; fi12++ {
                    switch p1 {
                    case -9223372036854775808:
                        v13 := GenS2{F0: -1, F1: ("foo bar" + "")}
                        _ = v13
                    default:
                        p1 -= -2147483648
                    }
                    v14 := GenS1{F0: (-(fi12)), F1: 4294967295, F2: 18446744073709551615}
                    fmt.Println(p0, fi12)
                    _ = fi12
                    _ = v14
                }
            } else {
                v15 := GenS3{F0: (p1 | p1), F1: 995}
                if false {
                    p0 = 2
                } else {
                    var v16 int = -182
                    v17 := GenS2{F0: -764, F1: ("foo bar" + ("\n\t" + (("\n\t" + "a") + "foo bar")))}
                    _ = v16
                    _ = v17
                }
                var v18 float32 = 1e10
                _ = v15
                _ = v18
            }
            var v19 string = "a"
            _ = v19
        }
    }
    v20 := GenS3{F0: p1, F1: 65535}
    v21 := GenS2{F0: 0, F1: "\n\t"}
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    _ = v8
    return 152
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 4294967295
    var v3 uint16 = 2
    v4 := []string{("\n\t" + "hello"), ""}
    v3 -= v3
    v2 = 4294967295
    var v22 string = ""
    var v23 uint32 = 996
    switch ((v22 + v22) + "a") {
    case "\n\t":
        for fi24 := 0; fi24 < 2; fi24++ {
            for fi25 := 0; fi25 < 1; fi25++ {
                _ = len(v1)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v26 string = ("a" + ((v22 + "") + (v22 + "a")))
                _ = fi25
                _ = v26
            }
            var v27 uint = 507
            _ = fi24
            _ = v27
        }
    }
    var v28 int64 = 925
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 int16 = 163
        _ = fi29
        _ = v30
    }
    if (!false) {
        _ = len(v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
}
