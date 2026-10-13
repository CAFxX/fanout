package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]string)
    var v3 int8 = 67
    var v4 int16 = -14
    var v5 float32 = (3.14159 / (-2.25 * 1e10))
    v6 := make(map[string]int)
    var v7 float32 = 0.1
    v4 = (v4 & (v4 ^ (v4 - v4)))
    var v8 int = 127
    if true {
        if (128 >= 1) {
            v7 -= v5
            v7 -= (3.14159 + (v5 / (v5 * 0.1)))
        } else {
            v7 -= v7
            if ((^(v8)) != v8) {
                var v9 uint16 = 719
                _ = v9
            } else {
                fmt.Println(v8)
                if (-9223372036854775808 == 739) {
                    v5 = v7
                    if false {
                        v10 := GenS1{F0: 1}
                        _ = v10
                    } else {
                        var v11 uint8 = 127
                        v12 := []int{}
                        var v13 int32 = 821
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    v8 = 9223372036854775807
                } else {
                    var v14 bool = (((!true) && ("\n\t" < "foo bar")) && (2147483647 <= 4294967295))
                    _ = v14
                }
            }
        }
        if true {
            var v15 int32 = 42
            v16 := GenS1{F0: (v15 ^ (-(v15)))}
            v7 += 1.5
            _ = v15
            _ = v16
        } else {
            var v17 int = (^(v8))
            _ = v17
        }
    } else {
        switch ("hello" + ("\n\t" + (("foo bar" + (("" + ((("foo bar" + "foo bar") + "hello") + "hello")) + "hello")) + ("a" + ("\n\t" + ("foo bar" + ("" + (("\n\t" + "hello") + "foo bar")))))))) {
        case "":
            v18 := GenS1{F0: 477}
            _ = v18
        case "hello":
            v4 += v4
        case "a":
            fmt.Println(v8, v4)
        }
    }
    v19 := GenS1{F0: 1}
    var v20 float32 = 0.0
    if true {
        v21 := GenS1{F0: 276}
        v5 += 1.5
        _ = v21
    } else {
        var v22 int32 = 542
        v23 := GenS1{F0: 170}
        _ = v22
        _ = v23
    }
    if (((("\n\t" + "\n\t") + "") + "\n\t") == (("\n\t" + "foo bar") + "\n\t")) {
        v24 := []int{42, v8, 32767}
        _ = v24
    } else {
        switch v8 {
        case -1:
            if (127 == 1) {
                if (!true) {
                    v1[(("foo bar" + "a") + ("a" + ("hello" + "")))] = 579
                    fmt.Println(v7)
                } else {
                    var v25 uint64 = 4294967295
                    v20 = (v20 * (((1.5 / (1e10 * v7)) / 1.5) / float32(3.14159)))
                    v20 -= v7
                    _ = v25
                }
                var v26 int = 128
                v20 = (0.1 / (v5 + v5))
                _ = v26
            }
        case -853:
            v27 := GenS2{F0: 65535, F1: false}
            _ = v27
        }
        v28 := []int{32768, v8, -73, 32768}
        _ = v28
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 int16 = v4
            var v32 int32 = 0
            v33 := GenS1{F0: -8}
            v34 := GenS1{F0: -374}
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
            _ = v34
        }
        var v35 string = "foo bar"
        var v36 int64 = -661
        _ = fi29
        _ = v35
        _ = v36
    }
    v37 := []int{9223372036854775807, 42, v8}
    if len(v37) >= 2 {
        v37 = v37[1:len(v37)]
    }
    if (((true && (v7 != v7)) || false) && false) {
        var v38 int64 = 394
        _ = v38
    }
    fmt.Println(v4)
    v39 := GenS1{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
