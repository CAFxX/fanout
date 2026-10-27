package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint64 }
func genFn0(p0 int16, p1 int, p2 int) int {
    var v4 int32 = -125
    var v5 int = (p1 & int(v4))
    fmt.Println(p2, p0)
    if ("\n\t" != ("foo bar" + ("a" + ""))) {
        v4 -= 2147483647
        p0 += (p0 ^ (p0 >> ((2) & 63)))
        for fi6 := 0; fi6 < 20; fi6++ {
            p1 = 0
            fmt.Println(p1, p2, p2)
            if false {
                v7 := []int{v5}
                var v8 int = -727
                _ = v7
                _ = v8
            } else {
                v9 := []int{32767}
                switch "foo bar" {
                case "":
                    fmt.Println(fi6, v4)
                case "\n\t":
                    v4 = (v4 & -813)
                default:
                    if len(v9) > 0 {
                        v9[0] = (p2 + p2)
                    }
                }
                _ = v9
            }
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p1
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{844, 521, 2147483648}
    var v3 int32 = 1
    v10 := GenS1{F0: -257, F1: false}
    v1 = v1
    for fi11 := 0; fi11 < 20; fi11++ {
        if true {
            if false {
                _ = len(v2)
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 int64 = 1
                    fmt.Println(v1, fi11, v1)
                    _ = fi12
                    _ = v13
                }
                var v14 uint16 = 2
                _ = v14
            } else {
                v1 = v1
            }
        } else {
            fmt.Println(v3, v3)
            for fi15 := 0; fi15 < 7; fi15++ {
                if (1 > 100) {
                    v16 := GenS2{F0: v1}
                    _ = v16
                } else {
                    v3 += v3
                    v1 = ((v1 + "foo bar") + v1)
                    var v17 int = fi11
                    _ = v17
                }
                v18 := GenS2{F0: ""}
                v19 := GenS2{F0: (("hello" + (v1 + v1)) + "foo bar")}
                _ = fi15
                _ = v18
                _ = v19
            }
            fmt.Println(v3)
        }
        _ = fi11
    }
    if (0 == -1) {
        v20 := GenS3{F0: 2147483648}
        _ = v20
    }
    v3 -= v3
    if len(v2) > 0 {
        v2[0] = 1
    }
    v1 = (v1 + ("" + (("" + v1) + (v1 + v1))))
    v21 := GenS4{F0: 819, F1: 0}
    v22 := GenS1{F0: 0, F1: ((v3 >> ((255) & 63)) <= v3)}
    switch 4294967295 {
    case 4294967295:
        v23 := GenS1{F0: 32767, F1: (42 > 528)}
        _ = v23
    case 898:
        if (787 <= 0) {
            _ = len(v2)
            v1 = ("hello" + "\n\t")
            v3 = v3
        }
    default:
        var v24 uint16 = 42
        _ = v24
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
