package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 int32; F1 uint }
func genFn0(p0 string) int64 {
    p0 = p0
    if (!false) {
        v5 := GenS1{F0: 1000}
        v6 := []int{}
        _ = v5
        _ = v6
    } else {
        if (14 > -85) {
            var v7 uint16 = 108
            fmt.Println(p0, p0)
            v7 += 1
            _ = v7
        } else {
            v8 := GenS2{F0: 65535}
            switch 4294967295 {
            case 786:
                if (!(!true)) {
                    for fi9 := 0; fi9 < 4; fi9++ {
                        var v10 uint16 = 42
                        _ = fi9
                        _ = v10
                    }
                } else {
                    p0 = (p0 + p0)
                    for fi11 := 0; fi11 < 17; fi11++ {
                        fmt.Println(fi11, fi11)
                        _ = fi11
                    }
                    for fi12 := 0; fi12 < 11; fi12++ {
                        for fi13 := 0; fi13 < 20; fi13++ {
                            if true {
                                var v14 int = (-(fi12))
                                _ = v14
                            }
                            _ = fi13
                        }
                        v15 := []int{(fi12 * (fi12 / ((fi12) | 1))), (-(fi12)), fi12}
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = fi12
                        _ = v15
                    }
                }
            case 32768:
                if (127 == 42) {
                    p0 = "a"
                    v16 := []int{}
                    _ = len(v16)
                    _ = v16
                } else {
                    if true {
                        var v17 uint32 = 0
                        v17 = v17
                        if false {
                            v18 := GenS1{F0: 1000}
                            p0 = (p0 + p0)
                            _ = v18
                        }
                        _ = v17
                    } else {
                        fmt.Println(p0)
                        for fi19 := 0; fi19 < 6; fi19++ {
                            v20 := GenS1{F0: 739}
                            _ = fi19
                            _ = v20
                        }
                    }
                    var v21 uint = 450
                    fmt.Println(p0)
                    _ = v21
                }
            default:
                if (4294967295 == 2147483647) {
                    if (128 > 1) {
                        fmt.Println(p0)
                        fmt.Println(p0)
                        v22 := GenS3{F0: 35, F1: -1}
                        _ = v22
                    }
                } else {
                    var v23 uint8 = 69
                    var v24 bool = (0.0 != (1.5 / 1.5))
                    _ = v23
                    _ = v24
                }
            }
            _ = v8
        }
        var v25 uint64 = 2
        _ = v25
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        fmt.Println(p0, p0)
        for fi27 := 0; fi27 < 20; fi27++ {
            fmt.Println(fi27, p0)
            _ = fi27
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            v29 := []int{42}
            fmt.Println(fi26, p0, p0)
            p0 = "a"
            _ = fi28
            _ = v29
        }
        _ = fi26
    }
    v30 := GenS2{F0: 455}
    switch 128 {
    case 1:
        fmt.Println(p0)
    case 9223372036854775807:
        v31 := GenS3{F0: 34, F1: -867}
        _ = v31
    }
    _ = p0
    _ = v30
    return -695
}

func main() {
    var v1 int64 = -1
    var v2 bool = true
    v3 := []int{32767, -331, 117}
    v4 := make(map[uint32]string)
    var v32 int64 = 9223372036854775807
    if (4294967295 >= 4294967295) {
        var v33 bool = v2
        if true {
            v34 := GenS4{F0: 2, F1: 1}
            _ = v34
        } else {
            v35 := v4[2147483648]
            v33 = v2
            for fi36 := 0; fi36 < 2; fi36++ {
                if (0 < fi36) {
                    var v37 int = fi36
                    var v38 int16 = -1
                    _ = v37
                    _ = v38
                }
                _ = fi36
            }
            _ = v35
        }
        delete(v4, 1)
        _ = v33
    } else {
        switch "foo bar" {
        case "\n\t":
            for fi39 := 0; fi39 < 2; fi39++ {
                v40 := GenS5{F0: (fi39 * 127), F1: 1}
                v1 += v32
                _ = fi39
                _ = v40
            }
        case "":
            v2 = true
        case "hello":
            if false {
                fmt.Println(v2, v1, v1)
                switch -15 {
                case 32768:
                    var v41 float32 = (1.5 / 1.5)
                    _ = v41
                case -332:
                    v2 = (0 >= 127)
                case 0:
                    v2 = (!(!v2))
                default:
                    if v2 {
                        v2 = true
                    }
                }
            } else {
                v1 += (-(v1))
                for fi42 := 0; fi42 < 3; fi42++ {
                    v43 := GenS5{F0: fi42, F1: -2147483648}
                    var v44 int64 = 9223372036854775807
                    v3 = append(v3, fi42)
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            }
        }
        var v45 int64 = (v32 + v32)
        _ = v45
    }
    v46 := v4[930]
    delete(v4, 141)
    fmt.Println(v32)
    for fi47 := 0; fi47 < 3; fi47++ {
        v48 := GenS6{F0: -757, F1: 57}
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 8; fi49++ {
        v32 = genFn0(v46)
        v50 := GenS2{F0: 65535}
        if len(v3) > 0 {
            v3[0] = (fi49 >> ((213) & 63))
        }
        _ = fi49
        _ = v50
    }
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v46
}
