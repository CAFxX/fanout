package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint64 }
func genFn0() string {
    v4 := []int{1, -585, 9223372036854775807}
    var v5 bool = false
    _ = v4
    _ = v5
    return (("foo bar" + ((("" + ("a" + "foo bar")) + ("hello" + "foo bar")) + ((("" + "") + "foo bar") + "foo bar"))) + (("foo bar" + ("" + "a")) + ("\n\t" + "a")))
}

func main() {
    v1 := []uint32{2}
    var v2 float64 = ((3.14159 + 1.5) + float64((0.1 * 0.0)))
    var v3 int32 = 42
    for fi6 := 0; fi6 < 19; fi6++ {
        fmt.Println(v3, fi6)
        var v7 uint = 4294967295
        fmt.Println(v3, v3, fi6)
        _ = fi6
        _ = v7
    }
    switch ("hello" + "foo bar") {
    case "":
        var v8 bool = false
        _ = v8
    default:
        var v9 int64 = 1
        _ = v9
    }
    switch -93 {
    case 128:
        fmt.Println(v2, v3)
    case 127:
        var v10 float32 = (float32(((v2 / v2) * (3.14159 / (v2 - v2)))) - 1e10)
        _ = v10
    }
    switch "" {
    case "foo bar":
        switch "" {
        case "":
            var v11 int64 = -506
            _ = v11
        case "a":
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := GenS1{F0: -308}
                v2 = 1e10
                _ = fi12
                _ = v13
            }
        case "hello":
            for fi14 := 0; fi14 < 3; fi14++ {
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 uint32 = 1
                    var v17 bool = true
                    fmt.Println(v17, v2, fi14)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                var v18 int64 = 0
                _ = fi14
                _ = v18
            }
        }
    }
    v3 -= v3
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := GenS1{F0: 137}
        switch "foo bar" {
        case "a":
            if (!(0 != 588)) {
                for fi21 := 0; fi21 < 3; fi21++ {
                    switch -1 {
                    case 128:
                        v22 := GenS2{F0: (fi21 << ((661) & 63)), F1: (fi19 | (-(fi19)))}
                        _ = v22
                    case 143:
                        fmt.Println(fi21)
                    default:
                        var v23 uint32 = 65535
                        _ = v23
                    }
                    var v24 int64 = -1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi21
                    _ = v24
                }
                if true {
                    v2 = v2
                } else {
                    var v25 int = 32767
                    fmt.Println(v3)
                    _ = len(v1)
                    _ = v25
                }
                v26 := GenS2{F0: 32768, F1: -484}
                _ = v26
            }
        default:
            v27 := GenS3{F0: true, F1: 1}
            _ = v27
        }
        fmt.Println(v3, v2)
        _ = fi19
        _ = v20
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        if (1 >= 42) {
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 int16 = 0
                v2 += v2
                _ = fi29
                _ = v30
            }
        } else {
            v31 := GenS1{F0: -32768}
            var v32 bool = ((-(v3)) > v3)
            _ = v31
            _ = v32
        }
        _ = fi28
    }
    v33 := GenS4{F0: 4294967295, F1: -1, F2: 1}
    for fi34 := 0; fi34 < 7; fi34++ {
        v2 = v2
        v1 = append(v1, 2147483647)
        _ = fi34
    }
    v3 = 1
    var v35 float64 = (v2 * v2)
    var v36 int8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
}
