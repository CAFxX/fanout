package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int16, p1 int64, p2 uint8) int64 {
    fmt.Println(p0)
    p2 -= p2
    p0 = 178
    v7 := GenS1{F0: 65535, F1: "\n\t", F2: ("a" + "a")}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}

func main() {
    var v1 uint32 = 2
    var v2 uint32 = v1
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 0
    var v5 int64 = -1
    var v6 uint64 = 4294967295
    switch 742 {
    case 2147483647:
        v5 = v5
    case -9223372036854775808:
        var v8 int64 = 42
        _ = v8
    case 0:
        switch 9223372036854775807 {
        case 32768:
            switch -1 {
            case 416:
                _ = len(v3)
            case 128:
                switch "hello" {
                case "a":
                    var v9 uint8 = 255
                    _ = v9
                case "foo bar":
                    v6 += (v6 % ((v6) | 1))
                default:
                    var v10 uint = 874
                    _ = v10
                }
            case 955:
                for fi11 := 0; fi11 < 10; fi11++ {
                    fmt.Println(v1, v2)
                    v12 := []int{}
                    _ = fi11
                    _ = v12
                }
            default:
                v5 = 0
            }
        case -233:
            if false {
                v13 := GenS1{F0: v2, F1: (("a" + "") + ""), F2: "hello"}
                switch int(v5) {
                case -9223372036854775808:
                    for fi14 := 0; fi14 < 19; fi14++ {
                        if true {
                            fmt.Println(fi14)
                            v5 += v5
                        }
                        if (509 > (v1 & (v1 << ((42) & 63)))) {
                            var v15 int32 = -2147483648
                            var v16 uint = 65535
                            _ = v15
                            _ = v16
                        }
                        var v17 bool = false
                        _ = fi14
                        _ = v17
                    }
                case 32768:
                    var v18 uint16 = (v4 & v4)
                    _ = v18
                case -2147483648:
                    fmt.Println(v6)
                }
                _ = v13
            }
        case 373:
            for fi19 := 0; fi19 < 7; fi19++ {
                v20 := v3[1]
                _ = fi19
                _ = v20
            }
        }
    default:
        v4 += (v4 & (v4 / ((255) | 1)))
    }
    v21 := GenS2{F0: 2, F1: 0, F2: 18446744073709551615}
    var v22 bool = (true && (true && true))
    switch "a" {
    case "a":
        switch (("\n\t" + ("a" + (("\n\t" + "") + ""))) + "a") {
        case "a":
            var v23 string = (("a" + "foo bar") + ("hello" + ("a" + "")))
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            v4 = 0
            _ = fi25
        }
        v4 = (v4 + (v4 + v4))
        var v26 int64 = genFn0(42, (v5 % ((genFn0(-962, genFn0(1000, v5, 202), 185)) | 1)), 93)
        _ = fi24
        _ = v26
    }
    var v27 int16 = -703
    v22 = v22
    v4 -= v4
    if (v22 || v22) {
        if v22 {
            for fi28 := 0; fi28 < 7; fi28++ {
                delete(v3, (v1 | 42))
                if (int32(v6) >= -2147483648) {
                    v29 := v3[(v1 / ((2) | 1))]
                    _ = v29
                }
                _ = fi28
            }
            var v30 uint32 = 481
            _ = v30
        }
        var v31 uint16 = v4
        _ = v31
    } else {
        v32 := GenS3{F0: v4}
        var v33 int = -790
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
