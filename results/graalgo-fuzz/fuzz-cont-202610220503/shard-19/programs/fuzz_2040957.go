package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int, p1 int, p2 uint16) int64 {
    fmt.Println(p0, p1, p0)
    switch "" {
    case "\n\t":
        fmt.Println(p2, p0)
    case "":
        p2 = 28
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 bool = true
    var v2 int8 = 1
    v3 := make(map[uint32]int)
    var v4 uint = 65535
    if v1 {
        v5 := []int{16, 588, 9223372036854775807}
        v2 -= v2
        v4 = v4
        _ = v5
    } else {
        if (2 == 162) {
            if v1 {
                switch -1 {
                case -801:
                    v1 = v1
                case -58:
                    if v1 {
                        v4 -= (^(v4))
                    }
                }
                var v6 uint8 = 128
                _ = v6
            } else {
                v7 := GenS1{F0: "\n\t"}
                if (!(0 > 857)) {
                    v1 = v1
                }
                for fi8 := 0; fi8 < 9; fi8++ {
                    var v9 bool = (!v1)
                    v10 := []int{}
                    var v11 uint = v4
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                _ = v7
            }
        } else {
            v3[2147483648] = 2147483647
            var v12 uint32 = 42
            _ = v12
        }
        switch "\n\t" {
        case "":
            var v13 float32 = (0.1 - ((0.1 * (0.0 + 1e10)) / -2.25))
            _ = v13
        case "foo bar":
            for fi14 := 0; fi14 < 15; fi14++ {
                for fi15 := 0; fi15 < 8; fi15++ {
                    v4 -= v4
                    v16 := GenS1{F0: ("foo bar" + "\n\t")}
                    var v17 uint32 = 425
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v4 = 4294967295
                v18 := GenS2{F0: 29}
                _ = fi14
                _ = v18
            }
        case "a":
            v1 = (42 <= 0)
        default:
            v19 := []int{42}
            _ = v19
        }
    }
    v3[1] = 42
    v1 = false
    v20 := GenS1{F0: "hello"}
    for fi21 := 0; fi21 < 12; fi21++ {
        fmt.Println(fi21, v4)
        var v22 bool = (v1 && v1)
        var v23 int64 = genFn0((fi21 / (((^(fi21))) | 1)), (^(fi21)), 2)
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        switch fi24 {
        case 9223372036854775807:
            if false {
                var v25 uint64 = 500
                fmt.Println(v1, v25)
                v25 -= (v25 ^ v25)
                _ = v25
            }
        case -1:
            fmt.Println(v1, fi24, fi24)
        }
        if v1 {
            var v26 uint16 = 32767
            _ = v26
        } else {
            switch fi24 {
            case 4294967295:
                for fi27 := 0; fi27 < 2; fi27++ {
                    switch 32767 {
                    case -1:
                        delete(v3, 2147483647)
                    }
                    var v28 int = -9223372036854775808
                    _ = fi27
                    _ = v28
                }
            default:
                if (v1 && (2147483647 < -565)) {
                    switch (("a" + "hello") + (("hello" + "a") + "hello")) {
                    case "hello":
                        var v29 int32 = 846
                        _ = v29
                    case "\n\t":
                        v30 := GenS2{F0: 986}
                        _ = v30
                    case "foo bar":
                        v1 = v1
                    }
                } else {
                    v31 := v3[0]
                    _ = v31
                }
            }
            var v32 bool = (158 != 2)
            _ = v32
        }
        if (2147483647 == 2147483647) {
            fmt.Println(v2, fi24, fi24)
            var v33 int8 = v2
            for fi34 := 0; fi34 < 14; fi34++ {
                fmt.Println(v1)
                var v35 uint64 = 570
                _ = fi34
                _ = v35
            }
            _ = v33
        } else {
            v4 = (v4 / (((v4 & 495)) | 1))
        }
        _ = fi24
    }
    var v36 int16 = 42
    v37 := GenS3{F0: (v4 + 18446744073709551615)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
