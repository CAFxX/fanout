package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint16; F1 string; F2 int }
func genFn0(p0 uint32, p1 string, p2 uint) int32 {
    switch 4294967295 {
    case 4294967295:
        var v4 uint = p2
        _ = v4
    case 938:
        var v5 uint8 = 128
        _ = v5
    default:
        if (true && (true && true)) {
            for fi6 := 0; fi6 < 10; fi6++ {
                fmt.Println(p0)
                _ = fi6
            }
        }
    }
    v7 := []int{923, 734}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1() uint8 {
    var v8 float64 = (1.5 + float64((1.5 + (1e10 / 1e10))))
    v8 = v8
    var v9 int64 = 0
    if (-32768 != 42) {
        v10 := GenS1{F0: 744, F1: 1, F2: 4294967295}
        var v11 uint = 255
        _ = v10
        _ = v11
    } else {
        if ((-32768 < -956) && false) {
            v12 := GenS1{F0: 65535, F1: 4294967295, F2: 2}
            v13 := []int{}
            v14 := GenS1{F0: 1, F1: 220, F2: 4294967295}
            _ = v12
            _ = v13
            _ = v14
        } else {
            v15 := []int{-206, 127, -32, -1}
            _ = v15
        }
        if true {
            v16 := GenS2{F0: true}
            v8 = (-2.25 / ((v8 + -2.25) * float64(1e10)))
            v17 := []int{42}
            _ = v16
            _ = v17
        } else {
            v8 -= 3.14159
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS2{F0: false}
                switch ("" + "\n\t") {
                case "hello":
                    var v20 uint8 = 1
                    _ = v20
                default:
                    if (true && false) {
                        var v21 uint8 = 11
                        v22 := []int{}
                        _ = v21
                        _ = v22
                    } else {
                        v23 := []int{}
                        fmt.Println(v9, v8)
                        for fi24 := 0; fi24 < 12; fi24++ {
                            v8 = (0.1 + v8)
                            _ = fi24
                        }
                        _ = v23
                    }
                }
                if true {
                    switch fi18 {
                    case -234:
                        if false {
                            switch "hello" {
                            case "a":
                                v9 -= (v9 & (v9 % (((v9 * v9)) | 1)))
                            }
                            var v25 string = "a"
                            _ = v25
                        } else {
                            var v26 uint = 372
                            _ = v26
                        }
                    }
                } else {
                    var v27 float64 = float64(0.0)
                    v8 = v27
                    v28 := GenS3{F0: 1}
                    _ = v27
                    _ = v28
                }
                _ = fi18
                _ = v19
            }
        }
    }
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint16 = 0
    v2 := []uint32{722, 341}
    v3 := []string{"a", "a"}
    var v29 int32 = 1
    v1 = 2
    v30 := GenS3{F0: 1}
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := GenS1{F0: v1, F1: 648, F2: 315}
        var v33 string = "foo bar"
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 bool = (false || (127 == -36))
    v35 := GenS4{F0: 42, F1: v34}
    var v36 uint16 = v1
    switch -1 {
    case 4294967295:
        if (0 != -1) {
            var v37 uint8 = 127
            if v34 {
                for fi38 := 0; fi38 < 9; fi38++ {
                    for fi39 := 0; fi39 < 12; fi39++ {
                        var v40 uint32 = 57
                        fmt.Println(v29, v34, v37)
                        _ = fi39
                        _ = v40
                    }
                    v29 = 2147483647
                    _ = fi38
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v41 int8 = 42
                _ = v41
            }
            _ = v37
        } else {
            var v42 string = "foo bar"
            var v43 float64 = 0.0
            _ = v42
            _ = v43
        }
    }
    v44 := GenS2{F0: (v34 && (2 <= 969))}
    if len(v2) > 0 {
        v2[0] = 65535
    }
    switch 947 {
    case -964:
        var v45 int32 = (v29 >> ((898) & 63))
        _ = v45
    case 32768:
        for fi46 := 0; fi46 < 7; fi46++ {
            var v47 int8 = 0
            v48 := GenS2{F0: (42 <= 4294967295)}
            _ = fi46
            _ = v47
            _ = v48
        }
    case 1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v49 uint32 = 36
    fmt.Println(v34, v1)
    if v34 {
        var v50 uint64 = 390
        for fi51 := 0; fi51 < 17; fi51++ {
            v36 = (v1 & v1)
            _ = fi51
        }
        _ = v50
    }
    v49 += 0
    for fi52 := 0; fi52 < 10; fi52++ {
        v53 := GenS4{F0: v49, F1: v34}
        fmt.Println(v1)
        var v54 bool = v34
        _ = fi52
        _ = v53
        _ = v54
    }
    v55 := GenS5{F0: (v36 ^ v36), F1: "foo bar", F2: 2147483647}
    var v56 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v49:", v49)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v44
    _ = v49
    _ = v55
    _ = v56
}
