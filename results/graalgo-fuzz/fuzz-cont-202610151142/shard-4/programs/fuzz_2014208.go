package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint32; F2 string }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint64; F1 int8; F2 int16 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 42
    v3 := make(map[int]int)
    v4 := make(map[uint32]int)
    v5 := make(map[string]uint32)
    v6 := GenS1{F0: -32768}
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 string = "a"
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := GenS2{F0: (fi8 | (^(fi11))), F1: 314, F2: "\n\t"}
            _ = fi11
            _ = v12
        }
        _ = fi8
    }
    fmt.Println(v7, v7)
    switch 4294967295 {
    case 478:
        fmt.Println(v7)
    default:
        v13 := GenS1{F0: (^(v7))}
        _ = v13
    }
    if (("hello" <= "") || (!true)) {
        v14 := v1[4294967295]
        if ((128 > 32768) && (!((756 >= 32767) || ((v14 & v14) < (^(v14)))))) {
            v15 := GenS3{F0: 255, F1: (((("a" + "foo bar") + "") + "hello") + "\n\t")}
            _ = v15
        }
        v2 = (v2 ^ 0)
        _ = v14
    }
    var v16 int64 = -3
    if false {
        if true {
            var v17 bool = (722 <= 2147483647)
            _ = v17
        }
        var v18 bool = ((v2 & v2) != 18446744073709551615)
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 14; fi20++ {
            if true {
                v21 := GenS3{F0: 42, F1: "a"}
                var v22 string = "\n\t"
                _ = v21
                _ = v22
            }
            v16 = 1
            _ = fi20
        }
        var v23 int16 = -161
        switch "a" {
        case "a":
            v16 = (v16 >> ((332) & 63))
        case "":
            switch "a" {
            case "foo bar":
                if true {
                    v24 := GenS4{F0: v2, F1: -102, F2: int16(v16)}
                    _ = v24
                }
            case "hello":
                switch fi19 {
                case -462:
                    fmt.Println(v16, v23)
                case 32768:
                    v25 := GenS1{F0: v7}
                    _ = v25
                case -917:
                    v26 := []int{fi19}
                    _ = v26
                }
            case "\n\t":
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 string = ("hello" + ("\n\t" + "hello"))
                    fmt.Println(v2, v16, v28)
                    v29 := []int{9223372036854775807, fi27, fi27}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
        case "hello":
            if false {
                fmt.Println(v23, v2)
            } else {
                v30 := GenS2{F0: -1, F1: 2, F2: "a"}
                var v31 float64 = (1.5 / 3.14159)
                _ = v30
                _ = v31
            }
        }
        _ = fi19
        _ = v23
    }
    var v32 float64 = 0.1
    fmt.Println(v16, v32, v2)
    v33 := []int{}
    var v34 float32 = -2.25
    var v35 bool = ((true || (127 != 9223372036854775807)) || (42 <= -16))
    switch 4294967295 {
    case -2147483648:
        v36 := GenS4{F0: 2, F1: 42, F2: (v7 | (v7 - v7))}
        _ = v36
    case -663:
        v37 := GenS4{F0: 18446744073709551615, F1: 1, F2: (v7 & v7)}
        _ = v37
    case 32768:
        v38 := v4[4294967295]
        _ = v38
    }
    v39 := GenS5{F0: -53}
    v40 := GenS3{F0: 593, F1: "foo bar"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
