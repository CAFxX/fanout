package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 string }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 bool }
func genFn0() int {
    var v4 string = "foo bar"
    v4 = v4
    if ((2 == 397) || true) {
        switch v4 {
        case "foo bar":
            if false {
                v5 := GenS1{F0: 682, F1: (0 != 128), F2: ""}
                _ = v5
            }
        case "hello":
            fmt.Println(v4)
        case "a":
            fmt.Println(v4)
        }
    } else {
        switch 32767 {
        case 579:
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 bool = (0.1 <= -2.25)
                v8 := []int{fi6, (^(fi6))}
                _ = fi6
                _ = v7
                _ = v8
            }
        case -2147483648:
            v4 = "foo bar"
        case 0:
            v4 = v4
        default:
            if true {
                v9 := []int{457, 2147483647, -946}
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 int16 = 1000
                    _ = fi10
                    _ = v11
                }
                _ = v9
            }
        }
        if true {
            fmt.Println(v4)
        }
    }
    _ = v4
    return 32768
}

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{4294967295, 1, 487, 0}
    var v3 float64 = 3.14159
    v3 = 0.0
    if len(v2) > 0 {
        v2[0] = 65535
    }
    v12 := GenS1{F0: 734, F1: true, F2: ("" + ("foo bar" + "a"))}
    var v13 float32 = 1e10
    if false {
        if true {
            fmt.Println(v3, v1)
        } else {
            v2 = append(v2, 1)
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := GenS1{F0: 639, F1: (genFn0() != genFn0()), F2: "a"}
            v1 = v1
            v3 = v3
            _ = fi14
            _ = v15
        }
    }
    if false {
        fmt.Println(v3, v1, v3)
    } else {
        var v16 uint32 = 2147483647
        var v17 float64 = (float64((v13 - v13)) - float64(-2.25))
        for fi18 := 0; fi18 < 6; fi18++ {
            fmt.Println(v13, fi18, fi18)
            v2 = append(v2, v16)
            v19 := GenS2{F0: 0, F1: 134}
            _ = fi18
            _ = v19
        }
        _ = v16
        _ = v17
    }
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    v13 = float32(float64(3.14159))
    var v20 int32 = 2147483647
    var v21 uint32 = 2
    if true {
        fmt.Println(v13)
    } else {
        var v22 int = -201
        var v23 int8 = 1
        _ = v22
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = v21
    }
    if ("a" <= "hello") {
        var v24 bool = ((v20 * -301) <= (v20 % (((v20 % ((398) | 1))) | 1)))
        v25 := GenS3{F0: ((v20 > v20) || true)}
        _ = v24
        _ = v25
    } else {
        var v26 int8 = -57
        var v27 int8 = 100
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
