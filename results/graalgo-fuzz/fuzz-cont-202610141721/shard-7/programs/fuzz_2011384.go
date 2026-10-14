package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 bool; F1 int32; F2 uint8 }
func genFn0(p0 uint) int16 {
    if ((true && (true || (32767 >= 2))) && (1 == 0)) {
        if true {
            var v6 int64 = 499
            v7 := []int{-1}
            _ = v6
            _ = v7
        } else {
            switch (("" + "hello") + "\n\t") {
            case "foo bar":
                v8 := GenS1{F0: 705, F1: 534}
                _ = v8
            case "":
                if (!true) {
                    v9 := []int{}
                    switch "a" {
                    case "a":
                        p0 += p0
                    }
                    v10 := GenS1{F0: 2147483647, F1: -2147483648}
                    _ = v10
                    _ = v9
                }
            default:
                for fi11 := 0; fi11 < 3; fi11++ {
                    p0 = 665
                    _ = fi11
                }
            }
            v12 := []int{-551}
            _ = v12
        }
        var v13 uint = p0
        _ = v13
    } else {
        switch ((("foo bar" + "hello") + "foo bar") + "") {
        case "a":
            var v14 float32 = (0.1 - -2.25)
            _ = v14
        }
        v15 := []int{127}
        _ = v15
    }
    var v16 bool = (!(-2.25 == (-2.25 * (3.14159 / -2.25))))
    _ = p0
    _ = v16
    return 0
}

func main() {
    var v1 bool = (!false)
    var v2 int32 = 0
    v3 := make(map[string]bool)
    var v4 uint8 = 2
    var v5 uint32 = 620
    if ((42 == -32768) || (1 > 0)) {
        var v17 uint16 = 0
        v18 := []int{-9223372036854775808, -942, -9223372036854775808}
        switch -2147483648 {
        case 4294967295:
            var v19 int64 = 9223372036854775807
            _ = v19
        case 9223372036854775807:
            delete(v3, "foo bar")
        case 127:
            var v20 uint8 = (v4 / (((v4 ^ (v4 * 42))) | 1))
            _ = v20
        }
        _ = v17
        _ = v18
    }
    v21 := GenS1{F0: 65535, F1: 0}
    fmt.Println(v2, v2, v1)
    if v1 {
        v22 := v3[""]
        v23 := GenS2{F0: (v1 && (((^(v5)) == v5) && v22)), F1: 2147483647, F2: 127}
        _ = v22
        _ = v23
    }
    var v24 uint = 121
    var v25 uint16 = 779
    var v26 bool = (255 <= v24)
    v4 = (^(v4))
    _ = len(v3)
    fmt.Println(v24)
    v26 = (v1 || v1)
    var v27 uint32 = (v5 - (^(v5)))
    if ((v25 / ((v25) | 1)) == 2) {
        var v28 float64 = -2.25
        for fi29 := 0; fi29 < 10; fi29++ {
            v25 = v25
            switch "\n\t" {
            case "a":
                fmt.Println(v4, v5, v2)
            default:
                v5 += (v5 >> ((v24) & 63))
            }
            v27 += (^(v27))
            _ = fi29
        }
        _ = v28
    } else {
        var v30 int8 = int8(v24)
        _ = v30
    }
    fmt.Println(v25, v2, v1)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
