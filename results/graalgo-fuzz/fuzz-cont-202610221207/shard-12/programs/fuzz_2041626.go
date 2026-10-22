package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int8; F2 uint8 }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: "", F1: 857, F2: -1}
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 42
    v10 := GenS2{F0: false}
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 int8 = -128
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v9
    return 255
}
func genFn1() bool {
    switch 4294967295 {
    case 4294967295:
        v13 := []int{}
        _ = v13
    case 9223372036854775807:
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 bool = true
            if v15 {
                var v16 bool = false
                if (-32768 >= 0) {
                    v16 = ("\n\t" != "hello")
                    if (!("\n\t" >= "\n\t")) {
                        v17 := GenS2{F0: v15}
                        v16 = true
                        var v18 int64 = -935
                        _ = v17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v15 = false
                        _ = fi19
                    }
                } else {
                    fmt.Println(v15, fi14, v15)
                }
                v16 = true
                _ = v16
            } else {
                v20 := []int{(fi14 >> ((255) & 63)), -2147483648}
                v15 = (42 <= 9223372036854775807)
                _ = v20
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                switch fi14 {
                case 32768:
                    if v15 {
                        v15 = (107 != 18446744073709551615)
                    }
                default:
                    fmt.Println(fi21, fi14)
                }
                fmt.Println(fi21, fi14)
                fmt.Println(fi14, v15)
                _ = fi21
            }
            _ = fi14
            _ = v15
        }
    case 517:
        var v22 float64 = 3.14159
        _ = v22
    }
    v23 := []int{-9223372036854775808, -122, 1}
    var v24 uint8 = 127
    _ = v23
    _ = v24
    return ("\n\t" >= ("a" + "foo bar"))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 851
    var v3 string = "hello"
    var v4 uint8 = 1
    var v5 int64 = -353
    v6 := make(map[uint32]uint32)
    var v25 uint8 = v4
    for fi26 := 0; fi26 < 1; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            v28 := GenS2{F0: true}
            fmt.Println(v25, v5)
            v4 = v25
            _ = fi27
            _ = v28
        }
        var v29 string = ("foo bar" + ((((v3 + v3) + v3) + ((v3 + v3) + v3)) + v3))
        v2 = v2
        _ = fi26
        _ = v29
    }
    if (-128 < -20) {
        for fi30 := 0; fi30 < 11; fi30++ {
            v31 := GenS3{F0: v4}
            _ = fi30
            _ = v31
        }
    } else {
        v4 -= v4
        for fi32 := 0; fi32 < 11; fi32++ {
            v25 = genFn0()
            switch "foo bar" {
            case "a":
                for fi33 := 0; fi33 < 14; fi33++ {
                    var v34 uint = 4294967295
                    v25 -= (v25 ^ (v4 | (v25 >> ((v34) & 63))))
                    _ = fi33
                    _ = v34
                }
            case "\n\t":
                v35 := []int{fi32, 128, fi32, fi32}
                _ = v35
            case "hello":
                switch ((v3 + v3) + "foo bar") {
                case "\n\t":
                    for fi36 := 0; fi36 < 11; fi36++ {
                        var v37 int = (fi32 / ((9223372036854775807) | 1))
                        v6[646] = 42
                        v38 := GenS4{F0: 485, F1: -69, F2: (v25 >> ((255) & 63))}
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                case "a":
                    v4 += 2
                default:
                    var v39 float32 = (-2.25 * 1e10)
                    _ = v39
                }
            }
            fmt.Println(fi32)
            _ = fi32
        }
    }
    v40 := GenS2{F0: (42 == 100)}
    v3 = ((v3 + "\n\t") + (v3 + ("hello" + (v3 + ((v3 + (v3 + v3)) + "a")))))
    var v41 int64 = -1
    fmt.Println(v5, v41, v2)
    if true {
        v3 = v3
        if (!false) {
            var v42 uint = 255
            _ = v42
        }
        v43 := []int{}
        _ = v43
    }
    fmt.Println(v5, v2, v4)
    v44 := []int{0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v5
    _ = v6
}
