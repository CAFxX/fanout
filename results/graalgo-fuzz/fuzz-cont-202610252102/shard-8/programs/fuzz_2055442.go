package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint }

func main() {
    var v1 float64 = (0.0 + float64(1e10))
    var v2 string = "hello"
    var v3 uint16 = 210
    switch ("foo bar" + "hello") {
    case "a":
        v2 = ""
    case "foo bar":
        if (-9223372036854775808 == 127) {
            switch "" {
            case "hello":
                v4 := GenS1{F0: 42}
                _ = v4
            case "foo bar":
                var v5 int32 = 1
                _ = v5
            case "a":
                for fi6 := 0; fi6 < 4; fi6++ {
                    var v7 int = 127
                    var v8 uint = 0
                    for fi9 := 0; fi9 < 19; fi9++ {
                        v10 := []int{(v7 | fi6)}
                        _ = fi9
                        _ = v10
                    }
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            default:
                v11 := []int{0}
                _ = v11
            }
            v12 := GenS1{F0: 0}
            var v13 uint64 = 2
            _ = v12
            _ = v13
        } else {
            if true {
                if (1e10 >= (0.1 / float32(((v1 + (v1 / 3.14159)) + v1)))) {
                    if (255 <= 445) {
                        v14 := []int{0, 2147483647, 2147483647, 42}
                        if true {
                            if true {
                                fmt.Println(v3)
                                fmt.Println(v1, v1)
                                if len(v14) > 0 {
                                    v14[0] = 223
                                }
                            }
                            switch v2 {
                            case "\n\t":
                                for fi15 := 0; fi15 < 11; fi15++ {
                                    v3 += v3
                                    var v16 uint = 901
                                    switch v2 {
                                    case "a":
                                        var v17 int64 = 9223372036854775807
                                        _ = v17
                                    case "\n\t":
                                        if true {
                                            v2 = v2
                                            if true {
                                                var v18 uint = (v16 - v16)
                                                var v19 int8 = 42
                                                if len(v14) > 0 {
                                                    v14[0] = (^(fi15))
                                                }
                                                _ = v18
                                                _ = v19
                                            }
                                            var v20 uint32 = 1
                                            _ = v20
                                        }
                                    case "hello":
                                        var v21 int16 = 1000
                                        _ = v21
                                    }
                                    _ = fi15
                                    _ = v16
                                }
                            case "hello":
                                v1 = 1.5
                            }
                            v22 := GenS1{F0: -9223372036854775808}
                            _ = v22
                        } else {
                            var v23 bool = (0 > 42)
                            fmt.Println(v23, v23)
                            var v24 int8 = 1
                            _ = v23
                            _ = v24
                        }
                        var v25 float64 = 0.1
                        _ = v14
                        _ = v25
                    } else {
                        v26 := []int{-2147483648, 32768}
                        v2 = (((v2 + "hello") + "\n\t") + (v2 + "foo bar"))
                        _ = v26
                    }
                    if false {
                        switch "a" {
                        case "\n\t":
                            for fi27 := 0; fi27 < 13; fi27++ {
                                var v28 uint32 = 42
                                switch v2 {
                                case "":
                                    fmt.Println(v28, v3, v28)
                                case "foo bar":
                                    if false {
                                        v2 = (v2 + "\n\t")
                                        var v29 int8 = 15
                                        v2 = (((("foo bar" + (v2 + "")) + (v2 + v2)) + v2) + v2)
                                        _ = v29
                                    }
                                default:
                                    v2 = v2
                                }
                                var v30 string = v2
                                _ = fi27
                                _ = v28
                                _ = v30
                            }
                        case "a":
                            for fi31 := 0; fi31 < 1; fi31++ {
                                switch v2 {
                                case "foo bar":
                                    v32 := []int{-9223372036854775808, 9223372036854775807, (fi31 << ((178) & 63))}
                                    _ = v32
                                case "":
                                    v33 := []int{fi31, (-(fi31))}
                                    _ = v33
                                case "\n\t":
                                    for fi34 := 0; fi34 < 3; fi34++ {
                                        var v35 uint16 = v3
                                        v1 = 0.0
                                        fmt.Println(fi34, v3)
                                        _ = fi34
                                        _ = v35
                                    }
                                }
                                _ = fi31
                            }
                        case "":
                            v1 += v1
                        default:
                            var v36 bool = true
                            _ = v36
                        }
                        v1 -= float64(float32(float64(-2.25)))
                        v37 := GenS2{F0: ("\n\t" + (v2 + "")), F1: 42}
                        _ = v37
                    }
                }
                for fi38 := 0; fi38 < 20; fi38++ {
                    if (!true) {
                        switch "a" {
                        case "":
                            var v39 int = (^(fi38))
                            _ = v39
                        case "hello":
                            v40 := GenS2{F0: "", F1: 0}
                            _ = v40
                        }
                        var v41 bool = false
                        for fi42 := 0; fi42 < 3; fi42++ {
                            v2 = (((v2 + (v2 + v2)) + v2) + v2)
                            _ = fi42
                        }
                        _ = v41
                    }
                    _ = fi38
                }
            }
            v43 := GenS1{F0: 1}
            _ = v43
        }
    }
    if false {
        v44 := GenS2{F0: v2, F1: -732}
        _ = v44
    } else {
        v45 := []int{32767, -968, -185}
        _ = v45
    }
    v2 = v2
    v46 := GenS2{F0: v2, F1: 84}
    v2 = v2
    fmt.Println(v1, v3)
    fmt.Println(v1, v1)
    switch (v2 + v2) {
    case "a":
        v47 := GenS2{F0: v2, F1: 0}
        _ = v47
    case "":
        v1 += (v1 - float64(1.5))
    case "\n\t":
        if (!("hello" >= v2)) {
            var v48 int32 = -655
            v49 := []int{47, -1}
            _ = v48
            _ = v49
        }
    default:
        var v50 string = "foo bar"
        _ = v50
    }
    v51 := []int{813, -283, 128, 42}
    var v52 uint16 = (v3 - (v3 >> ((0) & 63)))
    v1 -= (v1 - ((3.14159 + v1) * (v1 / float64(1e10))))
    v53 := GenS3{F0: 42, F1: 283, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v51
    _ = v52
    _ = v53
}
