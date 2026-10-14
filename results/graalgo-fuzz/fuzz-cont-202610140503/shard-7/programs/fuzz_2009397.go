package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int16 }
func genFn0() uint {
    v7 := GenS1{F0: 1, F1: 0, F2: 0}
    var v8 int32 = 1
    var v9 int64 = 42
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int = (fi10 % (((-(fi10))) | 1))
        fmt.Println(v11)
        v8 = (^(v8))
        _ = fi10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() int64 {
    var v12 int8 = -128
    var v13 int8 = (v12 << ((genFn0()) & 63))
    v14 := []int{2147483647, 546}
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int = int(v13)
        _ = fi15
        _ = v16
    }
    var v17 uint16 = 255
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    return 672
}

func main() {
    var v1 uint = 2
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    var v5 uint8 = 1
    var v6 uint64 = 4294967295
    switch ("hello" + "hello") {
    case "":
        if true {
            if (65535 != 32767) {
                fmt.Println(v6, v5)
                v18 := v4[""]
                v18 = v18
                _ = v18
            }
            for fi19 := 0; fi19 < 6; fi19++ {
                if true {
                    var v20 uint16 = 32767
                    _ = v20
                }
                switch (fi19 >> ((genFn0()) & 63)) {
                case 32768:
                    var v21 uint64 = 1
                    _ = v21
                case 0:
                    delete(v4, "")
                case 9223372036854775807:
                    switch fi19 {
                    case 32768:
                        for fi22 := 0; fi22 < 9; fi22++ {
                            var v23 float64 = (((3.14159 - float64((float32(1.5) / 3.14159))) + (((1e10 + -2.25) * 1.5) + (3.14159 * (1.5 / (((0.0 + 0.0) - 1.5) + (-2.25 * 1e10)))))) - float64(1.5))
                            switch (fi19 & (fi19 << ((v1) & 63))) {
                            case 128:
                                v24 := GenS1{F0: -2147483648, F1: v6, F2: (^(v5))}
                                _ = v24
                            case 127:
                                v25 := v2[("hello" + (("foo bar" + "foo bar") + "\n\t"))]
                                _ = v25
                            case 0:
                                fmt.Println(v6, v6)
                            }
                            _ = fi22
                            _ = v23
                        }
                    case 898:
                        for fi26 := 0; fi26 < 17; fi26++ {
                            fmt.Println(v6, v1)
                            v6 = (v6 | uint64(v1))
                            _ = fi26
                        }
                    case 4294967295:
                        for fi27 := 0; fi27 < 20; fi27++ {
                            switch "\n\t" {
                            case "foo bar":
                                v28 := GenS1{F0: (fi19 % (((-(fi19))) | 1)), F1: (v6 % ((2) | 1)), F2: (v5 & v5)}
                                _ = v28
                            case "a":
                                v1 += (^(v1))
                            case "":
                                var v29 int64 = 9223372036854775807
                                _ = v29
                            default:
                                var v30 uint16 = 2
                                _ = v30
                            }
                            var v31 float32 = ((1.5 - 0.1) + 3.14159)
                            v6 = v6
                            _ = fi27
                            _ = v31
                        }
                    default:
                        v32 := GenS2{F0: "\n\t"}
                        _ = v32
                    }
                default:
                    v33 := GenS1{F0: (fi19 + fi19), F1: (^(v6)), F2: (^(v5))}
                    _ = v33
                }
                var v34 uint16 = 255
                _ = fi19
                _ = v34
            }
            var v35 float32 = 1.5
            _ = v35
        }
    case "hello":
        if (v5 != v5) {
            v5 = (^(v5))
        } else {
            if true {
                v1 = (v1 - genFn0())
                v36 := GenS3{F0: -1}
                _ = v36
            } else {
                v6 += 881
                switch "foo bar" {
                case "\n\t":
                    v37 := GenS3{F0: 787}
                    _ = v37
                default:
                    v1 -= v1
                }
                v6 = v6
            }
        }
    case "foo bar":
        if false {
            v5 += 29
            v6 = 63
        }
    default:
        var v38 string = "a"
        _ = v38
    }
    v1 += v1
    for fi39 := 0; fi39 < 5; fi39++ {
        switch "foo bar" {
        case "a":
            var v40 int32 = -2147483648
            _ = v40
        default:
            for fi41 := 0; fi41 < 16; fi41++ {
                v5 -= v5
                v42 := []int{(fi39 & fi39), fi39}
                if (1 == 894) {
                    v1 += (^(v1))
                    v1 = (^(v1))
                    var v43 int16 = 42
                    _ = v43
                }
                _ = fi41
                _ = v42
            }
        }
        fmt.Println(fi39)
        var v44 bool = (-1 >= -1)
        _ = fi39
        _ = v44
    }
    v1 += 18446744073709551615
    var v45 uint = 42
    v5 += v5
    fmt.Println(v45)
    v46 := GenS2{F0: ("hello" + ("a" + ""))}
    v47 := GenS4{F0: 4294967295, F1: v6, F2: 78}
    fmt.Println(v6)
    v1 = v1
    var v48 string = "a"
    fmt.Println(v1, v5, v1)
    var v49 int64 = 0
    fmt.Println(v45, v49)
    v49 = v49
    v50 := GenS4{F0: 385, F1: v6, F2: -32768}
    v1 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
