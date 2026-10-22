package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 uint8, p1 uint64, p2 int32) bool {
    var v5 int16 = 42
    p2 = 42
    for fi6 := 0; fi6 < 6; fi6++ {
        switch "" {
        case "a":
            p0 = (p0 | (p0 + p0))
        case "foo bar":
            fmt.Println(v5, p2)
        case "\n\t":
            if (1 < 2147483647) {
                v7 := []int{-1, fi6, (fi6 / (((fi6 + (fi6 * (fi6 << ((0) & 63))))) | 1)), 4294967295}
                v8 := GenS1{F0: 18446744073709551615, F1: 33, F2: (((!true) && false) && true)}
                var v9 float32 = 0.0
                _ = v7
                _ = v8
                _ = v9
            }
        }
        if false {
            fmt.Println(p1)
            v10 := []int{0}
            _ = v10
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return true
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = -6
    var v3 int16 = 691
    v4 := make(map[uint32]uint32)
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: 421, F1: 4294967295, F2: true}
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 1, F1: 816, F2: (2147483648 == 2147483648)}
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v2, fi14)
        fmt.Println(v3)
        _ = fi14
    }
    v15 := GenS2{F0: v3}
    var v16 bool = true
    delete(v4, 4294967295)
    delete(v1, "hello")
    var v17 int8 = 90
    var v18 string = "\n\t"
    for fi19 := 0; fi19 < 14; fi19++ {
        if true {
            v3 += v3
            for fi20 := 0; fi20 < 1; fi20++ {
                _ = len(v4)
                _ = fi20
            }
        } else {
            var v21 uint = 18446744073709551615
            v16 = ((v16 || (2147483647 == -660)) || true)
            switch v18 {
            case "hello":
                var v22 int32 = -1
                _ = v22
            }
            _ = v21
        }
        fmt.Println(v2)
        _ = fi19
    }
    if (255 != 0) {
        v23 := v4[164]
        _ = v23
    } else {
        v16 = false
        var v24 uint32 = 233
        v25 := GenS3{F0: 806}
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := GenS1{F0: 1, F1: 336, F2: (!(1 >= 127))}
        var v28 uint16 = 452
        var v29 uint16 = 124
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: 42, F1: 65535, F2: v16}
    for fi31 := 0; fi31 < 1; fi31++ {
        v18 = v18
        for fi32 := 0; fi32 < 12; fi32++ {
            if true {
                if (!(75 == 937)) {
                    _ = len(v1)
                    v33 := GenS1{F0: 4294967295, F1: 42, F2: v16}
                    fmt.Println(v16, v17)
                    v16 = ((v2 != (fi31 & (fi31 + 2147483647))) && ((v3 >> ((65535) & 63)) > v3))
                    if (fi32 == 4294967295) {
                        switch v18 {
                        case "a":
                            v34 := GenS4{F0: v18, F1: 42}
                            _ = v34
                        default:
                            var v35 int8 = -116
                            _ = v35
                        }
                    } else {
                        var v36 int64 = -9223372036854775808
                        var v37 bool = false
                        v36 = (v36 << ((2) & 63))
                        if (2147483648 >= 0) {
                            v38 := v4[4294967295]
                            v37 = (!true)
                            _ = v38
                        } else {
                            var v39 uint64 = 414
                            _ = v39
                        }
                        if (v36 != 433) {
                            var v40 float64 = (1.5 * -2.25)
                            v18 = (v18 + v18)
                            var v41 int64 = 390
                            _ = v40
                            _ = v41
                        }
                        _ = v36
                        _ = v37
                    }
                    _ = v33
                } else {
                    var v42 float32 = (1.5 + -2.25)
                    v2 = 865
                    _ = v42
                }
                v2 = 0
                var v43 bool = (0.0 < 1.5)
                _ = v43
            } else {
                v18 = "hello"
                if (!((!(255 <= 18446744073709551615)) || v16)) {
                    delete(v1, v18)
                    var v44 int = (fi32 / ((v2) | 1))
                    delete(v1, v18)
                    _ = v44
                }
                if ("foo bar" < ((("hello" + "\n\t") + (v18 + v18)) + ("\n\t" + ("\n\t" + ((v18 + ("foo bar" + "foo bar")) + (v18 + v18)))))) {
                    v45 := v4[65535]
                    _ = v45
                }
            }
            if v16 {
                fmt.Println(v18, fi31)
                v3 = (v3 | 32767)
                v3 = v3
            }
            v46 := GenS3{F0: 374}
            _ = fi32
            _ = v46
        }
        for fi47 := 0; fi47 < 17; fi47++ {
            fmt.Println(v2)
            v3 += v3
            fmt.Println(v3, v17, v16)
            _ = fi47
        }
        _ = fi31
    }
    var v48 uint16 = 80
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v48
}
