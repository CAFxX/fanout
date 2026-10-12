package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 int32; F2 uint8 }
func genFn0() int64 {
    var v4 bool = false
    v5 := []int{-2147483648, 4294967295}
    _ = v4
    _ = v5
    return -729
}
func genFn1(p0 int, p1 uint32) int16 {
    if true {
        v6 := GenS1{F0: 176}
        fmt.Println(p0, p1)
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        p0 += fi7
        v8 := []int{p0, fi7}
        var v9 float64 = -2.25
        _ = fi7
        _ = v8
        _ = v9
    }
    p1 += 4294967295
    var v10 uint16 = 65535
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{}
        switch "hello" {
        case "\n\t":
            var v13 int32 = -2147483648
            _ = v13
        case "":
            v14 := GenS2{F0: genFn0()}
            _ = v14
        case "a":
            fmt.Println(p0, fi11)
        }
        var v15 uint8 = uint8(p1)
        _ = fi11
        _ = v12
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v10
    return -1
}

func main() {
    var v1 int16 = -700
    v2 := make(map[string]string)
    var v3 uint = 4294967295
    v3 = 283
    fmt.Println(v1, v3)
    if false {
        v1 = (v1 * v1)
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v1)
    default:
        var v16 uint16 = 304
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(v3, v3, fi17)
        v2["a"] = "a"
        _ = fi17
    }
    switch ("a" + ("hello" + ("foo bar" + "\n\t"))) {
    case "foo bar":
        v18 := v2["foo bar"]
        _ = v18
    case "\n\t":
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 18; fi20++ {
                v21 := GenS3{F0: fi19, F1: 606}
                var v22 int32 = -2147483648
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi19
        }
    case "a":
        v23 := GenS4{F0: (v3 % ((v3) | 1)), F1: 0}
        _ = v23
    default:
        switch 32767 {
        case 128:
            fmt.Println(v3, v1)
        case -1:
            if false {
                _ = len(v2)
            } else {
                var v24 int8 = -128
                _ = v24
            }
        case 42:
            for fi25 := 0; fi25 < 15; fi25++ {
                v1 += (^(v1))
                v26 := GenS3{F0: 711, F1: int32(v3)}
                _ = fi25
                _ = v26
            }
        default:
            v27 := GenS3{F0: 32767, F1: -407}
            _ = v27
        }
    }
    var v28 uint32 = 620
    v3 = 302
    for fi29 := 0; fi29 < 11; fi29++ {
        if false {
            var v30 int = 128
            for fi31 := 0; fi31 < 11; fi31++ {
                if (!(0.1 > 0.1)) {
                    var v32 uint8 = 114
                    var v33 string = ""
                    var v34 uint64 = 150
                    _ = v32
                    _ = v33
                    _ = v34
                } else {
                    var v35 float64 = (0.0 + 0.0)
                    v3 = (v3 & (v3 % (((v3 >> ((v3) & 63))) | 1)))
                    v36 := []int{v30, v30}
                    v37 := GenS5{F0: (v30 | 1)}
                    _ = v35
                    _ = v36
                    _ = v37
                }
                v38 := GenS6{F0: (v28 % (((v28 | 315)) | 1)), F1: 1, F2: 1}
                _ = fi31
                _ = v38
            }
            for fi39 := 0; fi39 < 18; fi39++ {
                if ((!((v1 ^ 42) <= v1)) && (2 != 952)) {
                    var v40 bool = (3.14159 < float64(0.0))
                    _ = v40
                } else {
                    var v41 bool = (2 > uint64(v3))
                    v2["a"] = (((("hello" + "\n\t") + ("hello" + "a")) + "foo bar") + "\n\t")
                    fmt.Println(v3, v3, v30)
                    if (9223372036854775807 == genFn0()) {
                        v42 := []int{}
                        _ = v42
                    } else {
                        var v43 bool = v41
                        fmt.Println(v28)
                        v43 = ((false || (-9223372036854775808 <= 9223372036854775807)) || true)
                        _ = v43
                    }
                    _ = v41
                }
                v44 := GenS6{F0: 2147483647, F1: 1, F2: 127}
                var v45 string = (("\n\t" + (("hello" + ("foo bar" + "foo bar")) + "")) + ("foo bar" + ("\n\t" + ("hello" + "a"))))
                _ = fi39
                _ = v44
                _ = v45
            }
            _ = v30
        } else {
            var v46 int64 = 945
            _ = v46
        }
        for fi47 := 0; fi47 < 6; fi47++ {
            v48 := []int{fi29, fi29, (fi29 / (((-(fi29))) | 1)), 128}
            _ = fi47
            _ = v48
        }
        if (!true) {
            var v49 float32 = 1e10
            _ = v49
        }
        _ = fi29
    }
    for fi50 := 0; fi50 < 20; fi50++ {
        v51 := []int{(fi50 - (-(fi50))), int(v28), fi50, fi50}
        _ = fi50
        _ = v51
    }
    var v52 uint16 = 0
    for fi53 := 0; fi53 < 5; fi53++ {
        for fi54 := 0; fi54 < 7; fi54++ {
            if false {
                v1 = -142
            }
            _ = fi54
        }
        for fi55 := 0; fi55 < 8; fi55++ {
            var v56 int = fi55
            if (!true) {
                var v57 uint32 = (v28 * v28)
                _ = v57
            }
            _ = fi55
            _ = v56
        }
        _ = fi53
    }
    var v58 float32 = 1e10
    for fi59 := 0; fi59 < 5; fi59++ {
        var v60 uint64 = 340
        _ = fi59
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v52
    _ = v58
}
