package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 string; F1 int8; F2 int64 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 int64; F1 int64 }

func main() {
    var v1 int = -2147483648
    var v2 int16 = -1
    var v3 int64 = 9223372036854775807
    var v4 uint8 = 255
    v5 := make(map[uint32]int)
    v6 := []int{}
    if true {
        v3 -= -58
        var v7 string = ""
        if true {
            var v8 int32 = -66
            _ = v8
        }
        _ = v7
    } else {
        v9 := GenS1{F0: -1}
        _ = v9
    }
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v12 uint8 = 127
                var v13 uint8 = (v12 >> ((255) & 63))
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    }
    v5[750] = v1
    v14 := v5[2147483648]
    var v15 int32 = -2147483648
    _ = len(v6)
    delete(v5, 735)
    switch "a" {
    case "foo bar":
        switch (v1 * v14) {
        case 4294967295:
            var v16 int = (v1 & (v14 | 966))
            _ = v16
        case -9223372036854775808:
            v2 -= 1
        }
    case "hello":
        fmt.Println(v3)
    case "\n\t":
        _ = len(v5)
    default:
        if false {
            var v17 uint16 = 0
            for fi18 := 0; fi18 < 9; fi18++ {
                for fi19 := 0; fi19 < 8; fi19++ {
                    v14 += (fi19 % (((v1 >> ((4294967295) & 63))) | 1))
                    v20 := GenS2{F0: 0}
                    _ = fi19
                    _ = v20
                }
                switch "a" {
                case "a":
                    v21 := v5[65535]
                    _ = v21
                case "foo bar":
                    v3 = v3
                }
                _ = fi18
            }
            _ = v17
        } else {
            for fi22 := 0; fi22 < 1; fi22++ {
                for fi23 := 0; fi23 < 7; fi23++ {
                    v5[65535] = v1
                    var v24 uint = 65535
                    _ = len(v5)
                    _ = fi23
                    _ = v24
                }
                v1 += v1
                _ = fi22
            }
            v1 += (v1 - (v14 << ((410) & 63)))
            v25 := GenS3{F0: "hello", F1: -964}
            _ = v25
        }
    }
    switch v1 {
    case 1:
        v3 = (v3 / ((v3) | 1))
    case 2147483647:
        v26 := GenS4{F0: "foo bar", F1: 100, F2: (v3 << ((18446744073709551615) & 63))}
        _ = v26
    case -324:
        v2 = int16(v3)
    default:
        var v27 uint = 12
        _ = v27
    }
    fmt.Println(v1, v14, v2)
    for fi28 := 0; fi28 < 5; fi28++ {
        if false {
            if (false || (-32768 < -146)) {
                for fi29 := 0; fi29 < 15; fi29++ {
                    switch (^(v1)) {
                    case -1:
                        if false {
                            switch ("foo bar" + "hello") {
                            case "a":
                                v30 := GenS5{F0: 945, F1: 2147483648}
                                _ = v30
                            }
                            var v31 int32 = 42
                            _ = v31
                        }
                    }
                    v4 = 1
                    fmt.Println(v4)
                    _ = fi29
                }
                v4 += 2
            }
            delete(v5, 65535)
            if (-128 > 127) {
                v32 := GenS5{F0: 516, F1: 189}
                _ = v32
            } else {
                var v33 float32 = 0.0
                v14 -= (fi28 - v1)
                fmt.Println(v15, v1, v14)
                _ = v33
            }
        } else {
            if true {
                switch int(v3) {
                case 128:
                    v14 = -39
                }
            }
        }
        v34 := GenS6{F0: (v3 - v3), F1: (-(v3))}
        v15 = v15
        _ = fi28
        _ = v34
    }
    var v35 float64 = (3.14159 + (3.14159 + 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
