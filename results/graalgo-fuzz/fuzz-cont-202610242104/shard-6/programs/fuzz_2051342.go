package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32; F1 uint64; F2 int32 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint; F1 int32; F2 uint }
func genFn0(p0 uint, p1 uint16, p2 bool) int64 {
    for fi7 := 0; fi7 < 3; fi7++ {
        p2 = (-2.25 != ((1.5 * (float32(-2.25) * 1e10)) / 1.0))
        var v8 uint16 = p1
        v9 := GenS1{F0: ("hello" + ("a" + ((("\n\t" + "a") + "a") + "\n\t"))), F1: 65535}
        _ = fi7
        _ = v8
        _ = v9
    }
    p2 = (p2 || (4294967295 != 9223372036854775807))
    if p2 {
        v10 := []int{}
        var v11 uint32 = 2147483648
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 4; fi12++ {
            if p2 {
                switch "\n\t" {
                case "foo bar":
                    var v13 float64 = 0.0
                    _ = v13
                }
                var v14 string = (("hello" + "foo bar") + "")
                v15 := []int{fi12}
                _ = v14
                _ = v15
            }
            p1 += p1
            _ = fi12
        }
        for fi16 := 0; fi16 < 14; fi16++ {
            p0 = (p0 % (((^(p0))) | 1))
            _ = fi16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"hello"}
    v3 := make(map[int]uint32)
    var v4 int64 = -9223372036854775808
    var v5 string = "foo bar"
    v6 := make(map[string]string)
    for fi17 := 0; fi17 < 12; fi17++ {
        switch v5 {
        case "a":
            if ((false || ((!((!false) || true)) || true)) || (!true)) {
                delete(v6, "hello")
            }
        case "\n\t":
            var v18 int16 = 1000
            _ = v18
        }
        v5 = v5
        fmt.Println(v5, fi17)
        _ = fi17
    }
    v19 := GenS1{F0: v5, F1: 42}
    if ((v4 * v4) <= v4) {
        var v20 int64 = 1
        v5 = v5
        if true {
            v21 := GenS2{F0: 1, F1: 1, F2: 1}
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v5, v5)
            }
            v22 := GenS2{F0: 819, F1: 4294967295, F2: -2147483648}
            _ = v21
            _ = v22
        }
        _ = v20
    } else {
        var v23 int8 = -120
        v24 := GenS2{F0: 1, F1: 1, F2: 0}
        _ = v23
        _ = v24
    }
    if (false && (!false)) {
        if false {
            for fi25 := 0; fi25 < 12; fi25++ {
                v4 = v4
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 int = fi26
                    _ = len(v3)
                    v27 += v27
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
        } else {
            if len(v2) > 0 {
                v2[0] = ""
            }
        }
        var v28 uint32 = 920
        for fi29 := 0; fi29 < 3; fi29++ {
            for fi30 := 0; fi30 < 13; fi30++ {
                v31 := GenS3{F0: v4, F1: 42}
                if (1 != 127) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v5 = ((("" + v5) + (v5 + "foo bar")) + v5)
                    var v32 bool = (!(!(0.1 >= -2.25)))
                    _ = v32
                }
                if true {
                    if (!(!(106 >= 1))) {
                        var v33 uint64 = 2
                        var v34 uint16 = 785
                        _ = v33
                        _ = v34
                    } else {
                        var v35 uint8 = 255
                        _ = v35
                    }
                    var v36 uint32 = v28
                    var v37 bool = ((42 == int16(v4)) || (((1e10 + (1.5 * (3.14159 / (((3.14159 * (0.1 / 3.14159)) / 1e10) + 3.14159)))) - 1.5) < -2.25))
                    _ = v36
                    _ = v37
                } else {
                    var v38 uint64 = 0
                    _ = v38
                }
                _ = fi30
                _ = v31
            }
            v4 += v4
            _ = fi29
        }
        _ = v28
    } else {
        for fi39 := 0; fi39 < 3; fi39++ {
            for fi40 := 0; fi40 < 6; fi40++ {
                if (true && false) {
                    if (true || false) {
                        v4 -= (^(v4))
                    } else {
                        v41 := GenS3{F0: v4, F1: 4294967295}
                        v4 = v4
                        _ = v41
                    }
                    _ = len(v2)
                }
                _ = fi40
            }
            v4 = 42
            _ = fi39
        }
    }
    v4 = (v4 % ((v4) | 1))
    if (0 <= 4294967295) {
        if false {
            v4 += (v4 & -522)
            v42 := GenS4{F0: 90, F1: 0, F2: 902}
            _ = v42
        } else {
            v5 = v5
            if len(v2) > 0 {
                v2[0] = v5
            }
        }
    }
    v4 += 42
    if (false || ((172 > 750) || (!true))) {
        if len(v2) > 0 {
            v2[0] = (v5 + v5)
        }
        var v43 int = -415
        var v44 int32 = 42
        _ = v43
        _ = v44
    }
    var v45 uint = 2
    v45 = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v45)
    v46 := GenS4{F0: 255, F1: 899, F2: v45}
    var v47 bool = (!false)
    fmt.Println(v45, v5, v4)
    v4 = v4
    if v47 {
        for fi48 := 0; fi48 < 18; fi48++ {
            fmt.Println(fi48)
            v4 -= (v4 % ((v4) | 1))
            _ = fi48
        }
        v3[87] = 602
    } else {
        v2 = append(v2, ("a" + "\n\t"))
        var v49 int64 = (v4 + (v4 << (((^(v45))) & 63)))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
