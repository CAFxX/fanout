package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
func genFn0(p0 int64, p1 uint32, p2 uint16) int16 {
    p1 += (p1 % ((2147483647) | 1))
    if ((-2.25 - 0.0) != (0.0 - 0.0)) {
        v5 := []int{4294967295, -739, 32767}
        _ = v5
    } else {
        for fi6 := 0; fi6 < 15; fi6++ {
            p0 = 0
            _ = fi6
        }
        var v7 uint = 4294967295
        p1 += p1
        _ = v7
    }
    var v8 uint32 = (p1 % (((p1 << ((756) & 63))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return int16(v8)
}

func main() {
    var v1 int32 = -2147483648
    v2 := []uint32{2147483647, 4294967295}
    var v3 int32 = 42
    var v4 bool = false
    v1 -= v1
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    switch 42 {
    case 2147483647:
        if (!(!((970 >= 96) && (!(-2.25 <= float64(0.1)))))) {
            v9 := GenS1{F0: 1, F1: (v1 | 0)}
            _ = v9
        } else {
            if ((3.14159 - (3.14159 / -2.25)) < (0.0 - 1.5)) {
                var v10 bool = v4
                _ = v10
            } else {
                if ((v3 * -1) < -2147483648) {
                    for fi11 := 0; fi11 < 9; fi11++ {
                        if (243 > 159) {
                            var v12 int32 = (v1 & -2147483648)
                            if len(v2) > 0 {
                                v2[0] = 398
                            }
                            v13 := GenS1{F0: 214, F1: -2147483648}
                            _ = v12
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 20; fi14++ {
                            fmt.Println(fi14)
                            v15 := GenS2{F0: ""}
                            var v16 string = "foo bar"
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 7; fi17++ {
                            var v18 int16 = 32767
                            v18 = v18
                            _ = fi17
                            _ = v18
                        }
                        _ = fi11
                    }
                }
                var v19 int64 = 1
                _ = v19
            }
        }
    case 0:
        if len(v2) > 0 {
            v2[0] = 767
        }
    }
    var v20 bool = v4
    var v21 bool = false
    if (0 <= 100) {
        v3 += v1
        fmt.Println(v3, v20, v21)
        var v22 int32 = v1
        _ = v22
    }
    switch ("hello" + "") {
    case "a":
        var v23 string = ("a" + "\n\t")
        _ = v23
    case "\n\t":
        fmt.Println(v1, v4, v3)
    case "":
        if false {
            for fi24 := 0; fi24 < 9; fi24++ {
                fmt.Println(v20)
                _ = fi24
            }
        }
    }
    v1 += 1
    v25 := GenS3{F0: ("foo bar" + "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
}
