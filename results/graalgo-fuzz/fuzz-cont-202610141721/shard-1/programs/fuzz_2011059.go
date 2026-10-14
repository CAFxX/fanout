package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int; F2 int64 }
func genFn0(p0 bool, p1 bool) int8 {
    p0 = (-1 > 127)
    p1 = false
    fmt.Println(p1, p0)
    var v7 string = "hello"
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1(p0 int, p1 uint64) uint {
    if false {
        fmt.Println(p0, p1)
        p1 -= 4294967295
        p1 = (p1 % (((p1 / ((976) | 1))) | 1))
    } else {
        for fi8 := 0; fi8 < 3; fi8++ {
            if true {
                var v9 int = p0
                _ = v9
            }
            _ = fi8
        }
        v10 := GenS1{F0: "hello", F1: -2147483648, F2: -1}
        _ = v10
    }
    p1 += 4294967295
    v11 := []int{p0, p0, 2147483647}
    _ = p0
    _ = p1
    _ = v11
    return 42
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]uint32)
    v3 := []string{(("\n\t" + (v1 + (((("a" + "foo bar") + "\n\t") + (v1 + "foo bar")) + ("hello" + v1)))) + ((v1 + "hello") + ("\n\t" + v1))), v1, v1}
    v4 := []int{}
    var v5 uint = 130
    v6 := []int{-31, -9223372036854775808}
    var v12 uint8 = 112
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v12 += (v12 | (v12 << ((2) & 63)))
    for fi13 := 0; fi13 < 17; fi13++ {
        if true {
            v4 = append(v4, (fi13 >> ((v5) & 63)))
            _ = len(v2)
            fmt.Println(fi13, fi13, fi13)
        } else {
            v14 := GenS1{F0: v1, F1: 772, F2: 32767}
            _ = v14
        }
        var v15 bool = (((v1 + (v1 + v1)) + "a") <= ((v1 + "hello") + v1))
        _ = fi13
        _ = v15
    }
    v12 += v12
    fmt.Println(v5)
    var v16 uint = 4294967295
    if (100 <= -128) {
        v16 = v5
        fmt.Println(v5, v5, v12)
    } else {
        for fi17 := 0; fi17 < 13; fi17++ {
            fmt.Println(fi17, v1, v16)
            if (!true) {
                switch "\n\t" {
                case "hello":
                    if false {
                        var v18 int64 = 9223372036854775807
                        for fi19 := 0; fi19 < 5; fi19++ {
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                            fmt.Println(v16)
                            var v20 uint64 = 2
                            _ = fi19
                            _ = v20
                        }
                        var v21 uint32 = 165
                        _ = v18
                        _ = v21
                    }
                }
                if ((v12 >> ((v5) & 63)) > 0) {
                    v16 -= (v16 / ((v5) | 1))
                }
                for fi22 := 0; fi22 < 9; fi22++ {
                    v23 := v2[-270]
                    v12 += (^(v12))
                    v5 -= genFn1((fi22 % (((-(fi22))) | 1)), 2)
                    _ = fi22
                    _ = v23
                }
            } else {
                v24 := GenS2{F0: 0, F1: 65535}
                _ = v24
            }
            for fi25 := 0; fi25 < 14; fi25++ {
                delete(v2, -222)
                v5 += genFn1(fi25, 0)
                v26 := GenS3{F0: 0, F1: (-(fi25)), F2: int64(v16)}
                _ = fi25
                _ = v26
            }
            _ = fi17
        }
        v27 := GenS1{F0: (("hello" + "") + ((("hello" + v1) + v1) + "\n\t")), F1: -342, F2: 0}
        var v28 bool = ((v1 + v1) >= v1)
        _ = v27
        _ = v28
    }
    var v29 bool = (uint(v12) >= v16)
    var v30 uint32 = 2147483648
    var v31 uint32 = v30
    for fi32 := 0; fi32 < 14; fi32++ {
        v4 = append(v4, fi32)
        _ = fi32
    }
    v29 = (!(false && v29))
    for fi33 := 0; fi33 < 10; fi33++ {
        if (!false) {
            if false {
                var v34 int16 = 0
                var v35 uint64 = 224
                _ = len(v4)
                _ = v34
                _ = v35
            }
            v36 := GenS1{F0: "", F1: 685, F2: 32767}
            _ = v36
        }
        fmt.Println(v12, v30)
        v37 := GenS3{F0: 681, F1: 4294967295, F2: -9223372036854775808}
        _ = fi33
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
