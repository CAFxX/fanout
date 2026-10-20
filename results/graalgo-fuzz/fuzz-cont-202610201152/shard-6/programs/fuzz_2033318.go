package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint8) string {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 uint = 4294967295
        v6 := []int{int(p0), fi4}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int16 = -1
    _ = p0
    _ = v7
    return ((("" + "hello") + ("hello" + "")) + "")
}

func main() {
    var v1 float32 = 1.5
    var v2 float64 = 0.0
    var v3 uint = 2
    if true {
        switch "hello" {
        case "a":
            v2 -= v2
        case "":
            v3 -= 65535
        }
        var v8 uint32 = 358
        fmt.Println(v3, v1)
        _ = v8
    } else {
        switch ("a" + (("" + "\n\t") + "a")) {
        case "\n\t":
            for fi9 := 0; fi9 < 18; fi9++ {
                v2 += 3.14159
                if (42 == -9223372036854775808) {
                    if (!true) {
                        v3 = (^(v3))
                    }
                    v1 -= v1
                    var v10 uint8 = 2
                    _ = v10
                }
                var v11 string = "foo bar"
                _ = fi9
                _ = v11
            }
        case "a":
            if ((!("hello" > "hello")) || true) {
                v1 -= v1
                if (true && (127 < 1)) {
                    v1 -= -2.25
                } else {
                    switch -626 {
                    case 32768:
                        v12 := GenS1{F0: 127}
                        _ = v12
                    case -92:
                        v13 := []int{0}
                        _ = v13
                    case -506:
                        var v14 bool = ((v1 - v1) == (((0.0 - (v1 * 1e10)) / v1) / v1))
                        _ = v14
                    }
                    var v15 int16 = 1
                    var v16 string = "foo bar"
                    _ = v15
                    _ = v16
                }
                v17 := GenS2{F0: 127, F1: 88}
                _ = v17
            }
        case "":
            if true {
                v18 := GenS3{F0: 4294967295, F1: 0}
                var v19 uint8 = 1
                var v20 uint16 = 508
                _ = v18
                _ = v19
                _ = v20
            }
        default:
            for fi21 := 0; fi21 < 17; fi21++ {
                v22 := []int{(fi21 / (((fi21 + fi21)) | 1)), (^(fi21)), fi21, (fi21 / (((fi21 % ((fi21) | 1))) | 1))}
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 int16 = 42
                    _ = fi23
                    _ = v24
                }
                var v25 uint32 = 760
                _ = fi21
                _ = v22
                _ = v25
            }
        }
        v26 := GenS4{F0: 1}
        v27 := GenS4{F0: -9223372036854775808}
        _ = v26
        _ = v27
    }
    fmt.Println(v2)
    var v28 int16 = 1
    var v29 uint16 = 255
    switch 32768 {
    case 32767:
        fmt.Println(v1)
    case -9223372036854775808:
        var v30 uint64 = 2
        _ = v30
    case 856:
        var v31 uint64 = 2
        _ = v31
    default:
        if true {
            v3 = v3
            fmt.Println(v29, v3)
        } else {
            var v32 int = 9223372036854775807
            v33 := GenS3{F0: (v32 ^ int(v29)), F1: 685}
            _ = v32
            _ = v33
        }
    }
    v29 -= v29
    fmt.Println(v29, v1)
    if ((("" + "\n\t") == "foo bar") || (false || (v28 != (v28 + (v28 % ((-32768) | 1)))))) {
        v34 := GenS4{F0: 0}
        _ = v34
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
