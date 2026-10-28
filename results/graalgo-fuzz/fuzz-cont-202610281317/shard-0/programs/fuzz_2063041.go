package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 bool }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int64 }
func genFn0(p0 string, p1 int32) int8 {
    var v6 int8 = -1
    v7 := []int{993}
    var v8 uint32 = 2147483647
    if true {
        var v9 int64 = 356
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return (v6 + -90)
}
func genFn1(p0 uint) uint32 {
    var v10 uint16 = 10
    v11 := GenS1{F0: int64(v10), F1: -32768, F2: 78}
    if false {
        p0 -= 2
        var v12 uint16 = 65535
        _ = v12
    } else {
        var v13 int64 = 1
        _ = v13
    }
    switch ("foo bar" + "hello") {
    case "\n\t":
        var v14 float32 = 3.14159
        _ = v14
    }
    var v15 uint32 = 65535
    _ = p0
    _ = v10
    _ = v11
    _ = v15
    return 773
}

func main() {
    var v1 uint8 = 1
    var v2 float64 = 1.5
    v3 := make(map[uint32]string)
    var v4 uint8 = (v1 ^ 255)
    v5 := []int{42}
    v4 += v4
    if (-1 > -2147483648) {
        var v16 int16 = 1
        _ = v16
    } else {
        if (v2 <= (float64(0.1) * (-2.25 - v2))) {
            v4 = v1
            var v17 uint64 = 4294967295
            switch "hello" {
            case "a":
                _ = len(v5)
            default:
                var v18 uint64 = v17
                _ = v18
            }
            _ = v17
        } else {
            fmt.Println(v2, v2, v4)
            var v19 int64 = 0
            _ = v19
        }
        if true {
            for fi20 := 0; fi20 < 11; fi20++ {
                v1 += (v4 | (v4 / (((v1 + v1)) | 1)))
                _ = fi20
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    var v21 int16 = 1
    switch "hello" {
    case "foo bar":
        v1 = v4
    case "hello":
        switch 385 {
        case 374:
            v1 = v4
        case -9223372036854775808:
            v21 = v21
        }
    }
    v21 = v21
    v22 := GenS2{F0: 2, F1: 2147483647, F2: false}
    var v23 uint = 2
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 int8 = 1
        for fi26 := 0; fi26 < 6; fi26++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v2 += 0.0
            switch "a" {
            case "foo bar":
                var v27 uint16 = 32767
                _ = v27
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = (-(fi24))
                }
            }
            _ = fi26
        }
        v25 += (v25 + 115)
        _ = fi24
        _ = v25
    }
    var v28 uint = (v23 - (^(v23)))
    if (1 == 0) {
        var v29 float32 = -2.25
        v28 += v28
        var v30 string = ("a" + "")
        _ = v29
        _ = v30
    } else {
        if false {
            for fi31 := 0; fi31 < 14; fi31++ {
                fmt.Println(v4, fi31, v28)
                _ = fi31
            }
        } else {
            for fi32 := 0; fi32 < 7; fi32++ {
                v4 = v1
                _ = fi32
            }
            var v33 uint8 = (v1 | 127)
            for fi34 := 0; fi34 < 7; fi34++ {
                var v35 int = fi34
                if len(v5) > 0 {
                    v5[0] = fi34
                }
                _ = fi34
                _ = v35
            }
            _ = v33
        }
        v36 := GenS2{F0: genFn1((^(v28))), F1: 2147483648, F2: ((-1 < 0) || true)}
        _ = v36
    }
    if (0 == 2147483647) {
        v37 := v3[454]
        v38 := GenS2{F0: 2, F1: 2, F2: true}
        switch 765 {
        case -44:
            var v39 uint64 = 0
            _ = v39
        case 0:
            v1 -= (v1 & v4)
        default:
            v23 -= v23
        }
        _ = v37
        _ = v38
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        switch "foo bar" {
        case "\n\t":
            var v41 uint = 65535
            _ = v41
        case "a":
            if false {
                if false {
                    switch (((("hello" + "foo bar") + "hello") + ((("" + "foo bar") + ("\n\t" + "hello")) + "")) + "") {
                    case "\n\t":
                        v2 -= v2
                    }
                    var v42 uint32 = 0
                    _ = v42
                }
                var v43 int32 = 0
                _ = v43
            }
        case "":
            v44 := GenS3{F0: 0, F1: -709}
            _ = v44
        default:
            v23 -= (^(v23))
        }
        for fi45 := 0; fi45 < 20; fi45++ {
            v2 += (0.0 - v2)
            var v46 int8 = -1
            _ = fi45
            _ = v46
        }
        if true {
            for fi47 := 0; fi47 < 17; fi47++ {
                var v48 int16 = v21
                _ = fi47
                _ = v48
            }
            var v49 uint32 = genFn1(v28)
            for fi50 := 0; fi50 < 2; fi50++ {
                switch ("\n\t" + "") {
                case "a":
                    fmt.Println(fi40, v2)
                case "":
                    _ = len(v5)
                }
                var v51 uint8 = v1
                _ = fi50
                _ = v51
            }
            _ = v49
        }
        _ = fi40
    }
    v28 = 18446744073709551615
    v52 := GenS4{F0: -9223372036854775808}
    for fi53 := 0; fi53 < 7; fi53++ {
        fmt.Println(v21, v28, fi53)
        for fi54 := 0; fi54 < 7; fi54++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v55 uint = v28
            _ = fi54
            _ = v55
        }
        _ = fi53
    }
    _ = len(v5)
    if (42 <= 42) {
        v56 := GenS4{F0: 9223372036854775807}
        _ = v56
    }
    var v57 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v57
}
