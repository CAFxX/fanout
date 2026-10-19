package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint8; F1 int64 }
func genFn0(p0 int32) int {
    var v4 string = "hello"
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{fi5, fi5}
        _ = fi5
        _ = v6
    }
    p0 += 2147483647
    var v7 uint8 = 127
    v4 = v4
    _ = p0
    _ = v4
    _ = v7
    return 32768
}
func genFn1(p0 uint32, p1 int8, p2 uint32) int8 {
    fmt.Println(p2, p0, p1)
    v8 := GenS1{F0: 4294967295, F1: 0}
    if true {
        p2 = 2147483647
        if true {
            var v9 uint8 = 0
            _ = v9
        }
    }
    p2 = (p2 & 65535)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    var v3 uint8 = 1
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(v3)
        _ = fi10
    }
    var v11 int16 = -32768
    v12 := GenS2{F0: 1, F1: 65535}
    v11 += (-(v11))
    v11 += v11
    v2[("\n\t" + (("" + "foo bar") + ((("\n\t" + ("a" + "")) + "hello") + ("\n\t" + "\n\t"))))] = -658
    switch 4294967295 {
    case 0:
        v11 = (v11 - (v11 | -32768))
    case 401:
        switch ("\n\t" + ("foo bar" + "")) {
        case "":
            fmt.Println(v11, v3)
        }
    case -198:
        v3 = v3
    default:
        if (true && (910 > 42)) {
            var v13 uint32 = 4294967295
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS2{F0: 468, F1: 18446744073709551615}
        var v16 string = "a"
        v17 := []int{int(v11), fi14}
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    if (true || (2147483647 != 9223372036854775807)) {
        v18 := GenS3{F0: v3, F1: -985}
        _ = v18
    }
    fmt.Println(v11, v11)
    v11 += (v11 & (v11 << ((0) & 63)))
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := []int{fi19, fi19, (fi19 | 434)}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            if (!false) {
                var v23 uint8 = (v3 & (^(v3)))
                var v24 float64 = (((3.14159 * float64((((1e10 / -2.25) * -2.25) + 3.14159))) / (1.5 / (3.14159 + 0.1))) - (1e10 * (3.14159 - ((((0.0 * 0.1) + (1.5 + 3.14159)) / ((1.5 - 1e10) - -2.25)) / -2.25))))
                var v25 uint8 = v3
                _ = v23
                _ = v24
                _ = v25
            } else {
                switch "" {
                case "\n\t":
                    var v26 bool = (2 < 2)
                    _ = v26
                }
            }
            _ = len(v2)
            _ = fi22
        }
        _ = fi21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
