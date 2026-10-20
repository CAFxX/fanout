package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 string }
func genFn0(p0 string, p1 uint, p2 uint8) uint8 {
    p2 = (p2 << (((p1 >> ((0) & 63))) & 63))
    var v5 uint64 = 0
    p1 += (p1 % ((p1) | 1))
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}
func genFn1(p0 int8, p1 int16) bool {
    switch "a" {
    case "\n\t":
        v6 := []int{-90, 42, 127, 32767}
        _ = v6
    }
    p1 = p1
    if true {
        switch 1 {
        case 9223372036854775807:
            fmt.Println(p1, p1)
        case 882:
            v7 := []int{}
            _ = v7
        default:
            v8 := GenS1{F0: 2147483647}
            _ = v8
        }
        var v9 string = ("hello" + ("a" + "\n\t"))
        v10 := GenS2{F0: ((v9 + v9) + "hello"), F1: "foo bar"}
        _ = v10
        _ = v9
    }
    p0 = (^(p0))
    p0 -= p0
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 int8 = 0
    var v2 int = 0
    var v3 int16 = 1000
    var v4 int8 = (v1 ^ (v1 % ((v1) | 1)))
    fmt.Println(v3, v4)
    v1 += 127
    v4 -= v1
    v3 -= (v3 ^ v3)
    v11 := []int{v2}
    v12 := GenS2{F0: ("a" + (("hello" + ("\n\t" + "hello")) + "")), F1: "\n\t"}
    if false {
        fmt.Println(v3, v2)
    } else {
        switch 4294967295 {
        case 32767:
            for fi13 := 0; fi13 < 1; fi13++ {
                switch (fi13 << ((18446744073709551615) & 63)) {
                case 128:
                    fmt.Println(v4)
                case 224:
                    var v14 int8 = 127
                    _ = v14
                case 32768:
                    for fi15 := 0; fi15 < 14; fi15++ {
                        v16 := GenS2{F0: "foo bar", F1: "a"}
                        _ = len(v11)
                        if false {
                            v17 := GenS2{F0: "", F1: ("" + "\n\t")}
                            v4 = v4
                            v11 = append(v11, (fi13 & (fi13 >> ((82) & 63))))
                            _ = v17
                        } else {
                            v2 -= fi13
                            v2 -= (v2 >> ((18446744073709551615) & 63))
                            var v18 int32 = -2147483648
                            _ = v18
                        }
                        _ = fi15
                        _ = v16
                    }
                }
                fmt.Println(v1)
                _ = fi13
            }
        default:
            v19 := GenS2{F0: "\n\t", F1: ("\n\t" + "\n\t")}
            _ = v19
        }
        var v20 bool = false
        if (1 <= 42) {
            var v21 bool = v20
            v21 = (int(v4) == v2)
            _ = v21
        }
        _ = v20
    }
    if true {
        for fi22 := 0; fi22 < 7; fi22++ {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            if (v1 == 100) {
                if true {
                    v11 = append(v11, 396)
                }
                v2 += int(v3)
            }
            for fi23 := 0; fi23 < 11; fi23++ {
                v4 = v1
                _ = fi23
            }
            _ = fi22
        }
    } else {
        v4 += (v4 & 42)
        for fi24 := 0; fi24 < 14; fi24++ {
            if len(v11) > 0 {
                v11[0] = fi24
            }
            _ = fi24
        }
    }
    v25 := GenS1{F0: 258}
    var v26 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
