package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint32; F1 int8 }
func genFn0() int {
    v4 := GenS1{F0: 42, F1: -1}
    if true {
        if true {
            if (true || (215 != 0)) {
                v5 := GenS1{F0: 0, F1: -2147483648}
                v6 := []int{4}
                switch 78 {
                case 428:
                    for fi7 := 0; fi7 < 7; fi7++ {
                        fmt.Println(fi7)
                        v8 := GenS1{F0: -9223372036854775808, F1: 4294967295}
                        var v9 uint32 = 436
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                case 673:
                    var v10 int = 576
                    _ = v10
                default:
                    _ = len(v6)
                }
                _ = v5
                _ = v6
            }
            var v11 bool = (true && (!false))
            _ = v11
        }
        if false {
            v12 := []int{-859, 395, 42, 127}
            var v13 int64 = 9223372036854775807
            v14 := GenS2{F0: 127}
            _ = v12
            _ = v13
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(fi15)
            _ = fi15
        }
        switch 128 {
        case 127:
            if true {
                if (!(false && false)) {
                    var v16 uint64 = 0
                    v17 := GenS2{F0: 32768}
                    var v18 int32 = 0
                    _ = v16
                    _ = v17
                    _ = v18
                }
                for fi19 := 0; fi19 < 5; fi19++ {
                    var v20 float64 = 1e10
                    v21 := []int{(fi19 * (fi19 & (fi19 >> ((42) & 63))))}
                    _ = len(v21)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        case 42:
            var v22 int16 = 42
            _ = v22
        case 4294967295:
            v23 := GenS1{F0: 0, F1: 494}
            _ = v23
        }
        var v24 uint8 = 117
        _ = v24
    }
    if (false || (true && false)) {
        var v25 float32 = (((3.14159 - 3.14159) + 1.5) * ((((1.5 - float32(0.1)) - 0.1) + 0.0) / 1e10))
        v25 = 3.14159
        _ = v25
    }
    var v26 int16 = -564
    v26 += v26
    _ = v26
    _ = v4
    return -594
}
func genFn1(p0 uint64, p1 int16) uint {
    fmt.Println(p0, p1)
    p1 -= 32767
    p1 -= (p1 - (p1 & (p1 * p1)))
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 1
    var v3 uint16 = 728
    fmt.Println(v3)
    var v27 uint8 = 49
    var v28 uint8 = (v27 & (v27 / (((^(v27))) | 1)))
    if (-2.25 <= 0.0) {
        var v29 int8 = 0
        _ = v29
    }
    switch ("\n\t" + "hello") {
    case "hello":
        switch "a" {
        case "":
            switch ((("a" + "foo bar") + "") + "foo bar") {
            case "":
                var v30 float32 = -2.25
                _ = v30
            case "\n\t":
                v2 = 32767
            }
        }
    case "\n\t":
        v31 := GenS3{F0: v28, F1: "a"}
        _ = v31
    case "foo bar":
        var v32 int32 = -2147483648
        _ = v32
    default:
        var v33 int32 = 833
        _ = v33
    }
    delete(v1, ("hello" + ("a" + "hello")))
    for fi34 := 0; fi34 < 3; fi34++ {
        fmt.Println(v2, v2, v3)
        _ = fi34
    }
    var v35 float32 = (3.14159 / 1e10)
    var v36 uint16 = v2
    v37 := GenS4{F0: 2147483648, F1: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
