package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8, p1 uint) uint16 {
    var v5 uint8 = 1
    p0 -= 0
    switch "\n\t" {
    case "a":
        var v6 uint32 = 0
        _ = v6
    default:
        v7 := []int{-479, -9223372036854775808}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 65535
}

func main() {
    var v1 bool = false
    v2 := []string{"", "hello"}
    var v3 string = "foo bar"
    v4 := make(map[int]uint32)
    var v8 int32 = 0
    if len(v2) > 0 {
        v2[0] = v3
    }
    if v1 {
        switch v3 {
        case "":
            fmt.Println(v1, v1)
        case "foo bar":
            for fi9 := 0; fi9 < 9; fi9++ {
                delete(v4, fi9)
                var v10 int8 = 0
                _ = fi9
                _ = v10
            }
        case "a":
            fmt.Println(v1)
        }
    } else {
        var v11 uint16 = 0
        _ = v11
    }
    var v12 string = ""
    _ = len(v2)
    v8 = -892
    switch v3 {
    case "hello":
        var v13 string = (("\n\t" + v12) + v3)
        _ = v13
    case "":
        v1 = (v1 || (v1 || ((42 == 810) || false)))
    case "a":
        switch "foo bar" {
        case "foo bar":
            var v14 int8 = 42
            _ = v14
        case "a":
            var v15 int16 = 42
            _ = v15
        default:
            v1 = v1
        }
    default:
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    }
    if false {
        var v16 float32 = 1e10
        if len(v2) > 0 {
            v2[0] = ""
        }
        _ = v16
    }
    v8 += v8
    var v17 int16 = 503
    var v18 int16 = v17
    for fi19 := 0; fi19 < 12; fi19++ {
        fmt.Println(fi19)
        _ = fi19
    }
    var v20 float32 = 0.1
    for fi21 := 0; fi21 < 3; fi21++ {
        if (!v1) {
            switch v3 {
            case "":
                if (((-120 >= 127) && false) && (true && false)) {
                    var v22 int64 = -9223372036854775808
                    _ = v22
                }
            case "foo bar":
                var v23 bool = (!v1)
                _ = v23
            case "a":
                v24 := GenS1{F0: 18446744073709551615}
                _ = v24
            }
            for fi25 := 0; fi25 < 7; fi25++ {
                switch -1 {
                case -2147483648:
                    fmt.Println(fi21, v20)
                default:
                    var v26 uint8 = 127
                    if v1 {
                        v17 = int16(fi25)
                        v4[(fi25 >> ((65535) & 63))] = 129
                        _ = len(v4)
                    } else {
                        _ = len(v2)
                        switch fi21 {
                        case 4294967295:
                            v2 = append(v2, v12)
                        case 1:
                            switch v3 {
                            case "a":
                                v3 = "hello"
                            case "\n\t":
                                var v27 int64 = -9223372036854775808
                                _ = v27
                            case "":
                                v28 := GenS1{F0: 255}
                                switch v3 {
                                case "":
                                    v1 = v1
                                case "a":
                                    v29 := GenS1{F0: 18446744073709551615}
                                    _ = v29
                                default:
                                    _ = len(v4)
                                }
                                _ = v28
                            }
                            var v30 int32 = (v8 % ((v8) | 1))
                            _ = v30
                        default:
                            if ((0 > -9223372036854775808) && ((699 >= 42) || (!false))) {
                                v26 += 208
                            }
                        }
                        switch "foo bar" {
                        case "":
                            delete(v4, (fi21 ^ (fi21 % ((fi21) | 1))))
                        case "hello":
                            switch v3 {
                            case "foo bar":
                                v1 = false
                            case "a":
                                var v31 uint8 = v26
                                _ = v31
                            case "":
                                fmt.Println(v8, v26, v17)
                            }
                            v8 = 1
                        case "foo bar":
                            v32 := GenS1{F0: 0}
                            _ = v32
                        }
                    }
                    var v33 int16 = (-(v17))
                    _ = v26
                    _ = v33
                }
                var v34 int = 2147483647
                _ = fi25
                _ = v34
            }
        } else {
            var v35 uint8 = 153
            fmt.Println(fi21, v18)
            _ = v35
        }
        v36 := GenS1{F0: 18446744073709551615}
        _ = fi21
        _ = v36
    }
    v37 := GenS2{F0: v18}
    v18 = 125
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
    _ = v8
}
