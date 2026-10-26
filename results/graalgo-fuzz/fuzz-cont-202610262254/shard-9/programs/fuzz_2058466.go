package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int32 }
func genFn0(p0 int16) int {
    var v7 bool = false
    v8 := GenS1{F0: v7}
    v7 = false
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    _ = v8
    return -533
}

func main() {
    var v1 int = 4294967295
    var v2 int16 = -548
    v3 := []uint32{2, 222, 4294967295}
    v4 := make(map[int]bool)
    var v5 int = v1
    var v6 int = v1
    var v9 uint16 = 65535
    var v10 string = "\n\t"
    if true {
        switch (("" + v10) + ((v10 + ("foo bar" + v10)) + v10)) {
        case "foo bar":
            v6 = (^(v6))
        case "\n\t":
            if (!false) {
                v6 += (v6 / ((213) | 1))
                var v11 string = (v10 + ("foo bar" + v10))
                _ = v11
            }
        default:
            v6 -= (v5 ^ (v5 * 127))
        }
    } else {
        v5 = (v1 + 1)
        v9 = 650
        v12 := GenS1{F0: true}
        _ = v12
    }
    fmt.Println(v1)
    if false {
        v13 := GenS1{F0: false}
        var v14 string = (v10 + v10)
        var v15 int = (v1 & 32767)
        _ = v13
        _ = v14
        _ = v15
    } else {
        var v16 int64 = -1
        fmt.Println(v16)
        _ = v16
    }
    v10 = (v10 + "a")
    fmt.Println(v10)
    fmt.Println(v2, v9, v1)
    var v17 uint64 = 1
    var v18 int32 = -1
    v5 = (v6 << ((672) & 63))
    for fi19 := 0; fi19 < 19; fi19++ {
        switch v10 {
        case "hello":
            v20 := GenS1{F0: false}
            _ = v20
        case "\n\t":
            switch genFn0((-(v2))) {
            case -885:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case "foo bar":
            v2 += v2
        default:
            if (0 != 2147483647) {
                var v21 bool = (!true)
                for fi22 := 0; fi22 < 19; fi22++ {
                    fmt.Println(v5, v17, v17)
                    _ = fi22
                }
                if v21 {
                    var v23 int32 = int32(v9)
                    v23 += v18
                    _ = v23
                } else {
                    if v21 {
                        switch ("hello" + (("foo bar" + ("\n\t" + v10)) + v10)) {
                        case "hello":
                            v6 += (v6 * genFn0(32767))
                        default:
                            var v24 uint64 = 607
                            _ = v24
                        }
                        if (v17 == 42) {
                            var v25 int32 = (v18 + (v18 & v18))
                            _ = v25
                        }
                        switch (v1 - (v6 >> ((1) & 63))) {
                        case -706:
                            fmt.Println(v5, v17)
                        case 4294967295:
                            fmt.Println(v5, v5, v18)
                        case 32767:
                            v26 := GenS1{F0: v21}
                            _ = v26
                        default:
                            v2 = int16(v1)
                        }
                    }
                }
                _ = v21
            }
        }
        v10 = v10
        _ = fi19
    }
    fmt.Println(v10, v9)
    for fi27 := 0; fi27 < 4; fi27++ {
        switch v1 {
        case -9223372036854775808:
            v28 := GenS2{F0: -827, F1: (-(v18))}
            _ = v28
        default:
            var v29 bool = (!true)
            _ = v29
        }
        if true {
            v9 -= v9
        } else {
            v5 = (v6 % ((0) | 1))
        }
        _ = fi27
    }
    v5 = (-(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
