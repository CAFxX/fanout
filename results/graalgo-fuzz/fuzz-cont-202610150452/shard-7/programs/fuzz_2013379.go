package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
func genFn0(p0 int32, p1 bool, p2 int64) string {
    var v6 int32 = p0
    var v7 uint32 = 65535
    v8 := []int{1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return ("a" + ("hello" + "foo bar"))
}
func genFn1(p0 int16, p1 int16) int {
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := GenS1{F0: false, F1: 1, F2: 0}
            _ = fi10
            _ = v11
        }
        fmt.Println(p1, fi9)
        v12 := GenS1{F0: false, F1: 2147483647, F2: 690}
        _ = fi9
        _ = v12
    }
    if (95 != 4294967295) {
        v13 := GenS1{F0: false, F1: 2, F2: 1}
        v14 := GenS1{F0: (0 >= 927), F1: 0, F2: 0}
        _ = v13
        _ = v14
    } else {
        if false {
            if false {
                v15 := []int{32767, -2147483648}
                _ = v15
            }
        }
        var v16 int64 = -1
        var v17 uint16 = 42
        _ = v16
        _ = v17
    }
    var v18 uint64 = 2
    var v19 int64 = -1
    if true {
        if true {
            var v20 int8 = 42
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 uint8 = 42
                for fi23 := 0; fi23 < 6; fi23++ {
                    v24 := GenS1{F0: ((v22 >> ((789) & 63)) <= (v22 % (((v22 - (^(v22)))) | 1))), F1: 2, F2: 255}
                    if false {
                        switch fi21 {
                        case 1:
                            v25 := []int{(fi21 ^ fi23), 128, (^(fi21)), fi21}
                            _ = v25
                        default:
                            switch ("foo bar" + "hello") {
                            case "":
                                v26 := GenS1{F0: true, F1: 2, F2: 2}
                                _ = v26
                            case "foo bar":
                                v18 = v18
                            case "\n\t":
                                var v27 float32 = 1.5
                                _ = v27
                            default:
                                v28 := GenS2{F0: fi21, F1: v18}
                                _ = v28
                            }
                            v20 = v20
                            v18 = 4294967295
                            v19 += -9223372036854775808
                            v29 := GenS3{F0: (-(v19))}
                            if true {
                                p1 = (p1 ^ p1)
                                v30 := []int{(fi23 | (fi21 << ((2) & 63))), (fi21 - (-(fi21))), fi21}
                                _ = v30
                            } else {
                                p0 -= (p1 | (p1 >> ((310) & 63)))
                                var v31 string = "hello"
                                v22 = (^(v22))
                                _ = v31
                            }
                            var v32 uint8 = (^(v22))
                            var v33 uint32 = 2147483648
                            _ = v29
                            _ = v32
                            _ = v33
                        }
                        v34 := GenS3{F0: 1}
                        _ = v34
                    }
                    _ = fi23
                    _ = v24
                }
                _ = fi21
                _ = v22
            }
            var v35 uint16 = 1
            _ = v20
            _ = v35
        } else {
            fmt.Println(p0, v18)
            switch ("hello" + "") {
            case "foo bar":
                var v36 uint32 = 798
                _ = v36
            }
            var v37 int16 = p1
            _ = v37
        }
    } else {
        v38 := []int{32768}
        _ = v38
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return -404
}

func main() {
    var v1 int = -88
    var v2 uint32 = 743
    v3 := []uint32{0, v2, v2, v2}
    var v4 int8 = 42
    v5 := make(map[int]string)
    for fi39 := 0; fi39 < 20; fi39++ {
        var v40 int64 = 470
        v1 = v1
        if true {
            var v41 int64 = (v40 ^ v40)
            _ = v41
        } else {
            var v42 float32 = (0.1 - float32(1e10))
            _ = v42
        }
        _ = fi39
        _ = v40
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v43 int8 = -128
    fmt.Println(v4, v4)
    switch int(v4) {
    case 607:
        v5[32768] = "foo bar"
    }
    fmt.Println(v1, v43)
    if (true || (!(((-2.25 - -2.25) == (-2.25 / 1.0)) && ((1e10 * ((0.1 - (((float64((3.14159 / 3.14159)) * (1e10 - 0.1)) / 1.0) / ((1e10 - 0.0) + ((float64((1e10 * 3.14159)) - 0.1) - 3.14159)))) + -2.25)) != 0.1)))) {
        if false {
            v4 += v4
            var v44 float32 = (0.1 * 1.5)
            _ = v44
        } else {
            var v45 int64 = 0
            _ = v45
        }
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v46 := GenS1{F0: (42 != (v2 & 548)), F1: 2, F2: 4294967295}
            v3 = append(v3, v2)
            for fi47 := 0; fi47 < 20; fi47++ {
                v1 += fi47
                fmt.Println(v43)
                v48 := GenS1{F0: false, F1: v2, F2: 553}
                _ = fi47
                _ = v48
            }
            _ = v46
        }
        v4 = v43
    } else {
        switch -9223372036854775808 {
        case 32768:
            v5[(v1 >> ((42) & 63))] = ""
        case 128:
            fmt.Println(v4, v4)
        }
        if (0.0 > (3.14159 * (-2.25 - 1e10))) {
            var v49 int32 = 665
            var v50 string = (("hello" + ("" + "")) + "a")
            v51 := GenS4{F0: genFn1(-185, -412)}
            _ = v49
            _ = v50
            _ = v51
        }
    }
    for fi52 := 0; fi52 < 1; fi52++ {
        var v53 uint = 0
        v53 = v53
        _ = fi52
        _ = v53
    }
    v43 -= (-(v4))
    var v54 int64 = -9223372036854775808
    v4 += (^(v4))
    fmt.Println(v1, v43)
    v4 -= (^(v43))
    var v55 int64 = (v54 * -919)
    v3 = append(v3, 919)
    if (935 <= 1) {
        v55 = v54
    } else {
        if (((v1 % (((v1 + genFn1(1, 42))) | 1)) <= (v1 * int(v54))) || true) {
            var v56 float64 = 1.5
            _ = v56
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v54
    _ = v55
}
