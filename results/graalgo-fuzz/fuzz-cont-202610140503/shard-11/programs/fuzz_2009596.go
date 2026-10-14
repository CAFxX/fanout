package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint8 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0() string {
    var v6 bool = (!(!(!false)))
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint64 = 42
        fmt.Println(fi7, v8, v8)
        for fi9 := 0; fi9 < 4; fi9++ {
            v8 = (v8 << ((65535) & 63))
            v6 = v6
            v10 := GenS1{F0: 0, F1: 1, F2: 208}
            switch "foo bar" {
            case "a":
                fmt.Println(fi9)
            case "hello":
                var v11 int32 = 779
                v12 := []int{(^(fi7)), fi7}
                fmt.Println(fi7)
                switch (fi9 >> ((1) & 63)) {
                case 0:
                    v11 = 42
                case -66:
                    var v13 uint8 = 255
                    _ = v13
                default:
                    if len(v12) > 0 {
                        v12[0] = fi7
                    }
                }
                _ = v11
                _ = v12
            case "\n\t":
                v14 := GenS1{F0: 255, F1: -1, F2: 2}
                _ = v14
            }
            v15 := []int{(fi7 - (fi9 * -9223372036854775808))}
            _ = fi9
            _ = v10
            _ = v15
        }
        _ = fi7
        _ = v8
    }
    _ = v6
    return (("foo bar" + "foo bar") + "hello")
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]bool)
    var v3 uint8 = 2
    var v4 float32 = 0.1
    var v5 string = (("" + "\n\t") + "\n\t")
    if true {
        if ((1 == 579) || (!false)) {
            var v16 int = 1
            _ = v16
        }
        v4 = float32((1.5 / 3.14159))
    }
    fmt.Println(v4, v4)
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(v3, fi17, v5)
        _ = fi17
    }
    v18 := []int{0, int(v1)}
    var v19 uint32 = (v1 + v1)
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS2{F0: v5, F1: -1}
        var v22 string = v5
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int64 = 623
    fmt.Println(v5)
    if ((v5 + (v5 + (v5 + ("" + ((((("\n\t" + v5) + ("hello" + ("a" + (("\n\t" + "\n\t") + (("hello" + v5) + v5))))) + ("a" + v5)) + v5) + ("hello" + v5)))))) == "") {
        switch 548 {
        case 726:
            v3 = (v3 % ((2) | 1))
        case 4294967295:
            for fi24 := 0; fi24 < 18; fi24++ {
                fmt.Println(v3, v1, fi24)
                for fi25 := 0; fi25 < 8; fi25++ {
                    var v26 uint64 = 18446744073709551615
                    if (!true) {
                        var v27 float32 = 3.14159
                        v19 = (v19 * 2147483647)
                        var v28 uint8 = (^(v3))
                        _ = v27
                        _ = v28
                    } else {
                        v19 = v19
                    }
                    if len(v18) > 0 {
                        v18[0] = fi24
                    }
                    _ = fi25
                    _ = v26
                }
                if true {
                    if false {
                        var v29 float32 = 1e10
                        var v30 int8 = 0
                        var v31 int8 = -66
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        if (!((v3 << ((18446744073709551615) & 63)) == 1)) {
                            fmt.Println(fi24, fi24)
                        }
                        switch 32768 {
                        case 1:
                            var v32 int = fi24
                            _ = v32
                        case 32768:
                            _ = len(v18)
                        case -9223372036854775808:
                            for fi33 := 0; fi33 < 8; fi33++ {
                                v19 -= v1
                                _ = fi33
                            }
                        }
                    }
                }
                _ = fi24
            }
        case -2147483648:
            var v34 float32 = (v4 / v4)
            _ = v34
        default:
            v5 = (((v5 + "a") + "") + v5)
        }
        var v35 bool = true
        v4 -= float32((-2.25 + 0.0))
        _ = v35
    }
    v1 += (v19 >> ((0) & 63))
    var v36 float32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
