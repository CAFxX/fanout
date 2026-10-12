package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int64; F2 string }
type GenS3 struct { F0 bool }
func genFn0(p0 int64, p1 uint8, p2 uint) bool {
    var v7 string = ("foo bar" + "")
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int16 = 1000
        _ = fi8
        _ = v9
    }
    fmt.Println(p0)
    if false {
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return false
}

func main() {
    var v1 int32 = -1
    var v2 uint = 42
    var v3 bool = false
    var v4 uint16 = 42
    var v5 int = 0
    var v6 bool = false
    if v3 {
        if v3 {
            var v10 float32 = (0.0 * 1.5)
            _ = v10
        }
    } else {
        v11 := GenS1{F0: v1, F1: 1}
        var v12 float32 = (((float32((0.1 + 0.0)) + 1e10) * 0.0) - ((0.0 - 1.5) * (1.5 / 1.0)))
        switch "a" {
        case "a":
            if true {
                var v13 int16 = 1
                v14 := []int{v5, (-(v5)), 42}
                _ = v13
                _ = v14
            }
        case "hello":
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := []int{fi15, (fi15 >> (((v2 & (v2 * v2))) & 63)), fi15, 4294967295}
                _ = fi15
                _ = v16
            }
        case "":
            v17 := GenS2{F0: v4, F1: 269, F2: ((("\n\t" + ("hello" + "hello")) + "") + "\n\t")}
            _ = v17
        default:
            fmt.Println(v1)
        }
        _ = v11
        _ = v12
    }
    if (((^(v1)) == (v1 / ((42) | 1))) || false) {
        v1 = 1
        var v18 uint64 = 18446744073709551615
        var v19 uint64 = v18
        _ = v18
        _ = v19
    }
    v6 = (1e10 >= (1.5 / -2.25))
    var v20 float64 = 3.14159
    v4 = 996
    v6 = (v3 && v3)
    for fi21 := 0; fi21 < 20; fi21++ {
        fmt.Println(v3, v5)
        _ = fi21
    }
    var v22 uint64 = 245
    v20 = 1e10
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := []int{(fi23 & (-(v5))), v5}
        switch v5 {
        case -800:
            var v25 int8 = -128
            _ = v25
        case -1:
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
        }
        if len(v24) > 0 {
            v24[0] = (fi23 + v5)
        }
        _ = fi23
        _ = v24
    }
    var v26 uint32 = uint32(v4)
    v27 := GenS1{F0: (^(v1)), F1: 2}
    switch ("hello" + "foo bar") {
    case "foo bar":
        v6 = v6
    case "hello":
        for fi28 := 0; fi28 < 10; fi28++ {
            for fi29 := 0; fi29 < 2; fi29++ {
                switch "foo bar" {
                case "a":
                    if (1 == 255) {
                        if v3 {
                            v1 = (v1 - v1)
                            if ((v5 & 4294967295) > fi29) {
                                var v30 uint16 = 0
                                var v31 bool = (false || false)
                                _ = v30
                                _ = v31
                            }
                            var v32 float32 = -2.25
                            var v33 int16 = 755
                            var v34 int = v5
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        if (!(!(100 > 56))) {
                            var v35 uint = v2
                            var v36 uint64 = 18446744073709551615
                            _ = v35
                            _ = v36
                        }
                        fmt.Println(v4, v2)
                    } else {
                        v37 := GenS3{F0: (!false)}
                        v38 := []int{(fi28 + 32767), (-(fi29)), v5}
                        _ = v37
                        _ = v38
                    }
                case "\n\t":
                    v4 = 17
                case "foo bar":
                    v3 = ("" > "hello")
                }
                fmt.Println(v26, v3)
                var v39 int16 = 32767
                _ = fi29
                _ = v39
            }
            _ = fi28
        }
    default:
        for fi40 := 0; fi40 < 5; fi40++ {
            for fi41 := 0; fi41 < 17; fi41++ {
                v42 := []int{}
                _ = len(v42)
                _ = len(v42)
                _ = fi41
                _ = v42
            }
            _ = fi40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
