package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint32 }
func genFn0(p0 string, p1 string) string {
    for fi4 := 0; fi4 < 9; fi4++ {
        fmt.Println(p1, fi4)
        if false {
            if (((1.5 * -2.25) >= (0.0 + (1.5 + float64(0.0)))) || true) {
                v5 := []int{}
                _ = v5
            } else {
                v6 := []int{(^(fi4))}
                _ = v6
            }
        }
        var v7 uint = 2
        _ = fi4
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: 320, F1: p1}
        _ = fi8
        _ = v9
    }
    var v10 uint = 0
    v10 = (v10 << (((^(v10))) & 63))
    _ = p0
    _ = p1
    _ = v10
    return (p1 + "a")
}
func genFn1(p0 int, p1 int64, p2 bool) int8 {
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint64 = 4294967295
        if (p2 || false) {
            for fi13 := 0; fi13 < 7; fi13++ {
                p2 = (!p2)
                _ = fi13
            }
        }
        _ = fi11
        _ = v12
    }
    v14 := []int{32768, 127, (p0 % (((p0 >> ((42) & 63))) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return -128
}

func main() {
    var v1 int32 = 299
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    var v15 int = 41
    v16 := []int{(^(v15))}
    fmt.Println(v1, v15)
    var v17 uint32 = 65535
    v1 -= 0
    v17 = (v17 / ((v17) | 1))
    if false {
        for fi18 := 0; fi18 < 6; fi18++ {
            _ = len(v16)
            v19 := GenS2{F0: v1, F1: 836}
            _ = fi18
            _ = v19
        }
        if true {
            for fi20 := 0; fi20 < 7; fi20++ {
                v3[4294967295] = (((("hello" + "a") + ("" + "foo bar")) + "a") + "\n\t")
                for fi21 := 0; fi21 < 10; fi21++ {
                    v15 -= (fi20 << ((0) & 63))
                    _ = fi21
                }
                for fi22 := 0; fi22 < 3; fi22++ {
                    switch "a" {
                    case "\n\t":
                        if false {
                            v23 := GenS3{F0: -9223372036854775808, F1: 32767, F2: v17}
                            fmt.Println(v17, fi22, v15)
                            v24 := GenS1{F0: 0, F1: "a"}
                            _ = v23
                            _ = v24
                        }
                    case "a":
                        v1 = 0
                    case "foo bar":
                        _ = len(v16)
                    default:
                        v25 := GenS2{F0: v1, F1: 18446744073709551615}
                        _ = v25
                    }
                    fmt.Println(fi22, fi22)
                    var v26 float32 = (3.14159 + (float32(0.1) + ((0.0 + -2.25) + ((((1.5 - -2.25) / 1e10) / ((0.1 - 3.14159) / -2.25)) + float32((-2.25 - ((-2.25 - (1.5 * ((1e10 / (1.5 * 3.14159)) + 1.5))) * (float64(3.14159) + (1e10 / (1.5 - 0.0))))))))))
                    _ = fi22
                    _ = v26
                }
                _ = fi20
            }
            v27 := GenS2{F0: (v1 << ((2) & 63)), F1: 42}
            _ = v27
        } else {
            var v28 float64 = -2.25
            fmt.Println(v1, v17)
            switch "hello" {
            case "foo bar":
                v29 := GenS1{F0: 0, F1: ("\n\t" + "")}
                _ = v29
            default:
                v17 += v17
            }
            _ = v28
        }
    }
    var v30 uint32 = v17
    var v31 int = v15
    v30 += v30
    v17 = (v30 << ((973) & 63))
    for fi32 := 0; fi32 < 11; fi32++ {
        if true {
            v33 := GenS2{F0: v1, F1: 1}
            _ = v33
        }
        switch "\n\t" {
        case "foo bar":
            var v34 uint64 = 800
            _ = v34
        }
        _ = fi32
    }
    var v35 bool = (!(v17 <= 16))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
}
