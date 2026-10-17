package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint64; F2 string }
func genFn0(p0 uint) int8 {
    var v5 uint8 = 88
    for fi6 := 0; fi6 < 11; fi6++ {
        if (true && false) {
            for fi7 := 0; fi7 < 11; fi7++ {
                v5 = 255
                fmt.Println(fi6, v5)
                if false {
                    v8 := GenS1{F0: false, F1: -9223372036854775808}
                    v9 := []int{fi7, fi6}
                    fmt.Println(v5)
                    _ = v8
                    _ = v9
                } else {
                    v10 := GenS1{F0: (!(!(fi6 > fi6))), F1: fi6}
                    switch ("\n\t" + ("a" + ("\n\t" + (("foo bar" + "hello") + "a")))) {
                    case "a":
                        var v11 uint16 = 65535
                        _ = v11
                    default:
                        if (!true) {
                            var v12 int = fi6
                            v5 += (v5 - 255)
                            v13 := []int{}
                            _ = v12
                            _ = v13
                        }
                    }
                    _ = v10
                }
                _ = fi7
            }
            switch ("foo bar" + ("" + "hello")) {
            case "hello":
                var v14 float32 = (-2.25 / 3.14159)
                _ = v14
            case "\n\t":
                for fi15 := 0; fi15 < 19; fi15++ {
                    p0 = 532
                    v16 := []int{127, 195}
                    _ = fi15
                    _ = v16
                }
            default:
                v17 := GenS2{F0: fi6}
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            p0 = p0
            fmt.Println(fi6, fi18, v5)
            v19 := GenS2{F0: 127}
            _ = fi18
            _ = v19
        }
        var v20 string = "hello"
        _ = fi6
        _ = v20
    }
    p0 -= 4294967295
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := []int{127, (fi21 / (((fi21 << ((p0) & 63))) | 1))}
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 bool = (42 >= 0)
    v4 := make(map[string]int)
    fmt.Println(v3)
    var v23 uint64 = 42
    var v24 int32 = -1
    var v25 uint16 = 529
    if (128 < int(v24)) {
        for fi26 := 0; fi26 < 17; fi26++ {
            v3 = (1e10 >= (float64((0.0 * (0.1 + 0.0))) * 3.14159))
            switch "a" {
            case "hello":
                switch 0 {
                case -293:
                    v24 -= -1
                default:
                    var v27 int16 = -32768
                    _ = v27
                }
            }
            _ = fi26
        }
        var v28 string = "\n\t"
        var v29 uint = 2
        _ = v28
        _ = v29
    }
    v23 -= (v23 | 2)
    var v30 float64 = float64((float32(0.0) + 1e10))
    switch ("\n\t" + ((("hello" + "hello") + "foo bar") + (("a" + "\n\t") + "\n\t"))) {
    case "\n\t":
        if (-1 > 678) {
            var v31 uint16 = 255
            v32 := GenS3{F0: 1000, F1: 65535}
            _ = v31
            _ = v32
        } else {
            v30 -= float64(((-2.25 + (1.5 + (((1.5 * ((-2.25 * 0.1) / (float32(v30) / 1e10))) / float32(v30)) + 1e10))) / ((3.14159 / -2.25) * 3.14159)))
            var v33 int = -2147483648
            v34 := GenS4{F0: v23, F1: 0, F2: ("a" + "a")}
            _ = v33
            _ = v34
        }
    default:
        v23 = v23
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        var v36 float32 = 1e10
        var v37 uint32 = uint32(v25)
        _ = fi35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        v39 := GenS3{F0: 514, F1: 2147483648}
        var v40 float64 = float64(-2.25)
        v41 := []int{(fi38 | (fi38 * fi38))}
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
}
