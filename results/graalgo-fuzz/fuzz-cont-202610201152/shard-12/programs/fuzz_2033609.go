package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 uint16 }
func genFn0() int16 {
    if true {
        var v5 bool = true
        v5 = ((v5 && v5) || true)
        fmt.Println(v5)
        _ = v5
    } else {
        var v6 uint64 = 792
        var v7 string = "foo bar"
        var v8 int64 = 829
        _ = v6
        _ = v7
        _ = v8
    }
    if true {
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 9; fi10++ {
                switch fi9 {
                case 128:
                    v11 := GenS1{F0: 952}
                    _ = v11
                }
                v12 := []int{(-(fi10)), (fi10 >> ((2) & 63)), (fi10 - 4294967295), fi9}
                _ = fi10
                _ = v12
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := GenS2{F0: 711, F1: 464, F2: 0}
                var v15 int16 = -1
                var v16 uint32 = 42
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            switch -699 {
            case 102:
                v17 := GenS3{F0: 2, F1: 25, F2: 9223372036854775807}
                _ = v17
            case 2147483647:
                var v18 int8 = -99
                _ = v18
            case 4294967295:
                var v19 bool = (2 >= 910)
                _ = v19
            default:
                var v20 int64 = -702
                _ = v20
            }
            _ = fi9
        }
        v21 := GenS4{F0: -957}
        v22 := GenS1{F0: 1}
        _ = v21
        _ = v22
    }
    switch 1 {
    case 673:
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 int64 = -1
            v25 := []int{}
            for fi26 := 0; fi26 < 6; fi26++ {
                switch "\n\t" {
                case "hello":
                    switch "foo bar" {
                    case "a":
                        switch "hello" {
                        case "\n\t":
                            var v27 int16 = 42
                            _ = v27
                        case "foo bar":
                            var v28 bool = (1 < 834)
                            _ = v28
                        }
                    }
                default:
                    fmt.Println(fi23, v24, v24)
                }
                _ = fi26
            }
            _ = fi23
            _ = v24
            _ = v25
        }
    case 128:
        if true {
            switch "\n\t" {
            case "":
                var v29 string = ("\n\t" + "a")
                _ = v29
            case "hello":
                v30 := GenS4{F0: 1000}
                _ = v30
            }
        }
    }
    switch "\n\t" {
    case "foo bar":
        v31 := GenS5{F0: -1, F1: "\n\t"}
        _ = v31
    default:
        v32 := []int{-9223372036854775808, 0, 13}
        _ = v32
    }
    return 1
}
func genFn1(p0 uint64, p1 uint) bool {
    var v33 int64 = -9223372036854775808
    fmt.Println(p1, p1)
    fmt.Println(p1)
    v34 := []int{42, -2147483648}
    var v35 int32 = 1
    _ = p0
    _ = p1
    _ = v33
    _ = v34
    _ = v35
    return (!(!true))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    var v4 uint8 = 225
    fmt.Println(v4)
    if false {
        var v36 int = 32767
        v37 := GenS4{F0: 42}
        _ = v36
        _ = v37
    } else {
        v38 := GenS6{F0: 0}
        v4 -= 233
        v39 := v3[4294967295]
        _ = v38
        _ = v39
    }
    var v40 uint = 42
    var v41 int64 = 763
    v42 := GenS4{F0: 32767}
    v43 := GenS6{F0: 1}
    switch "a" {
    case "foo bar":
        var v44 uint = (v40 % (((v40 >> (((v40 + v40)) & 63))) | 1))
        _ = v44
    default:
        v41 -= (v41 * 9223372036854775807)
    }
    v45 := GenS7{F0: -799, F1: 65535}
    delete(v1, ("\n\t" + ""))
    v41 -= -608
    if true {
        v41 = v41
    }
    var v46 int = 4294967295
    fmt.Println(v4, v41, v40)
    var v47 float64 = 0.1
    if (!(false || true)) {
        fmt.Println(v41, v47)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v45
    _ = v46
    _ = v47
}
