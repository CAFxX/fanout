package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 string }
func genFn0(p0 int) int8 {
    var v7 uint32 = 2147483648
    var v8 float32 = (((0.1 - ((((-2.25 * 1e10) + 0.1) / (-2.25 + 1.5)) - (((0.1 / -2.25) * (0.0 / 0.1)) * -2.25))) / 1.5) * float32(3.14159))
    v7 += v7
    v7 += v7
    v7 += (^(v7))
    _ = p0
    _ = v7
    _ = v8
    return -128
}
func genFn1() int16 {
    var v9 int = 4294967295
    var v10 uint8 = 148
    _ = v10
    _ = v9
    return -825
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{((("" + "foo bar") + "\n\t") + "\n\t"), "hello", "hello"}
    var v3 int64 = -490
    var v4 uint32 = 1
    var v5 float64 = (-2.25 - ((float64(float32((((3.14159 / -2.25) * float64((((3.14159 / 1.0) / 1.0) - (1e10 - (float32(0.1) / ((float32(0.0) - 1.5) * 0.1)))))) - 1.5))) + 1e10) - 1e10))
    v6 := []string{("foo bar" + ""), ("\n\t" + "a"), ("foo bar" + "a"), (("foo bar" + "") + ("a" + "hello"))}
    var v11 int32 = 0
    fmt.Println(v3)
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(v3, v11)
        _ = fi12
    }
    if (v11 < 625) {
        v13 := GenS1{F0: (^(v11)), F1: 1, F2: 1000}
        _ = v13
    }
    v14 := GenS2{F0: (^(v4))}
    v5 += v5
    fmt.Println(v3, v3, v3)
    var v15 int64 = v3
    if ("hello" < "\n\t") {
        if (!true) {
            switch -508 {
            case 1:
                var v16 int16 = 617
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            v11 = (^(v11))
            v5 = v5
            _ = fi17
        }
    }
    v18 := GenS3{F0: 972, F1: v4, F2: "foo bar"}
    if false {
        var v19 float64 = 1e10
        _ = v19
    } else {
        v5 += 0.0
        for fi20 := 0; fi20 < 7; fi20++ {
            for fi21 := 0; fi21 < 17; fi21++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v15 += 1
                v1[v4] = (fi20 ^ fi20)
                _ = fi21
            }
            _ = fi20
        }
    }
    v22 := v1[(v4 >> ((42) & 63))]
    var v23 int64 = (v3 + (v3 & v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
