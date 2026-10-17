package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 float32 = ((((3.14159 / 1.0) - 3.14159) / -2.25) / 1e10)
    v2 := make(map[string]uint32)
    v3 := []int{1, 1, 128, -2147483648}
    for fi4 := 0; fi4 < 19; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            fmt.Println(v1, fi4, v1)
            _ = fi5
        }
        if true {
            _ = len(v2)
            if (false && (true && (!(-965 < 1)))) {
                if (true && (255 == 65535)) {
                    var v6 int16 = -1
                    v7 := GenS1{F0: true, F1: 2147483648}
                    _ = v6
                    _ = v7
                }
                v2[((((("hello" + "a") + "foo bar") + "hello") + (("a" + "foo bar") + "a")) + "foo bar")] = 866
            }
            fmt.Println(fi4)
        } else {
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 uint64 = 2
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: ((!true) && (!false)), F1: 4294967295}
            v1 -= v1
            _ = v10
        }
        v1 += v1
        _ = fi4
    }
    v3 = append(v3, -188)
    fmt.Println(v1)
    switch -1 {
    case -615:
        v11 := v2[("\n\t" + (("hello" + "hello") + ""))]
        _ = v11
    case 32768:
        var v12 float64 = (-2.25 + 1e10)
        _ = v12
    case 42:
        v1 -= ((v1 + (v1 + (v1 * v1))) * v1)
    default:
        if false {
            switch ("" + (("a" + (("\n\t" + "\n\t") + "a")) + ((("" + ("a" + ("hello" + ("" + "")))) + ("hello" + ("" + "hello"))) + ("" + "foo bar")))) {
            case "hello":
                var v13 int16 = 32767
                _ = v13
            }
        }
    }
    v1 = ((float32(1.5) / v1) * 0.1)
    var v14 float32 = (v1 - 1.5)
    if true {
        v15 := GenS2{F0: 1}
        _ = v15
    }
    v16 := GenS1{F0: true, F1: 2147483647}
    var v17 uint = 4294967295
    v18 := GenS1{F0: true, F1: 389}
    if (0 <= v17) {
        for fi19 := 0; fi19 < 11; fi19++ {
            fmt.Println(fi19, v14)
            if (4294967295 == 2) {
                fmt.Println(v14, v1, fi19)
                v14 = v14
                var v20 int8 = -63
                _ = v20
            } else {
                delete(v2, ("hello" + "foo bar"))
            }
            var v21 string = ""
            _ = fi19
            _ = v21
        }
    } else {
        v17 = v17
        fmt.Println(v14, v1, v14)
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        switch fi22 {
        case 60:
            v1 = (v1 - v1)
        case 32767:
            switch fi22 {
            case 1:
                var v23 uint = (v17 % (((v17 ^ v17)) | 1))
                _ = v23
            }
        case 9223372036854775807:
            switch ("hello" + ("hello" + "\n\t")) {
            case "hello":
                var v24 int32 = -874
                _ = v24
            case "a":
                v17 = (v17 % ((v17) | 1))
            default:
                if true {
                    fmt.Println(v17)
                    for fi25 := 0; fi25 < 4; fi25++ {
                        v17 -= v17
                        fmt.Println(fi22)
                        switch fi25 {
                        case 132:
                            fmt.Println(v14, v1, v17)
                        case -2147483648:
                            var v26 uint32 = 1
                            _ = v26
                        default:
                            v27 := GenS3{F0: (v17 | v17)}
                            fmt.Println(v14)
                            var v28 int16 = -32768
                            _ = v27
                            _ = v28
                        }
                        v3 = append(v3, fi22)
                        _ = fi25
                    }
                } else {
                    var v29 uint64 = uint64(v17)
                    v14 = (v1 * 1.5)
                    _ = v29
                }
            }
        default:
            switch "hello" {
            case "\n\t":
                v30 := GenS1{F0: (219 <= 32767), F1: 65535}
                _ = v30
            case "":
                var v31 uint = 63
                _ = v31
            case "a":
                if true {
                    for fi32 := 0; fi32 < 18; fi32++ {
                        v1 = v1
                        var v33 uint = 0
                        _ = fi32
                        _ = v33
                    }
                }
            default:
                v17 += (^(v17))
            }
        }
        _ = fi22
    }
    v17 += 0
    var v34 uint16 = 42
    switch ("a" + "") {
    case "a":
        _ = len(v2)
    case "hello":
        v1 = (v14 - (v1 + 3.14159))
    default:
        for fi35 := 0; fi35 < 14; fi35++ {
            v34 -= (v34 | v34)
            var v36 float32 = v1
            v1 = (float32(0.1) / 1e10)
            _ = fi35
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
}
