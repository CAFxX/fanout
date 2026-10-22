package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint64; F1 int8; F2 string }
type GenS4 struct { F0 uint; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    var v3 uint = 42
    v4 := make(map[uint32]int)
    var v5 uint16 = uint16(v3)
    v6 := []int{-9223372036854775808}
    if (false && false) {
        v3 += v3
        fmt.Println(v3, v3)
        for fi7 := 0; fi7 < 8; fi7++ {
            v5 = v5
            var v8 uint16 = (^(v5))
            _ = fi7
            _ = v8
        }
    } else {
        if false {
            var v9 uint16 = v5
            if false {
                delete(v2, -2147483648)
            }
            _ = v9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            v5 = v5
            v11 := GenS1{F0: v3}
            _ = fi10
            _ = v11
        }
        fmt.Println(v5, v3)
    }
    v12 := v2[9223372036854775807]
    var v13 float64 = 3.14159
    if (0.1 >= 0.1) {
        switch -1 {
        case 32767:
            v14 := GenS1{F0: v3}
            _ = v14
        case 1:
            v15 := GenS2{F0: -331, F1: (!v12)}
            _ = v15
        default:
            if ((!(!v12)) || (-1 == -9223372036854775808)) {
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 string = "hello"
                    var v18 uint64 = 18446744073709551615
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                v13 = (v13 * (float64((3.14159 + 0.0)) + v13))
                var v19 float64 = float64((0.0 / (1e10 + (((0.1 + 3.14159) * (1.5 / (0.1 - (0.1 + float32((v13 * v13)))))) + float32((1e10 * v13))))))
                _ = v19
            }
        }
        v5 = v5
    }
    fmt.Println(v3)
    var v20 bool = (2147483648 == 183)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            v6 = append(v6, -2147483648)
            if ((!v12) && v20) {
                v23 := v4[636]
                v6 = append(v6, fi22)
                _ = v23
            } else {
                var v24 int = -995
                switch (^(fi21)) {
                case 0:
                    var v25 float64 = ((v13 * 1.5) - v13)
                    _ = v25
                default:
                    var v26 uint8 = 117
                    _ = v26
                }
                switch (fi21 >> (((v3 >> ((v3) & 63))) & 63)) {
                case -690:
                    switch (((("a" + "a") + "foo bar") + (("a" + "a") + (("foo bar" + "hello") + "foo bar"))) + "foo bar") {
                    case "foo bar":
                        v24 = (^(fi22))
                    case "":
                        var v27 string = "hello"
                        _ = v27
                    }
                default:
                    fmt.Println(v5, fi22, v20)
                }
                _ = v24
            }
            _ = fi22
        }
        var v28 int8 = 1
        var v29 int64 = 1
        _ = fi21
        _ = v28
        _ = v29
    }
    if true {
        for fi30 := 0; fi30 < 17; fi30++ {
            for fi31 := 0; fi31 < 5; fi31++ {
                if (391 >= (v3 * (v3 + (v3 + (^(v3)))))) {
                    if (!v12) {
                        v32 := GenS3{F0: 1, F1: 127, F2: ""}
                        _ = v32
                    } else {
                        v3 -= (v3 + 0)
                    }
                    v33 := GenS3{F0: 42, F1: 127, F2: "\n\t"}
                    _ = v33
                }
                _ = fi31
            }
            v3 = v3
            switch "\n\t" {
            case "hello":
                v12 = (true || v20)
            case "foo bar":
                v34 := GenS1{F0: v3}
                _ = v34
            default:
                fmt.Println(v5, v3, fi30)
            }
            _ = fi30
        }
        delete(v1, 0)
        var v35 float32 = -2.25
        _ = v35
    } else {
        switch (("\n\t" + ("foo bar" + "foo bar")) + "") {
        case "hello":
            var v36 uint64 = 101
            _ = v36
        case "":
            v37 := GenS2{F0: 0, F1: v20}
            _ = v37
        case "foo bar":
            var v38 bool = v12
            _ = v38
        default:
            var v39 bool = ((3.14159 / (0.1 - 3.14159)) >= (float32(v13) + 3.14159))
            _ = v39
        }
        var v40 int8 = -14
        v41 := GenS4{F0: 994, F1: 140}
        _ = v40
        _ = v41
    }
    v20 = v20
    var v42 int = 128
    v13 += v13
    var v43 int64 = 0
    v43 = (^(v43))
    v44 := GenS4{F0: (^(v3)), F1: 15}
    v6 = append(v6, -9223372036854775808)
    switch v42 {
    case -2147483648:
        var v45 int8 = -40
        _ = v45
    default:
        var v46 bool = v20
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
