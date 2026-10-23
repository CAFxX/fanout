package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
func genFn0(p0 int32, p1 int8, p2 bool) string {
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(p2, p2)
        for fi8 := 0; fi8 < 15; fi8++ {
            switch "" {
            case "hello":
                fmt.Println(fi8, fi8, fi7)
            case "foo bar":
                var v9 uint64 = 209
                _ = v9
            case "a":
                var v10 int = fi7
                _ = v10
            default:
                switch (fi7 - (fi8 % (((-(fi8))) | 1))) {
                case -9223372036854775808:
                    var v11 float32 = -2.25
                    _ = v11
                case -356:
                    p1 = (p1 * p1)
                }
            }
            _ = fi8
        }
        if p2 {
            switch ("a" + "foo bar") {
            case "a":
                p1 -= p1
            case "hello":
                v12 := []int{-253, (fi7 >> ((830) & 63))}
                _ = v12
            case "":
                if (0.1 == 0.0) {
                    p1 = 127
                    fmt.Println(p0)
                } else {
                    p1 = (p1 + (p1 ^ p1))
                    var v13 uint = 255
                    _ = v13
                }
            }
            var v14 int8 = p1
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 int8 = v14
                v17 := []int{fi15, -1, (fi15 + (^(fi15)))}
                fmt.Println(v16, p0, fi15)
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
        _ = fi7
    }
    fmt.Println(p2, p0, p2)
    var v18 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return ("\n\t" + "\n\t")
}
func genFn1(p0 int32) uint16 {
    var v19 int64 = 0
    v19 = 9223372036854775807
    v20 := []int{}
    fmt.Println(v19)
    var v21 int8 = 127
    _ = p0
    _ = v19
    _ = v20
    _ = v21
    return 988
}

func main() {
    var v1 uint64 = 658
    v2 := make(map[int]bool)
    var v3 int32 = -2147483648
    v4 := []int{1}
    var v5 int = -9223372036854775808
    v6 := []string{"", (("a" + "") + "\n\t"), "a", "foo bar"}
    var v22 uint64 = 380
    v3 += -1
    var v23 uint = 65535
    for fi24 := 0; fi24 < 19; fi24++ {
        switch ("foo bar" + ("foo bar" + "hello")) {
        case "hello":
            var v25 int16 = -565
            _ = v25
        case "\n\t":
            var v26 int16 = -44
            _ = v26
        default:
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 int8 = -1
                if (-9223372036854775808 == 373) {
                    switch (fi27 ^ fi27) {
                    case -9223372036854775808:
                        if true {
                            v22 = (v1 | v22)
                            v23 = 4294967295
                            v3 -= (^(v3))
                        }
                    }
                } else {
                    v3 = (-(v3))
                    v23 = v23
                    if (true || (true && false)) {
                        v23 = (v23 - 1)
                        v1 = (v22 | 18446744073709551615)
                        var v29 uint8 = 0
                        _ = v29
                    } else {
                        if len(v4) > 0 {
                            v4[0] = -242
                        }
                        var v30 uint64 = v1
                        v31 := GenS1{F0: 784, F1: 394}
                        _ = v30
                        _ = v31
                    }
                    if ((1e10 == 3.14159) && false) {
                        var v32 uint32 = 2
                        fmt.Println(v1)
                        var v33 uint16 = 733
                        _ = v32
                        _ = v33
                    }
                }
                v34 := GenS1{F0: -9223372036854775808, F1: 2}
                if (568 != v3) {
                    v22 = (v22 >> (((v23 / (((v23 - uint(v1))) | 1))) & 63))
                    v4 = append(v4, 4294967295)
                }
                _ = fi27
                _ = v28
                _ = v34
            }
        }
        _ = fi24
    }
    var v35 uint8 = 2
    var v36 int = 0
    var v37 string = ("foo bar" + "")
    var v38 string = "\n\t"
    switch v37 {
    case "hello":
        v22 = 18446744073709551615
    }
    var v39 int16 = -877
    fmt.Println(v38, v36)
    var v40 int64 = 526
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
