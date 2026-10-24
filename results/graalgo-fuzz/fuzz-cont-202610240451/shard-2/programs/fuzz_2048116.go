package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 uint32 }
func genFn0(p0 string, p1 uint64) uint {
    v4 := GenS1{F0: (^(p1)), F1: 2147483647, F2: 0}
    p0 = ("foo bar" + "hello")
    v5 := GenS2{F0: 42, F1: 2, F2: 42}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 779
}

func main() {
    var v1 int8 = 32
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    var v6 uint64 = 0
    if (3.14159 != -2.25) {
        v6 -= 0
        switch 1 {
        case -457:
            v7 := GenS1{F0: 2, F1: -976, F2: 2147483648}
            _ = v7
        case 4294967295:
            if true {
                var v8 float64 = 0.1
                v9 := GenS1{F0: (v6 * v6), F1: 494, F2: 2147483647}
                switch 127 {
                case 127:
                    fmt.Println(v6)
                case -794:
                    fmt.Println(v8, v6, v8)
                case 42:
                    if false {
                        v10 := []int{779, -1, 128}
                        v11 := v2[4294967295]
                        v11 = v11
                        _ = v10
                        _ = v11
                    }
                }
                _ = v8
                _ = v9
            } else {
                v1 = v1
                v6 += 2
                switch "a" {
                case "foo bar":
                    fmt.Println(v6, v1)
                }
            }
        default:
            if (!false) {
                for fi12 := 0; fi12 < 4; fi12++ {
                    v1 -= (v1 | (-(v1)))
                    _ = fi12
                }
                for fi13 := 0; fi13 < 13; fi13++ {
                    v14 := GenS1{F0: (v6 % (((^(v6))) | 1)), F1: 1, F2: 0}
                    switch (^(fi13)) {
                    case -28:
                        var v15 bool = (-1 == fi13)
                        _ = v15
                    case 42:
                        v16 := GenS3{F0: 0}
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
                v6 += (v6 | (v6 ^ 679))
            }
        }
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            fmt.Println(fi17, v1)
            _ = fi17
        }
    }
    _ = len(v2)
    switch 9223372036854775807 {
    case 170:
        if true {
            v18 := []int{4294967295, -9223372036854775808, 32767}
            fmt.Println(v1, v6)
            v6 -= 42
            _ = v18
        } else {
            v19 := []int{-1}
            _ = v19
        }
    }
    if false {
        v20 := GenS2{F0: 0, F1: 65535, F2: 405}
        if (false || (!true)) {
            for fi21 := 0; fi21 < 11; fi21++ {
                v22 := []int{fi21, (-(fi21)), fi21}
                _ = fi21
                _ = v22
            }
        } else {
            v1 = -128
            v1 -= v1
            v23 := []int{}
            _ = v23
        }
        var v24 float32 = float32(1.5)
        _ = v20
        _ = v24
    }
    v1 = v1
    fmt.Println(v1, v1)
    v25 := GenS2{F0: -923, F1: 359, F2: 255}
    var v26 int64 = 1
    v27 := GenS1{F0: v6, F1: 804, F2: 2147483648}
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 bool = (127 >= 42)
        var v30 float64 = 0.1
        _ = fi28
        _ = v29
        _ = v30
    }
    v3["\n\t"] = 973
    if ((-32768 >= 1) || (("a" + ((("\n\t" + "hello") + "a") + ("foo bar" + ("" + "\n\t")))) >= "hello")) {
        var v31 string = "a"
        _ = v31
    }
    var v32 int = 4294967295
    if (!true) {
        v32 -= (^(v32))
        for fi33 := 0; fi33 < 2; fi33++ {
            fmt.Println(v26, fi33, v26)
            _ = fi33
        }
        if false {
            var v34 uint64 = 4294967295
            v35 := GenS3{F0: 65535}
            v32 = -706
            _ = v34
            _ = v35
        }
    } else {
        v6 += 2
    }
    var v36 float64 = 1e10
    fmt.Println(v26)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v36
    _ = v6
}
