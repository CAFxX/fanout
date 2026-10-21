package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 int64 }
func genFn0(p0 int64, p1 uint32) int16 {
    p0 = (p0 * (p0 << ((0) & 63)))
    v6 := GenS1{F0: 18446744073709551615}
    v7 := GenS1{F0: 4294967295}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1000
}

func main() {
    var v1 bool = true
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 string = ("" + ("foo bar" + ("a" + (("" + "foo bar") + "a"))))
    var v5 int64 = -9223372036854775808
    for fi8 := 0; fi8 < 6; fi8++ {
        v4 = ""
        var v9 uint32 = 4294967295
        _ = fi8
        _ = v9
    }
    v1 = ((true || (false && (-9223372036854775808 <= -187))) || ((!true) && v1))
    v10 := GenS2{F0: 255, F1: 162, F2: 42}
    fmt.Println(v4, v4)
    v1 = v1
    for fi11 := 0; fi11 < 1; fi11++ {
        if (genFn0((v5 % ((v5) | 1)), 804) == 32767) {
            v12 := GenS2{F0: 42, F1: 0, F2: (v5 & (^(v5)))}
            v13 := v2[(-(fi11))]
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: 728, F1: 0, F2: -197}
            _ = v14
        }
        if (2 >= 0) {
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 int64 = (v5 & 1)
                if (!true) {
                    var v17 uint32 = 0
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
            fmt.Println(v5)
            if (0.1 < float64(3.14159)) {
                v4 = ("hello" + v4)
            }
        }
        _ = fi11
    }
    fmt.Println(v5, v5)
    v18 := GenS2{F0: 341, F1: 0, F2: v5}
    v4 = (v4 + ("a" + v4))
    if true {
        _ = len(v2)
        v1 = (998 != 293)
    }
    v19 := GenS1{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
