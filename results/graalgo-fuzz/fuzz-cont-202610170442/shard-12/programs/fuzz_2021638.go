package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int; F1 uint; F2 bool }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint32; F1 string; F2 int }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 uint }
func genFn0(p0 int32, p1 string) int16 {
    v5 := []int{32767, 127, 0, -316}
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = v5
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := []int{}
    var v4 uint = 2
    switch -955 {
    case -2147483648:
        for fi6 := 0; fi6 < 18; fi6++ {
            v4 = v4
            v7 := GenS1{F0: fi6, F1: fi6}
            v8 := GenS2{F0: (^(fi6)), F1: 4294967295, F2: (9223372036854775807 < -1)}
            _ = fi6
            _ = v7
            _ = v8
        }
    case -1:
        v4 -= v4
    case 32768:
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 string = "a"
            var v11 int = fi9
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 bool = true
                fmt.Println(fi12, v10)
                v13 = (fi12 > 0)
                _ = fi12
                _ = v13
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        for fi14 := 0; fi14 < 16; fi14++ {
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 uint16 = 255
                _ = fi15
                _ = v16
            }
            var v17 int32 = 0
            fmt.Println(fi14, v17)
            _ = fi14
            _ = v17
        }
    }
    v18 := GenS2{F0: 790, F1: v4, F2: true}
    var v19 string = "a"
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 int16 = -1
        _ = fi20
        _ = v21
    }
    var v22 int16 = 32767
    var v23 string = "foo bar"
    var v24 uint64 = 955
    v25 := GenS3{F0: 4294967295, F1: v19}
    var v26 int64 = -277
    v22 -= int16(v26)
    _ = len(v2)
    if true {
        var v27 int16 = (v22 & (-(v22)))
        if false {
            if false {
                if (float64(0.0) > (1.5 + ((3.14159 * 3.14159) / (0.1 + (0.0 - 1e10))))) {
                    v26 += v26
                    for fi28 := 0; fi28 < 11; fi28++ {
                        v3 = append(v3, fi28)
                        v4 = (^(v4))
                        var v29 uint16 = 42
                        _ = fi28
                        _ = v29
                    }
                    v2[1] = 1
                }
            } else {
                for fi30 := 0; fi30 < 4; fi30++ {
                    v4 = v4
                    _ = fi30
                }
                for fi31 := 0; fi31 < 1; fi31++ {
                    delete(v2, fi31)
                    _ = fi31
                }
                var v32 int32 = -2147483648
                _ = v32
            }
        }
        v33 := GenS2{F0: -383, F1: (v4 + (v4 << ((42) & 63))), F2: false}
        _ = v27
        _ = v33
    } else {
        if (!false) {
            if false {
                v34 := v1[int(v26)]
                v23 = (v23 + (v19 + v23))
                _ = v34
            }
            if false {
                if true {
                    var v35 uint16 = 65535
                    _ = v35
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v36 uint = 65535
                _ = v36
            } else {
                for fi37 := 0; fi37 < 9; fi37++ {
                    if false {
                        switch "hello" {
                        case "a":
                            if true {
                                if true {
                                    var v38 int64 = 1
                                    _ = v38
                                }
                            } else {
                                v3 = append(v3, 2147483647)
                            }
                        case "foo bar":
                            v39 := GenS4{F0: -627, F1: (0.0 >= 0.0)}
                            _ = v39
                        case "hello":
                            var v40 uint64 = v24
                            _ = v40
                        }
                    }
                    v41 := GenS5{F0: 147, F1: v19, F2: 32767}
                    if true {
                        var v42 float32 = 1.5
                        _ = len(v2)
                        _ = v42
                    }
                    _ = fi37
                    _ = v41
                }
                for fi43 := 0; fi43 < 5; fi43++ {
                    v24 -= (v24 | v24)
                    switch 922 {
                    case -2147483648:
                        v44 := GenS6{F0: (v26 & -1), F1: 72}
                        _ = v44
                    case -9223372036854775808:
                        if true {
                            v45 := GenS6{F0: (-(v26)), F1: 142}
                            _ = v45
                        }
                    }
                    fmt.Println(fi43, fi43, fi43)
                    _ = fi43
                }
                v46 := GenS3{F0: (v4 - (v4 & v4)), F1: v19}
                _ = v46
            }
        } else {
            v19 = ""
            for fi47 := 0; fi47 < 4; fi47++ {
                if (42 >= v24) {
                    v48 := GenS3{F0: v4, F1: (v19 + "a")}
                    v49 := GenS6{F0: v26, F1: 54}
                    _ = v48
                    _ = v49
                }
                _ = fi47
            }
            for fi50 := 0; fi50 < 12; fi50++ {
                _ = len(v3)
                fmt.Println(fi50, v22, v24)
                _ = len(v2)
                _ = fi50
            }
        }
    }
    if (false || (v4 != v4)) {
        v19 = ""
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v51 int8 = -1
        delete(v2, 127)
        var v52 int = -364
        _ = v51
        _ = v52
    }
    v53 := GenS7{F0: 0}
    switch 240 {
    case -268:
        v22 = v22
    }
    v26 = (v26 << (((v4 & v4)) & 63))
    var v54 int16 = (^(v22))
    delete(v1, 157)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v53
    _ = v54
}
