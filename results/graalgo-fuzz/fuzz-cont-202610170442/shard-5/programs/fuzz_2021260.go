package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 string; F1 int32 }
func genFn0(p0 bool) int8 {
    p0 = (0 >= 128)
    if p0 {
        v4 := GenS1{F0: 255}
        _ = v4
    } else {
        switch ("foo bar" + ("" + "hello")) {
        case "\n\t":
            var v5 float32 = ((1e10 * 0.0) / 1.0)
            _ = v5
        }
    }
    v6 := []int{}
    p0 = false
    p0 = false
    _ = p0
    _ = v6
    return 42
}
func genFn1() int64 {
    v7 := GenS1{F0: 0}
    switch 128 {
    case -9223372036854775808:
        v8 := []int{2147483647, 1, 846}
        _ = v8
    case 4294967295:
        var v9 int = 0
        _ = v9
    case 9223372036854775807:
        var v10 int16 = 736
        _ = v10
    default:
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint32 = 94
            switch fi11 {
            case 0:
                fmt.Println(fi11, v12)
            case -2147483648:
                for fi13 := 0; fi13 < 1; fi13++ {
                    if false {
                        v14 := GenS1{F0: 595}
                        v15 := GenS1{F0: 18446744073709551615}
                        _ = v14
                        _ = v15
                    } else {
                        switch (fi11 << ((18446744073709551615) & 63)) {
                        case 0:
                            var v16 uint = 399
                            _ = v16
                        }
                    }
                    v17 := GenS2{F0: 255, F1: 1, F2: true}
                    v18 := GenS2{F0: 787, F1: 2, F2: (!false)}
                    fmt.Println(fi13)
                    _ = fi13
                    _ = v17
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 float32 = (-2.25 + 0.1)
                    _ = fi19
                    _ = v20
                }
            }
            var v21 int8 = 127
            _ = fi11
            _ = v12
            _ = v21
        }
    }
    var v22 int32 = 42
    v23 := GenS2{F0: 2, F1: 18446744073709551615, F2: false}
    _ = v22
    _ = v23
    _ = v7
    return -476
}

func main() {
    var v1 int32 = -1
    var v2 int8 = 100
    var v3 int32 = (v1 * v1)
    v24 := GenS2{F0: 4294967295, F1: 4294967295, F2: true}
    v1 = 377
    for fi25 := 0; fi25 < 13; fi25++ {
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := GenS2{F0: 0, F1: 18446744073709551615, F2: false}
            v28 := []int{452}
            v2 += v2
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    var v29 bool = ((!(464 <= 65535)) && (127 < 59))
    var v30 uint32 = 65535
    v31 := GenS2{F0: 65535, F1: 18446744073709551615, F2: v29}
    v2 -= genFn0(v29)
    switch -2147483648 {
    case -1:
        if v29 {
            v3 = (v3 >> ((2) & 63))
            switch ("a" + ("foo bar" + "")) {
            case "\n\t":
                v32 := []int{0, 127, 42, -454}
                _ = v32
            case "a":
                var v33 bool = v29
                _ = v33
            }
            v34 := []int{127}
            _ = v34
        } else {
            v1 = (v3 & v1)
            var v35 float64 = -2.25
            _ = v35
        }
    case 32768:
        var v36 float64 = (0.0 / 3.14159)
        _ = v36
    case -9223372036854775808:
        v37 := GenS3{F0: 255, F1: v29}
        _ = v37
    default:
        v38 := GenS4{F0: 47, F1: v2}
        _ = v38
    }
    if true {
        v39 := []int{128, -9223372036854775808, 9223372036854775807, -2147483648}
        var v40 int16 = 42
        _ = v39
        _ = v40
    }
    v30 -= v30
    fmt.Println(v3, v3)
    for fi41 := 0; fi41 < 1; fi41++ {
        switch 128 {
        case -898:
            v42 := []int{-936, fi41}
            _ = v42
        case -617:
            v43 := []int{-761}
            _ = v43
        case -2147483648:
            v44 := GenS5{F0: -220, F1: 1}
            _ = v44
        default:
            v2 = -5
        }
        var v45 int32 = (v3 & (v1 * 1))
        _ = fi41
        _ = v45
    }
    fmt.Println(v30)
    switch ((("" + ("" + "a")) + ("" + "hello")) + "hello") {
    case "foo bar":
        v46 := []int{9223372036854775807, 1, -793}
        _ = v46
    case "hello":
        for fi47 := 0; fi47 < 12; fi47++ {
            var v48 uint16 = 32767
            v49 := []int{(fi47 - fi47)}
            _ = len(v49)
            _ = fi47
            _ = v48
            _ = v49
        }
    default:
        var v50 string = (((("" + ("hello" + ((("foo bar" + "a") + "") + "\n\t"))) + "a") + "a") + ("" + "a"))
        _ = v50
    }
    switch -907 {
    case 128:
        v51 := []int{-2147483648, -9223372036854775808, 128, 9223372036854775807}
        _ = v51
    case -1:
        v52 := GenS1{F0: 4294967295}
        _ = v52
    case 891:
        switch -2147483648 {
        case 2147483647:
            v3 = (v3 / ((v1) | 1))
        case 4294967295:
            v2 = 127
        }
    }
    v53 := GenS6{F0: ("a" + "a"), F1: (v3 * (v1 % ((v3) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v53
}
