package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 string }
func genFn0(p0 string, p1 int32, p2 int) uint8 {
    fmt.Println(p0, p2)
    v7 := GenS1{F0: 195, F1: false, F2: 42}
    fmt.Println(p2)
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(p0, p2)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 string, p1 bool) uint16 {
    v9 := GenS2{F0: (255 <= 2)}
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := []int{(fi10 & (fi10 % ((fi10) | 1))), fi10, fi10}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "foo bar"
    v3 := []int{}
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    var v6 bool = (!true)
    delete(v5, -2147483648)
    if v6 {
        v2 = "hello"
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            if (-128 <= -1) {
                var v13 uint8 = 0
                var v14 int64 = 1
                _ = v13
                _ = v14
            } else {
                v4[(^(fi12))] = 65535
                if (("" + ((v2 + (v2 + (v2 + v2))) + (v2 + ("foo bar" + (v2 + v2))))) <= "a") {
                    v4[506] = 0
                } else {
                    v2 = ((("" + ("a" + "")) + "a") + v2)
                }
            }
            for fi15 := 0; fi15 < 14; fi15++ {
                v2 = ""
                var v16 int64 = 0
                var v17 uint64 = 2
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi12
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := GenS2{F0: true}
            fmt.Println(v2, v2)
            if true {
                var v20 int32 = 0
                v2 = v2
                _ = v20
            } else {
                if (2147483648 <= 2147483648) {
                    v2 = v2
                    switch v2 {
                    case "":
                        var v21 int32 = 2147483647
                        _ = v21
                    }
                } else {
                    var v22 int8 = 1
                    v23 := v1[2]
                    _ = v22
                    _ = v23
                }
                switch (v2 + "hello") {
                case "foo bar":
                    v6 = (v2 != (v2 + "hello"))
                default:
                    v2 = (v2 + v2)
                }
            }
            _ = fi18
            _ = v19
        }
        if (!true) {
            v5[1] = v2
        }
    }
    var v24 int64 = 17
    v2 = ""
    var v25 float64 = (3.14159 + 1.5)
    var v26 uint = 354
    if (!(65535 == 255)) {
        if len(v3) > 0 {
            v3[0] = -632
        }
    }
    if (127 >= -9223372036854775808) {
        if len(v3) > 0 {
            v3[0] = 128
        }
        for fi27 := 0; fi27 < 16; fi27++ {
            fmt.Println(v2, fi27)
            v26 += (v26 >> (((^(v26))) & 63))
            _ = fi27
        }
        switch -2147483648 {
        case 9223372036854775807:
            var v28 int8 = -128
            _ = v28
        case 2147483647:
            switch 32768 {
            case -770:
                if true {
                    v4[32767] = 618
                }
            case 1:
                v29 := v5[33]
                _ = v29
            }
        case 4294967295:
            fmt.Println(v6)
        default:
            v30 := GenS1{F0: 1, F1: (!false), F2: v26}
            _ = v30
        }
    }
    if (!(((v25 / v25) * 1e10) == 0.1)) {
        for fi31 := 0; fi31 < 10; fi31++ {
            fmt.Println(v26, v26, v24)
            _ = fi31
        }
        for fi32 := 0; fi32 < 1; fi32++ {
            v6 = true
            for fi33 := 0; fi33 < 11; fi33++ {
                v6 = v6
                _ = fi33
            }
            var v34 uint64 = 89
            _ = fi32
            _ = v34
        }
    }
    if v6 {
        v35 := GenS3{F0: 283, F1: v2}
        _ = v35
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        if len(v3) > 0 {
            v3[0] = fi36
        }
        fmt.Println(v25)
        _ = fi36
    }
    _ = len(v4)
    switch 9223372036854775807 {
    case -1:
        v25 = v25
    }
    if (1e10 != (0.0 / 1e10)) {
        v24 = -9223372036854775808
    } else {
        _ = len(v4)
    }
    v2 = v2
    v6 = (799 == -500)
    delete(v4, -1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
