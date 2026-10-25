package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 int8, p1 int32) uint64 {
    for fi7 := 0; fi7 < 9; fi7++ {
        p0 -= 100
        var v8 int16 = 0
        _ = fi7
        _ = v8
    }
    p1 -= 129
    fmt.Println(p0, p1)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 472
}

func main() {
    var v1 int = -2147483648
    var v2 int = v1
    v3 := make(map[string]string)
    var v4 int = (v2 ^ (v2 % ((v2) | 1)))
    var v5 int64 = 9223372036854775807
    v6 := make(map[int]uint32)
    if (2147483647 == 0) {
        switch v4 {
        case 572:
            var v9 uint8 = 127
            _ = v9
        case 32767:
            var v10 float64 = 0.1
            _ = v10
        default:
            v11 := []int{}
            _ = v11
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            for fi13 := 0; fi13 < 11; fi13++ {
                if (3.14159 <= (3.14159 * 0.0)) {
                    fmt.Println(v2)
                    v6[(^(v1))] = 2
                }
                if false {
                    v14 := v6[v2]
                    _ = v14
                }
                v3[(("hello" + ("a" + "\n\t")) + (("hello" + "\n\t") + "foo bar"))] = "foo bar"
                _ = fi13
            }
            _ = fi12
        }
        if true {
            var v15 int32 = 581
            _ = v15
        } else {
            fmt.Println(v1, v1)
            _ = len(v3)
        }
    }
    var v16 float64 = 1.5
    v5 -= (v5 << ((0) & 63))
    switch "a" {
    case "\n\t":
        if false {
            v5 = 42
        } else {
            if true {
                fmt.Println(v16)
            }
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := GenS1{F0: 172, F1: 18446744073709551615}
                _ = fi17
                _ = v18
            }
            v19 := GenS2{F0: (v5 - (v5 / (((v5 % ((v5) | 1))) | 1)))}
            _ = v19
        }
    case "":
        v5 = (v5 & v5)
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        if ((255 <= 2) || (0 > 255)) {
            if (!(true || false)) {
                for fi21 := 0; fi21 < 11; fi21++ {
                    switch "hello" {
                    case "foo bar":
                        v16 -= (v16 * 0.1)
                    case "\n\t":
                        if (v4 != (v1 + (v1 * 42))) {
                            fmt.Println(fi21, v2)
                        }
                    case "a":
                        var v22 uint = 2
                        _ = v22
                    default:
                        switch (("\n\t" + "\n\t") + ("hello" + "")) {
                        case "a":
                            var v23 string = ("hello" + (("a" + "foo bar") + ""))
                            _ = v23
                        case "\n\t":
                            v24 := []int{(v2 - fi20), (^(v2)), v4}
                            _ = v24
                        case "hello":
                            switch 127 {
                            case -9223372036854775808:
                                var v25 uint16 = 249
                                _ = v25
                            case -584:
                                if (fi21 <= (v4 % ((-9223372036854775808) | 1))) {
                                    fmt.Println(fi21, v1, fi20)
                                    switch ("" + "a") {
                                    case "foo bar":
                                        v5 = 1
                                    default:
                                        v26 := GenS1{F0: 127, F1: 801}
                                        _ = v26
                                    }
                                } else {
                                    v6[(v4 + fi21)] = 958
                                    v16 += v16
                                }
                            }
                        }
                    }
                    switch -128 {
                    case 9223372036854775807:
                        var v27 int8 = 127
                        _ = v27
                    default:
                        v28 := GenS1{F0: 82, F1: 2}
                        v4 = v2
                        v2 -= v1
                        _ = v28
                    }
                    v29 := []int{}
                    _ = fi21
                    _ = v29
                }
                v30 := []int{-1, (-(fi20)), 42}
                _ = v30
            }
        } else {
            var v31 uint = 4294967295
            _ = v31
        }
        if false {
            var v32 int16 = -1
            _ = v32
        }
        _ = fi20
    }
    if (756 == 42) {
        v16 += -2.25
    }
    v33 := []int{(v2 & v1), 9223372036854775807}
    var v34 uint = 812
    if true {
        v35 := GenS3{F0: 993, F1: -431}
        _ = v35
    }
    fmt.Println(v5, v2)
    fmt.Println(v16, v4, v16)
    delete(v3, ("a" + "hello"))
    for fi36 := 0; fi36 < 19; fi36++ {
        for fi37 := 0; fi37 < 6; fi37++ {
            var v38 int = fi37
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    fmt.Println(v34)
    for fi39 := 0; fi39 < 8; fi39++ {
        v34 -= (v34 >> ((2) & 63))
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
