package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 bool) string {
    var v4 uint = 4294967295
    if (42 < 1) {
        p0 = (-1 == 0)
        for fi5 := 0; fi5 < 12; fi5++ {
            v6 := []int{}
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 int = 589
            if (2 <= 42) {
                v8 -= (fi7 % ((-1) | 1))
                var v9 float32 = float32(-2.25)
                if ("hello" >= "") {
                    var v10 float64 = 1.5
                    _ = v10
                } else {
                    fmt.Println(fi7, fi7)
                    var v11 int64 = 9223372036854775807
                    _ = v11
                }
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    } else {
        if (910 >= 211) {
            switch ("foo bar" + (((("foo bar" + ("a" + "")) + "\n\t") + ((("foo bar" + ("a" + ("a" + "\n\t"))) + "a") + ("foo bar" + ("foo bar" + "")))) + ("a" + "\n\t"))) {
            case "a":
                if false {
                    v12 := GenS1{F0: 42, F1: 130}
                    v4 = v4
                    var v13 uint64 = 627
                    _ = v12
                    _ = v13
                }
            default:
                var v14 string = "hello"
                _ = v14
            }
            fmt.Println(v4, p0)
        }
        v15 := []int{0, 969, 32768, 194}
        var v16 float32 = float32(1e10)
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := []int{65, fi17}
        var v19 float32 = -2.25
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = v4
    return "hello"
}

func main() {
    var v1 int32 = 42
    var v2 int = 0
    var v3 int64 = 9223372036854775807
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := []int{(v2 ^ (fi20 ^ (fi20 | (v2 * -638)))), (v2 ^ v2), (v2 % (((-(v2))) | 1)), (fi20 % ((int(v3)) | 1))}
        fmt.Println(v1)
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: 128, F1: 127}
    if (0.0 != -2.25) {
        v3 -= v3
        if true {
            var v23 uint8 = 2
            if false {
                v24 := []int{v2}
                for fi25 := 0; fi25 < 9; fi25++ {
                    _ = len(v24)
                    switch fi25 {
                    case -1:
                        if false {
                            v26 := GenS1{F0: 42, F1: (v23 & 204)}
                            for fi27 := 0; fi27 < 12; fi27++ {
                                var v28 int32 = 1
                                v24 = append(v24, 9223372036854775807)
                                _ = fi27
                                _ = v28
                            }
                            v3 += (v3 - v3)
                            _ = v26
                        } else {
                            var v29 uint = 255
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                            _ = v29
                        }
                    }
                    _ = fi25
                }
                v2 = -9223372036854775808
                _ = v24
            }
            v3 = 1
            _ = v23
        }
    }
    v1 = (v1 / (((v1 | (v1 % (((v1 >> ((65535) & 63))) | 1)))) | 1))
    for fi30 := 0; fi30 < 19; fi30++ {
        v1 += 42
        for fi31 := 0; fi31 < 19; fi31++ {
            v32 := []int{(fi31 & (^(v2))), (v2 * 128), (fi31 | (v2 - (fi31 + fi31)))}
            if (64 != 42) {
                fmt.Println(fi30)
            }
            fmt.Println(fi31)
            v33 := GenS2{F0: 256}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        fmt.Println(fi34, v2)
        _ = fi34
    }
    switch "hello" {
    case "":
        v35 := []int{(v2 >> ((1) & 63))}
        _ = v35
    default:
        v1 -= v1
    }
    v1 = (v1 / (((v1 | v1)) | 1))
    var v36 bool = (true && false)
    fmt.Println(v3)
    if (("foo bar" + "foo bar") == ("" + ("" + "a"))) {
        for fi37 := 0; fi37 < 11; fi37++ {
            v38 := []int{fi37, fi37}
            _ = fi37
            _ = v38
        }
    } else {
        v39 := []int{(v2 >> ((65535) & 63))}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v36
}
