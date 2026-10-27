package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int64; F1 int32; F2 uint8 }
type GenS5 struct { F0 int8 }
func genFn0() int64 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint = 977
        for fi7 := 0; fi7 < 14; fi7++ {
            switch -1 {
            case 4294967295:
                var v8 float32 = ((0.0 + 0.0) + (0.0 + (1e10 * -2.25)))
                _ = v8
            case -1:
                switch fi5 {
                case 871:
                    v6 = 154
                case 4294967295:
                    var v9 uint8 = 235
                    _ = v9
                case -835:
                    v6 -= v6
                }
                if false {
                    v10 := GenS1{F0: 0, F1: 0, F2: "a"}
                    v11 := GenS1{F0: 2147483647, F1: -2147483648, F2: ("foo bar" + ("" + "hello"))}
                    _ = v10
                    _ = v11
                }
                var v12 int = (fi5 % ((fi7) | 1))
                _ = v12
            }
            var v13 float64 = 0.1
            _ = fi7
            _ = v13
        }
        _ = fi5
        _ = v6
    }
    var v14 uint32 = 0
    _ = v14
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 int32 = 42
    v4 -= (v4 * v4)
    switch 2147483647 {
    case 592:
        v4 -= v4
    case 42:
        var v15 int8 = -128
        _ = v15
    default:
        for fi16 := 0; fi16 < 16; fi16++ {
            if false {
                var v17 int = (fi16 + fi16)
                v1[65535] = ("a" + "\n\t")
                _ = v17
            } else {
                v18 := GenS2{F0: genFn0()}
                _ = v18
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                v20 := GenS1{F0: (v4 ^ v4), F1: 47, F2: ((("" + "") + "hello") + "")}
                var v21 uint16 = 0
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi16
        }
    }
    var v22 string = "foo bar"
    if ((true || (!(9223372036854775807 <= -184))) && true) {
        fmt.Println(v4, v4)
        v1[628] = v22
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 int64 = 155
            fmt.Println(v22, fi23, v4)
            var v25 int8 = 100
            _ = fi23
            _ = v24
            _ = v25
        }
    } else {
        var v26 int32 = v4
        v27 := GenS3{F0: 258, F1: 255}
        v22 = v22
        _ = v26
        _ = v27
    }
    fmt.Println(v22, v22)
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 int8 = -128
        v30 := []int{(fi28 ^ fi28), (fi28 / ((fi28) | 1)), fi28, 32768}
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 uint64 = 704
    for fi32 := 0; fi32 < 11; fi32++ {
        if false {
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 float64 = -2.25
                switch (v22 + (((((v22 + v22) + (v22 + (v22 + v22))) + "a") + v22) + "hello")) {
                case "hello":
                    fmt.Println(v31)
                case "foo bar":
                    var v35 uint = 1
                    _ = v35
                }
                v36 := []int{}
                _ = fi33
                _ = v34
                _ = v36
            }
            var v37 int = (fi32 * fi32)
            v38 := GenS2{F0: -40}
            _ = v37
            _ = v38
        } else {
            v39 := GenS4{F0: 1, F1: v4, F2: 127}
            if (128 == (fi32 - (fi32 | 4294967295))) {
                v4 = v4
                v22 = v22
            }
            v40 := GenS5{F0: -39}
            _ = v39
            _ = v40
        }
        v41 := GenS1{F0: v4, F1: v4, F2: "foo bar"}
        var v42 string = v22
        _ = fi32
        _ = v41
        _ = v42
    }
    var v43 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v43
}
