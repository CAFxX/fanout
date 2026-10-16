package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 string; F2 uint16 }
func genFn0(p0 int64, p1 int64, p2 string) int32 {
    if (!(100 >= 42)) {
        for fi4 := 0; fi4 < 9; fi4++ {
            if (0.1 == -2.25) {
                var v5 string = p2
                var v6 uint32 = 2147483647
                _ = v5
                _ = v6
            } else {
                v7 := []int{(^(fi4)), fi4, (fi4 ^ (fi4 + (fi4 >> ((65535) & 63)))), fi4}
                var v8 int = (fi4 >> ((203) & 63))
                _ = v7
                _ = v8
            }
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 string = p2
                switch 723 {
                case 127:
                    v11 := []int{(fi9 | fi9), -1}
                    _ = v11
                case -277:
                    fmt.Println(p0, v10, fi9)
                case -9223372036854775808:
                    switch "a" {
                    case "":
                        v12 := GenS1{F0: -128, F1: 1}
                        _ = v12
                    case "hello":
                        v10 = p2
                    case "foo bar":
                        fmt.Println(v10)
                    default:
                        var v13 uint8 = 200
                        _ = v13
                    }
                default:
                    fmt.Println(fi9, fi4, fi4)
                }
                _ = fi9
                _ = v10
            }
            _ = fi4
        }
        var v14 uint16 = 616
        v15 := []int{9223372036854775807, 215, 32768, 127}
        _ = v14
        _ = v15
    } else {
        p2 = p2
        v16 := []int{51}
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 int = (fi17 >> ((323) & 63))
        _ = fi17
        _ = v18
    }
    var v19 float32 = 0.0
    v19 += (v19 - (v19 + 3.14159))
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return -592
}
func genFn1() bool {
    switch ("\n\t" + ("" + (("" + ("a" + ("\n\t" + "foo bar"))) + "a"))) {
    case "a":
        for fi20 := 0; fi20 < 1; fi20++ {
            switch (fi20 * (fi20 % ((fi20) | 1))) {
            case -9223372036854775808:
                fmt.Println(fi20)
            case 2147483647:
                var v21 bool = (!false)
                _ = v21
            case -623:
                switch fi20 {
                case 42:
                    var v22 bool = true
                    _ = v22
                case -9223372036854775808:
                    v23 := GenS2{F0: 303, F1: 65535}
                    _ = v23
                case 32768:
                    switch fi20 {
                    case -2147483648:
                        fmt.Println(fi20)
                    case 889:
                        v24 := []int{fi20, (^(fi20))}
                        _ = v24
                    }
                }
            }
            var v25 uint32 = 0
            var v26 uint16 = 897
            _ = fi20
            _ = v25
            _ = v26
        }
    }
    switch -713 {
    case 32768:
        if false {
            var v27 int = -659
            for fi28 := 0; fi28 < 2; fi28++ {
                v29 := []int{-2147483648, fi28, fi28}
                v30 := GenS1{F0: -10, F1: 1}
                fmt.Println(v27)
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v27
        }
    case 861:
        var v31 int16 = 723
        _ = v31
    case 9223372036854775807:
        var v32 bool = false
        _ = v32
    }
    return true
}

func main() {
    var v1 int8 = 0
    var v2 float32 = (-2.25 - (1e10 + -2.25))
    v3 := []int{}
    v33 := GenS3{F0: 2147483648}
    v2 = v2
    v34 := GenS4{F0: -1, F1: "hello", F2: 20}
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = 42
    }
    if false {
        if len(v3) > 0 {
            v3[0] = -120
        }
        if (-2.25 <= 1.5) {
            var v35 int64 = 443
            v36 := GenS3{F0: 65535}
            _ = v35
            _ = v36
        } else {
            v37 := GenS3{F0: 509}
            _ = v37
        }
        v38 := GenS3{F0: 2147483648}
        _ = v38
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v39 uint = 0
    var v40 int32 = 0
    var v41 uint16 = 0
    if false {
        fmt.Println(v39, v2, v41)
        v42 := GenS2{F0: 9223372036854775807, F1: 424}
        fmt.Println(v39)
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v39
    _ = v40
    _ = v41
}
