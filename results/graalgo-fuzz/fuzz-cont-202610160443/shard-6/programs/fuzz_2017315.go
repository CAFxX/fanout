package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint }
func genFn0() uint64 {
    var v7 uint = 955
    v8 := []int{32768, -425}
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int = fi9
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi9
        _ = v10
    }
    v7 = (v7 % (((^(v7))) | 1))
    for fi11 := 0; fi11 < 12; fi11++ {
        if len(v8) > 0 {
            v8[0] = (^(fi11))
        }
        switch ("" + ("hello" + "foo bar")) {
        case "hello":
            _ = len(v8)
        default:
            var v12 float64 = 3.14159
            _ = v12
        }
        if len(v8) > 0 {
            v8[0] = 32768
        }
        _ = fi11
    }
    _ = v7
    _ = v8
    return 882
}
func genFn1(p0 int16, p1 uint8, p2 uint16) uint16 {
    var v13 uint16 = 0
    var v14 float32 = 1e10
    var v15 int16 = (p0 % (((p0 % (((p0 + 1)) | 1))) | 1))
    if (1 > 42) {
        v16 := []int{2147483647}
        var v17 uint32 = 2147483648
        fmt.Println(v17, p1, v15)
        _ = v16
        _ = v17
    } else {
        if true {
            if (("hello" + "foo bar") == "") {
                v18 := []int{}
                _ = v18
            }
        } else {
            v14 -= float32(float64(((-2.25 - v14) / 0.1)))
            var v19 uint = 645
            _ = v19
        }
        if (!true) {
            switch 4294967295 {
            case 32768:
                v20 := []int{-645, 196}
                _ = v20
            case 128:
                v15 -= v15
            case 32767:
                if ((!false) || true) {
                    v21 := []int{-34}
                    p2 = 535
                    switch 2147483647 {
                    case 128:
                        _ = len(v21)
                    default:
                        p1 = p1
                    }
                    _ = v21
                } else {
                    var v22 uint8 = 164
                    v13 -= (^(v13))
                    _ = v22
                }
            }
            for fi23 := 0; fi23 < 4; fi23++ {
                v24 := GenS1{F0: 4294967295}
                v25 := []int{}
                _ = fi23
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 9; fi26++ {
                if false {
                    var v27 uint64 = 603
                    switch "" {
                    case "hello":
                        v28 := GenS1{F0: 65535}
                        _ = v28
                    case "foo bar":
                        fmt.Println(p2)
                    case "":
                        var v29 int8 = -1
                        _ = v29
                    default:
                        v30 := []int{fi26}
                        _ = v30
                    }
                    _ = v27
                }
                _ = fi26
            }
        }
        var v31 uint64 = genFn0()
        _ = v31
    }
    if false {
        v32 := []int{42}
        fmt.Println(p2)
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return 32767
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{128, 0, -9223372036854775808}
    var v3 int64 = 0
    var v4 int32 = 1
    v5 := make(map[int]int)
    v6 := []int{127, 42}
    var v33 float64 = (-2.25 / 1.0)
    for fi34 := 0; fi34 < 9; fi34++ {
        v3 -= (v3 - v3)
        v35 := GenS1{F0: 65535}
        v4 = (v4 & v4)
        _ = fi34
        _ = v35
    }
    _ = len(v2)
    v36 := GenS1{F0: 76}
    if (255 == 503) {
        if ("foo bar" <= "hello") {
            if (2147483648 == 65535) {
                v3 -= (v3 | (v3 / ((-9223372036854775808) | 1)))
                if true {
                    var v37 uint8 = 29
                    _ = v37
                }
            }
            if false {
                if (42 > (v4 - 2147483647)) {
                    var v38 int64 = v3
                    v38 = (v3 + v3)
                    v39 := GenS1{F0: 896}
                    _ = v38
                    _ = v39
                }
            } else {
                v3 = (v3 - 0)
            }
            for fi40 := 0; fi40 < 7; fi40++ {
                fmt.Println(v4, v33)
                var v41 uint = 2
                _ = fi40
                _ = v41
            }
        } else {
            v42 := GenS2{F0: 4294967295}
            _ = v42
        }
    }
    v43 := GenS2{F0: 1}
    var v44 int8 = 100
    var v45 int64 = (v3 / ((42) | 1))
    switch 118 {
    case -1:
        if len(v2) > 0 {
            v2[0] = -9223372036854775808
        }
    default:
        for fi46 := 0; fi46 < 7; fi46++ {
            v45 += int64(v44)
            _ = fi46
        }
    }
    for fi47 := 0; fi47 < 2; fi47++ {
        var v48 int = fi47
        var v49 int16 = -418
        fmt.Println(fi47)
        _ = fi47
        _ = v48
        _ = v49
    }
    if false {
        delete(v1, "\n\t")
        if (true && ("\n\t" == "hello")) {
            fmt.Println(v33, v45)
        } else {
            v44 += (v44 & v44)
        }
        var v50 float32 = 1.5
        _ = v50
    }
    v51 := GenS3{F0: true, F1: 255}
    switch ("" + "a") {
    case "foo bar":
        var v52 bool = ((v3 / ((v45) | 1)) >= -1)
        _ = v52
    case "\n\t":
        for fi53 := 0; fi53 < 11; fi53++ {
            var v54 float64 = 0.1
            _ = fi53
            _ = v54
        }
    default:
        v33 = (v33 + v33)
    }
    v45 = (^(v45))
    var v55 uint32 = 4294967295
    var v56 uint16 = 1
    v45 = (v45 / (((-(v45))) | 1))
    var v57 uint32 = v55
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v51
    _ = v55
    _ = v56
    _ = v57
    _ = v6
}
