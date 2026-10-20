package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int16; F1 int8; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS6 struct { F0 string }
func genFn0(p0 uint64, p1 string, p2 int) string {
    var v6 int8 = 127
    v7 := []int{p2, (-(p2)), (p2 << ((484) & 63)), p2}
    v6 -= v6
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            switch fi8 {
            case 2147483647:
                var v10 int16 = 1
                _ = v10
            case 9223372036854775807:
                var v11 float32 = ((float32(float64(0.1)) / float32(float64(1.5))) + 1.5)
                _ = v11
            }
            v12 := GenS1{F0: 746}
            _ = fi9
            _ = v12
        }
        var v13 uint32 = 4294967295
        _ = fi8
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}

func main() {
    v1 := []string{("" + "a"), "foo bar"}
    var v2 uint64 = 0
    v3 := make(map[int]uint32)
    var v4 float64 = float64((3.14159 / 1.5))
    var v5 string = "\n\t"
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 uint16 = 2
        _ = fi14
        _ = v15
    }
    switch v5 {
    case "a":
        var v16 float32 = 0.0
        _ = v16
    default:
        v2 += v2
    }
    v5 = (v5 + v5)
    if false {
        v4 = v4
        fmt.Println(v4)
        switch 9223372036854775807 {
        case -1:
            fmt.Println(v2, v2, v2)
        case -2147483648:
            var v17 uint32 = 625
            _ = v17
        case 128:
            switch 125 {
            case 909:
                for fi18 := 0; fi18 < 14; fi18++ {
                    v19 := GenS2{F0: 42, F1: "a"}
                    _ = fi18
                    _ = v19
                }
            default:
                for fi20 := 0; fi20 < 13; fi20++ {
                    for fi21 := 0; fi21 < 13; fi21++ {
                        var v22 uint8 = 2
                        _ = fi21
                        _ = v22
                    }
                    _ = fi20
                }
            }
        }
    }
    var v23 uint = 255
    var v24 float64 = (v4 + 0.1)
    switch (v5 + ("foo bar" + (((("hello" + "a") + "a") + v5) + "hello"))) {
    case "hello":
        var v25 float64 = float64(3.14159)
        _ = v25
    }
    fmt.Println(v24, v5, v2)
    if (true || false) {
        if false {
            v4 -= v4
            var v26 uint = (v23 | 1)
            _ = v26
        }
        for fi27 := 0; fi27 < 5; fi27++ {
            v23 -= (v23 >> (((v23 - v23)) & 63))
            fmt.Println(v23)
            _ = fi27
        }
        switch v5 {
        case "foo bar":
            v2 += v2
        case "\n\t":
            fmt.Println(v2, v5, v4)
        default:
            if (!(36 <= 244)) {
                fmt.Println(v2)
            } else {
                var v28 int32 = 42
                _ = v28
            }
        }
    }
    if (v24 == ((v4 / (v4 / -2.25)) * (v4 * v4))) {
        var v29 uint64 = v2
        v30 := GenS1{F0: 0}
        v5 = ""
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        if (false || ((!((((0.1 * (1e10 / v4)) / (float64(1.5) / 1.5)) <= (0.1 / v4)) && true)) && (((-2.25 - float64(0.1)) - v24) == (v24 / (v24 / v4))))) {
            var v32 int8 = -1
            _ = v32
        }
        _ = fi31
    }
    if (0 < 255) {
        var v33 float64 = v4
        if false {
            if (true && false) {
                var v34 int64 = -836
                if false {
                    delete(v3, 4294967295)
                    var v35 uint = uint(v2)
                    _ = v35
                }
                v36 := GenS3{F0: -32768, F1: 42, F2: 729}
                _ = v34
                _ = v36
            }
        } else {
            if (((127 < 100) && (!true)) && (0 <= 127)) {
                var v37 int = 128
                _ = v37
            } else {
                var v38 uint8 = uint8(v2)
                v33 = v24
                _ = v38
            }
            v39 := GenS3{F0: 1, F1: 42, F2: 4294967295}
            v40 := GenS2{F0: 126, F1: v5}
            _ = v39
            _ = v40
        }
        if false {
            for fi41 := 0; fi41 < 1; fi41++ {
                switch v5 {
                case "\n\t":
                    fmt.Println(v2)
                case "a":
                    var v42 uint32 = 85
                    _ = v42
                }
                v43 := GenS4{F0: 127, F1: (v2 - (v2 >> (((^(v23))) & 63))), F2: 0}
                if false {
                    var v44 int64 = 1
                    var v45 string = (v5 + (v5 + "hello"))
                    for fi46 := 0; fi46 < 6; fi46++ {
                        v47 := GenS5{F0: (v23 >> ((v23) & 63)), F1: 2, F2: 723}
                        v5 = ""
                        v23 = v23
                        _ = fi46
                        _ = v47
                    }
                    _ = v44
                    _ = v45
                }
                _ = fi41
                _ = v43
            }
            var v48 int32 = int32(v2)
            for fi49 := 0; fi49 < 13; fi49++ {
                var v50 bool = (!((41 != 0) || (!(float32(1.5) == -2.25))))
                v51 := GenS2{F0: 127, F1: v5}
                _ = fi49
                _ = v50
                _ = v51
            }
            _ = v48
        } else {
            if (false && (!(false || true))) {
                delete(v3, 139)
            }
        }
        _ = v33
    } else {
        for fi52 := 0; fi52 < 10; fi52++ {
            v53 := GenS2{F0: 100, F1: v5}
            v54 := GenS5{F0: 0, F1: 65535, F2: 65535}
            var v55 uint32 = 0
            _ = fi52
            _ = v53
            _ = v54
            _ = v55
        }
        v5 = v5
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    _ = len(v3)
    var v56 int32 = -831
    v4 = ((v4 * float64((1e10 + 1e10))) + v24)
    _ = len(v1)
    v57 := GenS6{F0: v5}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
}
