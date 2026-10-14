package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int64; F1 int; F2 int16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS5 struct { F0 int16; F1 int }
func genFn0() uint8 {
    v5 := GenS1{F0: 1, F1: 2}
    if ("hello" < "hello") {
        var v6 uint16 = 367
        switch -862 {
        case 128:
            var v7 int64 = 1
            _ = v7
        case 9223372036854775807:
            var v8 uint = 368
            _ = v8
        }
        _ = v6
    } else {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint16 = 65535
            _ = fi9
            _ = v10
        }
    }
    var v11 int16 = 40
    for fi12 := 0; fi12 < 16; fi12++ {
        if true {
            var v13 float32 = float32(0.1)
            v14 := []int{9223372036854775807, -786, fi12}
            if ((true && false) || true) {
                var v15 int16 = v11
                v11 = (-(v15))
                fmt.Println(v13, v13)
                _ = v15
            } else {
                var v16 int16 = (-(v11))
                v17 := GenS2{F0: -954, F1: (-(fi12)), F2: -32768}
                _ = v16
                _ = v17
            }
            _ = v13
            _ = v14
        } else {
            for fi18 := 0; fi18 < 2; fi18++ {
                v11 = v11
                if true {
                    var v19 uint16 = 255
                    switch fi18 {
                    case 2147483647:
                        v20 := []int{fi18}
                        _ = v20
                    case 127:
                        v21 := []int{fi18}
                        _ = v21
                    case 64:
                        var v22 uint16 = (^(v19))
                        _ = v22
                    }
                    _ = v19
                }
                var v23 int64 = 606
                _ = fi18
                _ = v23
            }
            fmt.Println(fi12, v11)
            fmt.Println(fi12, fi12)
        }
        var v24 uint16 = 949
        switch "\n\t" {
        case "hello":
            v25 := []int{-883, 4294967295}
            _ = v25
        }
        _ = fi12
        _ = v24
    }
    switch 2147483647 {
    case 32767:
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 uint16 = 2
            _ = fi26
            _ = v27
        }
    }
    _ = v11
    _ = v5
    return 127
}

func main() {
    var v1 int16 = 422
    v2 := make(map[string]int)
    var v3 uint64 = 0
    var v4 int32 = 0
    var v28 uint16 = 42
    if true {
        v1 += 42
        v4 -= v4
        var v29 uint16 = (^(v28))
        _ = v29
    } else {
        var v30 bool = (361 <= 4294967295)
        _ = v30
    }
    var v31 int = -412
    if (0 >= 65535) {
        for fi32 := 0; fi32 < 10; fi32++ {
            for fi33 := 0; fi33 < 7; fi33++ {
                switch "a" {
                case "hello":
                    var v34 string = ("foo bar" + "a")
                    _ = v34
                case "\n\t":
                    _ = len(v2)
                }
                _ = fi33
            }
            v35 := GenS3{F0: 2, F1: ("a" + ("\n\t" + "hello"))}
            v28 += v28
            _ = fi32
            _ = v35
        }
        var v36 uint16 = v28
        v37 := []int{(v31 >> ((411) & 63)), -1}
        _ = v36
        _ = v37
    } else {
        delete(v2, "hello")
        var v38 int32 = 42
        if false {
            var v39 float32 = (1e10 * -2.25)
            _ = len(v2)
            v1 -= v1
            _ = v39
        } else {
            v40 := v2["a"]
            v28 -= v28
            _ = v40
        }
        _ = v38
    }
    var v41 int16 = (v1 + v1)
    for fi42 := 0; fi42 < 14; fi42++ {
        v1 -= v41
        _ = fi42
    }
    v41 -= v1
    fmt.Println(v31, v3)
    var v43 uint16 = v28
    for fi44 := 0; fi44 < 8; fi44++ {
        fmt.Println(v3, v31, v43)
        _ = fi44
    }
    var v45 int16 = -1
    var v46 uint32 = 4294967295
    v4 = 2147483647
    for fi47 := 0; fi47 < 16; fi47++ {
        for fi48 := 0; fi48 < 14; fi48++ {
            switch (("foo bar" + "a") + ("" + "")) {
            case "a":
                var v49 int16 = (v1 + 32767)
                _ = v49
            case "foo bar":
                v41 = v41
            case "\n\t":
                var v50 string = "foo bar"
                v28 -= (v28 & (v43 - (^(v28))))
                _ = v50
            default:
                v4 = v4
            }
            _ = fi48
        }
        v51 := v2["foo bar"]
        if true {
            v52 := GenS4{F0: (v46 ^ 1), F1: -1, F2: v3}
            _ = v52
        }
        _ = fi47
        _ = v51
    }
    if (0 == 42) {
        v1 -= (v45 | (-(v45)))
        v53 := GenS5{F0: v41, F1: (v31 >> ((26) & 63))}
        v3 = (v3 / ((v3) | 1))
        _ = v53
    }
    var v54 uint8 = 0
    var v55 uint8 = (v54 & v54)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v41
    _ = v43
    _ = v45
    _ = v46
    _ = v54
    _ = v55
}
