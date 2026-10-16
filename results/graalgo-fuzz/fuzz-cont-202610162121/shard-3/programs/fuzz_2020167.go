package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int8; F2 string }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]int)
    var v3 float64 = 1.5
    v4 := make(map[string]int)
    var v5 uint64 = 2
    for fi6 := 0; fi6 < 9; fi6++ {
        v1 += (^(v1))
        _ = fi6
    }
    v7 := GenS1{F0: 0}
    fmt.Println(v5, v1)
    _ = len(v4)
    if (true && false) {
        v3 -= 3.14159
        v8 := GenS1{F0: 4294967295}
        _ = v8
    } else {
        v9 := GenS1{F0: 65535}
        _ = len(v4)
        var v10 bool = (true || (1 >= 4294967295))
        _ = v10
        _ = v9
    }
    var v11 string = ("hello" + ("" + ("hello" + "\n\t")))
    v1 -= (v1 % (((v1 ^ (v1 ^ (v1 * (v1 ^ 2))))) | 1))
    var v12 uint16 = uint16(v1)
    if true {
        var v13 int = int(v12)
        v14 := []int{}
        _ = v13
        _ = v14
    } else {
        v11 = "foo bar"
        v15 := GenS2{F0: ((v5 * 42) >= v5), F1: (v12 + uint16(v1))}
        var v16 uint64 = 0
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v1 = 1
        v18 := GenS3{F0: 2147483647, F1: -95, F2: ((("foo bar" + (("" + v11) + "foo bar")) + v11) + v11)}
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: (0 != 32767), F1: (v12 >> ((250) & 63))}
    v12 -= 66
    for fi20 := 0; fi20 < 10; fi20++ {
        if true {
            for fi21 := 0; fi21 < 7; fi21++ {
                v22 := []int{fi20, fi21, 102}
                var v23 uint = 2
                switch 686 {
                case -9223372036854775808:
                    switch v11 {
                    case "hello":
                        v23 -= v23
                    default:
                        v23 = v23
                    }
                }
                v24 := GenS2{F0: (!(v3 > (v3 - 1.5))), F1: v12}
                var v25 int64 = 323
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 16; fi26++ {
                var v27 string = ""
                if false {
                    switch "foo bar" {
                    case "":
                        var v28 uint8 = (v1 | v1)
                        _ = v28
                    case "\n\t":
                        if (213 >= 255) {
                            var v29 uint = 255
                            v3 = (((float64(0.1) / (v3 - (1e10 + ((((v3 + v3) * v3) + v3) + (((v3 + -2.25) - v3) / v3))))) * 1.5) + v3)
                            if (!true) {
                                var v30 int32 = -1
                                if (((v3 - float64((1e10 / 3.14159))) * (v3 * 1e10)) != (1e10 * v3)) {
                                    v31 := GenS3{F0: 2147483647, F1: 100, F2: ""}
                                    var v32 uint = 65535
                                    _ = v31
                                    _ = v32
                                } else {
                                    fmt.Println(v5, v27)
                                }
                                _ = v30
                            }
                            _ = v29
                        } else {
                            var v33 float32 = 3.14159
                            v4[("foo bar" + (v11 + v11))] = (fi26 & fi26)
                            v34 := GenS1{F0: 183}
                            _ = v33
                            _ = v34
                        }
                    case "a":
                        v35 := v2[(fi20 - fi20)]
                        _ = v35
                    default:
                        switch (fi20 << ((2) & 63)) {
                        case 778:
                            v1 = v1
                        case 86:
                            var v36 int = -968
                            _ = v36
                        }
                    }
                    v37 := GenS2{F0: true, F1: 0}
                    var v38 float32 = 1.5
                    _ = v37
                    _ = v38
                } else {
                    fmt.Println(v5, fi20)
                    fmt.Println(v27, fi26)
                }
                _ = fi26
                _ = v27
            }
        } else {
            v39 := GenS3{F0: 2, F1: 107, F2: (("a" + ("foo bar" + ((v11 + "foo bar") + v11))) + "a")}
            v40 := GenS2{F0: true, F1: (v12 >> ((uint(v5)) & 63))}
            _ = v39
            _ = v40
        }
        delete(v4, "\n\t")
        _ = fi20
    }
    var v41 int16 = -32768
    for fi42 := 0; fi42 < 17; fi42++ {
        v12 -= 32767
        _ = fi42
    }
    v11 = (v11 + ((v11 + v11) + "\n\t"))
    switch (((("" + v11) + v11) + v11) + "hello") {
    case "hello":
        switch v11 {
        case "hello":
            delete(v4, (v11 + v11))
        case "":
            for fi43 := 0; fi43 < 14; fi43++ {
                var v44 int64 = 1
                _ = fi43
                _ = v44
            }
        }
    case "\n\t":
        v45 := GenS4{F0: -2147483648}
        _ = v45
    case "foo bar":
        var v46 string = v11
        _ = v46
    }
    if (-1 > -6) {
        fmt.Println(v12, v11)
        var v47 uint8 = (v1 - (v1 >> ((1) & 63)))
        var v48 int = -759
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v7
}
