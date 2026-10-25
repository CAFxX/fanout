package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int32; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint }
type GenS5 struct { F0 bool; F1 int16 }
func genFn0(p0 int32) int64 {
    v5 := []int{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v6 int64 = 42
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (-9223372036854775808 > 711) {
        v7 := GenS1{F0: 0}
        v6 = 1
        _ = v7
    }
    _ = p0
    _ = v5
    _ = v6
    return v6
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 uint32 = 42
    var v4 int32 = -203
    if (!true) {
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := []int{fi8}
            v3 = v3
            for fi10 := 0; fi10 < 2; fi10++ {
                switch 128 {
                case -1:
                    v11 := GenS1{F0: (fi10 & fi10)}
                    _ = v11
                default:
                    if (2 >= 0) {
                        v12 := GenS1{F0: 4294967295}
                        v3 = 488
                        var v13 string = ("a" + "a")
                        _ = v12
                        _ = v13
                    }
                }
                switch 32768 {
                case -954:
                    v4 -= (v4 >> ((657) & 63))
                    v3 = 803
                case -2147483648:
                    v3 -= 4294967295
                case -356:
                    switch "hello" {
                    case "a":
                        v14 := GenS2{F0: 65535, F1: (-(v4)), F2: 71}
                        _ = v14
                    case "hello":
                        v15 := GenS2{F0: 65535, F1: (v4 ^ (-(v4))), F2: 1}
                        _ = v15
                    }
                }
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
    }
    v4 += 313
    v3 += 0
    if (1000 > -934) {
        for fi16 := 0; fi16 < 9; fi16++ {
            _ = len(v1)
            var v17 int8 = -128
            _ = fi16
            _ = v17
        }
    } else {
        if true {
            switch "" {
            case "a":
                v18 := v2["foo bar"]
                _ = v18
            case "\n\t":
                fmt.Println(v3, v3)
            case "foo bar":
                if true {
                    v19 := []int{9223372036854775807}
                    v20 := GenS2{F0: 517, F1: (^(v4)), F2: 1}
                    if true {
                        v3 += v3
                        v3 += (v3 >> ((2) & 63))
                        for fi21 := 0; fi21 < 2; fi21++ {
                            for fi22 := 0; fi22 < 8; fi22++ {
                                v4 += 1
                                _ = fi22
                            }
                            if ((v3 * v3) != v3) {
                                if ((294 > 29) || (32767 > 32767)) {
                                    for fi23 := 0; fi23 < 8; fi23++ {
                                        if (!(float64((float32(3.14159) + 1e10)) != (float64(3.14159) + 1.5))) {
                                            v3 = (v3 ^ (v3 >> ((405) & 63)))
                                        } else {
                                            switch ((("foo bar" + "foo bar") + "") + ("\n\t" + ("a" + "a"))) {
                                            case "a":
                                                switch ("hello" + "foo bar") {
                                                case "a":
                                                    v3 -= (^(v3))
                                                }
                                            case "hello":
                                                v24 := GenS2{F0: 255, F1: (^(v4)), F2: 0}
                                                _ = v24
                                            default:
                                                fmt.Println(v4, v3, v4)
                                            }
                                        }
                                        switch (fi21 * (fi23 + (fi21 << ((1) & 63)))) {
                                        case 1:
                                            v3 = v3
                                        }
                                        _ = fi23
                                    }
                                    for fi25 := 0; fi25 < 1; fi25++ {
                                        v26 := GenS3{F0: 415, F1: 65535, F2: 255}
                                        delete(v1, fi21)
                                        if true {
                                            v4 += -760
                                        }
                                        _ = fi25
                                        _ = v26
                                    }
                                    delete(v1, 42)
                                } else {
                                    v4 += v4
                                    fmt.Println(v4, v4)
                                }
                            } else {
                                var v27 int32 = v4
                                switch (fi21 + fi21) {
                                case 685:
                                    if false {
                                        if false {
                                            fmt.Println(v4)
                                            fmt.Println(fi21, v4)
                                            for fi28 := 0; fi28 < 2; fi28++ {
                                                v4 = 522
                                                _ = fi28
                                            }
                                        }
                                        delete(v2, "hello")
                                    } else {
                                        if false {
                                            var v29 uint64 = 4294967295
                                            v3 = 65535
                                            _ = v29
                                        } else {
                                            if (!(!(("a" + ("foo bar" + ("" + "foo bar"))) != ("foo bar" + "a")))) {
                                                if (true || (-2.25 <= 1.5)) {
                                                    if true {
                                                        v30 := GenS2{F0: 1, F1: 0, F2: 1}
                                                        v31 := GenS4{F0: 255, F1: 42, F2: 42}
                                                        var v32 float64 = (float64(-2.25) / 1e10)
                                                        _ = v30
                                                        _ = v31
                                                        _ = v32
                                                    } else {
                                                        v33 := GenS3{F0: 255, F1: 167, F2: 0}
                                                        var v34 uint16 = 1
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    fmt.Println(v3)
                                                    if len(v19) >= 2 {
                                                        v19 = v19[1:len(v19)]
                                                    }
                                                }
                                                v27 = (v27 * (v27 / (((v27 >> ((791) & 63))) | 1)))
                                                v3 = v3
                                            } else {
                                                var v35 int16 = 32767
                                                _ = v35
                                            }
                                            var v36 int8 = -1
                                            _ = v36
                                        }
                                    }
                                case -9223372036854775808:
                                    fmt.Println(v4, v3, v27)
                                case 240:
                                    var v37 int8 = 29
                                    _ = v37
                                default:
                                    for fi38 := 0; fi38 < 17; fi38++ {
                                        v39 := GenS2{F0: 811, F1: (v27 * (^(v27))), F2: genFn0(v27)}
                                        _ = fi38
                                        _ = v39
                                    }
                                }
                                _ = v27
                            }
                            _ = fi21
                        }
                    }
                    _ = v19
                    _ = v20
                }
            default:
                for fi40 := 0; fi40 < 15; fi40++ {
                    v4 = v4
                    _ = fi40
                }
            }
        }
        fmt.Println(v3, v3)
    }
    for fi41 := 0; fi41 < 14; fi41++ {
        if true {
            var v42 uint64 = 18446744073709551615
            _ = v42
        }
        _ = fi41
    }
    for fi43 := 0; fi43 < 3; fi43++ {
        fmt.Println(fi43)
        _ = fi43
    }
    v44 := v1[-1]
    var v45 uint = 4294967295
    v46 := GenS5{F0: true, F1: -1}
    for fi47 := 0; fi47 < 4; fi47++ {
        _ = len(v1)
        _ = fi47
    }
    delete(v1, (^(v44)))
    var v48 int32 = (-(v4))
    var v49 float32 = (float32(0.1) / 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v48
    _ = v49
}
