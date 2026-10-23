package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int32 }
type GenS2 struct { F0 bool; F1 int32; F2 int8 }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0() bool {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint = 170
        _ = fi7
        _ = v8
    }
    switch "hello" {
    case "hello":
        v9 := []int{-19, 0, 127}
        _ = v9
    case "a":
        switch -771 {
        case 1:
            var v10 string = (("hello" + (("\n\t" + "hello") + ("foo bar" + "a"))) + "hello")
            _ = v10
        case 42:
            var v11 uint8 = 2
            _ = v11
        case 896:
            if false {
                var v12 string = "a"
                v13 := GenS1{F0: 314, F1: false, F2: 657}
                v14 := GenS1{F0: 2147483648, F1: (!false), F2: -916}
                _ = v12
                _ = v13
                _ = v14
            }
        default:
            v15 := GenS1{F0: 410, F1: false, F2: 2147483647}
            _ = v15
        }
    }
    switch 32767 {
    case -2147483648:
        var v16 string = "foo bar"
        _ = v16
    }
    if (true || (!false)) {
        var v17 string = "hello"
        _ = v17
    }
    return true
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 0
    var v3 uint32 = 2147483647
    var v4 int = -9223372036854775808
    var v5 float32 = 0.0
    v6 := make(map[uint32]uint32)
    if ((!false) || false) {
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS1{F0: 2147483647, F1: (0.1 < (float32(0.1) + 0.1)), F2: 2147483647}
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 int16 = 742
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := GenS2{F0: false, F1: 1, F2: 1}
            fmt.Println(v5, v4, fi22)
            var v24 uint8 = 165
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    v25 := GenS3{F0: -1, F1: -1}
    for fi26 := 0; fi26 < 5; fi26++ {
        v2 -= (v2 - v2)
        switch (fi26 ^ (fi26 << (((v2 / ((v2) | 1))) & 63))) {
        case -2147483648:
            v27 := []int{fi26, v4}
            _ = v27
        case -531:
            if false {
                var v28 string = (("hello" + ("" + "a")) + "\n\t")
                if (false || false) {
                    fmt.Println(v2)
                } else {
                    v4 += fi26
                }
                v5 -= 1e10
                _ = v28
            }
        case 435:
            var v29 uint64 = 623
            _ = v29
        }
        _ = fi26
    }
    v2 += (v2 & v2)
    for fi30 := 0; fi30 < 19; fi30++ {
        v31 := GenS1{F0: 42, F1: false, F2: 2147483647}
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        v3 = v3
        for fi33 := 0; fi33 < 5; fi33++ {
            v3 = v3
            var v34 int8 = 24
            v35 := v1[""]
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi32
    }
    var v36 int8 = 127
    v3 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
