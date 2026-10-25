package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint16; F2 int8 }
func genFn0(p0 uint, p1 int) uint {
    if false {
        p0 += (p0 | 842)
        p0 = p0
    }
    v4 := GenS1{F0: 0}
    var v5 uint8 = 9
    if (-880 != -245) {
        if (!(false && true)) {
            if (42 >= -1) {
                v5 += v5
            } else {
                var v6 string = ("hello" + "a")
                p1 = -792
                _ = v6
            }
            v7 := GenS2{F0: p0}
            _ = v7
        }
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p0
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    _ = len(v3)
    v8 := []int{-697}
    v9 := GenS3{F0: 9223372036854775807, F1: 1, F2: 127}
    var v10 int32 = 42
    v1[0] = false
    v10 = v10
    switch (("foo bar" + "") + "") {
    case "":
        if (9223372036854775807 == 42) {
            var v11 uint64 = 1
            _ = v11
        } else {
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 uint8 = 2
                for fi14 := 0; fi14 < 3; fi14++ {
                    v8 = append(v8, (fi12 >> ((255) & 63)))
                    switch ("foo bar" + ("hello" + "foo bar")) {
                    case "foo bar":
                        var v15 int = fi14
                        _ = v15
                    case "\n\t":
                        fmt.Println(fi14, fi14)
                    case "":
                        var v16 int16 = -32768
                        _ = v16
                    default:
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    }
                    if ((fi14 ^ (fi12 & fi14)) == fi14) {
                        if (-397 >= -9223372036854775808) {
                            v17 := GenS3{F0: 905, F1: 255, F2: 42}
                            _ = v17
                        } else {
                            var v18 uint32 = 2
                            var v19 int = 128
                            _ = v18
                            _ = v19
                        }
                        v10 += (v10 - (v10 << ((705) & 63)))
                    }
                    if ((!(true && ((v10 / ((894) | 1)) == (v10 % ((v10) | 1))))) || true) {
                        v10 += (v10 << ((255) & 63))
                    } else {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    }
                    _ = fi14
                }
                _ = fi12
                _ = v13
            }
            v10 -= v10
            switch -304 {
            case 0:
                var v20 int64 = -446
                _ = v20
            }
        }
    case "a":
        var v21 int8 = 110
        _ = v21
    case "hello":
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 int8 = 42
            for fi24 := 0; fi24 < 19; fi24++ {
                if (0 >= 223) {
                    v10 = v10
                }
                v10 = v10
                _ = len(v2)
                v25 := v1[fi22]
                v26 := GenS2{F0: 18446744073709551615}
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi22
            _ = v23
        }
    }
    v3[42] = 507
    for fi27 := 0; fi27 < 6; fi27++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
