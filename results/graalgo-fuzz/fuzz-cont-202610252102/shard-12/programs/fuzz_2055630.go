package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int16; F2 int64 }
func genFn0() int {
    if true {
        var v5 uint32 = 65535
        for fi6 := 0; fi6 < 7; fi6++ {
            v5 += 2147483648
            var v7 string = "a"
            for fi8 := 0; fi8 < 13; fi8++ {
                fmt.Println(fi8, fi6, fi8)
                _ = fi8
            }
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        var v9 int16 = 856
        switch "" {
        case "hello":
            v10 := GenS1{F0: 32767}
            _ = v10
        }
        _ = v9
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        switch (fi11 >> ((255) & 63)) {
        case -533:
            if true {
                v12 := []int{993, 127, (fi11 ^ 32767), 9223372036854775807}
                var v13 uint64 = 2
                _ = v12
                _ = v13
            } else {
                var v14 float64 = 1.5
                var v15 float32 = 3.14159
                _ = v14
                _ = v15
            }
        }
        if (!false) {
            v16 := []int{(^(fi11)), (fi11 % (((fi11 - 19)) | 1)), fi11, fi11}
            var v17 string = ""
            fmt.Println(fi11, v17)
            _ = v16
            _ = v17
        } else {
            if (!(true && false)) {
                v18 := []int{6, 623, fi11, (fi11 + 42)}
                var v19 uint = 1
                _ = v18
                _ = v19
            }
            var v20 uint32 = 2147483648
            _ = v20
        }
        _ = fi11
    }
    if ((0.1 + ((float32(((float64(3.14159) / 3.14159) - 3.14159)) - ((1.5 / 1.0) + (0.1 / -2.25))) * (3.14159 * -2.25))) != (float32((0.1 - (0.0 * ((3.14159 / 0.1) + 3.14159)))) / 3.14159)) {
        var v21 string = "\n\t"
        if true {
            v22 := []int{1, 798, 4294967295, 628}
            v23 := GenS2{F0: 140}
            _ = v22
            _ = v23
        }
        _ = v21
    } else {
        for fi24 := 0; fi24 < 4; fi24++ {
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 uint16 = 32767
                v27 := GenS3{F0: true, F1: 1000, F2: 9223372036854775807}
                _ = fi25
                _ = v26
                _ = v27
            }
            v28 := []int{887, fi24, 32767, fi24}
            var v29 bool = (0 >= 42)
            _ = fi24
            _ = v28
            _ = v29
        }
        if false {
            var v30 int8 = 0
            _ = v30
        } else {
            v31 := GenS1{F0: 320}
            _ = v31
        }
        var v32 int64 = 267
        _ = v32
    }
    var v33 int16 = 32767
    _ = v33
    return 1
}
func genFn1(p0 int32) int16 {
    for fi34 := 0; fi34 < 18; fi34++ {
        v35 := []int{fi34, (fi34 ^ (fi34 / ((-1) | 1))), (fi34 << ((65535) & 63)), (fi34 ^ fi34)}
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        switch ("a" + "") {
        case "a":
            p0 = p0
        case "":
            p0 -= p0
        }
        for fi37 := 0; fi37 < 14; fi37++ {
            p0 -= p0
            var v38 uint8 = 159
            switch ("hello" + "") {
            case "hello":
                switch fi36 {
                case 250:
                    fmt.Println(fi36)
                default:
                    switch fi36 {
                    case 127:
                        if (!(2147483647 <= 680)) {
                            var v39 uint8 = v38
                            _ = v39
                        } else {
                            var v40 uint32 = 65535
                            v41 := []int{fi37, int(v38)}
                            _ = v40
                            _ = v41
                        }
                    case 2147483647:
                        var v42 int32 = int32(v38)
                        _ = v42
                    case 1:
                        fmt.Println(p0)
                    default:
                        v43 := GenS2{F0: v38}
                        _ = v43
                    }
                }
            case "a":
                v44 := []int{fi36}
                _ = v44
            }
            v45 := []int{1, (fi37 << ((428) & 63)), (fi36 / (((fi37 & fi37)) | 1))}
            var v46 uint32 = 996
            _ = fi37
            _ = v38
            _ = v45
            _ = v46
        }
        p0 += p0
        _ = fi36
    }
    for fi47 := 0; fi47 < 18; fi47++ {
        v48 := []int{genFn0(), 9223372036854775807, (fi47 / (((fi47 % ((-2147483648) | 1))) | 1))}
        fmt.Println(p0)
        _ = fi47
        _ = v48
    }
    _ = p0
    return -32768
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]string)
    var v3 bool = ((true || (18446744073709551615 > 266)) || false)
    var v4 uint8 = 218
    var v49 uint64 = 4294967295
    if (v3 || v3) {
        fmt.Println(v1, v4)
    }
    v50 := GenS1{F0: 538}
    for fi51 := 0; fi51 < 7; fi51++ {
        var v52 int32 = 62
        if true {
            fmt.Println(v1)
            v53 := v2[2]
            v52 += 1
            _ = v53
        }
        v1 -= v1
        _ = fi51
        _ = v52
    }
    var v54 uint16 = 818
    var v55 int8 = 127
    v56 := []int{-2147483648, -420, 128, -9223372036854775808}
    var v57 uint16 = v54
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v54
    _ = v55
    _ = v56
    _ = v57
}
