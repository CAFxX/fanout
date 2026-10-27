package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int64; F1 int32; F2 int8 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint32, p1 int64) int64 {
    if true {
        v6 := GenS1{F0: -1, F1: "a"}
        v7 := GenS2{F0: -1, F1: 1, F2: 73}
        p0 -= 42
        _ = v6
        _ = v7
    } else {
        var v8 string = ("hello" + ("foo bar" + "a"))
        v9 := GenS2{F0: p1, F1: 666, F2: 100}
        _ = v8
        _ = v9
    }
    v10 := []int{-397}
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(fi11, fi11)
        _ = fi11
    }
    v12 := GenS2{F0: (-(p1)), F1: 42, F2: -102}
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    return (^(p1))
}
func genFn1(p0 int16, p1 int64) uint64 {
    p0 -= p0
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := []int{fi13, fi13, (fi13 * 4294967295)}
        var v15 uint32 = 1
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    v1 := []uint32{2, 709}
    var v2 uint64 = 196
    var v3 string = "hello"
    var v4 float64 = 3.14159
    v5 := make(map[int]uint32)
    switch 128 {
    case 2147483647:
        var v16 uint8 = 255
        _ = v16
    default:
        if false {
            v3 = ((((v3 + "foo bar") + v3) + ("a" + v3)) + "a")
            fmt.Println(v4, v4)
            if (!(3.14159 > 0.1)) {
                v2 += 18446744073709551615
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 uint32 = 4294967295
                    v19 := GenS1{F0: 2147483647, F1: (v3 + (("hello" + (("hello" + ("" + v3)) + "\n\t")) + v3))}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            } else {
                v4 = float64(float32(v4))
            }
        }
    }
    _ = len(v5)
    var v20 uint16 = 255
    if false {
        v5[42] = 42
    } else {
        v20 += uint16(v2)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v21 int = 127
        _ = v21
    }
    if (!((1e10 / 1.0) < 0.0)) {
        fmt.Println(v3)
        v5[2147483647] = 2147483647
    } else {
        for fi22 := 0; fi22 < 15; fi22++ {
            v23 := GenS1{F0: 977, F1: (("" + ("a" + (v3 + "\n\t"))) + (v3 + "a"))}
            for fi24 := 0; fi24 < 1; fi24++ {
                if false {
                    if true {
                        v1 = append(v1, 0)
                        v20 = v20
                        var v25 string = "foo bar"
                        _ = v25
                    }
                    v1 = append(v1, 42)
                } else {
                    if false {
                        _ = len(v1)
                    }
                    switch (^(fi24)) {
                    case -2147483648:
                        var v26 float32 = (1.5 / 1.0)
                        _ = v26
                    case 70:
                        v3 = v3
                    case -898:
                        v3 = "a"
                    }
                    v27 := GenS3{F0: 823, F1: -1, F2: 4294967295}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v28 := GenS1{F0: 42, F1: v3}
                    _ = v27
                    _ = v28
                }
                v2 = v2
                var v29 uint16 = v20
                fmt.Println(v20)
                _ = fi24
                _ = v29
            }
            v3 = ("hello" + ((((v3 + "\n\t") + v3) + "foo bar") + ((v3 + v3) + ("hello" + ("" + (v3 + v3))))))
            _ = fi22
            _ = v23
        }
    }
    var v30 uint32 = 2
    v31 := GenS4{F0: "a"}
    v32 := GenS5{F0: 847}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
