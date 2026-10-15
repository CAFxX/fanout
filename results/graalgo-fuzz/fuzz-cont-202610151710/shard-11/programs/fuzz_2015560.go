package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 bool }
type GenS3 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int16; F1 int64; F2 uint16 }
func genFn0(p0 uint16, p1 int16, p2 uint8) bool {
    switch ("hello" + ("hello" + (("foo bar" + "") + "\n\t"))) {
    case "a":
        if false {
            switch 42 {
            case -756:
                p2 += p2
            default:
                if true {
                    if (((((float32(1e10) * 0.0) + (0.1 - 0.1)) / -2.25) / (3.14159 * (0.1 - 1.5))) == (0.1 - -2.25)) {
                        v6 := []int{-2147483648, 470, 127}
                        var v7 uint16 = p0
                        var v8 uint64 = 59
                        _ = v6
                        _ = v7
                        _ = v8
                    }
                    for fi9 := 0; fi9 < 20; fi9++ {
                        v10 := []int{}
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        _ = fi9
                        _ = v10
                    }
                } else {
                    v11 := []int{-9223372036854775808, -339, 947}
                    fmt.Println(p2, p1, p0)
                    _ = v11
                }
            }
        }
    case "\n\t":
        var v12 uint64 = 18446744073709551615
        _ = v12
    case "foo bar":
        var v13 int = -9223372036854775808
        _ = v13
    default:
        if true {
            var v14 bool = false
            var v15 string = "a"
            _ = v14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(fi16)
        switch 453 {
        case -1:
            p0 -= 823
        case 919:
            fmt.Println(p2, fi16)
        case 2147483647:
            if (42 < 2147483647) {
                v17 := GenS1{F0: 0}
                v18 := []int{fi16, fi16, (fi16 / ((fi16) | 1))}
                _ = v17
                _ = v18
            }
        }
        _ = fi16
    }
    var v19 uint32 = 855
    switch 32767 {
    case 9223372036854775807:
        p2 = p2
    default:
        if (true || (p0 >= p0)) {
            var v20 uint8 = (p2 >> ((42) & 63))
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        fmt.Println(fi21, p1)
        for fi22 := 0; fi22 < 5; fi22++ {
            p2 -= p2
            v23 := []int{}
            _ = fi22
            _ = v23
        }
        fmt.Println(p0, p0)
        _ = fi21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return (!((v19 < (v19 * 228)) || true))
}
func genFn1() bool {
    for fi24 := 0; fi24 < 17; fi24++ {
        switch (("" + (("foo bar" + "\n\t") + "\n\t")) + ("\n\t" + "")) {
        case "foo bar":
            var v25 int8 = 0
            _ = v25
        case "hello":
            v26 := GenS2{F0: 0, F1: -806, F2: (true && (-48 > -96))}
            _ = v26
        case "a":
            var v27 uint32 = 2
            _ = v27
        default:
            var v28 uint8 = 128
            _ = v28
        }
        if (42 == 1) {
            var v29 uint16 = 168
            v29 = 891
            v30 := []int{(fi24 & (fi24 | -9223372036854775808)), (fi24 + fi24)}
            _ = v29
            _ = v30
        } else {
            fmt.Println(fi24)
            v31 := []int{}
            _ = v31
        }
        fmt.Println(fi24)
        _ = fi24
    }
    var v32 int8 = 0
    switch 898 {
    case -478:
        v33 := []int{}
        _ = v33
    }
    v34 := GenS1{F0: 430}
    _ = v32
    _ = v34
    return false
}

func main() {
    var v1 int16 = -812
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]uint32)
    var v5 float32 = 3.14159
    var v35 int8 = -1
    var v36 int64 = 0
    v35 = v35
    var v37 uint64 = 170
    switch "foo bar" {
    case "\n\t":
        v38 := []int{1, -993}
        _ = v38
    case "":
        v37 += 4294967295
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        var v40 float64 = 0.0
        for fi41 := 0; fi41 < 2; fi41++ {
            switch (fi41 ^ fi41) {
            case 155:
                v42 := GenS3{F0: 127, F1: 304, F2: 78}
                _ = v42
            case 280:
                var v43 int64 = (^(v36))
                _ = v43
            case 251:
                v44 := GenS4{F0: v1, F1: (v36 - 9223372036854775807), F2: 65535}
                _ = v44
            }
            _ = fi41
        }
        v45 := GenS1{F0: v37}
        _ = fi39
        _ = v40
        _ = v45
    }
    v36 += 9223372036854775807
    v46 := v4[618]
    for fi47 := 0; fi47 < 7; fi47++ {
        var v48 int32 = -943
        v49 := GenS2{F0: (v35 * 100), F1: (fi47 & fi47), F2: true}
        _ = fi47
        _ = v48
        _ = v49
    }
    v35 += -128
    v36 = (-(v36))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v46
    _ = v5
}
