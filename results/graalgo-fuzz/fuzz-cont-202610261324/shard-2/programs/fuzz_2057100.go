package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int; F1 int }
func genFn0() uint64 {
    v6 := []int{-193, 0, 9223372036854775807, 127}
    if true {
        var v7 float64 = (-2.25 / 0.1)
        v7 = ((v7 / (v7 / (v7 + v7))) - v7)
        _ = v7
    }
    v8 := GenS1{F0: 1000, F1: 42}
    _ = len(v6)
    _ = v6
    _ = v8
    return 0
}
func genFn1(p0 int32) uint64 {
    fmt.Println(p0)
    var v9 int64 = 42
    _ = p0
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    var v3 uint = 328
    var v4 int = 32767
    var v5 uint64 = 0
    if ((v4 & v4) != (v4 / ((9223372036854775807) | 1))) {
        var v10 int64 = -1
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 int64 = (v10 * (v10 & -312))
            if true {
                v12 = int64(v5)
                fmt.Println(v5)
                switch "foo bar" {
                case "hello":
                    switch ((("\n\t" + ((("a" + ("hello" + ("foo bar" + "hello"))) + ("foo bar" + "a")) + "")) + "\n\t") + (("a" + ("a" + "a")) + "a")) {
                    case "a":
                        for fi13 := 0; fi13 < 14; fi13++ {
                            fmt.Println(v12, v3, fi11)
                            switch -1 {
                            case -778:
                                v14 := []int{fi13, (^(v4)), -1000, 4294967295}
                                _ = v14
                            case 127:
                                var v15 uint16 = 740
                                _ = v15
                            case 1:
                                switch v4 {
                                case 42:
                                    v1[891] = (v4 ^ (fi13 & (v4 << ((42) & 63))))
                                }
                            }
                            _ = fi13
                        }
                    case "foo bar":
                        fmt.Println(v4, v10)
                    case "hello":
                        delete(v1, uint32(v5))
                    }
                default:
                    v16 := GenS1{F0: 1, F1: v3}
                    _ = v16
                }
            }
            _ = fi11
            _ = v12
        }
        var v17 string = ("" + ("\n\t" + "a"))
        _ = v10
        _ = v17
    }
    var v18 float64 = 0.1
    var v19 float32 = 0.0
    if false {
        fmt.Println(v18, v4, v3)
        v3 = (v3 + 1)
        v19 -= 3.14159
    } else {
        v5 += (v5 >> (((^(v3))) & 63))
        var v20 int16 = 32767
        _ = v20
    }
    v18 += ((((v18 + (v18 - ((v18 - v18) / v18))) / -2.25) + 0.0) / v18)
    if (!true) {
        var v21 int = v4
        fmt.Println(v3, v4, v4)
        _ = v21
    }
    fmt.Println(v4, v19, v19)
    if (!true) {
        var v22 int64 = -378
        var v23 int16 = -276
        _ = v22
        _ = v23
    } else {
        switch v4 {
        case 266:
            if (686 > 2147483647) {
                v24 := GenS2{F0: v3, F1: -437}
                _ = v24
            } else {
                v4 -= (v4 / (((^(v4))) | 1))
            }
        case 32767:
            v4 += v4
        case -9223372036854775808:
            for fi25 := 0; fi25 < 11; fi25++ {
                v26 := []int{(v4 * v4), 746}
                _ = fi25
                _ = v26
            }
        }
        v4 += (v4 ^ v4)
        for fi27 := 0; fi27 < 16; fi27++ {
            for fi28 := 0; fi28 < 14; fi28++ {
                var v29 uint64 = (v5 / (((v5 ^ 4294967295)) | 1))
                v5 = 0
                v30 := v1[820]
                _ = fi28
                _ = v29
                _ = v30
            }
            v31 := []int{v4}
            switch fi27 {
            case 128:
                var v32 int = v4
                _ = v32
            case 32767:
                if ((v3 >= v3) || (!(0.0 > 1e10))) {
                    if true {
                        var v33 int32 = 1
                        fmt.Println(v5, v4, v33)
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        _ = v33
                    }
                    if len(v31) > 0 {
                        v31[0] = (v4 >> ((2) & 63))
                    }
                }
            default:
                var v34 string = (("hello" + ("foo bar" + "foo bar")) + "\n\t")
                _ = v34
            }
            _ = fi27
            _ = v31
        }
    }
    v35 := []int{4294967295, 904, v4, (v4 | 9223372036854775807)}
    v36 := GenS3{F0: (^(v4)), F1: (-(v4))}
    for fi37 := 0; fi37 < 7; fi37++ {
        v38 := GenS3{F0: (fi37 | v4), F1: fi37}
        _ = fi37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
