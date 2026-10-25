package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int }
type GenS2 struct { F0 int; F1 bool; F2 int16 }
type GenS3 struct { F0 int; F1 bool; F2 int }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint64) uint8 {
    p0 -= 42
    switch 627 {
    case 42:
        for fi4 := 0; fi4 < 9; fi4++ {
            switch 264 {
            case 9223372036854775807:
                v5 := []int{(^(fi4)), fi4}
                _ = v5
            case 42:
                p0 = (p0 - p0)
            default:
                fmt.Println(fi4, p0)
            }
            if true {
                v6 := []int{fi4, (fi4 ^ (fi4 * (^(fi4))))}
                var v7 int8 = -128
                _ = v6
                _ = v7
            } else {
                if false {
                    v8 := GenS1{F0: 173, F1: ("a" + "a"), F2: (fi4 + fi4)}
                    var v9 int16 = -844
                    fmt.Println(p0, fi4)
                    _ = v8
                    _ = v9
                }
                fmt.Println(fi4, p0)
                fmt.Println(fi4)
            }
            var v10 uint32 = 1
            _ = fi4
            _ = v10
        }
    case 895:
        if false {
            if (-677 == 42) {
                var v11 uint8 = 128
                var v12 uint64 = (^(p0))
                _ = v11
                _ = v12
            } else {
                v13 := []int{}
                p0 = p0
                if len(v13) > 0 {
                    v13[0] = -538
                }
                _ = v13
            }
            v14 := GenS2{F0: 128, F1: true, F2: 0}
            _ = v14
        } else {
            v15 := []int{0, 128}
            _ = v15
        }
    case -1:
        fmt.Println(p0)
    default:
        if true {
            var v16 uint32 = 1
            var v17 float64 = (3.14159 - 0.1)
            v18 := GenS1{F0: 127, F1: "foo bar", F2: -698}
            _ = v16
            _ = v17
            _ = v18
        }
    }
    _ = p0
    return 128
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (9223372036854775807 >= 32767)
    var v3 string = ((("foo bar" + "a") + "") + "hello")
    v19 := GenS1{F0: genFn0(77), F1: "foo bar", F2: 32768}
    var v20 uint64 = 702
    var v21 float64 = ((3.14159 - 3.14159) - 0.0)
    v21 += v21
    if (v2 || v2) {
        v22 := GenS1{F0: 0, F1: ("" + (v3 + v3)), F2: -9223372036854775808}
        _ = v22
    } else {
        v23 := GenS2{F0: 785, F1: (169 <= 42), F2: 1000}
        v24 := GenS1{F0: 255, F1: (("foo bar" + v3) + ""), F2: 65}
        if v2 {
            for fi25 := 0; fi25 < 9; fi25++ {
                v1[4294967295] = (fi25 % ((fi25) | 1))
                _ = fi25
            }
            v26 := GenS3{F0: -2147483648, F1: v2, F2: 4294967295}
            _ = v26
        } else {
            v27 := GenS1{F0: genFn0((v20 | 0)), F1: (v3 + (v3 + "")), F2: 0}
            v2 = ((100 >= -1) && true)
            _ = v27
        }
        _ = v23
        _ = v24
    }
    if false {
        fmt.Println(v3)
        switch "foo bar" {
        case "a":
            if ("" <= v3) {
                v2 = true
                v28 := GenS1{F0: 231, F1: "", F2: -683}
                var v29 string = v3
                _ = v28
                _ = v29
            }
        case "\n\t":
            v3 = "a"
        case "hello":
            if v2 {
                v3 = ((v3 + (v3 + v3)) + v3)
            } else {
                var v30 string = v3
                _ = v30
            }
        default:
            v31 := []int{42}
            _ = v31
        }
        if (v20 < (v20 ^ (v20 % (((v20 - v20)) | 1)))) {
            v1[-2147483648] = 32768
        } else {
            var v32 uint = 65535
            v2 = false
            v3 = "\n\t"
            _ = v32
        }
    }
    v3 = v3
    for fi33 := 0; fi33 < 12; fi33++ {
        _ = len(v1)
        v34 := GenS4{F0: 145}
        _ = fi33
        _ = v34
    }
    var v35 uint64 = (v20 >> ((1) & 63))
    v21 = ((v21 * float64(float32((float64(1e10) / 1.0)))) + float64(0.1))
    var v36 uint64 = v20
    v37 := GenS4{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
