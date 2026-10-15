package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0() int {
    var v6 int32 = 0
    var v7 int64 = -1
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = (^(v1))
    v3 := make(map[int]uint32)
    var v4 int32 = 2147483647
    var v5 string = "\n\t"
    for fi8 := 0; fi8 < 10; fi8++ {
        if true {
            if (((-2.25 + float64(1e10)) + 0.1) > (0.0 * ((-2.25 / float64((0.1 / 1.0))) / ((1.5 + (((float64(1.5) - 3.14159) + 1.5) / 0.1)) + -2.25)))) {
                var v9 int16 = -312
                _ = v9
            }
        } else {
            for fi10 := 0; fi10 < 7; fi10++ {
                v2 = 18446744073709551615
                _ = fi10
            }
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 int32 = 2147483647
            _ = fi11
            _ = v12
        }
        v4 = (-(v4))
        _ = fi8
    }
    var v13 int64 = int64(v4)
    for fi14 := 0; fi14 < 7; fi14++ {
        switch v5 {
        case "foo bar":
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 uint = 1
                _ = fi15
                _ = v16
            }
        case "":
            v2 = (v2 & 881)
        case "hello":
            v17 := GenS1{F0: true, F1: (v5 + v5), F2: uint32(v1)}
            _ = v17
        default:
            fmt.Println(v5, fi14)
        }
        _ = fi14
    }
    v3[2147483647] = 670
    if false {
        for fi18 := 0; fi18 < 7; fi18++ {
            v19 := GenS2{F0: 786}
            _ = fi18
            _ = v19
        }
        switch 127 {
        case -9223372036854775808:
            if true {
                var v20 uint32 = 4294967295
                var v21 int64 = (^(v13))
                _ = v20
                _ = v21
            } else {
                v22 := GenS2{F0: -32768}
                v1 = (v1 - (v1 << ((65535) & 63)))
                v3[9223372036854775807] = 0
                _ = v22
            }
        case 4294967295:
            v23 := GenS1{F0: (3.14159 <= float64((((-2.25 / 1.5) + ((((float32(-2.25) / -2.25) * 0.0) - (((0.1 + float32((3.14159 * 3.14159))) / 1e10) / 3.14159)) + 1e10)) * -2.25))), F1: "foo bar", F2: 4294967295}
            _ = v23
        case -876:
            if true {
                for fi24 := 0; fi24 < 15; fi24++ {
                    if true {
                        fmt.Println(v13)
                        var v25 int8 = 1
                        delete(v3, fi24)
                        _ = v25
                    }
                    var v26 int16 = 1000
                    switch fi24 {
                    case 128:
                        v5 = "a"
                    }
                    _ = fi24
                    _ = v26
                }
            }
        }
    } else {
        var v27 uint32 = 4294967295
        v4 = v4
        _ = v27
    }
    v28 := []int{-1, 552}
    v29 := GenS2{F0: 0}
    if (((v5 + ("hello" + (v5 + v5))) + (v5 + ("\n\t" + v5))) <= "foo bar") {
        for fi30 := 0; fi30 < 15; fi30++ {
            if (!(519 != 2)) {
                v1 += 510
                var v31 uint8 = 208
                v2 = (v2 | (v2 & (v1 & (v1 << ((42) & 63)))))
                _ = v31
            } else {
                fmt.Println(v5, v5, v5)
                var v32 uint32 = 65535
                if (7 > 127) {
                    var v33 uint16 = 32767
                    if true {
                        fmt.Println(v13, v4, fi30)
                    }
                    var v34 int16 = 957
                    _ = v33
                    _ = v34
                }
                _ = v32
            }
            v5 = v5
            _ = fi30
        }
        fmt.Println(v13)
    } else {
        var v35 uint = 540
        _ = v35
    }
    v2 = 4294967295
    v36 := GenS3{F0: 0, F1: 0}
    var v37 int64 = 42
    if false {
        var v38 uint32 = 2147483648
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
