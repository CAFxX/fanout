package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int32; F1 uint }

func main() {
    var v1 string = (("" + ("" + (("foo bar" + ("" + "\n\t")) + "a"))) + (((("" + "a") + ("\n\t" + ("\n\t" + "hello"))) + "foo bar") + "hello"))
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 uint16 = 255
    v5 := GenS1{F0: 147, F1: 209}
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint = 705
        if false {
            v1 = (v1 + v1)
            if true {
                v1 = ((v1 + "a") + ("\n\t" + "\n\t"))
                v1 = v1
                if false {
                    v8 := v3[v1]
                    var v9 float32 = (1e10 + 3.14159)
                    v8 -= (-(v8))
                    _ = v8
                    _ = v9
                }
            }
            if false {
                v10 := GenS2{F0: -1}
                _ = v10
            } else {
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 uint16 = (v4 - 255)
                    var v13 bool = true
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v14 := GenS3{F0: 42, F1: 1}
                _ = v14
            }
        } else {
            v15 := []int{}
            _ = v15
        }
        v1 = ((("foo bar" + (v1 + "hello")) + ("hello" + "\n\t")) + "\n\t")
        _ = fi6
        _ = v7
    }
    v16 := GenS3{F0: 255, F1: 95}
    v4 -= (v4 - (^(v4)))
    if ((-2.25 == float32(0.1)) || (42 > v4)) {
        for fi17 := 0; fi17 < 2; fi17++ {
            var v18 float64 = 0.0
            fmt.Println(fi17)
            var v19 uint64 = 970
            _ = fi17
            _ = v18
            _ = v19
        }
        if true {
            v20 := GenS1{F0: 0, F1: 4294967295}
            fmt.Println(v4, v1)
            if true {
                var v21 uint = 24
                v22 := []int{9223372036854775807}
                _ = v21
                _ = v22
            } else {
                fmt.Println(v1)
                fmt.Println(v4)
                fmt.Println(v1, v4)
            }
            _ = v20
        } else {
            v23 := GenS4{F0: 0, F1: 1}
            var v24 int32 = 0
            _ = v23
            _ = v24
        }
        if ((0.0 * 0.1) < float64(1e10)) {
            fmt.Println(v1)
        }
    }
    var v25 uint = 466
    v4 += v4
    var v26 uint8 = 42
    delete(v2, 127)
    var v27 float32 = (3.14159 + 3.14159)
    v26 = 42
    if true {
        for fi28 := 0; fi28 < 14; fi28++ {
            v4 -= (^(v4))
            v25 = v25
            v29 := []int{fi28, 715}
            _ = fi28
            _ = v29
        }
        v30 := []int{1}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
