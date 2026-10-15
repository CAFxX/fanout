package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int32 }
func genFn0() int32 {
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    var v8 int32 = 1
    _ = v8
    return v8
}
func genFn1(p0 uint16, p1 uint8) int64 {
    fmt.Println(p0)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 1
    var v3 float64 = 3.14159
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    v6 := make(map[uint32]uint32)
    v9 := GenS1{F0: 9223372036854775807}
    v10 := GenS1{F0: -352}
    if (1.5 > (1e10 / float32(v3))) {
        if true {
            for fi11 := 0; fi11 < 4; fi11++ {
                v3 += v3
                var v12 int16 = 1000
                _ = fi11
                _ = v12
            }
            fmt.Println(v2)
            var v13 uint = 0
            _ = v13
        }
        fmt.Println(v3, v2)
        if false {
            if (162 == 2147483647) {
                v3 = float64(1.5)
                fmt.Println(v2, v3)
            } else {
                var v14 float64 = 3.14159
                _ = v14
            }
            switch "hello" {
            case "hello":
                fmt.Println(v3, v2)
            case "a":
                v3 += 1.5
            }
            v15 := []int{42, -536}
            _ = v15
        }
    }
    if true {
        var v16 float64 = ((0.1 * v3) / 1.5)
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int64 = 9223372036854775807
            var v19 uint8 = 2
            switch 199 {
            case 0:
                var v20 uint32 = 42
                _ = v20
            case -2147483648:
                for fi21 := 0; fi21 < 18; fi21++ {
                    v22 := []int{(fi17 | fi17)}
                    switch (("" + (("a" + "hello") + ("hello" + ""))) + "\n\t") {
                    case "":
                        v19 -= 101
                    }
                    _ = len(v22)
                    _ = fi21
                    _ = v22
                }
            default:
                v16 = float64(0.1)
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v16
    }
    v23 := GenS1{F0: 860}
    switch ("a" + (("foo bar" + "") + "a")) {
    case "foo bar":
        var v24 float64 = 3.14159
        _ = v24
    case "":
        fmt.Println(v2)
    default:
        for fi25 := 0; fi25 < 16; fi25++ {
            v3 = -2.25
            _ = fi25
        }
    }
    fmt.Println(v3)
    v26 := GenS2{F0: 134}
    v3 -= ((float64(0.0) / v3) - ((v3 - (v3 - ((v3 - 1e10) + 1e10))) * v3))
    v27 := GenS3{F0: v2, F1: 32768}
    var v28 int = 4294967295
    if true {
        for fi29 := 0; fi29 < 3; fi29++ {
            switch (fi29 ^ (^(fi29))) {
            case 9223372036854775807:
                for fi30 := 0; fi30 < 19; fi30++ {
                    v31 := GenS4{F0: 2147483647}
                    _ = fi30
                    _ = v31
                }
            default:
                var v32 int64 = -1
                _ = v32
            }
            if (true && true) {
                fmt.Println(v3)
                var v33 int64 = -1
                fmt.Println(v33)
                _ = v33
            } else {
                if true {
                    v34 := GenS2{F0: 255}
                    _ = v34
                } else {
                    v35 := GenS3{F0: (v2 % ((v2) | 1)), F1: 783}
                    v3 += (float64(float32(-2.25)) * (0.0 + v3))
                    _ = v35
                }
                v36 := GenS2{F0: 1}
                v2 = v2
                _ = v36
            }
            _ = fi29
        }
        v3 = (v3 / ((float64(1.5) - (float64(3.14159) - ((v3 / 1.5) / v3))) - v3))
        fmt.Println(v3, v28, v28)
    } else {
        v2 = (v2 & (v2 | (^(v2))))
        switch 32767 {
        case 4294967295:
            var v37 string = "hello"
            _ = v37
        default:
            var v38 uint8 = 255
            _ = v38
        }
    }
    v3 += v3
    v39 := []int{v28}
    v28 = v28
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
