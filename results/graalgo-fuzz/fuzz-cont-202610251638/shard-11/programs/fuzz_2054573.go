package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 string }
func genFn0(p0 uint32, p1 int) uint64 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 950}
        switch (fi7 | fi7) {
        case 699:
            p1 -= fi7
        case -472:
            p0 = (p0 * p0)
        default:
            for fi9 := 0; fi9 < 13; fi9++ {
                p1 += (fi9 % (((p1 % (((fi9 | -456)) | 1))) | 1))
                var v10 int16 = 0
                _ = fi9
                _ = v10
            }
        }
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint64 = 0
        v13 := GenS2{F0: 1, F1: (p0 | p0)}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int16, p1 uint64, p2 string) int8 {
    var v14 int32 = 42
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int16 = 0
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 1
}

func main() {
    var v1 float32 = -2.25
    var v2 uint = 193
    var v3 uint8 = 127
    v4 := make(map[uint32]int)
    var v5 uint16 = 65535
    var v6 string = "foo bar"
    fmt.Println(v6)
    var v17 uint16 = v5
    if false {
        var v18 bool = true
        v19 := GenS3{F0: 20}
        _ = v18
        _ = v19
    }
    v20 := GenS4{F0: (true && false), F1: (v6 + v6)}
    var v21 string = "foo bar"
    for fi22 := 0; fi22 < 10; fi22++ {
        v5 = uint16(v3)
        fmt.Println(v6, v3)
        _ = fi22
    }
    v5 += v17
    var v23 int16 = -1
    if true {
        if (680 >= v23) {
            fmt.Println(v21, v1, v2)
            switch "foo bar" {
            case "\n\t":
                var v24 uint32 = uint32(v3)
                _ = v24
            case "":
                fmt.Println(v5, v3, v23)
            case "foo bar":
                v25 := []int{9223372036854775807, 0}
                _ = v25
            }
            fmt.Println(v1, v3, v17)
        } else {
            var v26 int8 = 100
            _ = v26
        }
    }
    if false {
        var v27 int64 = -9223372036854775808
        switch "\n\t" {
        case "\n\t":
            var v28 int = 1
            _ = v28
        }
        var v29 int8 = 42
        _ = v27
        _ = v29
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        var v31 float32 = v1
        _ = fi30
        _ = v31
    }
    fmt.Println(v3)
    for fi32 := 0; fi32 < 18; fi32++ {
        switch fi32 {
        case 127:
            for fi33 := 0; fi33 < 11; fi33++ {
                v34 := GenS4{F0: (v2 != v2), F1: v21}
                v4[65535] = (fi33 + fi33)
                _ = fi33
                _ = v34
            }
        }
        v35 := GenS4{F0: true, F1: ("foo bar" + v21)}
        _ = fi32
        _ = v35
    }
    if (true || (!(v5 < 0))) {
        v3 += v3
        switch -2147483648 {
        case 1:
            for fi36 := 0; fi36 < 15; fi36++ {
                v23 += (v23 ^ int16(fi36))
                _ = len(v4)
                _ = fi36
            }
        }
        var v37 int8 = 1
        _ = v37
    } else {
        for fi38 := 0; fi38 < 3; fi38++ {
            var v39 uint32 = 506
            var v40 uint64 = 42
            for fi41 := 0; fi41 < 12; fi41++ {
                v6 = "hello"
                if (571 != 0) {
                    v2 += v2
                } else {
                    switch (v6 + v6) {
                    case "foo bar":
                        var v42 int16 = v23
                        _ = v42
                    case "":
                        v43 := v4[(v39 - v39)]
                        _ = v43
                    }
                    v44 := GenS1{F0: (v2 >> ((895) & 63))}
                    v3 -= 127
                    _ = v44
                }
                _ = fi41
            }
            _ = fi38
            _ = v39
            _ = v40
        }
        var v45 int = int(v17)
        _ = v45
    }
    var v46 float64 = ((1e10 * 3.14159) - 0.0)
    v47 := v4[4294967295]
    if true {
        switch ((v6 + v6) + ((v6 + v6) + (v21 + "a"))) {
        case "hello":
            var v48 int32 = -2147483648
            _ = v48
        }
    } else {
        for fi49 := 0; fi49 < 6; fi49++ {
            v50 := []int{}
            v50 = append(v50, 155)
            if len(v50) >= 2 {
                v50 = v50[1:len(v50)]
            }
            _ = fi49
            _ = v50
        }
        var v51 string = v21
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
