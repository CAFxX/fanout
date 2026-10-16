package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint) int32 {
    v7 := []int{42, 9223372036854775807}
    p0 = 1
    if true {
        if len(v7) > 0 {
            v7[0] = -187
        }
        var v8 uint64 = 4294967295
        fmt.Println(v8)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return -1
}
func genFn1() int64 {
    v9 := GenS1{F0: 98}
    if (false || true) {
        var v10 uint8 = 0
        if true {
            if false {
                if (true && true) {
                    v10 -= v10
                    if false {
                        switch "" {
                        case "":
                            var v11 uint16 = 65535
                            _ = v11
                        case "\n\t":
                            v12 := GenS1{F0: 42}
                            _ = v12
                        case "hello":
                            v13 := []int{-568, 4294967295}
                            _ = v13
                        }
                        switch (("\n\t" + "foo bar") + ("foo bar" + "foo bar")) {
                        case "":
                            v14 := GenS1{F0: 2}
                            _ = v14
                        }
                        var v15 float64 = (3.14159 + -2.25)
                        _ = v15
                    }
                    v16 := []int{2147483647, 4294967295, -1, 4294967295}
                    _ = v16
                }
                v17 := []int{}
                v10 = v10
                _ = v17
            } else {
                var v18 string = "foo bar"
                v18 = "a"
                _ = v18
            }
            for fi19 := 0; fi19 < 3; fi19++ {
                for fi20 := 0; fi20 < 18; fi20++ {
                    v21 := GenS1{F0: 65535}
                    var v22 int8 = 42
                    var v23 int64 = int64(v10)
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                fmt.Println(v10)
                var v24 bool = ((-232 >= -32768) && (100 < -1))
                _ = fi19
                _ = v24
            }
        }
        _ = v10
    }
    v25 := []int{2147483647, 32767, 2147483647}
    _ = v25
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := []uint32{}
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    v5 := []uint32{928, 2147483648, 708, 788}
    var v6 int16 = 1000
    v26 := v3[0]
    v4[""] = (v26 != v26)
    var v27 float32 = ((3.14159 - float32(0.0)) + -2.25)
    fmt.Println(v2, v6)
    v28 := GenS1{F0: 641}
    switch (v26 * (v26 * v26)) {
    case 0:
        var v29 uint64 = 307
        _ = v29
    case -9223372036854775808:
        var v30 float32 = (0.0 / v27)
        _ = v30
    }
    var v31 int64 = (v2 + v2)
    for fi32 := 0; fi32 < 4; fi32++ {
        v33 := GenS1{F0: 42}
        v34 := GenS2{F0: fi32, F1: -39}
        v35 := GenS3{F0: 2}
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        fmt.Println(v27, v6, v6)
        _ = fi36
    }
    v27 = v27
    switch "\n\t" {
    case "hello":
        switch (v26 + (^(v26))) {
        case 2147483647:
            for fi37 := 0; fi37 < 1; fi37++ {
                _ = len(v4)
                _ = fi37
            }
        }
    case "foo bar":
        var v38 int32 = -2147483648
        _ = v38
    case "":
        v26 = v26
    }
    fmt.Println(v26, v6)
    var v39 uint32 = 0
    if (((0.0 <= (3.14159 - ((3.14159 + ((3.14159 + 1e10) * 1.5)) * float64(-2.25)))) && (!(false && (810 > 255)))) && false) {
        if ((((v27 - ((v27 + (v27 * (((1.5 / v27) + 1e10) + v27))) / ((0.0 / -2.25) + -2.25))) - (((v27 * ((v27 - 3.14159) * (v27 - (v27 + 1.5)))) / 1.5) - (((1e10 * (float32(0.0) - (float32((-2.25 - -2.25)) + v27))) / float32(-2.25)) / ((v27 + (1.5 * (v27 / v27))) * v27)))) - (v27 / float32(float64(v27)))) <= 1e10) {
            v40 := GenS3{F0: 236}
            var v41 bool = false
            v42 := GenS1{F0: v39}
            _ = v40
            _ = v41
            _ = v42
        }
        v43 := GenS2{F0: v26, F1: 42}
        _ = v43
    }
    var v44 string = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
