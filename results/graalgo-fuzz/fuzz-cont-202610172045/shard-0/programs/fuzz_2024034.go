package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int16; F2 int64 }
type GenS4 struct { F0 string; F1 int32; F2 int8 }
func genFn0(p0 int16, p1 uint16) int {
    fmt.Println(p1, p1)
    v7 := []int{32767}
    p1 = p1
    if true {
        for fi8 := 0; fi8 < 19; fi8++ {
            for fi9 := 0; fi9 < 14; fi9++ {
                fmt.Println(fi9, p0)
                var v10 int8 = -75
                _ = fi9
                _ = v10
            }
            fmt.Println(p0)
            p0 = (p0 * (p0 ^ p0))
            _ = fi8
        }
    } else {
        var v11 int = 1
        v11 = 9223372036854775807
        v12 := GenS1{F0: p0, F1: 42}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    return 128
}

func main() {
    var v1 uint64 = 837
    var v2 uint64 = v1
    v3 := []string{"\n\t", "a", "hello", ("a" + (("\n\t" + "") + ("foo bar" + ("\n\t" + ((("hello" + "") + (("hello" + (("a" + "") + "hello")) + "a")) + "\n\t")))))}
    var v4 uint = 65535
    v5 := []uint32{994}
    v6 := make(map[uint32]string)
    switch ("" + ("\n\t" + "\n\t")) {
    case "a":
        switch -983 {
        case 941:
            for fi13 := 0; fi13 < 17; fi13++ {
                for fi14 := 0; fi14 < 19; fi14++ {
                    delete(v6, uint32(v4))
                    _ = fi14
                }
                _ = fi13
            }
        case -9223372036854775808:
            var v15 float32 = ((3.14159 * (0.1 + float32((0.1 + (float64(1.5) / (-2.25 - 1e10)))))) + (((0.1 / 1.0) / 1.0) * (float32((0.1 - float64((0.1 + -2.25)))) + 1.5)))
            _ = v15
        case 32768:
            v1 = (^(v2))
        }
    case "\n\t":
        _ = len(v6)
    case "":
        var v16 float32 = ((1.5 - 1e10) / 1.5)
        _ = v16
    default:
        var v17 string = (("hello" + ("a" + "foo bar")) + ("a" + "\n\t"))
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        switch (-(fi18)) {
        case 127:
            switch (fi18 + genFn0(0, 2)) {
            case 127:
                fmt.Println(v1)
            case 1:
                if (1000 == 984) {
                    switch fi18 {
                    case 127:
                        v2 = v2
                    case -2147483648:
                        if true {
                            v19 := GenS1{F0: -1, F1: 9223372036854775807}
                            v1 -= uint64(v4)
                            _ = v19
                        } else {
                            v20 := GenS1{F0: 12, F1: -809}
                            for fi21 := 0; fi21 < 5; fi21++ {
                                v22 := GenS2{F0: 2147483647}
                                v23 := GenS3{F0: false, F1: -32768, F2: 42}
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                            delete(v6, 2)
                            _ = v20
                        }
                    }
                    var v24 float64 = -2.25
                    _ = v24
                }
            default:
                v25 := GenS2{F0: 2147483647}
                _ = v25
            }
        case 285:
            delete(v6, 2147483648)
        case -1:
            v4 = v4
        }
        v2 += (v1 % (((v1 - (v2 / ((2) | 1)))) | 1))
        v1 = v2
        _ = fi18
    }
    v4 = (v4 >> (((v4 + (v4 % ((v4) | 1)))) & 63))
    var v26 float32 = 1.5
    if (2147483648 <= 577) {
        switch 127 {
        case 42:
            var v27 uint16 = 1
            _ = v27
        default:
            v2 = v2
        }
        v28 := GenS4{F0: "", F1: -1, F2: 0}
        _ = v28
    }
    var v29 int8 = 127
    var v30 uint64 = v1
    for fi31 := 0; fi31 < 2; fi31++ {
        v29 = v29
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 bool = ((^(v4)) >= v4)
            _ = fi32
            _ = v33
        }
        var v34 float64 = 1.5
        _ = fi31
        _ = v34
    }
    var v35 int16 = -32768
    if (!(!true)) {
        var v36 uint16 = 912
        for fi37 := 0; fi37 < 14; fi37++ {
            var v38 float64 = (3.14159 * (1.5 + (0.0 * -2.25)))
            v30 = 0
            fmt.Println(v4, v4)
            _ = fi37
            _ = v38
        }
        _ = v36
    }
    v39 := GenS3{F0: false, F1: v35, F2: 1}
    v29 = v29
    _ = len(v6)
    v5 = append(v5, 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
