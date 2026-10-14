package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int16; F2 uint }
type GenS3 struct { F0 string; F1 int16; F2 string }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint64 }
func genFn0(p0 int64, p1 int, p2 uint) string {
    if ((p0 * 42) < p0) {
        p1 = p1
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS1{F0: 0}
            p1 -= fi7
            var v9 uint16 = 2
            _ = fi7
            _ = v8
            _ = v9
        }
        p0 -= p0
    } else {
        var v10 uint = p2
        v11 := GenS1{F0: 3}
        _ = v10
        _ = v11
    }
    if false {
        v12 := []int{(p1 - p1), (p1 << ((p2) & 63)), p1, (p1 - p1)}
        var v13 float64 = -2.25
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: -303}
        if (!(!(192 >= 2))) {
            p0 += (p0 / ((p0) | 1))
        } else {
            var v15 float32 = (-2.25 / 3.14159)
            fmt.Println(v15, v15)
            fmt.Println(v15)
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 int8 = 42
            var v18 uint32 = 42
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v14
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}
func genFn1(p0 string, p1 uint8) bool {
    v19 := []int{846, 127}
    var v20 uint = 65535
    var v21 int = 488
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return (497 < 2147483647)
}

func main() {
    var v1 int32 = -1
    var v2 uint64 = 1
    v3 := make(map[int]bool)
    v4 := []string{}
    var v5 int = 0
    var v6 int16 = 42
    switch (v5 % (((v5 & (v5 + v5))) | 1)) {
    case -799:
        var v22 int32 = v1
        _ = v22
    default:
        v3[v5] = (65535 == 1)
    }
    v6 = v6
    delete(v3, 32768)
    if false {
        v4 = append(v4, ("\n\t" + ("hello" + (("a" + "hello") + ("a" + ("foo bar" + ""))))))
    } else {
        v23 := GenS2{F0: (v6 + (v6 + v6)), F1: (v6 / (((^(v6))) | 1)), F2: 18446744073709551615}
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 int16 = 1000
            _ = fi24
            _ = v25
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            if true {
                v27 := GenS2{F0: (v6 * (-(v6))), F1: v6, F2: 18446744073709551615}
                var v28 int64 = 252
                _ = v27
                _ = v28
            } else {
                _ = len(v3)
            }
            if true {
                switch (("a" + "\n\t") + "") {
                case "\n\t":
                    v2 -= 852
                default:
                    for fi29 := 0; fi29 < 3; fi29++ {
                        fmt.Println(v1, v5)
                        _ = fi29
                    }
                }
                v1 -= 2147483647
            }
            _ = fi26
        }
        _ = v23
    }
    var v30 int16 = v6
    v31 := GenS3{F0: "hello", F1: v30, F2: ""}
    v32 := GenS4{F0: 23, F1: 1}
    var v33 string = "\n\t"
    fmt.Println(v30)
    v30 += (v6 / ((-719) | 1))
    var v34 float64 = 0.0
    switch (^(v5)) {
    case 9223372036854775807:
        if false {
            v35 := GenS5{F0: 65535}
            v4 = append(v4, ((v33 + "") + (v33 + v33)))
            _ = v35
        }
    case 42:
        fmt.Println(v2)
    }
    v36 := GenS5{F0: 972}
    switch (^(v5)) {
    case 1:
        var v37 int16 = v6
        _ = v37
    }
    var v38 uint16 = 522
    if false {
        var v39 uint8 = 200
        switch v33 {
        case "\n\t":
            v3[(v5 * v5)] = false
        case "foo bar":
            if false {
                v40 := GenS6{F0: 189, F1: (v2 & 1), F2: (^(v2))}
                v33 = ((v33 + ((("hello" + v33) + "\n\t") + v33)) + v33)
                v4 = append(v4, "a")
                _ = v40
            }
        default:
            v2 = (^(v2))
        }
        _ = v39
    }
    v34 = -2.25
    for fi41 := 0; fi41 < 3; fi41++ {
        for fi42 := 0; fi42 < 6; fi42++ {
            if len(v4) > 0 {
                v4[0] = (v33 + (v33 + v33))
            }
            _ = fi42
        }
        v5 = v5
        v33 = "\n\t"
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
