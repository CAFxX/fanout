package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 string }
func genFn0(p0 int64) int64 {
    var v4 uint32 = 42
    fmt.Println(p0)
    p0 += -290
    var v5 bool = true
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1() int16 {
    if true {
        v6 := []int{1, 42, -2147483648, 9223372036854775807}
        if false {
            var v7 uint16 = 2
            _ = v7
        } else {
            v8 := GenS1{F0: 255, F1: ("foo bar" + ""), F2: 127}
            _ = v8
        }
        _ = v6
    } else {
        var v9 uint8 = 42
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 uint32 = 663
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: 515, F1: "", F2: 31}
        for fi14 := 0; fi14 < 17; fi14++ {
            if (696 != genFn0(665)) {
                var v15 float64 = 1e10
                _ = v15
            }
            switch (-(fi14)) {
            case -161:
                fmt.Println(fi10, fi14)
            case 127:
                switch (fi14 | fi14) {
                case 0:
                    v16 := GenS2{F0: 32767, F1: -394}
                    _ = v16
                case 128:
                    fmt.Println(fi14)
                }
            }
            v17 := GenS1{F0: 688, F1: "foo bar", F2: 42}
            _ = fi14
            _ = v17
        }
        _ = fi10
        _ = v13
    }
    v18 := GenS3{F0: "hello"}
    if true {
        var v19 int8 = -96
        var v20 float64 = 1.5
        _ = v19
        _ = v20
    } else {
        if false {
            for fi21 := 0; fi21 < 8; fi21++ {
                switch ("hello" + "\n\t") {
                case "\n\t":
                    var v22 int64 = 482
                    _ = v22
                case "foo bar":
                    for fi23 := 0; fi23 < 10; fi23++ {
                        var v24 float32 = 3.14159
                        var v25 float32 = 3.14159
                        v26 := []int{}
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                default:
                    var v27 uint8 = 128
                    _ = v27
                }
                var v28 uint8 = 128
                _ = fi21
                _ = v28
            }
            var v29 int16 = 0
            v30 := []int{int(v29), 2147483647, 32767}
            _ = v29
            _ = v30
        } else {
            var v31 uint64 = 42
            _ = v31
        }
        if false {
            v32 := []int{-1, 227, 335, 127}
            var v33 int8 = 127
            _ = v32
            _ = v33
        } else {
            var v34 bool = false
            _ = v34
        }
        var v35 float64 = (((0.1 / 0.1) - (1.5 * 1e10)) * (1e10 / 0.1))
        _ = v35
    }
    _ = v18
    return 868
}

func main() {
    var v1 uint8 = 42
    var v2 uint16 = uint16(v1)
    var v3 uint = 255
    v2 = v2
    for fi36 := 0; fi36 < 11; fi36++ {
        switch int(v3) {
        case 331:
            var v37 uint16 = (^(v2))
            _ = v37
        case 2147483647:
            v38 := GenS2{F0: int16(v1), F1: 218}
            _ = v38
        }
        _ = fi36
    }
    for fi39 := 0; fi39 < 20; fi39++ {
        v2 = v2
        fmt.Println(v3, v2)
        _ = fi39
    }
    v1 -= v1
    var v40 uint8 = (v1 + v1)
    switch ("" + "") {
    case "foo bar":
        v40 = v1
    }
    var v41 uint16 = (v2 * (v2 % (((^(v2))) | 1)))
    v2 -= (v2 << ((v3) & 63))
    v42 := []int{}
    v40 = (v1 ^ v40)
    v3 = v3
    for fi43 := 0; fi43 < 4; fi43++ {
        if false {
            if len(v42) > 0 {
                v42[0] = (fi43 % ((int(v3)) | 1))
            }
        } else {
            var v44 uint16 = (^(v2))
            _ = v44
        }
        _ = fi43
    }
    if len(v42) > 0 {
        v42[0] = 1
    }
    var v45 uint8 = (v1 | uint8(v3))
    switch -177 {
    case -9223372036854775808:
        v2 = (v41 % ((v2) | 1))
    case 128:
        fmt.Println(v2, v45)
    case 1:
        var v46 int8 = 1
        _ = v46
    default:
        v41 -= uint16(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v45
}
