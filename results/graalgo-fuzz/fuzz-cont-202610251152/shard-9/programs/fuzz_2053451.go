package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
func genFn0(p0 bool) uint64 {
    if p0 {
        for fi5 := 0; fi5 < 13; fi5++ {
            switch "foo bar" {
            case "hello":
                var v6 int16 = 1
                _ = v6
            case "":
                p0 = p0
            }
            _ = fi5
        }
        p0 = (9223372036854775807 >= 1)
    }
    if p0 {
        for fi7 := 0; fi7 < 17; fi7++ {
            if (-1 <= 1) {
                v8 := GenS1{F0: 353, F1: 779}
                _ = v8
            } else {
                p0 = p0
            }
            switch "a" {
            case "":
                fmt.Println(p0, fi7)
            }
            if ((!p0) || ((p0 && (!(p0 || p0))) && (!true))) {
                var v9 int = (fi7 - -2147483648)
                _ = v9
            }
            _ = fi7
        }
    } else {
        if p0 {
            p0 = (p0 && (193 > 65535))
        }
        var v10 float32 = 1.5
        var v11 int32 = -2147483648
        _ = v10
        _ = v11
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    v3 := []uint32{992, 42, 42}
    var v4 int64 = 9223372036854775807
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int16 = 336
        for fi14 := 0; fi14 < 8; fi14++ {
            v4 -= 0
            if (932 < 2147483648) {
                switch fi14 {
                case 9223372036854775807:
                    v15 := GenS1{F0: 2147483647, F1: 1}
                    _ = v15
                case 539:
                    var v16 uint = 1
                    _ = v16
                case -1:
                    v17 := GenS1{F0: 2147483647, F1: 4294967295}
                    var v18 uint32 = 118
                    _ = v17
                    _ = v18
                }
                if ((((false || (255 >= 0)) || false) && (!(128 >= 42))) && true) {
                    _ = len(v1)
                }
            } else {
                var v19 uint32 = 2
                _ = v19
            }
            var v20 int8 = -78
            _ = fi14
            _ = v20
        }
        v3 = append(v3, 65535)
        _ = fi12
        _ = v13
    }
    v4 -= (v4 ^ v4)
    v21 := GenS1{F0: 0, F1: 0}
    if (false || false) {
        if true {
            if true {
                switch 128 {
                case 2147483647:
                    var v22 uint32 = 2
                    _ = v22
                case 0:
                    v23 := GenS2{F0: -757}
                    _ = v23
                case 127:
                    fmt.Println(v4)
                default:
                    var v24 uint = 0
                    _ = v24
                }
                switch "hello" {
                case "foo bar":
                    switch 4294967295 {
                    case 128:
                        v4 += (v4 % ((v4) | 1))
                    }
                case "\n\t":
                    v4 = v4
                case "hello":
                    var v25 int = -296
                    _ = v25
                default:
                    delete(v1, "foo bar")
                }
                v4 = v4
            } else {
                v4 += v4
            }
        }
        for fi26 := 0; fi26 < 16; fi26++ {
            fmt.Println(fi26)
            for fi27 := 0; fi27 < 7; fi27++ {
                v4 = -1
                if false {
                    v28 := GenS2{F0: fi26}
                    var v29 float32 = 1.5
                    fmt.Println(v4, fi26)
                    _ = v28
                    _ = v29
                }
                _ = fi27
            }
            if (2 < 255) {
                v4 -= (-(v4))
                for fi30 := 0; fi30 < 18; fi30++ {
                    fmt.Println(fi30, fi30, fi30)
                    v4 -= (v4 | v4)
                    _ = len(v3)
                    _ = fi30
                }
                v4 = v4
            }
            _ = fi26
        }
    } else {
        v31 := GenS2{F0: -677}
        var v32 int8 = 127
        _ = v31
        _ = v32
    }
    v33 := GenS3{F0: "\n\t"}
    var v34 float32 = (3.14159 - 1e10)
    v4 += v4
    switch 42 {
    case -2147483648:
        var v35 int32 = 1
        _ = v35
    default:
        var v36 int16 = -32768
        _ = v36
    }
    switch ("hello" + "foo bar") {
    case "":
        var v37 string = (("" + "a") + "\n\t")
        _ = v37
    }
    fmt.Println(v4)
    fmt.Println(v34, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
