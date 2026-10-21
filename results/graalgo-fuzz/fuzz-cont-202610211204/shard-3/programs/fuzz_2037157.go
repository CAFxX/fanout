package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 string; F1 int; F2 uint8 }
func genFn0(p0 int32, p1 uint32, p2 uint32) int32 {
    fmt.Println(p1, p0, p1)
    v4 := []int{32768, 523, -152}
    _ = len(v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p0
}
func genFn1(p0 string) int16 {
    v5 := []int{}
    var v6 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 3; fi7++ {
        switch ("a" + (p0 + "\n\t")) {
        case "hello":
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS1{F0: 32767, F1: 1}
                _ = fi8
                _ = v9
            }
        case "a":
            p0 = p0
        default:
            if (false || ((true && false) || (0 > 100))) {
                if true {
                    v6 -= (v6 / ((v6) | 1))
                    p0 = p0
                }
                switch (fi7 | fi7) {
                case 2147483647:
                    v10 := GenS1{F0: 255, F1: 1}
                    _ = v10
                case -2147483648:
                    if (true || true) {
                        if false {
                            switch ((("foo bar" + ((p0 + p0) + "foo bar")) + (p0 + "")) + "") {
                            case "\n\t":
                                fmt.Println(v6)
                            case "foo bar":
                                fmt.Println(p0, v6, fi7)
                            case "a":
                                var v11 bool = false
                                _ = v11
                            }
                            v6 -= (v6 + v6)
                        }
                        if (false || (!false)) {
                            var v12 uint64 = 18446744073709551615
                            var v13 bool = (0 <= 42)
                            v13 = v13
                            _ = v12
                            _ = v13
                        }
                    } else {
                        if (true || false) {
                            var v14 uint16 = 42
                            _ = v14
                        }
                    }
                }
                v15 := GenS2{F0: 0, F1: (fi7 / (((fi7 ^ 32767)) | 1))}
                _ = v15
            }
        }
        if (2147483647 != -2147483648) {
            switch fi7 {
            case 9223372036854775807:
                v6 = (v6 * 42)
            }
            if (-864 >= -816) {
                var v16 int8 = 0
                if false {
                    v16 = 71
                } else {
                    v6 = (v6 << ((837) & 63))
                }
                _ = v16
            }
            v5 = append(v5, 229)
        } else {
            p0 = p0
        }
        p0 = ("hello" + ((p0 + p0) + (p0 + "foo bar")))
        _ = fi7
    }
    _ = p0
    _ = v5
    _ = v6
    return -32768
}

func main() {
    v1 := []int{128}
    v2 := make(map[int]string)
    var v3 uint64 = 18446744073709551615
    fmt.Println(v3)
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 += 1
        if false {
            v17 := GenS3{F0: (v3 & v3)}
            var v18 int16 = -32768
            var v19 uint32 = 582
            _ = v17
            _ = v18
            _ = v19
        } else {
            if true {
                v3 -= v3
                var v20 string = ("\n\t" + ("hello" + "\n\t"))
                _ = v20
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        switch "" {
        case "hello":
            if false {
                if len(v1) > 0 {
                    v1[0] = fi21
                }
                v22 := GenS4{F0: (^(v3)), F1: (fi21 ^ fi21)}
                for fi23 := 0; fi23 < 1; fi23++ {
                    switch fi21 {
                    case -911:
                        if (true || (!true)) {
                            v24 := GenS4{F0: (v3 * 4294967295), F1: 4294967295}
                            var v25 uint8 = 128
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v24
                            _ = v25
                        } else {
                            if true {
                                fmt.Println(fi21, fi21)
                                switch (fi21 - fi23) {
                                case 32768:
                                    if true {
                                        v26 := GenS1{F0: 0, F1: 199}
                                        var v27 int8 = -128
                                        _ = v26
                                        _ = v27
                                    }
                                default:
                                    switch (fi21 % ((fi21) | 1)) {
                                    case -928:
                                        v28 := GenS3{F0: 0}
                                        _ = v28
                                    case 9223372036854775807:
                                        v3 += (^(v3))
                                    }
                                }
                                if (255 == 2) {
                                    v29 := v2[fi21]
                                    var v30 float32 = -2.25
                                    switch (fi21 % ((fi21) | 1)) {
                                    case -1:
                                        v31 := GenS5{F0: -2147483648, F1: -109}
                                        _ = v31
                                    }
                                    _ = v29
                                    _ = v30
                                } else {
                                    fmt.Println(fi21, fi21, v3)
                                    var v32 uint = 0
                                    _ = v32
                                }
                            }
                        }
                    case 780:
                        v3 += v3
                    }
                    fmt.Println(v3, v3, fi23)
                    _ = fi23
                }
                _ = v22
            }
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        for fi33 := 0; fi33 < 18; fi33++ {
            var v34 bool = (42 == 42)
            fmt.Println(v34, v34, v34)
            switch "" {
            case "":
                if v34 {
                    v35 := GenS6{F0: 42, F1: (v3 >> ((2) & 63))}
                    v36 := GenS7{F0: "hello", F1: (fi33 + fi21), F2: 16}
                    _ = v35
                    _ = v36
                }
            case "\n\t":
                var v37 float64 = 1e10
                _ = v37
            }
            _ = fi33
            _ = v34
        }
        fmt.Println(fi21, v3)
        _ = fi21
    }
    v38 := GenS2{F0: 494, F1: -1}
    var v39 int = 9223372036854775807
    var v40 float32 = (1.5 * 3.14159)
    if false {
        var v41 int64 = -204
        _ = v41
    }
    var v42 uint16 = 2
    var v43 int32 = -1
    var v44 uint = 745
    v42 = 142
    v39 = int(v44)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
    _ = v42
    _ = v43
    _ = v44
}
