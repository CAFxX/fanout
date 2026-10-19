package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint8; F1 string; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int }
func genFn0() int {
    var v4 int = 2147483647
    switch "hello" {
    case "a":
        v4 += -35
    default:
        fmt.Println(v4)
    }
    var v5 uint = 506
    switch (("" + "") + "\n\t") {
    case "\n\t":
        fmt.Println(v5, v4)
    case "":
        var v6 int64 = 0
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 12; fi7++ {
            fmt.Println(v4, fi7, fi7)
            var v8 int32 = 0
            _ = fi7
            _ = v8
        }
    default:
        var v9 bool = (-2147483648 > 1)
        _ = v9
    }
    _ = v4
    _ = v5
    return v4
}
func genFn1(p0 int, p1 uint32) uint {
    v10 := []int{(p0 - (p0 ^ (p0 / (((p0 / (((^(p0))) | 1))) | 1)))), (^(p0)), p0, p0}
    v10 = append(v10, p0)
    v11 := GenS1{F0: -1, F1: 2147483647}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 517
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"foo bar", (("hello" + ("foo bar" + ("foo bar" + ""))) + ("foo bar" + "hello"))}
    var v3 float32 = (1.5 * -2.25)
    fmt.Println(v3)
    v12 := GenS2{F0: 1, F1: "\n\t", F2: 123}
    if false {
        if ((float64(v3) >= 3.14159) && ((42 > 9223372036854775807) || true)) {
            fmt.Println(v3)
            v3 = (v3 * (0.0 * (v3 * float32(0.1))))
            if ((1 >= 614) && (true && (genFn1(127, 328) != 386))) {
                if (42 > 129) {
                    var v13 uint8 = 255
                    _ = len(v1)
                    _ = v13
                }
                var v14 uint16 = 42
                var v15 uint32 = 220
                _ = v14
                _ = v15
            } else {
                v16 := GenS2{F0: 129, F1: "hello", F2: 0}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v3 += 0.1
                _ = v16
            }
        }
        delete(v1, "hello")
    }
    v17 := GenS3{F0: 4294967295, F1: 65535}
    var v18 int8 = -1
    switch 0 {
    case 2147483647:
        switch "foo bar" {
        case "":
            v19 := GenS1{F0: -1, F1: 32768}
            _ = v19
        default:
            v20 := GenS3{F0: 945, F1: 190}
            _ = v20
        }
    case 1:
        v3 = 0.0
    default:
        v21 := GenS3{F0: 1, F1: 2}
        _ = v21
    }
    if false {
        v3 = v3
    } else {
        v18 = -128
        v3 = -2.25
        v18 = 100
    }
    v22 := v1[(("foo bar" + "a") + "foo bar")]
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 string = v22
        switch ("foo bar" + "hello") {
        case "\n\t":
            for fi25 := 0; fi25 < 17; fi25++ {
                var v26 bool = (true || (1 > 255))
                fmt.Println(v24, fi25, fi23)
                v27 := GenS2{F0: 219, F1: "foo bar", F2: 103}
                _ = fi25
                _ = v26
                _ = v27
            }
        case "a":
            v3 = 0.1
        default:
            if len(v2) > 0 {
                v2[0] = ""
            }
        }
        _ = len(v2)
        _ = fi23
        _ = v24
    }
    v3 = (float32(float64(3.14159)) - (v3 * float32(float64(float32(0.0)))))
    v28 := GenS1{F0: 442, F1: -9223372036854775808}
    v18 = 0
    v29 := GenS4{F0: "hello"}
    var v30 int32 = 1
    v31 := GenS2{F0: 0, F1: v22, F2: 1}
    v32 := GenS5{F0: genFn0()}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
