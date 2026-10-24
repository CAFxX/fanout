package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint32, p1 uint64) string {
    v5 := []int{-2147483648}
    if (!(1.5 != 3.14159)) {
        if true {
            p0 = p0
            var v6 uint8 = 2
            p0 += (p0 << ((0) & 63))
            _ = v6
        }
    }
    var v7 int8 = 127
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return ("" + "")
}

func main() {
    var v1 int8 = -74
    var v2 string = ""
    var v3 uint64 = 1
    var v4 int8 = (v1 & (-(v1)))
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: -2147483648}
        v3 = (v3 >> ((1) & 63))
        if (1 > (^(v3))) {
            var v10 uint16 = 1
            var v11 int64 = 0
            _ = v10
            _ = v11
        } else {
            v12 := []int{fi8, -909}
            for fi13 := 0; fi13 < 1; fi13++ {
                v1 -= (-(v1))
                v4 = v4
                _ = fi13
            }
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := []int{(fi14 * (^(fi14))), -9223372036854775808, fi14}
        v1 = (v1 - (v1 >> ((0) & 63)))
        _ = fi14
        _ = v15
    }
    v1 = (v4 | 100)
    if true {
        switch -9223372036854775808 {
        case 1:
            v4 += -1
        case 32768:
            if false {
                var v16 uint16 = 65535
                v17 := []int{128}
                _ = v16
                _ = v17
            }
        }
    } else {
        for fi18 := 0; fi18 < 18; fi18++ {
            v3 = v3
            v19 := GenS2{F0: 0, F1: 0}
            _ = fi18
            _ = v19
        }
    }
    v4 = (^(v4))
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS1{F0: 459}
        v22 := GenS3{F0: -638, F1: 236}
        if (v1 > v4) {
            var v23 int32 = 2147483647
            v24 := []int{}
            _ = v23
            _ = v24
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        v3 = (^(v3))
        var v26 int32 = 0
        switch -561 {
        case 32767:
            v3 = 4294967295
        case -9223372036854775808:
            if (false || (false || (false || ((!true) && true)))) {
                for fi27 := 0; fi27 < 6; fi27++ {
                    if (!false) {
                        v26 += (v26 & (-(v26)))
                        v28 := GenS1{F0: (v26 + (v26 % ((v26) | 1)))}
                        v2 = "a"
                        _ = v28
                    }
                    v29 := GenS1{F0: int32(fi27)}
                    v30 := []int{-1, (fi27 / (((fi27 + fi27)) | 1))}
                    _ = fi27
                    _ = v29
                    _ = v30
                }
                var v31 int8 = (-(v1))
                _ = v31
            }
        }
        _ = fi25
        _ = v26
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        v2 = (v2 + v2)
        var v33 float32 = 1.5
        v34 := []int{fi32, (fi32 - fi32), (-(fi32))}
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 uint = 18446744073709551615
    var v36 int = 2147483647
    if (2147483647 != -1) {
        v36 -= v36
        var v37 int64 = 42
        v38 := GenS1{F0: 417}
        _ = v37
        _ = v38
    } else {
        switch v36 {
        case 1:
            var v39 int = (-(v36))
            _ = v39
        case 128:
            v40 := GenS4{F0: 0}
            _ = v40
        case 888:
            v4 += (v1 >> ((v35) & 63))
        }
        var v41 int32 = -1
        _ = v41
    }
    fmt.Println(v3)
    fmt.Println(v2, v1)
    switch "\n\t" {
    case "hello":
        switch (v36 ^ int(v35)) {
        case 599:
            v42 := []int{(v36 % ((-2147483648) | 1)), v36, v36}
            _ = v42
        case -982:
            var v43 int = v36
            _ = v43
        default:
            var v44 int64 = 759
            _ = v44
        }
    case "\n\t":
        switch v2 {
        case "hello":
            v45 := GenS1{F0: 2147483647}
            _ = v45
        }
    case "a":
        var v46 uint64 = (v3 * (v3 & 385))
        _ = v46
    }
    v47 := GenS5{F0: (v35 * 0)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v47
}
