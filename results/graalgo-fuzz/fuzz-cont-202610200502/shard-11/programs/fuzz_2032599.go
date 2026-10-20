package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 string; F2 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32 }
func genFn0(p0 int, p1 uint64, p2 uint) uint16 {
    v4 := GenS1{F0: 65535, F1: 0}
    for fi5 := 0; fi5 < 8; fi5++ {
        if (131 >= 42) {
            v6 := []int{1, 0, (fi5 % ((-141) | 1)), int(p2)}
            var v7 int64 = 1
            _ = v6
            _ = v7
        } else {
            if true {
                v8 := GenS1{F0: 602, F1: 32767}
                _ = v8
            } else {
                v9 := []int{p0, 2147483647}
                _ = v9
            }
            if ("foo bar" >= ("\n\t" + "a")) {
                var v10 float64 = 0.1
                _ = v10
            } else {
                p0 += 937
                fmt.Println(fi5)
                v11 := []int{(^(fi5)), (p0 & int(p2))}
                _ = v11
            }
            p2 = p2
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            fmt.Println(p1)
            p2 = 42
            fmt.Println(p0)
            _ = fi12
        }
        switch fi5 {
        case 546:
            v13 := []int{(fi5 * (p0 << ((p2) & 63))), p0, (p0 & (fi5 ^ fi5))}
            _ = v13
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1() int {
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int64 = 1
        if true {
            var v16 uint64 = 578
            v17 := GenS2{F0: "foo bar"}
            fmt.Println(v16, fi14, v15)
            _ = v16
            _ = v17
        } else {
            var v18 int64 = (v15 / (((v15 + v15)) | 1))
            var v19 int64 = -9223372036854775808
            _ = v18
            _ = v19
        }
        var v20 bool = false
        _ = fi14
        _ = v15
        _ = v20
    }
    var v21 string = "\n\t"
    _ = v21
    return -1
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 255
    v3 := make(map[string]int)
    if false {
        var v22 int16 = 1000
        _ = v22
    }
    v2 = (^(v2))
    for fi23 := 0; fi23 < 16; fi23++ {
        if ("" > "foo bar") {
            var v24 int32 = -2147483648
            var v25 int8 = 79
            switch "" {
            case "foo bar":
                if (true || (true || (!(false || false)))) {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        if ((32767 >= 42) && ((!(!false)) && (2 < 255))) {
                            var v27 uint64 = 0
                            var v28 int64 = 942
                            v29 := GenS3{F0: genFn1()}
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            v30 := GenS4{F0: -9223372036854775808, F1: "", F2: 242}
                            _ = v30
                        }
                        _ = fi26
                    }
                }
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "a":
                v31 := GenS4{F0: -9223372036854775808, F1: "foo bar", F2: uint(v2)}
                _ = v31
            }
            _ = v24
            _ = v25
        } else {
            var v32 uint = 114
            if ((v2 << (((v32 | (^(v32)))) & 63)) < uint8(v32)) {
                v2 = (^(v2))
            }
            _ = v32
        }
        if (!(fi23 >= fi23)) {
            v33 := GenS5{F0: ("foo bar" + "hello")}
            _ = v33
        } else {
            var v34 uint16 = 65535
            switch ("hello" + "foo bar") {
            case "a":
                v35 := GenS6{F0: -2147483648}
                _ = v35
            case "hello":
                v2 = v2
            case "\n\t":
                var v36 int = fi23
                _ = v36
            default:
                var v37 string = "hello"
                _ = v37
            }
            var v38 float32 = (0.0 + 1.5)
            _ = v34
            _ = v38
        }
        _ = fi23
    }
    v39 := GenS5{F0: "hello"}
    v2 = v2
    v2 = 2
    v2 = v2
    v2 += (v2 ^ v2)
    v2 = v2
    if true {
        for fi40 := 0; fi40 < 20; fi40++ {
            if true {
                var v41 float32 = ((0.1 / 0.1) - ((0.1 / -2.25) * 0.0))
                _ = v41
            }
            var v42 int = (-(fi40))
            _ = fi40
            _ = v42
        }
        switch ("hello" + "hello") {
        case "foo bar":
            v43 := GenS1{F0: 0, F1: 1000}
            _ = v43
        }
        v2 += (v2 * v2)
    }
    for fi44 := 0; fi44 < 10; fi44++ {
        var v45 uint = 2
        fmt.Println(fi44, v2)
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 11; fi46++ {
        fmt.Println(fi46, fi46)
        if len(v1) > 0 {
            v1[0] = 2
        }
        _ = fi46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
}
