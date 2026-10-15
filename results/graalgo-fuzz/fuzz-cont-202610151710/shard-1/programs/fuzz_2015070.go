package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint; F2 string }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 uint64; F1 string; F2 int64 }
func genFn0() int16 {
    if false {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 uint8 = 139
            _ = fi7
            _ = v8
        }
    } else {
        var v9 int32 = 2147483647
        var v10 int = -1
        var v11 float32 = 0.0
        _ = v10
        _ = v11
        _ = v9
    }
    if true {
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := []int{(fi12 % (((fi12 & 706)) | 1)), 363, fi12}
            if true {
                var v14 uint8 = 2
                _ = v14
            } else {
                v15 := GenS1{F0: 2147483648}
                v16 := GenS1{F0: 2}
                var v17 int64 = -393
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi12
            _ = v13
        }
    }
    switch 2147483647 {
    case 1:
        var v18 uint16 = 0
        _ = v18
    }
    v19 := GenS1{F0: 2147483648}
    if (-746 <= 141) {
        var v20 float64 = (0.0 + 3.14159)
        _ = v20
    }
    _ = v19
    return 1
}
func genFn1(p0 int64) uint8 {
    switch "hello" {
    case "\n\t":
        var v21 string = ""
        _ = v21
    default:
        v22 := GenS2{F0: p0}
        _ = v22
    }
    var v23 int64 = p0
    for fi24 := 0; fi24 < 18; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS3{F0: 1, F1: 1, F2: ("" + "")}
            v27 := GenS3{F0: 1, F1: 42, F2: "a"}
            _ = fi25
            _ = v26
            _ = v27
        }
        v28 := []int{}
        for fi29 := 0; fi29 < 7; fi29++ {
            switch "\n\t" {
            case "a":
                var v30 uint8 = 1
                _ = v30
            case "hello":
                v28 = append(v28, 32768)
            case "\n\t":
                v23 -= v23
            default:
                v23 = (v23 / (((p0 ^ 0)) | 1))
            }
            _ = fi29
        }
        _ = fi24
        _ = v28
    }
    switch "\n\t" {
    case "a":
        p0 += v23
    case "hello":
        var v31 uint32 = 280
        _ = v31
    }
    _ = p0
    _ = v23
    return 7
}

func main() {
    var v1 string = "\n\t"
    var v2 int8 = 42
    var v3 uint = 42
    var v4 bool = true
    var v5 float32 = (3.14159 * 3.14159)
    var v6 float32 = (((v5 - v5) - v5) - ((v5 - ((v5 - (((v5 + 3.14159) + v5) + (((v5 + (0.0 / ((v5 / v5) * v5))) + -2.25) - v5))) * 0.0)) * 1e10))
    if (0 != 985) {
        v1 = ("hello" + v1)
        if ((((1.5 - 0.0) / -2.25) / 1.0) > 1e10) {
            v5 = (v5 * v5)
        } else {
            fmt.Println(v1, v3)
            var v32 uint32 = 2
            v33 := []int{-1}
            _ = v32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        v4 = v4
        _ = fi34
    }
    var v35 uint32 = 2
    if (!(2147483647 < int(v3))) {
        v36 := GenS4{F0: 1, F1: 2}
        v5 -= (v6 / 1.5)
        _ = v36
    } else {
        switch v1 {
        case "a":
            v37 := []int{-691, int(v2), -2147483648, -9223372036854775808}
            _ = v37
        case "\n\t":
            v38 := GenS3{F0: 1, F1: v3, F2: v1}
            _ = v38
        default:
            var v39 bool = false
            _ = v39
        }
        v6 = v5
    }
    if ((genFn1(9223372036854775807) <= 42) || ((((v1 + v1) < v1) && (42 >= 1)) && (226 != genFn1(42)))) {
        v40 := []int{32768, 121, 32767, 9223372036854775807}
        for fi41 := 0; fi41 < 18; fi41++ {
            v5 -= float32((3.14159 + (-2.25 + 1e10)))
            _ = fi41
        }
        _ = v40
    } else {
        for fi42 := 0; fi42 < 18; fi42++ {
            if (!(v4 && v4)) {
                v43 := []int{}
                fmt.Println(fi42)
                _ = v43
            } else {
                if (((v35 * v35) >= v35) || ((v4 && v4) && (!(!(false || (((477 != 42) && (!(true || v4))) || v4)))))) {
                    var v44 bool = true
                    _ = v44
                }
            }
            _ = fi42
        }
        var v45 int8 = (v2 << ((v3) & 63))
        for fi46 := 0; fi46 < 12; fi46++ {
            v47 := []int{(fi46 - 521), fi46}
            v48 := GenS5{F0: 798, F1: v35, F2: 42}
            _ = fi46
            _ = v47
            _ = v48
        }
        _ = v45
    }
    v49 := GenS3{F0: 42, F1: v3, F2: "\n\t"}
    v50 := GenS3{F0: 2, F1: v3, F2: v1}
    v51 := GenS3{F0: 91, F1: (v3 << ((v3) & 63)), F2: v1}
    var v52 string = "foo bar"
    fmt.Println(v2)
    for fi53 := 0; fi53 < 7; fi53++ {
        v54 := GenS6{F0: 0, F1: 42}
        _ = fi53
        _ = v54
    }
    v3 -= 0
    v55 := GenS1{F0: (v35 * (v35 >> ((4294967295) & 63)))}
    v5 -= v5
    v56 := GenS6{F0: 32767, F1: 2147483647}
    if (864 == -2147483648) {
        v57 := GenS7{F0: v4, F1: -315}
        _ = v57
    }
    v58 := GenS7{F0: ((!((v4 || v4) && (1 < 42))) || v4), F1: 9223372036854775807}
    switch v52 {
    case "":
        if v4 {
            v59 := GenS8{F0: 42, F1: (("foo bar" + v52) + v52), F2: -601}
            v3 -= v3
            v4 = ((473 != 1) || (-2147483648 >= -2147483648))
            _ = v59
        }
    case "a":
        fmt.Println(v2)
    case "\n\t":
        for fi60 := 0; fi60 < 14; fi60++ {
            var v61 int = fi60
            _ = fi60
            _ = v61
        }
    default:
        var v62 int16 = 1
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v55
    _ = v56
    _ = v58
    _ = v6
}
