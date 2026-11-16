package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int32 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint64 }
func genFn0(p0 uint64) int32 {
    if (257 <= 0) {
        var v7 bool = true
        _ = v7
    } else {
        p0 = p0
        p0 = p0
        v8 := []int{}
        _ = v8
    }
    var v9 int = 32767
    _ = p0
    _ = v9
    return 42
}
func genFn1(p0 int16, p1 int64, p2 int64) uint16 {
    var v10 int64 = -1
    if (true && (42 < 65535)) {
        v11 := GenS1{F0: p0, F1: (("foo bar" + "a") + ("foo bar" + ("" + ("hello" + ((((("\n\t" + "a") + ("hello" + "foo bar")) + "a") + (("hello" + "") + ("\n\t" + ("foo bar" + "a")))) + ("hello" + "a")))))), F2: -834}
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS1{F0: 32767, F1: ("" + ("a" + ("foo bar" + ((("hello" + "foo bar") + "a") + "")))), F2: -2147483648}
            v14 := []int{9223372036854775807, int(p0), fi12, 4294967295}
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    } else {
        var v15 string = ("a" + (("\n\t" + "hello") + "a"))
        _ = v15
    }
    v16 := []int{-682, 811, 551, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v16
    return 0
}

func main() {
    var v1 uint = 36
    var v2 int = -9223372036854775808
    var v3 int16 = 588
    var v4 string = "hello"
    var v5 int16 = v3
    var v6 int = -779
    var v17 uint64 = 604
    for fi18 := 0; fi18 < 16; fi18++ {
        v17 = v17
        fmt.Println(v17)
        _ = fi18
    }
    v19 := GenS2{F0: 410, F1: (v5 / ((-1) | 1)), F2: v17}
    if true {
        for fi20 := 0; fi20 < 19; fi20++ {
            switch (("foo bar" + v4) + v4) {
            case "":
                for fi21 := 0; fi21 < 10; fi21++ {
                    v22 := GenS2{F0: 0, F1: (v5 | int16(v6)), F2: v17}
                    v23 := GenS2{F0: 1, F1: (v3 * (^(v3))), F2: v17}
                    var v24 uint64 = (v17 / ((v17) | 1))
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            case "hello":
                fmt.Println(v1, v4)
            }
            fmt.Println(v4)
            v17 -= v17
            _ = fi20
        }
        for fi25 := 0; fi25 < 11; fi25++ {
            switch "" {
            case "a":
                v6 = 38
            case "hello":
                var v26 string = v4
                _ = v26
            }
            _ = fi25
        }
        switch (v2 / ((127) | 1)) {
        case -9223372036854775808:
            var v27 uint = v1
            _ = v27
        case -774:
            v2 -= v6
        case -132:
            v1 = v1
        default:
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 uint32 = 2147483648
                _ = fi28
                _ = v29
            }
        }
    } else {
        v3 -= (v3 & 42)
        for fi30 := 0; fi30 < 7; fi30++ {
            v31 := []int{fi30, (fi30 - (v6 & int(v1))), (fi30 << (((v1 * (v1 & (v1 / ((v1) | 1))))) & 63))}
            _ = fi30
            _ = v31
        }
    }
    fmt.Println(v6, v4, v17)
    v3 = 1000
    var v32 int8 = 0
    var v33 uint32 = uint32(v1)
    var v34 float32 = (((0.0 * 1e10) / (-2.25 + (0.0 + (1.5 / 0.1)))) - -2.25)
    v35 := GenS1{F0: v3, F1: ((v4 + v4) + (("hello" + (v4 + "a")) + "a")), F2: genFn0(868)}
    v36 := []int{v6, 658, (v6 | (v2 / (((v2 - (-(v6)))) | 1)))}
    var v37 uint16 = 0
    fmt.Println(v3, v33)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
