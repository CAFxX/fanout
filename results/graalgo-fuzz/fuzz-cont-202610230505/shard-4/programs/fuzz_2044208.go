package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int32; F1 uint8; F2 int32 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint8; F1 int8; F2 bool }
func genFn0() uint32 {
    v5 := GenS1{F0: -28, F1: 65535, F2: 42}
    var v6 uint8 = 71
    switch 42 {
    case -2147483648:
        var v7 int = 31
        _ = v7
    default:
        var v8 uint16 = 42
        _ = v8
    }
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int8, p1 string) int {
    v9 := GenS1{F0: p0, F1: 255, F2: 4294967295}
    v10 := GenS1{F0: p0, F1: 42, F2: 172}
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            p0 = p0
            v13 := GenS1{F0: (^(p0)), F1: 42, F2: 42}
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 4; fi14++ {
            switch (((p1 + p1) + p1) + p1) {
            case "\n\t":
                fmt.Println(fi14, fi11)
            default:
                if true {
                    switch "foo bar" {
                    case "hello":
                        v15 := GenS2{F0: 375}
                        v16 := []int{(fi11 ^ (fi11 % (((fi11 * (fi11 % (((fi14 / ((fi14) | 1))) | 1)))) | 1)))}
                        _ = v15
                        _ = v16
                    case "":
                        v17 := GenS2{F0: -1}
                        _ = v17
                    case "foo bar":
                        if false {
                            var v18 int8 = p0
                            v19 := GenS2{F0: 1}
                            _ = v18
                            _ = v19
                        }
                    }
                    var v20 float64 = 3.14159
                    var v21 int32 = -200
                    v21 -= (v21 >> ((1) & 63))
                    var v22 string = (p1 + (p1 + p1))
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    fmt.Println(fi14, p0)
                    var v23 float64 = -2.25
                    p1 = ("a" + (p1 + ""))
                    _ = v23
                }
            }
            p1 = ((p1 + (p1 + "foo bar")) + "foo bar")
            _ = fi14
        }
        _ = fi11
    }
    var v24 uint8 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v24
    _ = v9
    return 1
}

func main() {
    v1 := []string{}
    var v2 bool = false
    v3 := make(map[int]int)
    var v4 uint64 = 628
    v4 = v4
    var v25 int8 = 100
    if (9223372036854775807 > 9223372036854775807) {
        var v26 int64 = -1
        _ = v26
    } else {
        v4 = v4
        var v27 float64 = (-2.25 + float64(0.0))
        _ = v27
    }
    if v2 {
        if v2 {
            if (289 > 4294967295) {
                switch ("\n\t" + "a") {
                case "":
                    var v28 float64 = 1e10
                    _ = v28
                }
                v4 -= 1
                v29 := v3[35]
                _ = v29
            } else {
                v30 := GenS3{F0: 18446744073709551615, F1: 59}
                v31 := GenS4{F0: 165, F1: 36, F2: 358}
                _ = v30
                _ = v31
            }
        }
        var v32 bool = (!(v2 && ((("hello" + ("\n\t" + "foo bar")) > "hello") || ((42 <= (v25 % ((v25) | 1))) && ("hello" <= "")))))
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        _ = v32
    }
    fmt.Println(v25)
    if true {
        if true {
            fmt.Println(v25, v4)
        } else {
            for fi33 := 0; fi33 < 18; fi33++ {
                v2 = ((v2 || (!v2)) || false)
                _ = fi33
            }
        }
        fmt.Println(v25, v2, v25)
        v2 = true
    } else {
        v2 = false
        var v34 int16 = 1
        _ = v34
    }
    if (-2.25 < 0.0) {
        var v35 uint8 = 0
        fmt.Println(v4, v2)
        _ = v35
    } else {
        for fi36 := 0; fi36 < 3; fi36++ {
            for fi37 := 0; fi37 < 4; fi37++ {
                if (!v2) {
                    var v38 uint64 = 649
                    if (-9223372036854775808 >= -1) {
                        fmt.Println(v38)
                    }
                    _ = v38
                } else {
                    if false {
                        if (-637 != 2147483647) {
                            var v39 uint32 = 1
                            v2 = ((v2 || (false || v2)) || v2)
                            _ = v39
                        } else {
                            fmt.Println(v2, v2, v25)
                        }
                        var v40 float32 = 1.5
                        v1 = append(v1, "a")
                        _ = v40
                    } else {
                        delete(v3, 4294967295)
                    }
                }
                v25 -= 100
                switch 1 {
                case -2147483648:
                    switch genFn1(v25, "") {
                    case -1:
                        var v41 float32 = float32(1e10)
                        _ = v41
                    }
                case 9223372036854775807:
                    v4 += 42
                }
                _ = fi37
            }
            if (v25 != (v25 ^ v25)) {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                v2 = true
                v1 = append(v1, "hello")
            }
            v42 := v3[genFn1((v25 / ((v25) | 1)), "hello")]
            _ = fi36
            _ = v42
        }
    }
    v3[42] = 127
    if v2 {
        var v43 uint32 = 2147483648
        var v44 uint = 207
        var v45 uint64 = (v4 / (((v4 >> ((v44) & 63))) | 1))
        _ = v43
        _ = v44
        _ = v45
    } else {
        if v2 {
            fmt.Println(v2, v25, v2)
            fmt.Println(v25)
            v2 = v2
        }
    }
    if v2 {
        switch "\n\t" {
        case "hello":
            if (0.0 >= (0.1 - -2.25)) {
                v46 := v3[-2147483648]
                v46 += v46
                v47 := GenS5{F0: 42, F1: (v4 | (v4 & v4)), F2: 18446744073709551615}
                _ = v46
                _ = v47
            } else {
                v48 := GenS5{F0: -9223372036854775808, F1: 734, F2: v4}
                var v49 string = (("hello" + ("hello" + (((("a" + ("" + "\n\t")) + "") + "\n\t") + ("\n\t" + "a")))) + "a")
                _ = v48
                _ = v49
            }
        case "":
            if false {
                var v50 uint64 = (v4 >> ((190) & 63))
                if (0.1 < 0.0) {
                    var v51 uint = 18446744073709551615
                    switch -2147483648 {
                    case 42:
                        _ = len(v1)
                    case 128:
                        if false {
                            if v2 {
                                v25 += v25
                            } else {
                                if v2 {
                                    _ = len(v1)
                                }
                                var v52 int64 = 9223372036854775807
                                fmt.Println(v2, v4, v52)
                                _ = v52
                            }
                            v53 := GenS3{F0: 42, F1: -2147483648}
                            v54 := GenS6{F0: 255, F1: v25, F2: ((v51 < v51) || v2)}
                            _ = v53
                            _ = v54
                        }
                    case 231:
                        v25 -= v25
                    default:
                        var v55 string = "a"
                        _ = v55
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v51
                }
                _ = v50
            }
        }
        v4 = 761
    } else {
        switch 4294967295 {
        case -1:
            v2 = (3.14159 <= 1e10)
        case 1:
            var v56 int64 = -9223372036854775808
            _ = v56
        case 2147483647:
            var v57 int = -2147483648
            _ = v57
        }
    }
    fmt.Println(v25, v2)
    v4 = 42
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
