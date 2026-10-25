package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint; F1 bool; F2 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8 }
func genFn0() int64 {
    v7 := GenS1{F0: 644, F1: 854, F2: 2}
    if true {
        v8 := []int{-15, -607}
        if true {
            for fi9 := 0; fi9 < 5; fi9++ {
                for fi10 := 0; fi10 < 2; fi10++ {
                    v11 := GenS1{F0: 4294967295, F1: 1, F2: 255}
                    _ = fi10
                    _ = v11
                }
                v12 := GenS2{F0: -104}
                _ = fi9
                _ = v12
            }
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 int32 = 0
                var v15 uint16 = 32767
                fmt.Println(v14, v14, fi13)
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            for fi16 := 0; fi16 < 3; fi16++ {
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 int32 = -475
                    _ = fi17
                    _ = v18
                }
                v19 := GenS3{F0: 1}
                _ = fi16
                _ = v19
            }
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = v8
    }
    _ = v7
    return 0
}

func main() {
    var v1 uint16 = 2
    v2 := []uint32{2147483648, 65535}
    v3 := []int{}
    var v4 float32 = 1.5
    var v5 int = 0
    var v6 uint32 = 0
    _ = len(v3)
    v6 += v6
    v6 -= uint32(v1)
    v20 := GenS3{F0: (v1 + v1)}
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 uint16 = (v1 + v1)
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS3{F0: (v22 % ((v22) | 1))}
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    var v25 uint32 = v6
    if false {
        var v26 uint = 54
        if false {
            fmt.Println(v6)
            _ = len(v3)
        } else {
            var v27 bool = ((-32768 > 42) || (false || (!(true || false))))
            _ = len(v3)
            var v28 uint16 = 255
            _ = v27
            _ = v28
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v26
    } else {
        var v29 bool = (!false)
        for fi30 := 0; fi30 < 10; fi30++ {
            fmt.Println(fi30, v1, v6)
            var v31 float64 = float64(float32(float64(v4)))
            var v32 int8 = 100
            _ = fi30
            _ = v31
            _ = v32
        }
        v33 := GenS4{F0: 42}
        _ = v29
        _ = v33
    }
    switch -9223372036854775808 {
    case -1:
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 int16 = 1000
            if true {
                var v36 float64 = 3.14159
                var v37 float32 = v4
                var v38 int8 = 127
                _ = v36
                _ = v37
                _ = v38
            } else {
                v1 = v1
                if (!true) {
                    if len(v2) > 0 {
                        v2[0] = (v6 - v25)
                    }
                    v39 := GenS4{F0: uint8(v1)}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v39
                } else {
                    if (65535 != 42) {
                        fmt.Println(v35, v35)
                    } else {
                        if true {
                            var v40 uint8 = 2
                            v41 := GenS5{F0: (v40 * 43)}
                            var v42 int64 = 9223372036854775807
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                        v5 = 4294967295
                        var v43 int = (-(fi34))
                        _ = v43
                    }
                }
            }
            _ = fi34
            _ = v35
        }
    }
    v6 += uint32(v1)
    for fi44 := 0; fi44 < 17; fi44++ {
        for fi45 := 0; fi45 < 4; fi45++ {
            v4 = (v4 - -2.25)
            _ = fi45
        }
        for fi46 := 0; fi46 < 11; fi46++ {
            v47 := GenS6{F0: 649}
            _ = fi46
            _ = v47
        }
        _ = fi44
    }
    v6 = 2147483648
    v48 := GenS7{F0: 65535, F1: ((v5 < v5) || (true && (1 > 106))), F2: 0}
    if (0 <= 199) {
        switch v5 {
        case 0:
            for fi49 := 0; fi49 < 14; fi49++ {
                for fi50 := 0; fi50 < 10; fi50++ {
                    if len(v2) > 0 {
                        v2[0] = v25
                    }
                    fmt.Println(v4, v1)
                    _ = fi50
                }
                _ = len(v3)
                _ = fi49
            }
        default:
            v51 := GenS8{F0: -533}
            _ = v51
        }
    } else {
        for fi52 := 0; fi52 < 17; fi52++ {
            v25 -= 810
            _ = fi52
        }
    }
    for fi53 := 0; fi53 < 13; fi53++ {
        v25 = 2
        switch "foo bar" {
        case "foo bar":
            for fi54 := 0; fi54 < 1; fi54++ {
                var v55 string = "foo bar"
                _ = fi54
                _ = v55
            }
        }
        _ = fi53
    }
    v56 := GenS9{F0: 2}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v57 int64 = -1
    var v58 float32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
