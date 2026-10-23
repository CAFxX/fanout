package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 bool; F1 int; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int; F2 int32 }
type GenS4 struct { F0 string; F1 uint32; F2 uint }
type GenS5 struct { F0 uint16; F1 int32; F2 int }
func genFn0(p0 int32, p1 uint32) bool {
    var v5 uint = 1
    var v6 uint8 = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 uint) string {
    if true {
        p0 -= p0
        for fi7 := 0; fi7 < 10; fi7++ {
            if (!(!((3.14159 / (1e10 - 1.5)) > float64(0.1)))) {
                p0 = 4294967295
            }
            v8 := GenS1{F0: 1, F1: fi7}
            _ = fi7
            _ = v8
        }
        p0 = (p0 % (((p0 % (((p0 >> ((0) & 63))) | 1))) | 1))
    } else {
        var v9 int32 = 0
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(p0, p0)
        var v11 uint8 = 2
        _ = fi10
        _ = v11
    }
    var v12 int16 = -32768
    v13 := []int{}
    _ = p0
    _ = v12
    _ = v13
    return ((("a" + "hello") + "a") + "a")
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    v3 := make(map[string]uint32)
    var v4 uint32 = 1
    var v14 uint64 = (v2 | (v2 + 1))
    if (42 == -1) {
        var v15 uint64 = (v2 % ((v14) | 1))
        var v16 bool = ((float64(0.0) - (1e10 - (1.5 * (float64(3.14159) / 0.1)))) >= ((float64(1.5) + (0.0 / 0.1)) - (0.0 / (float64(1e10) - 3.14159))))
        var v17 int8 = 77
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 1, F1: 32767}
    if (-1 >= 1) {
        for fi19 := 0; fi19 < 19; fi19++ {
            if ((((0.1 - ((-2.25 - 0.0) + 0.1)) < 0.1) && true) && false) {
                v20 := GenS2{F0: (v4 != v4), F1: 0, F2: v14}
                _ = v20
            }
            var v21 float64 = 0.0
            _ = fi19
            _ = v21
        }
        v22 := GenS3{F0: 634, F1: 127, F2: -411}
        switch "\n\t" {
        case "a":
            v23 := GenS4{F0: "a", F1: (v4 >> ((0) & 63)), F2: 42}
            _ = v23
        }
        _ = v22
    }
    if true {
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 float32 = -2.25
            _ = fi24
            _ = v25
        }
    }
    v26 := GenS4{F0: ("\n\t" + "hello"), F1: v4, F2: 65535}
    if false {
        v27 := []int{128}
        if false {
            if (true && true) {
                v4 += v4
            }
            v4 += (v4 ^ 4294967295)
            v27 = append(v27, -1)
        }
        _ = v27
    } else {
        for fi28 := 0; fi28 < 12; fi28++ {
            _ = len(v1)
            v29 := []int{}
            fmt.Println(fi28, v2, v2)
            _ = fi28
            _ = v29
        }
    }
    v30 := GenS1{F0: 42, F1: -2147483648}
    v4 -= 2
    if ("foo bar" < "") {
        v4 = (v4 << ((4294967295) & 63))
        var v31 uint = 42
        _ = v31
    } else {
        var v32 int = 1
        v33 := []int{v32, v32, v32}
        var v34 float64 = ((0.0 / -2.25) * 0.1)
        _ = v32
        _ = v33
        _ = v34
    }
    if ((998 >= 255) || (false && (!(43 != 42)))) {
        v2 = v14
        v35 := GenS5{F0: uint16(v14), F1: 0, F2: 258}
        for fi36 := 0; fi36 < 5; fi36++ {
            for fi37 := 0; fi37 < 1; fi37++ {
                v38 := GenS2{F0: (false && false), F1: (^(fi37)), F2: 1}
                v14 = (v2 / (((^(v14))) | 1))
                _ = fi37
                _ = v38
            }
            v14 = (v14 | 83)
            _ = fi36
        }
        _ = v35
    } else {
        if (-2147483648 == 2147483647) {
            var v39 float32 = 1e10
            _ = v39
        }
    }
    switch 342 {
    case 0:
        v40 := []int{682}
        _ = v40
    default:
        fmt.Println(v2, v4)
    }
    switch -9223372036854775808 {
    case 2147483647:
        v41 := GenS4{F0: "", F1: (v4 / ((v4) | 1)), F2: 816}
        _ = v41
    case 127:
        v14 = (v14 % (((v14 << ((369) & 63))) | 1))
    default:
        var v42 uint64 = v14
        _ = v42
    }
    if true {
        v4 = (v4 | v4)
    }
    var v43 float32 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v43
}
