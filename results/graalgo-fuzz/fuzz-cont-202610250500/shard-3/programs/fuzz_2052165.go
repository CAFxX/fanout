package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint }
func genFn0() uint64 {
    v7 := []int{32768}
    if true {
        var v8 float64 = 1.5
        fmt.Println(v8)
        _ = v8
    } else {
        var v9 uint64 = 450
        _ = v9
    }
    if false {
        var v10 int8 = 127
        for fi11 := 0; fi11 < 17; fi11++ {
            if false {
                if true {
                    v10 -= (v10 - (v10 << ((255) & 63)))
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    var v12 bool = (true && true)
                    _ = v12
                } else {
                    v7 = append(v7, (^(fi11)))
                }
            } else {
                v13 := GenS1{F0: 18446744073709551615, F1: true}
                v10 -= 0
                var v14 uint16 = 881
                _ = v13
                _ = v14
            }
            var v15 uint64 = 4294967295
            _ = fi11
            _ = v15
        }
        _ = v10
    } else {
        var v16 uint8 = 255
        _ = v16
    }
    switch "foo bar" {
    case "":
        if ((("hello" + ("" + "\n\t")) + ("a" + "foo bar")) < "") {
            var v17 int16 = 1
            _ = v17
        }
    case "a":
        if (!true) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        } else {
            if (((0 < -664) && (!true)) || true) {
                for fi18 := 0; fi18 < 17; fi18++ {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        var v20 float64 = 1e10
                        v20 = v20
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
                if false {
                    var v21 float32 = ((float32(float64(-2.25)) / (1.5 + 3.14159)) + 1e10)
                    _ = v21
                } else {
                    _ = len(v7)
                    if true {
                        var v22 bool = false
                        _ = v22
                    }
                    var v23 int8 = 127
                    _ = v23
                }
            } else {
                for fi24 := 0; fi24 < 19; fi24++ {
                    v25 := GenS2{F0: -1}
                    if false {
                        var v26 uint16 = 32767
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 3; fi27++ {
                            if false {
                                v28 := GenS1{F0: 69, F1: (-9223372036854775808 < 42)}
                                var v29 uint32 = 65535
                                _ = v28
                                _ = v29
                            }
                            _ = fi27
                        }
                    }
                    var v30 int32 = 1
                    _ = fi24
                    _ = v25
                    _ = v30
                }
                var v31 uint64 = 1
                var v32 int32 = 0
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                var v34 uint32 = 0
                _ = fi33
                _ = v34
            }
            if (!false) {
                var v35 int32 = 516
                var v36 string = "\n\t"
                if len(v7) > 0 {
                    v7[0] = -133
                }
                _ = v35
                _ = v36
            }
        }
    }
    _ = v7
    return 4294967295
}
func genFn1(p0 uint32) string {
    for fi37 := 0; fi37 < 15; fi37++ {
        v38 := GenS3{F0: 504, F1: 857}
        _ = fi37
        _ = v38
    }
    var v39 float32 = float32(-2.25)
    var v40 int64 = -9223372036854775808
    _ = p0
    _ = v39
    _ = v40
    return ("a" + "foo bar")
}

func main() {
    var v1 int16 = 216
    var v2 float64 = (0.0 + 1e10)
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 int8 = 42
    var v6 float64 = v2
    if false {
        v41 := GenS2{F0: (v5 << ((71) & 63))}
        _ = v41
    } else {
        var v42 int = 32767
        var v43 uint16 = 654
        for fi44 := 0; fi44 < 4; fi44++ {
            if ((0.0 / 1.0) == (0.1 / 3.14159)) {
                if true {
                    var v45 int16 = v1
                    var v46 uint8 = 0
                    for fi47 := 0; fi47 < 11; fi47++ {
                        v48 := []int{(^(fi47)), 387}
                        _ = fi47
                        _ = v48
                    }
                    _ = v45
                    _ = v46
                }
                v49 := []int{(v42 ^ -753), v42, 127, fi44}
                _ = v49
            } else {
                v1 -= 1
            }
            _ = fi44
        }
        _ = v42
        _ = v43
    }
    for fi50 := 0; fi50 < 5; fi50++ {
        var v51 uint64 = genFn0()
        if false {
            v52 := GenS2{F0: 1}
            _ = v52
        } else {
            v53 := GenS3{F0: int64(fi50), F1: 255}
            v1 = v1
            fmt.Println(v1, v51)
            _ = v53
        }
        _ = fi50
        _ = v51
    }
    v54 := GenS2{F0: (v5 * v5)}
    switch "a" {
    case "":
        for fi55 := 0; fi55 < 14; fi55++ {
            v2 -= v2
            v56 := []int{(fi55 * fi55), fi55, fi55, (fi55 / ((42) | 1))}
            _ = fi55
            _ = v56
        }
    }
    fmt.Println(v2, v1, v2)
    v57 := []int{0, -9223372036854775808, 911, -972}
    if len(v57) >= 2 {
        v57 = v57[1:len(v57)]
    }
    if (true || true) {
        var v58 float64 = v6
        var v59 uint = 2
        _ = v58
        _ = v59
    }
    if len(v57) > 0 {
        v57[0] = 32767
    }
    v2 = float64(1e10)
    var v60 int16 = (v1 ^ (v1 - 120))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v57
    _ = v6
    _ = v60
}
