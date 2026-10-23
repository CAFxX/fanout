package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int16; F2 uint16 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool }
func genFn0(p0 int8, p1 uint16) string {
    var v6 int16 = 42
    var v7 int8 = p0
    var v8 uint64 = 42
    fmt.Println(v8, p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 string, p1 uint16, p2 int8) uint32 {
    if false {
        v9 := []int{127, 128}
        if true {
            var v10 uint32 = 42
            fmt.Println(v10)
            _ = v10
        }
        switch 127 {
        case -222:
            for fi11 := 0; fi11 < 2; fi11++ {
                p2 -= (p2 - p2)
                v12 := GenS1{F0: 342, F1: 0, F2: 42}
                _ = fi11
                _ = v12
            }
        }
        _ = v9
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint = 724
        _ = fi13
        _ = v14
    }
    if true {
        var v15 uint8 = 226
        _ = v15
    } else {
        var v16 float64 = 0.0
        v16 += (v16 * v16)
        v16 = (v16 / v16)
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = (120 > 0)
    v3 := make(map[string]bool)
    var v4 float32 = ((-2.25 * (3.14159 - float32(0.0))) - 1e10)
    var v5 int8 = 100
    v2 = v2
    v4 += (v4 * (v4 / v4))
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := v3["a"]
        switch 0 {
        case -158:
            var v19 uint64 = 2
            _ = v19
        case -2147483648:
            if false {
                if (!v18) {
                    v20 := []int{fi17, (fi17 + fi17), 32767, 2147483647}
                    _ = v20
                } else {
                    v21 := GenS2{F0: 186}
                    _ = v21
                }
                var v22 int64 = 512
                _ = v22
            } else {
                if ((v5 / ((52) | 1)) < (v5 ^ v5)) {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        var v24 int = fi23
                        _ = fi23
                        _ = v24
                    }
                    v4 = ((-2.25 / 1.5) - v4)
                }
                var v25 float32 = 0.0
                if false {
                    v18 = false
                    v25 += v25
                    v26 := GenS3{F0: -1, F1: 1, F2: 1}
                    _ = v26
                }
                _ = v25
            }
        }
        if true {
            if v18 {
                var v27 uint8 = 123
                _ = v27
            }
            v28 := GenS1{F0: 2, F1: 447, F2: 593}
            v4 += v4
            _ = v28
        }
        _ = fi17
        _ = v18
    }
    v2 = false
    delete(v3, "\n\t")
    v29 := GenS4{F0: -9223372036854775808, F1: -258}
    fmt.Println(v4, v5, v2)
    v2 = (!(!v2))
    var v30 uint16 = 2
    var v31 float32 = v4
    switch (("a" + "foo bar") + "") {
    case "\n\t":
        fmt.Println(v2, v31, v30)
    }
    if (v2 && true) {
        if (0 >= 1) {
            switch -1 {
            case 127:
                var v32 float64 = 0.1
                _ = v32
            case -2147483648:
                for fi33 := 0; fi33 < 13; fi33++ {
                    switch fi33 {
                    case 4294967295:
                        var v34 int32 = 42
                        _ = v34
                    case -1:
                        if v2 {
                            var v35 int = fi33
                            for fi36 := 0; fi36 < 20; fi36++ {
                                if (v5 >= v5) {
                                    if (float64((v31 / 1.0)) == 0.0) {
                                        v30 -= v30
                                        fmt.Println(v5, v35, v35)
                                        var v37 int8 = (v5 | v5)
                                        _ = v37
                                    }
                                }
                                _ = fi36
                            }
                            v4 = float32(1.5)
                            _ = v35
                        }
                    default:
                        var v38 string = ((("hello" + ("" + ("foo bar" + ((("hello" + "a") + "a") + "foo bar")))) + "") + "foo bar")
                        _ = v38
                    }
                    v3[("foo bar" + "")] = v2
                    _ = fi33
                }
            default:
                var v39 int = -718
                _ = v39
            }
        } else {
            v40 := []int{}
            v41 := GenS5{F0: 44}
            _ = v40
            _ = v41
        }
    } else {
        switch ("foo bar" + (("" + "a") + "")) {
        case "foo bar":
            v42 := GenS6{F0: (!false)}
            _ = v42
        default:
            v43 := []int{127, 128, -232}
            _ = v43
        }
        var v44 int = 32767
        var v45 uint64 = 18446744073709551615
        _ = v44
        _ = v45
    }
    var v46 int64 = int64(v30)
    v47 := GenS2{F0: 128}
    var v48 int16 = -909
    var v49 uint8 = 2
    fmt.Println(v2)
    fmt.Println(v5, v46, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
