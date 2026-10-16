package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint32; F1 uint; F2 string }
type GenS3 struct { F0 int64; F1 uint32; F2 int64 }
func genFn0(p0 uint16, p1 bool) int {
    v4 := []int{}
    var v5 int = int(p0)
    if ((-2.25 + (1e10 * -2.25)) <= float32(-2.25)) {
        for fi6 := 0; fi6 < 2; fi6++ {
            fmt.Println(p0, v5)
            var v7 uint8 = 127
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: 42, F1: (("hello" + "") + ((("foo bar" + "hello") + ("\n\t" + (("foo bar" + "hello") + "hello"))) + "foo bar"))}
        var v9 bool = (!p1)
        _ = v8
        _ = v9
    }
    _ = len(v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (-(v5))
}

func main() {
    var v1 uint32 = 2147483648
    var v2 bool = (true && false)
    v3 := []int{-334, -538, 127, 166}
    var v10 uint32 = v1
    v11 := GenS1{F0: (v10 | v10), F1: "a"}
    var v12 float32 = 0.1
    fmt.Println(v10, v12, v12)
    var v13 float32 = ((float32(-2.25) - (v12 + v12)) / (((0.0 - 0.1) / 1.0) + (float32(0.0) + (float32(3.14159) / (3.14159 + v12)))))
    _ = len(v3)
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 string = "foo bar"
        switch "foo bar" {
        case "a":
            v16 := GenS1{F0: (v10 % (((v1 + 2147483647)) | 1)), F1: (v15 + v15)}
            _ = v16
        case "\n\t":
            var v17 int = (^(fi14))
            _ = v17
        }
        var v18 uint64 = 0
        _ = fi14
        _ = v15
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v3 = append(v3, (fi19 * fi19))
        if (v2 || true) {
            var v20 bool = false
            v21 := GenS1{F0: v10, F1: "foo bar"}
            v3 = append(v3, (fi19 | (fi19 >> ((2) & 63))))
            _ = v20
            _ = v21
        } else {
            for fi22 := 0; fi22 < 7; fi22++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi22
            }
            if (42 >= 100) {
                var v23 uint32 = v1
                v23 = (v1 | (v1 >> ((uint(v23)) & 63)))
                _ = v23
            }
            if (v2 || v2) {
                var v24 uint = 52
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v26 float64 = float64(0.1)
            var v27 float64 = v26
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi19
    }
    v28 := GenS2{F0: (v10 % ((v10) | 1)), F1: 1, F2: "hello"}
    v2 = v2
    if true {
        if len(v3) > 0 {
            v3[0] = -9223372036854775808
        }
    } else {
        v12 = v13
        v10 -= (v1 >> ((65535) & 63))
    }
    v29 := GenS3{F0: -1, F1: (v10 * v1), F2: 611}
    v3 = append(v3, 334)
    var v30 uint16 = 42
    v12 -= v12
    fmt.Println(v2, v12)
    v1 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
