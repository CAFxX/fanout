package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 int8; F1 int16 }
type GenS9 struct { F0 uint32; F1 uint16; F2 string }
func genFn0(p0 int16, p1 int16, p2 int16) uint16 {
    if false {
        p2 += -1
    }
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            fmt.Println(p2, fi4)
            _ = fi5
        }
        _ = fi4
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint = 0
        v8 := GenS1{F0: 206, F1: 2}
        var v9 int64 = 9223372036854775807
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            if (992 > 2) {
                if (fi11 != (fi10 - (fi10 + (fi10 << ((432) & 63))))) {
                    var v12 int64 = 1
                    _ = v12
                }
            } else {
                v13 := GenS1{F0: 185, F1: 2}
                v14 := GenS1{F0: 0, F1: 455}
                _ = v13
                _ = v14
            }
            v15 := GenS1{F0: 1, F1: 61}
            p1 += (p0 + -32768)
            _ = fi11
            _ = v15
        }
        p0 += (p1 & (p2 / ((p0) | 1)))
        var v16 uint8 = 2
        _ = fi10
        _ = v16
    }
    var v17 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 1
}
func genFn1(p0 int8) int {
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := GenS2{F0: "a", F1: "hello"}
            if true {
                p0 += p0
                var v21 bool = (1e10 <= ((0.0 * 1.5) + 1e10))
                switch (fi18 - fi19) {
                case 9223372036854775807:
                    var v22 float32 = 1.5
                    _ = v22
                case -1:
                    fmt.Println(p0, fi18, v21)
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        v23 := []int{(fi19 | (fi19 - (fi18 * (fi19 % (((fi19 ^ 42)) | 1))))), fi19, 4294967295, -218}
                        switch (fi18 / ((fi18) | 1)) {
                        case 128:
                            var v24 uint16 = 2
                            _ = v24
                        default:
                            v25 := GenS3{F0: 0, F1: fi19}
                            _ = v25
                        }
                        _ = v23
                    }
                    fmt.Println(fi19, fi19, p0)
                }
                var v26 int = (fi18 | -1)
                fmt.Println(fi19)
                var v27 string = (("hello" + "") + "a")
                _ = v21
                _ = v26
                _ = v27
            }
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    var v28 uint8 = 0
    v29 := []int{1, 32767}
    switch -63 {
    case -1:
        var v30 int64 = 12
        _ = v30
    case 0:
        for fi31 := 0; fi31 < 10; fi31++ {
            for fi32 := 0; fi32 < 14; fi32++ {
                v29 = append(v29, (fi32 + fi31))
                fmt.Println(fi32, v28, p0)
                _ = fi32
            }
            _ = fi31
        }
    case 120:
        v33 := GenS4{F0: 2}
        _ = v33
    }
    var v34 bool = false
    _ = p0
    _ = v28
    _ = v29
    _ = v34
    return -9223372036854775808
}

func main() {
    var v1 uint64 = 797
    var v2 uint32 = 4294967295
    var v3 bool = (123 != -20)
    v35 := []int{-2147483648, genFn1(-62)}
    v2 = 750
    fmt.Println(v3, v2)
    v2 = v2
    var v36 bool = (!v3)
    for fi37 := 0; fi37 < 5; fi37++ {
        if (1.5 <= 1.5) {
            _ = len(v35)
        } else {
            for fi38 := 0; fi38 < 9; fi38++ {
                if (true || (false && v3)) {
                    var v39 float32 = float32(((3.14159 / 0.1) - float64(1e10)))
                    v40 := GenS5{F0: 261}
                    switch (fi37 >> ((42) & 63)) {
                    case 1:
                        if (v2 > (v2 | v2)) {
                            var v41 int64 = 642
                            v39 = v39
                            var v42 float64 = (0.0 / (-2.25 / 1.0))
                            _ = v41
                            _ = v42
                        } else {
                            _ = len(v35)
                            v1 = 18446744073709551615
                            if (-128 <= 0) {
                                v43 := GenS1{F0: 2, F1: v2}
                                v36 = (61 == 1)
                                v35 = append(v35, int(v1))
                                _ = v43
                            } else {
                                fmt.Println(fi38)
                            }
                        }
                    case 9223372036854775807:
                        v44 := GenS6{F0: 32767}
                        _ = v44
                    default:
                        var v45 int = (fi38 & (fi37 & (^(fi37))))
                        _ = v45
                    }
                    _ = v39
                    _ = v40
                }
                _ = len(v35)
                fmt.Println(fi38)
                _ = fi38
            }
            for fi46 := 0; fi46 < 14; fi46++ {
                var v47 int32 = 956
                fmt.Println(v1, v36, v1)
                _ = fi46
                _ = v47
            }
            v1 = (v1 * v1)
        }
        var v48 int = (^(fi37))
        for fi49 := 0; fi49 < 11; fi49++ {
            if false {
                v50 := GenS7{F0: 867, F1: (v2 - v2)}
                _ = v50
            }
            v36 = ((^(v1)) > v1)
            v51 := GenS3{F0: 1, F1: 2147483647}
            _ = fi49
            _ = v51
        }
        _ = fi37
        _ = v48
    }
    fmt.Println(v1)
    var v52 float64 = (3.14159 / -2.25)
    v35 = append(v35, 2147483647)
    var v53 int = 2147483647
    if v3 {
        for fi54 := 0; fi54 < 15; fi54++ {
            _ = len(v35)
            v35 = append(v35, fi54)
            _ = fi54
        }
        v35 = append(v35, 1)
        var v55 int32 = -2147483648
        _ = v55
    } else {
        for fi56 := 0; fi56 < 15; fi56++ {
            var v57 int = genFn1(-1)
            for fi58 := 0; fi58 < 16; fi58++ {
                if true {
                    var v59 int8 = -128
                    fmt.Println(v59)
                    fmt.Println(fi56, v59)
                    _ = v59
                }
                fmt.Println(v2, v1)
                _ = fi58
            }
            v60 := GenS8{F0: 7, F1: 1}
            _ = fi56
            _ = v57
            _ = v60
        }
        v61 := GenS9{F0: (^(v2)), F1: uint16(v2), F2: ""}
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v52
    _ = v53
}
