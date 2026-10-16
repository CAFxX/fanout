package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint32; F1 string; F2 uint }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 string; F1 uint32; F2 int16 }

func main() {
    v1 := []uint32{371, 43, 2147483647, 942}
    var v2 uint = 2
    var v3 string = "a"
    var v4 float64 = 0.0
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: 18446744073709551615, F1: -1}
        if (42 >= 0) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            for fi7 := 0; fi7 < 3; fi7++ {
                if true {
                    fmt.Println(v4)
                    var v8 uint8 = 120
                    _ = len(v1)
                    fmt.Println(v2, v4, v3)
                    var v9 int16 = 0
                    _ = v8
                    _ = v9
                }
                _ = fi7
            }
            for fi10 := 0; fi10 < 19; fi10++ {
                v2 += (v2 << ((353) & 63))
                var v11 uint8 = 1
                fmt.Println(fi5, v4)
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 20; fi12++ {
                v3 = "a"
                if (false && ((0 > -50) && true)) {
                    v1 = append(v1, 1)
                    v13 := GenS2{F0: 590, F1: v3, F2: 255}
                    v1 = append(v1, 65535)
                    _ = v13
                } else {
                    var v14 int8 = 1
                    _ = v14
                }
                if false {
                    if true {
                        v4 = ((3.14159 * (v4 * v4)) * v4)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        fmt.Println(v2, v3, fi5)
                    }
                    fmt.Println(fi5)
                    var v15 bool = (!true)
                    _ = v15
                } else {
                    var v16 uint32 = 538
                    _ = v16
                }
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
                v2 = (v2 - v2)
                _ = fi12
            }
        }
        _ = fi5
        _ = v6
    }
    if len(v1) > 0 {
        v1[0] = 622
    }
    var v17 float64 = v4
    var v18 string = (("" + "a") + (v3 + "foo bar"))
    var v19 uint16 = 32767
    for fi20 := 0; fi20 < 3; fi20++ {
        if true {
            v17 -= ((v17 + (v17 * (float64((1.5 - (0.0 - -2.25))) - 0.1))) - float64(3.14159))
        } else {
            v1 = append(v1, 92)
            v21 := GenS2{F0: 673, F1: (v3 + (v18 + v18)), F2: 2}
            _ = v21
        }
        if true {
            fmt.Println(v2, fi20, v19)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        switch "a" {
        case "hello":
            switch fi20 {
            case -658:
                v22 := GenS2{F0: 690, F1: (v18 + v18), F2: (v2 % (((v2 >> ((65535) & 63))) | 1))}
                _ = v22
            case 127:
                if (100 == -1) {
                    v1 = append(v1, 65535)
                } else {
                    v23 := GenS3{F0: false, F1: -128}
                    if true {
                        var v24 uint = v2
                        _ = v24
                    }
                    _ = v23
                }
            }
        case "":
            fmt.Println(v19, v3, v2)
        case "a":
            var v25 float32 = (-2.25 / 1.5)
            _ = v25
        }
        _ = fi20
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 float32 = 0.1
        var v28 uint16 = (v19 | (v19 >> ((v2) & 63)))
        if (32767 >= 148) {
            v29 := GenS4{F0: "a", F1: 4294967295, F2: 1}
            switch ("a" + (v18 + (((("" + v18) + v18) + ("foo bar" + (v3 + (v3 + ("\n\t" + "foo bar"))))) + v3))) {
            case "a":
                v18 = v18
            default:
                v30 := GenS1{F0: 154, F1: 1000}
                _ = v30
            }
            _ = v29
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        v32 := GenS3{F0: true, F1: 127}
        _ = fi31
        _ = v32
    }
    if false {
        var v33 int16 = -925
        _ = v33
    }
    fmt.Println(v19)
    v18 = ((v3 + "\n\t") + v18)
    fmt.Println(v18, v19, v4)
    for fi34 := 0; fi34 < 14; fi34++ {
        _ = len(v1)
        _ = fi34
    }
    v35 := GenS3{F0: true, F1: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
