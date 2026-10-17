package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint16; F1 uint64; F2 bool }
type GenS6 struct { F0 int32; F1 int32; F2 uint64 }
type GenS7 struct { F0 string }
func genFn0(p0 bool) int64 {
    if ((942 == 254) || (!(((359 == 0) || (!p0)) && true))) {
        if p0 {
            v4 := GenS1{F0: 645, F1: (!false)}
            _ = v4
        }
        v5 := GenS1{F0: 1, F1: (!(209 > -1))}
        if ("a" > "\n\t") {
            p0 = (0.0 <= float32((0.1 / -2.25)))
            v6 := []int{}
            var v7 uint8 = 0
            _ = v6
            _ = v7
        }
        _ = v5
    }
    var v8 string = ""
    _ = p0
    _ = v8
    return -526
}
func genFn1(p0 int8, p1 string) uint32 {
    p1 = ("a" + (p1 + p1))
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(p0)
        if (75 >= 65535) {
            v10 := GenS1{F0: 428, F1: false}
            p1 = (("a" + ("\n\t" + p1)) + p1)
            _ = v10
        } else {
            p1 = ((p1 + (p1 + p1)) + "a")
            switch p1 {
            case "a":
                p1 = p1
            case "hello":
                var v11 bool = (42 == p0)
                _ = v11
            case "foo bar":
                v12 := GenS2{F0: 2, F1: 1}
                _ = v12
            }
            switch p1 {
            case "":
                switch fi9 {
                case 32767:
                    for fi13 := 0; fi13 < 9; fi13++ {
                        p0 = p0
                        var v14 uint16 = 65535
                        var v15 uint = 278
                        v15 += (v15 - v15)
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                case 2147483647:
                    if (-2147483648 != 2147483647) {
                        v16 := GenS3{F0: 0, F1: -1}
                        fmt.Println(fi9, p1, p1)
                        _ = v16
                    } else {
                        if (true && false) {
                            v17 := GenS3{F0: 873, F1: -2147483648}
                            _ = v17
                        } else {
                            v18 := GenS4{F0: 91, F1: 2}
                            v19 := GenS5{F0: 1, F1: 599, F2: (!(p1 < (p1 + p1)))}
                            fmt.Println(p0)
                            _ = v18
                            _ = v19
                        }
                        p0 += (p0 | (p0 | 1))
                    }
                }
            }
        }
        var v20 bool = (702 == 0)
        _ = fi9
        _ = v20
    }
    p0 += (p0 & 100)
    v21 := []int{-953, 4294967295}
    v22 := GenS3{F0: 65535, F1: 42}
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return 65535
}

func main() {
    var v1 float32 = (3.14159 * -2.25)
    var v2 int8 = -28
    var v3 uint8 = 255
    v2 = 0
    var v23 int64 = 0
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := []int{}
        _ = fi24
        _ = v25
    }
    fmt.Println(v1, v1)
    v26 := GenS6{F0: 725, F1: 973, F2: 83}
    var v27 float64 = (0.1 / 0.1)
    var v28 int32 = -564
    v29 := []int{32767, 9223372036854775807, 1, -2147483648}
    switch 128 {
    case 32768:
        if false {
            v3 += v3
            if (841 <= 4294967295) {
                if len(v29) > 0 {
                    v29[0] = 9223372036854775807
                }
                v3 += v3
                if true {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        v3 -= v3
                        var v31 int = (fi30 / ((-844) | 1))
                        _ = fi30
                        _ = v31
                    }
                }
            } else {
                v2 -= -1
                v32 := GenS6{F0: v28, F1: v28, F2: 1}
                v27 += v27
                _ = v32
            }
        } else {
            v23 -= 42
        }
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        if (1e10 != (v27 * v27)) {
            if ((("\n\t" + "foo bar") + "foo bar") == "foo bar") {
                var v34 string = "foo bar"
                for fi35 := 0; fi35 < 15; fi35++ {
                    v36 := GenS7{F0: "foo bar"}
                    if len(v29) > 0 {
                        v29[0] = 32768
                    }
                    v3 = (v3 % ((v3) | 1))
                    fmt.Println(fi35, fi35)
                    var v37 bool = false
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                _ = v34
            } else {
                for fi38 := 0; fi38 < 3; fi38++ {
                    fmt.Println(v23, v28, v1)
                    _ = fi38
                }
                var v39 uint8 = v3
                _ = v39
            }
            var v40 float64 = 1.5
            _ = v40
        }
        _ = fi33
    }
    v41 := GenS3{F0: 204, F1: 2147483647}
    v42 := GenS1{F0: 77, F1: (2 == 1)}
    fmt.Println(v3)
    var v43 uint = 466
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v41
    _ = v42
    _ = v43
}
