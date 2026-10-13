package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int32 }
func genFn0(p0 uint32) int16 {
    v7 := GenS1{F0: 231, F1: 255, F2: 42}
    fmt.Println(p0)
    v8 := []int{}
    _ = p0
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 int64, p1 uint16) uint64 {
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(fi9)
        fmt.Println(p0, fi9, p1)
        fmt.Println(p1)
        _ = fi9
    }
    var v10 uint64 = 2
    switch "hello" {
    case "\n\t":
        v10 -= v10
    case "hello":
        switch "hello" {
        case "a":
            switch 32767 {
            case -218:
                if true {
                    p0 = 9223372036854775807
                    var v11 string = ""
                    v12 := []int{705}
                    _ = v11
                    _ = v12
                }
            }
        case "":
            var v13 uint64 = v10
            _ = v13
        default:
            v14 := []int{4294967295, -649, -846}
            _ = v14
        }
    case "a":
        var v15 int32 = 1
        _ = v15
    }
    var v16 string = ("\n\t" + ("hello" + ("a" + (("foo bar" + "a") + "hello"))))
    _ = p0
    _ = p1
    _ = v10
    _ = v16
    return (v10 - (v10 | uint64(p1)))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 int = 127
    var v6 string = ("" + "\n\t")
    fmt.Println(v6, v6)
    var v17 uint = 2
    if false {
        fmt.Println(v5, v5)
        var v18 uint64 = 0
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS1{F0: 1, F1: 1, F2: -934}
        v6 = v6
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(v17, v5)
        v22 := GenS1{F0: 127, F1: 127, F2: -643}
        _ = fi21
        _ = v22
    }
    v23 := []int{32767}
    if true {
        if (9223372036854775807 < -1) {
            if (!(162 == 2147483647)) {
                fmt.Println(v5, v5)
            }
        }
    }
    var v24 float64 = (float64(3.14159) * ((1.5 * (-2.25 + -2.25)) - 1e10))
    switch (v5 ^ v5) {
    case 896:
        v25 := GenS1{F0: 0, F1: 42, F2: -562}
        _ = v25
    case -36:
        delete(v3, "")
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        if len(v23) > 0 {
            v23[0] = -2147483648
        }
        if ((255 < 686) || (true && true)) {
            v17 = 2
        }
        v27 := v2[(("" + (v6 + v6)) + (v6 + "foo bar"))]
        _ = fi26
        _ = v27
    }
    if len(v23) > 0 {
        v23[0] = v5
    }
    var v28 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
