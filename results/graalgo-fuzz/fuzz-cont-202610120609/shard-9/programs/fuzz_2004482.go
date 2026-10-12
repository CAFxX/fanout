package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint8, p1 int8, p2 uint16) bool {
    var v6 bool = (18446744073709551615 <= 503)
    v7 := GenS1{F0: 42, F1: 4294967295}
    p2 = p2
    if true {
        fmt.Println(v6, p2, p1)
        if v6 {
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS1{F0: 0, F1: 4294967295}
                switch ("hello" + ("a" + "a")) {
                case "\n\t":
                    if true {
                        v10 := GenS2{F0: p2, F1: 42}
                        for fi11 := 0; fi11 < 1; fi11++ {
                            var v12 int64 = -9223372036854775808
                            var v13 uint = 18446744073709551615
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        p0 = 42
                        _ = v10
                    }
                case "hello":
                    var v14 float64 = ((1e10 - 1e10) / 3.14159)
                    _ = v14
                case "a":
                    var v15 bool = (!(p1 != (p1 >> ((255) & 63))))
                    _ = v15
                }
                _ = fi8
                _ = v9
            }
        } else {
            v6 = (!false)
            var v16 uint16 = p2
            p0 = (^(p0))
            _ = v16
        }
        var v17 uint16 = p2
        _ = v17
    } else {
        var v18 uint8 = p0
        _ = v18
    }
    v19 := []int{int(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 uint, p1 uint16) uint64 {
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 float64 = (1.5 * 3.14159)
        for fi22 := 0; fi22 < 4; fi22++ {
            v21 = -2.25
            if true {
                var v23 bool = true
                fmt.Println(p1, fi20, p0)
                var v24 string = "\n\t"
                _ = v23
                _ = v24
            } else {
                var v25 string = ("" + ("a" + ""))
                p0 += p0
                v21 = v21
                _ = v25
            }
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    v26 := GenS2{F0: (p1 & (p1 % (((^(p1))) | 1))), F1: p0}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v26
    return uint64(p1)
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    var v3 uint64 = 1
    var v4 int64 = -480
    v5 := []int{}
    v2[103] = -1
    v27 := GenS1{F0: -52, F1: 0}
    v5 = append(v5, -2147483648)
    var v28 float64 = 0.1
    if (374 <= 866) {
        v29 := GenS3{F0: (^(v3)), F1: 0}
        _ = v29
    }
    var v30 uint = 42
    var v31 uint16 = 42
    var v32 int64 = 42
    v31 = v31
    if ((v3 % ((v3) | 1)) != v3) {
        if true {
            var v33 string = ""
            v1["hello"] = true
            _ = v33
        } else {
            if false {
                if (v30 != (v30 ^ v30)) {
                    v28 = v28
                    fmt.Println(v30, v31)
                } else {
                    v5 = append(v5, -741)
                    var v34 uint32 = 26
                    _ = v34
                }
                v1[""] = (true && (true && ("hello" <= "\n\t")))
                fmt.Println(v32)
            } else {
                for fi35 := 0; fi35 < 18; fi35++ {
                    if true {
                        v3 = genFn1(v30, 2)
                        v36 := v2[42]
                        _ = v36
                    }
                    fmt.Println(v30, v32)
                    fmt.Println(v4, v30, v4)
                    _ = fi35
                }
                fmt.Println(v32, v31)
                if (696 != 2147483647) {
                    fmt.Println(v32, v32, v30)
                } else {
                    var v37 bool = ((v31 >> ((2) & 63)) <= 65535)
                    _ = v37
                }
            }
        }
    } else {
        if (((v4 & v32) < (v32 / ((-9223372036854775808) | 1))) || (false && (-2.25 <= 0.1))) {
            v3 = (^(v3))
            v30 -= uint(v31)
            var v38 int32 = 42
            _ = v38
        } else {
            fmt.Println(v3, v32, v30)
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
