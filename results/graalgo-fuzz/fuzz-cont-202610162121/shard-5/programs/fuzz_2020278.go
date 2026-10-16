package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint8) bool {
    if false {
        v7 := []int{128, 937, 4294967295}
        var v8 bool = ((-9223372036854775808 < 1) && false)
        if v8 {
            if v8 {
                v8 = false
                p0 += p0
            } else {
                p0 += (p0 | (p0 | 128))
                for fi9 := 0; fi9 < 4; fi9++ {
                    fmt.Println(fi9)
                    _ = fi9
                }
                var v10 string = "foo bar"
                _ = v10
            }
        }
        _ = v7
        _ = v8
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            fmt.Println(fi11, p0)
            p0 = (p0 % (((p0 & (p0 | p0))) | 1))
            _ = fi11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                switch "hello" {
                case "foo bar":
                    if false {
                        v14 := GenS1{F0: 4294967295, F1: 4294967295}
                        v15 := GenS2{F0: p0}
                        switch "a" {
                        case "\n\t":
                            p0 = 128
                        case "a":
                            v16 := []int{fi13, fi13, fi12}
                            _ = v16
                        default:
                            fmt.Println(fi12)
                        }
                        switch fi12 {
                        case 32767:
                            switch ((((("" + "hello") + "") + "a") + "") + "") {
                            case "hello":
                                var v17 int64 = 0
                                _ = v17
                            case "a":
                                v18 := []int{}
                                _ = v18
                            case "\n\t":
                                fmt.Println(p0, fi13)
                            }
                        default:
                            if false {
                                var v19 uint16 = 65535
                                _ = v19
                            }
                        }
                        var v20 float32 = -2.25
                        _ = v14
                        _ = v15
                        _ = v20
                    }
                default:
                    switch (fi12 << ((877) & 63)) {
                    case 478:
                        var v21 string = (("foo bar" + "hello") + "\n\t")
                        _ = v21
                    case 127:
                        v22 := GenS2{F0: 170}
                        _ = v22
                    }
                }
                _ = fi13
            }
            v23 := GenS3{F0: 2}
            _ = fi12
            _ = v23
        }
    }
    var v24 int8 = 1
    var v25 float64 = (1e10 / (1e10 + ((-2.25 * float64(((3.14159 / -2.25) * 3.14159))) - 1.5)))
    _ = p0
    _ = v24
    _ = v25
    return (1 > 0)
}

func main() {
    var v1 float64 = 1e10
    var v2 string = "foo bar"
    var v3 bool = (-1 != -2147483648)
    v4 := make(map[uint32]bool)
    v5 := []string{v2, v2}
    v6 := make(map[uint32]string)
    for fi26 := 0; fi26 < 15; fi26++ {
        fmt.Println(v3, v1, v2)
        _ = fi26
    }
    if (!(4294967295 < 1)) {
        var v27 uint64 = 42
        fmt.Println(v27, v1)
        if (0 > (^(v27))) {
            if true {
                v27 -= v27
            }
        } else {
            fmt.Println(v3)
        }
        _ = v27
    } else {
        v28 := GenS3{F0: 18446744073709551615}
        _ = v28
    }
    var v29 uint64 = 18446744073709551615
    fmt.Println(v29, v3, v2)
    v30 := v6[0]
    var v31 string = (((v2 + v30) + "") + ((("" + v2) + ("hello" + v2)) + v2))
    v4[4294967295] = true
    if (v31 < "\n\t") {
        var v32 uint16 = 65535
        var v33 bool = v3
        v3 = (v1 >= (0.1 / v1))
        _ = v32
        _ = v33
    } else {
        for fi34 := 0; fi34 < 8; fi34++ {
            v3 = ((v1 == v1) || (1 == -1))
            _ = fi34
        }
        delete(v4, 157)
    }
    var v35 uint8 = 151
    delete(v4, 2)
    _ = len(v5)
    v29 -= 961
    v36 := GenS4{F0: -32768, F1: 1}
    v37 := GenS5{F0: -772}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
