package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 bool }
func genFn0() uint32 {
    var v4 int = 721
    for fi5 := 0; fi5 < 2; fi5++ {
        v4 += (-(fi5))
        _ = fi5
    }
    _ = v4
    return 0
}
func genFn1(p0 uint, p1 string, p2 int32) int8 {
    fmt.Println(p0)
    if true {
        switch ("a" + (("\n\t" + (p1 + (p1 + (p1 + p1)))) + p1)) {
        case "":
            var v6 int16 = -32768
            _ = v6
        default:
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := []int{(fi7 % ((0) | 1))}
                if true {
                    var v9 float32 = 0.1
                    _ = v9
                } else {
                    var v10 int8 = -119
                    if false {
                        p1 = (p1 + (("hello" + (p1 + ("hello" + (("a" + p1) + p1)))) + ("hello" + "a")))
                    }
                    _ = v10
                }
                fmt.Println(p2)
                _ = fi7
                _ = v8
            }
        }
        switch ("a" + (((("foo bar" + p1) + p1) + p1) + p1)) {
        case "a":
            v11 := GenS1{F0: ((p1 + p1) + p1)}
            _ = v11
        case "hello":
            var v12 int64 = 49
            _ = v12
        default:
            var v13 uint = p0
            _ = v13
        }
    } else {
        switch "" {
        case "\n\t":
            fmt.Println(p1, p1)
        default:
            v14 := []int{127, 127}
            _ = v14
        }
        if ((-116 == 127) || (!(587 != 65535))) {
            fmt.Println(p0, p0, p0)
        } else {
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 int16 = 1000
                fmt.Println(p2)
                for fi17 := 0; fi17 < 5; fi17++ {
                    fmt.Println(p2, v16)
                    p1 = p1
                    v18 := []int{(fi15 + (^(fi17)))}
                    _ = fi17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
            if (p0 == (p0 >> (((p0 & 243)) & 63))) {
                var v19 int = 250
                if true {
                    for fi20 := 0; fi20 < 14; fi20++ {
                        var v21 float64 = (float64(float32((1.5 * 1.5))) / 1e10)
                        var v22 bool = (1e10 != (float32(v21) - -2.25))
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v19 = (v19 & (-(v19)))
                }
                _ = v19
            }
            v23 := []int{-617, -1}
            _ = v23
        }
    }
    p1 = (p1 + p1)
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    var v1 int32 = 14
    var v2 uint32 = 908
    v3 := make(map[uint32]string)
    var v24 float64 = 3.14159
    v1 += v1
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := GenS2{F0: 128}
        v2 = (v2 >> ((18446744073709551615) & 63))
        v24 -= (float64((float32(v24) * (-2.25 - 3.14159))) + v24)
        _ = fi25
        _ = v26
    }
    var v27 int8 = -128
    switch 622 {
    case 1:
        if (2 <= 211) {
            v28 := []int{1, 647, 702, 128}
            _ = v28
        } else {
            v29 := []int{}
            _ = v29
        }
    case 128:
        delete(v3, genFn0())
    default:
        if true {
            switch "foo bar" {
            case "a":
                v30 := GenS3{F0: 32767, F1: true}
                _ = v30
            case "\n\t":
                v31 := GenS3{F0: -1, F1: true}
                _ = v31
            case "hello":
                v32 := GenS3{F0: 527, F1: false}
                _ = v32
            default:
                var v33 int64 = 1
                _ = v33
            }
        }
    }
    v34 := GenS1{F0: (("a" + ("\n\t" + "\n\t")) + "a")}
    var v35 string = "hello"
    v3[genFn0()] = "foo bar"
    v36 := v3[320]
    v37 := GenS1{F0: "a"}
    v1 = 900
    v36 = "foo bar"
    _ = len(v3)
    var v38 uint16 = 65535
    for fi39 := 0; fi39 < 17; fi39++ {
        var v40 uint = 1
        for fi41 := 0; fi41 < 15; fi41++ {
            if false {
                v42 := v3[uint32(v40)]
                v3[(v2 ^ v2)] = (("a" + (v35 + "a")) + "a")
                var v43 int = fi41
                _ = v42
                _ = v43
            } else {
                var v44 uint16 = 0
                _ = v44
            }
            v45 := GenS3{F0: fi39, F1: (v24 >= 1.5)}
            _ = fi41
            _ = v45
        }
        v2 -= (v2 % ((2) | 1))
        _ = fi39
        _ = v40
    }
    v46 := GenS3{F0: 706, F1: true}
    v38 -= v38
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v46
}
