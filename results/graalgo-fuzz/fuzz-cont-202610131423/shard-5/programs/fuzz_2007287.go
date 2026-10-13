package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 string, p1 string) uint32 {
    if ((3.14159 / float64(1e10)) >= 3.14159) {
        var v7 uint64 = 0
        var v8 string = p1
        if true {
            var v9 float64 = -2.25
            if (false && false) {
                switch p0 {
                case "\n\t":
                    v8 = p1
                case "a":
                    v10 := []int{int(v7), 128, 2147483647}
                    _ = v10
                }
                if (397 <= 4294967295) {
                    fmt.Println(p0, v7, p1)
                    var v11 uint64 = 4294967295
                    if true {
                        v9 -= (v9 / 0.1)
                    }
                    _ = v11
                } else {
                    var v12 float32 = (-2.25 * 3.14159)
                    _ = v12
                }
            }
            _ = v9
        }
        _ = v7
        _ = v8
    } else {
        for fi13 := 0; fi13 < 7; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                if false {
                    v15 := GenS1{F0: 4294967295, F1: p0, F2: 210}
                    _ = v15
                } else {
                    switch "" {
                    case "":
                        v16 := []int{fi14}
                        _ = v16
                    case "hello":
                        if true {
                            switch fi13 {
                            case -2147483648:
                                fmt.Println(fi13, fi14, fi14)
                            case 9223372036854775807:
                                var v17 int64 = -1
                                _ = v17
                            }
                            var v18 int32 = 269
                            _ = v18
                        } else {
                            var v19 int8 = 0
                            v19 = v19
                            _ = v19
                        }
                    default:
                        if false {
                            fmt.Println(p0)
                            v20 := []int{}
                            _ = v20
                        } else {
                            var v21 uint8 = 42
                            p0 = p0
                            switch "foo bar" {
                            case "hello":
                                var v22 int = fi14
                                _ = v22
                            case "":
                                var v23 string = p1
                                _ = v23
                            default:
                                switch fi13 {
                                case 1:
                                    p1 = "foo bar"
                                case 2147483647:
                                    v24 := GenS1{F0: 211, F1: "a", F2: 6}
                                    _ = v24
                                }
                            }
                            _ = v21
                        }
                    }
                    if ((0 != 1) || (fi14 == 926)) {
                        fmt.Println(fi14, p1)
                    }
                }
                _ = fi14
            }
            _ = fi13
        }
        v25 := GenS2{F0: -1}
        _ = v25
    }
    fmt.Println(p1, p1)
    if (2 <= 762) {
        v26 := GenS3{F0: (p1 != ("\n\t" + p0)), F1: 2147483647}
        if (531 < 650) {
            v27 := []int{2147483647, 555}
            _ = len(v27)
            v28 := GenS3{F0: true, F1: 273}
            _ = v27
            _ = v28
        } else {
            var v29 uint = 399
            var v30 uint = v29
            _ = v29
            _ = v30
        }
        var v31 int = 9223372036854775807
        _ = v26
        _ = v31
    } else {
        var v32 int = 748
        if true {
            for fi33 := 0; fi33 < 19; fi33++ {
                v34 := GenS2{F0: -2147483648}
                _ = fi33
                _ = v34
            }
            var v35 int32 = 141
            _ = v35
        }
        _ = v32
    }
    _ = p0
    _ = p1
    return 2147483648
}

func main() {
    v1 := []string{("" + "")}
    v2 := make(map[int]bool)
    var v3 float32 = 1e10
    v4 := make(map[int]bool)
    var v5 uint64 = 2
    var v6 int32 = -951
    v5 += (^(v5))
    var v36 uint16 = 2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((1.5 != 1.5) && false) {
        if (float32(1.5) > v3) {
            v5 += (v5 + 0)
            v36 -= v36
            v3 = 0.0
        } else {
            var v37 float64 = 0.1
            for fi38 := 0; fi38 < 16; fi38++ {
                switch ("a" + "foo bar") {
                case "a":
                    v37 += 0.1
                default:
                    fmt.Println(v6)
                }
                _ = fi38
            }
            _ = v37
        }
    } else {
        v39 := GenS2{F0: v6}
        _ = v39
    }
    v3 = v3
    v5 = 2
    fmt.Println(v6)
    v5 = v5
    var v40 uint16 = (v36 * v36)
    v41 := GenS3{F0: (!((v6 * v6) <= int32(v36))), F1: genFn0(("foo bar" + "\n\t"), "foo bar")}
    fmt.Println(v40, v40, v36)
    var v42 int32 = 0
    v3 = v3
    var v43 int16 = 32767
    var v44 int8 = 42
    v45 := GenS2{F0: v6}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
