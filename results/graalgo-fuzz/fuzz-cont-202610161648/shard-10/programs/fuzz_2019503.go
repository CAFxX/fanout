package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 uint8; F2 int16 }
func genFn0(p0 uint64, p1 int) uint32 {
    if ("" > "a") {
        for fi5 := 0; fi5 < 13; fi5++ {
            for fi6 := 0; fi6 < 3; fi6++ {
                v7 := []int{}
                p0 += 0
                if true {
                    if true {
                        if len(v7) > 0 {
                            v7[0] = int(p0)
                        }
                        if (((p0 | p0) <= 409) || ((175 == 2147483647) && (127 < 2))) {
                            p1 += 32767
                        } else {
                            if true {
                                fmt.Println(fi6, fi5)
                            }
                        }
                    } else {
                        p0 = p0
                    }
                } else {
                    _ = len(v7)
                }
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
    } else {
        v8 := GenS1{F0: (p0 << ((18446744073709551615) & 63)), F1: uint8(p0)}
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 string = ("foo bar" + ("hello" + ("\n\t" + ("\n\t" + "hello"))))
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    var v11 string = "\n\t"
    for fi12 := 0; fi12 < 17; fi12++ {
        p0 += 1
        v13 := []int{}
        p1 += (fi12 + (p1 * fi12))
        _ = fi12
        _ = v13
    }
    v14 := []int{42}
    _ = p0
    _ = p1
    _ = v11
    _ = v14
    return 36
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{42}
    v3 := make(map[string]int)
    var v4 int = -2147483648
    var v15 float64 = (float64(((1.5 / 0.1) - (0.0 * -2.25))) + (0.1 - float64((float32(3.14159) * 0.0))))
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 float64 = -2.25
        _ = fi16
        _ = v17
    }
    fmt.Println(v15, v15)
    var v18 uint = 625
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 bool = (false || false)
        if v20 {
            fmt.Println(v15, fi19)
            v20 = v20
        }
        _ = fi19
        _ = v20
    }
    _ = len(v3)
    var v21 float64 = (v15 - v15)
    v18 = v18
    if false {
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 float32 = 0.0
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v15, v23)
            _ = fi22
            _ = v23
        }
        if len(v2) > 0 {
            v2[0] = 468
        }
        v21 -= (v15 + (3.14159 - v21))
    } else {
        _ = len(v3)
        v21 += 1.5
        _ = len(v3)
    }
    v15 -= v21
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        v18 = v18
        _ = fi24
    }
    fmt.Println(v4, v18, v15)
    var v25 float32 = ((1.5 / (float32(v21) * 1e10)) + (-2.25 + 0.0))
    v26 := GenS2{F0: 1, F1: 1}
    for fi27 := 0; fi27 < 18; fi27++ {
        if true {
            for fi28 := 0; fi28 < 6; fi28++ {
                _ = len(v3)
                switch "\n\t" {
                case "":
                    v4 += fi28
                case "foo bar":
                    if ((v18 % ((1) | 1)) > (^(v18))) {
                        var v29 uint = (^(v18))
                        fmt.Println(v18, v4, v15)
                        v2 = append(v2, 413)
                        _ = v29
                    }
                    v30 := GenS3{F0: 1}
                    _ = v30
                }
                fmt.Println(v25, v15, v25)
                switch ("foo bar" + "") {
                case "foo bar":
                    v21 = 0.1
                    v31 := GenS4{F0: 32767, F1: 0, F2: 1}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch "" {
                    case "":
                        v18 = (v18 * (v18 >> (((v18 % (((v18 % (((v18 ^ (v18 + v18))) | 1))) | 1))) & 63)))
                    }
                    _ = v31
                case "a":
                    var v32 bool = ((42 > 812) && true)
                    _ = v32
                case "":
                    var v33 int8 = 127
                    _ = v33
                }
                _ = fi28
            }
            v21 -= 0.1
        } else {
            var v34 float32 = v25
            _ = v34
        }
        var v35 int64 = 178
        _ = fi27
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
