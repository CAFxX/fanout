package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int32; F2 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 int64; F1 string; F2 uint16 }
type GenS9 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint, p1 uint) uint {
    v7 := GenS1{F0: 288, F1: 73, F2: 151}
    var v8 bool = (164 != 503)
    for fi9 := 0; fi9 < 12; fi9++ {
        v8 = (v8 && (!(0.1 <= -2.25)))
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := []int{(^(fi10)), fi10, (-(fi10)), fi9}
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            var v12 int32 = 2147483647
            if ((fi10 < (fi9 + (-(fi9)))) && v8) {
                switch ("a" + "hello") {
                case "foo bar":
                    switch (fi9 + -357) {
                    case -9223372036854775808:
                        var v13 int = fi9
                        _ = v13
                    }
                case "a":
                    v8 = (!v8)
                }
                switch "hello" {
                case "a":
                    if false {
                        if v8 {
                            p1 = 18446744073709551615
                        }
                    }
                case "hello":
                    var v14 uint = (^(p0))
                    var v15 float32 = ((3.14159 / 1e10) / ((float32(((0.1 - (1e10 * 1e10)) - 1.5)) / (float32(1.5) - (1e10 + 0.0))) + ((float32((0.0 - 1e10)) + (float32(((1.5 + (1e10 + 1.5)) - 3.14159)) + -2.25)) / (1.5 / 3.14159))))
                    _ = v14
                    _ = v15
                case "\n\t":
                    var v16 uint64 = 1
                    _ = v16
                default:
                    var v17 bool = v8
                    _ = v17
                }
                var v18 float32 = 0.0
                _ = v18
            } else {
                var v19 int = fi9
                if true {
                    p1 += p1
                } else {
                    v11 = append(v11, -2147483648)
                    if len(v11) > 0 {
                        v11[0] = 42
                    }
                    fmt.Println(fi10)
                    var v20 uint = 938
                    _ = v20
                }
                var v21 float64 = 1e10
                _ = v19
                _ = v21
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(v8, v8)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (p0 & p1)
}
func genFn1() int8 {
    if (2 == 44) {
        switch "a" {
        case "":
            var v22 float64 = 3.14159
            _ = v22
        case "foo bar":
            var v23 uint8 = 127
            _ = v23
        case "\n\t":
            v24 := GenS2{F0: 0, F1: 35}
            _ = v24
        default:
            var v25 uint64 = 2
            _ = v25
        }
        var v26 int16 = 1000
        _ = v26
    } else {
        if (1 == 0) {
            v27 := []int{-51, 2147483647, 127}
            var v28 string = "hello"
            _ = v27
            _ = v28
        } else {
            var v29 float64 = 0.1
            for fi30 := 0; fi30 < 7; fi30++ {
                v31 := []int{736, fi30}
                if len(v31) >= 2 {
                    v31 = v31[1:len(v31)]
                }
                var v32 int8 = -128
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v29
        }
        for fi33 := 0; fi33 < 8; fi33++ {
            var v34 uint32 = 65535
            v35 := []int{fi33, -9223372036854775808, -686, (fi33 & (fi33 | fi33))}
            _ = fi33
            _ = v34
            _ = v35
        }
        var v36 int = 2147483647
        _ = v36
    }
    v37 := []int{}
    for fi38 := 0; fi38 < 3; fi38++ {
        v39 := GenS3{F0: 1}
        v40 := GenS4{F0: 123, F1: 2147483647, F2: fi38}
        _ = fi38
        _ = v39
        _ = v40
    }
    switch 185 {
    case 1:
        if (true && (!false)) {
            v41 := GenS5{F0: true}
            _ = v41
        } else {
            var v42 float64 = 3.14159
            if len(v37) >= 2 {
                v37 = v37[1:len(v37)]
            }
            _ = len(v37)
            _ = v42
        }
    default:
        for fi43 := 0; fi43 < 1; fi43++ {
            var v44 int64 = 42
            var v45 bool = (!(0.0 > (0.1 / (3.14159 - ((0.1 + 0.0) * 1e10)))))
            v44 -= (v44 + (v44 & 9223372036854775807))
            _ = fi43
            _ = v44
            _ = v45
        }
    }
    _ = v37
    return 127
}

func main() {
    var v1 int8 = 75
    var v2 float32 = (0.0 - 0.1)
    var v3 float32 = v2
    var v4 uint8 = 1
    v5 := make(map[string]string)
    var v6 string = "a"
    var v46 uint64 = 18446744073709551615
    fmt.Println(v3)
    if (((1000 == -32768) || true) || false) {
        var v47 int = 2147483647
        var v48 uint32 = 815
        v49 := GenS6{F0: true, F1: 537}
        _ = v47
        _ = v48
        _ = v49
    } else {
        for fi50 := 0; fi50 < 5; fi50++ {
            v51 := GenS7{F0: 0, F1: 18446744073709551615}
            fmt.Println(v2, v1, fi50)
            var v52 string = (("a" + (v6 + ("\n\t" + ""))) + (("" + (("a" + v6) + v6)) + "foo bar"))
            _ = fi50
            _ = v51
            _ = v52
        }
        v53 := GenS1{F0: v46, F1: 1, F2: v46}
        _ = v53
    }
    var v54 uint64 = (v46 << ((0) & 63))
    if false {
        v55 := GenS1{F0: v54, F1: genFn1(), F2: v54}
        v5["\n\t"] = "a"
        _ = v55
    }
    var v56 string = "foo bar"
    fmt.Println(v54, v46)
    var v57 float64 = ((((0.0 / -2.25) + (3.14159 * 0.1)) + float64(v3)) + 1e10)
    v57 = 1e10
    switch v56 {
    case "\n\t":
        v58 := v5[v56]
        _ = v58
    case "foo bar":
        fmt.Println(v56, v54, v54)
    default:
        v59 := GenS8{F0: 42, F1: v6, F2: 32767}
        _ = v59
    }
    var v60 uint16 = 362
    v61 := GenS9{F0: (v60 % ((v60) | 1)), F1: 42}
    v62 := []int{591, 510, 715, -593}
    if false {
        if len(v62) > 0 {
            v62[0] = 4294967295
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v54
    _ = v56
    _ = v57
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
