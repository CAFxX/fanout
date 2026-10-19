package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int64, p1 uint64, p2 int16) uint64 {
    p0 = (p0 << ((117) & 63))
    switch 9223372036854775807 {
    case -28:
        for fi6 := 0; fi6 < 2; fi6++ {
            p1 = (^(p1))
            if ((65535 > 0) && (0 != 1)) {
                p1 += (p1 - p1)
            }
            _ = fi6
        }
    case 32767:
        p0 += -9223372036854775808
    case 230:
        switch "foo bar" {
        case "\n\t":
            v7 := []int{32767, 9223372036854775807, 568, 307}
            _ = v7
        case "foo bar":
            p0 -= (p0 / (((p0 & -9223372036854775808)) | 1))
        case "a":
            var v8 uint8 = 127
            _ = v8
        default:
            p1 = 0
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 uint16, p1 uint) uint8 {
    if (false && true) {
        if (!(!((-2.25 < -2.25) && (!(18446744073709551615 <= 18446744073709551615))))) {
            if false {
                var v9 int64 = 257
                _ = v9
            }
            switch "" {
            case "hello":
                fmt.Println(p0, p1)
            }
            var v10 uint = (p1 % (((p1 & (p1 - p1))) | 1))
            _ = v10
        }
        p0 = p0
        for fi11 := 0; fi11 < 7; fi11++ {
            switch (fi11 >> ((812) & 63)) {
            case -410:
                v12 := []int{fi11, fi11, fi11, fi11}
                _ = v12
            case 128:
                v13 := []int{(fi11 / (((fi11 % ((fi11) | 1))) | 1)), (fi11 << ((p1) & 63)), (fi11 * -455)}
                _ = v13
            case 42:
                v14 := []int{fi11, 127, (fi11 ^ fi11), 128}
                _ = v14
            }
            p0 = p0
            _ = fi11
        }
    }
    v15 := []int{4294967295, 127}
    p1 -= (^(p1))
    if len(v15) > 0 {
        v15[0] = 9223372036854775807
    }
    _ = p0
    _ = p1
    _ = v15
    return 1
}

func main() {
    v1 := []uint32{18, 1, 2147483648}
    var v2 uint = 825
    var v3 int16 = 533
    v4 := []uint32{0, 2147483647, 676, 2}
    var v5 int = 1
    var v16 uint = v2
    var v17 string = "\n\t"
    var v18 uint = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v17 = v17
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((0.1 > (0.1 + 0.0)) || (false && ((^(v3)) <= v3))) {
        switch (v17 + "") {
        case "":
            if (v17 > (v17 + "foo bar")) {
                fmt.Println(v3, v17)
                for fi19 := 0; fi19 < 14; fi19++ {
                    switch (((v17 + v17) + v17) + v17) {
                    case "\n\t":
                        if true {
                            var v20 uint16 = 421
                            var v21 int64 = -9223372036854775808
                            _ = v20
                            _ = v21
                        } else {
                            fmt.Println(v18, v5, v18)
                            for fi22 := 0; fi22 < 10; fi22++ {
                                v5 = 42
                                v17 = v17
                                _ = fi22
                            }
                        }
                    default:
                        v18 = v18
                    }
                    _ = fi19
                }
            }
        case "\n\t":
            if ((v16 ^ v2) <= v2) {
                v17 = v17
            } else {
                var v23 uint32 = 1
                _ = v23
            }
        default:
            var v24 uint = 255
            _ = v24
        }
    } else {
        for fi25 := 0; fi25 < 5; fi25++ {
            switch "a" {
            case "foo bar":
                v5 = (^(v5))
            case "hello":
                fmt.Println(v5, v3)
            }
            v26 := GenS1{F0: v3}
            var v27 bool = (2 == 4294967295)
            _ = fi25
            _ = v26
            _ = v27
        }
        if len(v4) > 0 {
            v4[0] = 1
        }
    }
    var v28 int64 = 0
    v2 += v2
    var v29 int16 = (v3 * -32768)
    v30 := GenS1{F0: (v29 + (v29 % ((v3) | 1)))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
