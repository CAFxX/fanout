package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
type GenS3 struct { F0 int8; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint; F1 int16 }

func main() {
    var v1 string = "foo bar"
    var v2 uint = 4294967295
    var v3 uint16 = 42
    v4 := make(map[int]uint32)
    var v5 float32 = (3.14159 - (float32(1e10) * float32(1.5)))
    var v6 uint8 = uint8(v2)
    v3 -= (v3 | (v3 / ((v3) | 1)))
    fmt.Println(v3)
    switch (("" + v1) + ("\n\t" + v1)) {
    case "hello":
        if (((v6 * (v6 >> ((v2) & 63))) > 171) && (!(!true))) {
            var v7 float32 = 0.1
            switch (("hello" + ("" + (v1 + (v1 + (v1 + "a"))))) + (("a" + "a") + (v1 + ("hello" + "foo bar")))) {
            case "":
                for fi8 := 0; fi8 < 8; fi8++ {
                    v5 = v7
                    fmt.Println(v3, v3)
                    v6 -= (v6 << (((^(v2))) & 63))
                    _ = fi8
                }
            case "hello":
                v5 += v7
            default:
                if (v6 <= (v6 | 1)) {
                    for fi9 := 0; fi9 < 6; fi9++ {
                        v2 -= 2
                        fmt.Println(fi9)
                        v10 := []int{(-(fi9))}
                        _ = fi9
                        _ = v10
                    }
                    var v11 string = ((v1 + ("foo bar" + "a")) + v1)
                    _ = v11
                }
            }
            switch "\n\t" {
            case "hello":
                switch 153 {
                case 32768:
                    v12 := GenS1{F0: 2}
                    _ = v12
                case -2147483648:
                    v2 = v2
                }
            }
            _ = v7
        } else {
            v3 = (v3 ^ (v3 >> ((v2) & 63)))
            v13 := GenS1{F0: v6}
            var v14 uint32 = 2147483648
            _ = v13
            _ = v14
        }
    }
    v15 := GenS2{F0: 122, F1: 1, F2: 32767}
    var v16 uint8 = (^(v6))
    v17 := GenS2{F0: 18446744073709551615, F1: 1, F2: -251}
    switch v1 {
    case "":
        for fi18 := 0; fi18 < 14; fi18++ {
            switch "foo bar" {
            case "":
                var v19 float32 = v5
                _ = v19
            case "a":
                v20 := GenS2{F0: 42, F1: 42, F2: -1}
                _ = v20
            }
            var v21 uint8 = v6
            for fi22 := 0; fi22 < 11; fi22++ {
                var v23 float64 = (1.5 - 1.5)
                _ = fi22
                _ = v23
            }
            _ = fi18
            _ = v21
        }
    case "hello":
        v16 -= (v6 + 128)
    case "foo bar":
        var v24 int8 = 100
        _ = v24
    }
    var v25 int8 = 3
    v4[-2147483648] = 4294967295
    for fi26 := 0; fi26 < 11; fi26++ {
        if (42 == -315) {
            for fi27 := 0; fi27 < 2; fi27++ {
                v5 = float32(0.1)
                _ = fi27
            }
            for fi28 := 0; fi28 < 8; fi28++ {
                fmt.Println(v6)
                fmt.Println(fi28, v2)
                if (0 >= 0) {
                    var v29 int32 = -1
                    var v30 bool = (true && (42 <= 1))
                    _ = v29
                    _ = v30
                }
                _ = fi28
            }
        } else {
            v2 = v2
        }
        v31 := GenS3{F0: v25, F1: v2, F2: 0}
        v32 := []int{}
        _ = fi26
        _ = v31
        _ = v32
    }
    v33 := GenS4{F0: (v2 << ((v2) & 63)), F1: 32767}
    if false {
        v5 = 3.14159
        v2 += v2
        if (-9223372036854775808 > -9223372036854775808) {
            var v34 int = -461
            _ = v34
        }
    } else {
        switch v1 {
        case "":
            fmt.Println(v6)
        case "hello":
            v2 -= (v2 << ((65535) & 63))
        default:
            fmt.Println(v1, v1, v6)
        }
        switch 4294967295 {
        case 128:
            var v35 uint32 = 1
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        var v37 uint8 = v6
        if (42 != (v25 & v25)) {
            fmt.Println(v16, v5)
        }
        _ = fi36
        _ = v37
    }
    if ("\n\t" > "\n\t") {
        var v38 string = "a"
        _ = v38
    } else {
        for fi39 := 0; fi39 < 5; fi39++ {
            _ = len(v4)
            var v40 int16 = 576
            _ = fi39
            _ = v40
        }
    }
    v3 -= (v3 + 2)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
