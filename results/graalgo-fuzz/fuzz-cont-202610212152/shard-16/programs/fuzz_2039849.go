package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int8; F1 int }
func genFn0(p0 uint64, p1 uint32, p2 uint8) int {
    var v5 uint64 = (p0 & 1)
    if false {
        if (-2.25 > ((0.1 + 1e10) + ((3.14159 / 1e10) * 0.0))) {
            p0 = v5
        }
        var v6 float32 = ((1e10 - (3.14159 / 1.5)) + (float32(((0.0 - (3.14159 - 0.0)) * 0.0)) * (3.14159 / 0.1)))
        v7 := GenS1{F0: 0, F1: -9223372036854775808}
        _ = v6
        _ = v7
    } else {
        p2 = (p2 % (((p2 % ((p2) | 1))) | 1))
        if false {
            p2 = (^(p2))
            var v8 int8 = 26
            if false {
                v9 := GenS2{F0: (v8 + (v8 + v8)), F1: 9223372036854775807}
                fmt.Println(p2, p2, v5)
                _ = v9
            }
            _ = v8
        } else {
            v10 := []int{-1}
            _ = v10
        }
        if ((0 < 18446744073709551615) && (-1 >= int64(p1))) {
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS2{F0: int8(p2), F1: fi11}
                v13 := GenS2{F0: 0, F1: (fi11 & fi11)}
                p2 += (^(p2))
                _ = fi11
                _ = v12
                _ = v13
            }
            p0 = 976
            v14 := []int{2147483647, -2147483648, 2147483647}
            _ = v14
        } else {
            var v15 uint64 = (^(p0))
            _ = v15
        }
    }
    if ((!(false && (true && (-9223372036854775808 >= -1)))) || true) {
        switch (("hello" + ("hello" + "foo bar")) + "a") {
        case "hello":
            fmt.Println(p0, p1, p1)
        case "\n\t":
            if (0 > p2) {
                if (!(false && (((!(p1 > p1)) || false) || ((p0 % ((uint64(p1)) | 1)) < p0)))) {
                    var v16 float32 = (0.0 * 0.0)
                    fmt.Println(p2, p1)
                    _ = v16
                }
            }
        default:
            fmt.Println(p2)
        }
    } else {
        p2 = 1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 32768
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 int = 127
    v4 := []string{(("foo bar" + "hello") + ("foo bar" + "")), "", "foo bar"}
    var v17 uint16 = 0
    delete(v2, 4294967295)
    var v18 int64 = 9223372036854775807
    var v19 int32 = -1
    _ = len(v1)
    for fi20 := 0; fi20 < 1; fi20++ {
        _ = len(v4)
        _ = fi20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        if ((32767 <= 42) && (42 < 42)) {
            v22 := GenS1{F0: 766, F1: v3}
            var v23 int8 = 100
            switch ("hello" + ("hello" + "")) {
            case "foo bar":
                fmt.Println(v19)
            }
            _ = v22
            _ = v23
        }
        v17 = v17
        var v24 uint = 129
        _ = fi21
        _ = v24
    }
    if (2 <= 805) {
        fmt.Println(v18, v19)
    }
    if false {
        if true {
            var v25 uint32 = 67
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 uint64 = 42
                if false {
                    v27 += v27
                }
                _ = fi26
                _ = v27
            }
            if false {
                var v28 float64 = float64((0.1 - 3.14159))
                var v29 int32 = 2147483647
                for fi30 := 0; fi30 < 9; fi30++ {
                    for fi31 := 0; fi31 < 2; fi31++ {
                        v29 += -1
                        switch fi30 {
                        case -464:
                            fmt.Println(v3, fi31)
                        case 32768:
                            var v32 uint16 = (^(v17))
                            _ = v32
                        default:
                            var v33 string = (("hello" + (("hello" + "\n\t") + "hello")) + "")
                            _ = v33
                        }
                        if ((v19 % (((v29 / (((-(v19))) | 1))) | 1)) <= v29) {
                            var v34 uint16 = v17
                            if len(v4) > 0 {
                                v4[0] = "\n\t"
                            }
                            _ = v34
                        }
                        _ = fi31
                    }
                    v28 -= -2.25
                    v19 += v29
                    _ = fi30
                }
                _ = v28
                _ = v29
            } else {
                var v35 string = "foo bar"
                var v36 uint16 = 2
                for fi37 := 0; fi37 < 12; fi37++ {
                    var v38 uint32 = (v25 | v25)
                    _ = fi37
                    _ = v38
                }
                _ = v35
                _ = v36
            }
            _ = v25
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v39 uint16 = v17
        _ = v39
    } else {
        v2[v3] = "\n\t"
        v40 := GenS1{F0: 65535, F1: (v3 | (-(v3)))}
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
