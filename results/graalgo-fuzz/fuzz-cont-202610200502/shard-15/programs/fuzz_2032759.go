package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 string }

func main() {
    var v1 float64 = ((1e10 - 1.5) - 0.1)
    v2 := make(map[uint32]uint32)
    var v3 string = "a"
    var v4 int32 = 1
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(v3)
        var v6 int16 = -13
        var v7 int8 = 0
        _ = fi5
        _ = v6
        _ = v7
    }
    if ((v3 + "a") >= (("hello" + ("" + "hello")) + "")) {
        if false {
            var v8 uint32 = 2
            var v9 string = ("\n\t" + v3)
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: (((v3 + v3) + ("\n\t" + (v3 + (("foo bar" + ((((v3 + v3) + "") + "\n\t") + ("\n\t" + v3))) + "foo bar")))) + v3), F1: ("" + (v3 + "foo bar")), F2: v4}
        _ = v10
    } else {
        v11 := GenS1{F0: v3, F1: v3, F2: (v4 + 928)}
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := v2[495]
        v1 -= v1
        _ = fi12
        _ = v13
    }
    if (32767 < 127) {
        if (1 <= 831) {
            if true {
                fmt.Println(v1)
                if (true || false) {
                    switch 2147483647 {
                    case -310:
                        v14 := v2[42]
                        _ = v14
                    case 1:
                        switch v3 {
                        case "\n\t":
                            v15 := GenS1{F0: (v3 + v3), F1: v3, F2: (v4 & (v4 / (((^(v4))) | 1)))}
                            _ = v15
                        case "foo bar":
                            for fi16 := 0; fi16 < 17; fi16++ {
                                var v17 int16 = 851
                                _ = fi16
                                _ = v17
                            }
                        case "":
                            var v18 int32 = 1
                            _ = v18
                        }
                    }
                    v3 = v3
                    v3 = ((v3 + "hello") + ("a" + (v3 + v3)))
                }
            } else {
                v19 := GenS1{F0: v3, F1: v3, F2: -795}
                fmt.Println(v3, v1, v4)
                _ = v19
            }
            v20 := GenS1{F0: "", F1: "\n\t", F2: -2147483648}
            var v21 int64 = 1
            _ = v20
            _ = v21
        }
        if false {
            switch -2147483648 {
            case 32767:
                v4 -= -1
            case 127:
                for fi22 := 0; fi22 < 4; fi22++ {
                    v3 = (v3 + "hello")
                    _ = fi22
                }
            default:
                v4 = 0
            }
        } else {
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := []int{(fi23 | (fi23 & fi23))}
                v2[1] = 1
                _ = fi23
                _ = v24
            }
            fmt.Println(v1, v4, v1)
        }
    } else {
        fmt.Println(v4)
        v25 := []int{9223372036854775807}
        var v26 bool = ((((1e10 - (1.5 * 1e10)) / (v1 + (v1 - 3.14159))) / ((v1 + 0.0) + v1)) != 0.1)
        _ = v25
        _ = v26
    }
    var v27 string = v3
    var v28 float32 = 3.14159
    v29 := []int{0, 324}
    for fi30 := 0; fi30 < 3; fi30++ {
        v31 := GenS2{F0: 126}
        v32 := GenS2{F0: 382}
        switch (v27 + v27) {
        case "hello":
            fmt.Println(v27, v4)
        case "":
            v33 := GenS3{F0: 128, F1: "foo bar"}
            _ = v33
        case "\n\t":
            var v34 float32 = float32((v1 * (0.1 * float64(3.14159))))
            _ = v34
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    var v35 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v4
}
