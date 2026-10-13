package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 bool; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint; F2 int64 }
type GenS4 struct { F0 string; F1 int8; F2 string }

func main() {
    var v1 bool = (false && (true || (42 >= 0)))
    var v2 string = "hello"
    v3 := make(map[uint32]bool)
    var v4 uint64 = 0
    v5 := make(map[int]uint32)
    var v6 int = -90
    v1 = ((0 != 945) || (v1 && ("foo bar" == v2)))
    v7 := GenS1{F0: -940}
    var v8 int16 = 582
    for fi9 := 0; fi9 < 1; fi9++ {
        if false {
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 int = -9223372036854775808
                _ = fi10
                _ = v11
            }
            fmt.Println(v1)
            if (1 > 0) {
                if v1 {
                    switch (^(fi9)) {
                    case -177:
                        if ((3.14159 + 1.5) != float64(-2.25)) {
                            if (-564 == 43) {
                                v5[-95] = 0
                            }
                            switch v2 {
                            case "a":
                                v12 := v3[2]
                                _ = v12
                            case "hello":
                                fmt.Println(v8, fi9)
                            default:
                                var v13 uint16 = uint16(v4)
                                _ = v13
                            }
                        }
                    case 9223372036854775807:
                        for fi14 := 0; fi14 < 6; fi14++ {
                            var v15 uint16 = 2
                            var v16 int64 = 9223372036854775807
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                    v17 := GenS2{F0: -2147483648, F1: ((!v1) || (42 <= -920)), F2: (v4 ^ v4)}
                    var v18 int32 = 1
                    _ = v17
                    _ = v18
                }
                v19 := GenS2{F0: (-(fi9)), F1: (32767 <= (v8 >> ((1) & 63))), F2: v4}
                fmt.Println(v8, fi9)
                _ = v19
            } else {
                v6 += fi9
            }
        }
        _ = fi9
    }
    var v20 int8 = 100
    v21 := GenS1{F0: (v8 - 0)}
    v4 = (v4 << ((18446744073709551615) & 63))
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 uint64 = (v4 * 2)
        var v24 float64 = ((1e10 + 1.5) - 1.5)
        v6 -= 1
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        switch v2 {
        case "a":
            delete(v5, -2147483648)
        case "hello":
            v26 := GenS1{F0: v8}
            _ = v26
        case "":
            if (("a" < ((v2 + v2) + (v2 + (((v2 + (((v2 + ((v2 + v2) + ((("" + v2) + ("\n\t" + ((("hello" + (("a" + ((v2 + v2) + "")) + v2)) + (v2 + ((v2 + v2) + v2))) + v2))) + (v2 + "")))) + v2) + ("hello" + "foo bar"))) + (v2 + v2)) + (v2 + ("a" + ("" + ""))))))) && (2 < 255)) {
                v4 += (v4 * v4)
                for fi27 := 0; fi27 < 1; fi27++ {
                    v28 := []int{int(v8)}
                    _ = fi27
                    _ = v28
                }
                v29 := GenS3{F0: 2147483647, F1: 65535, F2: 42}
                _ = v29
            }
        default:
            var v30 float64 = -2.25
            _ = v30
        }
        if false {
            var v31 uint8 = 2
            v2 = (v2 + v2)
            _ = v31
        }
        _ = fi25
    }
    switch (v6 | (v6 - 32768)) {
    case -9223372036854775808:
        v32 := GenS4{F0: v2, F1: 20, F2: "\n\t"}
        _ = v32
    case 4294967295:
        for fi33 := 0; fi33 < 10; fi33++ {
            var v34 uint = 261
            v6 += (^(fi33))
            v34 -= 2
            _ = fi33
            _ = v34
        }
    default:
        switch "a" {
        case "foo bar":
            var v35 int = -374
            _ = v35
        case "a":
            var v36 int16 = (v8 + 1000)
            _ = v36
        case "\n\t":
            delete(v3, 2147483647)
        }
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        v4 = v4
        _ = fi37
    }
    v38 := v3[0]
    for fi39 := 0; fi39 < 19; fi39++ {
        fmt.Println(v8, v8)
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
