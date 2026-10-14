package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 int }
type GenS4 struct { F0 int }

func main() {
    var v1 uint8 = 68
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 string = "a"
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 9; fi7++ {
        switch ("hello" + (v4 + v4)) {
        case "hello":
            v4 = ""
        case "a":
            v8 := []int{4294967295, fi7, (fi7 >> ((2) & 63))}
            _ = v8
        case "":
            v9 := GenS1{F0: int8(fi7), F1: 4294967295}
            _ = v9
        default:
            fmt.Println(v4, fi7)
        }
        _ = fi7
    }
    v4 = ("a" + v4)
    v10 := []int{127, 2147483647, 32767}
    v11 := GenS1{F0: 63, F1: 4294967295}
    v4 = v4
    var v12 int8 = 42
    fmt.Println(v4, v4, v1)
    if (true && false) {
        if true {
            if (!(false && (1.5 == ((3.14159 / 1e10) / (float32(0.1) * 1.5))))) {
                if len(v10) > 0 {
                    v10[0] = 741
                }
            }
        } else {
            fmt.Println(v12, v12)
            delete(v5, 209)
            var v13 float64 = (3.14159 * -2.25)
            _ = v13
        }
        if false {
            if false {
                v14 := GenS1{F0: v12, F1: 2}
                v15 := GenS1{F0: -1, F1: 221}
                for fi16 := 0; fi16 < 19; fi16++ {
                    var v17 uint = 65535
                    v17 = (^(v17))
                    var v18 float64 = 0.1
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = v14
                _ = v15
            } else {
                v19 := GenS2{F0: (^(v1)), F1: -1}
                v4 = (v4 + v4)
                v4 = (v4 + ((v4 + "hello") + v4))
                _ = v19
            }
        }
        switch 2147483647 {
        case 4294967295:
            fmt.Println(v12)
        case 9223372036854775807:
            v1 += (v1 << ((65535) & 63))
        case 397:
            switch "hello" {
            case "hello":
                v12 += 61
            case "foo bar":
                var v20 uint8 = v1
                _ = v20
            case "\n\t":
                fmt.Println(v1, v4)
            default:
                if true {
                    v4 = ("foo bar" + "a")
                    switch v4 {
                    case "foo bar":
                        var v21 float64 = ((float64(3.14159) - float64(1e10)) * ((1e10 / 0.1) * float64((1.5 - -2.25))))
                        _ = v21
                    default:
                        var v22 bool = true
                        _ = v22
                    }
                    v12 -= v12
                }
            }
        default:
            v1 = 0
        }
    } else {
        v6[4294967295] = 42
        if true {
            for fi23 := 0; fi23 < 17; fi23++ {
                v24 := GenS2{F0: 1, F1: int64(v1)}
                _ = fi23
                _ = v24
            }
            var v25 int8 = v12
            if (!(false || false)) {
                v26 := GenS2{F0: 86, F1: -165}
                _ = v26
            } else {
                var v27 int64 = 42
                var v28 int8 = (v25 << ((42) & 63))
                _ = v27
                _ = v28
            }
            _ = v25
        }
    }
    switch (v4 + "") {
    case "foo bar":
        var v29 int16 = 1000
        _ = v29
    case "hello":
        v12 -= (v12 + v12)
    default:
        fmt.Println(v1, v4)
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        v31 := GenS2{F0: v1, F1: 42}
        for fi32 := 0; fi32 < 16; fi32++ {
            v33 := GenS3{F0: 1, F1: (-(fi32)), F2: (fi32 ^ -2147483648)}
            var v34 int8 = 100
            _ = fi32
            _ = v33
            _ = v34
        }
        var v35 string = v4
        _ = fi30
        _ = v31
        _ = v35
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        v12 -= (-(v12))
        _ = fi36
    }
    v10 = append(v10, 961)
    for fi37 := 0; fi37 < 14; fi37++ {
        v10 = append(v10, (fi37 << ((647) & 63)))
        _ = fi37
    }
    var v38 uint8 = (^(v1))
    v39 := GenS3{F0: 2, F1: 0, F2: -302}
    for fi40 := 0; fi40 < 8; fi40++ {
        if (v38 <= (^(v38))) {
            switch v4 {
            case "":
                v12 += v12
            case "a":
                if (-1 < -32768) {
                    fmt.Println(v12, fi40, v12)
                    fmt.Println(v12, v38, v1)
                }
            case "foo bar":
                for fi41 := 0; fi41 < 19; fi41++ {
                    if false {
                        var v42 string = v4
                        if false {
                            v43 := GenS2{F0: (v38 % ((42) | 1)), F1: -1}
                            v42 = v42
                            _ = v43
                        }
                        v44 := GenS3{F0: 2, F1: fi41, F2: 536}
                        _ = v42
                        _ = v44
                    } else {
                        v4 = ("\n\t" + ("\n\t" + ((("hello" + v4) + (((("\n\t" + v4) + v4) + v4) + v4)) + v4)))
                    }
                    v1 -= 217
                    fmt.Println(v38)
                    _ = fi41
                }
            default:
                var v45 uint16 = 32767
                _ = v45
            }
            for fi46 := 0; fi46 < 7; fi46++ {
                v4 = (v4 + v4)
                _ = fi46
            }
        } else {
            v38 = 255
        }
        _ = fi40
    }
    for fi47 := 0; fi47 < 7; fi47++ {
        v1 = (v38 ^ (v1 ^ v1))
        if ((!true) && (true && false)) {
            v3[fi47] = v4
        } else {
            var v48 int16 = -480
            v38 = (v1 | (v38 % (((v38 % ((v1) | 1))) | 1)))
            _ = v48
        }
        v49 := GenS4{F0: fi47}
        _ = fi47
        _ = v49
    }
    fmt.Println(v4, v38, v38)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
