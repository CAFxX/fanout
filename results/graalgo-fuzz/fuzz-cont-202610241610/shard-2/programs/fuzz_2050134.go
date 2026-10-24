package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 string }
type GenS3 struct { F0 int; F1 int64; F2 uint16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint64) int {
    switch "foo bar" {
    case "hello":
        for fi4 := 0; fi4 < 7; fi4++ {
            var v5 float32 = -2.25
            v6 := []int{(^(fi4))}
            p0 += 685
            _ = fi4
            _ = v5
            _ = v6
        }
    case "\n\t":
        var v7 uint8 = 127
        _ = v7
    case "foo bar":
        fmt.Println(p0)
    }
    v8 := GenS1{F0: 684, F1: -1, F2: 32767}
    if (42 > 930) {
        var v9 int32 = 0
        _ = v9
    } else {
        var v10 int64 = 42
        fmt.Println(p0, v10)
        v11 := []int{}
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := []int{(fi12 - -1)}
        _ = fi12
        _ = v13
    }
    var v14 float64 = (1e10 - 3.14159)
    _ = p0
    _ = v14
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 int = -805
    var v15 int32 = 261
    v15 += 42
    switch (("" + "a") + "a") {
    case "\n\t":
        fmt.Println(v3, v15)
    case "a":
        v15 = v15
    default:
        v3 -= (v3 | 42)
    }
    if (false || true) {
        v16 := []int{-9223372036854775808, 127, 1}
        if (!false) {
            var v17 float64 = (3.14159 + 3.14159)
            _ = v17
        }
        if (42 == 790) {
            var v18 int16 = -32768
            v19 := GenS2{F0: 664, F1: 1, F2: "a"}
            _ = v18
            _ = v19
        }
        _ = v16
    } else {
        if true {
            switch "a" {
            case "\n\t":
                for fi20 := 0; fi20 < 19; fi20++ {
                    v15 -= (v15 % (((v15 + int32(fi20))) | 1))
                    fmt.Println(v15, fi20)
                    for fi21 := 0; fi21 < 8; fi21++ {
                        var v22 float32 = 1e10
                        var v23 uint32 = 57
                        v15 += v15
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    _ = fi20
                }
            }
            v24 := GenS3{F0: v3, F1: -1, F2: 255}
            for fi25 := 0; fi25 < 7; fi25++ {
                v15 -= (v15 % ((2147483647) | 1))
                _ = fi25
            }
            _ = v24
        } else {
            var v26 uint64 = 2
            if ((("a" + "") >= ("\n\t" + "a")) || false) {
                v1[2147483648] = ("hello" > "")
                if ((v3 <= v3) || ((v15 - 2147483647) > v15)) {
                    fmt.Println(v15)
                } else {
                    var v27 int = 0
                    delete(v2, ("foo bar" + ""))
                    var v28 int8 = 127
                    _ = v27
                    _ = v28
                }
            }
            _ = v26
        }
        v29 := GenS4{F0: -1}
        var v30 int = (v3 % ((-710) | 1))
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 float64 = 0.1
        v33 := []int{}
        var v34 float32 = -2.25
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    v35 := GenS2{F0: 0, F1: 127, F2: (("\n\t" + "foo bar") + ("\n\t" + "hello"))}
    _ = len(v1)
    if true {
        v36 := []int{v3, v3, (v3 - v3)}
        var v37 uint16 = 876
        for fi38 := 0; fi38 < 11; fi38++ {
            v1[0] = ((true && true) && true)
            _ = fi38
        }
        _ = v36
        _ = v37
    } else {
        v39 := GenS2{F0: 418, F1: 128, F2: "foo bar"}
        _ = v39
    }
    if (65535 == 834) {
        v3 = v3
        switch v3 {
        case -9223372036854775808:
            var v40 float64 = ((0.0 * (3.14159 * (((3.14159 + -2.25) * 0.0) * 0.1))) * (float64(0.1) / -2.25))
            _ = v40
        case 2147483647:
            var v41 int32 = 2147483647
            _ = v41
        }
        for fi42 := 0; fi42 < 6; fi42++ {
            var v43 uint32 = 2147483648
            _ = fi42
            _ = v43
        }
    } else {
        for fi44 := 0; fi44 < 6; fi44++ {
            if (false || true) {
                var v45 float32 = (-2.25 - (0.0 + 1e10))
                if false {
                    fmt.Println(v3)
                    v46 := []int{fi44}
                    _ = v46
                } else {
                    var v47 uint32 = 0
                    _ = v47
                }
                _ = v45
            }
            _ = fi44
        }
        delete(v1, 207)
        if (2147483647 >= 65535) {
            v15 = 0
        }
    }
    v1[177] = false
    for fi48 := 0; fi48 < 17; fi48++ {
        v15 -= -1
        _ = fi48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
}
