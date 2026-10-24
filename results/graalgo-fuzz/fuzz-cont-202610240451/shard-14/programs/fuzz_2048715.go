package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint16) uint32 {
    for fi7 := 0; fi7 < 15; fi7++ {
        p0 = (^(p0))
        _ = fi7
    }
    p0 = p0
    switch 32768 {
    case 1:
        switch "foo bar" {
        case "hello":
            if (true && (true && false)) {
                switch -238 {
                case -462:
                    for fi8 := 0; fi8 < 7; fi8++ {
                        for fi9 := 0; fi9 < 3; fi9++ {
                            var v10 uint8 = 0
                            _ = fi9
                            _ = v10
                        }
                        if ("a" < ("" + ("" + "hello"))) {
                            fmt.Println(p0)
                        } else {
                            fmt.Println(fi8)
                            v11 := []int{(fi8 ^ fi8), -947}
                            if len(v11) > 0 {
                                v11[0] = fi8
                            }
                            _ = v11
                        }
                        p0 -= p0
                        _ = fi8
                    }
                case -634:
                    v12 := GenS1{F0: 4294967295}
                    _ = v12
                case -246:
                    switch "" {
                    case "\n\t":
                        p0 = p0
                    case "":
                        p0 = (p0 + (^(p0)))
                    case "foo bar":
                        if true {
                            if (true || true) {
                                p0 = (^(p0))
                                for fi13 := 0; fi13 < 8; fi13++ {
                                    var v14 uint8 = 255
                                    for fi15 := 0; fi15 < 7; fi15++ {
                                        v16 := []int{-869, fi15, (fi13 | (fi15 % ((-9223372036854775808) | 1))), fi15}
                                        fmt.Println(fi13, p0, fi15)
                                        _ = fi15
                                        _ = v16
                                    }
                                    var v17 string = ""
                                    _ = fi13
                                    _ = v14
                                    _ = v17
                                }
                                var v18 int = 4294967295
                                _ = v18
                            } else {
                                for fi19 := 0; fi19 < 5; fi19++ {
                                    var v20 uint32 = 2147483648
                                    v20 = v20
                                    for fi21 := 0; fi21 < 11; fi21++ {
                                        var v22 uint16 = p0
                                        _ = fi21
                                        _ = v22
                                    }
                                    _ = fi19
                                    _ = v20
                                }
                                p0 -= (^(p0))
                            }
                        } else {
                            switch ("hello" + ((("\n\t" + (("foo bar" + "foo bar") + "")) + "\n\t") + "foo bar")) {
                            case "hello":
                                var v23 float64 = float64((1e10 * 3.14159))
                                _ = v23
                            }
                        }
                    default:
                        for fi24 := 0; fi24 < 14; fi24++ {
                            var v25 uint8 = 2
                            v26 := GenS1{F0: 44}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                }
                var v27 int32 = -345
                var v28 uint = 690
                _ = v27
                _ = v28
            } else {
                if false {
                    p0 -= 1
                    p0 = (p0 << ((831) & 63))
                    v29 := GenS1{F0: 2}
                    _ = v29
                } else {
                    var v30 int8 = -119
                    v31 := []int{}
                    var v32 bool = (!(-689 > 0))
                    _ = v30
                    _ = v31
                    _ = v32
                }
                var v33 int16 = 32767
                _ = v33
            }
        case "":
            v34 := GenS1{F0: 0}
            _ = v34
        case "a":
            v35 := GenS1{F0: 2147483648}
            _ = v35
        }
    }
    _ = p0
    return 439
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]string)
    v3 := []int{-1, -2147483648, 51}
    var v4 int16 = int16(v1)
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    fmt.Println(v1)
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 string = "\n\t"
        _ = fi36
        _ = v37
    }
    v5["foo bar"] = "foo bar"
    v1 += -9223372036854775808
    if true {
        v38 := GenS2{F0: v1}
        _ = v38
    }
    fmt.Println(v1)
    var v39 int32 = 0
    fmt.Println(v1, v4)
    switch ("foo bar" + "a") {
    case "hello":
        if false {
            var v40 uint16 = 538
            switch -9223372036854775808 {
            case -707:
                var v41 bool = true
                _ = v41
            default:
                var v42 int16 = (v4 - v4)
                _ = v42
            }
            _ = v40
        }
    case "a":
        v43 := GenS2{F0: v1}
        _ = v43
    default:
        fmt.Println(v1, v4, v4)
    }
    fmt.Println(v4, v1, v1)
    var v44 float32 = 1.5
    switch 1 {
    case 1:
        for fi45 := 0; fi45 < 1; fi45++ {
            v46 := GenS2{F0: 0}
            v47 := GenS1{F0: 2147483647}
            _ = fi45
            _ = v46
            _ = v47
        }
    case 9223372036854775807:
        var v48 int8 = 1
        _ = v48
    default:
        v39 = v39
    }
    v4 = (v4 / ((v4) | 1))
    var v49 bool = true
    fmt.Println(v1, v49)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v44
    _ = v49
    _ = v5
    _ = v6
}
