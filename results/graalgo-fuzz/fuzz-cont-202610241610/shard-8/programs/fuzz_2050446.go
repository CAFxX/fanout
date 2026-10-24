package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint; F1 int32; F2 bool }
type GenS4 struct { F0 string }
func genFn0(p0 string, p1 int, p2 uint8) int8 {
    if false {
        switch p0 {
        case "":
            var v4 int8 = 100
            _ = v4
        case "foo bar":
            p2 = (^(p2))
        default:
            v5 := GenS1{F0: 787}
            _ = v5
        }
    }
    switch -1 {
    case -1:
        if (false || true) {
            p1 += (-(p1))
            var v6 int16 = -32768
            _ = v6
        }
    default:
        p0 = p0
    }
    var v7 uint16 = 0
    fmt.Println(p0, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 100
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    v8 := GenS1{F0: 255}
    var v9 int8 = -128
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS2{F0: 543, F1: 255, F2: 4294967295}
        var v12 int32 = 1
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint = 874
    if (false || (263 == 65535)) {
        fmt.Println(v13, v9)
        var v14 uint16 = 0
        _ = v14
    }
    v15 := GenS1{F0: 828}
    delete(v3, 42)
    if ((false || (1 > 2)) && true) {
        if true {
            if ((4294967295 >= 2) && false) {
                var v16 string = ""
                _ = v16
            } else {
                for fi17 := 0; fi17 < 20; fi17++ {
                    v18 := GenS1{F0: 42}
                    v9 -= v9
                    _ = fi17
                    _ = v18
                }
            }
        } else {
            if false {
                v19 := []int{128, -209, 1}
                _ = v19
            } else {
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 int16 = 0
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 int64 = 658
                        v9 -= v9
                        _ = fi22
                        _ = v23
                    }
                    _ = fi20
                    _ = v21
                }
                for fi24 := 0; fi24 < 5; fi24++ {
                    v25 := GenS2{F0: 1, F1: 253, F2: 1}
                    v26 := []int{(fi24 / ((fi24) | 1)), (fi24 / ((-539) | 1)), (fi24 >> (((v13 - (v13 + (v13 | v13)))) & 63))}
                    switch "foo bar" {
                    case "hello":
                        v13 = v13
                    default:
                        fmt.Println(fi24, fi24, fi24)
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                var v27 int32 = 74
                _ = v27
            }
        }
        _ = len(v2)
        fmt.Println(v9)
    } else {
        fmt.Println(v9)
    }
    v9 += 100
    for fi28 := 0; fi28 < 18; fi28++ {
        v29 := []int{fi28, (fi28 ^ 4294967295), (fi28 << (((v13 & v13)) & 63)), (fi28 - fi28)}
        v29 = append(v29, (fi28 - fi28))
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        var v31 int16 = -32768
        var v32 uint64 = 1
        _ = fi30
        _ = v31
        _ = v32
    }
    if false {
        switch ((("foo bar" + (("hello" + "a") + "foo bar")) + ("a" + ("" + "\n\t"))) + "foo bar") {
        case "hello":
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 int = fi33
                var v35 uint = (v13 | (v13 | 255))
                v1[(v34 % ((32768) | 1))] = (-(fi33))
                _ = fi33
                _ = v34
                _ = v35
            }
        case "\n\t":
            v36 := []int{32768, 1, 127}
            _ = v36
        }
        fmt.Println(v13, v9)
    } else {
        switch ((("a" + ("foo bar" + ("" + "foo bar"))) + ("\n\t" + (("\n\t" + "hello") + "hello"))) + ("foo bar" + "")) {
        case "a":
            var v37 int16 = -374
            _ = v37
        }
        v9 -= genFn0("foo bar", 32767, 128)
        v38 := []int{}
        _ = v38
    }
    v39 := GenS3{F0: v13, F1: 1, F2: (!(!false))}
    v40 := GenS4{F0: "hello"}
    var v41 uint16 = 1
    for fi42 := 0; fi42 < 5; fi42++ {
        v13 = 2
        _ = fi42
    }
    var v43 int16 = -995
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v43
    _ = v8
    _ = v9
}
