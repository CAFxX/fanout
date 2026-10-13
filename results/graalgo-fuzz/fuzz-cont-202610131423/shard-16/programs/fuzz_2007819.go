package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint8 }

func main() {
    var v1 uint16 = 2
    var v2 uint16 = v1
    var v3 int64 = 315
    var v4 uint16 = v2
    v5 := GenS1{F0: ("foo bar" + "hello")}
    if true {
        for fi6 := 0; fi6 < 11; fi6++ {
            fmt.Println(fi6)
            var v7 int64 = v3
            var v8 int = fi6
            _ = fi6
            _ = v7
            _ = v8
        }
        v9 := GenS2{F0: 127}
        _ = v9
    } else {
        v2 -= 22
    }
    v4 = 307
    if true {
        v10 := GenS2{F0: 128}
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        if true {
            v12 := []int{(fi11 >> ((18446744073709551615) & 63))}
            _ = v12
        }
        if false {
            v1 = (^(v4))
            var v13 float64 = ((((float64(3.14159) - 3.14159) / -2.25) / 1e10) + 1e10)
            var v14 int64 = (v3 - -723)
            _ = v13
            _ = v14
        } else {
            v4 = (v4 & 65535)
            var v15 uint16 = v4
            var v16 uint32 = 1
            _ = v15
            _ = v16
        }
        _ = fi11
    }
    var v17 float32 = 3.14159
    for fi18 := 0; fi18 < 13; fi18++ {
        fmt.Println(fi18, v4, v4)
        v1 = (v2 >> ((1) & 63))
        _ = fi18
    }
    if false {
        for fi19 := 0; fi19 < 6; fi19++ {
            v20 := []int{fi19, fi19, fi19}
            v17 = -2.25
            _ = fi19
            _ = v20
        }
        var v21 uint = 1
        _ = v21
    }
    switch 32767 {
    case 4294967295:
        if (true && false) {
            var v22 uint32 = 2147483647
            if (((146 == 935) && false) || false) {
                if ((true || false) && (!true)) {
                    var v23 uint32 = 694
                    v24 := []int{32767, -9223372036854775808, 966, -895}
                    if ((55 <= 167) || true) {
                        v17 = v17
                        if true {
                            fmt.Println(v4)
                            v1 = v2
                        }
                    }
                    _ = v23
                    _ = v24
                } else {
                    var v25 float64 = 1e10
                    v26 := []int{int(v3), 869, -689}
                    _ = v25
                    _ = v26
                }
            }
            _ = v22
        }
    case 598:
        v3 += 42
    case 32768:
        v27 := []int{-9223372036854775808, 209, 127, int(v3)}
        _ = v27
    }
    if (9223372036854775807 != -953) {
        if true {
            fmt.Println(v17)
            v1 += (v4 % ((2) | 1))
            v28 := GenS1{F0: ""}
            _ = v28
        } else {
            var v29 bool = false
            v1 += 651
            _ = v29
        }
        var v30 bool = false
        _ = v30
    } else {
        var v31 int64 = -1
        v32 := []int{395, 127, 320, 128}
        switch (("\n\t" + "\n\t") + "\n\t") {
        case "":
            var v33 int16 = 32767
            _ = v33
        case "hello":
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
        case "\n\t":
            var v34 uint8 = 236
            _ = v34
        default:
            switch ("\n\t" + "a") {
            case "":
                fmt.Println(v1, v17, v4)
            default:
                v17 += ((v17 + ((1e10 + (v17 / v17)) / ((v17 / v17) - v17))) + (0.1 + float32((0.1 + -2.25))))
            }
        }
        _ = v31
        _ = v32
    }
    var v35 uint16 = (^(v1))
    switch -9223372036854775808 {
    case -2147483648:
        v36 := GenS3{F0: "foo bar", F1: 100}
        _ = v36
    case -640:
        var v37 uint = 1
        _ = v37
    case 9223372036854775807:
        var v38 float32 = v17
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
