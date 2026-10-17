package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 uint32 }
func genFn0(p0 int64) int8 {
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: 65535, F1: "\n\t", F2: fi5}
        if (42 == -1) {
            var v7 uint = 2
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    if true {
        var v8 int16 = 0
        for fi9 := 0; fi9 < 7; fi9++ {
            p0 = p0
            v8 += (v8 ^ v8)
            _ = fi9
        }
        _ = v8
    }
    if (1 > 782) {
        fmt.Println(p0)
    }
    var v10 bool = (!false)
    fmt.Println(v10, p0)
    _ = p0
    _ = v10
    return 100
}

func main() {
    var v1 int32 = -1
    var v2 uint16 = 255
    var v3 int32 = (v1 ^ v1)
    v4 := []uint32{65535}
    var v11 int32 = (v3 / (((v3 & v1)) | 1))
    for fi12 := 0; fi12 < 12; fi12++ {
        v1 = 813
        v2 = 1
        _ = fi12
    }
    v13 := GenS1{F0: 4294967295, F1: "a", F2: 282}
    if len(v4) > 0 {
        v4[0] = 172
    }
    if false {
        if len(v4) > 0 {
            v4[0] = 2147483648
        }
    }
    v11 = v11
    fmt.Println(v3)
    fmt.Println(v11, v2, v1)
    if (!(!(103 == uint64(v2)))) {
        fmt.Println(v2)
        if (!false) {
            var v14 int16 = 1
            if (-946 <= -2147483648) {
                fmt.Println(v1, v1, v14)
                v15 := GenS2{F0: ("" + "foo bar")}
                if false {
                    v16 := GenS3{F0: 873}
                    var v17 int = 32767
                    var v18 int32 = v11
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = v15
            } else {
                fmt.Println(v2)
                fmt.Println(v2)
            }
            _ = v14
        } else {
            v19 := GenS4{F0: (-412 <= 32767), F1: 189}
            _ = v19
        }
    }
    var v20 int16 = -98
    v11 -= v11
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 int = fi21
        for fi23 := 0; fi23 < 8; fi23++ {
            v4 = append(v4, 2147483648)
            _ = fi23
        }
        v22 = 692
        _ = fi21
        _ = v22
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 uint16 = (v2 & (^(v2)))
            switch -222 {
            case 1:
                if false {
                    var v27 uint = 42
                    fmt.Println(v27, v26)
                    var v28 string = ("foo bar" + "hello")
                    var v29 string = "foo bar"
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    v26 -= 65535
                }
            case -9223372036854775808:
                var v30 float32 = 1.5
                _ = v30
            }
            v1 = v11
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        _ = len(v4)
        if (18446744073709551615 == 18446744073709551615) {
            fmt.Println(v20, v20, v20)
        }
        var v32 uint64 = 387
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
