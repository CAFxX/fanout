package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 string; F1 int; F2 string }
type GenS3 struct { F0 int8 }
func genFn0(p0 int8, p1 int16, p2 int) uint {
    if (928 <= 1) {
        v5 := []int{}
        _ = v5
    } else {
        p1 = p1
    }
    v6 := GenS1{F0: 0, F1: 0}
    v7 := GenS1{F0: 18446744073709551615, F1: 591}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []int{-9223372036854775808, 84, -2147483648, 32767}
    v2 := make(map[uint32]uint32)
    var v3 int = 311
    var v4 uint8 = 127
    for fi9 := 0; fi9 < 10; fi9++ {
        switch fi9 {
        case 4294967295:
            fmt.Println(fi9, v4)
        case 127:
            v2[65535] = 744
        default:
            var v10 uint64 = 2
            _ = v10
        }
        delete(v2, 2)
        var v11 bool = ((("foo bar" + "a") + "\n\t") != "foo bar")
        _ = fi9
        _ = v11
    }
    var v12 int = 847
    var v13 uint16 = 2
    var v14 uint64 = 42
    switch (v12 >> ((4294967295) & 63)) {
    case -212:
        v3 = (^(v12))
    case -878:
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS1{F0: v14, F1: 2147483647}
            _ = fi15
            _ = v16
        }
    case 127:
        v12 -= v12
    default:
        if (!(false && ((^(v14)) > (^(v14))))) {
            fmt.Println(v13, v14, v13)
            switch v3 {
            case -1:
                for fi17 := 0; fi17 < 17; fi17++ {
                    if (159 < 2) {
                        var v18 int = v3
                        v14 = 2
                        _ = v18
                    } else {
                        _ = len(v1)
                        var v19 int16 = 0
                        v20 := GenS2{F0: "\n\t", F1: (fi17 * (v3 * fi17)), F2: ((("" + "") + "\n\t") + ("hello" + "a"))}
                        _ = v19
                        _ = v20
                    }
                    _ = fi17
                }
            case 32768:
                fmt.Println(v12)
            case -956:
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 int32 = 2147483647
                    var v23 int64 = 1
                    v24 := GenS3{F0: 56}
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
            switch "a" {
            case "":
                switch (v12 << ((630) & 63)) {
                case -1:
                    var v25 uint32 = 4294967295
                    _ = v25
                case 128:
                    for fi26 := 0; fi26 < 10; fi26++ {
                        for fi27 := 0; fi27 < 13; fi27++ {
                            var v28 uint64 = v14
                            fmt.Println(fi27, v14)
                            _ = fi27
                            _ = v28
                        }
                        _ = fi26
                    }
                case 32767:
                    v14 = (v14 & (v14 << ((4294967295) & 63)))
                }
            default:
                var v29 uint8 = v4
                _ = v29
            }
        }
    }
    var v30 int8 = 0
    for fi31 := 0; fi31 < 13; fi31++ {
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := GenS2{F0: "hello", F1: fi31, F2: "hello"}
            if (2 < 2147483648) {
                var v34 int32 = 2147483647
                v13 += (^(v13))
                _ = v34
            }
            var v35 int = v12
            _ = fi32
            _ = v33
            _ = v35
        }
        _ = fi31
    }
    var v36 uint64 = v14
    v37 := GenS3{F0: v30}
    var v38 int16 = 0
    var v39 int64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
