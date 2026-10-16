package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint32 }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 int32 = -472
        _ = fi5
        _ = v6
    }
    switch 9223372036854775807 {
    case 128:
        var v7 uint8 = 42
        _ = v7
    case -138:
        var v8 uint8 = 201
        _ = v8
    }
    if ((((0.0 + (float32((float64(-2.25) - (1e10 * (0.0 - 1e10)))) / 1e10)) - 1e10) - (1e10 + (0.0 + float32(3.14159)))) > 1.5) {
        var v9 uint16 = 65535
        if true {
            var v10 int32 = 42
            if (("\n\t" != ("foo bar" + "")) || false) {
                var v11 uint = 1
                var v12 uint8 = 23
                for fi13 := 0; fi13 < 15; fi13++ {
                    if true {
                        if true {
                            v14 := []int{fi13, fi13, fi13, (^(fi13))}
                            var v15 uint64 = 465
                            _ = v14
                            _ = v15
                        }
                        v12 = 42
                        v16 := []int{128, (-(fi13))}
                        _ = v16
                    }
                    v10 += v10
                    _ = fi13
                }
                _ = v11
                _ = v12
            } else {
                for fi17 := 0; fi17 < 8; fi17++ {
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        v9 -= v9
                    }
                    _ = fi17
                }
                v9 -= (v9 - (v9 ^ (v9 & 1)))
                if true {
                    var v18 bool = true
                    var v19 int8 = 104
                    _ = v18
                    _ = v19
                }
            }
            v20 := GenS1{F0: 231}
            _ = v10
            _ = v20
        } else {
            v9 = 32767
        }
        _ = v9
    } else {
        for fi21 := 0; fi21 < 11; fi21++ {
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 float64 = 1.5
                v24 := GenS2{F0: 745, F1: 4294967295}
                _ = fi22
                _ = v23
                _ = v24
            }
            var v25 string = "\n\t"
            _ = fi21
            _ = v25
        }
    }
    return 2
}
func genFn1(p0 int64) uint64 {
    var v26 string = "\n\t"
    for fi27 := 0; fi27 < 3; fi27++ {
        fmt.Println(v26, fi27, p0)
        _ = fi27
    }
    v26 = (v26 + v26)
    switch -682 {
    case 4294967295:
        v26 = v26
    case 2147483647:
        if false {
            v28 := []int{}
            if true {
                _ = len(v28)
            }
            _ = v28
        }
    }
    _ = p0
    _ = v26
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"hello", "foo bar", "foo bar"}
    v3 := make(map[int]uint32)
    var v4 uint16 = 2
    var v29 int8 = -1
    v30 := GenS3{F0: 1}
    v4 += v4
    v29 -= v29
    v4 = (^(v4))
    fmt.Println(v29, v4)
    v29 += (-(v29))
    switch ((("\n\t" + "\n\t") + "foo bar") + "foo bar") {
    case "\n\t":
        fmt.Println(v4)
    case "foo bar":
        for fi31 := 0; fi31 < 12; fi31++ {
            fmt.Println(fi31, fi31, fi31)
            var v32 int32 = 312
            if (!true) {
                v32 -= v32
                var v33 int64 = 1
                _ = v33
            } else {
                var v34 float32 = 0.0
                v35 := GenS1{F0: 42}
                _ = v34
                _ = v35
            }
            _ = fi31
            _ = v32
        }
    case "hello":
        v29 = 100
    }
    var v36 uint16 = v4
    v37 := GenS3{F0: 65535}
    if (false && false) {
        _ = len(v2)
    }
    v38 := GenS3{F0: 0}
    v39 := GenS1{F0: 1}
    fmt.Println(v4, v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
