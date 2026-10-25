package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int64; F1 uint; F2 int }
type GenS4 struct { F0 bool }

func main() {
    v1 := []string{("" + "\n\t"), "a"}
    v2 := []uint32{65535}
    var v3 int32 = 891
    var v4 uint = 65535
    if (true || false) {
        v4 = (v4 % ((v4) | 1))
    } else {
        var v5 int64 = 1
        for fi6 := 0; fi6 < 18; fi6++ {
            if (!false) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v3 = (-(v3))
                var v7 bool = true
                _ = v7
            }
            for fi8 := 0; fi8 < 18; fi8++ {
                if (false && true) {
                    fmt.Println(v3, fi6)
                    fmt.Println(fi6)
                    var v9 float64 = 1e10
                    _ = v9
                }
                v5 = v5
                var v10 int16 = 42
                _ = fi8
                _ = v10
            }
            var v11 float32 = -2.25
            _ = fi6
            _ = v11
        }
        switch "a" {
        case "hello":
            for fi12 := 0; fi12 < 3; fi12++ {
                fmt.Println(v3, fi12)
                v13 := GenS1{F0: fi12}
                _ = fi12
                _ = v13
            }
        case "a":
            if (!(!false)) {
                for fi14 := 0; fi14 < 15; fi14++ {
                    switch ("foo bar" + "") {
                    case "":
                        var v15 int64 = v5
                        _ = v15
                    }
                    if (true && (1000 <= 42)) {
                        v1 = append(v1, "")
                        fmt.Println(fi14)
                    }
                    fmt.Println(fi14, v4, v3)
                    _ = fi14
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v16 int32 = v3
                _ = v16
            }
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = v5
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        v4 -= (v4 & v4)
        _ = fi17
    }
    v3 = -536
    var v18 uint = 1
    var v19 int64 = 0
    if false {
        var v20 uint32 = 2
        fmt.Println(v20, v3, v4)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 uint32 = 441
            _ = fi21
            _ = v22
        }
        switch "hello" {
        case "foo bar":
            fmt.Println(v18, v18, v18)
        case "hello":
            switch -9223372036854775808 {
            case 2147483647:
                v23 := GenS2{F0: 0, F1: v19}
                _ = v23
            case -1:
                var v24 int64 = (v19 | (v19 % (((v19 * v19)) | 1)))
                _ = v24
            case 1:
                var v25 bool = (!(!(45 < 32768)))
                _ = v25
            }
        case "a":
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 uint16 = 1
                if (false || false) {
                    if (!(!false)) {
                        switch fi26 {
                        case 32768:
                            var v28 int64 = 42
                            _ = v28
                        case 0:
                            v29 := GenS2{F0: 166, F1: 901}
                            _ = v29
                        case 870:
                            var v30 int = (fi26 + fi26)
                            _ = v30
                        }
                        fmt.Println(fi26)
                        v31 := GenS2{F0: 1, F1: v19}
                        _ = v31
                    } else {
                        v32 := GenS1{F0: (fi26 / ((-651) | 1))}
                        _ = len(v1)
                        _ = v32
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v19 = -122
                _ = fi26
                _ = v27
            }
        default:
            v33 := GenS3{F0: v19, F1: (v4 % ((18446744073709551615) | 1)), F2: 0}
            _ = v33
        }
        var v34 int8 = 10
        _ = v34
    }
    if (true && false) {
        v4 += v18
    }
    v18 -= v18
    switch "hello" {
    case "\n\t":
        v35 := GenS3{F0: 42, F1: 18446744073709551615, F2: 563}
        _ = v35
    case "a":
        if len(v1) > 0 {
            v1[0] = (("" + "hello") + "")
        }
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        v37 := GenS4{F0: (-121 == 32767)}
        fmt.Println(v19, fi36, v19)
        _ = fi36
        _ = v37
    }
    v1 = append(v1, ("hello" + ("" + "a")))
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 int8 = 127
        for fi40 := 0; fi40 < 7; fi40++ {
            var v41 int16 = 1
            var v42 uint8 = 255
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = fi38
        _ = v39
    }
    switch "hello" {
    case "\n\t":
        var v43 int = 2147483647
        _ = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
