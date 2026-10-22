package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint32, p2 uint64) int8 {
    for fi5 := 0; fi5 < 19; fi5++ {
        switch (fi5 / (((fi5 * fi5)) | 1)) {
        case 128:
            p0 -= 0
        default:
            v6 := []int{(fi5 - fi5), 32767}
            _ = v6
        }
        _ = fi5
    }
    var v7 float64 = (0.1 + float64(1.5))
    var v8 int16 = 0
    if (47 < 1) {
        var v9 int8 = 112
        var v10 int8 = (v9 + v9)
        var v11 int32 = (p0 | (p0 / ((p0) | 1)))
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 string, p1 int8) int {
    v12 := GenS1{F0: p0}
    switch (("" + "foo bar") + p0) {
    case "a":
        fmt.Println(p1)
    case "foo bar":
        switch (("hello" + p0) + p0) {
        case "foo bar":
            v13 := GenS1{F0: p0}
            _ = v13
        case "hello":
            if (791 < 515) {
                switch p0 {
                case "\n\t":
                    p0 = ("foo bar" + p0)
                case "a":
                    var v14 uint16 = 1
                    _ = v14
                case "foo bar":
                    v15 := GenS2{F0: 726, F1: 32767}
                    _ = v15
                }
            }
        }
    default:
        var v16 int = 470
        _ = v16
    }
    p0 = ("a" + (p0 + "a"))
    v17 := []int{}
    _ = p0
    _ = p1
    _ = v12
    _ = v17
    return 721
}

func main() {
    var v1 uint16 = 255
    var v2 uint16 = (v1 << ((489) & 63))
    var v3 uint16 = v2
    var v4 uint32 = 213
    if true {
        var v18 bool = (!(9223372036854775807 <= 849))
        _ = v18
    }
    if true {
        v1 = 2
        v1 -= 158
        v4 = v4
    } else {
        v3 -= 1
        var v19 uint16 = (v1 * 255)
        for fi20 := 0; fi20 < 13; fi20++ {
            if true {
                for fi21 := 0; fi21 < 6; fi21++ {
                    fmt.Println(fi20, v2, v19)
                    var v22 uint64 = 4294967295
                    _ = fi21
                    _ = v22
                }
                var v23 float64 = 3.14159
                _ = v23
            } else {
                if false {
                    var v24 int = (fi20 << ((65535) & 63))
                    if (((((1.5 * float32(0.1)) / 1e10) / 1.0) * ((3.14159 + (1e10 + (3.14159 + float32(1.5)))) - (0.0 + 1.5))) == (-2.25 - (1e10 / ((3.14159 - 3.14159) + (0.1 * 1.5))))) {
                        if false {
                            v1 = v2
                        }
                        v25 := GenS1{F0: "foo bar"}
                        _ = v25
                    } else {
                        v26 := []int{}
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        _ = v26
                    }
                    _ = v24
                } else {
                    var v27 bool = true
                    var v28 string = ("a" + "hello")
                    fmt.Println(v28, v2, v19)
                    _ = v27
                    _ = v28
                }
                switch (fi20 % ((0) | 1)) {
                case -9223372036854775808:
                    v29 := GenS3{F0: -1, F1: 307}
                    _ = v29
                case -475:
                    v30 := []int{(fi20 - 0), fi20, fi20, fi20}
                    _ = v30
                default:
                    v31 := []int{-717, fi20, (fi20 * int(v2)), fi20}
                    _ = v31
                }
            }
            switch 42 {
            case 438:
                fmt.Println(v3, fi20, v1)
            default:
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v4, v3, v3)
                default:
                    var v32 uint32 = 2
                    _ = v32
                }
            }
            _ = fi20
        }
        _ = v19
    }
    fmt.Println(v4, v4, v2)
    var v33 int32 = -2147483648
    v4 += v4
    v34 := GenS4{F0: (v4 << ((969) & 63))}
    var v35 int64 = int64(v33)
    v2 += 0
    var v36 int = genFn1("", 62)
    switch genFn1("\n\t", -128) {
    case 42:
        for fi37 := 0; fi37 < 13; fi37++ {
            if (false && true) {
                var v38 uint64 = uint64(v4)
                var v39 int64 = v35
                _ = v38
                _ = v39
            } else {
                var v40 uint = 1
                _ = v40
            }
            var v41 float32 = ((3.14159 / -2.25) - 1e10)
            _ = fi37
            _ = v41
        }
    case -186:
        fmt.Println(v2, v1, v36)
    default:
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
