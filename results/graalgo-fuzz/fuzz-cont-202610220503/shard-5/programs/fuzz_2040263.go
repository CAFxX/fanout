package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 string }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint32, p1 uint64) int8 {
    v6 := GenS1{F0: 128, F1: "foo bar", F2: 42}
    var v7 float32 = 0.1
    for fi8 := 0; fi8 < 8; fi8++ {
        if (0 > -1) {
            p0 -= (p0 ^ p0)
        } else {
            var v9 uint32 = p0
            switch "\n\t" {
            case "\n\t":
                var v10 int8 = -1
                _ = v10
            case "":
                p1 = uint64(p0)
            }
            for fi11 := 0; fi11 < 1; fi11++ {
                switch (fi11 / ((127) | 1)) {
                case -425:
                    v9 -= (p0 & 35)
                }
                _ = fi11
            }
            _ = v9
        }
        if true {
            var v12 bool = ((-185 < -6) && true)
            _ = v12
        } else {
            p1 += (^(p1))
            fmt.Println(p1, fi8, v7)
        }
        p0 = p0
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint32) int8 {
    for fi13 := 0; fi13 < 7; fi13++ {
        p0 += (^(p0))
        _ = fi13
    }
    v14 := []int{1}
    _ = p0
    _ = v14
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{"a", "a", "a"}
    v3 := make(map[int]int)
    var v4 int8 = 2
    v5 := make(map[string]bool)
    if (!true) {
        fmt.Println(v4)
    } else {
        if (!true) {
            var v15 int8 = (v4 >> ((18446744073709551615) & 63))
            fmt.Println(v15)
            _ = v15
        } else {
            fmt.Println(v4)
        }
        v2 = append(v2, "\n\t")
    }
    var v16 int32 = 0
    var v17 float64 = (((0.1 / 1.5) / (3.14159 + (3.14159 * -2.25))) * 0.0)
    if false {
        if (!true) {
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 uint64 = (^(v1))
                _ = fi18
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 18; fi20++ {
                fmt.Println(v4, v1)
                v2 = append(v2, ((("hello" + "") + "foo bar") + ("hello" + "a")))
                _ = fi20
            }
            var v21 uint = 4294967295
            _ = v21
        }
        fmt.Println(v4)
    } else {
        if (v17 <= (v17 + ((1e10 / (3.14159 + v17)) + (float64((3.14159 + -2.25)) / (v17 - (-2.25 + ((v17 / ((3.14159 - ((v17 * (v17 * v17)) / ((v17 * 1.5) - (v17 - 0.1)))) - (v17 - -2.25))) + 0.0))))))) {
            fmt.Println(v16, v17)
            var v22 float32 = 1.5
            _ = v22
        } else {
            v17 += (v17 / v17)
            fmt.Println(v17)
        }
        var v23 float32 = 1e10
        v24 := GenS2{F0: (9223372036854775807 > 821), F1: 255, F2: "hello"}
        _ = v23
        _ = v24
    }
    var v25 int8 = v4
    v5[("a" + "")] = true
    if false {
        v16 = (v16 | (^(v16)))
    } else {
        if false {
            for fi26 := 0; fi26 < 17; fi26++ {
                switch "" {
                case "":
                    v1 = 784
                case "a":
                    _ = len(v2)
                case "foo bar":
                    var v27 float64 = v17
                    _ = v27
                default:
                    v1 -= (v1 << ((354) & 63))
                }
                _ = fi26
            }
            fmt.Println(v4)
        }
        for fi28 := 0; fi28 < 4; fi28++ {
            switch fi28 {
            case -122:
                switch ("" + (("foo bar" + (("" + "foo bar") + ("a" + ("a" + "\n\t")))) + "\n\t")) {
                case "\n\t":
                    var v29 uint8 = 127
                    _ = v29
                case "foo bar":
                    v16 -= (v16 - v16)
                case "a":
                    v30 := GenS3{F0: v1, F1: (v4 | -96)}
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 14; fi31++ {
                        if false {
                            _ = len(v5)
                            v16 += 419
                            v32 := GenS3{F0: v1, F1: (v25 * (v25 * v25))}
                            _ = v32
                        } else {
                            v17 -= (1e10 * v17)
                            v16 += (v16 * (v16 & v16))
                        }
                        _ = fi31
                    }
                }
            case -303:
                if (217 > (fi28 % ((fi28) | 1))) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (4294967295 > 42) {
                        v25 -= genFn1(4294967295)
                    }
                }
            case 32768:
                v3[1] = fi28
            default:
                delete(v3, 127)
            }
            v4 -= genFn0(42, v1)
            var v33 int = (-(fi28))
            _ = fi28
            _ = v33
        }
        for fi34 := 0; fi34 < 9; fi34++ {
            v4 = (v25 | v4)
            _ = fi34
        }
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        var v36 bool = (59 <= uint8(v1))
        switch "a" {
        case "foo bar":
            v5["a"] = ((true || true) && (124 > 2))
        case "hello":
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
        case "a":
            fmt.Println(v4, v25, v16)
        default:
            v36 = (!v36)
        }
        var v37 uint64 = v1
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 string = ""
    var v39 uint32 = 4294967295
    v40 := GenS2{F0: ((("a" + ("" + v38)) != "hello") && false), F1: 981, F2: ""}
    fmt.Println(v17, v39)
    if false {
        for fi41 := 0; fi41 < 14; fi41++ {
            var v42 float32 = -2.25
            fmt.Println(v4, v42, v17)
            _ = fi41
            _ = v42
        }
        if len(v2) > 0 {
            v2[0] = v38
        }
        v1 = v1
    } else {
        v43 := GenS3{F0: 0, F1: genFn0(2147483647, (v1 % ((v1) | 1)))}
        var v44 bool = false
        fmt.Println(v1, v16)
        _ = v43
        _ = v44
    }
    switch (v38 + "foo bar") {
    case "hello":
        v45 := GenS3{F0: v1, F1: 127}
        _ = v45
    case "\n\t":
        v46 := GenS2{F0: (-6 != -9223372036854775808), F1: 255, F2: (v38 + v38)}
        _ = v46
    case "":
        v25 = 42
    }
    switch int(v39) {
    case -1:
        v47 := GenS2{F0: (!(v1 != (v1 * v1))), F1: 600, F2: v38}
        _ = v47
    case 127:
        var v48 uint = 123
        _ = v48
    default:
        fmt.Println(v16, v17)
    }
    var v49 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v49
    _ = v5
}
