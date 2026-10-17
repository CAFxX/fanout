package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS2 struct { F0 string }
func genFn0(p0 uint, p1 uint8, p2 uint) int32 {
    fmt.Println(p1)
    v7 := []int{4294967295, -1, 282}
    var v8 string = "a"
    v9 := GenS1{F0: (1 <= 42), F1: 2147483648, F2: 2}
    switch ((v8 + v8) + "hello") {
    case "\n\t":
        if (((42 >= p1) && (!(-2147483648 <= -169))) || false) {
            if (true || false) {
                var v10 uint64 = 42
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                v11 := GenS2{F0: (v8 + "foo bar")}
                _ = v10
                _ = v11
            } else {
                p1 -= p1
                v12 := GenS2{F0: v8}
                fmt.Println(p1)
                _ = v12
            }
            p0 = p2
            var v13 float32 = -2.25
            _ = v13
        }
    case "a":
        var v14 uint64 = 552
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    var v4 uint16 = 65535
    v5 := make(map[uint32]bool)
    v6 := []uint32{}
    for fi15 := 0; fi15 < 9; fi15++ {
        if (!false) {
            var v16 string = (("hello" + ("foo bar" + "")) + "\n\t")
            var v17 int16 = 32767
            if (false && (true || (v17 != v17))) {
                switch (fi15 / (((fi15 ^ fi15)) | 1)) {
                case 4294967295:
                    if ((!(1.5 > (1.5 / 1.5))) || true) {
                        v18 := GenS1{F0: true, F1: 2147483648, F2: 364}
                        _ = v18
                    }
                case 114:
                    var v19 float64 = 1.5
                    _ = v19
                default:
                    var v20 uint16 = (v4 * v4)
                    _ = v20
                }
                switch (fi15 + (fi15 << ((255) & 63))) {
                case 540:
                    switch v16 {
                    case "hello":
                        fmt.Println(v16, v16, v17)
                    default:
                        var v21 uint8 = 233
                        _ = v21
                    }
                case -2147483648:
                    if (42 < -14) {
                        var v22 int = 127
                        fmt.Println(v1)
                        _ = v22
                    } else {
                        _ = len(v6)
                    }
                case -191:
                    fmt.Println(v1)
                default:
                    v16 = v16
                }
                v23 := GenS2{F0: (("\n\t" + (("" + v16) + "hello")) + (v16 + v16))}
                _ = v23
            }
            _ = v16
            _ = v17
        } else {
            for fi24 := 0; fi24 < 16; fi24++ {
                v25 := GenS2{F0: "foo bar"}
                v4 += 42
                _ = fi24
                _ = v25
            }
            v26 := GenS2{F0: ""}
            _ = v26
        }
        var v27 float64 = 1e10
        v28 := GenS1{F0: false, F1: 0, F2: 2}
        _ = fi15
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        var v30 float32 = (3.14159 / ((float32(0.0) / (1.5 - 0.1)) + (((3.14159 - ((float32(0.0) * 3.14159) / 1.0)) - float32((-2.25 / (0.1 - (((0.0 + 1e10) + (0.0 / 0.1)) / (0.0 + (-2.25 - ((3.14159 / (0.1 + 3.14159)) / -2.25)))))))) + ((0.0 + (3.14159 * -2.25)) / 1.5))))
        _ = fi29
        _ = v30
    }
    if (!true) {
        var v31 bool = (-10 == 85)
        v31 = (v4 >= v4)
        _ = v31
    }
    fmt.Println(v4)
    var v32 float64 = 1e10
    v1 = (v1 * (-(v1)))
    v33 := v5[42]
    var v34 uint64 = 797
    var v35 int = -747
    fmt.Println(v34, v32)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
