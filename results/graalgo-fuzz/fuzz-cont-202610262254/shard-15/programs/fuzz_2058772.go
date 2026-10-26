package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64; F1 bool }
func genFn0(p0 uint, p1 uint, p2 int64) int64 {
    v7 := []int{236}
    switch (("\n\t" + (("hello" + ("\n\t" + "hello")) + "\n\t")) + "hello") {
    case "a":
        p1 = (^(p0))
    }
    v8 := GenS1{F0: 74, F1: "a", F2: 35}
    var v9 int = -251
    v10 := GenS1{F0: 174, F1: "\n\t", F2: 197}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return p2
}
func genFn1(p0 uint) uint64 {
    if ("" != "foo bar") {
        var v11 int = 9223372036854775807
        _ = v11
    }
    p0 = (p0 / (((p0 + 199)) | 1))
    fmt.Println(p0)
    var v12 int64 = -9223372036854775808
    _ = p0
    _ = v12
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -128
    var v3 float64 = ((-2.25 - float64(0.1)) / 1e10)
    var v4 uint32 = 2
    var v5 float64 = 3.14159
    var v6 string = "foo bar"
    var v13 int32 = 565
    for fi14 := 0; fi14 < 15; fi14++ {
        v4 = (v4 ^ v4)
        _ = fi14
    }
    var v15 int8 = 127
    if (false && false) {
        var v16 int32 = (-(v13))
        switch -9223372036854775808 {
        case 42:
            v16 = 1
        case 222:
            v17 := v1[v4]
            _ = v17
        case 9223372036854775807:
            var v18 uint = 255
            _ = v18
        }
        _ = v16
    }
    v13 = (-(v13))
    v19 := GenS2{F0: 1, F1: 0}
    v20 := []int{}
    var v21 string = (v6 + v6)
    fmt.Println(v21)
    for fi22 := 0; fi22 < 10; fi22++ {
        if ((("foo bar" != ((v6 + ("a" + ("foo bar" + v21))) + ("hello" + "foo bar"))) || false) || (!(v2 < 63))) {
            if false {
                for fi23 := 0; fi23 < 17; fi23++ {
                    v1[(v4 - 2147483648)] = (v6 + (("a" + ("" + "foo bar")) + v21))
                    v24 := v1[v4]
                    fmt.Println(v4)
                    _ = fi23
                    _ = v24
                }
                v2 = 0
                var v25 int = 4294967295
                _ = v25
            } else {
                v13 += (v13 >> ((42) & 63))
            }
            var v26 uint8 = 15
            _ = v26
        } else {
            v27 := GenS3{F0: 18446744073709551615, F1: false}
            var v28 string = ("hello" + ((("a" + v21) + v6) + v21))
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = v27
            _ = v28
        }
        var v29 uint64 = 18446744073709551615
        if (true || true) {
            var v30 uint8 = 123
            _ = v30
        } else {
            if true {
                v5 -= ((0.0 / v3) * v5)
                var v31 int16 = 720
                var v32 uint = 442
                _ = v31
                _ = v32
            } else {
                v2 -= v2
                v21 = "\n\t"
                v21 = v21
            }
        }
        _ = fi22
        _ = v29
    }
    var v33 int32 = (v13 << ((2) & 63))
    v33 = (v13 % (((v13 | (v33 | (v33 / (((-(v13))) | 1))))) | 1))
    var v34 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
