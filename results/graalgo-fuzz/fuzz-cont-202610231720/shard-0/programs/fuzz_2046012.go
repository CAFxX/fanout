package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 int8 }

func main() {
    var v1 bool = (32767 != -9223372036854775808)
    var v2 float32 = float32((3.14159 + -2.25))
    v3 := make(map[string]bool)
    var v4 int16 = 1000
    v5 := make(map[int]string)
    switch -1 {
    case 128:
        for fi6 := 0; fi6 < 2; fi6++ {
            if (!(((v4 * 0) <= 1) && (!false))) {
                var v7 uint32 = 4294967295
                var v8 float64 = float64((float32(1e10) - 0.0))
                _ = v7
                _ = v8
            }
            switch fi6 {
            case -822:
                var v9 string = "a"
                _ = v9
            case -2147483648:
                v10 := []int{}
                _ = v10
            case -2:
                if v1 {
                    if v1 {
                        if true {
                            switch "hello" {
                            case "hello":
                                v11 := GenS1{F0: 971, F1: 1}
                                _ = v11
                            case "":
                                var v12 string = "hello"
                                _ = v12
                            }
                            if (64 != 42) {
                                if false {
                                    for fi13 := 0; fi13 < 13; fi13++ {
                                        v14 := GenS1{F0: 487, F1: 2147483647}
                                        v15 := GenS2{F0: -1}
                                        v2 = float32((0.1 - ((3.14159 - float64(1.5)) - -2.25)))
                                        _ = fi13
                                        _ = v14
                                        _ = v15
                                    }
                                    v16 := GenS3{F0: 2}
                                    var v17 int = (fi6 - fi6)
                                    _ = v16
                                    _ = v17
                                } else {
                                    fmt.Println(v4)
                                }
                                var v18 int = (fi6 - fi6)
                                switch (fi6 / ((v18) | 1)) {
                                case 14:
                                    v4 = (v4 % (((v4 >> ((4294967295) & 63))) | 1))
                                case 127:
                                    var v19 int64 = 1
                                    _ = v19
                                case 128:
                                    v18 = 128
                                }
                                _ = v18
                            } else {
                                v20 := GenS2{F0: -2147483648}
                                v2 = v2
                                v2 = v2
                                _ = v20
                            }
                            v21 := GenS4{F0: false}
                            _ = v21
                        } else {
                            v1 = true
                            v22 := GenS3{F0: 129}
                            _ = v22
                        }
                    }
                    var v23 uint32 = 736
                    _ = v23
                } else {
                    switch ("" + ("a" + ("hello" + "hello"))) {
                    case "hello":
                        if true {
                            if (true && (42 > 2147483647)) {
                                var v24 uint64 = 338
                                v25 := GenS5{F0: -2147483648, F1: 255}
                                _ = v24
                                _ = v25
                            }
                            switch ("" + "foo bar") {
                            case "a":
                                fmt.Println(fi6)
                            case "hello":
                                if (2 <= 255) {
                                    var v26 uint8 = 0
                                    v27 := v5[fi6]
                                    _ = v26
                                    _ = v27
                                } else {
                                    v2 += float32(0.1)
                                    if false {
                                        v28 := GenS5{F0: fi6, F1: 127}
                                        delete(v5, fi6)
                                        var v29 uint32 = 2147483647
                                        _ = v28
                                        _ = v29
                                    } else {
                                        var v30 string = ("foo bar" + "hello")
                                        _ = v30
                                    }
                                }
                            }
                        } else {
                            var v31 string = ("hello" + "a")
                            if false {
                                v31 = v31
                            }
                            _ = v31
                        }
                    case "foo bar":
                        switch (fi6 >> ((818) & 63)) {
                        case 4294967295:
                            switch "\n\t" {
                            case "a":
                                for fi32 := 0; fi32 < 17; fi32++ {
                                    var v33 uint = 2
                                    v34 := GenS6{F0: 45}
                                    _ = fi32
                                    _ = v33
                                    _ = v34
                                }
                            case "hello":
                                v2 = (1.5 * v2)
                            }
                        case 0:
                            switch "\n\t" {
                            case "a":
                                var v35 int8 = 15
                                _ = v35
                            case "foo bar":
                                v2 += v2
                            case "":
                                var v36 int8 = 72
                                _ = v36
                            default:
                                var v37 float32 = v2
                                _ = v37
                            }
                        }
                    case "a":
                        v1 = false
                    default:
                        for fi38 := 0; fi38 < 16; fi38++ {
                            v1 = (3.14159 == (1.5 - float64(v2)))
                            v39 := []int{}
                            _ = fi38
                            _ = v39
                        }
                    }
                }
            }
            _ = fi6
        }
    case 9223372036854775807:
        for fi40 := 0; fi40 < 14; fi40++ {
            v2 = 1e10
            for fi41 := 0; fi41 < 6; fi41++ {
                var v42 uint16 = 742
                var v43 int32 = 42
                var v44 float32 = (float32((0.0 - 0.0)) - 1e10)
                _ = fi41
                _ = v42
                _ = v43
                _ = v44
            }
            v2 = (v2 * 1.5)
            _ = fi40
        }
    case 32768:
        if v1 {
            if (v1 && (42 != -2147483648)) {
                v45 := GenS2{F0: 42}
                _ = v45
            }
        } else {
            fmt.Println(v4)
            for fi46 := 0; fi46 < 1; fi46++ {
                var v47 float32 = (v2 + v2)
                _ = fi46
                _ = v47
            }
            var v48 bool = v1
            _ = v48
        }
    }
    v2 -= (v2 - v2)
    var v49 bool = true
    var v50 bool = (v49 && (!(-2147483648 < 0)))
    v2 += 0.0
    v51 := GenS3{F0: 1}
    fmt.Println(v49, v49, v1)
    var v52 string = ""
    v3[(v52 + v52)] = v49
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
