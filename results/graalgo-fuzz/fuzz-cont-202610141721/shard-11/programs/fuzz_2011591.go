package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int64 }
func genFn0(p0 string) int32 {
    switch -139 {
    case 128:
        for fi4 := 0; fi4 < 10; fi4++ {
            for fi5 := 0; fi5 < 5; fi5++ {
                switch (p0 + "a") {
                case "":
                    v6 := []int{fi4}
                    _ = v6
                case "a":
                    var v7 uint64 = 1
                    p0 = "\n\t"
                    v8 := GenS1{F0: p0, F1: 0, F2: 65535}
                    _ = v7
                    _ = v8
                default:
                    p0 = "foo bar"
                }
                _ = fi5
            }
            _ = fi4
        }
    case -2147483648:
        p0 = ""
    case -106:
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                p0 = p0
                v11 := GenS1{F0: (p0 + "foo bar"), F1: -2147483648, F2: 42}
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 2; fi12++ {
                if ((4294967295 < 4294967295) && true) {
                    var v13 float32 = float32(0.1)
                    v14 := []int{fi9, -931, 127, (-(fi12))}
                    var v15 int = (fi12 + (fi9 * fi12))
                    var v16 uint64 = 0
                    _ = v13
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    v17 := []int{fi9}
                    p0 = ""
                    _ = v17
                }
                _ = fi12
            }
            _ = fi9
        }
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        if (true || true) {
            switch (fi18 * fi18) {
            case 128:
                switch (^(fi18)) {
                case -860:
                    var v19 int8 = -104
                    _ = v19
                case 4294967295:
                    switch ("" + p0) {
                    case "":
                        p0 = (p0 + (p0 + ("foo bar" + p0)))
                    case "foo bar":
                        v20 := GenS1{F0: p0, F1: 42, F2: 65535}
                        _ = v20
                    default:
                        p0 = ("hello" + "")
                    }
                case -545:
                    fmt.Println(fi18, p0)
                default:
                    p0 = "a"
                }
            case 32768:
                fmt.Println(fi18, p0)
            default:
                v21 := []int{fi18, (fi18 << ((4294967295) & 63)), -9223372036854775808, 0}
                _ = v21
            }
            switch (fi18 * fi18) {
            case -9223372036854775808:
                if false {
                    fmt.Println(fi18, fi18)
                    var v22 int = (fi18 * fi18)
                    _ = v22
                } else {
                    var v23 float64 = 0.0
                    _ = v23
                }
            case 515:
                v24 := GenS2{F0: 742, F1: -1}
                _ = v24
            case -2147483648:
                fmt.Println(p0, p0)
            }
        }
        _ = fi18
    }
    p0 = p0
    var v25 uint8 = 2
    _ = p0
    _ = v25
    return 1
}
func genFn1(p0 int8, p1 int64) uint {
    switch -1 {
    case 32767:
        v26 := []int{2147483647}
        _ = v26
    default:
        fmt.Println(p0, p0)
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 uint64 = 444
        _ = fi27
        _ = v28
    }
    p0 = 0
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 919
    var v2 float32 = 0.0
    var v3 uint16 = 1
    if (true && (!((((0.1 / (1.5 + (0.1 / 1.0))) - 3.14159) - 3.14159) == (3.14159 - 1.5)))) {
        fmt.Println(v2, v3)
        if (127 != 46) {
            if false {
                var v29 int8 = 42
                _ = v29
            }
        } else {
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 uint16 = v3
                var v32 uint = 15
                _ = fi30
                _ = v31
                _ = v32
            }
            fmt.Println(v2)
        }
    }
    v2 -= v2
    switch (("hello" + "") + "a") {
    case "\n\t":
        for fi33 := 0; fi33 < 14; fi33++ {
            v1 -= 2147483648
            _ = fi33
        }
    case "foo bar":
        if (v1 != (v1 ^ 810)) {
            if true {
                v2 += v2
                if (!(42 >= 42)) {
                    for fi34 := 0; fi34 < 18; fi34++ {
                        v1 += v1
                        _ = fi34
                    }
                } else {
                    v2 = float32(((3.14159 - 0.1) / ((((3.14159 + 1e10) - 1.5) * (0.1 - 0.0)) + 0.1)))
                    for fi35 := 0; fi35 < 5; fi35++ {
                        v36 := []int{32767, (fi35 / ((fi35) | 1))}
                        for fi37 := 0; fi37 < 1; fi37++ {
                            var v38 uint = genFn1(0, -994)
                            _ = fi37
                            _ = v38
                        }
                        _ = fi35
                        _ = v36
                    }
                }
                switch 128 {
                case 0:
                    fmt.Println(v1)
                case -421:
                    v39 := []int{2147483647}
                    _ = v39
                default:
                    fmt.Println(v3, v3)
                }
            }
        }
    case "a":
        v40 := GenS1{F0: "foo bar", F1: 42, F2: 2147483648}
        _ = v40
    default:
        v2 = v2
    }
    switch 128 {
    case 42:
        if true {
            for fi41 := 0; fi41 < 16; fi41++ {
                var v42 int8 = 0
                if true {
                    v1 -= (v1 >> ((862) & 63))
                    v43 := []int{fi41}
                    _ = v43
                }
                _ = fi41
                _ = v42
            }
        } else {
            switch "a" {
            case "\n\t":
                fmt.Println(v3, v3, v1)
            default:
                var v44 int8 = 100
                _ = v44
            }
        }
    case 128:
        v3 = v3
    }
    var v45 uint = 0
    v46 := GenS2{F0: v45, F1: -13}
    var v47 uint32 = 4294967295
    fmt.Println(v1, v45)
    var v48 int64 = 42
    v49 := []int{0, 635, 32768, 42}
    switch "" {
    case "a":
        if false {
            fmt.Println(v48, v48, v45)
        }
    default:
        if (!true) {
            v50 := GenS2{F0: (v45 % ((v45) | 1)), F1: (v48 / (((v48 << (((v45 ^ v45)) & 63))) | 1))}
            if len(v49) >= 2 {
                v49 = v49[1:len(v49)]
            }
            _ = v50
        } else {
            fmt.Println(v2, v1)
        }
    }
    v45 += genFn1(0, v48)
    var v51 string = "foo bar"
    if len(v49) > 0 {
        v49[0] = int(v1)
    }
    for fi52 := 0; fi52 < 13; fi52++ {
        if ((v3 ^ 0) != (v3 % ((v3) | 1))) {
            var v53 uint8 = uint8(v3)
            var v54 int = fi52
            if len(v49) > 0 {
                v49[0] = v54
            }
            _ = v53
            _ = v54
        } else {
            switch (v51 + "hello") {
            case "\n\t":
                v55 := GenS3{F0: 473, F1: 218, F2: (v48 & v48)}
                _ = v55
            case "hello":
                v56 := GenS1{F0: v51, F1: genFn0(v51), F2: v47}
                _ = v56
            case "a":
                v3 -= v3
            }
        }
        switch v51 {
        case "":
            switch v51 {
            case "":
                fmt.Println(v48, v3)
            }
        }
        _ = fi52
    }
    v51 = (v51 + "hello")
    var v57 uint32 = (v1 ^ v47)
    v49 = append(v49, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v51:", v51)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v51
    _ = v57
}
