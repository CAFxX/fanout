package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 string }
func genFn0(p0 uint16, p1 string, p2 int8) uint8 {
    var v7 uint = 742
    v8 := []int{}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 184
}
func genFn1() uint {
    switch -9223372036854775808 {
    case -715:
        v9 := []int{783, -97}
        _ = v9
    case -2147483648:
        switch "foo bar" {
        case "\n\t":
            var v10 int16 = -977
            _ = v10
        case "":
            if (false || (true && (0.0 <= float64((((1.5 * (1.5 * 3.14159)) - (float32(((0.0 - -2.25) + float64((0.0 + (0.1 + ((((0.1 + 0.1) * 0.0) + 3.14159) * -2.25)))))) + float32(1.5))) * 1.5))))) {
                var v11 int = -9223372036854775808
                _ = v11
            } else {
                switch (("" + "hello") + (("hello" + (("\n\t" + "hello") + "\n\t")) + "foo bar")) {
                case "hello":
                    var v12 bool = (!(!(0 >= -1)))
                    _ = v12
                }
                v13 := []int{1, -656, -544}
                _ = v13
            }
        case "hello":
            var v14 bool = (!(false && (3.14159 != (-2.25 / (float32(1e10) * -2.25)))))
            _ = v14
        default:
            if (!true) {
                var v15 float64 = 0.1
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 int = fi16
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
        }
    case 221:
        v18 := []int{9223372036854775807}
        _ = v18
    }
    var v19 int8 = 0
    v19 -= (v19 % (((^(v19))) | 1))
    v19 += v19
    _ = v19
    return 721
}

func main() {
    v1 := []uint32{2}
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    v5 := make(map[int]int)
    var v6 int32 = 2147483647
    fmt.Println(v6)
    v20 := GenS1{F0: 2, F1: 127, F2: 2147483647}
    delete(v3, 4294967295)
    v3[272] = -2147483648
    if (!("foo bar" >= "hello")) {
        v21 := GenS2{F0: true, F1: 42}
        var v22 int16 = -596
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: 705, F1: -472, F2: 2147483647}
    v6 -= (v6 & v6)
    switch "hello" {
    case "\n\t":
        v6 -= v6
    }
    var v24 int16 = int16(v6)
    fmt.Println(v6, v6)
    switch -2147483648 {
    case 29:
        fmt.Println(v6)
    }
    var v25 int64 = -1
    var v26 uint = genFn1()
    v27 := GenS3{F0: -1}
    if (42 != 50) {
        switch "a" {
        case "hello":
            var v28 int32 = (v6 >> ((4294967295) & 63))
            _ = v28
        }
    } else {
        var v29 int64 = (v25 >> ((v26) & 63))
        v30 := GenS4{F0: v26, F1: "a"}
        _ = v29
        _ = v30
    }
    var v31 uint64 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
