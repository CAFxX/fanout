package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 uint32 }
func genFn0(p0 string, p1 uint64, p2 string) int8 {
    if false {
        fmt.Println(p1, p2, p2)
    }
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{-1, (fi4 + (fi4 * (fi4 / (((-(fi4))) | 1)))), (fi4 % (((-(fi4))) | 1)), (^(fi4))}
        var v6 uint = 0
        var v7 string = ("hello" + p0)
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 float64 = 1e10
    v9 := GenS1{F0: 9223372036854775807, F1: "foo bar"}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 100
    v3 := []uint32{699, 2}
    switch 535 {
    case 1:
        var v10 int64 = 0
        _ = v10
    case 32767:
        var v11 float64 = (float64(1.5) + 0.1)
        _ = v11
    default:
        if false {
            v1[("foo bar" + "hello")] = ("a" + "")
            if true {
                v2 -= (v2 % (((v2 | (v2 & genFn0("a", 0, "hello")))) | 1))
                v2 -= (^(v2))
            } else {
                for fi12 := 0; fi12 < 2; fi12++ {
                    var v13 float32 = 3.14159
                    _ = fi12
                    _ = v13
                }
            }
            var v14 uint8 = 1
            _ = v14
        }
    }
    v15 := v1["a"]
    if (42 < 673) {
        fmt.Println(v15, v2)
        fmt.Println(v15, v2)
    }
    v2 = v2
    for fi16 := 0; fi16 < 5; fi16++ {
        switch v15 {
        case "hello":
            v2 = (v2 >> ((1) & 63))
        case "":
            fmt.Println(v15)
        case "a":
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := GenS2{F0: -1, F1: 255, F2: 42}
                var v19 bool = (true && (((v15 + v15) + (("hello" + ("\n\t" + (v15 + "a"))) + v15)) < v15))
                _ = fi17
                _ = v18
                _ = v19
            }
        default:
            for fi20 := 0; fi20 < 5; fi20++ {
                v21 := GenS2{F0: 1, F1: 128, F2: 39}
                _ = fi20
                _ = v21
            }
        }
        v15 = "\n\t"
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := GenS2{F0: 42, F1: 1, F2: 0}
            v24 := GenS2{F0: 0, F1: 0, F2: 128}
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi16
    }
    if (42 != 206) {
        if false {
            v3 = append(v3, 1)
            _ = len(v1)
        } else {
            v25 := v1[v15]
            _ = v25
        }
        var v26 uint = 255
        fmt.Println(v26, v15, v2)
        _ = v26
    }
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    if false {
        for fi27 := 0; fi27 < 18; fi27++ {
            v2 = (v2 - (v2 ^ (^(v2))))
            if ((-78 != (v2 % (((v2 | (v2 * v2))) | 1))) || (false || true)) {
                if (671 == 4294967295) {
                    v2 = 51
                    v3 = append(v3, 2)
                    v28 := GenS3{F0: 42}
                    _ = v28
                }
                fmt.Println(fi27)
                var v29 int = fi27
                _ = v29
            } else {
                v30 := GenS3{F0: 2}
                _ = v30
            }
            var v31 int64 = 271
            _ = fi27
            _ = v31
        }
        v32 := GenS4{F0: 9223372036854775807, F1: 4294967295}
        var v33 string = v15
        _ = v32
        _ = v33
    }
    var v34 int32 = -1
    var v35 float32 = 0.1
    var v36 int8 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
