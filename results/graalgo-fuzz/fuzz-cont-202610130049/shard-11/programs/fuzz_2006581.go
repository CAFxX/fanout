package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 string; F1 int16 }
func genFn0(p0 int8) int64 {
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            switch fi6 {
            case 92:
                switch (("hello" + "foo bar") + "\n\t") {
                case "hello":
                    var v7 uint64 = 18446744073709551615
                    _ = v7
                case "\n\t":
                    if (!(!false)) {
                        fmt.Println(fi6, fi5)
                        var v8 string = ("" + ("a" + "foo bar"))
                        var v9 int = -961
                        _ = v8
                        _ = v9
                    }
                case "a":
                    switch (("" + (("a" + "foo bar") + "hello")) + ("a" + "a")) {
                    case "a":
                        var v10 uint8 = 0
                        var v11 int16 = 1
                        if ((v11 + v11) <= 1) {
                            v11 -= (v11 - v11)
                            v12 := []int{fi6, fi5, fi6}
                            _ = v12
                        }
                        _ = v10
                        _ = v11
                    case "foo bar":
                        p0 = -68
                    default:
                        fmt.Println(fi5, p0)
                    }
                default:
                    if true {
                        v13 := []int{fi6}
                        if false {
                            fmt.Println(p0, fi5)
                            p0 += p0
                            v14 := GenS1{F0: p0}
                            v13 = append(v13, (fi6 - (fi6 / ((fi5) | 1))))
                            _ = v14
                        }
                        _ = v13
                    } else {
                        v15 := []int{9223372036854775807, (fi6 / ((128) | 1)), -2147483648, fi5}
                        p0 += 0
                        v16 := GenS2{F0: 1, F1: fi5}
                        _ = v15
                        _ = v16
                    }
                }
            case 42:
                if (255 != 65535) {
                    fmt.Println(fi6)
                    switch "" {
                    case "\n\t":
                        p0 -= p0
                    case "foo bar":
                        var v17 float64 = 3.14159
                        _ = v17
                    }
                    var v18 int = (fi5 ^ 42)
                    _ = v18
                }
            case 32768:
                if (0 >= 18446744073709551615) {
                    if (65535 < 2) {
                        if ("" >= "\n\t") {
                            v19 := []int{fi5, (fi6 / ((fi5) | 1)), (^(fi6))}
                            p0 -= (p0 - p0)
                            _ = v19
                        } else {
                            var v20 int16 = 1000
                            _ = v20
                        }
                    }
                    var v21 float64 = 1e10
                    _ = v21
                } else {
                    v22 := GenS3{F0: (((("\n\t" + "foo bar") + ("a" + ("a" + "foo bar"))) + ("\n\t" + "")) + "a"), F1: 32767}
                    _ = v22
                }
            }
            fmt.Println(p0, fi6)
            v23 := []int{}
            _ = fi6
            _ = v23
        }
        var v24 uint64 = 0
        _ = fi5
        _ = v24
    }
    v25 := []int{-501}
    fmt.Println(p0)
    _ = p0
    _ = v25
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = float64((1.5 * (1e10 / 1.5)))
    var v3 uint8 = 1
    v4 := []string{("a" + ("" + ""))}
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 uint8 = (v3 >> ((1) & 63))
        switch fi26 {
        case 220:
            for fi28 := 0; fi28 < 5; fi28++ {
                v3 += v3
                var v29 bool = true
                _ = fi28
                _ = v29
            }
        case 42:
            for fi30 := 0; fi30 < 1; fi30++ {
                switch "hello" {
                case "foo bar":
                    var v31 uint = 42
                    _ = v31
                }
                _ = fi30
            }
        }
        _ = fi26
        _ = v27
    }
    var v32 int16 = 179
    if false {
        fmt.Println(v3, v32)
    } else {
        var v33 float64 = (1.5 + 0.0)
        fmt.Println(v33, v32)
        _ = v33
    }
    var v34 int64 = -640
    if true {
        if false {
            v34 += genFn0(100)
            v32 = -597
        }
        var v35 uint16 = 0
        _ = v35
    } else {
        var v36 uint32 = 517
        var v37 uint16 = 42
        fmt.Println(v2, v32)
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        var v39 string = "hello"
        _ = len(v1)
        v3 = (^(v3))
        _ = fi38
        _ = v39
    }
    switch 2147483647 {
    case 1:
        var v40 uint64 = 0
        _ = v40
    case -673:
        if (true || (v32 < (-(v32)))) {
            v3 -= v3
            v32 -= (v32 ^ (^(v32)))
        }
    }
    fmt.Println(v34)
    for fi41 := 0; fi41 < 10; fi41++ {
        if (!(-88 != 1)) {
            if (!(42 != (v3 / (((v3 | v3)) | 1)))) {
                var v42 uint8 = (^(v3))
                fmt.Println(v32, v3)
                _ = v42
            } else {
                var v43 uint64 = 323
                _ = v43
            }
            fmt.Println(fi41)
            v1[""] = fi41
        }
        var v44 int64 = 42
        _ = fi41
        _ = v44
    }
    delete(v1, ("hello" + "a"))
    fmt.Println(v2, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v45 := GenS3{F0: "", F1: v32}
    v46 := GenS3{F0: ((("a" + "") + "a") + ("foo bar" + "a")), F1: -432}
    for fi47 := 0; fi47 < 13; fi47++ {
        var v48 uint64 = 0
        v34 = (v34 | (v34 - v34))
        fmt.Println(v3)
        _ = fi47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v4
    _ = v45
    _ = v46
}
