package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint64; F2 bool }
type GenS5 struct { F0 string; F1 uint8; F2 uint64 }
func genFn0(p0 uint8) bool {
    v6 := GenS1{F0: false}
    switch "\n\t" {
    case "hello":
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 11; fi8++ {
                v9 := []int{-9223372036854775808, -1, (fi8 << ((255) & 63)), (fi8 ^ -33)}
                var v10 uint = 0
                v10 = (v10 * 501)
                if len(v9) > 0 {
                    v9[0] = (fi7 / (((fi8 % ((-9223372036854775808) | 1))) | 1))
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            if true {
                switch (fi7 << ((4294967295) & 63)) {
                case 984:
                    for fi11 := 0; fi11 < 5; fi11++ {
                        if ((!((!true) || true)) && (42 <= 65535)) {
                            var v12 uint8 = 255
                            var v13 int32 = -942
                            _ = v12
                            _ = v13
                        }
                        var v14 int8 = 40
                        _ = fi11
                        _ = v14
                    }
                }
            }
            _ = fi7
        }
    }
    fmt.Println(p0)
    switch 2147483647 {
    case -2147483648:
        p0 -= p0
    default:
        v15 := GenS2{F0: 42, F1: 91, F2: 32768}
        _ = v15
    }
    _ = p0
    _ = v6
    return (!(1 <= 118))
}
func genFn1(p0 int8, p1 uint64) int16 {
    var v16 int16 = 1000
    p0 = (-(p0))
    p1 = (p1 + (p1 + (p1 >> ((162) & 63))))
    _ = p0
    _ = p1
    _ = v16
    return 591
}

func main() {
    v1 := []string{"a"}
    var v2 int = 32767
    v3 := []string{"hello", (("foo bar" + ("" + "\n\t")) + "a"), "hello"}
    var v4 uint8 = 96
    var v5 string = ("\n\t" + ("\n\t" + "\n\t"))
    v5 = ("\n\t" + v5)
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int = (v2 + (v2 + (v2 % ((4294967295) | 1))))
        _ = fi17
        _ = v18
    }
    if (false || (v2 != v2)) {
        v19 := GenS3{F0: 255, F1: genFn1(0, 18446744073709551615)}
        v20 := GenS4{F0: (v4 & v4), F1: uint64(v4), F2: false}
        _ = v19
        _ = v20
    } else {
        switch 4294967295 {
        case 128:
            var v21 uint8 = (^(v4))
            _ = v21
        case -102:
            v22 := GenS4{F0: 2, F1: 979, F2: true}
            _ = v22
        default:
            v23 := GenS1{F0: (842 >= 805)}
            _ = v23
        }
        v5 = v5
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            v4 += v4
            var v26 float32 = 1e10
            _ = fi25
            _ = v26
        }
        fmt.Println(v5, v4, v2)
        _ = fi24
    }
    var v27 string = "hello"
    if true {
        if true {
            var v28 int8 = 0
            var v29 int16 = -32768
            if true {
                var v30 int32 = -2147483648
                _ = v30
            }
            _ = v28
            _ = v29
        } else {
            fmt.Println(v27)
        }
        var v31 uint16 = 0
        var v32 uint16 = (v31 / ((32767) | 1))
        _ = v31
        _ = v32
    }
    v33 := GenS5{F0: v5, F1: 1, F2: 4294967295}
    fmt.Println(v4, v4)
    for fi34 := 0; fi34 < 7; fi34++ {
        var v35 int8 = 42
        switch 2147483647 {
        case -226:
            fmt.Println(v5, v2, v4)
        case -1:
            fmt.Println(v35, v35)
        default:
            v35 += 37
        }
        _ = fi34
        _ = v35
    }
    v5 = v5
    var v36 float32 = (1.5 - 0.1)
    switch (v2 << ((581) & 63)) {
    case 9223372036854775807:
        switch (v2 << ((2) & 63)) {
        case 579:
            if len(v3) > 0 {
                v3[0] = v27
            }
        case 0:
            switch -352 {
            case 32767:
                var v37 uint64 = 954
                _ = v37
            case 42:
                for fi38 := 0; fi38 < 5; fi38++ {
                    for fi39 := 0; fi39 < 1; fi39++ {
                        var v40 int32 = 0
                        _ = fi39
                        _ = v40
                    }
                    for fi41 := 0; fi41 < 12; fi41++ {
                        var v42 int8 = 35
                        v4 -= (^(v4))
                        fmt.Println(fi38, v36)
                        _ = fi41
                        _ = v42
                    }
                    _ = fi38
                }
            }
        case 4294967295:
            for fi43 := 0; fi43 < 4; fi43++ {
                fmt.Println(v2, v2, v36)
                _ = fi43
            }
        }
    }
    v44 := GenS3{F0: 35, F1: 0}
    fmt.Println(v5, v36)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v44
    _ = v5
}
