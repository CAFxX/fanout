package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint64 }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 uint32) uint32 {
    v6 := []int{9223372036854775807, 0}
    fmt.Println(p0)
    p0 -= p0
    var v7 string = (("foo bar" + "foo bar") + "hello")
    switch v7 {
    case "":
        var v8 uint = 42
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: 42, F1: "foo bar", F2: 1}
            _ = fi9
            _ = v10
        }
    default:
        v7 = v7
    }
    _ = p0
    _ = v6
    _ = v7
    return (p0 * p0)
}
func genFn1(p0 int64, p1 int16, p2 int8) uint {
    fmt.Println(p1, p1)
    for fi11 := 0; fi11 < 14; fi11++ {
        p2 = (p2 * 127)
        _ = fi11
    }
    switch "\n\t" {
    case "\n\t":
        var v12 int = -993
        _ = v12
    case "hello":
        if true {
            p1 += (-(p1))
            if (!(!(-128 != p2))) {
                for fi13 := 0; fi13 < 11; fi13++ {
                    fmt.Println(fi13, p0)
                    switch (^(fi13)) {
                    case -9223372036854775808:
                        p1 += p1
                    case -303:
                        for fi14 := 0; fi14 < 5; fi14++ {
                            p2 -= 127
                            fmt.Println(p1, fi14)
                            _ = fi14
                        }
                    }
                    v15 := []int{}
                    _ = fi13
                    _ = v15
                }
                if true {
                    v16 := GenS2{F0: 9223372036854775807, F1: ("hello" + "hello")}
                    var v17 uint32 = 720
                    if false {
                        p0 = 9223372036854775807
                    } else {
                        for fi18 := 0; fi18 < 1; fi18++ {
                            v19 := []int{4294967295}
                            var v20 int = fi18
                            var v21 uint8 = 127
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    }
                    _ = v16
                    _ = v17
                } else {
                    var v22 uint16 = 2
                    _ = v22
                }
                v23 := GenS1{F0: 175, F1: "\n\t", F2: 4294967295}
                _ = v23
            }
        }
    case "a":
        v24 := GenS2{F0: 784, F1: "hello"}
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 float32 = ((1e10 * -2.25) + (3.14159 / float32(float64(-2.25))))
    var v2 uint = 18446744073709551615
    var v3 int16 = 42
    var v4 uint64 = 42
    v5 := []int{1}
    var v25 float64 = 1e10
    v3 = 42
    v3 += v3
    var v26 float64 = (0.1 / v25)
    v26 = 1.5
    v27 := GenS2{F0: 993, F1: ("" + ("" + "a"))}
    var v28 float32 = 0.1
    switch 0 {
    case -52:
        fmt.Println(v26, v1)
    case 196:
        switch "hello" {
        case "\n\t":
            for fi29 := 0; fi29 < 16; fi29++ {
                for fi30 := 0; fi30 < 11; fi30++ {
                    var v31 float32 = 1e10
                    _ = fi30
                    _ = v31
                }
                var v32 float32 = v1
                fmt.Println(v2, v32)
                _ = fi29
                _ = v32
            }
        }
    default:
        if false {
            if (false || true) {
                v1 += 3.14159
                var v33 int16 = v3
                var v34 int8 = 0
                _ = v33
                _ = v34
            }
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (-1 >= (v3 & (v3 - (v3 + (v3 & v3))))) {
        v35 := GenS1{F0: v4, F1: "foo bar", F2: 1}
        _ = v35
    } else {
        v36 := GenS2{F0: -977, F1: ""}
        if ((!(!false)) && false) {
            fmt.Println(v25)
            for fi37 := 0; fi37 < 10; fi37++ {
                fmt.Println(v28)
                for fi38 := 0; fi38 < 1; fi38++ {
                    v39 := GenS1{F0: v4, F1: "a", F2: (^(v4))}
                    fmt.Println(v26)
                    v25 -= (((((v25 - float64(3.14159)) * 1e10) + float64((-2.25 - ((v1 / v28) - -2.25)))) * (v26 - v25)) + v25)
                    _ = fi38
                    _ = v39
                }
                for fi40 := 0; fi40 < 4; fi40++ {
                    v5 = append(v5, fi40)
                    var v41 uint16 = 65535
                    _ = len(v5)
                    v1 += (0.0 + 1e10)
                    fmt.Println(v4, v3, v25)
                    _ = fi40
                    _ = v41
                }
                _ = fi37
            }
        } else {
            for fi42 := 0; fi42 < 7; fi42++ {
                _ = len(v5)
                _ = fi42
            }
            for fi43 := 0; fi43 < 5; fi43++ {
                v44 := GenS2{F0: (fi43 << (((v2 + 290)) & 63)), F1: ("a" + ("\n\t" + ("a" + ("" + "\n\t"))))}
                v4 -= (v4 / ((608) | 1))
                var v45 int8 = 42
                _ = fi43
                _ = v44
                _ = v45
            }
        }
        if (true && false) {
            v1 = v1
            v28 += -2.25
        }
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
