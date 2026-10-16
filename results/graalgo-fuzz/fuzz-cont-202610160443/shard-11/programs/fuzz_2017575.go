package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int32; F2 uint8 }
type GenS3 struct { F0 int; F1 bool; F2 uint32 }
func genFn0(p0 int, p1 int32) int64 {
    var v6 int8 = -128
    var v7 int32 = (p1 >> ((255) & 63))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -636
}
func genFn1() uint8 {
    var v8 uint16 = 42
    v8 = v8
    _ = v8
    return 186
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"hello", "hello"}
    var v3 int64 = -1
    var v4 float64 = (1.5 * 3.14159)
    v5 := make(map[int]int)
    var v9 uint64 = 1
    switch 32768 {
    case 2147483647:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 1:
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int = fi10
            _ = fi10
            _ = v11
        }
    }
    v12 := GenS1{F0: (v9 * v9)}
    switch 0 {
    case 852:
        var v13 uint32 = 2
        _ = v13
    case 202:
        for fi14 := 0; fi14 < 10; fi14++ {
            if true {
                delete(v1, (fi14 / (((^(fi14))) | 1)))
                v15 := GenS1{F0: 260}
                _ = v15
            }
            v3 += -9223372036854775808
            _ = fi14
        }
    default:
        switch 32768 {
        case -9223372036854775808:
            for fi16 := 0; fi16 < 4; fi16++ {
                fmt.Println(v4)
                _ = fi16
            }
        case 42:
            v9 += 634
        case -95:
            v3 = (^(v3))
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        switch (fi17 - fi17) {
        case 108:
            var v18 float64 = float64((-2.25 - (float32(0.1) / float32(v4))))
            _ = v18
        case -1:
            for fi19 := 0; fi19 < 11; fi19++ {
                fmt.Println(fi17, v3)
                v4 += 0.1
                _ = fi19
            }
        case 0:
            if false {
                var v20 int8 = -128
                var v21 uint16 = 1
                var v22 uint8 = 231
                _ = v20
                _ = v21
                _ = v22
            }
        default:
            var v23 string = ("a" + ("hello" + "\n\t"))
            _ = v23
        }
        _ = fi17
    }
    switch -9223372036854775808 {
    case 4294967295:
        v24 := GenS2{F0: 618, F1: -1, F2: 1}
        _ = v24
    default:
        delete(v5, 114)
    }
    v25 := GenS1{F0: v9}
    v4 -= (((v4 - float64(float32(0.1))) / (v4 / v4)) * v4)
    switch 127 {
    case 42:
        var v26 int8 = 93
        _ = v26
    case 32768:
        if true {
            v27 := GenS2{F0: 42, F1: 2147483647, F2: 77}
            var v28 int32 = 1
            v9 = v9
            _ = v27
            _ = v28
        }
    default:
        var v29 int16 = 42
        _ = v29
    }
    var v30 uint16 = 43
    var v31 int16 = 42
    switch 1 {
    case 128:
        var v32 uint16 = (v30 | v30)
        _ = v32
    }
    switch 298 {
    case -2147483648:
        if (false || (-9223372036854775808 < 4294967295)) {
            var v33 int32 = -2147483648
            switch 32767 {
            case -788:
                var v34 float32 = 1e10
                _ = v34
            case 128:
                if (false && (v33 >= (v33 + 0))) {
                    for fi35 := 0; fi35 < 4; fi35++ {
                        for fi36 := 0; fi36 < 14; fi36++ {
                            var v37 bool = (!true)
                            var v38 int16 = v31
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        _ = fi35
                    }
                    v30 = v30
                    var v39 uint64 = 793
                    _ = v39
                } else {
                    if true {
                        v40 := GenS1{F0: v9}
                        for fi41 := 0; fi41 < 9; fi41++ {
                            v42 := GenS3{F0: fi41, F1: (!false), F2: 921}
                            _ = fi41
                            _ = v42
                        }
                        _ = v40
                    } else {
                        for fi43 := 0; fi43 < 5; fi43++ {
                            v44 := GenS3{F0: 0, F1: false, F2: 543}
                            var v45 uint8 = genFn1()
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                        if true {
                            v46 := v1[42]
                            var v47 uint64 = 4294967295
                            _ = v46
                            _ = v47
                        } else {
                            for fi48 := 0; fi48 < 18; fi48++ {
                                for fi49 := 0; fi49 < 11; fi49++ {
                                    var v50 uint16 = (v30 ^ v30)
                                    _ = fi49
                                    _ = v50
                                }
                                var v51 float32 = (-2.25 * 1.5)
                                _ = fi48
                                _ = v51
                            }
                            _ = len(v2)
                        }
                    }
                    for fi52 := 0; fi52 < 10; fi52++ {
                        v5[(fi52 % ((127) | 1))] = fi52
                        v53 := GenS1{F0: (v9 >> ((4294967295) & 63))}
                        _ = fi52
                        _ = v53
                    }
                }
            }
            _ = len(v2)
            _ = v33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
