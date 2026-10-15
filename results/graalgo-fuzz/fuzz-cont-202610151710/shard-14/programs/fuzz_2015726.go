package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int8; F2 int8 }
func genFn0() uint16 {
    if false {
        var v4 bool = ((!(!(100 != -71))) || true)
        var v5 uint64 = 502
        _ = v4
        _ = v5
    }
    switch 127 {
    case 250:
        var v6 int = 0
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        switch ("" + (("\n\t" + "\n\t") + ((("hello" + "foo bar") + "\n\t") + ("hello" + "foo bar")))) {
        case "hello":
            v8 := []int{fi7, 42, (-(fi7))}
            _ = v8
        case "":
            v9 := GenS1{F0: 2147483648, F1: 288}
            _ = v9
        default:
            var v10 uint = 772
            _ = v10
        }
        v11 := []int{}
        _ = fi7
        _ = v11
    }
    var v12 float32 = 0.1
    _ = v12
    return 65535
}
func genFn1() uint32 {
    var v13 uint = 0
    var v14 uint32 = 65535
    var v15 int16 = 833
    _ = v13
    _ = v14
    _ = v15
    return (v14 - (^(v14)))
}

func main() {
    var v1 uint64 = 864
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    switch "hello" {
    case "a":
        v16 := GenS1{F0: 75, F1: v1}
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS1{F0: 577, F1: v1}
        var v19 string = ("a" + (("\n\t" + "foo bar") + "foo bar"))
        v19 = "foo bar"
        _ = fi17
        _ = v18
        _ = v19
    }
    switch 2147483647 {
    case 647:
        var v20 int8 = 127
        _ = v20
    case 42:
        var v21 bool = (828 != 1)
        _ = v21
    case 16:
        var v22 float64 = 0.0
        _ = v22
    default:
        if (!false) {
            var v23 uint16 = 2
            _ = v23
        } else {
            for fi24 := 0; fi24 < 10; fi24++ {
                fmt.Println(fi24, v1)
                _ = fi24
            }
        }
    }
    if false {
        if true {
            v1 += v1
            var v25 uint32 = 0
            _ = v25
        } else {
            if false {
                var v26 int16 = 32767
                _ = v26
            }
            switch ("a" + "\n\t") {
            case "hello":
                fmt.Println(v1)
            case "\n\t":
                var v27 uint16 = 396
                _ = v27
            case "":
                if true {
                    v1 = (v1 % (((^(v1))) | 1))
                    v28 := GenS1{F0: 1, F1: (v1 >> ((42) & 63))}
                    _ = v28
                }
            default:
                fmt.Println(v1)
            }
            var v29 int = 32768
            _ = v29
        }
        if false {
            switch -853 {
            case -2147483648:
                fmt.Println(v1)
            default:
                var v30 uint64 = (v1 << ((65535) & 63))
                _ = v30
            }
            fmt.Println(v1)
        } else {
            var v31 uint64 = v1
            v32 := []int{-1, 42, -9223372036854775808}
            v33 := GenS1{F0: 0, F1: v31}
            _ = v31
            _ = v32
            _ = v33
        }
    } else {
        for fi34 := 0; fi34 < 15; fi34++ {
            v35 := []int{(fi34 * (fi34 & fi34)), (fi34 % ((32768) | 1))}
            _ = fi34
            _ = v35
        }
        v36 := []int{4294967295, 1}
        fmt.Println(v1)
        _ = v36
    }
    v37 := GenS1{F0: 42, F1: v1}
    if false {
        delete(v2, 1)
    } else {
        v1 += 799
        switch 128 {
        case 2147483647:
            v38 := GenS1{F0: 460, F1: v1}
            _ = v38
        default:
            var v39 int64 = 9223372036854775807
            _ = v39
        }
        var v40 uint32 = 2147483647
        _ = v40
    }
    var v41 uint8 = 255
    switch 32768 {
    case 9223372036854775807:
        v1 += v1
    case 32768:
        for fi42 := 0; fi42 < 8; fi42++ {
            var v43 int8 = -128
            _ = fi42
            _ = v43
        }
    case 4294967295:
        switch -425 {
        case 542:
            v44 := GenS1{F0: genFn1(), F1: v1}
            _ = v44
        case 127:
            v45 := GenS2{F0: 1, F1: int8(v41), F2: 66}
            _ = v45
        case 552:
            v1 = (^(v1))
        default:
            fmt.Println(v1, v41)
        }
    }
    fmt.Println(v41, v1)
    if (!false) {
        var v46 uint64 = 42
        v47 := GenS1{F0: 407, F1: (v1 * (v46 << ((1) & 63)))}
        var v48 float64 = 1.5
        _ = v46
        _ = v47
        _ = v48
    }
    v1 = (v1 << ((2) & 63))
    v49 := []int{-1, 128}
    v1 -= 1
    if (127 < -113) {
        v50 := GenS2{F0: 0, F1: -22, F2: -61}
        _ = v50
    }
    var v51 int8 = 66
    var v52 int = -9223372036854775808
    var v53 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v41
    _ = v49
    _ = v51
    _ = v52
    _ = v53
}
