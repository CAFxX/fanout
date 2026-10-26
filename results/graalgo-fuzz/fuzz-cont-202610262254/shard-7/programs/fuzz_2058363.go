package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint32 }
func genFn0(p0 uint64, p1 int) uint16 {
    p0 = 4294967295
    v7 := []int{-290, -363, (p1 & (p1 % ((p1) | 1)))}
    for fi8 := 0; fi8 < 9; fi8++ {
        p1 = p1
        if (2147483647 > 0) {
            _ = len(v7)
            var v9 bool = (!(((0 != (^(p0))) || true) || true))
            v10 := GenS1{F0: ("foo bar" + ((("\n\t" + "a") + "") + "\n\t")), F1: ""}
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = v7
    return 255
}
func genFn1(p0 string) uint {
    if (-1 >= -1) {
        switch 42 {
        case -2147483648:
            v11 := []int{42, 9223372036854775807, 290, 32767}
            _ = v11
        case 0:
            p0 = p0
        default:
            switch (p0 + "\n\t") {
            case "\n\t":
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 float32 = -2.25
                    _ = fi12
                    _ = v13
                }
            default:
                switch p0 {
                case "":
                    var v14 int64 = -1
                    _ = v14
                case "foo bar":
                    v15 := GenS2{F0: 170, F1: 18446744073709551615}
                    _ = v15
                }
            }
        }
    } else {
        p0 = "a"
    }
    fmt.Println(p0)
    _ = p0
    return 1
}

func main() {
    var v1 float32 = 0.1
    v2 := []string{"\n\t", ("foo bar" + "foo bar"), ""}
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    var v5 string = (("hello" + ("hello" + ("a" + "hello"))) + "a")
    var v6 uint = 1
    var v16 int32 = -385
    for fi17 := 0; fi17 < 1; fi17++ {
        if false {
            v6 = (v6 ^ v6)
            var v18 bool = (!true)
            _ = v18
        } else {
            v16 += (-(v16))
            v19 := GenS3{F0: 42, F1: 18446744073709551615, F2: 525}
            var v20 float64 = 3.14159
            _ = v19
            _ = v20
        }
        v3[v5] = fi17
        v1 = v1
        _ = fi17
    }
    v16 = v16
    var v21 uint32 = 1
    v21 = (^(v21))
    v1 += v1
    v21 += 0
    if true {
        switch 446 {
        case 9223372036854775807:
            var v22 uint = 2
            _ = v22
        case -605:
            v6 -= v6
        case -9223372036854775808:
            v5 = ("foo bar" + (v5 + (v5 + v5)))
        }
        v6 -= v6
        for fi23 := 0; fi23 < 5; fi23++ {
            fmt.Println(fi23, v5)
            if false {
                _ = len(v4)
                if false {
                    delete(v4, fi23)
                    fmt.Println(v16, v6, v21)
                }
            } else {
                if false {
                    v24 := GenS2{F0: 23, F1: v6}
                    _ = v24
                } else {
                    v6 -= uint(v21)
                    v25 := GenS1{F0: "foo bar", F1: "\n\t"}
                    _ = v25
                }
                for fi26 := 0; fi26 < 19; fi26++ {
                    v5 = (v5 + ((v5 + v5) + ("a" + v5)))
                    if (127 <= 42) {
                        var v27 float32 = ((float32(0.0) + float32(3.14159)) + (1e10 * v1))
                        v16 -= (-(v16))
                        var v28 string = v5
                        _ = v27
                        _ = v28
                    } else {
                        v6 = (v6 ^ (v6 + v6))
                        if false {
                            var v29 uint = genFn1("a")
                            var v30 float64 = 1e10
                            _ = v29
                            _ = v30
                        }
                    }
                    _ = fi26
                }
            }
            fmt.Println(v1)
            _ = fi23
        }
    }
    var v31 int32 = -2147483648
    var v32 uint8 = 116
    var v33 uint16 = 522
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
