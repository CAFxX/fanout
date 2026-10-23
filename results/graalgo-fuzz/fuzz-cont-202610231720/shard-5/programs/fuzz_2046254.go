package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint; F2 int }
type GenS6 struct { F0 uint8; F1 int8; F2 bool }
type GenS7 struct { F0 uint16; F1 uint32; F2 uint16 }
func genFn0() int {
    var v7 uint16 = 65535
    if true {
        v8 := []int{9223372036854775807, 4294967295, 42, 0}
        switch 743 {
        case 2147483647:
            switch "a" {
            case "foo bar":
                fmt.Println(v7)
            case "":
                fmt.Println(v7)
            case "\n\t":
                v7 -= 1
            }
        default:
            fmt.Println(v7)
        }
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        if true {
            if (1 == 539) {
                v7 -= 65535
                if ((18446744073709551615 < 266) || false) {
                    v10 := GenS1{F0: -128, F1: (("foo bar" + "a") + "hello")}
                    _ = v10
                } else {
                    switch "a" {
                    case "hello":
                        fmt.Println(fi9)
                    default:
                        v7 += (v7 | (v7 >> ((131) & 63)))
                    }
                    v11 := []int{(fi9 << ((1) & 63))}
                    if (fi9 < (fi9 % ((-2147483648) | 1))) {
                        var v12 int64 = 113
                        _ = v12
                    } else {
                        var v13 int64 = -1
                        _ = v13
                    }
                    _ = v11
                }
            }
        }
        v14 := GenS2{F0: (v7 ^ v7)}
        _ = fi9
        _ = v14
    }
    _ = v7
    return 32768
}
func genFn1(p0 bool, p1 string) uint32 {
    var v15 uint = 2
    switch 9223372036854775807 {
    case 127:
        v16 := GenS1{F0: -128, F1: "foo bar"}
        _ = v16
    }
    v17 := GenS3{F0: 0}
    switch 0 {
    case 922:
        v18 := []int{-323, genFn0(), 166, 297}
        _ = v18
    case 682:
        v19 := GenS1{F0: -37, F1: "a"}
        _ = v19
    case 9223372036854775807:
        var v20 int16 = 0
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v17
    return 0
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    var v4 bool = true
    var v5 uint64 = 18446744073709551615
    var v6 uint32 = 929
    fmt.Println(v4, v5)
    if (540 < 265) {
        var v21 float32 = -2.25
        fmt.Println(v1, v6, v5)
        _ = v21
    } else {
        var v22 string = "a"
        _ = v22
    }
    if v4 {
        if false {
            v6 = (v6 + v6)
        } else {
            v23 := GenS2{F0: 577}
            var v24 int16 = 370
            _ = v23
            _ = v24
        }
    } else {
        v4 = (!(!((1e10 / 1.0) > (-2.25 / (1.5 - (3.14159 / 0.1))))))
        var v25 float64 = 0.1
        v4 = v4
        _ = v25
    }
    v26 := GenS4{F0: -955, F1: 1}
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 8; fi28++ {
            v29 := []int{(fi27 * (^(fi28))), (fi27 - fi27)}
            switch fi28 {
            case -1:
                switch (-(fi27)) {
                case 32767:
                    v30 := GenS5{F0: v5, F1: 672, F2: 474}
                    _ = v30
                case 714:
                    var v31 int32 = (v1 & v1)
                    switch "hello" {
                    case "\n\t":
                        v32 := GenS6{F0: 255, F1: 0, F2: true}
                        _ = v32
                    default:
                        delete(v3, 123)
                    }
                    v33 := v3[genFn0()]
                    _ = v31
                    _ = v33
                default:
                    var v34 int32 = (^(v1))
                    switch (-(fi28)) {
                    case 2147483647:
                        var v35 uint8 = 127
                        _ = v35
                    case -555:
                        v36 := GenS7{F0: 32767, F1: v6, F2: 686}
                        _ = v36
                    }
                    var v37 int32 = v34
                    _ = v34
                    _ = v37
                }
            case 4294967295:
                fmt.Println(fi28, v1, v4)
            case 311:
                v6 += uint32(v5)
            }
            var v38 bool = ((v5 ^ v5) > (v5 % (((v5 ^ 386)) | 1)))
            _ = fi28
            _ = v29
            _ = v38
        }
        _ = fi27
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 float64 = 1.5
        if (v5 == v5) {
            var v41 uint16 = 0
            _ = v41
        } else {
            fmt.Println(v6, v40)
            fmt.Println(v6, v6, v5)
        }
        var v42 uint = 255
        _ = fi39
        _ = v40
        _ = v42
    }
    switch "hello" {
    case "a":
        v5 = 865
    case "foo bar":
        switch ("\n\t" + "\n\t") {
        case "a":
            if (uint8(v6) > 255) {
                v4 = v4
            } else {
                v1 -= v1
            }
        default:
            for fi43 := 0; fi43 < 13; fi43++ {
                v44 := GenS1{F0: 0, F1: ((("a" + ((("foo bar" + "hello") + ((("hello" + (("hello" + ("hello" + "\n\t")) + "foo bar")) + "foo bar") + "")) + "\n\t")) + ("foo bar" + "\n\t")) + ("foo bar" + ""))}
                _ = fi43
                _ = v44
            }
        }
    }
    fmt.Println(v6, v5, v4)
    fmt.Println(v5)
    var v45 uint64 = (v5 + v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
