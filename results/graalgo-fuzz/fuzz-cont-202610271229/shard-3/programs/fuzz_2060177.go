package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS3 struct { F0 string; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{}
    var v3 int16 = 1
    var v4 int16 = (v3 * v3)
    for fi5 := 0; fi5 < 4; fi5++ {
        if false {
            switch fi5 {
            case -477:
                if len(v2) > 0 {
                    v2[0] = -582
                }
            }
            if false {
                _ = len(v1)
            } else {
                fmt.Println(fi5, fi5, v4)
                v3 += v3
                var v6 int16 = (v3 / ((-57) | 1))
                _ = v6
            }
        }
        _ = fi5
    }
    if (true || false) {
        v2 = append(v2, -283)
        if (4294967295 <= 4294967295) {
            switch 2147483647 {
            case 1:
                switch "a" {
                case "foo bar":
                    if (-1 <= 1) {
                        v4 = (v3 + v3)
                        var v7 uint8 = 255
                        _ = v7
                    }
                case "":
                    for fi8 := 0; fi8 < 3; fi8++ {
                        var v9 string = "\n\t"
                        v1[154] = v9
                        v4 = v4
                        _ = fi8
                        _ = v9
                    }
                }
            }
            v3 -= (^(v4))
        } else {
            delete(v1, 1)
            var v10 uint64 = 373
            if (!true) {
                var v11 bool = false
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi12 := 0; fi12 < 9; fi12++ {
                    v4 = (v4 << ((255) & 63))
                    _ = len(v2)
                    _ = fi12
                }
                _ = v11
            }
            _ = v10
        }
        var v13 uint32 = 815
        _ = v13
    }
    var v14 uint32 = 2
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 uint32 = (v14 / (((v14 >> ((0) & 63))) | 1))
        fmt.Println(v16)
        v4 -= 1000
        _ = fi15
        _ = v16
    }
    var v17 uint16 = 255
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := GenS1{F0: 113}
        switch "" {
        case "":
            var v20 int64 = -9223372036854775808
            _ = v20
        case "foo bar":
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 int32 = 681
                _ = fi21
                _ = v22
            }
        }
        if ("" == "foo bar") {
            v23 := GenS2{F0: (^(v17)), F1: true, F2: 4294967295}
            for fi24 := 0; fi24 < 3; fi24++ {
                v4 += (v4 ^ (v4 << ((0) & 63)))
                _ = fi24
            }
            _ = v23
        } else {
            if true {
                for fi25 := 0; fi25 < 7; fi25++ {
                    v1[v14] = ""
                    _ = fi25
                }
            }
            v1[v14] = ("a" + "hello")
        }
        _ = fi18
        _ = v19
    }
    v26 := v1[(v14 % ((v14) | 1))]
    for fi27 := 0; fi27 < 19; fi27++ {
        switch v26 {
        case "a":
            var v28 uint8 = 128
            _ = v28
        case "hello":
            for fi29 := 0; fi29 < 8; fi29++ {
                v30 := GenS3{F0: v26, F1: 353, F2: 4294967295}
                var v31 int = (fi29 >> ((255) & 63))
                v32 := GenS4{F0: 1, F1: v31}
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        }
        _ = fi27
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v26 = v26
    if true {
        if ((v14 - v14) > (v14 / ((v14) | 1))) {
            var v33 int8 = int8(v14)
            var v34 int16 = (v4 / (((v4 + (v4 % (((-(v3))) | 1)))) | 1))
            _ = v33
            _ = v34
        }
    } else {
        v14 -= v14
        var v35 uint = 595
        v36 := GenS4{F0: -1, F1: 32767}
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        if false {
            var v38 uint32 = v14
            var v39 float32 = float32(3.14159)
            v40 := GenS4{F0: 42, F1: -9223372036854775808}
            _ = v38
            _ = v39
            _ = v40
        } else {
            var v41 string = v26
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v41
        }
        for fi42 := 0; fi42 < 1; fi42++ {
            v2 = append(v2, 76)
            var v43 bool = (1000 == (v3 ^ (v3 | 1)))
            switch fi37 {
            case -9223372036854775808:
                var v44 uint = 688
                _ = v44
            }
            var v45 int32 = 1
            _ = fi42
            _ = v43
            _ = v45
        }
        _ = fi37
    }
    v46 := GenS1{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v46
}
