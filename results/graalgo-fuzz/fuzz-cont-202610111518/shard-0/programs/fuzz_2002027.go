package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
func genFn0(p0 int32, p1 int) uint64 {
    switch 128 {
    case 4294967295:
        var v5 uint8 = 186
        _ = v5
    case 159:
        p1 += 0
    default:
        p0 += p0
    }
    v6 := GenS1{F0: -325, F1: "\n\t", F2: 353}
    for fi7 := 0; fi7 < 1; fi7++ {
        p0 = (p0 | (p0 % ((-278) | 1)))
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return 4294967295
}
func genFn1(p0 int64, p1 int64, p2 uint32) string {
    switch ("a" + "hello") {
    case "hello":
        v8 := GenS1{F0: 42, F1: "", F2: 42}
        _ = v8
    default:
        for fi9 := 0; fi9 < 10; fi9++ {
            p1 = (p1 | 9223372036854775807)
            _ = fi9
        }
    }
    var v10 uint = 536
    p0 = (p1 | p0)
    p1 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return "\n\t"
}

func main() {
    var v1 int16 = -1
    var v2 int = 32768
    v3 := make(map[int]string)
    var v4 float32 = ((1.5 - -2.25) - -2.25)
    v11 := GenS1{F0: 1, F1: "foo bar", F2: 423}
    var v12 float32 = v4
    var v13 float64 = (((((((0.1 / float64(float32((0.1 / 1.5)))) / 1e10) + 0.0) + 1.5) + -2.25) * 3.14159) + 1.5)
    v14 := []int{32767, v2}
    var v15 float64 = float64(1e10)
    var v16 float64 = v13
    v17 := GenS1{F0: -497, F1: "hello", F2: 1}
    v18 := GenS2{F0: "hello"}
    var v19 uint8 = 128
    v20 := GenS1{F0: 9223372036854775807, F1: "hello", F2: 32767}
    v1 = v1
    if true {
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := v3[fi21]
            v13 -= (v16 - (v16 + (float64(((float32((v13 / (0.1 - (v16 * (v16 - v15))))) * (((float32(v13) - 1.5) - (v4 / v12)) - (v12 + 1.5))) / float32(v16))) * (-2.25 / v13))))
            v23 := GenS3{F0: 738}
            _ = fi21
            _ = v22
            _ = v23
        }
        if (!(!false)) {
            v15 -= (0.0 + v15)
        } else {
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 int32 = 2147483647
                _ = fi24
                _ = v25
            }
            if ((v2 / (((-(v2))) | 1)) >= -2147483648) {
                if (v2 <= (v2 - (v2 & 42))) {
                    v4 = (0.0 + (v4 + (3.14159 / v4)))
                }
                v26 := v3[v2]
                _ = v26
            }
        }
        var v27 uint32 = 263
        _ = v27
    } else {
        for fi28 := 0; fi28 < 17; fi28++ {
            var v29 string = ("" + (("\n\t" + "\n\t") + "a"))
            var v30 int32 = 0
            _ = fi28
            _ = v29
            _ = v30
        }
        v12 -= 1.5
        v3[(v2 | -2147483648)] = (("foo bar" + "a") + (("" + ("" + ("hello" + ("\n\t" + "\n\t")))) + ""))
    }
    v31 := GenS1{F0: 9223372036854775807, F1: ("foo bar" + "foo bar"), F2: 1}
    for fi32 := 0; fi32 < 7; fi32++ {
        v12 -= (0.0 / (v4 * (0.0 / ((float32(float64(v4)) / 1.5) - v12))))
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
}
