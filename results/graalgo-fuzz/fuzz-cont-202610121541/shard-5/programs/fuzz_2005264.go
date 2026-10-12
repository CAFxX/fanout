package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint16 }
func genFn0(p0 int16) uint8 {
    v7 := []int{}
    var v8 int32 = -1
    _ = p0
    _ = v7
    _ = v8
    return 42
}
func genFn1() string {
    var v9 int8 = -1
    v9 += 1
    v10 := GenS1{F0: 2, F1: 18446744073709551615}
    if (true || true) {
        if (true && ("\n\t" > "")) {
            var v11 uint64 = 2
            v12 := GenS2{F0: "hello", F1: "\n\t"}
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 6; fi13++ {
                v14 := GenS2{F0: "a", F1: "hello"}
                v9 += (v9 >> ((18446744073709551615) & 63))
                for fi15 := 0; fi15 < 11; fi15++ {
                    v16 := []int{15, fi13}
                    if (false && (!true)) {
                        fmt.Println(v9, fi13, v9)
                        fmt.Println(fi15, v9)
                    } else {
                        v17 := GenS2{F0: "hello", F1: ("a" + "a")}
                        _ = v17
                    }
                    _ = fi15
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
        }
        var v18 uint8 = 128
        if (false && false) {
            var v19 uint64 = 1
            v20 := GenS1{F0: 42, F1: (v19 >> ((65535) & 63))}
            var v21 uint16 = 255
            _ = v19
            _ = v20
            _ = v21
        }
        _ = v18
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := GenS1{F0: 2, F1: 4294967295}
        _ = fi22
        _ = v23
    }
    _ = v10
    _ = v9
    return ("" + "a")
}

func main() {
    var v1 bool = true
    v2 := []uint32{}
    var v3 float64 = (0.1 - ((1e10 - (float64(3.14159) - (0.1 / 3.14159))) * -2.25))
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    v6 := []int{}
    _ = len(v4)
    switch 32768 {
    case 854:
        fmt.Println(v1)
    default:
        delete(v5, 65535)
    }
    fmt.Println(v3)
    fmt.Println(v1, v3)
    switch 1 {
    case -1:
        for fi24 := 0; fi24 < 2; fi24++ {
            v1 = v1
            var v25 uint16 = 567
            var v26 uint = 0
            _ = fi24
            _ = v25
            _ = v26
        }
    default:
        v3 -= (float64(0.0) + 1.5)
    }
    var v27 uint16 = 32767
    v1 = (false && (4294967295 <= -711))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        v29 := GenS1{F0: 2, F1: 18446744073709551615}
        _ = fi28
        _ = v29
    }
    if (4294967295 < 973) {
        var v30 string = "foo bar"
        v3 = v3
        _ = v30
    }
    var v31 int = 9223372036854775807
    var v32 float64 = v3
    switch v31 {
    case -9223372036854775808:
        switch ("hello" + (((((("hello" + ("foo bar" + ("\n\t" + ("foo bar" + ("hello" + "foo bar"))))) + "") + "") + "") + ((("hello" + "\n\t") + ("foo bar" + "hello")) + (("" + "") + "a"))) + "\n\t")) {
        case "":
            for fi33 := 0; fi33 < 7; fi33++ {
                v34 := GenS3{F0: 1, F1: 255, F2: 0}
                _ = fi33
                _ = v34
            }
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            if (!(!(false || (-310 != -746)))) {
                var v35 bool = v1
                _ = v35
            }
        }
    case -2147483648:
        for fi36 := 0; fi36 < 2; fi36++ {
            switch v31 {
            case -959:
                switch "foo bar" {
                case "foo bar":
                    var v37 int64 = 1
                    _ = v37
                }
            default:
                v31 = fi36
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v38 float64 = v32
            _ = fi36
            _ = v38
        }
    }
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
