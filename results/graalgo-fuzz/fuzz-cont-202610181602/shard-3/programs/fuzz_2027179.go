package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int8; F1 uint64; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 uint8; F2 int8 }
type GenS6 struct { F0 int8; F1 int32 }
func genFn0(p0 int64, p1 bool, p2 uint) bool {
    var v6 uint = p2
    v7 := GenS1{F0: 42}
    var v8 uint32 = 789
    switch 128 {
    case -2147483648:
        p1 = p1
    case 2147483647:
        if (!true) {
            v9 := []int{}
            _ = v9
        }
    case 32767:
        switch "a" {
        case "":
            v8 -= (v8 - 42)
        default:
            for fi10 := 0; fi10 < 9; fi10++ {
                fmt.Println(v6, v8)
                var v11 string = ("foo bar" + ("a" + ("foo bar" + ("hello" + "foo bar"))))
                var v12 string = (v11 + "")
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    default:
        p2 = (v6 - (^(p2)))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (127 <= -1)
}
func genFn1(p0 int8, p1 int8) int16 {
    if ("\n\t" != ("\n\t" + (("" + "a") + ("foo bar" + "hello")))) {
        switch "a" {
        case "":
            v13 := GenS2{F0: 745, F1: 1}
            _ = v13
        }
        v14 := []int{338, 32768}
        _ = v14
    } else {
        if (!true) {
            v15 := GenS3{F0: p0, F1: 1, F2: (("\n\t" + "\n\t") + "foo bar")}
            p1 -= p1
            _ = v15
        } else {
            if (0 < -324) {
                var v16 float32 = 1e10
                var v17 uint64 = 42
                p0 = p0
                _ = v16
                _ = v17
            } else {
                var v18 bool = ((!(293 <= -2147483648)) && ((0 > 65535) || false))
                v18 = v18
                _ = v18
            }
            var v19 uint32 = 2
            _ = v19
        }
    }
    v20 := []int{-2147483648}
    fmt.Println(p0, p0)
    switch -173 {
    case 2147483647:
        p0 -= p0
    case 32768:
        fmt.Println(p1, p1)
    case 4294967295:
        if (-982 != -78) {
            switch "hello" {
            case "hello":
                fmt.Println(p0)
            case "a":
                var v21 float32 = 3.14159
                _ = v21
            case "":
                switch "\n\t" {
                case "":
                    var v22 uint32 = 2147483648
                    _ = v22
                case "foo bar":
                    fmt.Println(p1)
                }
            }
            var v23 int16 = int16(p0)
            for fi24 := 0; fi24 < 3; fi24++ {
                p0 = (p0 + p0)
                if false {
                    var v25 int16 = (v23 | v23)
                    _ = v25
                } else {
                    fmt.Println(p0)
                    p0 = p0
                }
                _ = fi24
            }
            _ = v23
        } else {
            var v26 float64 = (-2.25 - float64((float32(((1.5 / -2.25) + (float64(1.5) / 0.1))) / (float32((1e10 - 0.0)) * -2.25))))
            fmt.Println(p1)
            var v27 string = "foo bar"
            _ = v26
            _ = v27
        }
    default:
        v28 := GenS4{F0: "a"}
        _ = v28
    }
    if true {
        p1 = 50
        v29 := GenS4{F0: (("\n\t" + "hello") + ("a" + "foo bar"))}
        _ = v29
    }
    _ = p0
    _ = p1
    _ = v20
    return 877
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]bool)
    var v3 int8 = 127
    v4 := make(map[string]uint32)
    var v5 int8 = v3
    if true {
        v3 = v5
        delete(v4, (v1 + "\n\t"))
    } else {
        v30 := []int{-806, 42, 9223372036854775807}
        v31 := GenS5{F0: genFn1(100, v3), F1: 0, F2: (v3 % (((v5 | (v5 | v5))) | 1))}
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        v3 += (-(v5))
        var v33 bool = false
        v34 := []int{4294967295, (fi32 / ((fi32) | 1))}
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        var v36 int16 = 1000
        for fi37 := 0; fi37 < 8; fi37++ {
            v38 := []int{fi37, fi37, (-(fi35))}
            switch fi35 {
            case 4294967295:
                if (2147483647 < -677) {
                    if len(v38) > 0 {
                        v38[0] = fi37
                    }
                    var v39 int = -435
                    fmt.Println(v39)
                    _ = v39
                } else {
                    switch fi37 {
                    case 4294967295:
                        var v40 uint32 = 2
                        _ = v40
                    default:
                        _ = len(v2)
                    }
                    v3 += v3
                }
            }
            _ = fi37
            _ = v38
        }
        for fi41 := 0; fi41 < 2; fi41++ {
            var v42 int = fi35
            _ = fi41
            _ = v42
        }
        _ = fi35
        _ = v36
    }
    v1 = (v1 + v1)
    fmt.Println(v5)
    var v43 int64 = 0
    for fi44 := 0; fi44 < 14; fi44++ {
        v45 := GenS5{F0: genFn1(127, (v5 * 100)), F1: 26, F2: v5}
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        var v47 int32 = 1
        if true {
            v48 := GenS6{F0: v3, F1: (-(v47))}
            _ = v48
        }
        _ = fi46
        _ = v47
    }
    v49 := GenS5{F0: 42, F1: 0, F2: (v5 | (^(v5)))}
    switch 102 {
    case -9223372036854775808:
        v50 := []int{}
        _ = v50
    }
    var v51 bool = true
    switch -663 {
    case -2147483648:
        for fi52 := 0; fi52 < 1; fi52++ {
            if ((false || (128 < 207)) && true) {
                var v53 string = v1
                _ = v53
            }
            v54 := GenS3{F0: v3, F1: 4294967295, F2: "a"}
            fmt.Println(fi52, v3, v1)
            _ = fi52
            _ = v54
        }
    case 1:
        for fi55 := 0; fi55 < 11; fi55++ {
            for fi56 := 0; fi56 < 5; fi56++ {
                var v57 int32 = 0
                _ = len(v2)
                _ = fi56
                _ = v57
            }
            _ = fi55
        }
    default:
        if (v51 && false) {
            v58 := []int{2147483647, -82}
            for fi59 := 0; fi59 < 7; fi59++ {
                v5 += 1
                switch v1 {
                case "hello":
                    v3 -= (v3 >> ((0) & 63))
                default:
                    var v60 int32 = 547
                    _ = v60
                }
                delete(v4, v1)
                _ = fi59
            }
            _ = v58
        } else {
            var v61 uint16 = 32767
            _ = v61
        }
    }
    v62 := v4["hello"]
    v3 -= 100
    v63 := v2[34]
    var v64 int = 127
    var v65 int8 = v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v49
    _ = v5
    _ = v51
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
