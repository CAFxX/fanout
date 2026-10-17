package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0(p0 uint, p1 uint8) string {
    v4 := []int{480}
    var v5 uint64 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "hello"
}

func main() {
    var v1 string = ""
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]bool)
    if (((1e10 * 1e10) - 3.14159) == 0.1) {
        v1 = (v1 + "a")
        fmt.Println(v2, v1)
    } else {
        fmt.Println(v1, v2)
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
        switch 412 {
        case -2147483648:
            v7 := []int{128, 32767, int(v2)}
            _ = v7
        case 160:
            var v8 int = 0
            _ = v8
        case -9223372036854775808:
            var v9 uint8 = 119
            _ = v9
        }
    }
    v10 := GenS1{F0: 42}
    if false {
        v3[1] = ((-2147483648 != 790) && true)
    }
    if (1 < 1) {
        v11 := GenS1{F0: 254}
        _ = v11
    } else {
        var v12 uint8 = 255
        _ = v12
    }
    v13 := GenS1{F0: 153}
    for fi14 := 0; fi14 < 10; fi14++ {
        if (2 < 2) {
            var v15 float32 = (0.1 + (0.0 + ((3.14159 / float32((1e10 - float64(0.1)))) * 3.14159)))
            _ = v15
        } else {
            v1 = "\n\t"
            v16 := GenS2{F0: 255}
            var v17 bool = false
            _ = v16
            _ = v17
        }
        var v18 float32 = float32(0.0)
        v19 := []int{-1, (fi14 >> ((255) & 63)), (fi14 | 431), fi14}
        _ = fi14
        _ = v18
        _ = v19
    }
    v20 := []int{128}
    var v21 string = (v1 + ((v1 + (v1 + "foo bar")) + (v1 + "\n\t")))
    if true {
        var v22 int16 = 423
        switch 42 {
        case 127:
            v23 := GenS2{F0: 1}
            _ = v23
        default:
            v22 += v22
        }
        _ = v22
    } else {
        if true {
            v3[1] = false
            v24 := GenS2{F0: 31}
            _ = v24
        }
        v25 := v3[-1]
        if (0 == 2) {
            v21 = (v1 + v21)
        }
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        if false {
            switch ("" + v21) {
            case "":
                fmt.Println(v21)
            }
            v27 := GenS1{F0: 255}
            _ = len(v3)
            _ = v27
        }
        v1 = v21
        for fi28 := 0; fi28 < 3; fi28++ {
            if true {
                v29 := GenS1{F0: 84}
                if true {
                    if (v2 <= 9223372036854775807) {
                        v1 = (v21 + (v21 + "\n\t"))
                        var v30 float64 = 1e10
                        if false {
                            if true {
                                var v31 uint16 = 0
                                var v32 bool = (127 == 42)
                                if v32 {
                                    var v33 int32 = -2147483648
                                    _ = v33
                                } else {
                                    v34 := v3[fi28]
                                    switch (v1 + "foo bar") {
                                    case "":
                                        v35 := GenS3{F0: true}
                                        _ = v35
                                    case "a":
                                        var v36 int8 = -93
                                        _ = v36
                                    }
                                    _ = v34
                                }
                                _ = v31
                                _ = v32
                            }
                            v37 := GenS4{F0: 255, F1: -2147483648}
                            _ = v37
                        } else {
                            delete(v3, (fi26 * -1))
                            var v38 float32 = (0.1 / (((1e10 - float32((float64(((3.14159 / 1.0) * 3.14159)) - (v30 / v30)))) * -2.25) * (-2.25 * (float32(-2.25) * 0.0))))
                            _ = v38
                        }
                        v39 := GenS4{F0: 891, F1: 826}
                        _ = v30
                        _ = v39
                    } else {
                        var v40 bool = false
                        var v41 uint = 545
                        _ = v40
                        _ = v41
                    }
                    if (!(-2.25 > 1.5)) {
                        var v42 float32 = 3.14159
                        var v43 int32 = 1
                        _ = v42
                        _ = v43
                    }
                    switch (fi26 - (fi26 / (((fi28 / ((fi26) | 1))) | 1))) {
                    case 32767:
                        if len(v20) > 0 {
                            v20[0] = (fi28 + fi26)
                        }
                    case 127:
                        var v44 string = v1
                        _ = v44
                    default:
                        switch (v21 + v1) {
                        case "foo bar":
                            if (("a" + (v1 + v21)) == v1) {
                                if len(v20) > 0 {
                                    v20[0] = fi26
                                }
                                v21 = v1
                            }
                        case "":
                            var v45 int16 = -673
                            _ = v45
                        }
                    }
                } else {
                    fmt.Println(v21, v1)
                    if len(v20) > 0 {
                        v20[0] = (^(fi26))
                    }
                    _ = len(v20)
                }
                var v46 uint8 = 127
                v47 := GenS3{F0: false}
                v48 := GenS3{F0: true}
                _ = v29
                _ = v46
                _ = v47
                _ = v48
            }
            _ = fi28
        }
        _ = fi26
    }
    fmt.Println(v1)
    delete(v3, 217)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
