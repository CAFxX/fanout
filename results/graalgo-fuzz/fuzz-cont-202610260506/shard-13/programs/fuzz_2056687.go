package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint8, p2 int) uint {
    if p0 {
        p0 = (p0 && p0)
    } else {
        switch (p2 / ((p2) | 1)) {
        case -1:
            if ((-128 >= 42) || (71 != 255)) {
                v7 := GenS1{F0: ("a" + (("\n\t" + ("\n\t" + ("hello" + ("" + ((("" + "foo bar") + (("hello" + "a") + "\n\t")) + "foo bar"))))) + "\n\t")), F1: (p1 * p1)}
                fmt.Println(p2)
                var v8 uint8 = (p1 ^ p1)
                _ = v7
                _ = v8
            }
        case 4294967295:
            if p0 {
                if p0 {
                    switch ("a" + "hello") {
                    case "":
                        v9 := GenS1{F0: ("" + ("hello" + "foo bar")), F1: (p1 >> ((255) & 63))}
                        _ = v9
                    case "foo bar":
                        p0 = true
                    default:
                        fmt.Println(p0, p0, p0)
                    }
                    v10 := []int{p2, p2, -9223372036854775808}
                    if p0 {
                        v11 := GenS2{F0: -9223372036854775808}
                        _ = v11
                    } else {
                        if p0 {
                            var v12 uint = 2
                            var v13 uint32 = 323
                            _ = v12
                            _ = v13
                        }
                        v14 := GenS1{F0: "foo bar", F1: 81}
                        _ = v14
                    }
                    _ = v10
                }
                p1 = 255
            } else {
                for fi15 := 0; fi15 < 6; fi15++ {
                    if (p1 == p1) {
                        var v16 uint = 328
                        p0 = p0
                        for fi17 := 0; fi17 < 1; fi17++ {
                            fmt.Println(v16)
                            v16 -= (v16 << ((v16) & 63))
                            var v18 uint8 = p1
                            _ = fi17
                            _ = v18
                        }
                        _ = v16
                    }
                    var v19 uint32 = 898
                    _ = fi15
                    _ = v19
                }
                if p0 {
                    var v20 float64 = (3.14159 / 0.1)
                    _ = v20
                }
                fmt.Println(p2, p2)
            }
        case 9223372036854775807:
            v21 := []int{(p2 * 42), (^(p2))}
            _ = v21
        }
    }
    fmt.Println(p1)
    v22 := []int{-725, 4294967295, p2, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 42
}
func genFn1(p0 uint64, p1 uint) int16 {
    var v23 bool = true
    v24 := GenS2{F0: 9223372036854775807}
    var v25 int8 = 1
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 int16 = -1
        v28 := []int{fi26, fi26, -831, (fi26 * (fi26 / (((fi26 & fi26)) | 1)))}
        for fi29 := 0; fi29 < 13; fi29++ {
            v23 = (float64(0.1) != 1.5)
            v30 := GenS3{F0: 0}
            _ = fi29
            _ = v30
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v25
    return 1000
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int = 32768
    var v3 uint = 264
    v4 := make(map[string]int)
    var v5 uint16 = 0
    v6 := make(map[string]uint32)
    if true {
        switch v2 {
        case 4294967295:
            var v31 int32 = 1
            _ = v31
        default:
            var v32 float32 = 0.1
            _ = v32
        }
    }
    var v33 bool = (genFn0((!(false && (int(v1) != 2147483647))), 81, -2147483648) == (v3 << ((v3) & 63)))
    for fi34 := 0; fi34 < 16; fi34++ {
        v3 = (v3 - v3)
        var v35 uint8 = 2
        _ = fi34
        _ = v35
    }
    v1 -= v1
    var v36 uint8 = 0
    v37 := []int{(v2 - v2), v2, -407}
    v38 := GenS3{F0: 0}
    var v39 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
