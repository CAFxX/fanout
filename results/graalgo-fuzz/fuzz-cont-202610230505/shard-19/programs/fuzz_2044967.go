package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 bool; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint32 }
func genFn0() uint16 {
    for fi4 := 0; fi4 < 15; fi4++ {
        if false {
            var v5 int64 = -1
            v5 -= 9223372036854775807
            switch "hello" {
            case "hello":
                v6 := GenS1{F0: (^(fi4))}
                _ = v6
            case "\n\t":
                if (243 <= -1) {
                    var v7 int16 = 32767
                    v8 := GenS1{F0: (^(fi4))}
                    for fi9 := 0; fi9 < 4; fi9++ {
                        v10 := []int{fi4, (fi4 & 0), (^(fi4))}
                        _ = fi9
                        _ = v10
                    }
                    _ = v7
                    _ = v8
                }
            }
            _ = v5
        }
        var v11 int64 = 42
        v11 -= (v11 + (v11 + (v11 + (v11 + (v11 % (((v11 << ((835) & 63))) | 1))))))
        _ = fi4
        _ = v11
    }
    var v12 float32 = (0.1 * float32(((0.0 / ((1.5 - float64(1e10)) + 3.14159)) * -2.25)))
    for fi13 := 0; fi13 < 16; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 float32 = 0.1
            if (185 <= -392) {
                v16 := GenS1{F0: (fi13 >> ((65535) & 63))}
                fmt.Println(fi14, fi14, v15)
                v12 += float32((((1e10 + 0.1) / 1.5) - ((1e10 * 0.1) * (0.1 * 0.0))))
                v17 := []int{}
                _ = v16
                _ = v17
            }
            v18 := GenS2{F0: true, F1: false, F2: 0}
            _ = fi14
            _ = v15
            _ = v18
        }
        _ = fi13
    }
    _ = v12
    return 0
}
func genFn1(p0 uint32, p1 string) uint16 {
    switch 128 {
    case 563:
        var v19 uint64 = 4294967295
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        if (0 != -1) {
            p1 = p1
            v21 := []int{32767, fi20, (fi20 & (-(fi20)))}
            _ = v21
        }
        var v22 string = ("foo bar" + ((p1 + p1) + "foo bar"))
        for fi23 := 0; fi23 < 16; fi23++ {
            fmt.Println(p1)
            _ = fi23
        }
        _ = fi20
        _ = v22
    }
    v24 := GenS3{F0: (-32768 != 32767)}
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := []int{}
        if len(v26) > 0 {
            v26[0] = fi25
        }
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v24
    return genFn0()
}

func main() {
    var v1 uint16 = 32767
    var v2 string = "hello"
    v3 := make(map[string]uint32)
    if ((942 >= 255) || (false || (((true && true) || false) || (42 <= 0)))) {
        v27 := GenS4{F0: 0, F1: 880, F2: 2147483648}
        var v28 uint32 = 42
        v29 := GenS4{F0: 42, F1: 2, F2: v28}
        _ = v27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        fmt.Println(v1, v1)
        var v31 float32 = (float32(1.5) - ((3.14159 - -2.25) * (3.14159 - (((3.14159 / 0.1) * (0.0 - (float32(3.14159) / 1.0))) / 0.1))))
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        for fi33 := 0; fi33 < 6; fi33++ {
            v2 = ("\n\t" + v2)
            _ = fi33
        }
        _ = fi32
    }
    v2 = v2
    for fi34 := 0; fi34 < 4; fi34++ {
        v2 = v2
        for fi35 := 0; fi35 < 18; fi35++ {
            var v36 int32 = 2147483647
            v37 := []int{-984, 525, (^(fi35)), 128}
            v38 := GenS1{F0: fi35}
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        _ = fi34
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        for fi40 := 0; fi40 < 1; fi40++ {
            v2 = "\n\t"
            _ = fi40
        }
        var v41 float32 = 0.0
        _ = fi39
        _ = v41
    }
    switch v2 {
    case "hello":
        v1 = genFn0()
    }
    fmt.Println(v2)
    for fi42 := 0; fi42 < 10; fi42++ {
        for fi43 := 0; fi43 < 16; fi43++ {
            v44 := []int{-9223372036854775808, 4294967295}
            fmt.Println(fi43, v1, fi43)
            v1 = v1
            _ = fi43
            _ = v44
        }
        var v45 string = v2
        for fi46 := 0; fi46 < 17; fi46++ {
            v45 = (v2 + v2)
            var v47 uint32 = 154
            _ = fi46
            _ = v47
        }
        _ = fi42
        _ = v45
    }
    if ((!true) || false) {
        for fi48 := 0; fi48 < 7; fi48++ {
            var v49 int64 = -637
            v50 := GenS2{F0: (!false), F1: (-764 > 1), F2: -35}
            _ = fi48
            _ = v49
            _ = v50
        }
    }
    delete(v3, "")
    for fi51 := 0; fi51 < 10; fi51++ {
        var v52 int = fi51
        v2 = (("foo bar" + ("\n\t" + "a")) + "a")
        _ = fi51
        _ = v52
    }
    v1 = v1
    v1 -= v1
    var v53 int32 = 1
    delete(v3, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
}
