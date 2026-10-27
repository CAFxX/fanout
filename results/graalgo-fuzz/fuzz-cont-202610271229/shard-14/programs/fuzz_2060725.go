package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 string; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int; F1 string; F2 int32 }
type GenS6 struct { F0 uint; F1 uint; F2 int32 }
func genFn0(p0 string, p1 int) uint32 {
    v7 := GenS1{F0: 9223372036854775807, F1: 0}
    p1 = p1
    p0 = ("\n\t" + "a")
    p1 = p1
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{}
        if (2 >= 133) {
            p0 = p0
            switch p0 {
            case "foo bar":
                for fi10 := 0; fi10 < 6; fi10++ {
                    v11 := GenS2{F0: ("" + "foo bar")}
                    v9 = append(v9, 2147483647)
                    _ = fi10
                    _ = v11
                }
            case "\n\t":
                var v12 string = (((("\n\t" + (((p0 + ("\n\t" + ("\n\t" + (("foo bar" + (p0 + p0)) + p0)))) + "\n\t") + "\n\t")) + ((p0 + (p0 + p0)) + p0)) + ("" + (("a" + ("foo bar" + ((((p0 + ("\n\t" + p0)) + "\n\t") + "") + "foo bar"))) + p0))) + "a")
                _ = v12
            case "":
                for fi13 := 0; fi13 < 13; fi13++ {
                    if true {
                        v14 := GenS1{F0: -9223372036854775808, F1: 42}
                        _ = v14
                    }
                    p1 += (fi8 >> ((4294967295) & 63))
                    if false {
                        switch (p0 + p0) {
                        case "\n\t":
                            p0 = p0
                        }
                        p0 = p0
                        p0 = p0
                    }
                    var v15 uint32 = 4294967295
                    switch (fi8 >> ((688) & 63)) {
                    case 4294967295:
                        var v16 float32 = float32(1e10)
                        _ = v16
                    default:
                        fmt.Println(fi13, fi8, p1)
                    }
                    _ = fi13
                    _ = v15
                }
            default:
                if len(v9) > 0 {
                    v9[0] = (fi8 % (((p1 + (fi8 ^ fi8))) | 1))
                }
            }
            switch fi8 {
            case 673:
                p0 = "a"
            case -1:
                if (!(1e10 < 1e10)) {
                    v17 := GenS3{F0: 1, F1: (680 > 358)}
                    var v18 bool = true
                    _ = v17
                    _ = v18
                } else {
                    if (1 == 0) {
                        v19 := GenS4{F0: p0, F1: 255, F2: 1}
                        _ = v19
                    } else {
                        p1 -= (-(p1))
                    }
                }
            case 42:
                fmt.Println(p0, p1)
            }
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]int)
    var v3 bool = false
    v4 := []int{42, 4294967295, -9223372036854775808, -2147483648}
    v5 := []string{"\n\t", "\n\t", "foo bar", (("" + "hello") + ("\n\t" + ((("hello" + (("foo bar" + "") + "hello")) + ("\n\t" + "\n\t")) + "foo bar")))}
    var v6 int32 = -285
    var v20 int32 = v6
    var v21 uint16 = 0
    v6 -= v20
    v20 = v6
    v1 += 907
    v1 = v1
    fmt.Println(v6, v6)
    v22 := GenS5{F0: 174, F1: ("foo bar" + (("" + "foo bar") + "foo bar")), F2: (^(v6))}
    v23 := GenS6{F0: 1, F1: 0, F2: v6}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
