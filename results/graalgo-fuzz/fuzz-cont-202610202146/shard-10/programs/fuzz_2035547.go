package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 string; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint64; F1 bool; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int64 }
func genFn0(p0 string, p1 bool, p2 int32) bool {
    if (("a" + "hello") > p0) {
        v5 := GenS1{F0: (("a" + "hello") + p0), F1: p0}
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 bool = true
            p1 = v7
            var v8 uint16 = 0
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        p2 -= (p2 + p2)
        switch (p0 + p0) {
        case "foo bar":
            fmt.Println(fi9, p2, p1)
        case "hello":
            v10 := []int{(fi9 - -2147483648), fi9}
            _ = v10
        default:
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := []int{(fi9 * 2147483647), (^(fi9))}
                _ = fi11
                _ = v12
            }
        }
        _ = fi9
    }
    p1 = (p1 || (!(4294967295 != 1)))
    switch -9223372036854775808 {
    case 2147483647:
        var v13 bool = (0 > 2147483647)
        _ = v13
    case 0:
        for fi14 := 0; fi14 < 14; fi14++ {
            p2 = p2
            if false {
                fmt.Println(fi14, fi14)
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 int8 = 42
                    v17 := GenS1{F0: p0, F1: "hello"}
                    var v18 uint32 = 154
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                fmt.Println(p1, p2, p1)
            }
            _ = fi14
        }
    case 42:
        for fi19 := 0; fi19 < 5; fi19++ {
            p1 = (!(483 != 42))
            _ = fi19
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return true
}
func genFn1() string {
    v20 := []int{-1, 127, 128}
    if true {
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 uint8 = 1
            switch fi21 {
            case -991:
                v23 := GenS2{F0: "hello", F1: -35, F2: 562}
                _ = v23
            case 261:
                switch -221 {
                case 4294967295:
                    var v24 uint8 = v22
                    _ = v24
                case 42:
                    v25 := GenS1{F0: "\n\t", F1: "foo bar"}
                    _ = v25
                default:
                    var v26 int64 = 0
                    _ = v26
                }
            case 127:
                var v27 string = ("a" + "")
                _ = v27
            default:
                v28 := GenS3{F0: 18446744073709551615, F1: (398 < 2), F2: 127}
                _ = v28
            }
            _ = fi21
            _ = v22
        }
    }
    _ = v20
    return ("a" + ("foo bar" + "a"))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    var v4 uint64 = 0
    switch "\n\t" {
    case "foo bar":
        for fi29 := 0; fi29 < 8; fi29++ {
            fmt.Println(fi29, v4)
            for fi30 := 0; fi30 < 1; fi30++ {
                v31 := GenS4{F0: v4, F1: 65535, F2: -9223372036854775808}
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
    case "a":
        v4 = 1
    default:
        switch -2147483648 {
        case 781:
            switch "a" {
            case "hello":
                v4 = v4
            }
        case 231:
            var v32 bool = true
            _ = v32
        case -2147483648:
            switch -9223372036854775808 {
            case -2147483648:
                v4 += (^(v4))
            }
        }
    }
    v4 += v4
    for fi33 := 0; fi33 < 11; fi33++ {
        v4 = (v4 * v4)
        _ = fi33
    }
    v4 -= v4
    for fi34 := 0; fi34 < 8; fi34++ {
        for fi35 := 0; fi35 < 5; fi35++ {
            v36 := []int{}
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    switch 2147483647 {
    case 852:
        v4 = v4
    case -9223372036854775808:
        v37 := GenS3{F0: (v4 ^ v4), F1: false, F2: 66}
        _ = v37
    }
    var v38 int = 32768
    fmt.Println(v4)
    v39 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
