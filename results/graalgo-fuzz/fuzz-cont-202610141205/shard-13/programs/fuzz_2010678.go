package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int64 }
func genFn0(p0 int, p1 uint16) uint64 {
    var v6 string = ("a" + ("foo bar" + "\n\t"))
    switch v6 {
    case "foo bar":
        var v7 string = ("" + "\n\t")
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(p0, p1)
            _ = fi8
        }
    case "a":
        switch "\n\t" {
        case "":
            if (2 < 738) {
                v6 = "hello"
                var v9 string = v6
                var v10 int32 = -566
                _ = v10
                _ = v9
            }
        case "a":
            v11 := GenS1{F0: ("" + (((("\n\t" + ((v6 + v6) + "a")) + "hello") + v6) + "a")), F1: 954, F2: p0}
            _ = v11
        case "\n\t":
            if (p0 < p0) {
                v12 := GenS1{F0: ((v6 + (v6 + ((("a" + "a") + v6) + "hello"))) + "hello"), F1: (p1 >> ((42) & 63)), F2: (p0 << ((42) & 63))}
                v13 := GenS1{F0: (v6 + v6), F1: (p1 << ((818) & 63)), F2: p0}
                _ = v12
                _ = v13
            } else {
                var v14 int64 = -719
                p1 = (^(p1))
                _ = v14
            }
        default:
            p0 = (p0 >> ((1) & 63))
        }
    default:
        for fi15 := 0; fi15 < 14; fi15++ {
            if (9223372036854775807 <= -33) {
                if true {
                    if false {
                        v16 := []int{}
                        var v17 bool = (!(!(((false && (v6 != (v6 + "foo bar"))) || ((-2.25 - 0.1) >= 1e10)) || true)))
                        _ = v16
                        _ = v17
                    }
                    p1 = 32767
                    var v18 string = ""
                    _ = v18
                } else {
                    var v19 int64 = -1
                    _ = v19
                }
            }
            _ = fi15
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 206
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -564
    var v3 int64 = 866
    v4 := []string{"foo bar", "", "foo bar"}
    var v5 int8 = 1
    for fi20 := 0; fi20 < 20; fi20++ {
        v1[280] = (3.14159 != 0.0)
        _ = fi20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        v22 := GenS1{F0: "a", F1: 0, F2: fi21}
        v3 += 42
        for fi23 := 0; fi23 < 12; fi23++ {
            if false {
                v5 -= (v5 * v5)
            }
            v4 = append(v4, "a")
            v4 = append(v4, ("foo bar" + ""))
            switch "foo bar" {
            case "":
                switch (^(fi21)) {
                case 42:
                    delete(v1, 2)
                case 1:
                    v24 := v1[855]
                    _ = v24
                case 32767:
                    var v25 string = "foo bar"
                    _ = v25
                }
            default:
                v5 = (v5 / (((v5 / (((v5 ^ (v5 & (v5 % ((v5) | 1))))) | 1))) | 1))
            }
            if ((false || false) || (65535 != 852)) {
                switch "\n\t" {
                case "a":
                    v5 += 1
                }
                v2 -= v2
                _ = len(v1)
                var v26 int = fi21
                _ = v26
            }
            v27 := GenS1{F0: ("hello" + "\n\t"), F1: 65535, F2: 32768}
            v2 = int32(v5)
            _ = fi23
            _ = v27
        }
        _ = fi21
        _ = v22
    }
    var v28 uint = 4294967295
    var v29 int32 = v2
    fmt.Println(v28)
    v28 = v28
    fmt.Println(v3, v2, v2)
    v30 := GenS2{F0: (v29 + (v29 / (((v29 - v29)) | 1))), F1: false}
    v2 = v29
    var v31 int = 692
    v32 := GenS3{F0: -497}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
