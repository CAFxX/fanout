package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }
type GenS2 struct { F0 string; F1 uint64; F2 uint }
func genFn0() uint16 {
    v7 := []int{-222}
    switch 127 {
    case 287:
        var v8 int8 = -128
        _ = v8
    default:
        switch ("" + ("a" + "foo bar")) {
        case "":
            if (-2.25 > ((-2.25 / float32((1e10 / 3.14159))) * float32(0.0))) {
                v9 := GenS1{F0: 2147483647, F1: 183, F2: 0}
                var v10 int8 = -128
                var v11 uint64 = 1
                _ = v10
                _ = v11
                _ = v9
            }
        }
    }
    switch "foo bar" {
    case "\n\t":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case "":
        var v12 uint = 1
        _ = v12
    default:
        var v13 uint64 = 18446744073709551615
        _ = v13
    }
    _ = v7
    return 0
}

func main() {
    v1 := []string{"", ("foo bar" + "foo bar"), ("a" + "")}
    v2 := []string{"foo bar"}
    v3 := []int{1}
    var v4 int8 = 0
    var v5 int64 = 9223372036854775807
    v6 := make(map[int]string)
    fmt.Println(v5)
    v4 = (v4 % (((v4 & -128)) | 1))
    v14 := GenS2{F0: "foo bar", F1: 2, F2: 1}
    switch "foo bar" {
    case "\n\t":
        if true {
            var v15 uint8 = 128
            _ = v15
        }
    case "a":
        var v16 float32 = 0.1
        _ = v16
    case "hello":
        fmt.Println(v4)
    default:
        if true {
            v4 += v4
            switch ("" + "") {
            case "foo bar":
                var v17 int = 464
                _ = v17
            default:
                _ = len(v6)
            }
        }
    }
    v18 := v6[0]
    if false {
        switch 9223372036854775807 {
        case 4294967295:
            if (!true) {
                switch -9223372036854775808 {
                case 4294967295:
                    for fi19 := 0; fi19 < 2; fi19++ {
                        _ = len(v2)
                        fmt.Println(fi19, fi19)
                        _ = fi19
                    }
                case 127:
                    v20 := v6[-564]
                    _ = v20
                default:
                    var v21 string = (v18 + "foo bar")
                    _ = v21
                }
            } else {
                v18 = v18
            }
        case -1:
            v22 := GenS2{F0: ("hello" + (("\n\t" + "a") + (v18 + v18))), F1: 2, F2: 2}
            _ = v22
        }
        switch v18 {
        case "":
            if ((v5 & (-(v5))) <= (v5 * -9223372036854775808)) {
                var v23 uint = 443
                var v24 int16 = 42
                _ = v23
                _ = v24
            } else {
                fmt.Println(v18, v4, v18)
            }
        default:
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 uint8 = 135
                _ = fi25
                _ = v26
            }
        }
        var v27 uint64 = 2
        _ = v27
    }
    v2 = append(v2, v18)
    v28 := GenS2{F0: "", F1: 0, F2: 836}
    _ = len(v6)
    _ = len(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
