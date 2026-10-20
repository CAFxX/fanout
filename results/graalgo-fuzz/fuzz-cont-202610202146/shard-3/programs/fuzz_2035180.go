package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 string }
type GenS2 struct { F0 int }
func genFn0(p0 uint, p1 bool, p2 uint64) int16 {
    if ((4294967295 != 197) && ((p2 & p2) >= p2)) {
        p2 -= uint64(p0)
        var v5 int16 = 1
        _ = v5
    }
    v6 := []int{1, 127, -9223372036854775808, 42}
    var v7 uint16 = 850
    if ((p1 && (!((1 > 1) && (228 != 253)))) && p1) {
        var v8 int8 = 1
        fmt.Println(p2)
        _ = v8
    } else {
        fmt.Println(p0, p2)
    }
    v9 := GenS1{F0: (("" + ("" + "\n\t")) + "\n\t"), F1: 32767, F2: (("hello" + ("foo bar" + "\n\t")) + "foo bar")}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v9
    return 42
}
func genFn1(p0 int16, p1 uint) bool {
    fmt.Println(p0, p0)
    p1 += p1
    fmt.Println(p0)
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 42
    var v3 int = (^(v2))
    v4 := []int{(v3 << ((197) & 63)), (v2 | 128), (v3 * 127), (v2 | v3)}
    v2 += (v3 | (v3 * (v2 - v3)))
    v3 += (v2 % ((v2) | 1))
    switch "hello" {
    case "\n\t":
        switch ((("foo bar" + "\n\t") + "hello") + "\n\t") {
        case "foo bar":
            for fi10 := 0; fi10 < 20; fi10++ {
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 int32 = -2147483648
                    v2 = fi10
                    _ = fi11
                    _ = v12
                }
                v3 = fi10
                for fi13 := 0; fi13 < 2; fi13++ {
                    var v14 uint16 = 32767
                    v14 = v14
                    _ = len(v1)
                    switch fi10 {
                    case -192:
                        v3 = (v2 + 32)
                    case 32767:
                        var v15 float32 = 0.1
                        _ = v15
                    default:
                        var v16 string = ""
                        _ = v16
                    }
                    if (32767 <= 1000) {
                        var v17 uint = 1
                        v14 = 65535
                        v14 = 0
                        _ = v17
                    } else {
                        _ = len(v1)
                    }
                    _ = fi13
                    _ = v14
                }
                _ = fi10
            }
        }
    }
    var v18 uint8 = 0
    var v19 int8 = 0
    if (-2147483648 > -1) {
        v20 := GenS1{F0: "a", F1: v3, F2: "hello"}
        _ = v20
    }
    v21 := v1[4294967295]
    v4 = append(v4, -1)
    v22 := GenS2{F0: (-(v2))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
