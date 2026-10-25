package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64, p1 int32) bool {
    switch 1 {
    case -2147483648:
        for fi6 := 0; fi6 < 5; fi6++ {
            switch 4294967295 {
            case 93:
                if false {
                    var v7 string = ""
                    var v8 uint = 255
                    _ = v7
                    _ = v8
                } else {
                    var v9 int64 = p0
                    _ = v9
                }
            case 32768:
                v10 := GenS1{F0: 1, F1: "a", F2: 42}
                _ = v10
            case 128:
                if false {
                    if false {
                        switch fi6 {
                        case 314:
                            var v11 int8 = 0
                            _ = v11
                        default:
                            var v12 float64 = ((1.5 + float64(0.0)) * float64(0.1))
                            _ = v12
                        }
                        v13 := GenS1{F0: 42, F1: "", F2: 18446744073709551615}
                        if true {
                            var v14 int32 = p1
                            _ = v14
                        }
                        _ = v13
                    }
                    p0 -= p0
                } else {
                    switch "" {
                    case "foo bar":
                        fmt.Println(p1, p0)
                    }
                }
            default:
                var v15 bool = (false || true)
                _ = v15
            }
            var v16 int = -240
            _ = fi6
            _ = v16
        }
    default:
        p0 = -1
    }
    v17 := GenS1{F0: 560, F1: "foo bar", F2: 2}
    if ((-(p1)) == 0) {
        v18 := GenS2{F0: p0}
        _ = v18
    } else {
        if true {
            if false {
                v19 := []int{42, -320, 32768}
                for fi20 := 0; fi20 < 13; fi20++ {
                    for fi21 := 0; fi21 < 13; fi21++ {
                        fmt.Println(fi20, p0)
                        fmt.Println(p0, p0, p1)
                        _ = fi21
                    }
                    if true {
                        var v22 int16 = -1
                        _ = v22
                    }
                    if true {
                        if (true || (1e10 < 0.1)) {
                            fmt.Println(p0, p0)
                            var v23 float32 = 1e10
                            v23 -= 1.5
                            _ = v23
                        } else {
                            var v24 int16 = 42
                            for fi25 := 0; fi25 < 12; fi25++ {
                                switch ("a" + "foo bar") {
                                case "":
                                    fmt.Println(fi20)
                                    fmt.Println(fi25, p1)
                                case "foo bar":
                                    v26 := GenS2{F0: 1}
                                    _ = v26
                                }
                                var v27 float64 = (0.1 + 0.0)
                                _ = fi25
                                _ = v27
                            }
                            _ = v24
                        }
                    }
                    _ = fi20
                }
                fmt.Println(p1, p1)
                _ = v19
            }
            fmt.Println(p0)
            fmt.Println(p1, p0)
        }
        for fi28 := 0; fi28 < 2; fi28++ {
            p0 = -9223372036854775808
            v29 := GenS2{F0: p0}
            var v30 bool = (212 < 244)
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    p1 += (-(p1))
    _ = p0
    _ = p1
    _ = v17
    return (2 > 263)
}

func main() {
    v1 := []int{}
    var v2 int = 0
    v3 := []uint32{0, 2}
    var v4 int32 = 42
    var v5 string = ("foo bar" + "")
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 uint8 = 128
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v32 += 210
        _ = fi31
        _ = v32
    }
    v33 := GenS1{F0: 0, F1: "", F2: 128}
    v2 = (^(v2))
    var v34 int16 = -228
    var v35 uint8 = 255
    fmt.Println(v5, v5, v34)
    if false {
        v36 := GenS2{F0: -9223372036854775808}
        v5 = "foo bar"
        _ = len(v1)
        _ = v36
    } else {
        var v37 float32 = 0.1
        var v38 float64 = 3.14159
        var v39 int = v2
        _ = v37
        _ = v38
        _ = v39
    }
    v3 = append(v3, 0)
    v35 -= 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
