package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint32; F2 bool }
type GenS5 struct { F0 uint64; F1 int64 }
func genFn0(p0 int8, p1 string, p2 uint8) int32 {
    if (128 != 1) {
        var v6 int16 = 873
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int = (fi7 | fi7)
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        p1 = p1
    }
    p0 = p0
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS1{F0: 2}
        var v11 float32 = -2.25
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 535
    var v3 uint32 = 65535
    var v4 uint = 42
    var v5 uint8 = 1
    var v12 float32 = -2.25
    switch ("foo bar" + "") {
    case "":
        if (42 > 42) {
            var v13 float32 = v12
            v5 = 35
            _ = v13
        } else {
            for fi14 := 0; fi14 < 13; fi14++ {
                fmt.Println(v5, v12, v4)
                var v15 float64 = (float64(v12) * (0.0 * ((((((3.14159 - float64(v12)) + 0.0) - float64(v12)) + 0.0) + -2.25) - -2.25)))
                var v16 uint64 = (^(v2))
                _ = fi14
                _ = v15
                _ = v16
            }
        }
    case "a":
        var v17 uint = uint(v5)
        _ = v17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            fmt.Println(v4, fi18)
            if (0 > 42) {
                v3 -= (v3 / ((v3) | 1))
            }
            v3 = 2147483648
            _ = fi19
        }
        var v20 bool = false
        _ = fi18
        _ = v20
    }
    fmt.Println(v2, v3, v5)
    if true {
        v2 += v2
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 int64 = 9223372036854775807
            _ = fi21
            _ = v22
        }
    }
    switch 9223372036854775807 {
    case 1:
        for fi23 := 0; fi23 < 9; fi23++ {
            v24 := GenS2{F0: (v2 % (((v2 >> (((v4 - (v4 | v4))) & 63))) | 1))}
            var v25 uint8 = (v5 << ((1) & 63))
            _ = fi23
            _ = v24
            _ = v25
        }
    case 32767:
        v4 += (v4 % (((^(v4))) | 1))
    case 974:
        v5 += (v5 | (v5 | v5))
    default:
        if true {
            fmt.Println(v4, v3)
            v26 := GenS3{F0: (v5 / ((v5) | 1)), F1: -51}
            var v27 uint8 = (v5 << ((688) & 63))
            _ = v26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        fmt.Println(v5, v5)
        var v29 uint64 = v2
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        delete(v1, fi30)
        _ = fi30
    }
    var v31 uint64 = uint64(v4)
    for fi32 := 0; fi32 < 16; fi32++ {
        if true {
            var v33 float64 = -2.25
            v12 += (v12 + 0.0)
            if true {
                v5 = (v5 & (v5 >> (((v4 * (v4 ^ (^(v4))))) & 63)))
                var v34 int64 = 971
                switch "\n\t" {
                case "a":
                    v35 := v1[fi32]
                    _ = v35
                case "":
                    _ = len(v1)
                case "\n\t":
                    var v36 int = int(v3)
                    _ = v36
                }
                _ = v34
            } else {
                var v37 float32 = (3.14159 + 1.5)
                v2 = 791
                fmt.Println(fi32, v3)
                _ = v37
            }
            _ = v33
        } else {
            v3 -= 4294967295
        }
        v38 := GenS4{F0: v5, F1: v3, F2: (v2 <= v2)}
        v39 := GenS2{F0: v2}
        _ = fi32
        _ = v38
        _ = v39
    }
    if ((v4 % ((1) | 1)) < (v4 + (^(v4)))) {
        v2 -= (^(v31))
        v40 := []int{-612, 224, 42}
        var v41 uint8 = v5
        _ = v40
        _ = v41
    } else {
        for fi42 := 0; fi42 < 7; fi42++ {
            fmt.Println(v5, v12)
            _ = fi42
        }
    }
    var v43 int32 = -2147483648
    for fi44 := 0; fi44 < 14; fi44++ {
        if false {
            fmt.Println(v4)
            if ((false || false) && true) {
                var v45 bool = (true || ((v43 << (((v4 << (((v4 % (((v4 * 0)) | 1))) & 63))) & 63)) < (v43 ^ genFn0(0, "", (^(v5))))))
                _ = v45
            }
            fmt.Println(v4)
        } else {
            v3 += (v3 % ((2) | 1))
            for fi46 := 0; fi46 < 10; fi46++ {
                v47 := []int{fi44, 852}
                var v48 string = ("foo bar" + "")
                v12 -= (3.14159 - v12)
                fmt.Println(v48, v5, v12)
                fmt.Println(v31, v2, v5)
                if false {
                    fmt.Println(v31, v12)
                    switch ((("" + v48) + "a") + v48) {
                    case "":
                        v43 = (-(v43))
                    default:
                        v49 := GenS2{F0: (v2 + (v31 & 1))}
                        v12 = v12
                        _ = v49
                    }
                }
                _ = fi46
                _ = v47
                _ = v48
            }
            var v50 uint64 = (v31 % (((v2 & (v31 + v31))) | 1))
            _ = v50
        }
        v3 = (v3 - v3)
        _ = fi44
    }
    var v51 uint64 = (v2 ^ v31)
    if ("hello" > ("hello" + "hello")) {
        if (592 < 32767) {
            delete(v1, 32768)
            switch 0 {
            case -518:
                v51 = v31
            case 2147483647:
                v52 := GenS1{F0: (v5 << ((v4) & 63))}
                _ = v52
            case -9223372036854775808:
                v51 = (^(v51))
            }
            switch 627 {
            case 32768:
                fmt.Println(v2, v31, v51)
            default:
                var v53 uint32 = (v3 | v3)
                _ = v53
            }
        } else {
            v31 = 42
            switch 32767 {
            case 891:
                var v54 uint = v4
                _ = v54
            }
        }
        switch 128 {
        case 127:
            fmt.Println(v5, v51)
        case 676:
            for fi55 := 0; fi55 < 5; fi55++ {
                v43 = v43
                delete(v1, fi55)
                v56 := GenS5{F0: (^(v2)), F1: 0}
                _ = fi55
                _ = v56
            }
        }
        v51 = uint64(v3)
    }
    fmt.Println(v51)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v43
    _ = v5
    _ = v51
}
