package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = false
    v3 := []uint32{0, 2147483648}
    v4 := make(map[string]int)
    v5 := make(map[int]bool)
    v6 := []uint32{2147483647, 4294967295}
    fmt.Println(v2)
    _ = len(v4)
    v2 = true
    _ = len(v4)
    v7 := GenS1{F0: 1, F1: 5}
    v2 = v2
    var v8 int32 = 2147483647
    v3 = append(v3, 65535)
    for fi9 := 0; fi9 < 7; fi9++ {
        if true {
            v8 = (v8 / ((-2147483648) | 1))
            if len(v3) > 0 {
                v3[0] = 333
            }
        } else {
            var v10 string = "hello"
            v8 = -2147483648
            _ = v10
        }
        var v11 uint64 = 1
        var v12 int8 = int8(fi9)
        _ = fi9
        _ = v11
        _ = v12
    }
    if v2 {
        v2 = (!v2)
    } else {
        delete(v5, 970)
        switch "\n\t" {
        case "foo bar":
            if (-304 != -603) {
                v1[-46] = "hello"
                v2 = (1.5 != ((-2.25 + 1.5) + 3.14159))
            } else {
                v2 = (v2 && (1 > 0))
                var v13 uint = 0
                _ = v13
            }
        default:
            if v2 {
                v14 := GenS1{F0: 533, F1: 100}
                if true {
                    fmt.Println(v8)
                    v2 = (1 == 96)
                    v2 = true
                } else {
                    var v15 string = "a"
                    var v16 uint64 = 42
                    _ = v15
                    _ = v16
                }
                v17 := v4["a"]
                _ = v14
                _ = v17
            } else {
                if true {
                    v18 := GenS2{F0: 2147483647}
                    _ = v18
                } else {
                    v19 := v5[-9223372036854775808]
                    switch ((("foo bar" + "\n\t") + "a") + ("a" + "a")) {
                    case "\n\t":
                        if (127 >= -128) {
                            var v20 int16 = 42
                            _ = v20
                        }
                    case "hello":
                        if (128 > 788) {
                            v6 = append(v6, 2147483648)
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = len(v1)
                        } else {
                            if len(v6) > 0 {
                                v6[0] = 2
                            }
                        }
                    case "a":
                        v21 := GenS3{F0: "hello"}
                        _ = v21
                    default:
                        v19 = ((0 != 1) || (((1 != 1) || v2) && (!v2)))
                    }
                    var v22 int8 = 42
                    _ = v19
                    _ = v22
                }
            }
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println(v8)
    var v23 int64 = -1
    fmt.Println(v23)
    v24 := GenS4{F0: 100, F1: 42, F2: v8}
    var v25 uint64 = 205
    switch 842 {
    case 32768:
        if false {
            v26 := GenS5{F0: 1, F1: 255}
            v27 := GenS1{F0: -32768, F1: -1}
            for fi28 := 0; fi28 < 8; fi28++ {
                var v29 float32 = (-2.25 + 3.14159)
                v30 := GenS2{F0: 1}
                v2 = v2
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v26
            _ = v27
        } else {
            if (true || (!true)) {
                for fi31 := 0; fi31 < 13; fi31++ {
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 float64 = -2.25
                        delete(v4, "a")
                        var v34 uint8 = 127
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    switch 4294967295 {
                    case 620:
                        v35 := GenS6{F0: 42, F1: 2}
                        _ = v35
                    case 32767:
                        var v36 uint8 = 128
                        _ = v36
                    case -175:
                        for fi37 := 0; fi37 < 18; fi37++ {
                            var v38 uint32 = 355
                            if len(v6) > 0 {
                                v6[0] = uint32(v25)
                            }
                            var v39 float64 = ((0.0 + 1.5) + (1e10 - (0.1 - float64((0.0 - -2.25)))))
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    }
                    _ = fi31
                }
            } else {
                for fi40 := 0; fi40 < 8; fi40++ {
                    v23 += 0
                    var v41 bool = (2 <= uint32(v25))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi40
                    _ = v41
                }
            }
            var v42 uint64 = (v25 + v25)
            _ = v42
        }
    case -9223372036854775808:
        if false {
            var v43 bool = v2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v43
        } else {
            v5[2147483647] = false
        }
    }
    if v2 {
        for fi44 := 0; fi44 < 16; fi44++ {
            var v45 uint32 = 2
            var v46 uint64 = 18446744073709551615
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
