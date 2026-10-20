package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
func genFn0(p0 string, p1 int64) int8 {
    v5 := GenS1{F0: "\n\t"}
    var v6 uint32 = 847
    var v7 int = 1
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := []int{fi8}
        v9 = append(v9, 42)
        if true {
            _ = len(v9)
        } else {
            var v10 int32 = -1
            if true {
                v11 := GenS1{F0: p0}
                _ = v11
            } else {
                if false {
                    fmt.Println(v6, v7, fi8)
                    if (!(((472 == (v10 << ((18446744073709551615) & 63))) || (v10 == v10)) || (-947 == v10))) {
                        var v12 int = v7
                        _ = v12
                    }
                    fmt.Println(v7, fi8, v7)
                }
                p0 = p0
                switch p0 {
                case "foo bar":
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                case "a":
                    _ = len(v9)
                }
            }
            if (true && (!false)) {
                fmt.Println(v6, v10)
            } else {
                v7 -= fi8
                fmt.Println(fi8)
            }
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 64
}
func genFn1(p0 int32) int {
    fmt.Println(p0)
    var v13 uint32 = 1
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 int = (fi14 * 1)
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: v13}
        v13 = (v13 & v13)
        _ = fi14
        _ = v17
    }
    var v18 uint32 = v13
    _ = p0
    _ = v13
    _ = v18
    return 128
}

func main() {
    var v1 float32 = 0.0
    var v2 int16 = -32768
    v3 := make(map[string]bool)
    var v4 int = -2147483648
    v19 := GenS1{F0: "foo bar"}
    v1 += (0.0 - (0.1 * v1))
    var v20 uint = 0
    v1 += 3.14159
    var v21 uint = v20
    v1 += float32(float64(float32(-2.25)))
    switch (v4 - v4) {
    case -312:
        if (true || false) {
            v1 = (float32(-2.25) - (3.14159 - v1))
            fmt.Println(v21, v20, v20)
            if true {
                v22 := GenS1{F0: "\n\t"}
                v2 = (-(v2))
                _ = v22
            }
        } else {
            v1 = (1.5 + (float32(0.1) / v1))
        }
    case 526:
        switch "hello" {
        case "hello":
            var v23 uint8 = 128
            _ = v23
        default:
            for fi24 := 0; fi24 < 13; fi24++ {
                v25 := []int{genFn1(761), fi24, (v4 * genFn1(241)), genFn1(42)}
                v21 += (v20 >> (((^(v20))) & 63))
                if true {
                    v21 = v21
                } else {
                    var v26 int = (v4 / ((v4) | 1))
                    var v27 string = (((("hello" + "a") + "") + ("" + "hello")) + ("foo bar" + (((("a" + (("a" + ("a" + "a")) + "")) + "\n\t") + "hello") + ("hello" + "a"))))
                    var v28 int8 = genFn0(v27, 0)
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi24
                _ = v25
            }
        }
    }
    var v29 int8 = int8(v2)
    v30 := v3[(("hello" + ((("hello" + ("a" + ("\n\t" + "foo bar"))) + "") + "hello")) + "a")]
    v31 := []int{v4, genFn1(-417)}
    v1 += v1
    _ = len(v31)
    var v32 int8 = genFn0(((("\n\t" + "hello") + (("\n\t" + "a") + "hello")) + ""), -9223372036854775808)
    for fi33 := 0; fi33 < 6; fi33++ {
        var v34 int = v4
        _ = fi33
        _ = v34
    }
    var v35 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
}
