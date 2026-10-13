package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 int; F1 uint32; F2 bool }
type GenS3 struct { F0 string; F1 int; F2 int32 }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]string)
    var v3 bool = true
    var v4 int8 = 1
    var v5 uint8 = 1
    v6 := make(map[int]uint32)
    switch "foo bar" {
    case "foo bar":
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 127, F1: 18446744073709551615, F2: 1}
            v9 := GenS1{F0: -1, F1: 18446744073709551615, F2: 403}
            v10 := []int{(fi7 & 689), fi7}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    case "":
        switch 2147483647 {
        case 128:
            fmt.Println(v3, v1, v4)
        }
    }
    var v11 uint16 = 2
    if (((^(v1)) < v1) && false) {
        for fi12 := 0; fi12 < 17; fi12++ {
            v11 = 42
            _ = fi12
        }
    } else {
        fmt.Println(v11, v1)
        if false {
            v1 += (v1 + (v1 - v1))
            var v13 uint8 = v5
            var v14 int8 = 0
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        switch "a" {
        case "a":
            fmt.Println(v11, v3)
        case "foo bar":
            v4 += (v4 << ((393) & 63))
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            v3 = (1 <= -851)
            v4 = (v4 & v4)
            _ = fi16
        }
        _ = fi15
    }
    _ = len(v2)
    var v17 float32 = 0.0
    v11 += uint16(v5)
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 int = (fi18 % ((int(v4)) | 1))
        var v20 uint16 = (v11 | 291)
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 bool = v3
        var v23 int = (fi21 % ((fi21) | 1))
        _ = fi21
        _ = v22
        _ = v23
    }
    if v3 {
        for fi24 := 0; fi24 < 10; fi24++ {
            if (((v5 / ((v5) | 1)) == 128) || (42 > 140)) {
                switch "hello" {
                case "hello":
                    for fi25 := 0; fi25 < 2; fi25++ {
                        var v26 uint32 = uint32(v11)
                        v26 -= (^(v1))
                        _ = fi25
                        _ = v26
                    }
                case "foo bar":
                    var v27 int8 = (v4 % ((v4) | 1))
                    _ = v27
                case "":
                    v28 := GenS2{F0: fi24, F1: uint32(v5), F2: v3}
                    _ = v28
                }
                switch fi24 {
                case 2147483647:
                    v29 := GenS2{F0: fi24, F1: v1, F2: (-32768 < 637)}
                    _ = v29
                case -763:
                    v30 := GenS1{F0: (-(v4)), F1: 549, F2: 1}
                    _ = v30
                default:
                    v1 = (v1 | (v1 * (v1 - 758)))
                }
            } else {
                v4 = (v4 & (v4 >> ((347) & 63)))
                var v31 string = "hello"
                _ = v31
            }
            v32 := GenS2{F0: fi24, F1: (^(v1)), F2: true}
            _ = fi24
            _ = v32
        }
    } else {
        var v33 int8 = v4
        _ = v33
    }
    fmt.Println(v17, v5, v3)
    if v3 {
        if v3 {
            v11 -= (v11 | (v11 / ((v11) | 1)))
            v34 := v2[32768]
            _ = v34
        }
    } else {
        if v3 {
            v4 = (-(v4))
            if ((1 >= 2) || (false || (!(((0.0 <= float32(1.5)) || false) && v3)))) {
                var v35 int8 = v4
                v36 := []int{-77, -9223372036854775808, -93, -98}
                _ = v35
                _ = v36
            }
        } else {
            v11 = v11
            if (!v3) {
                v3 = (-32768 >= 1000)
            } else {
                for fi37 := 0; fi37 < 2; fi37++ {
                    v38 := GenS3{F0: "foo bar", F1: fi37, F2: 913}
                    var v39 uint64 = 6
                    v4 = v4
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            v40 := GenS2{F0: 2147483647, F1: (v1 >> ((4294967295) & 63)), F2: ((((((1.5 - v17) / v17) - (v17 + v17)) / 1.0) / float32(1.5)) >= (1.5 - v17))}
            _ = v40
        }
    }
    if (-633 <= 1) {
        v41 := []int{967, -9223372036854775808, 2147483647, 775}
        _ = v41
    } else {
        v5 -= (v5 >> ((834) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
