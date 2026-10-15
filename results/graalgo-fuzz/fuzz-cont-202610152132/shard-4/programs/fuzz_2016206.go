package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint; F2 string }
func genFn0() int {
    var v7 uint8 = 42
    v8 := GenS1{F0: 42, F1: 1000, F2: 4294967295}
    for fi9 := 0; fi9 < 9; fi9++ {
        if ((fi9 ^ (fi9 / ((fi9) | 1))) >= 127) {
            v10 := []int{(fi9 % ((32767) | 1)), 4294967295}
            var v11 int32 = 42
            _ = v10
            _ = v11
        } else {
            v7 = (^(v7))
        }
        v12 := []int{(-(fi9)), 273, 177}
        _ = fi9
        _ = v12
    }
    _ = v7
    _ = v8
    return -865
}
func genFn1(p0 uint8, p1 uint) int8 {
    var v13 uint8 = 2
    v13 += 168
    switch "a" {
    case "hello":
        for fi14 := 0; fi14 < 20; fi14++ {
            if true {
                v15 := []int{fi14}
                _ = v15
            } else {
                for fi16 := 0; fi16 < 7; fi16++ {
                    v17 := GenS1{F0: 192, F1: 32767, F2: 777}
                    v13 += (p0 & v13)
                    _ = fi16
                    _ = v17
                }
                switch "\n\t" {
                case "a":
                    v18 := GenS1{F0: 0, F1: 1, F2: 42}
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 12; fi19++ {
                        var v20 uint8 = 255
                        v21 := GenS2{F0: 302, F1: (p1 ^ (p1 % ((p1) | 1))), F2: "hello"}
                        v13 = v13
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                v13 -= 202
            }
            _ = fi14
        }
    case "a":
        p0 -= v13
    default:
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 int32 = 1
            v13 += uint8(p1)
            var v24 int32 = (v23 * -1)
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 18; fi26++ {
            var v27 int8 = -128
            p1 = 1
            var v28 bool = (9223372036854775807 < 0)
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 uint32 = 661
        v13 = (p0 % ((v13) | 1))
        _ = fi25
        _ = v29
    }
    v30 := GenS2{F0: 65535, F1: 18446744073709551615, F2: ("a" + (("\n\t" + "") + "\n\t"))}
    _ = p0
    _ = p1
    _ = v13
    _ = v30
    return -128
}

func main() {
    var v1 int16 = 615
    v2 := []uint32{2147483648}
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 996
    var v5 bool = false
    v6 := []int{}
    fmt.Println(v1, v4)
    var v31 uint16 = 2
    var v32 int16 = v1
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v5, v31, v32)
        _ = fi33
    }
    if len(v2) > 0 {
        v2[0] = 2
    }
    if v5 {
        fmt.Println(v4, v1, v4)
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 string = ("a" + "foo bar")
            if len(v2) > 0 {
                v2[0] = 32
            }
            for fi36 := 0; fi36 < 15; fi36++ {
                v6 = append(v6, genFn0())
                _ = len(v3)
                v37 := GenS2{F0: v31, F1: 0, F2: ("a" + v35)}
                _ = fi36
                _ = v37
            }
            _ = fi34
            _ = v35
        }
    }
    var v38 int64 = -9223372036854775808
    var v39 int32 = 2147483647
    v39 -= (v39 >> ((662) & 63))
    switch "\n\t" {
    case "hello":
        _ = len(v3)
    case "foo bar":
        fmt.Println(v38, v31)
    }
    v40 := GenS2{F0: v31, F1: 255, F2: "hello"}
    var v41 int8 = 39
    switch "a" {
    case "\n\t":
        var v42 int16 = (v1 + v32)
        _ = v42
    case "foo bar":
        if true {
            v43 := GenS1{F0: 32767, F1: (v1 << ((65535) & 63)), F2: (v4 >> ((196) & 63))}
            _ = v43
        } else {
            v4 -= 1
            var v44 int8 = 127
            v39 = int32(v38)
            _ = v44
        }
    case "a":
        _ = len(v3)
    default:
        v5 = true
    }
    v4 -= v4
    fmt.Println(v5, v41)
    var v45 int8 = v41
    var v46 uint16 = v31
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
