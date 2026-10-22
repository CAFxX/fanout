package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0() int {
    var v7 uint8 = 0
    fmt.Println(v7)
    v8 := []int{9223372036854775807, 1, 999}
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 uint = 65535
    v2 := []int{}
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    var v5 uint = v1
    v6 := make(map[string]int)
    if false {
        _ = len(v4)
    }
    if (false || true) {
        if (0.1 == float32(1e10)) {
            v9 := GenS1{F0: -1, F1: (true || (false && (false && false))), F2: 255}
            var v10 int16 = 730
            _ = v10
            _ = v9
        }
        if true {
            var v11 float64 = float64((0.1 / 3.14159))
            _ = v11
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            v3[("hello" + "foo bar")] = "a"
            var v13 int8 = 79
            v13 = (v13 / ((v13) | 1))
            _ = fi12
            _ = v13
        }
    }
    if ((v5 - v1) >= v5) {
        var v14 int8 = 100
        for fi15 := 0; fi15 < 1; fi15++ {
            v1 = 255
            switch "\n\t" {
            case "":
                switch ("hello" + (("foo bar" + "") + "hello")) {
                case "hello":
                    v16 := GenS2{F0: 928}
                    _ = v16
                case "\n\t":
                    if (-9223372036854775808 > -896) {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi17
                        }
                        var v18 int32 = 435
                        _ = v18
                    } else {
                        var v19 int32 = -2147483648
                        if (true && (true && ((!(((fi15 >> ((v1) & 63)) != genFn0()) && (!true))) || false))) {
                            fmt.Println(fi15, fi15)
                            var v20 int32 = (v19 * (v19 << (((v5 + (v5 % ((1) | 1)))) & 63)))
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 10; fi21++ {
                            var v22 uint8 = 249
                            _ = fi21
                            _ = v22
                        }
                        _ = v19
                    }
                }
            case "a":
                var v23 bool = (!(4294967295 > 25))
                _ = v23
            case "foo bar":
                for fi24 := 0; fi24 < 12; fi24++ {
                    v14 -= v14
                    _ = fi24
                }
            default:
                if len(v2) > 0 {
                    v2[0] = 32768
                }
            }
            switch "hello" {
            case "hello":
                if (496 <= 2147483647) {
                    var v25 string = ""
                    for fi26 := 0; fi26 < 13; fi26++ {
                        v14 -= v14
                        _ = fi26
                    }
                    _ = v25
                }
            case "foo bar":
                var v27 uint8 = 127
                _ = v27
            case "":
                fmt.Println(v14)
            }
            _ = fi15
        }
        _ = v14
    } else {
        var v28 uint16 = 2
        _ = v28
    }
    if (false || true) {
        v29 := v3[(("a" + ("hello" + "foo bar")) + ("a" + ""))]
        v2 = append(v2, 0)
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 1; fi31++ {
            var v32 string = "\n\t"
            _ = fi31
            _ = v32
        }
        if true {
            if (128 < 89) {
                v5 -= v1
            } else {
                switch "a" {
                case "\n\t":
                    v5 -= 1
                default:
                    var v33 uint = v1
                    _ = v33
                }
                switch ("\n\t" + "\n\t") {
                case "":
                    v34 := GenS2{F0: 1}
                    _ = v34
                }
            }
            v2 = append(v2, (fi30 | 2147483647))
            if len(v2) > 0 {
                v2[0] = (fi30 >> ((v5) & 63))
            }
        }
        fmt.Println(v5)
        _ = fi30
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 int8 = -74
        var v37 int64 = -1
        var v38 uint16 = 912
        _ = fi35
        _ = v36
        _ = v37
        _ = v38
    }
    v39 := v4[-1]
    var v40 uint64 = 1
    v1 = (v1 << ((809) & 63))
    var v41 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
