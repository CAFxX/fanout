package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint32 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0() uint32 {
    switch -242 {
    case 128:
        v7 := []int{-684}
        _ = v7
    case -678:
        var v8 uint32 = 0
        _ = v8
    }
    v9 := GenS1{F0: 65535, F1: 127, F2: 42}
    v10 := []int{-458, 0, -221}
    var v11 uint64 = 4294967295
    _ = v10
    _ = v11
    _ = v9
    return 2147483648
}
func genFn1(p0 int16) int64 {
    p0 = 42
    p0 += p0
    p0 += p0
    v12 := GenS1{F0: 873, F1: 0, F2: 943}
    switch -2147483648 {
    case -853:
        p0 -= 1000
    case 9223372036854775807:
        p0 += (-(p0))
    default:
        p0 = (p0 & -1)
    }
    _ = p0
    _ = v12
    return 183
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"", "hello", "a"}
    v3 := make(map[uint32]uint32)
    var v4 int32 = 615
    var v5 int16 = 830
    v6 := []string{"hello", "a", "hello"}
    if ((v5 + (v5 / (((^(v5))) | 1))) < (v5 << ((65535) & 63))) {
        var v13 int32 = -2147483648
        for fi14 := 0; fi14 < 20; fi14++ {
            _ = len(v2)
            v15 := GenS2{F0: "a", F1: 42}
            _ = fi14
            _ = v15
        }
        switch 9223372036854775807 {
        case 0:
            if false {
                v4 = (v13 << ((83) & 63))
            }
        case -51:
            v16 := GenS1{F0: 65535, F1: 0, F2: 55}
            _ = v16
        case 115:
            v13 -= (v4 + (v4 | v13))
        default:
            var v17 uint16 = 65535
            _ = v17
        }
        _ = v13
    }
    var v18 string = (((("a" + (("\n\t" + (("foo bar" + "\n\t") + "hello")) + "\n\t")) + "") + "hello") + "a")
    var v19 int32 = v4
    if ((!(int(v4) == -9223372036854775808)) && true) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        fmt.Println(v4)
        v20 := GenS3{F0: 368, F1: 32767}
        _ = v20
    }
    var v21 int32 = -997
    var v22 int8 = 77
    var v23 uint16 = 65535
    v24 := GenS1{F0: 42, F1: v22, F2: 2147483647}
    _ = len(v6)
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 uint64 = 307
        if false {
            switch v18 {
            case "":
                if true {
                    if (!(280 >= 2)) {
                        _ = len(v2)
                    } else {
                        var v27 bool = (false || (((0.1 / ((((3.14159 / 0.1) * 1e10) - (3.14159 / ((0.0 - 1.5) * float64(1e10)))) - float64(0.1))) - 1e10) > (float64(1.5) * -2.25)))
                        v5 = (v5 & (v5 % ((v5) | 1)))
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 13; fi28++ {
                        fmt.Println(fi28)
                        _ = len(v2)
                        if false {
                            fmt.Println(fi25)
                            var v29 int8 = (v22 + v22)
                            delete(v3, 65535)
                            var v30 uint = 18446744073709551615
                            _ = v29
                            _ = v30
                        } else {
                            if true {
                                fmt.Println(v18, v21)
                            }
                            fmt.Println(v23, fi28, v18)
                            var v31 float64 = ((1e10 * -2.25) + -2.25)
                            _ = v31
                        }
                        _ = fi28
                    }
                    v6 = append(v6, "a")
                } else {
                    fmt.Println(v19, v23, v26)
                    var v32 string = v18
                    v26 += 42
                    _ = v32
                }
            case "hello":
                v3[65535] = 454
            }
        } else {
            var v33 uint = 4294967295
            _ = v33
        }
        for fi34 := 0; fi34 < 11; fi34++ {
            var v35 string = ("hello" + "")
            var v36 int8 = 100
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi25
        _ = v26
    }
    if (v22 <= (v22 / ((v22) | 1))) {
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 float64 = 0.0
            v38 = -2.25
            _ = fi37
            _ = v38
        }
    }
    var v39 float32 = float32((-2.25 / -2.25))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v18 = (v18 + v18)
    var v40 int32 = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
