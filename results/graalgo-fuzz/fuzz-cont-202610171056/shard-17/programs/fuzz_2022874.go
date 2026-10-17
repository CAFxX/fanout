package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 string; F2 uint32 }
type GenS3 struct { F0 int; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16; F1 bool; F2 uint }
type GenS7 struct { F0 int32; F1 int32 }
func genFn0() bool {
    v7 := []int{-198, 4294967295, 4294967295}
    var v8 uint = 369
    v8 = (v8 & v8)
    v8 = (v8 & (^(v8)))
    if len(v7) > 0 {
        v7[0] = -599
    }
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 int8, p1 int32) bool {
    v9 := []int{127, 128, 4294967295}
    switch 128 {
    case -1:
        var v10 int = -53
        _ = v10
    }
    var v11 float64 = -2.25
    fmt.Println(v11, p0, v11)
    _ = p0
    _ = p1
    _ = v11
    _ = v9
    return false
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]int)
    v3 := []int{127, -775, -410, 1}
    var v4 uint64 = 42
    var v5 int32 = -1
    var v6 bool = ("\n\t" <= ("hello" + ""))
    var v12 string = ((("a" + "\n\t") + "a") + ("\n\t" + "a"))
    var v13 uint32 = 326
    switch 42 {
    case 42:
        v14 := GenS1{F0: (v5 & -581)}
        _ = v14
    case 0:
        v13 = (^(v13))
    case 127:
        v15 := v2[(^(v13))]
        _ = v15
    default:
        v16 := GenS1{F0: (^(v5))}
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS2{F0: 42, F1: "hello", F2: v13}
        _ = fi17
        _ = v18
    }
    var v19 int = -994
    v20 := GenS1{F0: v5}
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS3{F0: (fi21 - -2147483648), F1: v5, F2: 42}
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: -88, F1: v12, F2: (v13 | 965)}
    for fi24 := 0; fi24 < 12; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 uint8 = 255
            v27 := GenS4{F0: 1, F1: 1}
            v5 = -2147483648
            _ = fi25
            _ = v26
            _ = v27
        }
        if (v5 == v5) {
            if ((!(v13 < v13)) || false) {
                if v6 {
                    delete(v2, v13)
                } else {
                    v13 = v13
                    var v28 uint32 = v13
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 5; fi29++ {
                switch -89 {
                case 128:
                    var v30 bool = true
                    fmt.Println(v30, fi29)
                    _ = v30
                case 127:
                    v19 = fi24
                case 2147483647:
                    if ((!true) || (-1 < -1)) {
                        v4 = (v4 | 42)
                    } else {
                        var v31 int8 = -103
                        v32 := GenS3{F0: 1, F1: v5, F2: 128}
                        _ = v31
                        _ = v32
                    }
                }
                _ = fi29
            }
            var v33 int32 = (v5 % (((v5 << ((65535) & 63))) | 1))
            _ = v33
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        for fi34 := 0; fi34 < 11; fi34++ {
            v6 = v6
            v13 -= 572
            v3 = append(v3, -1)
            _ = fi34
        }
        _ = fi24
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        var v36 int8 = -128
        var v37 int8 = (-(v36))
        _ = fi35
        _ = v36
        _ = v37
    }
    if false {
        if false {
            var v38 int16 = 589
            for fi39 := 0; fi39 < 10; fi39++ {
                switch v19 {
                case 4294967295:
                    v40 := GenS5{F0: 0}
                    _ = v40
                case 32768:
                    v4 -= 688
                case -240:
                    var v41 float64 = ((3.14159 - 1.5) - 1e10)
                    _ = v41
                default:
                    var v42 bool = (!true)
                    _ = v42
                }
                v12 = "a"
                _ = fi39
            }
            _ = v38
        } else {
            if true {
                for fi43 := 0; fi43 < 15; fi43++ {
                    v44 := GenS4{F0: 128, F1: -1}
                    v19 -= fi43
                    fmt.Println(v5)
                    _ = fi43
                    _ = v44
                }
            } else {
                v13 = 0
                for fi45 := 0; fi45 < 8; fi45++ {
                    v46 := GenS6{F0: (-(v1)), F1: v6, F2: 788}
                    if v6 {
                        for fi47 := 0; fi47 < 20; fi47++ {
                            if (v1 >= 0) {
                                fmt.Println(fi45, v1, v12)
                                fmt.Println(v6, v5)
                            }
                            switch (v19 << ((255) & 63)) {
                            case -119:
                                var v48 string = v12
                                _ = v48
                            case 32768:
                                var v49 int = (^(fi45))
                                _ = v49
                            case -1:
                                v13 = (v13 ^ (v13 / (((^(v13))) | 1)))
                            default:
                                if (3.14159 == 3.14159) {
                                    v6 = (v5 != -57)
                                    var v50 uint64 = 42
                                    _ = v50
                                }
                                var v51 uint32 = 4294967295
                                var v52 float64 = 1e10
                                _ = v51
                                _ = v52
                            }
                            var v53 int = fi45
                            _ = fi47
                            _ = v53
                        }
                        v54 := GenS7{F0: (-(v5)), F1: (^(v5))}
                        v55 := GenS2{F0: 100, F1: (((("" + "a") + (v12 + (v12 + "foo bar"))) + ("hello" + "")) + v12), F2: v13}
                        _ = v54
                        _ = v55
                    }
                    _ = fi45
                    _ = v46
                }
                v19 += (v19 << ((127) & 63))
            }
            v2[(v13 * v13)] = (v19 >> ((549) & 63))
        }
        for fi56 := 0; fi56 < 12; fi56++ {
            v4 = 336
            fmt.Println(v6)
            for fi57 := 0; fi57 < 3; fi57++ {
                if len(v3) > 0 {
                    v3[0] = (fi57 + (v19 & fi57))
                }
                fmt.Println(fi57, v6, fi56)
                _ = fi57
            }
            _ = fi56
        }
        switch int(v5) {
        case 127:
            v19 = v19
        }
    }
    v58 := GenS5{F0: -263}
    v6 = false
    var v59 uint16 = 32767
    fmt.Println(v59, v6, v5)
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
}
