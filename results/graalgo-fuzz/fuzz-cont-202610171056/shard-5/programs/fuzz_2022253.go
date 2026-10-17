package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint; F2 int8 }
type GenS5 struct { F0 bool }
func genFn0(p0 int32, p1 int32, p2 uint8) bool {
    switch ("foo bar" + ("hello" + (("a" + "hello") + "hello"))) {
    case "a":
        var v7 float64 = -2.25
        _ = v7
    case "":
        var v8 uint32 = 57
        _ = v8
    }
    p0 -= p0
    var v9 int8 = 0
    if (!false) {
        var v10 uint = 42
        _ = v10
    }
    switch -657 {
    case -164:
        var v11 uint8 = (p2 / ((p2) | 1))
        _ = v11
    case 547:
        var v12 float32 = (1.5 / 1e10)
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (100 < -1)
}
func genFn1(p0 uint16, p1 int, p2 int) bool {
    v13 := GenS1{F0: 737, F1: 100}
    switch "a" {
    case "hello":
        p0 = (p0 - p0)
    case "":
        if ((1 < -32768) || true) {
            if true {
                var v14 string = "foo bar"
                _ = v14
            }
            v15 := []int{(p2 % ((32768) | 1)), (p2 & 737)}
            p0 += p0
            _ = v15
        } else {
            var v16 uint16 = (p0 - (p0 << ((255) & 63)))
            p1 = (p2 & p1)
            fmt.Println(p2, p1, p0)
            _ = v16
        }
    case "a":
        switch (("\n\t" + "\n\t") + ("" + ("\n\t" + "foo bar"))) {
        case "foo bar":
            fmt.Println(p1, p2, p0)
        case "hello":
            fmt.Println(p1, p2)
        case "\n\t":
            fmt.Println(p2)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return false
}

func main() {
    var v1 string = (("" + "hello") + "foo bar")
    var v2 bool = true
    v3 := make(map[uint32]uint32)
    v4 := []uint32{347, 519, 2}
    v5 := []int{461}
    v6 := make(map[string]string)
    var v17 uint32 = 340
    var v18 float64 = 1.5
    switch 1 {
    case 127:
        var v19 uint64 = 42
        _ = v19
    case 128:
        delete(v6, (v1 + (v1 + v1)))
    case -789:
        var v20 uint64 = 18446744073709551615
        _ = v20
    }
    switch 2147483647 {
    case 42:
        for fi21 := 0; fi21 < 4; fi21++ {
            if (!(315 >= 757)) {
                switch (v1 + v1) {
                case "":
                    v22 := GenS1{F0: 949, F1: 127}
                    _ = v22
                }
                v5 = append(v5, (^(fi21)))
            } else {
                v23 := GenS1{F0: 646, F1: -1}
                _ = v23
            }
            for fi24 := 0; fi24 < 3; fi24++ {
                v25 := GenS2{F0: (v17 >> ((931) & 63)), F1: 18446744073709551615, F2: 116}
                v26 := GenS3{F0: 122}
                _ = fi24
                _ = v25
                _ = v26
            }
            switch v1 {
            case "foo bar":
                var v27 uint16 = 42
                _ = v27
            case "a":
                v18 = (3.14159 * v18)
            case "hello":
                v28 := GenS4{F0: v2, F1: 255, F2: 2}
                _ = v28
            default:
                v1 = v1
            }
            _ = fi21
        }
    case -9223372036854775808:
        var v29 float32 = 0.1
        _ = v29
    case 2147483647:
        v30 := GenS4{F0: (2147483647 < -612), F1: 570, F2: 27}
        _ = v30
    default:
        var v31 int8 = -114
        _ = v31
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        for fi33 := 0; fi33 < 17; fi33++ {
            switch fi32 {
            case -2147483648:
                v17 = (v17 >> ((4294967295) & 63))
            case 1:
                switch (-(fi32)) {
                case 32768:
                    if true {
                        v18 = (v18 / v18)
                    } else {
                        v18 -= (((1.5 * float64(float32((v18 / v18)))) - (v18 / 1e10)) * (3.14159 / v18))
                    }
                }
            case 2147483647:
                var v34 uint8 = 127
                _ = v34
            default:
                if (0.1 < (1.5 + 1.5)) {
                    fmt.Println(v18, fi32)
                    fmt.Println(fi33, v1)
                    v18 = 3.14159
                }
            }
            v18 -= 1e10
            v18 = (((0.1 / v18) + (float64(-2.25) / v18)) * ((v18 - 0.0) - float64(((-2.25 + 1e10) / 1e10))))
            v35 := GenS3{F0: 0}
            switch "hello" {
            case "":
                _ = len(v6)
            case "a":
                v36 := GenS3{F0: 64}
                _ = v36
            default:
                var v37 bool = ("hello" != v1)
                var v38 int8 = int8(fi32)
                _ = v37
                _ = v38
            }
            _ = fi33
            _ = v35
        }
        _ = fi32
    }
    if v2 {
        var v39 string = v1
        fmt.Println(v2)
        if false {
            for fi40 := 0; fi40 < 4; fi40++ {
                var v41 bool = true
                _ = fi40
                _ = v41
            }
            for fi42 := 0; fi42 < 10; fi42++ {
                v6[(("a" + v1) + v1)] = "a"
                _ = fi42
            }
            if len(v4) > 0 {
                v4[0] = (v17 ^ v17)
            }
        }
        _ = v39
    } else {
        v17 = (v17 << ((255) & 63))
    }
    if v2 {
        for fi43 := 0; fi43 < 4; fi43++ {
            v2 = false
            var v44 uint32 = v17
            var v45 int64 = -409
            _ = fi43
            _ = v44
            _ = v45
        }
        if v2 {
            v46 := GenS5{F0: (1 == -9223372036854775808)}
            var v47 int8 = 0
            switch 42 {
            case -9223372036854775808:
                var v48 uint8 = 127
                _ = v48
            case 904:
                v49 := GenS1{F0: 0, F1: v47}
                _ = v49
            default:
                if v2 {
                    delete(v3, (^(v17)))
                }
            }
            _ = v46
            _ = v47
        }
        var v50 uint16 = 32767
        _ = v50
    } else {
        var v51 uint16 = 32767
        _ = v51
    }
    v18 -= v18
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
