package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 int8; F2 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 int32; F2 uint64 }
type GenS8 struct { F0 uint16; F1 uint32; F2 string }

func main() {
    var v1 int = 9223372036854775807
    var v2 int16 = -707
    var v3 float32 = float32(0.0)
    v4 := make(map[uint32]int)
    var v5 string = ""
    var v6 int64 = 7
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := v4[0]
        delete(v4, 1)
        v9 := []int{(v8 - (v1 / ((v8) | 1))), (-(v1)), fi7, fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v3, v6)
    var v10 string = v5
    if ((!true) || false) {
        switch 948 {
        case 9223372036854775807:
            for fi11 := 0; fi11 < 10; fi11++ {
                v3 -= (v3 + v3)
                v3 += 1e10
                var v12 bool = (100 <= 0)
                _ = fi11
                _ = v12
            }
        default:
            fmt.Println(v10)
        }
        v13 := GenS1{F0: 1, F1: -1}
        _ = v13
    } else {
        var v14 uint64 = 42
        _ = v14
    }
    var v15 int32 = 0
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 11; fi17++ {
            fmt.Println(v3, fi16)
            var v18 int = (v1 * fi16)
            _ = fi17
            _ = v18
        }
        var v19 int16 = (-(v2))
        _ = fi16
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        v6 = v6
        if true {
            v10 = "foo bar"
            if (!false) {
                for fi21 := 0; fi21 < 5; fi21++ {
                    var v22 int16 = (-(v2))
                    _ = fi21
                    _ = v22
                }
            }
            v23 := GenS2{F0: 225, F1: v2}
            _ = v23
        } else {
            switch ((v5 + ("hello" + ("foo bar" + (("hello" + ("" + ((v5 + ((v5 + v10) + "\n\t")) + (((v5 + "hello") + (v10 + "hello")) + ("\n\t" + v5))))) + (((v10 + (("foo bar" + v10) + ("foo bar" + (v5 + v10)))) + "hello") + (("\n\t" + ((("\n\t" + (v10 + v5)) + (v10 + ("\n\t" + ("hello" + "hello")))) + "a")) + v10)))))) + "\n\t") {
            case "a":
                fmt.Println(v15, v6, v1)
            case "hello":
                var v24 int = (fi20 << ((2) & 63))
                _ = v24
            case "\n\t":
                var v25 uint = 65535
                _ = v25
            }
            v26 := GenS3{F0: (v6 + (v6 - (v6 * (v6 + v6))))}
            _ = v26
        }
        v27 := GenS2{F0: 2, F1: v2}
        _ = fi20
        _ = v27
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        v29 := []int{(v1 | fi28), (fi28 & v1)}
        _ = fi28
        _ = v29
    }
    if false {
        if false {
            fmt.Println(v1, v6, v5)
            var v30 float64 = 0.1
            _ = v30
        }
        if false {
            v5 = ""
            fmt.Println(v15, v6)
            v6 += (v6 ^ 0)
        } else {
            v31 := []int{127}
            v2 = (v2 | 32767)
            v3 += v3
            _ = v31
        }
    } else {
        v3 += float32(float64(((3.14159 - v3) / (((((float32(1.5) / (v3 * v3)) * 3.14159) + v3) - (1.5 / (((((-2.25 * v3) + -2.25) * float32(-2.25)) / v3) / v3))) / 1.5))))
        var v32 uint32 = 1
        v33 := GenS4{F0: 2}
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        for fi35 := 0; fi35 < 13; fi35++ {
            var v36 int32 = 394
            _ = fi35
            _ = v36
        }
        var v37 uint = 0
        v38 := []int{}
        _ = fi34
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        switch v10 {
        case "":
            v40 := GenS5{F0: 42, F1: 100, F2: false}
            _ = v40
        case "foo bar":
            var v41 uint8 = 1
            _ = v41
        }
        var v42 uint8 = 11
        for fi43 := 0; fi43 < 12; fi43++ {
            var v44 uint = 0
            v6 = (v6 >> (((v44 << (((v44 % ((v44) | 1))) & 63))) & 63))
            var v45 uint8 = v42
            v46 := GenS1{F0: 1, F1: -84}
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
        _ = fi39
        _ = v42
    }
    for fi47 := 0; fi47 < 4; fi47++ {
        var v48 uint16 = 42
        var v49 string = ("" + "a")
        _ = fi47
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 18; fi50++ {
        var v51 float32 = v3
        if (v51 >= v3) {
            if (-21 == 100) {
                v3 = v51
            } else {
                if ((v51 - -2.25) < (float32(-2.25) * 1e10)) {
                    var v52 int = v1
                    _ = v52
                } else {
                    var v53 int8 = 0
                    v54 := []int{fi50}
                    _ = len(v54)
                    _ = v53
                    _ = v54
                }
            }
        }
        v5 = v5
        _ = fi50
        _ = v51
    }
    for fi55 := 0; fi55 < 18; fi55++ {
        switch ("foo bar" + "a") {
        case "":
            v56 := []int{(fi55 & 32767), (-(v1)), fi55, (fi55 & fi55)}
            _ = v56
        case "\n\t":
            var v57 float64 = ((1e10 / 1e10) * 1e10)
            _ = v57
        }
        if true {
            v58 := GenS6{F0: 0}
            fmt.Println(v6)
            _ = v58
        }
        for fi59 := 0; fi59 < 7; fi59++ {
            var v60 bool = ((true && (0 <= v2)) || (fi55 == 42))
            v61 := GenS7{F0: v60, F1: -2147483648, F2: 0}
            var v62 uint8 = 33
            _ = fi59
            _ = v60
            _ = v61
            _ = v62
        }
        _ = fi55
    }
    var v63 string = "foo bar"
    for fi64 := 0; fi64 < 12; fi64++ {
        v65 := GenS8{F0: 42, F1: 738, F2: "hello"}
        _ = fi64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
}
