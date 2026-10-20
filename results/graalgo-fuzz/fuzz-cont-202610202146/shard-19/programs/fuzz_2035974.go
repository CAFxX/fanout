package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint; F1 int8; F2 int8 }
func genFn0(p0 int64) int64 {
    switch 32768 {
    case 1:
        if true {
            var v6 bool = (-87 != 1)
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := []int{892, 1, (fi7 << ((65535) & 63))}
                switch (("a" + "") + "foo bar") {
                case "a":
                    if false {
                        var v9 uint64 = 42
                        v9 = (v9 << ((0) & 63))
                        var v10 float32 = 0.0
                        _ = v10
                        _ = v9
                    } else {
                        switch ("\n\t" + ("hello" + ("a" + "hello"))) {
                        case "":
                            var v11 uint = 4294967295
                            _ = v11
                        case "a":
                            v12 := GenS1{F0: 255, F1: 127, F2: ("foo bar" + "foo bar")}
                            _ = v12
                        }
                        var v13 int32 = -480
                        var v14 string = "hello"
                        _ = v13
                        _ = v14
                    }
                case "\n\t":
                    if len(v8) > 0 {
                        v8[0] = fi7
                    }
                default:
                    var v15 int16 = -791
                    _ = v15
                }
                _ = fi7
                _ = v8
            }
            _ = v6
        }
    case -625:
        var v16 uint64 = 42
        _ = v16
    case -31:
        if (3.14159 <= ((1e10 + 0.1) / 1e10)) {
            v17 := []int{2147483647, 97}
            if len(v17) > 0 {
                v17[0] = 32768
            }
            v18 := GenS1{F0: 42, F1: 136, F2: (("\n\t" + "") + "a")}
            _ = v17
            _ = v18
        } else {
            v19 := GenS2{F0: 790}
            switch 170 {
            case 1:
                if false {
                    var v20 int32 = 42
                    _ = v20
                }
            case -1:
                var v21 float64 = -2.25
                _ = v21
            case -361:
                switch "" {
                case "\n\t":
                    var v22 uint64 = 2
                    _ = v22
                case "":
                    v23 := GenS1{F0: 1, F1: 128, F2: ""}
                    _ = v23
                case "a":
                    v24 := []int{2147483647}
                    _ = v24
                }
            }
            _ = v19
        }
    }
    fmt.Println(p0)
    _ = p0
    return 1
}
func genFn1(p0 bool, p1 int8) uint {
    var v25 uint8 = 2
    for fi26 := 0; fi26 < 2; fi26++ {
        p0 = p0
        _ = fi26
    }
    v27 := GenS1{F0: 208, F1: v25, F2: ""}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v25
    _ = v27
    return 1
}

func main() {
    var v1 string = "a"
    var v2 uint = 255
    v3 := []uint32{2, 2147483648, 342, 2147483648}
    var v4 float64 = 0.1
    var v5 string = v1
    for fi28 := 0; fi28 < 13; fi28++ {
        v4 -= v4
        v3 = append(v3, 2)
        var v29 uint = (v2 * 647)
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        v3 = append(v3, 42)
        _ = len(v3)
        _ = fi30
    }
    v1 = "a"
    if true {
        var v31 uint8 = 127
        _ = v31
    }
    v32 := GenS2{F0: (v2 ^ 255)}
    var v33 uint64 = 18446744073709551615
    if (genFn1((!(-1 == 1)), -128) != v2) {
        var v34 uint = v2
        if false {
            switch -87 {
            case -485:
                var v35 int16 = -1
                _ = v35
            case 901:
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
            }
        } else {
            if (!(v33 > v33)) {
                for fi36 := 0; fi36 < 15; fi36++ {
                    var v37 float32 = -2.25
                    _ = fi36
                    _ = v37
                }
                switch 32767 {
                case 2147483647:
                    fmt.Println(v34, v4, v33)
                case 4294967295:
                    fmt.Println(v2, v5, v33)
                }
                switch v1 {
                case "hello":
                    for fi38 := 0; fi38 < 10; fi38++ {
                        v39 := GenS1{F0: 127, F1: 2, F2: v1}
                        _ = fi38
                        _ = v39
                    }
                case "foo bar":
                    fmt.Println(v2, v2)
                case "":
                    if ((-1 >= -1) || false) {
                        if true {
                            var v40 bool = (1 == uint8(v2))
                            var v41 uint64 = (v33 << ((genFn1(((v1 + ("a" + v1)) != v5), -1)) & 63))
                            _ = v40
                            _ = v41
                        }
                    }
                default:
                    var v42 uint = 1
                    _ = v42
                }
            }
            var v43 int64 = -658
            if false {
                var v44 int8 = 42
                var v45 uint64 = (v33 + 18446744073709551615)
                _ = v44
                _ = v45
            }
            _ = v43
        }
        _ = v34
    }
    var v46 uint64 = v33
    v33 -= uint64(v2)
    var v47 uint = 131
    for fi48 := 0; fi48 < 8; fi48++ {
        v49 := GenS3{F0: genFn1((v4 != v4), 104), F1: 0}
        var v50 int32 = 2147483647
        _ = fi48
        _ = v49
        _ = v50
    }
    if true {
        var v51 uint32 = 2147483648
        fmt.Println(v5)
        _ = v51
    } else {
        var v52 uint16 = 42
        _ = v52
    }
    _ = len(v3)
    for fi53 := 0; fi53 < 3; fi53++ {
        var v54 uint = v47
        if true {
            for fi55 := 0; fi55 < 7; fi55++ {
                var v56 int32 = 0
                v33 = v46
                switch (v1 + v5) {
                case "\n\t":
                    v57 := GenS2{F0: 773}
                    _ = v57
                }
                _ = fi55
                _ = v56
            }
            switch -1 {
            case 32768:
                v58 := GenS4{F0: ((!((9223372036854775807 < -1) || false)) || (!true)), F1: true}
                _ = v58
            case 562:
                var v59 int32 = 0
                _ = v59
            case 4294967295:
                for fi60 := 0; fi60 < 9; fi60++ {
                    var v61 int64 = 1
                    _ = fi60
                    _ = v61
                }
            }
            var v62 uint64 = v33
            _ = v62
        } else {
            for fi63 := 0; fi63 < 11; fi63++ {
                var v64 string = "hello"
                v46 = (v46 + v33)
                _ = fi63
                _ = v64
            }
            v65 := GenS4{F0: (false || true), F1: false}
            _ = v65
        }
        if (false && (v4 < (v4 / (v4 - float64(0.0))))) {
            switch (v5 + "foo bar") {
            case "foo bar":
                var v66 float32 = ((0.0 + -2.25) * ((1.5 * (1.5 * 0.1)) / 1.0))
                _ = v66
            default:
                v33 -= 0
            }
            v67 := GenS5{F0: v2, F1: 115, F2: 1}
            _ = v67
        }
        _ = fi53
        _ = v54
    }
    for fi68 := 0; fi68 < 5; fi68++ {
        v33 += 692
        _ = fi68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
