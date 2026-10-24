package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16; F1 uint; F2 int }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int8; F1 uint64 }
func genFn0() int8 {
    v5 := GenS1{F0: 255}
    v6 := GenS1{F0: 127}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int8 = 42
        v8 += 127
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint8 = 0
    var v2 float64 = 1.5
    var v3 uint32 = 877
    var v4 string = ""
    var v10 bool = (-88 < genFn0())
    switch v4 {
    case "":
        v2 += (1.5 * v2)
    case "foo bar":
        v3 = v3
    default:
        v4 = v4
    }
    switch "foo bar" {
    case "foo bar":
        var v11 int32 = 872
        _ = v11
    case "a":
        v2 += 1e10
    }
    v12 := GenS2{F0: 32767, F1: 1}
    if (2 >= 555) {
        v13 := GenS1{F0: v1}
        if (127 < 0) {
            v10 = true
            var v14 string = (v4 + "")
            _ = v14
        }
        _ = v13
    } else {
        v10 = v10
        var v15 uint = 65535
        var v16 float64 = 1.5
        _ = v15
        _ = v16
    }
    v4 = v4
    if v10 {
        switch 128 {
        case 32768:
            var v17 string = v4
            _ = v17
        case 240:
            v4 = v4
        default:
            if true {
                for fi18 := 0; fi18 < 17; fi18++ {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        if (!(642 >= 0)) {
                            v1 = (v1 & 26)
                        }
                        v20 := GenS2{F0: 1, F1: 33}
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
                if (!false) {
                    if false {
                        v3 -= uint32(v1)
                        v21 := []int{42, -2147483648, -9223372036854775808, -2147483648}
                        _ = v21
                    } else {
                        var v22 uint64 = 337
                        if (127 >= -1) {
                            var v23 uint32 = 2147483647
                            var v24 uint16 = 255
                            fmt.Println(v10)
                            _ = v23
                            _ = v24
                        } else {
                            if (1 <= v22) {
                                v2 += (3.14159 * ((v2 * ((v2 * v2) + (v2 + (3.14159 + 3.14159)))) - v2))
                                v25 := GenS2{F0: 819, F1: 65535}
                                _ = v25
                            } else {
                                switch 4294967295 {
                                case 0:
                                    for fi26 := 0; fi26 < 6; fi26++ {
                                        v27 := GenS3{F0: -32768, F1: 642, F2: 32768}
                                        _ = fi26
                                        _ = v27
                                    }
                                default:
                                    v28 := []int{}
                                    _ = v28
                                }
                                v10 = ((!v10) && (255 != 264))
                            }
                        }
                        var v29 int16 = 32767
                        _ = v22
                        _ = v29
                    }
                    fmt.Println(v3)
                    v3 -= 4294967295
                }
                var v30 float32 = 0.1
                _ = v30
            }
        }
    } else {
        v31 := GenS4{F0: v10, F1: 9223372036854775807}
        fmt.Println(v1)
        _ = v31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        v33 := []int{fi32}
        _ = fi32
        _ = v33
    }
    if true {
        fmt.Println(v10)
        v34 := GenS2{F0: 755, F1: 475}
        _ = v34
    } else {
        var v35 bool = v10
        var v36 uint = 133
        _ = v35
        _ = v36
    }
    v4 = "foo bar"
    if (!v10) {
        if (!v10) {
            fmt.Println(v10, v2)
            switch (((v4 + (v4 + v4)) + ("" + ((v4 + "\n\t") + v4))) + "a") {
            case "a":
                switch "hello" {
                case "\n\t":
                    var v37 float32 = 0.0
                    _ = v37
                case "foo bar":
                    switch -2147483648 {
                    case -660:
                        for fi38 := 0; fi38 < 9; fi38++ {
                            var v39 int16 = 32767
                            _ = fi38
                            _ = v39
                        }
                    }
                }
            case "foo bar":
                var v40 uint8 = (^(v1))
                _ = v40
            case "hello":
                for fi41 := 0; fi41 < 1; fi41++ {
                    var v42 uint = 2
                    v43 := []int{(fi41 & fi41)}
                    var v44 float32 = -2.25
                    _ = fi41
                    _ = v42
                    _ = v43
                    _ = v44
                }
            }
        }
        v45 := GenS5{F0: -1, F1: 42}
        _ = v45
    }
    if v10 {
        fmt.Println(v1, v3)
        var v46 int8 = -105
        _ = v46
    } else {
        var v47 int64 = -1
        _ = v47
    }
    fmt.Println(v1, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
