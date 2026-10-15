package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint8) string {
    for fi4 := 0; fi4 < 10; fi4++ {
        p0 = p0
        var v5 float64 = 1e10
        _ = fi4
        _ = v5
    }
    if (((!true) || true) && ((4294967295 != 41) || false)) {
        for fi6 := 0; fi6 < 9; fi6++ {
            if (65535 != 983) {
                for fi7 := 0; fi7 < 6; fi7++ {
                    fmt.Println(fi7)
                    _ = fi7
                }
                p0 -= (p0 & 255)
                p0 = 42
            } else {
                var v8 int8 = 127
                var v9 uint16 = 1
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: 20, F1: -63}
            if true {
                for fi11 := 0; fi11 < 10; fi11++ {
                    if (3.14159 < -2.25) {
                        switch "\n\t" {
                        case "a":
                            p0 += 255
                        case "hello":
                            switch fi11 {
                            case 949:
                                v12 := []int{(fi6 ^ 0), (fi11 << ((608) & 63)), -216}
                                _ = v12
                            default:
                                if (((3.14159 <= -2.25) && ((!(float32(1.5) <= (-2.25 - 1.5))) || true)) || (false && true)) {
                                    v13 := []int{(fi6 - 9223372036854775807), (fi11 ^ (^(fi11))), fi11, (fi11 | -1)}
                                    _ = v13
                                }
                            }
                        default:
                            if false {
                                var v14 int16 = 1000
                                switch "\n\t" {
                                case "\n\t":
                                    v15 := GenS2{F0: 2}
                                    _ = v15
                                case "":
                                    v16 := []int{(fi11 ^ (fi11 | fi11)), 178}
                                    _ = v16
                                case "foo bar":
                                    var v17 string = ""
                                    _ = v17
                                default:
                                    var v18 int64 = -249
                                    _ = v18
                                }
                                _ = v14
                            } else {
                                var v19 uint64 = 146
                                v20 := GenS1{F0: 1, F1: -1}
                                _ = v19
                                _ = v20
                            }
                        }
                        p0 = 2
                    }
                    var v21 float32 = 3.14159
                    switch ((("foo bar" + "a") + "foo bar") + "foo bar") {
                    case "\n\t":
                        switch "foo bar" {
                        case "\n\t":
                            v21 -= (float32(((-2.25 + 1.5) * float64(0.0))) - float32(0.1))
                        }
                    case "hello":
                        fmt.Println(fi11)
                    default:
                        fmt.Println(fi6, p0)
                    }
                    switch -241 {
                    case -919:
                        var v22 int8 = 1
                        _ = v22
                    }
                    _ = fi11
                    _ = v21
                }
                var v23 uint = 313
                var v24 uint = v23
                _ = v23
                _ = v24
            }
            _ = fi6
            _ = v10
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 int8 = 30
            _ = fi25
            _ = v26
        }
    }
    p0 = p0
    v27 := []int{9223372036854775807}
    _ = p0
    _ = v27
    return "foo bar"
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]string)
    var v3 uint16 = v1
    fmt.Println(v3, v1)
    switch 32768 {
    case 805:
        var v28 uint8 = 1
        _ = v28
    case 2147483647:
        v1 -= 0
    }
    fmt.Println(v3)
    if false {
        var v29 int32 = 2147483647
        for fi30 := 0; fi30 < 2; fi30++ {
            v29 = v29
            _ = fi30
        }
        switch "" {
        case "foo bar":
            var v31 float64 = 0.0
            _ = v31
        case "a":
            v29 = v29
        }
        _ = v29
    } else {
        var v32 uint64 = 4294967295
        var v33 int64 = -1
        var v34 string = "a"
        _ = v32
        _ = v33
        _ = v34
    }
    switch 42 {
    case 805:
        if (128 > 128) {
            v35 := GenS1{F0: -124, F1: 100}
            var v36 int32 = -2147483648
            v37 := GenS1{F0: 127, F1: -38}
            _ = v35
            _ = v36
            _ = v37
        } else {
            _ = len(v2)
            if (false || ((42 >= 32767) || false)) {
                delete(v2, 6)
                for fi38 := 0; fi38 < 18; fi38++ {
                    v39 := []int{}
                    _ = fi38
                    _ = v39
                }
                v3 = (v3 ^ 42)
            } else {
                var v40 uint64 = 42
                v41 := GenS2{F0: v40}
                _ = v40
                _ = v41
            }
        }
    case 1:
        if false {
            delete(v2, -497)
        }
    case -1:
        v3 -= (v3 * (v1 ^ v1))
    }
    v42 := []int{0}
    delete(v2, 9223372036854775807)
    v43 := GenS1{F0: 127, F1: -62}
    if len(v42) > 0 {
        v42[0] = 27
    }
    if (42 <= 876) {
        v3 = v1
    } else {
        v1 = 65535
        v3 -= (v1 * v3)
        if true {
            var v44 uint64 = 0
            _ = v44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v42:", len(v42), v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
}
