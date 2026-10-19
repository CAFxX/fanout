package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS5 struct { F0 string; F1 int16 }
func genFn0(p0 uint, p1 int8, p2 bool) string {
    var v5 uint64 = 2
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 664, F1: 65535}
        p2 = (!p2)
        _ = fi6
        _ = v7
    }
    v8 := []int{-2147483648, -601, 2147483647}
    v8 = append(v8, 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return "foo bar"
}
func genFn1(p0 int64, p1 int32, p2 int32) int16 {
    v9 := GenS2{F0: 59}
    v10 := []int{0, 502}
    p1 = p1
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS3{F0: -1, F1: p2, F2: 0}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 float64 = (1e10 - (0.0 + 1.5))
    v4 := make(map[string]string)
    var v13 uint8 = 255
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int = (^(fi14))
        if false {
            fmt.Println(v13, v15)
        }
        v15 += (v15 & 32768)
        _ = fi14
        _ = v15
    }
    fmt.Println(v13)
    v1["hello"] = ("hello" + "hello")
    if ((-128 == 100) && (!false)) {
        v16 := GenS1{F0: 32767, F1: 823}
        delete(v2, 368)
        fmt.Println(v3, v13)
        _ = v16
    } else {
        switch "a" {
        case "foo bar":
            v17 := GenS3{F0: 42, F1: 2147483647, F2: 1}
            _ = v17
        default:
            switch 4294967295 {
            case 127:
                v1[""] = ("a" + "a")
            case 195:
                var v18 uint64 = 2
                _ = v18
            }
        }
        var v19 int8 = 42
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 int16 = -275
            v22 := GenS1{F0: 2, F1: 2147483648}
            fmt.Println(fi20, fi20, v13)
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v19
    }
    switch -2147483648 {
    case 42:
        v3 = v3
    case 400:
        if true {
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 float64 = (v3 * float64(3.14159))
                _ = fi23
                _ = v24
            }
            if (-343 >= 0) {
                var v25 string = "foo bar"
                v26 := GenS4{F0: 65535, F1: 1, F2: 42}
                _ = v25
                _ = v26
            } else {
                var v27 string = "hello"
                _ = v27
            }
            if false {
                var v28 float64 = (v3 * v3)
                _ = v28
            }
        }
    }
    v29 := GenS3{F0: 42, F1: 2147483647, F2: 42}
    var v30 bool = (370 < 504)
    if v30 {
        var v31 int = 32768
        _ = v31
    } else {
        if true {
            var v32 float32 = -2.25
            switch -9223372036854775808 {
            case -764:
                fmt.Println(v3, v3)
            case -2147483648:
                v3 += v3
            case -67:
                v33 := GenS5{F0: "foo bar", F1: -32768}
                _ = v33
            default:
                fmt.Println(v30, v30)
            }
            _ = v32
        }
    }
    v34 := GenS3{F0: 9223372036854775807, F1: 42, F2: uint64(v13)}
    var v35 int8 = -1
    switch "a" {
    case "hello":
        v35 = v35
    }
    if (32767 <= 1) {
        v13 -= v13
        v13 += 42
        v30 = v30
    } else {
        for fi36 := 0; fi36 < 19; fi36++ {
            var v37 uint = 689
            _ = fi36
            _ = v37
        }
        fmt.Println(v35)
        var v38 string = ((("foo bar" + ("foo bar" + "")) + "hello") + "a")
        _ = v38
    }
    if v30 {
        if (v35 > v35) {
            var v39 uint16 = 65535
            _ = v39
        } else {
            switch "" {
            case "foo bar":
                switch ("a" + ("\n\t" + "a")) {
                case "\n\t":
                    var v40 uint = 65535
                    _ = v40
                case "hello":
                    fmt.Println(v3, v35)
                }
            default:
                v30 = (1 >= -32768)
            }
        }
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
}
