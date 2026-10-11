package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint }
func genFn0() string {
    var v5 int8 = 0
    v6 := GenS1{F0: "foo bar", F1: 42}
    var v7 int64 = 1
    _ = v5
    _ = v6
    _ = v7
    return ""
}
func genFn1(p0 uint, p1 uint16, p2 int32) bool {
    var v8 float32 = (0.1 + 3.14159)
    if false {
        v9 := GenS2{F0: 0}
        v10 := GenS1{F0: ("" + ("hello" + "\n\t")), F1: 2147483647}
        var v11 bool = true
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        if ((p0 & (p0 & (p0 >> ((p0) & 63)))) == (p0 / ((p0) | 1))) {
            p0 = p0
            var v13 float64 = (1e10 / (float64(float32(-2.25)) * 1e10))
            _ = v13
        } else {
            if (1 < int64(p1)) {
                if false {
                    var v14 uint32 = 2
                    _ = v14
                }
                var v15 uint8 = 220
                _ = v15
            }
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return ((0.0 > 0.1) || (("" < "\n\t") || (564 > (p1 / ((p1) | 1)))))
}

func main() {
    var v1 int32 = -1
    var v2 int = 439
    var v3 int8 = 100
    v4 := make(map[uint32]int)
    for fi16 := 0; fi16 < 8; fi16++ {
        if true {
            switch (fi16 << ((65535) & 63)) {
            case 0:
                v17 := GenS2{F0: 186}
                _ = v17
            case 128:
                for fi18 := 0; fi18 < 11; fi18++ {
                    var v19 bool = (!true)
                    _ = fi18
                    _ = v19
                }
            case -1:
                v2 = int(v1)
            }
        } else {
            v3 = v3
            if true {
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 float32 = -2.25
                    v22 := []int{(v2 - v2), v2, (v2 * fi16), (fi20 / (((fi16 | (fi20 & 816))) | 1))}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v23 := []int{fi16}
                v3 -= int8(v1)
                _ = v23
            }
            v2 = -860
        }
        v1 += int32(v2)
        v24 := GenS1{F0: "", F1: -1}
        _ = fi16
        _ = v24
    }
    var v25 uint = 478
    v1 = 1
    if (42 >= 255) {
        switch "hello" {
        case "foo bar":
            v26 := v4[657]
            _ = v26
        case "":
            switch v2 {
            case 677:
                for fi27 := 0; fi27 < 1; fi27++ {
                    v28 := GenS1{F0: "", F1: (v1 ^ 0)}
                    if ((fi27 & 2147483647) < (fi27 + (v2 << ((18446744073709551615) & 63)))) {
                        v4[0] = fi27
                        v29 := GenS3{F0: (v25 - (^(v25))), F1: 1, F2: 124}
                        fmt.Println(v1, fi27, v3)
                        _ = v29
                    }
                    _ = fi27
                    _ = v28
                }
            case 0:
                switch (-(v2)) {
                case 2147483647:
                    var v30 uint8 = 255
                    _ = v30
                case 1:
                    var v31 int64 = -1
                    _ = v31
                case 4294967295:
                    v25 -= v25
                default:
                    v2 = v2
                }
            }
        case "\n\t":
            var v32 int16 = 20
            _ = v32
        default:
            var v33 uint32 = 423
            _ = v33
        }
    } else {
        var v34 float64 = (3.14159 / -2.25)
        fmt.Println(v34)
        var v35 string = ("foo bar" + "\n\t")
        _ = v34
        _ = v35
    }
    if (false || (!false)) {
        var v36 int8 = v3
        _ = v36
    } else {
        for fi37 := 0; fi37 < 5; fi37++ {
            v38 := GenS3{F0: v25, F1: 477, F2: 255}
            var v39 uint = v25
            _ = fi37
            _ = v38
            _ = v39
        }
        v3 = -65
    }
    delete(v4, 2147483647)
    v2 = (v2 - 768)
    var v40 bool = true
    for fi41 := 0; fi41 < 10; fi41++ {
        for fi42 := 0; fi42 < 3; fi42++ {
            var v43 int8 = 1
            _ = fi42
            _ = v43
        }
        fmt.Println(v3, v3)
        fmt.Println(v3, fi41)
        _ = fi41
    }
    var v44 int16 = -1
    if true {
        v25 = 255
        var v45 uint8 = 103
        _ = v45
    }
    var v46 int32 = v1
    _ = len(v4)
    for fi47 := 0; fi47 < 8; fi47++ {
        fmt.Println(v40, v40, v40)
        if ((false && ((1e10 * float32(((0.0 - (3.14159 / 1.0)) * ((3.14159 / -2.25) / -2.25)))) > 0.0)) || v40) {
            v46 += -2147483648
            v48 := GenS4{F0: 2}
            if v40 {
                if v40 {
                    v3 += v3
                    var v49 bool = false
                    _ = v49
                }
                v40 = ((-(v3)) != (v3 - (v3 & (v3 - v3))))
                if true {
                    var v50 int64 = -9223372036854775808
                    _ = v50
                }
            }
            _ = v48
        }
        _ = fi47
    }
    var v51 int64 = -175
    delete(v4, 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v44
    _ = v46
    _ = v51
}
