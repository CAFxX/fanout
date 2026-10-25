package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 string; F2 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint) uint64 {
    v7 := []int{128, 895, 888}
    var v8 int = -731
    _ = p0
    _ = v7
    _ = v8
    return 403
}
func genFn1(p0 uint) uint {
    v9 := GenS1{F0: 255}
    p0 = p0
    var v10 int32 = 42
    v11 := GenS2{F0: p0, F1: "", F2: -9223372036854775808}
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return (p0 + (^(p0)))
}

func main() {
    var v1 int32 = 339
    v2 := []string{"a", "\n\t", ("foo bar" + ""), ((("hello" + "hello") + "foo bar") + "hello")}
    var v3 uint = 572
    v4 := []uint32{582, 525, 1}
    var v5 float64 = 0.1
    var v6 uint16 = 255
    var v12 int8 = -1
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS1{F0: 127}
        v15 := GenS1{F0: 2}
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS2{F0: (v3 + genFn1((v3 | (^(v3))))), F1: "\n\t", F2: 654}
    switch "foo bar" {
    case "\n\t":
        v1 = (v1 | 42)
    default:
        switch 42 {
        case -9223372036854775808:
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 int8 = (^(v12))
                for fi19 := 0; fi19 < 12; fi19++ {
                    switch (^(fi17)) {
                    case 42:
                        var v20 float32 = 3.14159
                        v21 := GenS3{F0: 2}
                        var v22 uint64 = 4294967295
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    v18 = (v18 * (v18 | 1))
                    v4 = append(v4, 818)
                    _ = fi19
                }
                fmt.Println(v12, v18)
                _ = fi17
                _ = v18
            }
        case 1:
            var v23 uint32 = 1
            _ = v23
        case 0:
            fmt.Println(v12)
        }
    }
    v12 -= (v12 / (((v12 | v12)) | 1))
    var v24 string = "\n\t"
    v3 += genFn1((v3 & (v3 / ((v3) | 1))))
    if (65 > 42) {
        if len(v2) > 0 {
            v2[0] = ""
        }
    } else {
        v25 := GenS1{F0: 42}
        var v26 uint64 = 722
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        if ((v6 ^ 42) < (v6 / ((v6) | 1))) {
            switch (^(fi27)) {
            case 0:
                for fi28 := 0; fi28 < 5; fi28++ {
                    if (2147483648 <= 2) {
                        fmt.Println(v12, v6, fi28)
                        v6 = (^(v6))
                    } else {
                        var v29 uint = (v3 + v3)
                        _ = v29
                    }
                    v2 = append(v2, "hello")
                    _ = fi28
                }
            case 128:
                fmt.Println(v5, v6, v5)
            }
            switch fi27 {
            case -345:
                var v30 int32 = (v1 * v1)
                _ = v30
            case -225:
                if (false || (2147483647 != -532)) {
                    v24 = "a"
                    var v31 uint32 = 1
                    _ = v31
                }
            }
            switch "a" {
            case "foo bar":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "":
                v24 = v24
            case "a":
                if false {
                    fmt.Println(fi27)
                } else {
                    v2 = append(v2, v24)
                }
            default:
                var v32 uint = 65535
                _ = v32
            }
        } else {
            for fi33 := 0; fi33 < 4; fi33++ {
                var v34 uint8 = 172
                if (!(32767 != -1)) {
                    v24 = v24
                }
                fmt.Println(v6, v1, v1)
                v5 = float64((((((3.14159 / 1e10) + 3.14159) - float32(-2.25)) + 1.5) / ((1e10 * float32((v5 + 0.0))) / 1.5)))
                v2 = append(v2, v24)
                _ = fi33
                _ = v34
            }
            for fi35 := 0; fi35 < 14; fi35++ {
                var v36 int32 = 182
                v24 = ""
                var v37 int8 = (v12 + v12)
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        var v38 int8 = 0
        v2 = append(v2, "\n\t")
        _ = fi27
        _ = v38
    }
    var v39 uint32 = 1
    v40 := GenS3{F0: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
