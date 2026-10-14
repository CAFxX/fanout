package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int16; F2 uint }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0() int64 {
    var v6 float64 = float64((1e10 - float32(0.0)))
    v7 := []int{}
    fmt.Println(v6)
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    var v4 int64 = 1
    var v5 uint16 = 32767
    switch "foo bar" {
    case "a":
        var v8 float64 = v1
        _ = v8
    case "foo bar":
        fmt.Println(v5, v1)
    case "hello":
        switch "a" {
        case "":
            v5 -= 42
        case "hello":
            v4 = (v4 - (v4 >> ((1) & 63)))
        case "foo bar":
            var v9 float32 = 0.0
            _ = v9
        default:
            v10 := GenS1{F0: (42 < 61)}
            _ = v10
        }
    }
    v11 := GenS1{F0: true}
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: false}
        var v14 int8 = -1
        v14 = v14
        _ = fi12
        _ = v13
        _ = v14
    }
    if ((v4 & (-(v4))) >= (v4 * v4)) {
        if (74 <= 0) {
            if (false && (!true)) {
                for fi15 := 0; fi15 < 6; fi15++ {
                    if false {
                        switch "hello" {
                        case "hello":
                            v1 -= 1.5
                        case "a":
                            v1 = v1
                        }
                        v16 := v2[(("\n\t" + (("foo bar" + ("hello" + "\n\t")) + ((("foo bar" + "a") + "foo bar") + (((("foo bar" + "foo bar") + "a") + "\n\t") + "a")))) + "hello")]
                        v17 := []int{}
                        _ = v16
                        _ = v17
                    } else {
                        if false {
                            switch 127 {
                            case 128:
                                fmt.Println(v5)
                            case 32767:
                                v18 := GenS2{F0: (fi15 << ((295) & 63)), F1: 275, F2: 42}
                                _ = v18
                            }
                            v19 := GenS2{F0: fi15, F1: 42, F2: 185}
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 12; fi20++ {
                            v1 = float64(0.1)
                            switch "\n\t" {
                            case "":
                                var v21 string = ("a" + ((("" + ("foo bar" + "foo bar")) + "hello") + (("a" + ("foo bar" + "hello")) + ("" + ""))))
                                _ = v21
                            case "\n\t":
                                v1 -= v1
                            default:
                                if (-156 <= fi20) {
                                    fmt.Println(v1, fi20, v1)
                                } else {
                                    v22 := GenS1{F0: false}
                                    _ = v22
                                }
                            }
                            _ = fi20
                        }
                        var v23 uint = 14
                        _ = v23
                    }
                    delete(v2, (("foo bar" + ((((("" + ("hello" + "")) + "hello") + "a") + (("" + (((("foo bar" + ("a" + "foo bar")) + (("a" + (("" + "a") + "")) + "a")) + "hello") + "")) + "foo bar")) + "hello")) + "foo bar"))
                    v24 := GenS3{F0: 490, F1: 378}
                    _ = fi15
                    _ = v24
                }
                v25 := GenS2{F0: 2147483647, F1: 0, F2: 109}
                _ = v25
            } else {
                var v26 int8 = 0
                v26 += (v26 | (v26 / (((v26 - 65)) | 1)))
                v27 := GenS2{F0: 42, F1: 820, F2: 822}
                _ = v26
                _ = v27
            }
        }
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        v29 := []int{fi28}
        _ = fi28
        _ = v29
    }
    var v30 int8 = -1
    var v31 float64 = v1
    v32 := []int{-2147483648, 32768, 615, 32767}
    var v33 int8 = v30
    delete(v3, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
