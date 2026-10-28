package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int8; F2 int32 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int32, p1 int8, p2 string) uint {
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        p2 = ("foo bar" + "hello")
    case "\n\t":
        for fi6 := 0; fi6 < 8; fi6++ {
            if true {
                var v7 int = fi6
                _ = v7
            }
            v8 := GenS1{F0: false, F1: 996, F2: 489}
            _ = fi6
            _ = v8
        }
    case "a":
        v9 := GenS2{F0: 2, F1: p1, F2: p0}
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS2{F0: 42, F1: 42, F2: p0}
        v12 := GenS1{F0: (269 >= 526), F1: 32767, F2: 42}
        switch fi10 {
        case -2147483648:
            v13 := GenS1{F0: true, F1: 242, F2: 42}
            _ = v13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 string = ("" + ("" + ("a" + ("hello" + "a"))))
    v3 := make(map[string]string)
    var v4 int64 = -9223372036854775808
    v5 := make(map[int]bool)
    var v14 int = 673
    switch ((((v2 + v2) + (v2 + v2)) + "\n\t") + v2) {
    case "a":
        switch (v14 / ((v14) | 1)) {
        case 418:
            var v15 uint16 = 65535
            _ = v15
        }
    case "":
        fmt.Println(v4, v4, v2)
    }
    v16 := GenS2{F0: 1, F1: 1, F2: -956}
    v17 := []int{(-(v14)), v14, v14, v14}
    if true {
        fmt.Println(v14, v14)
        v18 := GenS3{F0: v14, F1: v4}
        switch v14 {
        case 42:
            for fi19 := 0; fi19 < 12; fi19++ {
                v20 := GenS4{F0: (!(!true)), F1: 1}
                var v21 int64 = v4
                _ = fi19
                _ = v20
                _ = v21
            }
        case 9223372036854775807:
            if true {
                for fi22 := 0; fi22 < 14; fi22++ {
                    if (true || false) {
                        v23 := GenS2{F0: 42, F1: 107, F2: 91}
                        fmt.Println(v4, v2, fi22)
                        _ = v23
                    }
                    if true {
                        v2 = ""
                    }
                    _ = fi22
                }
            }
        case 128:
            v4 = v4
        default:
            for fi24 := 0; fi24 < 5; fi24++ {
                v2 = "hello"
                _ = fi24
            }
        }
        _ = v18
    } else {
        v25 := GenS5{F0: 91, F1: (v4 % ((v4) | 1))}
        if true {
            v17 = append(v17, (v14 ^ v14))
            var v26 uint = 18446744073709551615
            _ = v26
        }
        _ = v25
    }
    if (((3.14159 - 1e10) <= (1e10 * (0.1 * 0.1))) || true) {
        if true {
            v4 = 436
            fmt.Println(v4)
        } else {
            v2 = "foo bar"
            for fi27 := 0; fi27 < 14; fi27++ {
                v2 = "hello"
                var v28 string = ((((v2 + v2) + v2) + "foo bar") + v2)
                _ = fi27
                _ = v28
            }
            v2 = v2
        }
        switch (v14 | (^(v14))) {
        case -1:
            delete(v1, (v14 << ((0) & 63)))
        case -850:
            var v29 uint32 = 2
            _ = v29
        default:
            if (42 != 42) {
                fmt.Println(v4, v4, v2)
                v30 := GenS5{F0: 65535, F1: 0}
                var v31 float32 = (1.5 / 0.1)
                _ = v30
                _ = v31
            }
        }
        v32 := GenS6{F0: 42, F1: 2147483647}
        _ = v32
    } else {
        var v33 float32 = 1.5
        _ = v33
    }
    v14 = v14
    fmt.Println(v14, v2, v14)
    switch (v14 ^ v14) {
    case -867:
        for fi34 := 0; fi34 < 1; fi34++ {
            for fi35 := 0; fi35 < 12; fi35++ {
                if true {
                    v36 := GenS1{F0: true, F1: -763, F2: 255}
                    v14 = fi34
                    _ = v36
                } else {
                    var v37 float64 = 1e10
                    _ = v37
                }
                _ = len(v1)
                _ = fi35
            }
            fmt.Println(v4)
            _ = fi34
        }
    case 32768:
        for fi38 := 0; fi38 < 18; fi38++ {
            v14 = fi38
            if (350 <= 0) {
                v2 = "a"
                _ = len(v1)
                v39 := GenS6{F0: 1, F1: 635}
                _ = v39
            }
            v14 -= (fi38 - fi38)
            _ = fi38
        }
    case 2147483647:
        for fi40 := 0; fi40 < 16; fi40++ {
            v2 = v2
            fmt.Println(v2, v4, fi40)
            v41 := GenS7{F0: 4294967295}
            _ = fi40
            _ = v41
        }
    }
    fmt.Println(v4, v14, v14)
    if false {
        var v42 int16 = -1
        if len(v17) > 0 {
            v17[0] = 4294967295
        }
        _ = v42
    } else {
        for fi43 := 0; fi43 < 19; fi43++ {
            for fi44 := 0; fi44 < 9; fi44++ {
                var v45 int = fi43
                v4 += v4
                _ = fi44
                _ = v45
            }
            _ = fi43
        }
        v14 -= v14
    }
    if true {
        v4 -= v4
        if true {
            for fi46 := 0; fi46 < 19; fi46++ {
                for fi47 := 0; fi47 < 5; fi47++ {
                    var v48 int16 = -174
                    v2 = "hello"
                    v3[(v2 + v2)] = (v2 + ((v2 + ("" + v2)) + "hello"))
                    _ = fi47
                    _ = v48
                }
                switch (((v2 + v2) + v2) + v2) {
                case "a":
                    var v49 float32 = (0.1 - 0.1)
                    _ = v49
                }
                var v50 int32 = 2147483647
                _ = fi46
                _ = v50
            }
            for fi51 := 0; fi51 < 17; fi51++ {
                if (!(!true)) {
                    v4 += -9223372036854775808
                    v2 = v2
                }
                _ = len(v3)
                _ = fi51
            }
            fmt.Println(v14, v4)
        } else {
            switch v14 {
            case 128:
                var v52 bool = ((!false) && true)
                _ = v52
            case -9223372036854775808:
                v1[32768] = 127
            }
            v53 := GenS8{F0: 2, F1: 0}
            _ = v53
        }
        if true {
            if (false && true) {
                _ = len(v1)
            }
            if (-2147483648 > 2147483647) {
                var v54 float32 = 0.0
                _ = v54
            }
            for fi55 := 0; fi55 < 13; fi55++ {
                switch fi55 {
                case -9223372036854775808:
                    var v56 int64 = (v4 * int64(fi55))
                    _ = v56
                case 128:
                    switch (v2 + "") {
                    case "a":
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                    case "":
                        var v57 uint = 18446744073709551615
                        _ = v57
                    default:
                        var v58 int16 = 1
                        _ = v58
                    }
                }
                _ = fi55
            }
        }
    } else {
        v14 = -2147483648
        var v59 float32 = 1.5
        v4 += v4
        _ = v59
    }
    switch v14 {
    case 4294967295:
        v4 += v4
    case 0:
        switch (v2 + v2) {
        case "a":
            if (-128 == -1) {
                var v60 uint8 = 42
                _ = v60
            }
        }
    }
    fmt.Println(v2)
    for fi61 := 0; fi61 < 17; fi61++ {
        v4 = -9223372036854775808
        _ = fi61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
