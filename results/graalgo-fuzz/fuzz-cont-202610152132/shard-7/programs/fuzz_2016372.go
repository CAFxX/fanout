package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int64 = 0
    v2 := []string{(("foo bar" + ("hello" + ("a" + ("a" + "foo bar")))) + "foo bar"), "hello", "foo bar", "foo bar"}
    var v3 float32 = -2.25
    var v4 uint32 = 2147483647
    var v5 bool = (!((^(v1)) == (^(v1))))
    fmt.Println(v4, v3, v1)
    fmt.Println(v4)
    _ = len(v2)
    v5 = v5
    v6 := GenS1{F0: 1000, F1: 0, F2: 209}
    for fi7 := 0; fi7 < 10; fi7++ {
        _ = len(v2)
        _ = fi7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int16 = -32768
        fmt.Println(v4)
        _ = fi8
        _ = v9
    }
    var v10 uint = 1
    for fi11 := 0; fi11 < 15; fi11++ {
        v4 += v4
        fmt.Println(v5, v5, v4)
        _ = fi11
    }
    fmt.Println(v4, v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        fmt.Println(v10)
    } else {
        if (v5 || v5) {
            if (-1 > -2147483648) {
                if ((v10 < v10) && v5) {
                    v12 := GenS2{F0: 1000}
                    var v13 float64 = (0.1 + 1e10)
                    _ = v12
                    _ = v13
                } else {
                    var v14 int64 = 42
                    v10 += v10
                    v15 := GenS2{F0: 32767}
                    _ = v14
                    _ = v15
                }
            } else {
                v16 := GenS1{F0: 1000, F1: 42, F2: 42}
                _ = v16
            }
            var v17 int = -314
            _ = v17
        } else {
            v2 = append(v2, "a")
            if true {
                v18 := GenS1{F0: 1, F1: (^(v10)), F2: 65535}
                _ = v18
            }
            if v5 {
                v19 := GenS2{F0: 0}
                _ = v19
            } else {
                if (42 >= 2147483647) {
                    fmt.Println(v10, v1)
                }
                v4 += v4
            }
        }
        var v20 int32 = 2147483647
        fmt.Println(v3)
        _ = v20
    }
    var v21 uint8 = 44
    v3 -= v3
    if ((v4 - v4) == v4) {
        v4 = (v4 * v4)
    } else {
        v22 := GenS2{F0: 42}
        v4 = (v4 + 2147483648)
        _ = v22
    }
    v23 := GenS1{F0: 0, F1: 850, F2: 1}
    v10 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
