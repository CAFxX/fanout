package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int8; F1 int; F2 int }
func genFn0(p0 uint64) uint8 {
    v5 := []int{32768}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 128
}

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    var v3 uint64 = 0
    v4 := []int{}
    v1 = (361 >= 65535)
    switch -1 {
    case 32768:
        _ = len(v2)
    case 42:
        v3 = v3
    default:
        if v1 {
            if false {
                v3 = (v3 % ((v3) | 1))
            }
        }
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := v2[("\n\t" + ((("" + "\n\t") + "foo bar") + (("hello" + "") + "foo bar")))]
        _ = fi6
        _ = v7
    }
    var v8 int8 = 127
    v1 = false
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 string = "hello"
        v11 := GenS1{F0: 1, F1: 32767}
        var v12 int16 = 32767
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: genFn0(v3), F1: 768}
    for fi14 := 0; fi14 < 7; fi14++ {
        if (-9223372036854775808 <= -412) {
            if true {
                var v15 int8 = (^(v8))
                _ = v15
            }
            v16 := GenS3{F0: v8, F1: (fi14 + fi14), F2: (fi14 - (-(fi14)))}
            fmt.Println(fi14, v3, fi14)
            _ = v16
        } else {
            switch fi14 {
            case 32768:
                if (347 == 1) {
                    for fi17 := 0; fi17 < 7; fi17++ {
                        var v18 bool = v1
                        var v19 uint32 = 4294967295
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    v20 := v2["\n\t"]
                    _ = v20
                }
            case 9223372036854775807:
                for fi21 := 0; fi21 < 15; fi21++ {
                    v3 -= (v3 * v3)
                    _ = fi21
                }
            default:
                v22 := GenS2{F0: 93, F1: 671}
                _ = v22
            }
            var v23 uint16 = 42
            _ = v23
        }
        _ = fi14
    }
    v24 := GenS3{F0: -5, F1: -2147483648, F2: 2147483647}
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    var v25 int8 = (v8 >> ((42) & 63))
    fmt.Println(v25)
    _ = len(v2)
    v26 := GenS2{F0: 127, F1: 918}
    v8 = (^(v8))
    switch -906 {
    case 42:
        switch (("a" + ((("" + ("hello" + "a")) + "a") + "foo bar")) + "\n\t") {
        case "\n\t":
            switch 4294967295 {
            case 111:
                v8 = (^(v8))
            case 706:
                if true {
                    if (882 < 42) {
                        v27 := GenS1{F0: -402, F1: 42}
                        var v28 bool = false
                        v25 = -128
                        _ = v27
                        _ = v28
                    } else {
                        v8 = (v25 << ((uint(v3)) & 63))
                    }
                }
            case -2147483648:
                var v29 string = ""
                _ = v29
            }
        case "hello":
            var v30 uint8 = 255
            _ = v30
        case "a":
            v8 = v25
        default:
            v1 = false
        }
    case -1:
        for fi31 := 0; fi31 < 15; fi31++ {
            v25 += (v25 - -3)
            if ((-2.25 * float64(3.14159)) >= (0.1 * 0.1)) {
                for fi32 := 0; fi32 < 15; fi32++ {
                    var v33 float32 = 0.1
                    _ = fi32
                    _ = v33
                }
                v34 := GenS1{F0: (fi31 & -731), F1: 42}
                v1 = v1
                _ = v34
            } else {
                v25 -= v25
                v3 -= 42
                v8 += 1
            }
            _ = fi31
        }
    default:
        for fi35 := 0; fi35 < 11; fi35++ {
            if v1 {
                for fi36 := 0; fi36 < 19; fi36++ {
                    v37 := GenS3{F0: 100, F1: 0, F2: -2147483648}
                    _ = fi36
                    _ = v37
                }
            } else {
                v1 = ("" > "foo bar")
                var v38 string = (("hello" + (("a" + "") + "\n\t")) + "foo bar")
                var v39 uint64 = 2
                _ = v38
                _ = v39
            }
            var v40 uint64 = (v3 << ((42) & 63))
            v41 := GenS3{F0: (v25 / (((v25 | v25)) | 1)), F1: fi35, F2: fi35}
            _ = fi35
            _ = v40
            _ = v41
        }
    }
    var v42 uint16 = 886
    var v43 int32 = 591
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v8
}
