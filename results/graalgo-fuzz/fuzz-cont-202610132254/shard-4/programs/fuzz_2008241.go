package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint64) int64 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 4; fi4++ {
        p0 -= 2
        _ = fi4
    }
    switch int(p0) {
    case 127:
        switch "\n\t" {
        case "":
            for fi5 := 0; fi5 < 7; fi5++ {
                var v6 int32 = 877
                _ = fi5
                _ = v6
            }
        }
    }
    var v7 int8 = -1
    if ((true && true) || (-1 == 2147483647)) {
        v7 = v7
        v8 := GenS1{F0: 42, F1: 1000}
        if (true && ((1 == 32767) && true)) {
            var v9 int32 = 2147483647
            v10 := GenS1{F0: 42, F1: 0}
            v11 := []int{-743, 32768}
            _ = v10
            _ = v11
            _ = v9
        }
        _ = v8
    }
    _ = p0
    _ = v7
    return -23
}
func genFn1(p0 int32, p1 int8) bool {
    for fi12 := 0; fi12 < 3; fi12++ {
        if true {
            var v13 int64 = genFn0(0)
            var v14 float64 = 3.14159
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := GenS1{F0: 2, F1: 29}
            _ = fi15
            _ = v16
        }
        switch (-(fi12)) {
        case -981:
            v17 := []int{(fi12 - fi12)}
            _ = v17
        case 42:
            fmt.Println(p0, p0)
        default:
            v18 := []int{fi12, fi12, (fi12 + fi12), fi12}
            _ = v18
        }
        _ = fi12
    }
    var v19 uint16 = 42
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 uint16 = 32767
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := GenS1{F0: 127, F1: 460}
            var v24 uint8 = 127
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v19
    return true
}

func main() {
    var v1 int8 = 42
    var v2 int8 = (-(v1))
    var v3 float64 = ((0.0 - 3.14159) - 0.0)
    fmt.Println(v2, v2, v3)
    switch "" {
    case "\n\t":
        var v25 uint32 = 677
        _ = v25
    }
    v1 += v2
    var v26 uint = 255
    var v27 bool = ((v26 / (((v26 & (v26 % (((v26 * 4294967295)) | 1)))) | 1)) == (v26 * v26))
    if (1.5 == float32((v3 * v3))) {
        var v28 uint = v26
        switch "hello" {
        case "a":
            v3 = float64((3.14159 / 3.14159))
        case "\n\t":
            for fi29 := 0; fi29 < 2; fi29++ {
                v30 := []int{fi29, fi29}
                fmt.Println(v1, fi29)
                switch (fi29 * (fi29 % (((-(fi29))) | 1))) {
                case 297:
                    for fi31 := 0; fi31 < 10; fi31++ {
                        if false {
                            fmt.Println(v2, v28)
                        } else {
                            fmt.Println(v2, v27)
                            fmt.Println(v27, v27, fi31)
                            v32 := GenS1{F0: 230, F1: 32767}
                            _ = v32
                        }
                        v26 = 18446744073709551615
                        _ = fi31
                    }
                case 194:
                    v26 -= (v28 + v28)
                case 9223372036854775807:
                    var v33 uint = v28
                    _ = v33
                default:
                    v1 += int8(v26)
                }
                _ = fi29
                _ = v30
            }
        case "foo bar":
            for fi34 := 0; fi34 < 9; fi34++ {
                v35 := GenS1{F0: 0, F1: 1000}
                fmt.Println(v2, v27)
                var v36 uint64 = 2
                _ = fi34
                _ = v35
                _ = v36
            }
        default:
            for fi37 := 0; fi37 < 17; fi37++ {
                v38 := []int{fi37, (fi37 - -451), fi37, (-(fi37))}
                var v39 uint32 = 2147483648
                if len(v38) >= 2 {
                    v38 = v38[1:len(v38)]
                }
                _ = fi37
                _ = v38
                _ = v39
            }
        }
        _ = v28
    } else {
        if ("" < "hello") {
            var v40 int = 127
            v3 = v3
            _ = v40
        }
        v2 = 1
    }
    v1 += (^(v2))
    var v41 bool = (v27 && (1 > -143))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v41
}
