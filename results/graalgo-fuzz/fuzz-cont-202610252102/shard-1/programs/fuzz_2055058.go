package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint64; F2 uint }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 uint64, p1 uint64, p2 uint) uint64 {
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := []int{(^(fi5))}
            var v8 int = (fi5 * fi5)
            v7 = append(v7, fi6)
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    p0 = 1
    fmt.Println(p2)
    if false {
        if (false || true) {
            switch -921 {
            case 128:
                var v9 int32 = 42
                _ = v9
            case 32767:
                var v10 int8 = 0
                _ = v10
            default:
                v11 := GenS1{F0: 42}
                _ = v11
            }
            var v12 uint = uint(p0)
            _ = v12
        } else {
            switch -1 {
            case 693:
                v13 := GenS1{F0: 255}
                _ = v13
            case 32767:
                for fi14 := 0; fi14 < 2; fi14++ {
                    fmt.Println(p2)
                    var v15 int = (-(fi14))
                    v16 := GenS1{F0: p2}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case 1:
                var v17 int64 = 9223372036854775807
                _ = v17
            }
        }
        p1 -= (p1 * (p1 + (p1 % (((p0 | p1)) | 1))))
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 - p0)
}
func genFn1(p0 int32, p1 int64) int8 {
    p1 = (^(p1))
    p1 -= 0
    _ = p0
    _ = p1
    return -11
}

func main() {
    var v1 int = 32768
    var v2 string = "hello"
    v3 := make(map[int]uint32)
    var v4 string = v2
    v2 = ""
    v18 := GenS1{F0: 1}
    v1 = (v1 * v1)
    v4 = (((((v4 + v4) + "foo bar") + "foo bar") + v2) + ("a" + ("hello" + v4)))
    var v19 bool = true
    for fi20 := 0; fi20 < 5; fi20++ {
        switch (^(v1)) {
        case 401:
            var v21 bool = false
            _ = v21
        case 101:
            for fi22 := 0; fi22 < 19; fi22++ {
                v1 = fi22
                if false {
                    var v23 int64 = -500
                    var v24 int = 127
                    _ = len(v3)
                    v24 += fi22
                    _ = v23
                    _ = v24
                } else {
                    if false {
                        v1 = fi22
                        v25 := []int{}
                        fmt.Println(v19)
                        v26 := GenS1{F0: 896}
                        switch (fi22 << ((1) & 63)) {
                        case 128:
                            var v27 float32 = 0.1
                            _ = v27
                        case -971:
                            var v28 uint16 = 0
                            var v29 int8 = 0
                            switch fi22 {
                            case 128:
                                switch v4 {
                                case "a":
                                    var v30 int64 = 9223372036854775807
                                    _ = v30
                                case "":
                                    if v19 {
                                        _ = len(v3)
                                    }
                                case "hello":
                                    var v31 float64 = 1e10
                                    _ = v31
                                default:
                                    fmt.Println(v1, v1, v1)
                                }
                            case 348:
                                if v19 {
                                    delete(v3, fi22)
                                } else {
                                    var v32 uint8 = 127
                                    _ = v32
                                }
                            }
                            _ = v28
                            _ = v29
                        default:
                            v19 = ((2 >= 4294967295) || (true && ((42 == 18446744073709551615) && v19)))
                        }
                        _ = v25
                        _ = v26
                    }
                    v33 := []int{42, fi22, (v1 - fi20), v1}
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                    _ = v33
                }
                fmt.Println(fi20, v1, fi22)
                _ = fi22
            }
        default:
            v34 := []int{}
            _ = v34
        }
        for fi35 := 0; fi35 < 14; fi35++ {
            var v36 uint64 = 42
            var v37 uint = 2
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi20
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        if v19 {
            v39 := []int{-2147483648}
            v19 = v19
            _ = v39
        } else {
            v4 = "hello"
            v2 = (("\n\t" + "\n\t") + v4)
            v40 := GenS1{F0: 348}
            _ = v40
        }
        _ = fi38
    }
    v41 := GenS2{F0: -869}
    for fi42 := 0; fi42 < 9; fi42++ {
        var v43 int16 = 1000
        switch ((("" + "hello") + v4) + "foo bar") {
        case "hello":
            v44 := []int{fi42, fi42, (v1 - v1), 483}
            _ = v44
        }
        _ = fi42
        _ = v43
    }
    v2 = "\n\t"
    switch (v2 + v2) {
    case "hello":
        v19 = v19
    case "foo bar":
        if v19 {
            var v45 float64 = 0.1
            _ = v45
        }
    case "":
        for fi46 := 0; fi46 < 16; fi46++ {
            var v47 float64 = (1.5 - float64(1.5))
            v48 := []int{2147483647, (^(v1)), -1}
            if v19 {
                fmt.Println(v19)
                if false {
                    v49 := GenS1{F0: 176}
                    _ = v49
                } else {
                    v19 = v19
                    for fi50 := 0; fi50 < 1; fi50++ {
                        v47 = ((float64(1.5) * (((1e10 / v47) / v47) * 1.5)) / float64(-2.25))
                        if (65535 >= 65535) {
                            if true {
                                if true {
                                    var v51 int32 = 42
                                    _ = v51
                                }
                                _ = len(v3)
                            }
                        } else {
                            switch (v2 + ("a" + (v4 + "foo bar"))) {
                            case "a":
                                if ((0.0 + 1e10) >= -2.25) {
                                    v47 = (v47 - float64(0.1))
                                }
                            case "":
                                v52 := GenS2{F0: -66}
                                _ = v52
                            default:
                                fmt.Println(v19, v47)
                            }
                        }
                        _ = fi50
                    }
                }
                v53 := GenS2{F0: -1}
                _ = v53
            }
            _ = fi46
            _ = v47
            _ = v48
        }
    }
    if (255 >= 141) {
        var v54 uint16 = 0
        _ = v54
    }
    v55 := GenS3{F0: (v4 + v2), F1: 110, F2: 624}
    fmt.Println(v2, v2)
    var v56 int16 = 0
    v57 := GenS3{F0: ("a" + v2), F1: 1, F2: 4294967295}
    v58 := GenS4{F0: v4, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
