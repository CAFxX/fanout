package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 bool; F1 int8 }
func genFn0(p0 bool, p1 bool) int32 {
    var v7 int32 = -20
    if false {
        switch "foo bar" {
        case "\n\t":
            var v8 int8 = 100
            _ = v8
        case "hello":
            v9 := []int{32767, 52}
            _ = v9
        case "":
            fmt.Println(p1)
        default:
            v10 := GenS1{F0: 213}
            _ = v10
        }
        var v11 string = "foo bar"
        var v12 uint64 = 591
        _ = v11
        _ = v12
    } else {
        var v13 uint = 926
        fmt.Println(v13, v13)
        _ = v13
    }
    switch 32768 {
    case 32768:
        var v14 int = 9223372036854775807
        _ = v14
    default:
        v15 := GenS2{F0: 255, F1: 15}
        _ = v15
    }
    v16 := []int{127, 128, 4294967295}
    v17 := GenS3{F0: -1}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v7
    return v7
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]int)
    var v3 int = 42
    v4 := []uint32{2147483648}
    v5 := []int{v3, (v3 >> ((2) & 63))}
    var v6 string = ""
    v18 := GenS1{F0: 18446744073709551615}
    v1 = (v1 % ((v1) | 1))
    v19 := GenS4{F0: 1, F1: 450, F2: 42}
    fmt.Println(v3, v6, v6)
    if (!false) {
        if false {
            fmt.Println(v6)
            if (128 >= (v3 % (((^(v3))) | 1))) {
                _ = len(v4)
                switch -1 {
                case -976:
                    var v20 uint16 = 0
                    _ = v20
                case 42:
                    if false {
                        if (((!true) || false) || (!true)) {
                            v6 = v6
                            v1 = 1
                            var v21 uint32 = 585
                            _ = v21
                        }
                        v3 = 9223372036854775807
                    } else {
                        v1 = 42
                        v1 -= (v1 & 1)
                    }
                case 32768:
                    var v22 int32 = (-(v1))
                    _ = v22
                default:
                    var v23 string = v6
                    _ = v23
                }
                var v24 float64 = float64((float32(float64((-2.25 / 1.5))) + ((1e10 + 0.1) - 0.1)))
                _ = v24
            }
            var v25 int16 = 409
            _ = v25
        } else {
            var v26 int8 = 42
            v27 := GenS4{F0: 0, F1: 4294967295, F2: 4294967295}
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 int16 = -233
            _ = fi28
            _ = v29
        }
        switch v3 {
        case -1:
            v3 = v3
        case 128:
            if false {
                var v30 uint64 = 704
                _ = v30
            }
        case -503:
            v6 = v6
        }
    } else {
        v31 := GenS5{F0: 0, F1: ("a" + (v6 + "foo bar"))}
        v32 := GenS1{F0: 65535}
        var v33 string = (v6 + v6)
        _ = v31
        _ = v32
        _ = v33
    }
    var v34 float64 = 3.14159
    var v35 uint64 = 4294967295
    fmt.Println(v34, v1)
    var v36 bool = (false && (!((-2.25 <= -2.25) && ((3.14159 > (-2.25 + 3.14159)) || (!true)))))
    switch "\n\t" {
    case "a":
        if (2 < 0) {
            fmt.Println(v35)
        }
    default:
        v37 := GenS6{F0: (255 < 4294967295), F1: -1}
        _ = v37
    }
    v36 = true
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 int32 = (v1 / ((v1) | 1))
        _ = fi38
        _ = v39
    }
    v35 += 976
    fmt.Println(v35, v3)
    var v40 int32 = (-(v1))
    if (!v36) {
        v2[("foo bar" + "a")] = v3
    } else {
        var v41 int64 = 1
        var v42 uint16 = 287
        _ = v41
        _ = v42
    }
    fmt.Println(v35)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
