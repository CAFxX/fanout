package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint; F2 bool }
func genFn0(p0 string, p1 uint8, p2 int64) string {
    p2 = (^(p2))
    var v4 float64 = ((0.1 * float64((1.5 + 0.1))) - float64(0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 + ("\n\t" + ("a" + (("a" + "") + (p0 + (("\n\t" + "a") + "\n\t"))))))
}
func genFn1(p0 uint32, p1 int) int16 {
    fmt.Println(p0, p0)
    for fi5 := 0; fi5 < 9; fi5++ {
        if false {
            var v6 float64 = float64((((1.5 - (1.5 / 0.1)) / 0.1) + 1e10))
            for fi7 := 0; fi7 < 20; fi7++ {
                fmt.Println(p1)
                fmt.Println(v6, p0)
                _ = fi7
            }
            v8 := GenS1{F0: 1}
            _ = v6
            _ = v8
        }
        _ = fi5
    }
    var v9 uint64 = 857
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    v1 := []int{4294967295, -92}
    v2 := make(map[string]string)
    var v3 uint16 = 42
    var v10 int16 = 32767
    v3 -= v3
    for fi11 := 0; fi11 < 18; fi11++ {
        switch int(v3) {
        case 128:
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 uint16 = v3
                _ = fi12
                _ = v13
            }
        case -678:
            v10 = v10
        case 2147483647:
            v3 = v3
        }
        if (!true) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v3 = 791
            var v14 int16 = 32767
            _ = v14
        }
        delete(v2, ("hello" + ("" + "")))
        _ = fi11
    }
    var v15 int64 = 631
    switch "a" {
    case "":
        switch 32768 {
        case 128:
            if len(v1) > 0 {
                v1[0] = 128
            }
        }
    case "hello":
        if len(v1) > 0 {
            v1[0] = 959
        }
    case "a":
        var v16 int8 = 1
        _ = v16
    }
    if (2 < 1) {
        var v17 int16 = genFn1(398, 941)
        var v18 float64 = -2.25
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 19; fi19++ {
            if (0 <= v15) {
                fmt.Println(v3)
            } else {
                v3 = v3
            }
            v10 += (-(v10))
            _ = fi19
        }
        v3 = (v3 - v3)
    }
    var v20 bool = (32767 <= 128)
    var v21 int64 = (-(v15))
    v10 = v10
    if false {
        if (!false) {
            v15 = 1
            if (!(!(0 >= 0))) {
                switch "hello" {
                case "":
                    var v22 int8 = 1
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 5; fi23++ {
                        for fi24 := 0; fi24 < 18; fi24++ {
                            v2[""] = "foo bar"
                            var v25 uint8 = 153
                            _ = fi24
                            _ = v25
                        }
                        v20 = false
                        for fi26 := 0; fi26 < 11; fi26++ {
                            fmt.Println(v10, v21)
                            _ = fi26
                        }
                        _ = fi23
                    }
                }
                v21 -= v15
            }
        } else {
            for fi27 := 0; fi27 < 5; fi27++ {
                _ = len(v2)
                _ = fi27
            }
            v28 := GenS2{F0: 42}
            _ = v28
        }
        v20 = (0 >= 1)
    } else {
        if (v20 || (1e10 > (((1e10 + 0.0) * -2.25) - (-2.25 - (0.0 / 1e10))))) {
            delete(v2, "")
            v29 := GenS3{F0: 255, F1: 389, F2: ((v3 >> ((186) & 63)) < 32767)}
            fmt.Println(v3, v21)
            _ = v29
        }
        v30 := GenS3{F0: (v3 / (((v3 % ((32767) | 1))) | 1)), F1: 2, F2: v20}
        switch "hello" {
        case "hello":
            if false {
                if true {
                    var v31 int8 = 1
                    fmt.Println(v15, v15)
                    _ = v31
                }
            }
        case "":
            v20 = false
        default:
            v20 = v20
        }
        _ = v30
    }
    fmt.Println(v3, v10)
    v32 := v2[((((("" + "hello") + (("\n\t" + (("\n\t" + ("" + "a")) + "a")) + "a")) + "\n\t") + "") + "")]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
}
