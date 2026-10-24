package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64; F1 int8; F2 bool }
type GenS10 struct { F0 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 165
    var v3 float64 = (-2.25 - (1.5 * 1.5))
    v4 := GenS1{F0: -144, F1: 211, F2: 821}
    switch 42 {
    case 128:
        switch 0 {
        case 42:
            v5 := GenS2{F0: 2, F1: 2, F2: 65535}
            _ = v5
        case 127:
            if false {
                if false {
                    if (4294967295 >= -496) {
                        switch "" {
                        case "a":
                            var v6 int16 = 1000
                            _ = v6
                        }
                        v2 -= 42
                        for fi7 := 0; fi7 < 14; fi7++ {
                            fmt.Println(v3)
                            v2 = (^(v2))
                            v8 := GenS2{F0: 2, F1: 1, F2: 255}
                            _ = fi7
                            _ = v8
                        }
                    } else {
                        v2 += v2
                    }
                }
                fmt.Println(v3, v2)
            } else {
                for fi9 := 0; fi9 < 16; fi9++ {
                    switch (fi9 + fi9) {
                    case 32768:
                        if (!true) {
                            v2 += v2
                            v2 += v2
                        } else {
                            if true {
                                switch fi9 {
                                case -1:
                                    v3 += (3.14159 / (((v3 * v3) * 0.0) * 3.14159))
                                case -2147483648:
                                    if true {
                                        v10 := GenS2{F0: 653, F1: 0, F2: 372}
                                        v3 = 1e10
                                        _ = v10
                                    } else {
                                        if (!((3.14159 != float64((1.5 * 1e10))) && (!true))) {
                                            v11 := GenS3{F0: 42, F1: 32767, F2: 146}
                                            _ = v11
                                        }
                                        fmt.Println(v2, fi9)
                                        delete(v1, ("foo bar" + ((("" + "hello") + (("\n\t" + "hello") + "")) + ("" + "foo bar"))))
                                    }
                                }
                            }
                            v12 := []int{fi9, (^(fi9)), fi9}
                            _ = v12
                        }
                    case 859:
                        switch (fi9 * -1) {
                        case 777:
                            v13 := GenS4{F0: fi9}
                            _ = v13
                        case 2147483647:
                            var v14 int16 = 0
                            _ = v14
                        case -2147483648:
                            v15 := GenS5{F0: 255, F1: 319}
                            _ = v15
                        default:
                            var v16 uint32 = 188
                            _ = v16
                        }
                    case 4294967295:
                        v17 := GenS6{F0: -32768, F1: 2147483648}
                        _ = v17
                    default:
                        var v18 uint8 = 127
                        _ = v18
                    }
                    v3 = 0.1
                    for fi19 := 0; fi19 < 20; fi19++ {
                        v20 := GenS6{F0: 27, F1: 2147483647}
                        v3 += v3
                        var v21 uint8 = 128
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    _ = fi9
                }
                if false {
                    if true {
                        v22 := GenS7{F0: 127}
                        v3 -= (v3 / v3)
                        v23 := GenS8{F0: -2147483648}
                        _ = v22
                        _ = v23
                    }
                    v24 := []int{458}
                    var v25 string = "hello"
                    _ = v24
                    _ = v25
                } else {
                    if (18446744073709551615 == 65535) {
                        v26 := []int{-1, 701, 4294967295}
                        delete(v1, "")
                        switch 128 {
                        case -2147483648:
                            v27 := GenS8{F0: 128}
                            _ = v27
                        case -421:
                            v3 -= -2.25
                        default:
                            v1[(("hello" + ("foo bar" + "hello")) + "\n\t")] = (((!(!true)) || (1e10 >= float32((v3 * (v3 - v3))))) || false)
                        }
                        _ = v26
                    } else {
                        var v28 uint = 65535
                        _ = v28
                    }
                }
                for fi29 := 0; fi29 < 12; fi29++ {
                    for fi30 := 0; fi30 < 2; fi30++ {
                        v2 = v2
                        _ = fi30
                    }
                    v3 = (float64(1e10) / 1e10)
                    _ = fi29
                }
            }
        case -16:
            fmt.Println(v3)
        }
    case 2147483647:
        if (2147483648 <= 4294967295) {
            for fi31 := 0; fi31 < 16; fi31++ {
                var v32 string = (("a" + (("hello" + "\n\t") + "foo bar")) + "a")
                v33 := GenS9{F0: 1, F1: -1, F2: (813 > 509)}
                _ = fi31
                _ = v32
                _ = v33
            }
            v34 := GenS7{F0: -13}
            v3 -= (v3 - (((v3 / (0.0 - (float64((1e10 + 0.0)) / ((1e10 / ((0.0 * v3) + 1e10)) + (-2.25 + float64(((float32((0.0 * v3)) * float32((0.0 / 3.14159))) - (1e10 / 1.5)))))))) / v3) * (3.14159 / v3)))
            _ = v34
        } else {
            var v35 bool = false
            _ = v35
        }
    case 127:
        var v36 uint8 = 0
        _ = v36
    }
    var v37 int32 = -805
    if false {
        var v38 uint8 = 8
        fmt.Println(v37)
        var v39 string = ("a" + ("a" + ("foo bar" + "foo bar")))
        _ = v38
        _ = v39
    }
    var v40 float64 = float64(float32(-2.25))
    for fi41 := 0; fi41 < 11; fi41++ {
        var v42 uint = 2
        v37 -= (v37 - (v37 >> ((65535) & 63)))
        for fi43 := 0; fi43 < 1; fi43++ {
            v44 := []int{(fi41 << (((v42 + (v42 * v42))) & 63)), 32767, 4294967295, fi43}
            var v45 int16 = 0
            _ = fi43
            _ = v44
            _ = v45
        }
        _ = fi41
        _ = v42
    }
    v46 := GenS10{F0: 20}
    for fi47 := 0; fi47 < 8; fi47++ {
        var v48 float64 = float64(1e10)
        if (955 == 18446744073709551615) {
            v2 -= v2
        }
        v49 := v1["a"]
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v46
}
