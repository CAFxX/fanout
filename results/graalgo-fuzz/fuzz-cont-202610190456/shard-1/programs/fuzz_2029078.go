package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int8 }
func genFn0() int {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 string = "a"
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: 0, F1: 255}
            var v11 string = "\n\t"
            _ = fi9
            _ = v10
            _ = v11
        }
        switch 128 {
        case 32768:
            var v12 uint32 = 65535
            _ = v12
        case -1:
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := GenS1{F0: 0, F1: 328}
                _ = fi13
                _ = v14
            }
        }
        _ = fi7
        _ = v8
    }
    var v15 int16 = -1
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 uint32 = 843
        _ = fi16
        _ = v17
    }
    switch -2147483648 {
    case 32767:
        v18 := []int{}
        _ = v18
    case -2147483648:
        v19 := GenS1{F0: 18446744073709551615, F1: 42}
        _ = v19
    case 630:
        v20 := []int{1, -9223372036854775808}
        _ = v20
    default:
        v21 := []int{128}
        _ = v21
    }
    _ = v15
    return -9223372036854775808
}
func genFn1() uint8 {
    switch 257 {
    case 1:
        var v22 uint = 606
        _ = v22
    default:
        var v23 string = ""
        _ = v23
    }
    var v24 uint64 = 435
    v24 -= 1
    _ = v24
    return 174
}

func main() {
    var v1 int64 = 42
    var v2 uint32 = 483
    var v3 int = -456
    var v4 uint64 = 4294967295
    var v5 bool = false
    v6 := make(map[int]int)
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := v6[genFn0()]
        _ = fi25
        _ = v26
    }
    v2 = v2
    v3 = v3
    fmt.Println(v2, v1)
    var v27 bool = false
    var v28 uint8 = 128
    if (false && v5) {
        if (((0 < 789) && ((v2 << ((4294967295) & 63)) == (^(v2)))) && true) {
            fmt.Println(v3)
            var v29 int32 = 473
            _ = v29
        } else {
            var v30 int = (-(v3))
            switch ("a" + "foo bar") {
            case "hello":
                var v31 float32 = -2.25
                _ = v31
            case "":
                fmt.Println(v5, v4)
            default:
                var v32 int = (-(v3))
                _ = v32
            }
            _ = v30
        }
    }
    v33 := GenS1{F0: 71, F1: 255}
    v34 := GenS1{F0: 18446744073709551615, F1: 329}
    var v35 int64 = -1
    if v5 {
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 string = ("foo bar" + "")
            var v38 string = ("a" + (v37 + ""))
            if ("" >= (v38 + (v38 + (((("a" + ("foo bar" + v38)) + "foo bar") + "foo bar") + v38)))) {
                var v39 float32 = float32(-2.25)
                var v40 string = (v37 + "\n\t")
                _ = v39
                _ = v40
            }
            _ = fi36
            _ = v37
            _ = v38
        }
    } else {
        v6[v3] = 4294967295
        v28 += 128
    }
    for fi41 := 0; fi41 < 3; fi41++ {
        var v42 uint = 1
        _ = fi41
        _ = v42
    }
    if true {
        v43 := []int{v3, (v3 - int(v4)), (-(v3))}
        for fi44 := 0; fi44 < 6; fi44++ {
            v45 := GenS2{F0: -128}
            _ = fi44
            _ = v45
        }
        _ = v43
    } else {
        v46 := []int{}
        for fi47 := 0; fi47 < 2; fi47++ {
            v48 := GenS1{F0: 42, F1: 2}
            _ = fi47
            _ = v48
        }
        var v49 float64 = float64((float32(float64(3.14159)) + 1.5))
        _ = v46
        _ = v49
    }
    v6[(v3 + genFn0())] = genFn0()
    for fi50 := 0; fi50 < 2; fi50++ {
        v51 := []int{-383, v3, (fi50 ^ (^(fi50))), v3}
        v3 = (fi50 - (v3 + fi50))
        _ = fi50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
