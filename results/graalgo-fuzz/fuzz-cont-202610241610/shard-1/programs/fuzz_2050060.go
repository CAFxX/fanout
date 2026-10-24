package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int; F2 string }
type GenS3 struct { F0 bool; F1 int64; F2 int }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint64; F1 uint8 }
type GenS7 struct { F0 int16; F1 int32; F2 int8 }
func genFn0() int16 {
    switch "a" {
    case "":
        var v5 int = -1
        _ = v5
    case "hello":
        for fi6 := 0; fi6 < 7; fi6++ {
            switch ("a" + "") {
            case "foo bar":
                if (false && (2147483647 <= 2147483647)) {
                    if (fi6 > (fi6 * -359)) {
                        var v7 bool = (4294967295 < 4294967295)
                        _ = v7
                    }
                    var v8 int = fi6
                    v9 := GenS1{F0: 4294967295, F1: 4294967295, F2: 255}
                    _ = v8
                    _ = v9
                } else {
                    var v10 uint16 = 255
                    var v11 uint = 876
                    _ = v10
                    _ = v11
                }
            case "a":
                for fi12 := 0; fi12 < 14; fi12++ {
                    var v13 bool = true
                    _ = fi12
                    _ = v13
                }
            case "\n\t":
                if (!false) {
                    for fi14 := 0; fi14 < 20; fi14++ {
                        switch ("" + ("" + "a")) {
                        case "\n\t":
                            switch "\n\t" {
                            case "\n\t":
                                var v15 int = -1
                                _ = v15
                            default:
                                var v16 int32 = -1
                                _ = v16
                            }
                        default:
                            if true {
                                var v17 float64 = 0.1
                                _ = v17
                            }
                        }
                        var v18 float32 = ((0.0 - 3.14159) - 3.14159)
                        v18 = 1.5
                        _ = fi14
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 5; fi19++ {
                        if false {
                            if false {
                                var v20 float32 = -2.25
                                _ = v20
                            } else {
                                var v21 bool = true
                                v21 = v21
                                _ = v21
                            }
                            var v22 uint8 = 42
                            var v23 uint32 = 212
                            _ = v22
                            _ = v23
                        } else {
                            var v24 int = -561
                            v24 = (fi19 / (((^(fi19))) | 1))
                            _ = v24
                        }
                        var v25 uint64 = 18446744073709551615
                        v25 = (v25 & (v25 % ((42) | 1)))
                        _ = fi19
                        _ = v25
                    }
                    v26 := GenS1{F0: 2, F1: 2, F2: 226}
                    _ = v26
                } else {
                    fmt.Println(fi6)
                }
            }
            v27 := GenS1{F0: 137, F1: 2147483648, F2: 42}
            _ = fi6
            _ = v27
        }
    case "\n\t":
        var v28 int16 = -203
        _ = v28
    default:
        var v29 uint16 = 933
        _ = v29
    }
    switch "\n\t" {
    case "":
        var v30 int8 = -128
        _ = v30
    case "\n\t":
        var v31 int = -2147483648
        _ = v31
    case "foo bar":
        var v32 uint16 = 65535
        _ = v32
    }
    var v33 bool = (42 == 42)
    var v34 int = 0
    _ = v33
    _ = v34
    return 478
}

func main() {
    var v1 int = 32768
    var v2 int16 = 1000
    v3 := make(map[uint32]string)
    var v4 uint32 = 65535
    v4 = 199
    v2 = 0
    v35 := v3[v4]
    v36 := GenS2{F0: 973, F1: (v1 - 4294967295), F2: (v35 + "a")}
    var v37 float32 = (3.14159 / 3.14159)
    if (-28 >= 0) {
        v2 += v2
        v3[(v4 & v4)] = v35
    }
    if (!(-1 == 1)) {
        v38 := []int{int(v2), 127}
        v39 := GenS2{F0: 32767, F1: v1, F2: "foo bar"}
        if true {
            v40 := GenS3{F0: (-1 > 42), F1: 0, F2: -724}
            v37 = (0.1 + 1e10)
            fmt.Println(v37, v4)
            _ = v40
        } else {
            v41 := GenS4{F0: 42, F1: v35}
            v42 := GenS4{F0: 0, F1: "\n\t"}
            _ = v41
            _ = v42
        }
        _ = v38
        _ = v39
    }
    for fi43 := 0; fi43 < 4; fi43++ {
        var v44 int16 = genFn0()
        fmt.Println(v44, v2, fi43)
        _ = fi43
        _ = v44
    }
    var v45 int64 = 0
    if true {
        var v46 string = ("\n\t" + (((v35 + v35) + "a") + v35))
        _ = v46
    } else {
        v47 := GenS5{F0: v4, F1: 42, F2: 1}
        _ = v47
    }
    switch v35 {
    case "\n\t":
        v48 := GenS6{F0: 0, F1: 1}
        _ = v48
    case "":
        fmt.Println(v37)
    case "hello":
        v2 = genFn0()
    default:
        v49 := GenS7{F0: v2, F1: 807, F2: 100}
        _ = v49
    }
    v50 := GenS3{F0: (!(18446744073709551615 >= 2)), F1: v45, F2: -293}
    switch -637 {
    case -24:
        fmt.Println(v35)
    case 42:
        v37 -= 3.14159
    case 1:
        v51 := GenS7{F0: -338, F1: 0, F2: -128}
        _ = v51
    default:
        var v52 uint = 255
        _ = v52
    }
    v35 = v35
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v45
    _ = v50
}
