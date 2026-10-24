package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 string; F1 int16 }
func genFn0(p0 int64, p1 uint32) uint8 {
    var v4 uint32 = (p1 / (((p1 << ((870) & 63))) | 1))
    v5 := GenS1{F0: 1000}
    v6 := GenS2{F0: 33, F1: 183}
    v7 := GenS1{F0: 42}
    var v8 string = ((("hello" + "foo bar") + ((("foo bar" + "\n\t") + ("" + "")) + "a")) + ("" + "a"))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 227
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v9 uint = 255
    v9 = 18446744073709551615
    v9 = 18446744073709551615
    v10 := GenS1{F0: 1}
    v9 = v9
    switch 389 {
    case 350:
        var v11 uint16 = 2
        _ = v11
    }
    if false {
        if (2 <= 1) {
            switch ("" + "hello") {
            case "hello":
                v12 := []int{613, 2147483647}
                _ = v12
            case "a":
                for fi13 := 0; fi13 < 20; fi13++ {
                    v9 -= (v9 * v9)
                    _ = fi13
                }
            case "\n\t":
                var v14 int32 = 42
                _ = v14
            }
            var v15 uint8 = genFn0(42, 1)
            for fi16 := 0; fi16 < 1; fi16++ {
                switch (fi16 + fi16) {
                case -442:
                    for fi17 := 0; fi17 < 5; fi17++ {
                        var v18 float32 = (float32(1.5) + 3.14159)
                        if (100 <= 127) {
                            v19 := v2[uint32(v9)]
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                }
                v9 = v9
                fmt.Println(v15)
                _ = fi16
            }
            _ = v15
        } else {
            v20 := GenS2{F0: 32767, F1: 127}
            if true {
                var v21 int64 = 1
                _ = v21
            } else {
                v9 = v9
                v9 = 40
            }
            _ = v20
        }
        v22 := GenS2{F0: 289, F1: genFn0(0, 888)}
        _ = v22
    }
    v9 += 0
    var v23 uint = (v9 & 4294967295)
    for fi24 := 0; fi24 < 11; fi24++ {
        fmt.Println(v9, fi24, v23)
        for fi25 := 0; fi25 < 7; fi25++ {
            v26 := []int{4294967295, fi24}
            v9 -= v9
            if true {
                var v27 uint32 = 0
                v23 = v9
                var v28 string = "\n\t"
                _ = v27
                _ = v28
            } else {
                var v29 int16 = 0
                v9 = 2
                delete(v2, 2)
                if false {
                    var v30 uint16 = 2
                    _ = v30
                }
                switch (^(fi24)) {
                case 1:
                    var v31 float64 = 3.14159
                    switch fi24 {
                    case 9223372036854775807:
                        fmt.Println(v9, fi24)
                    case -691:
                        v32 := GenS3{F0: 1}
                        _ = v32
                    case -9223372036854775808:
                        v9 = (v9 / ((v23) | 1))
                    }
                    var v33 int16 = (v29 ^ (v29 + v29))
                    v23 -= (v23 / (((v23 & (v23 + 548))) | 1))
                    _ = v31
                    _ = v33
                case -2147483648:
                    v34 := GenS1{F0: v29}
                    _ = v34
                case -72:
                    v9 -= v23
                }
                if (128 <= 20) {
                    v23 += (v23 ^ (v23 ^ v23))
                } else {
                    switch "\n\t" {
                    case "hello":
                        var v35 uint8 = 128
                        v35 = genFn0(0, 256)
                        _ = v35
                    case "a":
                        fmt.Println(v9, v29)
                    }
                    v26 = append(v26, (fi25 ^ (fi25 * (fi24 - fi25))))
                    fmt.Println(fi25, fi25, fi25)
                }
                _ = v29
            }
            if true {
                switch fi25 {
                case -9223372036854775808:
                    v36 := GenS4{F0: 42, F1: 127}
                    _ = v36
                case 374:
                    fmt.Println(fi25)
                case 9223372036854775807:
                    v23 = 18446744073709551615
                default:
                    if (!false) {
                        v9 += (^(v23))
                        v9 = (v23 ^ (v9 | v9))
                    } else {
                        var v37 int8 = 100
                        delete(v1, ("foo bar" + ""))
                        v38 := GenS3{F0: 1}
                        v39 := GenS5{F0: fi25, F1: 497}
                        v23 -= 42
                        v37 += (-(v37))
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                }
            } else {
                fmt.Println(v9)
            }
            v9 = (v9 * v23)
            v9 -= (v9 >> (((v9 / (((v23 ^ 89)) | 1))) & 63))
            if false {
                fmt.Println(fi24)
            }
            _ = fi25
            _ = v26
        }
        fmt.Println(fi24, v9)
        _ = fi24
    }
    switch "\n\t" {
    case "hello":
        if false {
            var v40 uint8 = 0
            _ = v40
        }
    case "\n\t":
        var v41 float32 = (3.14159 + (1.5 * (1e10 - (0.0 + (((1.5 - (3.14159 - 3.14159)) + (1.5 - (float32(1e10) * (3.14159 / -2.25)))) - (1.5 * (1.5 / 1.5)))))))
        _ = v41
    case "a":
        if ((true || (((-182 <= -32768) || (false || (1e10 >= 3.14159))) || true)) || (42 > 4294967295)) {
            v9 += v9
        }
    default:
        v42 := GenS6{F0: ("a" + "foo bar"), F1: 32767}
        _ = v42
    }
    v43 := []int{}
    v23 = 255
    v9 -= v23
    var v44 int16 = 42
    v9 = (v23 | v9)
    if (("foo bar" < "\n\t") && false) {
        var v45 int64 = -565
        v9 = v9
        var v46 uint32 = 563
        _ = v45
        _ = v46
    }
    v47 := GenS4{F0: 1, F1: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v43
    _ = v44
    _ = v47
    _ = v9
}
