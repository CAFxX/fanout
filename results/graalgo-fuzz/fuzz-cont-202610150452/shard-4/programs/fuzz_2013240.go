package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 uint8; F2 int64 }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]int)
    var v3 string = ("" + "\n\t")
    var v4 int32 = -1
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            if (-2147483648 <= 4294967295) {
                v4 = -1
                v7 := GenS1{F0: 0}
                _ = v7
            } else {
                v3 = v3
                switch fi5 {
                case 270:
                    v8 := GenS2{F0: (v3 + "hello")}
                    _ = v8
                }
                switch v3 {
                case "\n\t":
                    switch ((v3 + v3) + (v3 + (v3 + "hello"))) {
                    case "hello":
                        var v9 string = (v3 + v3)
                        _ = len(v2)
                        var v10 uint32 = 65535
                        _ = v10
                        _ = v9
                    case "":
                        if (float64(1.5) < -2.25) {
                            var v11 uint8 = 0
                            var v12 bool = ((^(fi6)) == fi6)
                            if false {
                                if v1 {
                                    if (3.14159 < 1e10) {
                                        v13 := v2[2147483647]
                                        _ = v13
                                    } else {
                                        v12 = v12
                                        v3 = v3
                                        if (((fi5 / (((fi5 ^ 32768)) | 1)) < fi6) && v12) {
                                            switch (v3 + v3) {
                                            case "foo bar":
                                                v14 := GenS3{F0: false}
                                                _ = v14
                                            }
                                            if v1 {
                                                v15 := GenS4{F0: 884}
                                                v16 := GenS1{F0: 42}
                                                if v12 {
                                                    v17 := []int{}
                                                    _ = v17
                                                }
                                                _ = v15
                                                _ = v16
                                            }
                                            v18 := []int{200, (fi5 & fi5), (fi6 * 404), fi6}
                                            _ = v18
                                        } else {
                                            v19 := GenS5{F0: 0, F1: 1, F2: 42}
                                            var v20 int = -218
                                            v2[4294967295] = 42
                                            _ = v19
                                            _ = v20
                                        }
                                    }
                                }
                                fmt.Println(v4, v1, v12)
                            }
                            _ = v11
                            _ = v12
                        } else {
                            v21 := v2[65535]
                            var v22 int64 = 0
                            v4 = v4
                            _ = v21
                            _ = v22
                        }
                        v23 := GenS5{F0: 70, F1: 129, F2: 9223372036854775807}
                        _ = v23
                    default:
                        var v24 uint = 670
                        _ = v24
                    }
                }
                v25 := GenS5{F0: -1, F1: 128, F2: -355}
                _ = v25
            }
            v4 = -1
            v26 := GenS2{F0: (("" + "a") + v3)}
            _ = fi6
            _ = v26
        }
        v1 = (!v1)
        _ = fi5
    }
    var v27 int16 = 0
    v1 = true
    if ("hello" >= v3) {
        fmt.Println(v27)
    }
    v4 -= (v4 << ((612) & 63))
    switch v3 {
    case "foo bar":
        v1 = v1
    case "":
        v1 = v1
    case "hello":
        v28 := []int{32768, -9223372036854775808, -1, 1}
        _ = v28
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 uint = 42
        v31 := GenS1{F0: -128}
        var v32 int16 = -32768
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    fmt.Println(v3)
    v33 := GenS4{F0: 255}
    fmt.Println(v27, v27, v3)
    var v34 bool = true
    for fi35 := 0; fi35 < 15; fi35++ {
        for fi36 := 0; fi36 < 16; fi36++ {
            fmt.Println(v27, v34, fi36)
            if true {
                var v37 int32 = (v4 + v4)
                switch fi36 {
                case 584:
                    switch 128 {
                    case 42:
                        var v38 uint = 642
                        _ = v38
                    case -1:
                        switch ("\n\t" + v3) {
                        case "a":
                            var v39 uint64 = 2
                            _ = v39
                        case "hello":
                            v27 += v27
                        case "\n\t":
                            v3 = "a"
                        }
                    }
                case -1:
                    var v40 uint8 = 2
                    _ = v40
                }
                _ = v37
            }
            var v41 int64 = -888
            _ = fi36
            _ = v41
        }
        _ = fi35
    }
    v42 := []int{9223372036854775807}
    fmt.Println(v34)
    v4 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v42
}
