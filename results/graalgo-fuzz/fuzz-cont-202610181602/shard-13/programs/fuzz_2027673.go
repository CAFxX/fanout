package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 uint; F2 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 int32 }
func genFn0() uint16 {
    if true {
        var v5 int = 9223372036854775807
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 string = ("foo bar" + "hello")
            if ((-21 <= 127) || (987 <= 533)) {
                var v8 int = v5
                _ = v8
            } else {
                v9 := GenS1{F0: 30, F1: -2147483648, F2: (0.1 >= float32(1e10))}
                v7 = (v7 + v7)
                v10 := []int{}
                _ = v10
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    v11 := []int{2147483647, 32767, 464, 598}
    v12 := GenS1{F0: -461, F1: -1, F2: (((!(false && true)) && true) && false)}
    _ = v11
    _ = v12
    return 432
}
func genFn1(p0 bool, p1 uint64, p2 uint32) uint32 {
    fmt.Println(p1)
    v13 := GenS1{F0: 42, F1: -2147483648, F2: p0}
    p1 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p2
}

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 255
    var v3 string = "foo bar"
    v4 := make(map[string]bool)
    v14 := []int{-1}
    v1 = ""
    v3 = v3
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint8 = 127
        switch v1 {
        case "\n\t":
            var v17 int16 = 42
            _ = v17
        case "":
            if false {
                v18 := v4[(v1 + "a")]
                _ = v18
            } else {
                var v19 int32 = -207
                _ = v19
            }
        default:
            var v20 int16 = -1
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v1, v2, v2)
    switch ((("\n\t" + v3) + v1) + "\n\t") {
    case "hello":
        if (0.1 < float64(3.14159)) {
            v21 := GenS2{F0: 268}
            for fi22 := 0; fi22 < 17; fi22++ {
                for fi23 := 0; fi23 < 17; fi23++ {
                    v1 = v1
                    v1 = ("\n\t" + v1)
                    _ = fi23
                }
                v24 := GenS3{F0: 581, F1: fi22, F2: 42}
                _ = fi22
                _ = v24
            }
            for fi25 := 0; fi25 < 10; fi25++ {
                if (-995 <= 42) {
                    var v26 uint = 18446744073709551615
                    _ = v26
                } else {
                    v27 := GenS3{F0: 0, F1: fi25, F2: 4294967295}
                    var v28 float64 = (0.0 + float64(-2.25))
                    var v29 string = (v1 + (v1 + "foo bar"))
                    _ = v27
                    _ = v28
                    _ = v29
                }
                v1 = (v3 + ((((v3 + v1) + (v1 + ("a" + v3))) + ((v3 + ((("a" + v3) + v1) + (v1 + (v1 + ("hello" + ("foo bar" + v1)))))) + "hello")) + (((v1 + "hello") + (((v1 + v3) + "") + "\n\t")) + v1)))
                v30 := GenS4{F0: 2}
                _ = fi25
                _ = v30
            }
            _ = v21
        }
    case "\n\t":
        v2 = (v2 / ((139) | 1))
    case "":
        v3 = v3
    default:
        switch -2147483648 {
        case 127:
            for fi31 := 0; fi31 < 10; fi31++ {
                for fi32 := 0; fi32 < 15; fi32++ {
                    var v33 uint8 = 148
                    var v34 uint64 = 524
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                v1 = v1
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                _ = fi31
            }
        case -9223372036854775808:
            var v35 bool = true
            _ = v35
        }
    }
    v1 = (v3 + (v1 + (v1 + v1)))
    switch 42 {
    case 32768:
        v2 = v2
    case 2147483647:
        v36 := GenS5{F0: "hello", F1: 0, F2: v3}
        _ = v36
    case 0:
        switch 9223372036854775807 {
        case -139:
            var v37 uint16 = v2
            _ = v37
        default:
            _ = len(v14)
        }
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        v39 := GenS6{F0: genFn0()}
        fmt.Println(v1, v2, v1)
        switch "foo bar" {
        case "a":
            for fi40 := 0; fi40 < 18; fi40++ {
                var v41 int16 = -1
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v3)
                case "hello":
                    v42 := GenS7{F0: true, F1: -430}
                    _ = v42
                case "\n\t":
                    fmt.Println(fi40, fi38)
                default:
                    _ = len(v14)
                }
                _ = fi40
                _ = v41
            }
        case "foo bar":
            var v43 float32 = -2.25
            _ = v43
        }
        _ = fi38
        _ = v39
    }
    v3 = v3
    if (!true) {
        if (false || true) {
            v44 := GenS1{F0: -2147483648, F1: 42, F2: (((0.1 / 1.0) + 0.1) == 3.14159)}
            _ = v44
        } else {
            var v45 int16 = -1
            var v46 int16 = v45
            v47 := GenS2{F0: -9223372036854775808}
            _ = v45
            _ = v46
            _ = v47
        }
    } else {
        v1 = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
