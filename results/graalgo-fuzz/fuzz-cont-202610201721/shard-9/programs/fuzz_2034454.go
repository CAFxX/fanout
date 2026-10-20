package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int }

func main() {
    var v1 uint = 0
    v2 := make(map[string]string)
    var v3 uint = (v1 * v1)
    v4 := make(map[string]string)
    var v5 int = 32767
    var v6 uint16 = 999
    v6 -= v6
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 255, F1: 1}
        _ = fi7
        _ = v8
    }
    var v9 int16 = -1
    fmt.Println(v9, v3)
    switch ("hello" + "a") {
    case "foo bar":
        if (1 == 104) {
            if (128 != (v5 >> ((42) & 63))) {
                v10 := []int{}
                _ = v10
            }
        } else {
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 int = (v5 - -940)
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 14; fi13++ {
                if true {
                    var v14 bool = true
                    fmt.Println(v5)
                    _ = v14
                } else {
                    v9 = (-(v9))
                    switch -9223372036854775808 {
                    case 127:
                        delete(v2, (("a" + ("hello" + "")) + "foo bar"))
                    default:
                        var v15 uint = 0
                        _ = v15
                    }
                    if (0 != fi13) {
                        for fi16 := 0; fi16 < 9; fi16++ {
                            v2["foo bar"] = "\n\t"
                            v17 := GenS1{F0: 236, F1: 18446744073709551615}
                            v5 = (^(v5))
                            v18 := []int{fi13, (fi16 + int(v9)), (fi13 - (fi16 | 9223372036854775807))}
                            v18 = append(v18, v5)
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        v19 := GenS2{F0: -128}
                        var v20 int64 = -1
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi13
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                switch ("a" + "") {
                case "":
                    fmt.Println(v3, v6, v6)
                }
                v22 := []int{(-(v5))}
                _ = fi21
                _ = v22
            }
        }
    case "hello":
        _ = len(v4)
    default:
        fmt.Println(v1, v5)
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 float32 = (float32(1e10) * 3.14159)
        switch ("a" + "hello") {
        case "a":
            var v25 bool = false
            _ = v25
        case "":
            var v26 float64 = 1e10
            _ = v26
        case "foo bar":
            v24 = 3.14159
        default:
            v27 := GenS3{F0: 846, F1: v5}
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    if true {
        var v28 int16 = v9
        _ = v28
    } else {
        if (((0.0 / 1.0) / 1.0) < (1e10 - 1.5)) {
            var v29 uint64 = 4294967295
            _ = v29
        }
        if false {
            fmt.Println(v6, v3)
        }
    }
    if (917 < v3) {
        v9 = v9
        v30 := GenS3{F0: 1, F1: -9223372036854775808}
        _ = v30
    } else {
        var v31 int = v5
        fmt.Println(v6, v1, v3)
        _ = v31
    }
    var v32 string = ""
    for fi33 := 0; fi33 < 18; fi33++ {
        v9 = v9
        var v34 int16 = 42
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        var v36 string = ((v32 + (("foo bar" + v32) + ("hello" + (("\n\t" + (("\n\t" + "foo bar") + "a")) + v32)))) + "")
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
