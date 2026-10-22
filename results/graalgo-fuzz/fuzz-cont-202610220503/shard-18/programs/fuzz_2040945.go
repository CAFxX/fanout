package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool; F1 int8; F2 int8 }
func genFn0() bool {
    var v4 int = 701
    v4 += v4
    fmt.Println(v4)
    v5 := GenS1{F0: "\n\t", F1: ("\n\t" + ("hello" + "\n\t"))}
    v4 -= (v4 * (v4 & 42))
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 bool, p1 bool, p2 int) int {
    switch "foo bar" {
    case "":
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 float32 = (0.0 / float32(0.1))
            _ = fi6
            _ = v7
        }
    case "a":
        var v8 bool = p1
        _ = v8
    }
    p1 = false
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            if (-1 < -9223372036854775808) {
                fmt.Println(fi10, p0, p2)
            }
            _ = fi10
        }
        _ = fi9
    }
    v11 := GenS1{F0: "hello", F1: ""}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p2
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "a"
    v3 := []int{33}
    v2 = (v2 + "\n\t")
    var v12 uint16 = 65535
    for fi13 := 0; fi13 < 7; fi13++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v2 = ((((((v2 + (("\n\t" + (v2 + ("a" + v2))) + (v2 + "hello"))) + v2) + (((("foo bar" + v2) + (v2 + "a")) + v2) + v2)) + v2) + "\n\t") + (v2 + (v2 + v2)))
        _ = fi13
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v14 float64 = float64(((1.5 * (((-2.25 - 3.14159) / -2.25) - ((1e10 - (float32(1e10) - 0.0)) - (0.0 / -2.25)))) + (-2.25 / ((0.0 + ((0.1 / (0.0 - 1e10)) - 3.14159)) - (((1.5 / 0.1) + 1.5) + float32((1e10 - float64(3.14159))))))))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v15 := GenS1{F0: v2, F1: (v2 + v2)}
    if true {
        for fi16 := 0; fi16 < 12; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 int = genFn1(true, (!(!true)), fi16)
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        if true {
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 uint16 = 1
                _ = fi19
                _ = v20
            }
            if (1 >= 37) {
                v21 := v1[v2]
                for fi22 := 0; fi22 < 4; fi22++ {
                    if (-1 > 1) {
                        v21 = (874 < 958)
                    }
                    fmt.Println(v14)
                    v14 -= v14
                    _ = fi22
                }
                _ = v21
            } else {
                switch v2 {
                case "foo bar":
                    _ = len(v3)
                default:
                    fmt.Println(v2, v12)
                }
            }
            var v23 uint32 = 4294967295
            _ = v23
        } else {
            if true {
                switch 4294967295 {
                case 128:
                    var v24 float32 = 3.14159
                    _ = v24
                case -9223372036854775808:
                    v3 = append(v3, 1)
                case 127:
                    fmt.Println(v14, v2)
                }
                var v25 int16 = 32767
                if true {
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v25 -= (v25 - v25)
                        v27 := GenS2{F0: (v12 << ((0) & 63))}
                        _ = fi26
                        _ = v27
                    }
                    if false {
                        v12 -= (v12 ^ v12)
                        v2 = (v2 + "foo bar")
                    }
                } else {
                    v28 := GenS3{F0: (v2 + (v2 + ""))}
                    _ = v28
                }
                _ = v25
            } else {
                var v29 uint8 = 76
                _ = v29
            }
            v14 += (0.0 / 0.1)
        }
        v2 = (v2 + v2)
    } else {
        _ = len(v1)
        for fi30 := 0; fi30 < 7; fi30++ {
            switch ("" + "") {
            case "hello":
                switch "" {
                case "hello":
                    switch fi30 {
                    case -9223372036854775808:
                        v3 = append(v3, (-(fi30)))
                    default:
                        v12 = v12
                    }
                case "foo bar":
                    var v31 uint8 = 127
                    _ = v31
                case "a":
                    v32 := GenS2{F0: v12}
                    _ = v32
                }
            case "a":
                _ = len(v1)
            case "\n\t":
                v12 += v12
            default:
                if ((1000 != -1) && (("\n\t" == "a") && false)) {
                    v33 := GenS4{F0: 2, F1: v2}
                    _ = v33
                }
            }
            v34 := GenS3{F0: v2}
            fmt.Println(v12, v2)
            _ = fi30
            _ = v34
        }
        v3 = append(v3, 32768)
    }
    v35 := v1[(v2 + ("\n\t" + v2))]
    switch v2 {
    case "":
        var v36 float64 = float64((-2.25 + float32(v14)))
        _ = v36
    case "foo bar":
        fmt.Println(v14)
    case "\n\t":
        for fi37 := 0; fi37 < 19; fi37++ {
            for fi38 := 0; fi38 < 14; fi38++ {
                v39 := GenS3{F0: v2}
                v2 = "\n\t"
                _ = fi38
                _ = v39
            }
            fmt.Println(v14, v35, v12)
            delete(v1, v2)
            _ = fi37
        }
    default:
        var v40 float64 = (v14 + v14)
        _ = v40
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        v1["a"] = true
        fmt.Println(v2)
        switch "foo bar" {
        case "":
            delete(v1, (v2 + "hello"))
        }
        _ = fi41
    }
    v12 = 65535
    var v42 float32 = 0.1
    if true {
        if (true || (255 > 65535)) {
            switch 951 {
            case 469:
                v42 += float32(v14)
            }
        } else {
            v2 = v2
        }
        for fi43 := 0; fi43 < 13; fi43++ {
            v35 = v35
            for fi44 := 0; fi44 < 5; fi44++ {
                if (241 != 1) {
                    v14 += v14
                    v45 := GenS5{F0: fi43, F1: 0}
                    switch (fi44 * 9223372036854775807) {
                    case 0:
                        v46 := GenS6{F0: 2}
                        var v47 uint32 = 65535
                        v2 = ((v2 + "") + "")
                        _ = v46
                        _ = v47
                    }
                    _ = v45
                }
                _ = fi44
            }
            fmt.Println(v2)
            _ = fi43
        }
    }
    _ = len(v3)
    v48 := GenS6{F0: 127}
    v42 = (0.0 * (((-2.25 / ((1e10 / (1.5 / float32(v14))) / v42)) / v42) * 1e10))
    for fi49 := 0; fi49 < 16; fi49++ {
        if (1 == 4294967295) {
            if false {
                fmt.Println(v12)
                for fi50 := 0; fi50 < 14; fi50++ {
                    v35 = ((v2 + "") >= v2)
                    v51 := GenS6{F0: 255}
                    v1[(v2 + "hello")] = (2147483647 < 469)
                    _ = fi50
                    _ = v51
                }
            } else {
                v42 = v42
                var v52 string = "hello"
                _ = v52
            }
            fmt.Println(v42, v2)
            v2 = (((v2 + "a") + ("" + v2)) + v2)
        } else {
            var v53 float64 = v14
            v54 := GenS7{F0: (!(1 != 42)), F1: 100, F2: 0}
            var v55 int64 = -9223372036854775808
            _ = v53
            _ = v54
            _ = v55
        }
        _ = len(v3)
        fmt.Println(v35)
        _ = fi49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v42
    _ = v48
}
