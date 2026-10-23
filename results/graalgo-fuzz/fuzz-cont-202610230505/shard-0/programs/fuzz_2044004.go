package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint16; F1 string; F2 int32 }
type GenS6 struct { F0 int16; F1 int64; F2 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint, p1 int, p2 uint16) int32 {
    var v5 uint = (p0 << (((^(p0))) & 63))
    var v6 uint32 = 945
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 int8 = 100
    var v2 int = 296
    var v3 uint16 = 32767
    var v4 int8 = (v1 - (-(v1)))
    v7 := GenS1{F0: 65535}
    v4 = v4
    fmt.Println(v2, v2)
    switch (v2 + v2) {
    case 4294967295:
        switch ("foo bar" + "hello") {
        case "\n\t":
            var v8 int8 = (v1 << ((42) & 63))
            _ = v8
        case "":
            v9 := []int{v2, v2}
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 7; fi10++ {
            v1 = v1
            for fi11 := 0; fi11 < 6; fi11++ {
                v12 := GenS2{F0: -780}
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    if (!false) {
        v4 -= -64
    } else {
        switch "\n\t" {
        case "foo bar":
            v13 := []int{v2}
            _ = v13
        case "":
            v14 := []int{v2, v2}
            _ = v14
        case "\n\t":
            fmt.Println(v4)
        }
    }
    switch (v2 >> ((655) & 63)) {
    case -9223372036854775808:
        fmt.Println(v2)
    case 782:
        if false {
            switch "\n\t" {
            case "":
                if (255 <= 0) {
                    v4 += v4
                    fmt.Println(v2)
                }
            case "foo bar":
                v15 := GenS1{F0: 419}
                _ = v15
            case "hello":
                v16 := []int{(v2 & 128), (^(v2)), v2}
                _ = v16
            }
            v17 := GenS1{F0: 36}
            for fi18 := 0; fi18 < 8; fi18++ {
                v19 := []int{(fi18 % (((fi18 % ((fi18) | 1))) | 1)), (v2 >> ((159) & 63)), v2, 128}
                v20 := GenS1{F0: 65535}
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v17
        } else {
            v21 := []int{}
            var v22 float32 = -2.25
            fmt.Println(v22, v3, v4)
            _ = v21
            _ = v22
        }
    case -82:
        for fi23 := 0; fi23 < 3; fi23++ {
            if false {
                var v24 uint32 = 2
                var v25 int32 = 1
                _ = v24
                _ = v25
            }
            v26 := GenS3{F0: -618, F1: 128}
            if false {
                v27 := []int{(fi23 << ((2) & 63)), 9223372036854775807, 1}
                var v28 int64 = 1
                _ = v27
                _ = v28
            } else {
                v29 := []int{v2, 4294967295, fi23, (fi23 & (fi23 >> ((0) & 63)))}
                _ = v29
            }
            _ = fi23
            _ = v26
        }
    default:
        var v30 int64 = 1
        _ = v30
    }
    var v31 uint32 = 697
    v32 := GenS2{F0: -1}
    var v33 string = ("\n\t" + (("" + ((("" + "hello") + "") + "\n\t")) + ("foo bar" + "foo bar")))
    if (2147483647 < 1) {
        v34 := GenS4{F0: 1000, F1: 599}
        v33 = v33
        _ = v34
    } else {
        var v35 uint = 2
        v36 := GenS2{F0: 42}
        _ = v35
        _ = v36
    }
    v31 += (v31 >> ((uint(v31)) & 63))
    if (true && false) {
        fmt.Println(v2)
        v4 = v4
    } else {
        var v37 string = v33
        v38 := GenS5{F0: (^(v3)), F1: "\n\t", F2: 0}
        _ = v37
        _ = v38
    }
    v31 = (^(v31))
    if ((v3 >> ((939) & 63)) >= 733) {
        v39 := GenS5{F0: (v3 >> ((255) & 63)), F1: (v33 + ((v33 + v33) + ((v33 + v33) + (v33 + ("" + "a"))))), F2: -1}
        v40 := GenS6{F0: 490, F1: 9223372036854775807, F2: 128}
        switch 32768 {
        case 0:
            for fi41 := 0; fi41 < 8; fi41++ {
                if true {
                    for fi42 := 0; fi42 < 2; fi42++ {
                        var v43 int = fi42
                        _ = fi42
                        _ = v43
                    }
                    switch (^(v2)) {
                    case 0:
                        v44 := []int{-378}
                        _ = v44
                    case 32767:
                        fmt.Println(v33, v33)
                    }
                }
                _ = fi41
            }
        case -821:
            v31 = v31
        case 127:
            v45 := GenS7{F0: 0}
            _ = v45
        default:
            var v46 float64 = 1.5
            _ = v46
        }
        _ = v39
        _ = v40
    }
    for fi47 := 0; fi47 < 8; fi47++ {
        fmt.Println(v2)
        for fi48 := 0; fi48 < 13; fi48++ {
            v49 := GenS8{F0: 0, F1: 2}
            _ = fi48
            _ = v49
        }
        _ = fi47
    }
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v7
}
