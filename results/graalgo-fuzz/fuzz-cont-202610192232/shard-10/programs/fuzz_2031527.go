package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int8; F1 uint; F2 uint8 }
type GenS5 struct { F0 bool; F1 int16; F2 int8 }
func genFn0(p0 int64, p1 int32, p2 uint32) int {
    v4 := []int{42, 2147483647, 705, -1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v5 uint32 = p2
    switch "\n\t" {
    case "hello":
        for fi6 := 0; fi6 < 16; fi6++ {
            if (!false) {
                var v7 uint32 = (v5 >> ((42) & 63))
                _ = v7
            } else {
                switch ("" + "") {
                case "hello":
                    p2 += v5
                }
                if (true && true) {
                    v8 := GenS1{F0: (-86 > 127), F1: 2}
                    var v9 float32 = 0.1
                    _ = v8
                    _ = v9
                } else {
                    p0 += (-(p0))
                    var v10 int8 = int8(fi6)
                    _ = v10
                }
            }
            _ = len(v4)
            v11 := GenS2{F0: p1, F1: 0}
            _ = fi6
            _ = v11
        }
    case "a":
        for fi12 := 0; fi12 < 17; fi12++ {
            p0 = 1
            v13 := GenS3{F0: -1, F1: 181}
            _ = fi12
            _ = v13
        }
    case "":
        var v14 uint16 = 65535
        _ = v14
    default:
        if (13 == -128) {
            if (1.5 > 3.14159) {
                var v15 float32 = 1.5
                if false {
                    v15 = (-2.25 - 1e10)
                } else {
                    var v16 int32 = (p1 % ((p1) | 1))
                    var v17 int64 = 9223372036854775807
                    for fi18 := 0; fi18 < 10; fi18++ {
                        p1 = (p1 * (v16 ^ p1))
                        _ = fi18
                    }
                    _ = v16
                    _ = v17
                }
                _ = v15
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -736
}
func genFn1(p0 int64, p1 uint16) int64 {
    v19 := GenS3{F0: 1000, F1: 42}
    p0 -= (p0 << ((65535) & 63))
    if false {
        var v20 bool = (!(133 == 1000))
        var v21 string = "hello"
        _ = v20
        _ = v21
    } else {
        p1 = (^(p1))
    }
    v22 := []int{0, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v19
    _ = v22
    return (p0 ^ 605)
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 1e10
    var v3 uint16 = 803
    var v23 int32 = 830
    if true {
        v24 := GenS4{F0: 42, F1: 65535, F2: 0}
        _ = v24
    }
    v25 := GenS5{F0: ((v2 * v2) > v2), F1: -1, F2: 0}
    v26 := GenS5{F0: true, F1: 0, F2: 100}
    if (("a" < "\n\t") && false) {
        v2 = v2
    }
    fmt.Println(v2)
    var v27 int8 = -22
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 int8 = (-(v27))
        var v30 uint = 2
        var v31 uint16 = 414
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    v32 := GenS5{F0: ("foo bar" >= ("a" + "\n\t")), F1: 134, F2: 127}
    var v33 int = genFn0(528, v23, 863)
    switch v33 {
    case 32768:
        var v34 bool = true
        _ = v34
    case -852:
        v35 := GenS5{F0: (int64(v33) != 9223372036854775807), F1: -769, F2: (v27 << ((125) & 63))}
        _ = v35
    case 9223372036854775807:
        if false {
            var v36 int8 = v27
            _ = v36
        } else {
            for fi37 := 0; fi37 < 13; fi37++ {
                for fi38 := 0; fi38 < 19; fi38++ {
                    if (18446744073709551615 <= 4294967295) {
                        if (!false) {
                            var v39 int = (^(v33))
                            _ = v39
                        }
                        v3 = 2
                        v23 += v23
                    }
                    if false {
                        if false {
                            v40 := GenS3{F0: 1, F1: 0}
                            var v41 int16 = 714
                            _ = v40
                            _ = v41
                        }
                        v23 = (v23 % ((v23) | 1))
                    }
                    var v42 int64 = 1
                    _ = fi38
                    _ = v42
                }
                _ = fi37
            }
            if false {
                if (0 >= 975) {
                    var v43 int = genFn0(1, v23, 469)
                    if (968 != 1) {
                        v2 = (v2 + 3.14159)
                        var v44 uint32 = 65535
                        _ = v44
                    }
                    _ = v43
                } else {
                    fmt.Println(v33)
                    v45 := GenS1{F0: false, F1: 42}
                    delete(v1, "")
                    _ = v45
                }
                fmt.Println(v33)
            }
        }
    default:
        for fi46 := 0; fi46 < 7; fi46++ {
            var v47 int32 = (v23 | 1)
            _ = fi46
            _ = v47
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
}
