package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int64 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int64; F1 uint8; F2 int8 }
func genFn0(p0 int32, p1 int32, p2 string) int16 {
    var v7 int16 = -1
    switch 42 {
    case 128:
        for fi8 := 0; fi8 < 3; fi8++ {
            p2 = ""
            for fi9 := 0; fi9 < 9; fi9++ {
                v7 -= v7
                p1 += p0
                _ = fi9
            }
            _ = fi8
        }
    case 0:
        p2 = ("a" + p2)
    default:
        p2 = "foo bar"
    }
    if true {
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 uint16 = 234
            var v12 int = -9223372036854775808
            if ((127 >= -1) || true) {
                p2 = p2
                switch p2 {
                case "hello":
                    v7 -= v7
                case "foo bar":
                    var v13 uint64 = 42
                    _ = v13
                }
            } else {
                fmt.Println(p2)
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return v7
}

func main() {
    v1 := []string{"", "\n\t", "foo bar"}
    var v2 uint64 = 965
    var v3 bool = true
    var v4 string = "hello"
    v5 := []int{891, -1, -2147483648, -9223372036854775808}
    var v6 uint8 = 255
    var v14 bool = (v3 || false)
    v15 := GenS1{F0: -1, F1: 860, F2: 42}
    v5 = append(v5, 1)
    v16 := GenS1{F0: 2147483647, F1: 221, F2: -1}
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS1{F0: -1, F1: 387, F2: -1}
            switch (fi18 - fi18) {
            case 212:
                v2 += 2
            }
            var v20 uint32 = uint32(v2)
            _ = fi18
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 uint = 1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v23 float32 = 0.0
            if (!true) {
                switch fi17 {
                case -1:
                    if v3 {
                        v22 -= v22
                    }
                case 256:
                    var v24 int32 = 42
                    _ = v24
                }
                v4 = v4
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi17
    }
    var v25 float32 = 3.14159
    if (v25 > v25) {
        fmt.Println(v25, v14)
        v4 = v4
        if (v14 || (v14 || v14)) {
            if (true && v3) {
                var v26 float64 = 1.5
                _ = v26
            } else {
                if (3.14159 == (1.5 / 3.14159)) {
                    var v27 float64 = (float64(v25) / 0.1)
                    v3 = ((v2 % ((v2) | 1)) < 18446744073709551615)
                    _ = v27
                } else {
                    if (!v14) {
                        v6 = (^(v6))
                        fmt.Println(v3, v4, v2)
                    }
                    fmt.Println(v4, v4)
                    v2 = (v2 | v2)
                }
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v28 int32 = 241
            _ = v28
        }
    }
    if (true || v3) {
        v25 = 3.14159
        if false {
            var v29 int8 = 101
            var v30 int32 = 42
            _ = v29
            _ = v30
        }
    }
    v6 = v6
    var v31 float32 = ((float32(0.0) / (v25 / (-2.25 / v25))) / v25)
    v32 := GenS2{F0: 100, F1: 1}
    v33 := GenS2{F0: 127, F1: 252}
    v3 = (2147483647 >= 2147483647)
    v34 := GenS3{F0: 9223372036854775807, F1: 1, F2: 42}
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
