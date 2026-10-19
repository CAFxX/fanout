package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 string }
type GenS2 struct { F0 int32; F1 int32; F2 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 bool; F2 uint }
func genFn0(p0 int16, p1 uint64) int16 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint64 = p1
        var v7 uint = 718
        var v8 uint = (v7 % (((v7 % (((v7 ^ (^(v7)))) | 1))) | 1))
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := []int{fi9, (fi9 >> ((2) & 63)), -408, fi9}
        if true {
            if ((!true) || true) {
                _ = len(v10)
            } else {
                if len(v10) > 0 {
                    v10[0] = 32768
                }
                if true {
                    if false {
                        p0 -= (-(p0))
                    }
                } else {
                    p0 = p0
                    v11 := GenS1{F0: ((0.0 + float32(float64(-2.25))) != -2.25), F1: p0, F2: (("hello" + "hello") + "hello")}
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v13 := GenS2{F0: 554, F1: -1, F2: 42}
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                }
            }
            fmt.Println(p1, fi9)
            v14 := GenS3{F0: 65535}
            _ = v14
        } else {
            p1 = (p1 ^ 2)
        }
        _ = fi9
        _ = v10
    }
    p0 -= (p0 + (p0 << ((515) & 63)))
    _ = p0
    _ = p1
    return 32767
}
func genFn1(p0 int, p1 int64, p2 int16) uint32 {
    var v15 bool = false
    v15 = ((p0 >> ((42) & 63)) <= -2147483648)
    v16 := GenS1{F0: v15, F1: (p2 ^ genFn0((p2 + 514), 2)), F2: "\n\t"}
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 1
}

func main() {
    v1 := []uint32{2147483647, 1, 65535, 2}
    v2 := make(map[string]int)
    v3 := []int{}
    var v4 uint64 = 18446744073709551615
    v2[("" + "foo bar")] = 32768
    fmt.Println(v4)
    v17 := GenS4{F0: 0, F1: (18446744073709551615 <= 2), F2: 17}
    var v18 uint16 = 1
    v18 += (v18 & (^(v18)))
    delete(v2, "")
    var v19 int64 = 0
    fmt.Println(v4)
    for fi20 := 0; fi20 < 4; fi20++ {
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 float32 = 0.1
            var v23 string = ("a" + ((("foo bar" + "foo bar") + ("" + "\n\t")) + ("a" + (("a" + ("\n\t" + "foo bar")) + "foo bar"))))
            v22 -= ((((float32((0.1 + float64(1.5))) * v22) * (v22 + v22)) + float32(-2.25)) - 0.0)
            _ = fi21
            _ = v22
            _ = v23
        }
        v24 := GenS1{F0: false, F1: 0, F2: ("" + "foo bar")}
        v18 = (v18 % ((v18) | 1))
        _ = fi20
        _ = v24
    }
    v18 = v18
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
