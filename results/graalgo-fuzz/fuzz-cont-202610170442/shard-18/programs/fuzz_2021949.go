package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint32 }
type GenS2 struct { F0 uint; F1 int64; F2 int32 }
type GenS3 struct { F0 bool; F1 int8 }
func genFn0(p0 bool, p1 int8) uint64 {
    var v4 uint32 = 173
    if p0 {
        for fi5 := 0; fi5 < 10; fi5++ {
            p0 = (!false)
            _ = fi5
        }
        v4 = v4
    } else {
        v6 := []int{32767}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 778
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -547
    var v3 uint32 = 2147483648
    switch "hello" {
    case "\n\t":
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := GenS1{F0: ("foo bar" + ""), F1: ("foo bar" + "a"), F2: (v3 % ((v3) | 1))}
            var v9 int8 = 21
            var v10 float32 = (((float32(-2.25) + 1e10) / 3.14159) + (0.0 - ((((((-2.25 - 3.14159) - (3.14159 * 1.5)) * -2.25) * (((0.0 / 1.5) - float32(0.1)) + (1.5 - (1.5 / 1e10)))) * ((float32((1e10 + 0.1)) - float32((0.1 / -2.25))) + (1.5 / 0.1))) / -2.25)))
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    case "foo bar":
        v11 := GenS1{F0: "a", F1: "", F2: v3}
        _ = v11
    case "hello":
        if (false || (128 > 42)) {
            v3 = v3
        } else {
            if false {
                var v12 int32 = 0
                v12 += v12
                v12 = int32(v2)
                _ = v12
            }
            v13 := GenS2{F0: 42, F1: 0, F2: 1}
            _ = v13
        }
    }
    fmt.Println(v3, v2)
    switch "\n\t" {
    case "":
        fmt.Println(v3, v2)
    case "foo bar":
        v1[981] = 32768
    }
    v2 -= 1
    v14 := GenS1{F0: "hello", F1: ("foo bar" + ("" + "a")), F2: v3}
    if false {
        if (true || false) {
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 int64 = 885
                _ = fi15
                _ = v16
            }
        } else {
            var v17 uint = 40
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 uint8 = 1
        v3 = 419
        if true {
            var v20 int64 = -601
            for fi21 := 0; fi21 < 20; fi21++ {
                v22 := GenS1{F0: "hello", F1: "", F2: 463}
                _ = fi21
                _ = v22
            }
            var v23 uint64 = 18446744073709551615
            _ = v20
            _ = v23
        }
        _ = fi18
        _ = v19
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        v25 := GenS2{F0: 841, F1: 9223372036854775807, F2: 2147483647}
        v26 := []int{(fi24 - (fi24 * (fi24 & fi24))), fi24, 1}
        v3 = (v3 - 1)
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println(v2, v2)
    v2 += (v2 | (v2 << ((1) & 63)))
    for fi27 := 0; fi27 < 17; fi27++ {
        fmt.Println(fi27)
        switch (fi27 & 32767) {
        case 42:
            var v28 string = "a"
            _ = v28
        case 2147483647:
            switch "hello" {
            case "\n\t":
                v2 += (v2 ^ v2)
            case "a":
                v2 -= (v2 >> ((18446744073709551615) & 63))
            case "hello":
                v29 := GenS1{F0: ("a" + ""), F1: "a", F2: 65535}
                _ = v29
            }
        default:
            v2 -= 0
        }
        var v30 uint32 = v3
        _ = fi27
        _ = v30
    }
    if (!false) {
        switch "foo bar" {
        case "\n\t":
            _ = len(v1)
        case "hello":
            delete(v1, -415)
        case "":
            var v31 float64 = 1.5
            _ = v31
        }
        v2 += -238
    }
    v32 := GenS3{F0: true, F1: -29}
    var v33 int8 = 100
    for fi34 := 0; fi34 < 3; fi34++ {
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 float32 = 1e10
            if (1 >= 42) {
                var v37 uint8 = 128
                _ = v37
            }
            v38 := []int{fi34, fi34, (-(fi35)), (fi35 % ((fi34) | 1))}
            _ = fi35
            _ = v36
            _ = v38
        }
        _ = fi34
    }
    var v39 int8 = (v33 | v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v39
}
