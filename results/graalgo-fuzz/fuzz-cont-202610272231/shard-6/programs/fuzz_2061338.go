package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 int64; F2 uint32 }

func main() {
    var v1 int16 = -674
    var v2 uint = 4294967295
    v3 := make(map[int]bool)
    var v4 int = -1
    var v5 float32 = 1e10
    v6 := make(map[string]string)
    v4 += (v4 + (v4 / (((^(v4))) | 1)))
    v4 = (v4 % ((0) | 1))
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            switch (fi7 - 0) {
            case 32767:
                var v9 float32 = 0.0
                v1 += 32767
                v9 = v5
                var v10 int16 = v1
                _ = v10
                _ = v9
            case 127:
                if (!false) {
                    switch (-(v4)) {
                    case -698:
                        var v11 uint32 = 1
                        _ = v11
                    }
                    delete(v3, (v4 << ((v2) & 63)))
                    if (v1 >= v1) {
                        var v12 int64 = 897
                        var v13 uint = 18446744073709551615
                        v14 := GenS1{F0: fi7}
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        v15 := GenS2{F0: 1}
                        if true {
                            var v16 float64 = 1.5
                            v2 += v2
                            v17 := GenS3{F0: (fi8 | 2147483647), F1: 42}
                            _ = v16
                            _ = v17
                        } else {
                            var v18 uint64 = 731
                            if (922 >= 1) {
                                v4 = 4294967295
                            }
                            _ = v18
                        }
                        _ = v15
                    }
                } else {
                    var v19 int32 = 2147483647
                    switch "hello" {
                    case "foo bar":
                        v20 := GenS4{F0: fi7, F1: (!false)}
                        _ = v20
                    }
                    v19 = v19
                    _ = v19
                }
            case -4:
                switch 42 {
                case 491:
                    v21 := []int{(fi8 % (((^(fi8))) | 1)), fi8, fi8}
                    _ = v21
                }
            }
            _ = fi8
        }
        v22 := []int{}
        _ = fi7
        _ = v22
    }
    if (0 == 42) {
        v1 = v1
        if true {
            v23 := GenS4{F0: -561, F1: true}
            _ = v23
        } else {
            var v24 int = (v4 - 128)
            _ = v24
        }
        var v25 int64 = 1
        _ = v25
    }
    var v26 bool = (!true)
    fmt.Println(v2)
    switch (("hello" + ("a" + ("\n\t" + "\n\t"))) + ("foo bar" + "hello")) {
    case "foo bar":
        v27 := GenS1{F0: v4}
        _ = v27
    case "":
        v28 := GenS5{F0: 255, F1: 1}
        _ = v28
    }
    v29 := []int{(^(v4))}
    fmt.Println(v5, v26, v4)
    if (!(v5 > (3.14159 * v5))) {
        v30 := GenS1{F0: v4}
        v29 = append(v29, (v4 << ((v2) & 63)))
        switch (v4 << ((1) & 63)) {
        case 600:
            var v31 int32 = 42
            _ = v31
        case 1:
            if (42 > 996) {
                v32 := GenS6{F0: v2}
                _ = v32
            } else {
                if (1 == 42) {
                    var v33 uint64 = 27
                    if (true || ((v26 && v26) || v26)) {
                        if (1e10 <= 0.1) {
                            v2 = uint(v33)
                        }
                        var v34 uint32 = 4294967295
                        _ = v34
                    }
                    var v35 uint64 = (v33 - v33)
                    _ = v33
                    _ = v35
                }
            }
        default:
            switch "foo bar" {
            case "\n\t":
                v1 -= v1
            case "foo bar":
                if (v4 >= v4) {
                    v36 := GenS7{F0: 100}
                    var v37 int16 = v1
                    delete(v6, "a")
                    _ = v36
                    _ = v37
                }
            case "":
                v2 = (v2 - 2)
            }
        }
        _ = v30
    } else {
        v5 = -2.25
    }
    v38 := GenS4{F0: (v4 ^ (v4 >> (((v2 + v2)) & 63))), F1: ("hello" == ("foo bar" + "a"))}
    for fi39 := 0; fi39 < 20; fi39++ {
        fmt.Println(v1, v26)
        var v40 int16 = -654
        if (255 == 255) {
            if (((1e10 / 0.1) < 1.5) || v26) {
                var v41 uint = v2
                _ = v41
            }
            _ = len(v6)
        } else {
            v42 := GenS3{F0: fi39, F1: 128}
            v43 := GenS8{F0: 1000, F1: -9223372036854775808, F2: 189}
            var v44 float32 = v5
            _ = v42
            _ = v43
            _ = v44
        }
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
