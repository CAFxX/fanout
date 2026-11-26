package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int64 }
func genFn0(p0 int32, p1 int) uint16 {
    p0 = 2147483647
    fmt.Println(p0)
    var v7 uint16 = 723
    if ((v7 & v7) < (v7 / ((255) | 1))) {
        var v8 int8 = 100
        v9 := []int{p1, (p1 | (p1 % ((592) | 1)))}
        p1 -= (^(p1))
        _ = v8
        _ = v9
    }
    var v10 uint8 = 12
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return (v7 >> ((938) & 63))
}
func genFn1(p0 uint16, p1 int32, p2 string) int16 {
    v11 := []int{2147483647, 42, -90}
    if false {
        p1 = p1
        fmt.Println(p1, p2)
        if (!(!(1.5 != (-2.25 * 1.5)))) {
            if (p2 != (("a" + (p2 + "")) + "hello")) {
                switch 346 {
                case -472:
                    if (false || true) {
                        p0 -= 42
                        var v12 bool = (893 > 65535)
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 9; fi13++ {
                            p1 = (p1 & p1)
                            if true {
                                p0 = (p0 ^ genFn0((p1 | (^(p1))), fi13))
                            }
                            _ = fi13
                        }
                    }
                }
            } else {
                p0 += 405
                p1 = -823
                fmt.Println(p2, p2, p2)
            }
        } else {
            var v14 int = 128
            v15 := GenS1{F0: 18446744073709551615, F1: (^(v14)), F2: -1}
            _ = len(v11)
            _ = v14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        p1 = p1
        _ = fi16
    }
    if false {
        var v17 int8 = -10
        var v18 int32 = (p1 >> ((65535) & 63))
        _ = v17
        _ = v18
    } else {
        v19 := GenS1{F0: 477, F1: 2147483647, F2: -263}
        _ = len(v11)
        for fi20 := 0; fi20 < 15; fi20++ {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            var v21 uint16 = (p0 & genFn0((p1 % ((-259) | 1)), fi20))
            switch (^(fi20)) {
            case -432:
                for fi22 := 0; fi22 < 20; fi22++ {
                    fmt.Println(fi20, fi20, fi22)
                    v23 := GenS1{F0: 18446744073709551615, F1: 300, F2: -403}
                    v11 = append(v11, 42)
                    _ = fi22
                    _ = v23
                }
            default:
                var v24 int8 = -128
                _ = v24
            }
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        if false {
            _ = len(v11)
            var v26 int16 = 295
            fmt.Println(p1)
            _ = v26
        } else {
            if len(v11) > 0 {
                v11[0] = -9223372036854775808
            }
            if len(v11) > 0 {
                v11[0] = fi25
            }
            for fi27 := 0; fi27 < 2; fi27++ {
                v28 := GenS1{F0: 4294967295, F1: fi27, F2: 0}
                p2 = (p2 + ("a" + (p2 + p2)))
                var v29 int32 = (p1 >> ((294) & 63))
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        var v30 float64 = 1e10
        _ = fi25
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -468
}

func main() {
    var v1 uint = 42
    var v2 int = -2147483648
    var v3 uint16 = 376
    var v4 uint16 = v3
    v5 := []string{}
    var v6 int16 = 32767
    var v31 uint64 = 4294967295
    if true {
        v5 = append(v5, "\n\t")
    } else {
        switch -581 {
        case -31:
            for fi32 := 0; fi32 < 14; fi32++ {
                var v33 uint8 = 127
                v1 += v1
                _ = fi32
                _ = v33
            }
        }
    }
    v1 += v1
    var v34 int8 = 0
    switch (-(v2)) {
    case 128:
        v4 += 65535
    }
    v35 := GenS1{F0: (^(v31)), F1: v2, F2: int64(v6)}
    for fi36 := 0; fi36 < 11; fi36++ {
        if true {
            switch "\n\t" {
            case "foo bar":
                v2 -= -693
            }
            fmt.Println(v4)
        }
        fmt.Println(v1)
        _ = fi36
    }
    if len(v5) > 0 {
        v5[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
