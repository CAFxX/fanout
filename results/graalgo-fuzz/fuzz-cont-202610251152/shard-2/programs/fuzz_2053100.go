package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 int32; F1 bool; F2 uint64 }
type GenS4 struct { F0 int64; F1 int8; F2 uint16 }
func genFn0(p0 int32) int8 {
    p0 += (^(p0))
    p0 = 1
    v4 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = v4
    return -95
}
func genFn1(p0 string, p1 int64) bool {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 bool = false
        fmt.Println(fi5, fi5)
        v7 := []int{}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int8 = 0
    for fi9 := 0; fi9 < 18; fi9++ {
        if ((p0 + p0) >= ("" + (p0 + p0))) {
            p1 = (p1 * (p1 >> ((387) & 63)))
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return false
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -275
    var v3 string = "\n\t"
    for fi10 := 0; fi10 < 5; fi10++ {
        if (!false) {
            v11 := GenS1{F0: 0, F1: 658, F2: 143}
            for fi12 := 0; fi12 < 3; fi12++ {
                v3 = ""
                v2 += (v2 ^ v2)
                _ = fi12
            }
            _ = v11
        } else {
            v13 := GenS1{F0: 0, F1: 4294967295, F2: 42}
            v2 += v2
            v2 += (^(v2))
            _ = v13
        }
        v14 := GenS1{F0: 4294967295, F1: 0, F2: -1}
        _ = fi10
        _ = v14
    }
    v2 = (v2 >> ((65535) & 63))
    v3 = "a"
    v15 := GenS1{F0: 4294967295, F1: 2147483647, F2: -1}
    fmt.Println(v3)
    v3 = v3
    switch 4294967295 {
    case 187:
        if false {
            fmt.Println(v2, v2)
            if false {
                if (2 <= 0) {
                    var v16 float32 = 1e10
                    _ = v16
                } else {
                    fmt.Println(v3)
                }
                v3 = v3
                v17 := GenS2{F0: 42, F1: 65535, F2: false}
                _ = v17
            }
            switch "\n\t" {
            case "a":
                var v18 uint16 = 375
                _ = v18
            case "foo bar":
                v2 -= v2
            default:
                switch "foo bar" {
                case "\n\t":
                    if true {
                        fmt.Println(v3, v3)
                        v2 -= -32768
                    }
                default:
                    for fi19 := 0; fi19 < 2; fi19++ {
                        var v20 int16 = 300
                        v21 := []int{}
                        v22 := GenS1{F0: 18446744073709551615, F1: 65535, F2: int64(v2)}
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                }
            }
        } else {
            var v23 float32 = 1e10
            v24 := v1[-1]
            var v25 bool = ((float64(v23) + 0.1) == 1e10)
            _ = v23
            _ = v24
            _ = v25
        }
    case 688:
        v26 := GenS3{F0: int32(v2), F1: (((float64(3.14159) - ((3.14159 + 0.0) / float64(1e10))) >= (float64(3.14159) - 0.1)) || true), F2: 42}
        _ = v26
    case 128:
        v3 = (v3 + v3)
    }
    v2 = v2
    var v27 uint64 = 87
    v27 -= v27
    var v28 string = v3
    v27 = v27
    fmt.Println(v27, v2, v2)
    v27 -= v27
    switch (((((v3 + "") + ("a" + (v3 + "hello"))) + v28) + "hello") + "foo bar") {
    case "hello":
        if (!false) {
            v29 := []int{362, -9223372036854775808, 185}
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
            for fi30 := 0; fi30 < 11; fi30++ {
                for fi31 := 0; fi31 < 20; fi31++ {
                    v3 = "foo bar"
                    v28 = ""
                    _ = fi31
                }
                _ = fi30
            }
            _ = v29
        }
    case "":
        var v32 float64 = -2.25
        _ = v32
    case "a":
        delete(v1, 1)
    default:
        fmt.Println(v2, v28)
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        switch ("" + (((v3 + ("foo bar" + v28)) + "foo bar") + "foo bar")) {
        case "a":
            if (1 > 364) {
                var v34 int = (fi33 / ((fi33) | 1))
                _ = len(v1)
                _ = v34
            } else {
                v35 := GenS3{F0: 2147483647, F1: true, F2: v27}
                v36 := GenS3{F0: 2147483647, F1: false, F2: v27}
                _ = v35
                _ = v36
            }
        }
        _ = fi33
    }
    var v37 uint8 = 2
    v38 := GenS4{F0: -240, F1: 1, F2: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v38
}
