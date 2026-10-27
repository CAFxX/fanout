package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 bool; F2 uint8 }
type GenS5 struct { F0 int16 }
func genFn0(p0 string, p1 uint8) int64 {
    v7 := GenS1{F0: (p1 * p1)}
    var v8 uint64 = 127
    var v9 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return -683
}
func genFn1(p0 bool, p1 int16) uint64 {
    fmt.Println(p0, p0)
    p0 = ((1.5 * (3.14159 + -2.25)) <= (3.14159 - 0.1))
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := []string{("foo bar" + "hello"), "", ("hello" + "")}
    var v2 float32 = (1e10 / 1.5)
    v3 := make(map[uint32]bool)
    v4 := []string{"", "hello", "a"}
    var v5 int16 = -1
    var v6 uint64 = 998
    fmt.Println(v2, v5, v6)
    var v10 int32 = 428
    v11 := GenS2{F0: 230, F1: (v6 * v6)}
    v6 = (v6 + 42)
    fmt.Println(v10, v5)
    if (!(int16(v10) < v5)) {
        for fi12 := 0; fi12 < 4; fi12++ {
            switch ("" + ("hello" + "foo bar")) {
            case "":
                if (!((v5 >> ((256) & 63)) != 32767)) {
                    switch fi12 {
                    case 128:
                        delete(v3, 1)
                    case 345:
                        for fi13 := 0; fi13 < 2; fi13++ {
                            var v14 bool = true
                            fmt.Println(v10, v14)
                            fmt.Println(v2, v5)
                            fmt.Println(v5, fi13)
                            v15 := GenS1{F0: 1}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    default:
                        v16 := GenS1{F0: 228}
                        _ = v16
                    }
                    if true {
                        v3[2] = false
                    } else {
                        var v17 uint64 = 1
                        switch (("a" + ("\n\t" + "foo bar")) + ("hello" + ("" + ("\n\t" + "hello")))) {
                        case "a":
                            switch (fi12 ^ (^(fi12))) {
                            case -1:
                                v2 += 0.1
                            case 1:
                                var v18 float64 = 3.14159
                                _ = v18
                            case 32767:
                                switch 1 {
                                case 229:
                                    switch 782 {
                                    case -32:
                                        var v19 uint64 = genFn1((!false), (v5 / (((v5 & (v5 / ((-657) | 1)))) | 1)))
                                        _ = v19
                                    }
                                default:
                                    var v20 int = fi12
                                    _ = v20
                                }
                            default:
                                v21 := GenS2{F0: 0, F1: 42}
                                _ = v21
                            }
                        }
                        _ = v17
                    }
                } else {
                    v5 += (v5 >> ((264) & 63))
                    v5 = 42
                }
            }
            switch "\n\t" {
            case "":
                v22 := GenS1{F0: 229}
                _ = v22
            case "hello":
                v23 := GenS3{F0: false}
                _ = v23
            }
            v24 := GenS1{F0: 143}
            _ = fi12
            _ = v24
        }
        if (false && (v2 >= -2.25)) {
            v25 := GenS3{F0: (!false)}
            var v26 bool = (v5 >= (v5 | v5))
            var v27 float32 = v2
            _ = v25
            _ = v26
            _ = v27
        } else {
            v28 := GenS1{F0: 255}
            fmt.Println(v10, v10)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v28
        }
        fmt.Println(v5, v10)
    }
    var v29 float32 = (v2 * v2)
    switch 32767 {
    case -12:
        if (true || true) {
            if true {
                switch 32768 {
                case -9223372036854775808:
                    _ = len(v1)
                case -237:
                    v29 = (float32(3.14159) * 1.5)
                default:
                    fmt.Println(v2, v10, v10)
                }
                var v30 int64 = 42
                _ = v30
            }
        } else {
            fmt.Println(v10, v2, v10)
            if false {
                var v31 int16 = v5
                fmt.Println(v6, v2, v2)
                _ = v31
            } else {
                v32 := GenS1{F0: 255}
                _ = v32
            }
            v29 = ((((float32(3.14159) * float32(1.5)) * float32(1.5)) - v2) * v2)
        }
    }
    fmt.Println(v10)
    v5 += 0
    if false {
        if true {
            if true {
                v33 := GenS3{F0: false}
                v6 = (^(v6))
                var v34 float64 = 3.14159
                _ = v33
                _ = v34
            }
            v10 = int32(v6)
            var v35 int16 = v5
            _ = v35
        }
        v10 -= -1
        v6 = genFn1(false, (v5 >> ((0) & 63)))
    } else {
        v36 := GenS4{F0: ("" + "a"), F1: (!(-1 != v5)), F2: 239}
        switch ("foo bar" + ("" + "\n\t")) {
        case "hello":
            var v37 uint32 = 42
            _ = v37
        case "":
            var v38 int64 = -758
            _ = v38
        default:
            v39 := GenS5{F0: v5}
            _ = v39
        }
        for fi40 := 0; fi40 < 7; fi40++ {
            v41 := v3[2147483648]
            if false {
                var v42 uint16 = 168
                v43 := GenS4{F0: "\n\t", F1: (!(9223372036854775807 <= 9223372036854775807)), F2: 1}
                _ = v42
                _ = v43
            } else {
                fmt.Println(v6, v5)
                for fi44 := 0; fi44 < 5; fi44++ {
                    v5 = (v5 << ((618) & 63))
                    v29 += 1.5
                    var v45 string = "\n\t"
                    _ = fi44
                    _ = v45
                }
            }
            if (uint(v6) < 255) {
                var v46 bool = v41
                _ = v46
            }
            _ = fi40
            _ = v41
        }
        _ = v36
    }
    v47 := GenS3{F0: true}
    var v48 uint8 = 145
    var v49 string = (((("hello" + "foo bar") + "foo bar") + "a") + (("foo bar" + ("foo bar" + "")) + ("foo bar" + "foo bar")))
    var v50 int8 = 0
    for fi51 := 0; fi51 < 2; fi51++ {
        var v52 int8 = -9
        _ = fi51
        _ = v52
    }
    v5 = v5
    switch 32767 {
    case 128:
        for fi53 := 0; fi53 < 10; fi53++ {
            v2 -= v2
            _ = len(v4)
            _ = fi53
        }
    case 127:
        v54 := GenS2{F0: (v48 ^ (^(v48))), F1: genFn1(true, -874)}
        _ = v54
    case 457:
        v55 := GenS1{F0: v48}
        _ = v55
    default:
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
