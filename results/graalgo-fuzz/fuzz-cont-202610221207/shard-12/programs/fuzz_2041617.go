package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 bool; F1 int16; F2 int32 }
type GenS3 struct { F0 int16; F1 uint64 }
func genFn0(p0 bool, p1 string, p2 int32) int8 {
    var v7 string = p1
    v8 := []int{9223372036854775807}
    var v9 float64 = 1e10
    fmt.Println(v7, v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return int8(p2)
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]string)
    var v3 float64 = (0.1 / 3.14159)
    v4 := make(map[string]string)
    var v5 uint = 2
    var v6 int8 = 97
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 bool = false
        _ = fi10
        _ = v11
    }
    delete(v2, 32767)
    if false {
        var v12 int32 = -2147483648
        fmt.Println(v12, v5)
        _ = v12
    }
    var v13 bool = (-1 <= -191)
    switch 32768 {
    case 32767:
        var v14 uint = (v5 / ((v5) | 1))
        _ = v14
    default:
        fmt.Println(v13, v1)
    }
    if (v5 != (v5 & v5)) {
        v13 = (!(v13 || true))
        switch 127 {
        case 554:
            switch "foo bar" {
            case "foo bar":
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 uint8 = 13
                    v17 := GenS1{F0: (v1 & v1), F1: v13}
                    v16 = 177
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
        default:
            for fi18 := 0; fi18 < 13; fi18++ {
                switch 128 {
                case 761:
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v20 := []int{(-(fi18)), fi18, (fi18 - (fi18 / ((159) | 1))), 4294967295}
                        _ = fi19
                        _ = v20
                    }
                default:
                    fmt.Println(fi18, v5, v13)
                }
                _ = fi18
            }
        }
        v21 := []int{1, -490}
        _ = v21
    } else {
        v4[""] = ("hello" + "")
    }
    v3 = v3
    for fi22 := 0; fi22 < 18; fi22++ {
        switch fi22 {
        case 815:
            v23 := GenS2{F0: v13, F1: 1000, F2: -2147483648}
            _ = v23
        case 32767:
            fmt.Println(fi22, v1)
        case 9223372036854775807:
            v13 = (620 >= (v5 / ((v5) | 1)))
        }
        v3 = v3
        fmt.Println(v5, v1, v13)
        _ = fi22
    }
    var v24 int = 589
    var v25 float32 = (0.1 * -2.25)
    v25 += float32(0.1)
    v26 := v4["foo bar"]
    fmt.Println(v26, v5)
    for fi27 := 0; fi27 < 8; fi27++ {
        switch "\n\t" {
        case "a":
            var v28 bool = (v13 || ((v6 <= v6) && true))
            _ = v28
        }
        v29 := []int{735}
        var v30 bool = (!(!(-972 != 1)))
        _ = fi27
        _ = v29
        _ = v30
    }
    if (998 == 449) {
        for fi31 := 0; fi31 < 7; fi31++ {
            for fi32 := 0; fi32 < 12; fi32++ {
                switch (^(v24)) {
                case 376:
                    if (42 == 65535) {
                        if (v1 >= v1) {
                            if (v5 != 966) {
                                v33 := GenS2{F0: ((v26 == v26) && v13), F1: -690, F2: 42}
                                v13 = true
                                _ = v33
                            }
                            switch v26 {
                            case "":
                                v34 := []int{fi31, fi31, fi32}
                                _ = v34
                            }
                        } else {
                            v24 = v24
                            v35 := []int{(fi32 * (fi31 + fi31)), 42, -1}
                            fmt.Println(v3, v3, v24)
                            switch (v26 + v26) {
                            case "\n\t":
                                v6 -= v6
                            default:
                                var v36 uint64 = 42
                                _ = v36
                            }
                            fmt.Println(fi31, v26)
                            _ = v35
                        }
                        v1 = v1
                    } else {
                        if true {
                            v37 := GenS3{F0: 32767, F1: 1}
                            _ = v37
                        } else {
                            v1 = 2
                            v38 := GenS1{F0: (v1 << ((2) & 63)), F1: v13}
                            v26 = v26
                            if true {
                                v39 := GenS3{F0: 0, F1: 1}
                                _ = v39
                            } else {
                                v13 = (65535 > 0)
                            }
                            v24 -= (fi32 * (fi31 + fi32))
                            _ = v38
                        }
                        switch v26 {
                        case "a":
                            if v13 {
                                fmt.Println(fi31)
                                v40 := []int{}
                                _ = v40
                            }
                        }
                    }
                case 42:
                    v5 = v5
                }
                _ = fi32
            }
            v1 = (v1 + 157)
            _ = fi31
        }
        v41 := GenS3{F0: 984, F1: 0}
        v25 -= 1e10
        _ = v41
    } else {
        v1 = (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
