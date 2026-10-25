package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }
func genFn0() uint64 {
    v7 := GenS1{F0: 572, F1: 42}
    v8 := []int{127}
    var v9 uint16 = 42
    v9 = (v9 / (((v9 | 1)) | 1))
    _ = len(v8)
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() int32 {
    v10 := []int{-375, -9223372036854775808}
    for fi11 := 0; fi11 < 5; fi11++ {
        v10 = append(v10, (fi11 * (fi11 / ((-319) | 1))))
        switch ("hello" + ("\n\t" + "hello")) {
        case "a":
            v12 := GenS1{F0: 971, F1: 2}
            _ = v12
        case "hello":
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 int8 = 42
                var v15 uint = 0
                if true {
                    if false {
                        v10 = append(v10, (fi13 ^ fi13))
                    } else {
                        v16 := GenS2{F0: 1}
                        _ = v16
                    }
                    switch fi11 {
                    case 385:
                        switch fi11 {
                        case 717:
                            v15 = (v15 - v15)
                        }
                    }
                    var v17 int8 = (v14 ^ -83)
                    _ = v17
                } else {
                    v15 -= (v15 + v15)
                    v18 := GenS1{F0: 42, F1: 0}
                    fmt.Println(fi13, fi13, v14)
                    _ = v18
                }
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            var v19 uint64 = 2
            _ = v19
        }
        _ = fi11
    }
    var v20 uint = 0
    var v21 int8 = 100
    var v22 float32 = 1.5
    _ = v10
    _ = v20
    _ = v21
    _ = v22
    return 1
}

func main() {
    var v1 int64 = 0
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    v4 := []string{("hello" + "hello"), "a", ("a" + ("a" + ("\n\t" + ("" + ("a" + "\n\t"))))), "hello"}
    var v5 uint = 4294967295
    v6 := []int{}
    for fi23 := 0; fi23 < 6; fi23++ {
        fmt.Println(v5)
        var v24 string = "a"
        for fi25 := 0; fi25 < 4; fi25++ {
            v24 = ("foo bar" + v24)
            if (-1 <= (fi25 * fi25)) {
                v26 := GenS3{F0: 42}
                _ = v26
            }
            switch fi23 {
            case 127:
                var v27 uint8 = 1
                _ = v27
            case 875:
                v5 -= v5
            case -2147483648:
                v28 := GenS1{F0: 4294967295, F1: 297}
                _ = v28
            }
            switch 343 {
            case 4294967295:
                v29 := GenS3{F0: 42}
                _ = v29
            case 127:
                var v30 string = v24
                _ = v30
            case -1:
                var v31 int8 = 46
                _ = v31
            }
            if false {
                var v32 float64 = (0.0 * 3.14159)
                v5 -= 4294967295
                _ = v32
            } else {
                v33 := GenS4{F0: -1}
                _ = v33
            }
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    v5 = 693
    var v34 int64 = (v1 * v1)
    if true {
        v5 = 255
    }
    v3["hello"] = false
    v1 += v1
    v1 = 42
    var v35 string = ("\n\t" + "a")
    var v36 string = v35
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
