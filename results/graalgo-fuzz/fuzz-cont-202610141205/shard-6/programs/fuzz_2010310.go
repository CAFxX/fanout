package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int; F2 string }
type GenS3 struct { F0 uint8; F1 string; F2 string }
func genFn0(p0 bool) int16 {
    for fi7 := 0; fi7 < 14; fi7++ {
        p0 = true
        var v8 uint32 = 1
        var v9 int16 = -1
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    v10 := GenS1{F0: 4294967295}
    if (true && p0) {
        if (42 > 2) {
            v11 := []int{-1, -9223372036854775808, -69}
            switch -347 {
            case -1:
                p0 = p0
            case 374:
                p0 = true
            }
            var v12 bool = p0
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := []int{fi13, fi13, fi13, 0}
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = v10
    return 0
}
func genFn1(p0 int32, p1 string, p2 uint) uint64 {
    p1 = p1
    p2 = 2
    _ = p0
    _ = p1
    _ = p2
    return 752
}

func main() {
    v1 := []uint32{}
    var v2 bool = (128 <= 127)
    var v3 int8 = 42
    var v4 int16 = 396
    var v5 int8 = v3
    var v6 uint = 929
    v15 := GenS1{F0: 1}
    for fi16 := 0; fi16 < 2; fi16++ {
        v2 = (!v2)
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 string = "\n\t"
            var v19 uint8 = 1
            v18 = "hello"
            _ = fi17
            _ = v18
            _ = v19
        }
        if false {
            v1 = append(v1, 881)
        }
        _ = fi16
    }
    v20 := GenS2{F0: -1, F1: 42, F2: "hello"}
    var v21 uint32 = 4294967295
    for fi22 := 0; fi22 < 6; fi22++ {
        var v23 float32 = 1.5
        fmt.Println(v2, v2, fi22)
        _ = fi22
        _ = v23
    }
    var v24 bool = (v2 || v2)
    _ = len(v1)
    var v25 uint8 = 78
    var v26 uint32 = (v21 / ((v21) | 1))
    switch ("\n\t" + ((("\n\t" + "foo bar") + (("a" + "hello") + "hello")) + "\n\t")) {
    case "hello":
        v24 = (("foo bar" + ("foo bar" + "foo bar")) < "\n\t")
    case "":
        v27 := GenS3{F0: (v25 & (^(v25))), F1: "a", F2: "hello"}
        _ = v27
    default:
        v28 := GenS1{F0: 42}
        _ = v28
    }
    fmt.Println(v21, v6)
    v4 -= v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
