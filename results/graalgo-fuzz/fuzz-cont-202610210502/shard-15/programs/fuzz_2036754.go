package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 int16 }
func genFn0(p0 int16, p1 uint8) int32 {
    p1 = p1
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint32 = 402
        var v9 bool = true
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []uint32{1, 42}
    var v2 uint = 2
    var v3 int32 = 748
    var v4 float32 = -2.25
    var v5 uint8 = 255
    var v6 float64 = 0.0
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS1{F0: 9223372036854775807, F1: ("foo bar" + "foo bar")}
            _ = fi11
            _ = v12
        }
        if (!(0 >= 1)) {
            var v13 int8 = -128
            if (!false) {
                if len(v1) > 0 {
                    v1[0] = 65535
                }
            } else {
                v14 := GenS2{F0: "a", F1: v13, F2: int8(v3)}
                switch "\n\t" {
                case "foo bar":
                    switch 1 {
                    case -1:
                        var v15 uint = v2
                        _ = v15
                    }
                }
                _ = v14
            }
            if false {
                var v16 uint8 = (v5 % ((206) | 1))
                fmt.Println(v5)
                _ = v16
            }
            _ = v13
        }
        _ = fi10
    }
    v4 -= 0.0
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 string = ("hello" + "")
        _ = fi17
        _ = v18
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    var v19 uint8 = (v5 % (((^(v5))) | 1))
    if false {
        fmt.Println(v6)
    } else {
        var v20 float32 = v4
        fmt.Println(v6)
        _ = v20
    }
    v21 := GenS3{F0: -865}
    var v22 string = ("foo bar" + "a")
    var v23 float64 = (v6 * 3.14159)
    v23 -= v6
    var v24 int64 = 42
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi26 := 0; fi26 < 7; fi26++ {
                v27 := GenS2{F0: v22, F1: 100, F2: 71}
                v24 = 0
                _ = fi26
                _ = v27
            }
            var v28 float64 = float64(v4)
            _ = v28
        }
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 uint64 = 953
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            _ = fi29
            _ = v30
        }
        _ = fi25
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        fmt.Println(v19, v23)
        v1 = append(v1, 0)
        _ = fi31
    }
    v32 := GenS3{F0: -434}
    for fi33 := 0; fi33 < 6; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            fmt.Println(v4, fi33, v22)
            fmt.Println(v23, v19)
            v23 += (float64(-2.25) - (v23 * (v23 - float64(v4))))
            v35 := GenS4{F0: 1}
            _ = fi34
            _ = v35
        }
        if true {
            if (int(v2) <= fi33) {
                v2 -= v2
            }
            if len(v1) > 0 {
                v1[0] = 781
            }
        } else {
            v36 := GenS1{F0: v24, F1: (((((((v22 + "a") + (("foo bar" + "foo bar") + v22)) + "foo bar") + "\n\t") + ("" + "a")) + v22) + (("foo bar" + "\n\t") + ("\n\t" + v22)))}
            _ = v36
        }
        _ = fi33
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        v24 += 444
        v38 := GenS5{F0: 0, F1: int16(v24)}
        _ = fi37
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
