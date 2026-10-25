package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8, p1 int64, p2 int16) uint64 {
    v6 := []int{4294967295, 128, 19, -9223372036854775808}
    var v7 uint = 1
    fmt.Println(p1)
    v7 -= v7
    if (2147483648 <= 0) {
        var v8 uint = v7
        _ = v8
    } else {
        if (1 <= 2) {
            if false {
                p2 += (^(p2))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if false {
                    var v9 int32 = 631
                    v7 = 2
                    _ = v9
                } else {
                    v10 := GenS1{F0: (v7 % ((2) | 1))}
                    _ = v10
                }
            } else {
                p0 = 2
                p0 = 2
            }
            if false {
                for fi11 := 0; fi11 < 18; fi11++ {
                    v7 -= (^(v7))
                    _ = fi11
                }
                for fi12 := 0; fi12 < 12; fi12++ {
                    var v13 int32 = -1
                    _ = fi12
                    _ = v13
                }
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int32 = 2147483647
                _ = fi14
                _ = v15
            }
        }
        var v16 int = -9223372036854775808
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 uint32) uint8 {
    var v17 uint32 = 0
    if false {
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 bool = false
            _ = fi18
            _ = v19
        }
    }
    var v20 uint8 = 248
    switch "a" {
    case "":
        var v21 float32 = (1e10 - 1.5)
        _ = v21
    case "\n\t":
        fmt.Println(p0, v17, v20)
    }
    var v22 int32 = 2147483647
    _ = p0
    _ = v17
    _ = v20
    _ = v22
    return v20
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 255
    var v3 bool = ("\n\t" == ("" + ("foo bar" + ((("foo bar" + (("foo bar" + "foo bar") + "foo bar")) + "foo bar") + "hello"))))
    var v4 uint = 866
    var v5 uint64 = 2
    v23 := []int{928, -839, 9223372036854775807}
    var v24 int = -612
    var v25 int32 = 2147483647
    var v26 int64 = 1
    for fi27 := 0; fi27 < 1; fi27++ {
        if true {
            var v28 int16 = -761
            if (!true) {
                var v29 uint64 = (v5 << ((42) & 63))
                _ = v29
            }
            _ = v28
        } else {
            for fi30 := 0; fi30 < 15; fi30++ {
                var v31 bool = true
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                v32 := GenS2{F0: 4294967295, F1: 93}
                fmt.Println(v31, v2)
                v33 := GenS3{F0: 0}
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
        }
        v3 = (-16 > 42)
        for fi34 := 0; fi34 < 11; fi34++ {
            v25 -= -1
            v2 -= (v2 * (v2 / ((1) | 1)))
            _ = fi34
        }
        _ = fi27
    }
    v25 += v25
    v24 -= 2147483647
    v2 -= (v2 - (v2 >> ((uint(v2)) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
