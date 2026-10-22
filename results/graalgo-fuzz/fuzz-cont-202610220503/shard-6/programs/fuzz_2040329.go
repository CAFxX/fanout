package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int8; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint32; F1 int64; F2 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool; F1 int16 }
type GenS10 struct { F0 uint16 }
func genFn0() int32 {
    var v4 int16 = 0
    v5 := GenS1{F0: 2, F1: 0}
    _ = v4
    _ = v5
    return 184
}
func genFn1() int {
    var v6 int64 = -1
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 739, F1: 484}
        v6 = -9223372036854775808
        v9 := []int{fi7, -1, fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    if (0 > 212) {
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS2{F0: 269}
            if (true || ((!((false && (!((1 == 38) || (107 > -21)))) && (682 >= 18446744073709551615))) || false)) {
                if (-833 >= -1) {
                    v12 := GenS3{F0: v6, F1: 127}
                    v13 := []int{(fi10 * -9223372036854775808), fi10}
                    _ = v12
                    _ = v13
                } else {
                    switch fi10 {
                    case 115:
                        v6 = v6
                    case 2147483647:
                        var v14 int8 = -1
                        _ = v14
                    case 217:
                        v15 := GenS1{F0: 0, F1: 2}
                        _ = v15
                    }
                    v16 := GenS3{F0: (v6 ^ v6), F1: -128}
                    _ = v16
                }
            }
            _ = fi10
            _ = v11
        }
        switch 128 {
        case -9223372036854775808:
            var v17 uint32 = 389
            _ = v17
        case 9223372036854775807:
            var v18 float64 = 0.0
            _ = v18
        case -170:
            v6 += (v6 ^ (v6 >> ((239) & 63)))
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := GenS4{F0: 11, F1: 0, F2: -42}
            _ = fi19
            _ = v20
        }
    }
    _ = v6
    return 360
}

func main() {
    var v1 int64 = -1
    var v2 int64 = (v1 << ((2) & 63))
    var v3 uint = 2
    switch "a" {
    case "hello":
        v21 := []int{128}
        _ = v21
    case "foo bar":
        for fi22 := 0; fi22 < 19; fi22++ {
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS4{F0: 1, F1: 1, F2: 97}
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    default:
        var v25 uint16 = 255
        _ = v25
    }
    var v26 float32 = 3.14159
    fmt.Println(v3)
    v1 = v1
    for fi27 := 0; fi27 < 15; fi27++ {
        if (!true) {
            v3 = (v3 * 0)
        } else {
            switch ("foo bar" + ("\n\t" + ((("\n\t" + ("" + "a")) + ("foo bar" + (("a" + ("foo bar" + "\n\t")) + "a"))) + ((("\n\t" + "\n\t") + (("a" + "\n\t") + "\n\t")) + ("foo bar" + ("a" + "foo bar")))))) {
            case "hello":
                v28 := GenS5{F0: 2147483648, F1: v1, F2: -1}
                _ = v28
            case "":
                fmt.Println(v2, v3)
            case "\n\t":
                fmt.Println(v1)
            default:
                v29 := []int{(fi27 + 0), 947, 262}
                _ = v29
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                fmt.Println(fi30)
                _ = fi30
            }
        }
        switch "" {
        case "\n\t":
            var v31 uint32 = 2147483647
            _ = v31
        case "":
            if false {
                v32 := []int{(fi27 | -1), -2147483648, genFn1()}
                if true {
                    for fi33 := 0; fi33 < 10; fi33++ {
                        var v34 bool = true
                        v35 := GenS1{F0: (v3 + (v3 ^ (v3 ^ (v3 ^ (v3 / (((v3 / (((v3 - v3)) | 1))) | 1)))))), F1: 32767}
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                } else {
                    switch "a" {
                    case "a":
                        var v36 float64 = ((3.14159 * 1.5) - (3.14159 + 0.1))
                        _ = v36
                    case "":
                        v26 = v26
                    default:
                        switch (fi27 ^ -649) {
                        case 42:
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                        }
                    }
                    _ = len(v32)
                    var v37 int8 = 100
                    _ = v37
                }
                _ = len(v32)
                _ = v32
            } else {
                v38 := GenS6{F0: 2147483647}
                v39 := []int{(fi27 >> ((779) & 63)), (fi27 * int(v2)), fi27, (fi27 << (((v3 >> (((v3 % ((v3) | 1))) & 63))) & 63))}
                _ = v38
                _ = v39
            }
        case "foo bar":
            v40 := GenS2{F0: 178}
            _ = v40
        default:
            fmt.Println(v26, v26, v3)
        }
        v41 := []int{(fi27 ^ genFn1())}
        _ = fi27
        _ = v41
    }
    v26 = v26
    v42 := GenS7{F0: 127, F1: 0}
    v43 := GenS4{F0: -128, F1: 2, F2: 0}
    switch "foo bar" {
    case "\n\t":
        for fi44 := 0; fi44 < 7; fi44++ {
            if ((v3 % (((v3 * v3)) | 1)) != (v3 | 42)) {
                if false {
                    fmt.Println(v1)
                    var v45 int32 = 0
                    var v46 bool = (false && false)
                    _ = v45
                    _ = v46
                }
                v47 := GenS8{F0: 127}
                _ = v47
            }
            v48 := GenS9{F0: true, F1: 222}
            switch genFn1() {
            case -2147483648:
                switch fi44 {
                case 4294967295:
                    v1 = -9223372036854775808
                case 128:
                    v2 = -1
                default:
                    var v49 float32 = (float32(1.5) * v26)
                    _ = v49
                }
            case 132:
                for fi50 := 0; fi50 < 4; fi50++ {
                    v51 := GenS10{F0: 0}
                    _ = fi50
                    _ = v51
                }
            case 32768:
                v52 := []int{}
                _ = v52
            default:
                v3 = v3
            }
            _ = fi44
            _ = v48
        }
    case "foo bar":
        for fi53 := 0; fi53 < 1; fi53++ {
            for fi54 := 0; fi54 < 19; fi54++ {
                fmt.Println(v3, v2)
                _ = fi54
            }
            for fi55 := 0; fi55 < 4; fi55++ {
                fmt.Println(v1, fi55)
                _ = fi55
            }
            for fi56 := 0; fi56 < 12; fi56++ {
                fmt.Println(v2)
                var v57 int8 = 0
                _ = fi56
                _ = v57
            }
            _ = fi53
        }
    }
    v58 := []int{1, 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v42
    _ = v43
    _ = v58
}
