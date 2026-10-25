package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int16; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int16 }
func genFn0() int {
    switch 357 {
    case 42:
        switch (("\n\t" + "foo bar") + (("\n\t" + "") + "a")) {
        case "":
            for fi4 := 0; fi4 < 3; fi4++ {
                var v5 bool = (-751 < 1)
                _ = fi4
                _ = v5
            }
        case "\n\t":
            var v6 int8 = -22
            _ = v6
        case "foo bar":
            var v7 int8 = 0
            _ = v7
        default:
            var v8 uint16 = 520
            _ = v8
        }
    case -517:
        v9 := []int{913, -9223372036854775808}
        _ = v9
    case 2147483647:
        v10 := []int{-2147483648, 986, 360, 128}
        _ = v10
    }
    var v11 int32 = -872
    fmt.Println(v11)
    var v12 uint8 = 128
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 int8 = -1
        v15 := GenS1{F0: 599, F1: (!false), F2: 0}
        if true {
            if false {
                var v16 float64 = (3.14159 - float64(-2.25))
                v14 -= (v14 + -128)
                fmt.Println(v12, fi13, v11)
                _ = v16
            } else {
                v17 := GenS1{F0: (^(fi13)), F1: true, F2: 0}
                v18 := GenS2{F0: false, F1: (fi13 - fi13)}
                _ = v17
                _ = v18
            }
            v12 -= (v12 ^ (v12 % ((v12) | 1)))
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v11
    _ = v12
    return 2147483647
}
func genFn1() uint32 {
    var v19 uint16 = 2
    var v20 string = "a"
    if true {
        switch 0 {
        case -301:
            var v21 int64 = -1
            _ = v21
        default:
            var v22 int = -508
            _ = v22
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 uint64 = 1
            if false {
                var v25 uint64 = 2
                for fi26 := 0; fi26 < 14; fi26++ {
                    switch (fi23 & 42) {
                    case 1:
                        v20 = ""
                    default:
                        if (18446744073709551615 <= 42) {
                            fmt.Println(fi23, v19)
                            v27 := []int{fi26, 32768, fi26}
                            v20 = ((v20 + v20) + ("a" + v20))
                            _ = v27
                        }
                    }
                    _ = fi26
                }
                if false {
                    v28 := GenS3{F0: -635, F1: 42, F2: 0}
                    if false {
                        v20 = (v20 + v20)
                    } else {
                        v25 += 739
                        v29 := GenS4{F0: v19, F1: -479}
                        v30 := GenS3{F0: -865, F1: -1, F2: 42}
                        _ = v29
                        _ = v30
                    }
                    _ = v28
                }
                _ = v25
            } else {
                fmt.Println(v24, fi23)
                if (true || (-569 < 1000)) {
                    if true {
                        fmt.Println(v20, fi23, v24)
                        var v31 uint16 = v19
                        _ = v31
                    } else {
                        if (866 < v19) {
                            fmt.Println(v24, v20, v24)
                        } else {
                            v32 := GenS1{F0: (fi23 | (fi23 | -9223372036854775808)), F1: true, F2: 254}
                            _ = v32
                        }
                    }
                }
                if false {
                    var v33 float32 = 0.1
                    var v34 uint = 65535
                    v35 := GenS4{F0: v19, F1: -1}
                    _ = v33
                    _ = v34
                    _ = v35
                } else {
                    v36 := []int{(fi23 + fi23), fi23, 1, (fi23 * (^(fi23)))}
                    var v37 uint16 = 628
                    _ = v36
                    _ = v37
                }
            }
            _ = fi23
            _ = v24
        }
        v38 := GenS2{F0: ("foo bar" != "\n\t"), F1: genFn0()}
        _ = v38
    }
    var v39 uint16 = v19
    _ = v19
    _ = v20
    _ = v39
    return 2147483648
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 32767
    var v3 float32 = 0.1
    v3 = (-2.25 - float32(((3.14159 + 0.1) + 0.1)))
    var v40 float64 = (1.5 * (0.1 / float64(((v3 * float32(0.1)) * v3))))
    var v41 int8 = 100
    v3 -= 1e10
    v42 := []int{32768, 143, -2147483648}
    v43 := GenS1{F0: -1, F1: true, F2: -419}
    v2 = (^(v2))
    var v44 int16 = 0
    v3 = v3
    var v45 uint16 = (v2 ^ (v2 ^ 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
