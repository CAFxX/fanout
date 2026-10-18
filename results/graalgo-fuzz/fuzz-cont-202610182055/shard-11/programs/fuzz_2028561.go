package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int; F2 int }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8 }
func genFn0() uint32 {
    var v4 int16 = 1
    v5 := GenS1{F0: 2}
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS2{F0: 42, F1: fi6, F2: (fi6 / ((-2147483648) | 1))}
        if true {
            var v8 string = "a"
            fmt.Println(v8, v8)
            _ = v8
        } else {
            for fi9 := 0; fi9 < 8; fi9++ {
                switch fi6 {
                case 9223372036854775807:
                    v4 += -57
                case 42:
                    switch "a" {
                    case "hello":
                        v4 = (v4 * v4)
                    case "":
                        var v10 uint = 65535
                        _ = v10
                    }
                case -1:
                    v11 := GenS3{F0: fi6, F1: 234}
                    _ = v11
                }
                v12 := []int{-191}
                _ = fi9
                _ = v12
            }
            v13 := GenS3{F0: fi6, F1: 2147483647}
            _ = v13
        }
        var v14 uint32 = 3
        _ = fi6
        _ = v14
        _ = v7
    }
    _ = v4
    _ = v5
    return 2147483648
}
func genFn1() int32 {
    if (((!(0 >= 0)) || ((!false) || (338 >= 42))) || (1 <= 42)) {
        var v15 int32 = 42
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int8 = 106
        var v18 uint64 = 4294967295
        if (!(false || (421 >= 76))) {
            v19 := GenS4{F0: 4294967295, F1: (false && false)}
            v18 = v18
            v18 -= (v18 ^ (v18 >> ((384) & 63)))
            _ = v19
        } else {
            v20 := []int{4294967295}
            v20 = append(v20, fi16)
            _ = v20
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    return 1
}

func main() {
    var v1 int64 = 0
    v2 := make(map[string]uint32)
    var v3 string = ("\n\t" + ("" + "\n\t"))
    v3 = (v3 + v3)
    if true {
        v21 := []int{42, 2147483647}
        _ = v21
    } else {
        if true {
            fmt.Println(v3, v3)
            var v22 float64 = 0.0
            _ = v22
        }
        if true {
            var v23 uint16 = 1
            fmt.Println(v23, v1, v3)
            v24 := GenS3{F0: 128, F1: 787}
            _ = v23
            _ = v24
        } else {
            switch ("hello" + v3) {
            case "a":
                v3 = ("" + "hello")
            case "foo bar":
                v1 = (^(v1))
            }
        }
    }
    v25 := GenS5{F0: genFn0()}
    var v26 int16 = 1
    v1 -= v1
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := GenS6{F0: 2}
        v29 := []int{fi27}
        _ = fi27
        _ = v28
        _ = v29
    }
    v1 = v1
    var v30 uint8 = 127
    fmt.Println(v1, v26)
    fmt.Println(v30, v26)
    v31 := GenS6{F0: 0}
    v32 := []int{450, -9223372036854775808}
    _ = len(v2)
    if true {
        fmt.Println(v1)
    } else {
        var v33 uint32 = 42
        var v34 string = v3
        v35 := GenS1{F0: uint64(v30)}
        _ = v33
        _ = v34
        _ = v35
    }
    var v36 int8 = 118
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
}
