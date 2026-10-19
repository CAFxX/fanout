package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int64; F1 int; F2 int8 }
type GenS6 struct { F0 uint16; F1 int; F2 uint }
func genFn0(p0 int64, p1 int8) uint64 {
    fmt.Println(p1, p0)
    if false {
        fmt.Println(p0, p1)
        p0 = p0
    }
    _ = p0
    _ = p1
    return 4294967295
}
func genFn1(p0 int16, p1 uint64, p2 uint16) int {
    if (302 == p1) {
        p2 -= p2
    } else {
        for fi4 := 0; fi4 < 14; fi4++ {
            if ((p1 % (((p1 % ((1) | 1))) | 1)) <= genFn0(int64(p2), 30)) {
                if (!(633 != (p1 - p1))) {
                    if false {
                        p0 = 1
                    } else {
                        var v5 int = fi4
                        fmt.Println(p2)
                        _ = v5
                    }
                    switch (((("hello" + "\n\t") + ("a" + (("a" + "hello") + "foo bar"))) + "") + "\n\t") {
                    case "a":
                        p1 += p1
                    case "\n\t":
                        p1 = p1
                    case "hello":
                        var v6 float64 = 0.0
                        _ = v6
                    }
                    var v7 uint16 = (^(p2))
                    _ = v7
                } else {
                    if true {
                        for fi8 := 0; fi8 < 9; fi8++ {
                            p0 = (p0 - p0)
                            p2 = 65535
                            _ = fi8
                        }
                    } else {
                        v9 := []int{fi4}
                        _ = v9
                    }
                }
                switch fi4 {
                case 42:
                    v10 := GenS1{F0: ((false || (729 == 9223372036854775807)) && (p1 < (^(p1))))}
                    _ = v10
                default:
                    p0 = int16(p1)
                }
                fmt.Println(p2)
            } else {
                var v11 uint8 = 1
                _ = v11
            }
            _ = fi4
        }
        fmt.Println(p1, p0, p0)
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS1{F0: ((false || false) || true)}
        _ = fi12
        _ = v13
    }
    v14 := []int{-2147483648, 350, 32768}
    v15 := GenS1{F0: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]bool)
    var v3 uint16 = 200
    if true {
        v16 := GenS2{F0: (v1 / ((0) | 1)), F1: -30}
        switch 4294967295 {
        case 127:
            var v17 uint32 = 1
            _ = v17
        case -948:
            v3 = (v3 ^ (v3 << ((255) & 63)))
        }
        switch "hello" {
        case "foo bar":
            v18 := []int{179, -9223372036854775808, -679}
            _ = v18
        case "a":
            v19 := GenS1{F0: (((((1e10 / (3.14159 * ((-2.25 + (1.5 / ((3.14159 / 1.0) - ((3.14159 / 1e10) + 1.5)))) + (1e10 - -2.25)))) * (float64(-2.25) * (float64(-2.25) * (3.14159 / 1.0)))) / (0.1 + 1.5)) + 3.14159) == 3.14159)}
            _ = v19
        case "\n\t":
            var v20 uint16 = (^(v3))
            _ = v20
        }
        _ = v16
    } else {
        delete(v2, ("" + ("\n\t" + ("\n\t" + (("hello" + "foo bar") + "foo bar")))))
    }
    fmt.Println(v3)
    if true {
        v1 = v1
        switch -408 {
        case -874:
            switch "foo bar" {
            case "":
                if false {
                    var v21 bool = (!(!(genFn0(222, 127) != 91)))
                    v3 = v3
                    _ = v21
                }
            case "foo bar":
                fmt.Println(v3)
            }
        }
        v3 -= 65535
    }
    fmt.Println(v3)
    var v22 uint16 = v3
    v23 := GenS2{F0: (^(v1)), F1: -110}
    var v24 uint8 = (v1 ^ v1)
    v25 := GenS1{F0: (true || false)}
    for fi26 := 0; fi26 < 15; fi26++ {
        v3 += 255
        _ = fi26
    }
    if false {
        for fi27 := 0; fi27 < 8; fi27++ {
            for fi28 := 0; fi28 < 10; fi28++ {
                v1 -= v24
                v22 = v22
                v29 := GenS1{F0: (-9223372036854775808 < int64(fi28))}
                _ = fi28
                _ = v29
            }
            v30 := []int{fi27}
            v31 := GenS3{F0: v22, F1: v3, F2: v22}
            _ = fi27
            _ = v30
            _ = v31
        }
        fmt.Println(v3, v24)
    } else {
        var v32 int16 = -685
        switch 32767 {
        case 127:
            for fi33 := 0; fi33 < 16; fi33++ {
                for fi34 := 0; fi34 < 13; fi34++ {
                    var v35 uint = uint(v24)
                    _ = fi34
                    _ = v35
                }
                v36 := GenS4{F0: 255, F1: (v24 / (((v1 % (((v1 & v24)) | 1))) | 1))}
                _ = fi33
                _ = v36
            }
        case -687:
            fmt.Println(v32)
        case 0:
            v37 := GenS2{F0: (v24 % ((v24) | 1)), F1: 127}
            _ = v37
        }
        v24 += v1
        _ = v32
    }
    v1 = (v1 << ((4294967295) & 63))
    for fi38 := 0; fi38 < 3; fi38++ {
        if ((v24 * v24) <= (v24 + v1)) {
            if false {
                fmt.Println(v24, fi38)
            } else {
                var v39 bool = false
                _ = v39
            }
            v40 := GenS5{F0: -9223372036854775808, F1: (fi38 * -204), F2: 12}
            if false {
                fmt.Println(fi38, v1, v3)
                fmt.Println(v3, v22)
            } else {
                var v41 int32 = 397
                for fi42 := 0; fi42 < 9; fi42++ {
                    v24 = 1
                    v24 += 2
                    v43 := GenS6{F0: 65535, F1: (fi38 << ((18446744073709551615) & 63)), F2: 65535}
                    var v44 uint32 = 65535
                    if (1e10 <= 0.0) {
                        var v45 int = fi38
                        fmt.Println(v41)
                        _ = v45
                    }
                    var v46 int = 645
                    v47 := GenS3{F0: (v22 & (v22 ^ (v22 + (v22 | uint16(v1))))), F1: (^(v3)), F2: uint16(v24)}
                    v3 = 42
                    _ = fi42
                    _ = v43
                    _ = v44
                    _ = v46
                    _ = v47
                }
                v48 := GenS4{F0: 32767, F1: (v24 + (v1 / ((42) | 1)))}
                _ = v41
                _ = v48
            }
            _ = v40
        } else {
            if (true && ((2 >= v1) && (false || false))) {
                fmt.Println(fi38, v3, fi38)
                var v49 int = fi38
                _ = v49
            } else {
                fmt.Println(fi38, v22)
            }
        }
        for fi50 := 0; fi50 < 8; fi50++ {
            v22 -= 523
            var v51 uint32 = 0
            var v52 uint8 = v1
            _ = fi50
            _ = v51
            _ = v52
        }
        fmt.Println(v1)
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
