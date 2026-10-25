package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int16; F1 string; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 int16; F2 uint }
func genFn0(p0 int8) int16 {
    if true {
        if ((p0 <= 46) || (42 >= 127)) {
            var v4 uint8 = 187
            v5 := GenS1{F0: 144, F1: 866}
            var v6 uint8 = (v4 % ((v4) | 1))
            _ = v4
            _ = v5
            _ = v6
        }
        switch "a" {
        case "\n\t":
            for fi7 := 0; fi7 < 20; fi7++ {
                for fi8 := 0; fi8 < 6; fi8++ {
                    var v9 string = "\n\t"
                    _ = fi8
                    _ = v9
                }
                p0 -= (-(p0))
                _ = fi7
            }
        }
        v10 := []int{}
        _ = v10
    }
    v11 := GenS2{F0: 1000, F1: "\n\t", F2: 1}
    if true {
        v12 := GenS3{F0: true}
        _ = v12
    }
    _ = p0
    _ = v11
    return 915
}
func genFn1() uint {
    var v13 uint16 = 255
    v13 = (^(v13))
    _ = v13
    return 202
}

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 0
    v3 := make(map[int]uint32)
    var v14 float32 = 1e10
    fmt.Println(v1)
    if false {
        switch "" {
        case "foo bar":
            if ((-2.25 > 3.14159) || true) {
                v1 = (v1 & (v1 + (^(v1))))
            }
        case "\n\t":
            v15 := GenS2{F0: 1, F1: "a", F2: 0}
            _ = v15
        case "":
            v2 = (v2 * 2)
        default:
            v14 = float32(0.0)
        }
        v16 := GenS4{F0: -128, F1: 32767}
        var v17 int = 32767
        _ = v16
        _ = v17
    }
    v18 := GenS4{F0: 1, F1: 65535}
    for fi19 := 0; fi19 < 10; fi19++ {
        if (!false) {
            if true {
                var v20 uint8 = (^(v2))
                var v21 uint = genFn1()
                _ = v20
                _ = v21
            } else {
                v2 = 255
            }
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 int8 = 127
        _ = fi22
        _ = v23
    }
    if true {
        var v24 int8 = -128
        v1 = v1
        v25 := GenS2{F0: 834, F1: ("hello" + (("\n\t" + ("a" + "a")) + "a")), F2: -1}
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        v2 = 0
        delete(v3, (fi26 / ((128) | 1)))
        _ = fi26
    }
    fmt.Println(v2)
    v27 := []int{}
    delete(v3, 127)
    for fi28 := 0; fi28 < 15; fi28++ {
        for fi29 := 0; fi29 < 13; fi29++ {
            if (0 != 578) {
                v30 := GenS1{F0: 128, F1: 0}
                if false {
                    var v31 uint32 = 501
                    _ = v31
                }
                var v32 int32 = -2147483648
                _ = v30
                _ = v32
            }
            v2 -= 0
            _ = fi29
        }
        if (42 > 0) {
            v14 = ((v14 * v14) * 1e10)
            v33 := v3[fi28]
            _ = v33
        }
        switch fi28 {
        case 0:
            v3[(-(fi28))] = 0
        }
        _ = fi28
    }
    v34 := GenS5{F0: (false && false)}
    if true {
        if false {
            var v35 float32 = -2.25
            for fi36 := 0; fi36 < 2; fi36++ {
                switch fi36 {
                case 0:
                    v14 = v35
                case -832:
                    fmt.Println(v35, v35, fi36)
                case 2147483647:
                    if true {
                        fmt.Println(v35, v14, v1)
                    }
                }
                fmt.Println(v14)
                _ = fi36
            }
            v37 := GenS3{F0: (false || false)}
            _ = v35
            _ = v37
        } else {
            v1 += (v1 ^ 2)
        }
        if (22 <= 1) {
            v38 := GenS6{F0: 1, F1: -129, F2: 65535}
            v39 := GenS3{F0: (!true)}
            _ = v38
            _ = v39
        }
    }
    v40 := v3[195]
    v41 := GenS3{F0: (false && false)}
    v1 += (v1 + v1)
    var v42 uint = genFn1()
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v40
    _ = v41
    _ = v42
}
