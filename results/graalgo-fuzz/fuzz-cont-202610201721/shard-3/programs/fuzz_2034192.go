package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 string; F2 string }

func main() {
    var v1 int8 = -35
    var v2 uint32 = 8
    var v3 int8 = v1
    v4 := make(map[uint32]int)
    switch 42 {
    case 4294967295:
        delete(v4, (v2 % ((907) | 1)))
    }
    v1 -= 42
    if true {
        if (0 <= 18446744073709551615) {
            v5 := []int{-9223372036854775808}
            _ = v5
        } else {
            v2 = (v2 / (((v2 * v2)) | 1))
            fmt.Println(v2, v3, v2)
        }
        var v6 uint32 = (v2 ^ (^(v2)))
        _ = v6
    } else {
        v7 := GenS1{F0: -1, F1: 2}
        for fi8 := 0; fi8 < 1; fi8++ {
            v1 = (v1 + (v3 % (((v3 ^ 1)) | 1)))
            fmt.Println(v1)
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v3)
            case "a":
                fmt.Println(v3, v2, v3)
            case "hello":
                var v9 uint32 = 65535
                _ = v9
            default:
                var v10 int8 = 1
                _ = v10
            }
            _ = fi8
        }
        _ = len(v4)
        _ = v7
    }
    v1 = -58
    switch 4294967295 {
    case -170:
        delete(v4, 0)
    default:
        if (true && (true || false)) {
            v11 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615}
            _ = v11
        }
    }
    var v12 float32 = (float32(1.5) / float32(((1.5 - 0.1) / -2.25)))
    switch "\n\t" {
    case "foo bar":
        v13 := []int{int(v2), 32768, 128}
        _ = v13
    default:
        var v14 uint32 = 27
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            if (!false) {
                fmt.Println(fi16, v2)
            } else {
                v3 = v1
            }
            _ = fi16
        }
        _ = fi15
    }
    var v17 float64 = 1e10
    for fi18 := 0; fi18 < 8; fi18++ {
        if ("" <= "\n\t") {
            if true {
                for fi19 := 0; fi19 < 19; fi19++ {
                    fmt.Println(v3, v1)
                    v12 -= v12
                    _ = fi19
                }
                var v20 uint = 42
                var v21 bool = (false && false)
                _ = v20
                _ = v21
            }
        }
        var v22 float64 = (v17 - 0.1)
        for fi23 := 0; fi23 < 5; fi23++ {
            if (uint(v2) < 960) {
                fmt.Println(fi23, v17, v12)
            } else {
                v24 := []int{-9223372036854775808, fi18}
                switch ("foo bar" + ("" + ("a" + "\n\t"))) {
                case "hello":
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                case "":
                    if true {
                        var v25 int32 = -1
                        var v26 float32 = (v12 / -2.25)
                        _ = v25
                        _ = v26
                    }
                case "a":
                    v2 += v2
                }
                _ = v24
            }
            _ = fi23
        }
        _ = fi18
        _ = v22
    }
    if true {
        v27 := GenS1{F0: 1, F1: 63}
        _ = v27
    } else {
        fmt.Println(v17, v12, v1)
    }
    var v28 float32 = v12
    v17 = (-2.25 / (v17 + 0.1))
    switch -464 {
    case -510:
        var v29 int32 = 1
        _ = v29
    case -9223372036854775808:
        if ((v17 < -2.25) || (false && false)) {
            v30 := v4[v2]
            if (true || true) {
                v4[v2] = v30
                v31 := []int{v30}
                fmt.Println(v1)
                _ = v31
            }
            _ = v30
        } else {
            v32 := []int{32767}
            var v33 float64 = v17
            _ = v32
            _ = v33
        }
    case 4294967295:
        switch (("hello" + "a") + "hello") {
        case "foo bar":
            for fi34 := 0; fi34 < 18; fi34++ {
                for fi35 := 0; fi35 < 10; fi35++ {
                    v17 = (v17 + -2.25)
                    _ = fi35
                }
                fmt.Println(v3, v1)
                var v36 string = (("foo bar" + ("foo bar" + ("a" + "a"))) + "\n\t")
                _ = fi34
                _ = v36
            }
        case "":
            v37 := GenS1{F0: -596, F1: 1}
            _ = v37
        case "\n\t":
            var v38 int16 = 42
            _ = v38
        }
    }
    _ = len(v4)
    v39 := GenS2{F0: 42, F1: ((("hello" + ("\n\t" + "")) + ("a" + "")) + (("foo bar" + "foo bar") + (("" + "\n\t") + "a"))), F2: "hello"}
    v28 += v28
    switch "a" {
    case "\n\t":
        v17 = ((float64(v28) / (((3.14159 - ((v17 / v17) - v17)) - (1e10 + v17)) - ((v17 * ((v17 - (float64((v12 / v12)) - ((((1.5 * 1.5) - (float64((v12 * float32(v17))) + (v17 + float64(1.5)))) + (1e10 * 3.14159)) + (1e10 + v17)))) / 3.14159)) + v17))) * 3.14159)
    case "":
        v1 += 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
}
