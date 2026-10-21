package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 string }
type GenS3 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0() int8 {
    var v6 float64 = (0.1 - (-2.25 / 1.5))
    var v7 float32 = (-2.25 + (float32(float64(-2.25)) - ((-2.25 + (1e10 * 1.5)) - float32(v6))))
    v8 := []int{32767, -644, 127}
    for fi9 := 0; fi9 < 17; fi9++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi9
    }
    var v10 uint64 = 42
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 bool) int8 {
    switch (("" + "") + "a") {
    case "hello":
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 string = ("hello" + "")
            v13 := GenS1{F0: 0, F1: 114, F2: 0}
            v12 = v12
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        var v14 int = 0
        _ = v14
    }
    p0 = false
    var v15 string = ""
    _ = p0
    _ = v15
    return 1
}

func main() {
    var v1 uint32 = 956
    v2 := []uint32{}
    var v3 uint = 784
    var v4 uint16 = 255
    v5 := []string{"foo bar", "a", "", ""}
    fmt.Println(v3, v1, v4)
    switch -717 {
    case -223:
        switch "hello" {
        case "foo bar":
            for fi16 := 0; fi16 < 12; fi16++ {
                v1 = (v1 << ((v3) & 63))
                if ((!true) && true) {
                    v4 = 1
                }
                v4 += 65535
                _ = fi16
            }
        }
    case 128:
        v5 = append(v5, ("a" + ("\n\t" + "foo bar")))
    case -380:
        switch "foo bar" {
        case "":
            var v17 uint8 = 128
            _ = v17
        case "foo bar":
            switch ("foo bar" + (((("" + "foo bar") + "") + "hello") + ("a" + "a"))) {
            case "\n\t":
                v18 := GenS1{F0: 0, F1: 0, F2: v1}
                _ = v18
            }
        case "a":
            if true {
                var v19 float64 = 0.0
                v19 += v19
                _ = v19
            } else {
                var v20 int64 = -9223372036854775808
                _ = v20
            }
        default:
            v4 = 0
        }
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        _ = fi21
    }
    if false {
        v22 := GenS1{F0: genFn0(), F1: 42, F2: 2147483648}
        v1 = 933
        v1 -= (^(v1))
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        v3 = (v3 | v3)
        v1 += (v1 & v1)
        v24 := GenS2{F0: 100, F1: v4, F2: ("" + "hello")}
        _ = fi23
        _ = v24
    }
    fmt.Println(v4)
    switch 2147483647 {
    case 0:
        var v25 string = "foo bar"
        _ = v25
    case 742:
        if len(v5) > 0 {
            v5[0] = ("foo bar" + ("a" + "a"))
        }
    case 128:
        var v26 uint32 = (v1 - v1)
        _ = v26
    default:
        v27 := GenS3{F0: 500, F1: 408, F2: (v4 & (v4 - 122))}
        _ = v27
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        for fi29 := 0; fi29 < 2; fi29++ {
            switch fi28 {
            case 0:
                v3 -= 2
            case 128:
                fmt.Println(fi29, fi29, v1)
            }
            v1 -= v1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi29
        }
        v4 = (v4 ^ (v4 ^ v4))
        if (!(int64(v1) < int64(fi28))) {
            for fi30 := 0; fi30 < 19; fi30++ {
                var v31 int = 4294967295
                if len(v5) > 0 {
                    v5[0] = ("\n\t" + "")
                }
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 3; fi32++ {
                v4 -= v4
                fmt.Println(v1, v3)
                _ = fi32
            }
            v1 = (v1 & (v1 / ((308) | 1)))
        } else {
            var v33 uint = 357
            _ = v33
        }
        _ = fi28
    }
    if (false && (true && false)) {
        for fi34 := 0; fi34 < 20; fi34++ {
            if len(v2) > 0 {
                v2[0] = v1
            }
            _ = fi34
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v4 = (^(v4))
    } else {
        var v35 uint = v3
        for fi36 := 0; fi36 < 5; fi36++ {
            fmt.Println(v4, fi36, v1)
            v1 -= 241
            _ = fi36
        }
        switch "\n\t" {
        case "":
            fmt.Println(v35, v35)
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = v35
    }
    v1 -= (v1 / ((v1) | 1))
    v37 := GenS4{F0: 32767, F1: 42}
    var v38 uint16 = 255
    if true {
        if true {
            var v39 uint32 = (v1 >> ((v3) & 63))
            _ = v39
        } else {
            var v40 int32 = int32(v1)
            fmt.Println(v1)
            v41 := GenS3{F0: 901, F1: 4294967295, F2: (v4 & 2)}
            _ = v40
            _ = v41
        }
    } else {
        var v42 float64 = 1e10
        switch "foo bar" {
        case "hello":
            v38 = (v4 - (^(v4)))
        case "\n\t":
            if false {
                v3 = v3
            } else {
                var v43 float64 = (v42 - v42)
                _ = v43
            }
        case "":
            if (182 >= 300) {
                for fi44 := 0; fi44 < 5; fi44++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi44
                }
            }
        default:
            if ((!true) || (!(false || (false && true)))) {
                var v45 float32 = 1e10
                _ = v45
            } else {
                var v46 uint64 = 1
                var v47 uint = (v3 * v3)
                _ = v46
                _ = v47
            }
        }
        v48 := GenS1{F0: -1, F1: genFn1((0 != -9223372036854775808)), F2: (v1 * uint32(v3))}
        _ = v42
        _ = v48
    }
    var v49 int8 = genFn0()
    var v50 int32 = 659
    switch "\n\t" {
    case "":
        for fi51 := 0; fi51 < 12; fi51++ {
            v38 += 0
            for fi52 := 0; fi52 < 20; fi52++ {
                v2 = append(v2, (v1 ^ (v1 / ((uint32(v4)) | 1))))
                _ = fi52
            }
            _ = fi51
        }
    default:
        fmt.Println(v38)
    }
    v3 = v3
    fmt.Println(v38, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v49
    _ = v5
    _ = v50
}
