package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint64; F1 string; F2 int16 }
type GenS5 struct { F0 string; F1 uint; F2 string }
type GenS6 struct { F0 int8 }
func genFn0() int8 {
    var v6 uint = 18446744073709551615
    v7 := GenS1{F0: -994}
    v8 := []int{128, 32767, 9223372036854775807, 173}
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint32) uint16 {
    fmt.Println(p0)
    var v9 int32 = 2147483647
    _ = p0
    _ = v9
    return 65535
}

func main() {
    var v1 int8 = 0
    var v2 uint8 = 0
    v3 := []string{"a", "hello"}
    var v4 int = 42
    var v5 int8 = v1
    var v10 bool = (!(-1 > 42))
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS2{F0: 42}
        v3 = append(v3, "")
        var v13 bool = (v10 && ((2147483648 != 987) || (float64(3.14159) <= 0.1)))
        _ = fi11
        _ = v12
        _ = v13
    }
    if v10 {
        var v14 int16 = 1000
        _ = v14
    }
    switch ("" + "") {
    case "a":
        switch v4 {
        case -127:
            var v15 int = 127
            _ = v15
        case 0:
            switch ("foo bar" + "a") {
            case "hello":
                v1 = (v5 / ((-112) | 1))
            case "\n\t":
                var v16 int8 = v5
                _ = v16
            }
        case 493:
            switch 887 {
            case 235:
                var v17 int16 = -1
                _ = v17
            }
        default:
            v2 = v2
        }
    }
    if len(v3) > 0 {
        v3[0] = ("hello" + "hello")
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    var v18 bool = (938 <= 76)
    if ((((!(v2 < v2)) || v10) && v10) && (0.1 >= 0.0)) {
        v19 := GenS1{F0: 32767}
        _ = v19
    } else {
        if true {
            var v20 uint8 = v2
            if false {
                v3 = append(v3, "foo bar")
            } else {
                if (1 <= v2) {
                    v21 := GenS3{F0: v4, F1: 738}
                    for fi22 := 0; fi22 < 20; fi22++ {
                        var v23 int64 = 725
                        _ = fi22
                        _ = v23
                    }
                    switch "" {
                    case "hello":
                        if len(v3) > 0 {
                            v3[0] = ("" + "hello")
                        }
                    case "foo bar":
                        if v18 {
                            for fi24 := 0; fi24 < 8; fi24++ {
                                fmt.Println(fi24)
                                if ((v4 * -2147483648) >= v4) {
                                    var v25 uint64 = 868
                                    _ = v25
                                } else {
                                    fmt.Println(v4, v4)
                                    fmt.Println(v4)
                                    if (905 != 65535) {
                                        if true {
                                            switch "\n\t" {
                                            case "hello":
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v26 := GenS2{F0: v2}
                                                    _ = v26
                                                default:
                                                    v10 = v10
                                                }
                                            default:
                                                switch (v4 * -108) {
                                                case 4294967295:
                                                    for fi27 := 0; fi27 < 5; fi27++ {
                                                        fmt.Println(fi27, v10, v2)
                                                        fmt.Println(v20)
                                                        _ = fi27
                                                    }
                                                case 42:
                                                    if (!false) {
                                                        switch 32767 {
                                                        case -1:
                                                            fmt.Println(v10, v2)
                                                        case -2147483648:
                                                            var v28 int32 = 42
                                                            _ = v28
                                                        }
                                                    } else {
                                                        v20 -= v20
                                                        if (65535 > 930) {
                                                            var v29 string = ("hello" + ("a" + (("" + ("" + "")) + "foo bar")))
                                                            _ = v29
                                                        }
                                                    }
                                                case -964:
                                                    for fi30 := 0; fi30 < 5; fi30++ {
                                                        v10 = true
                                                        if true {
                                                            switch v4 {
                                                            case 695:
                                                                var v31 float64 = -2.25
                                                                _ = v31
                                                            }
                                                        } else {
                                                            switch v4 {
                                                            case -516:
                                                                if v18 {
                                                                    if false {
                                                                        var v32 uint = 2
                                                                        fmt.Println(fi30, v32, v5)
                                                                        _ = v32
                                                                    }
                                                                    fmt.Println(v1)
                                                                } else {
                                                                    switch (fi24 % ((v4) | 1)) {
                                                                    case -1:
                                                                        v33 := GenS1{F0: 42}
                                                                        _ = v33
                                                                    case -758:
                                                                        if (1.5 <= float32(((1.5 * ((0.1 + 1e10) + -2.25)) / (3.14159 / 3.14159)))) {
                                                                            v34 := GenS3{F0: (fi24 << ((1) & 63)), F1: 1}
                                                                            v35 := GenS3{F0: v4, F1: 2}
                                                                            switch "" {
                                                                            case "hello":
                                                                                _ = len(v3)
                                                                            case "\n\t":
                                                                                var v36 uint64 = 0
                                                                                _ = v36
                                                                            case "a":
                                                                                var v37 int = (v4 / ((fi24) | 1))
                                                                                _ = v37
                                                                            }
                                                                            _ = v34
                                                                            _ = v35
                                                                        }
                                                                    default:
                                                                        fmt.Println(v10, v2, v4)
                                                                        v18 = (127 <= (v5 >> ((972) & 63)))
                                                                        v18 = (0 <= 2147483648)
                                                                        _ = len(v3)
                                                                        switch v4 {
                                                                        case 42:
                                                                            fmt.Println(v18, v2)
                                                                        }
                                                                        v38 := GenS3{F0: fi30, F1: 4294967295}
                                                                        _ = v38
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        if ((v1 % ((42) | 1)) > genFn0()) {
                                                                            v18 = v10
                                                                            if (42 < 1) {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                var v39 int8 = -25
                                                                                _ = v39
                                                                            } else {
                                                                                v20 = v2
                                                                            }
                                                                            var v40 int16 = -1
                                                                            _ = v40
                                                                        }
                                                                    case "hello":
                                                                        v41 := GenS3{F0: -2147483648, F1: 2}
                                                                        var v42 uint64 = 1
                                                                        v1 -= 0
                                                                        _ = v41
                                                                        _ = v42
                                                                    case "a":
                                                                        var v43 string = "a"
                                                                        _ = v43
                                                                    }
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v2 -= (v20 & (v2 | (v20 | (v20 ^ 42))))
                                                        _ = fi30
                                                    }
                                                }
                                            }
                                        } else {
                                            v18 = v10
                                        }
                                        v44 := GenS4{F0: 0, F1: "", F2: 682}
                                        _ = v44
                                    }
                                }
                                var v45 int16 = -1
                                _ = fi24
                                _ = v45
                            }
                            var v46 uint8 = (v20 & 42)
                            if v10 {
                                v10 = false
                            } else {
                                fmt.Println(v20, v2, v1)
                                v2 = v20
                                if v10 {
                                    var v47 uint32 = 918
                                    _ = v47
                                } else {
                                    var v48 uint32 = 2147483648
                                    fmt.Println(v46, v46)
                                    _ = v48
                                }
                            }
                            _ = v46
                        } else {
                            v4 = 32768
                        }
                    }
                    _ = v21
                } else {
                    switch (v4 << ((2) & 63)) {
                    case 127:
                        if v18 {
                            if false {
                                var v49 uint64 = 312
                                _ = v49
                            }
                            v50 := GenS1{F0: 32767}
                            _ = v50
                        }
                    case -721:
                        v3 = append(v3, ("foo bar" + (("hello" + "\n\t") + ("foo bar" + "foo bar"))))
                    }
                }
                var v51 uint16 = 65535
                switch v4 {
                case 127:
                    v52 := GenS5{F0: "foo bar", F1: 435, F2: "\n\t"}
                    _ = v52
                }
                _ = v51
            }
            _ = v20
        } else {
            if (float64(float32((0.1 + 1e10))) <= -2.25) {
                var v53 int32 = 0
                for fi54 := 0; fi54 < 1; fi54++ {
                    for fi55 := 0; fi55 < 18; fi55++ {
                        v56 := GenS6{F0: -128}
                        _ = fi55
                        _ = v56
                    }
                    _ = fi54
                }
                _ = v53
            }
            v5 -= (v5 >> ((1) & 63))
            switch v4 {
            case 4294967295:
                if (341 != 0) {
                    var v57 int32 = 2147483647
                    v58 := GenS6{F0: (v1 / (((^(v5))) | 1))}
                    _ = v57
                    _ = v58
                }
            case 9223372036854775807:
                v4 -= 9223372036854775807
            }
        }
        fmt.Println(v1, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
