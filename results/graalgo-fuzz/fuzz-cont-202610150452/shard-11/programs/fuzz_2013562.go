package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int, p1 string) uint16 {
    v7 := []int{-465, (-(p0)), (-(p0)), (p0 - (p0 + (p0 / (((^(p0))) | 1))))}
    if true {
        p1 = p1
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v7
    return 65535
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]bool)
    var v3 int8 = 127
    var v4 int64 = 42
    v5 := []string{(("foo bar" + "\n\t") + ("a" + ("a" + "hello"))), "", "\n\t"}
    v6 := []string{""}
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: (fi8 <= -283)}
        var v10 int64 = 9223372036854775807
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: true}
    v12 := GenS2{F0: 0, F1: 586}
    v2[649] = true
    var v13 string = "\n\t"
    if (true || false) {
        switch (v13 + v13) {
        case "hello":
            if true {
                var v14 uint = 826
                v15 := v2[65535]
                var v16 int = 32767
                _ = v14
                _ = v15
                _ = v16
            } else {
                v13 = ((((("hello" + v13) + (v13 + v13)) + v13) + v13) + v13)
                var v17 int32 = 1
                _ = v17
            }
        }
        v3 -= v3
    } else {
        for fi18 := 0; fi18 < 16; fi18++ {
            v1 = (v1 & 1)
            switch (fi18 & (fi18 / ((128) | 1))) {
            case -1:
                var v19 uint16 = 1
                _ = v19
            case 1:
                v20 := GenS1{F0: (!(true && ((-1 != int32(v3)) && (false && false))))}
                _ = v20
            case 4294967295:
                v21 := GenS3{F0: 255, F1: v13}
                _ = v21
            }
            fmt.Println(v3)
            _ = fi18
        }
    }
    var v22 float64 = (-2.25 - 3.14159)
    v13 = v13
    if (42 == 65535) {
        v23 := GenS3{F0: 42, F1: (v13 + v13)}
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := GenS3{F0: 622, F1: (((("\n\t" + (v13 + "\n\t")) + (v13 + v13)) + "") + (v13 + v13))}
            v22 = v22
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 bool = true
                _ = fi26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        v4 += (v4 + (v4 ^ v4))
        _ = v23
    }
    v13 = v13
    v13 = v13
    var v28 uint8 = 127
    if true {
        for fi29 := 0; fi29 < 13; fi29++ {
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 int8 = (v3 | (v3 / (((v3 + v3)) | 1)))
                _ = fi30
                _ = v31
            }
            if true {
                var v32 int = fi29
                var v33 string = "\n\t"
                v34 := GenS4{F0: (v4 * v4)}
                _ = v32
                _ = v33
                _ = v34
            } else {
                switch v13 {
                case "":
                    for fi35 := 0; fi35 < 16; fi35++ {
                        v13 = ("hello" + (v13 + (v13 + (("\n\t" + v13) + ("" + ("hello" + v13))))))
                        _ = fi35
                    }
                case "a":
                    fmt.Println(fi29, v3)
                }
                for fi36 := 0; fi36 < 2; fi36++ {
                    if true {
                        if (float32(1e10) == 0.0) {
                            var v37 uint64 = 2
                            var v38 bool = (v1 <= v1)
                            _ = v37
                            _ = v38
                        }
                    }
                    fmt.Println(v4, v1)
                    fmt.Println(v1, v28, fi36)
                    v39 := GenS4{F0: v4}
                    v22 += float64(1.5)
                    v3 = v3
                    _ = fi36
                    _ = v39
                }
                var v40 int8 = 42
                _ = v40
            }
            var v41 uint16 = genFn0(fi29, ((((v13 + v13) + "\n\t") + ("foo bar" + v13)) + v13))
            _ = fi29
            _ = v41
        }
        for fi42 := 0; fi42 < 13; fi42++ {
            if (true || (!(19 > 0))) {
                if false {
                    v1 = (v1 >> ((749) & 63))
                    v43 := GenS5{F0: 4294967295}
                    _ = v43
                } else {
                    v22 = float64(1.5)
                }
            }
            v22 = v22
            if ((fi42 ^ (fi42 << ((704) & 63))) >= (-(fi42))) {
                v4 += v4
            }
            _ = fi42
        }
    } else {
        if (-2147483648 >= 182) {
            var v44 int = 32768
            delete(v2, 2147483647)
            switch (v13 + v13) {
            case "":
                v3 = (-(v3))
            case "hello":
                if len(v6) > 0 {
                    v6[0] = v13
                }
            default:
                switch "\n\t" {
                case "a":
                    switch v44 {
                    case -2147483648:
                        switch 1 {
                        case 32768:
                            var v45 float32 = (float32(v22) / (1.5 / 0.1))
                            _ = v45
                        case 127:
                            v46 := v2[631]
                            _ = v46
                        case -9223372036854775808:
                            v4 = 1
                        }
                    }
                case "\n\t":
                    fmt.Println(v44, v22, v22)
                case "":
                    for fi47 := 0; fi47 < 7; fi47++ {
                        fmt.Println(fi47, v28)
                        var v48 uint16 = 255
                        _ = fi47
                        _ = v48
                    }
                }
            }
            _ = v44
        }
        for fi49 := 0; fi49 < 7; fi49++ {
            _ = len(v5)
            _ = fi49
        }
    }
    fmt.Println(v1, v13, v3)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
