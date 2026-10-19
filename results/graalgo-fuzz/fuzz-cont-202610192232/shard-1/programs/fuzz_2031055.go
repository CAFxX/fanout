package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 uint16; F1 string; F2 int32 }
type GenS4 struct { F0 uint16; F1 int32; F2 int }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int8; F1 int32; F2 uint8 }
type GenS7 struct { F0 bool; F1 int32; F2 uint64 }
type GenS8 struct { F0 int; F1 bool; F2 uint8 }
type GenS9 struct { F0 int8 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    var v4 int = 32767
    v5 := make(map[string]uint32)
    v6 := []int{499, (v4 & 1), (v4 + (v4 - v4))}
    if false {
        if true {
            v7 := GenS1{F0: false, F1: "", F2: 65535}
            var v8 int8 = 20
            _ = v7
            _ = v8
        } else {
            v4 -= (v4 / ((v4) | 1))
            var v9 uint32 = v1
            fmt.Println(v1, v4)
            _ = v9
        }
    }
    v10 := GenS2{F0: "", F1: false, F2: -968}
    if len(v6) > 0 {
        v6[0] = v4
    }
    fmt.Println(v1)
    v11 := v2[("" + "a")]
    v12 := GenS3{F0: 32767, F1: v11, F2: 2147483647}
    v1 += 2147483647
    if false {
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := GenS4{F0: 65535, F1: 1, F2: 9223372036854775807}
            v1 += 410
            _ = fi13
            _ = v14
        }
        var v15 uint = 627
        _ = v15
    } else {
        v1 = v1
        v1 += v1
    }
    if true {
        switch "hello" {
        case "foo bar":
            v16 := GenS5{F0: 127, F1: 324}
            _ = v16
        default:
            v4 = v4
        }
        v5[(v11 + v11)] = (v1 & 1)
    } else {
        fmt.Println(v1, v4, v11)
        switch (v11 + v11) {
        case "a":
            for fi17 := 0; fi17 < 9; fi17++ {
                v4 -= (fi17 * (fi17 * v4))
                if true {
                    for fi18 := 0; fi18 < 7; fi18++ {
                        v19 := GenS6{F0: 127, F1: -1, F2: 128}
                        var v20 uint64 = 42
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    v1 = 293
                }
                _ = fi17
            }
        }
        var v21 int16 = 42
        _ = v21
    }
    v22 := GenS5{F0: 128, F1: 339}
    for fi23 := 0; fi23 < 16; fi23++ {
        switch v4 {
        case 42:
            switch v4 {
            case 0:
                v4 -= -1
            }
        default:
            for fi24 := 0; fi24 < 8; fi24++ {
                fmt.Println(v11, v4)
                v1 = v1
                if ((v1 ^ (v1 - v1)) > (v1 % (((v1 / ((v1) | 1))) | 1))) {
                    v25 := v5[v11]
                    var v26 bool = false
                    _ = v25
                    _ = v26
                }
                _ = fi24
            }
        }
        if (1e10 == -2.25) {
            v27 := GenS7{F0: true, F1: 1, F2: 0}
            var v28 string = "hello"
            v29 := GenS8{F0: (v4 | (v4 % ((32768) | 1))), F1: (!false), F2: 225}
            _ = v27
            _ = v28
            _ = v29
        } else {
            var v30 uint = 48
            _ = v30
        }
        var v31 float32 = -2.25
        _ = fi23
        _ = v31
    }
    v1 = 42
    if false {
        switch v4 {
        case 2147483647:
            if (false && true) {
                var v32 int64 = 422
                _ = len(v5)
                _ = v32
            }
        default:
            for fi33 := 0; fi33 < 1; fi33++ {
                for fi34 := 0; fi34 < 3; fi34++ {
                    fmt.Println(v4)
                    v35 := GenS2{F0: "foo bar", F1: false, F2: 42}
                    if (65535 >= 1) {
                        fmt.Println(v11)
                        switch (v4 << ((964) & 63)) {
                        case -1:
                            var v36 uint32 = (v1 - 4294967295)
                            _ = v36
                        }
                    }
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
        }
        if true {
            v11 = (v11 + (v11 + "a"))
            var v37 uint64 = 0
            var v38 uint16 = 840
            _ = v37
            _ = v38
        }
        v39 := GenS7{F0: true, F1: 300, F2: 4294967295}
        _ = v39
    }
    var v40 uint16 = 65535
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v41 := GenS9{F0: 100}
    if false {
        v4 = 965
    } else {
        var v42 int16 = -80
        v43 := v3[2147483647]
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
