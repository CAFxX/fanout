package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }
type GenS2 struct { F0 bool; F1 int64; F2 string }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 int64, p1 uint) int8 {
    p1 -= p1
    v4 := []int{736, 2147483647}
    var v5 int8 = -128
    if true {
        p1 += (^(p1))
        if true {
            if false {
                for fi6 := 0; fi6 < 12; fi6++ {
                    for fi7 := 0; fi7 < 3; fi7++ {
                        v4 = append(v4, (fi6 % ((fi6) | 1)))
                        _ = fi7
                    }
                    var v8 uint = p1
                    _ = fi6
                    _ = v8
                }
                for fi9 := 0; fi9 < 18; fi9++ {
                    for fi10 := 0; fi10 < 6; fi10++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(fi10)
                        fmt.Println(fi9)
                        _ = fi10
                    }
                    var v11 float64 = (1.5 - ((1.5 / 3.14159) - -2.25))
                    _ = fi9
                    _ = v11
                }
            }
        }
    }
    if true {
        var v12 int64 = 0
        _ = v12
    } else {
        if (!(!((!false) || (3.14159 != 1e10)))) {
            v5 += v5
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return v5
}

func main() {
    var v1 uint8 = 179
    var v2 float64 = -2.25
    var v3 string = (("hello" + "a") + ("a" + "\n\t"))
    if true {
        v2 += v2
        v13 := []int{456, 42}
        if (2 < 255) {
            if true {
                var v14 float64 = v2
                v1 += (v1 / (((^(v1))) | 1))
                if (!true) {
                    switch 127 {
                    case 1:
                        v3 = v3
                    }
                    fmt.Println(v1, v3)
                    var v15 float64 = v14
                    _ = v15
                } else {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                _ = v14
            } else {
                if len(v13) > 0 {
                    v13[0] = 0
                }
                v16 := GenS1{F0: 1, F1: int16(v1), F2: ((v3 + (v3 + "foo bar")) + (((((("foo bar" + v3) + (v3 + v3)) + "foo bar") + v3) + "a") + ((v3 + (("\n\t" + "\n\t") + v3)) + "foo bar")))}
                var v17 uint32 = 2147483647
                _ = v16
                _ = v17
            }
            fmt.Println(v2, v1, v2)
            if (true && true) {
                switch "foo bar" {
                case "":
                    fmt.Println(v3, v2)
                case "foo bar":
                    for fi18 := 0; fi18 < 9; fi18++ {
                        if (true || (!false)) {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            var v19 uint = 0
                            _ = v19
                        } else {
                            switch fi18 {
                            case 127:
                                fmt.Println(fi18, v2)
                            }
                            v13 = append(v13, (fi18 | fi18))
                        }
                        _ = fi18
                    }
                default:
                    v20 := GenS2{F0: true, F1: -451, F2: (v3 + ((((("\n\t" + "hello") + ((((v3 + v3) + (("" + ((((v3 + "hello") + v3) + v3) + v3)) + (("hello" + ("\n\t" + ("foo bar" + ("foo bar" + (v3 + ((v3 + v3) + v3)))))) + v3))) + ((v3 + v3) + "\n\t")) + v3)) + ("hello" + "\n\t")) + ("a" + "")) + "hello"))}
                    _ = v20
                }
                fmt.Println(v3, v2, v2)
            } else {
                var v21 uint16 = 48
                _ = v21
            }
        } else {
            for fi22 := 0; fi22 < 2; fi22++ {
                v13 = append(v13, fi22)
                v2 -= (v2 / (v2 + v2))
                _ = fi22
            }
            var v23 float32 = (1e10 - float32(((float64(1.5) + 0.1) * (v2 - ((v2 / (v2 - v2)) / v2)))))
            if true {
                switch v3 {
                case "a":
                    v23 = ((1e10 + v23) / (v23 + (1.5 + v23)))
                default:
                    if true {
                        for fi24 := 0; fi24 < 12; fi24++ {
                            var v25 int16 = 32767
                            var v26 int8 = -1
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    } else {
                        fmt.Println(v2, v1)
                        if true {
                            var v27 int16 = -247
                            _ = v27
                        }
                        v28 := GenS1{F0: 32767, F1: int16(v1), F2: v3}
                        _ = v28
                    }
                }
                var v29 uint = 244
                _ = v29
            }
            _ = v23
        }
        _ = v13
    } else {
        v30 := []int{0, 1}
        _ = v30
    }
    var v31 int = 127
    switch v3 {
    case "a":
        switch (("hello" + v3) + (("hello" + "a") + v3)) {
        case "":
            fmt.Println(v2, v1)
        case "a":
            v32 := GenS1{F0: 32767, F1: -310, F2: (("hello" + v3) + "\n\t")}
            _ = v32
        }
    }
    v33 := []int{(v31 >> ((4294967295) & 63))}
    _ = len(v33)
    v34 := GenS3{F0: v31, F1: 4294967295}
    var v35 bool = true
    var v36 uint16 = 255
    v36 = (v36 >> ((2) & 63))
    switch v3 {
    case "":
        v33 = append(v33, (v31 >> ((864) & 63)))
    default:
        var v37 bool = v35
        _ = v37
    }
    if (float64(1.5) <= (0.0 * v2)) {
        v33 = append(v33, (v31 + 0))
        switch v3 {
        case "foo bar":
            for fi38 := 0; fi38 < 12; fi38++ {
                for fi39 := 0; fi39 < 2; fi39++ {
                    v33 = append(v33, v31)
                    _ = fi39
                }
                _ = fi38
            }
        case "\n\t":
            v40 := GenS2{F0: (!v35), F1: -563, F2: (v3 + v3)}
            _ = v40
        default:
            v35 = false
        }
        for fi41 := 0; fi41 < 8; fi41++ {
            fmt.Println(v35, v1)
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            _ = fi41
        }
    }
    v42 := GenS1{F0: 32767, F1: 1000, F2: v3}
    var v43 float32 = float32(-2.25)
    v33 = append(v33, v31)
    fmt.Println(v2)
    var v44 uint64 = 42
    v44 -= 1
    switch "\n\t" {
    case "hello":
        if v35 {
            var v45 uint16 = v36
            v33 = append(v33, 128)
            _ = v45
        }
    case "":
        v46 := GenS1{F0: -1, F1: 0, F2: ""}
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v42
    _ = v43
    _ = v44
}
