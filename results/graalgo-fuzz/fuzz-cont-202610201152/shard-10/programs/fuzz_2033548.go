package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 string; F2 int }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 int }
func genFn0(p0 int16, p1 uint8) uint8 {
    var v6 string = "hello"
    p0 += p0
    v7 := []int{-315, 128, 32767}
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 float64 = -2.25
        v10 := GenS1{F0: 4294967295}
        switch int(p1) {
        case -9223372036854775808:
            var v11 float32 = -2.25
            _ = v11
        case 245:
            if len(v7) > 0 {
                v7[0] = fi8
            }
        case 0:
            var v12 float32 = -2.25
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 | (p1 << ((uint(p1)) & 63)))
}
func genFn1(p0 int64, p1 uint16) int16 {
    p1 = 42
    var v13 float64 = float64(1e10)
    fmt.Println(p0, p0)
    var v14 int16 = 32767
    if true {
        var v15 uint16 = 1
        v16 := GenS1{F0: 4294967295}
        var v17 string = "\n\t"
        _ = v15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return (v14 / (((v14 | v14)) | 1))
}

func main() {
    v1 := []int{42}
    var v2 int32 = -279
    var v3 float64 = (3.14159 * 1e10)
    var v4 int32 = v2
    var v5 uint32 = 42
    if (true && (!false)) {
        v5 = v5
        if true {
            fmt.Println(v5, v2, v3)
            var v18 int32 = (v4 >> ((245) & 63))
            var v19 int8 = -72
            _ = v18
            _ = v19
        }
        v20 := GenS2{F0: 2, F1: ("hello" + ("\n\t" + "\n\t")), F2: -854}
        _ = v20
    } else {
        var v21 int16 = -32768
        var v22 uint32 = v5
        var v23 uint = 1
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := GenS3{F0: ("hello" + (("foo bar" + (("a" + ("a" + "")) + "foo bar")) + ("hello" + ("a" + "\n\t")))), F1: 48}
    for fi25 := 0; fi25 < 14; fi25++ {
        v4 = -2147483648
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 bool = false
            switch (fi25 - fi26) {
            case 128:
                v4 += (v4 / (((v4 ^ int32(fi25))) | 1))
            }
            if v27 {
                var v28 int64 = 9223372036854775807
                var v29 string = "a"
                v3 = 1e10
                _ = v28
                _ = v29
            }
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    switch ("\n\t" + ("foo bar" + (((("" + ("hello" + "foo bar")) + "a") + "a") + "\n\t"))) {
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "a":
        v30 := GenS2{F0: 668, F1: "\n\t", F2: 4294967295}
        _ = v30
    default:
        v1 = append(v1, -639)
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 uint16 = 65535
        _ = fi31
        _ = v32
    }
    v33 := GenS4{F0: 4294967295}
    switch 32768 {
    case 32767:
        if len(v1) > 0 {
            v1[0] = -32
        }
    case 127:
        v34 := GenS5{F0: 2, F1: 0}
        _ = v34
    case 4294967295:
        v3 += (v3 / v3)
    }
    fmt.Println(v5, v3, v4)
    var v35 string = "a"
    _ = len(v1)
    for fi36 := 0; fi36 < 5; fi36++ {
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 string = (((v35 + "\n\t") + (v35 + v35)) + ((("hello" + ((v35 + v35) + (v35 + "a"))) + v35) + ((v35 + "") + v35)))
            _ = fi37
            _ = v38
        }
        var v39 int8 = -128
        _ = fi36
        _ = v39
    }
    switch 127 {
    case 32768:
        switch -2147483648 {
        case 351:
            var v40 int16 = 704
            _ = v40
        default:
            if (true && (0.1 <= (v3 + (float64(-2.25) * (0.1 * 1e10))))) {
                v4 += (v2 + v4)
            }
        }
    case -9223372036854775808:
        v41 := GenS4{F0: 232}
        _ = v41
    }
    _ = len(v1)
    if (genFn1(-9223372036854775808, 67) >= -72) {
        if false {
            switch 1 {
            case 0:
                v4 += v4
            case 32768:
                var v42 int32 = v2
                _ = v42
            case 1:
                if len(v1) > 0 {
                    v1[0] = 1
                }
            }
        }
        fmt.Println(v35, v35)
        for fi43 := 0; fi43 < 3; fi43++ {
            v44 := GenS4{F0: 65535}
            _ = fi43
            _ = v44
        }
    } else {
        var v45 int = -631
        v46 := GenS4{F0: 1}
        _ = v45
        _ = v46
    }
    var v47 float32 = ((float32(1e10) + (((float32(3.14159) / (-2.25 / 1.0)) + 0.1) * (((0.1 + ((0.1 + float32((v3 - 0.0))) / 1.0)) * -2.25) + (-2.25 / (1e10 - -2.25))))) + ((0.1 * 3.14159) - -2.25))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v47
    _ = v5
}
